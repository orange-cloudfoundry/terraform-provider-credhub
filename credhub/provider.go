package credhub

import (
	"encoding/json"
	"fmt"
	"github.com/cloudfoundry-incubator/credhub-cli/credhub"
	"github.com/cloudfoundry-incubator/credhub-cli/credhub/auth"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	TOKENS_FILENAME = "tf-credhub-tokens.json"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func Provider() terraform.ResourceProvider {

	// The actual provider
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"credhub_server": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CREDHUB_SERVER", ""),
				Description: "Your credhub api url. (Note: https will be enforced)",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CREDHUB_USERNAME", ""),
				Description: "The username of an UAA user credhub.write and credhub.read scopes. (Optional if you use an client_id/client_secret)",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CREDHUB_PASSWORD", ""),
				Description: "The password of an UAA user credhub.write and credhub.read scopes. (Optional if you use an client_id/client_secret)",
			},
			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CREDHUB_CLIENT", ""),
				Default:     "credhub_cli",
				Description: "The client_id of an UAA client credhub.write and credhub.read scopes. (Optional if you use an username/password)",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CREDHUB_SECRET", ""),
				Description: "The client_secret of an UAA client credhub.write and credhub.read scopes. (Optional if you use an username/password)",
			},
			"ca_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CREDHUB_CA_CERT", ""),
				Description: "Trusted CA for API and UAA TLS connections.",
			},
			"skip_ssl_validation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Set to true to skip verification of the API endpoint. Not recommended!",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"credhub_certificate": LoadGenerateResource(&GenerateCertificateResource{}),
			"credhub_password":    LoadGenerateResource(&GeneratePasswordResource{}),
			"credhub_rsa":         LoadGenerateResource(&GenerateRSAResource{}),
			"credhub_ssh":         LoadGenerateResource(&GenerateSSHResource{}),
			"credhub_user":        LoadGenerateResource(&GenerateUserResource{}),
			"credhub_generic":     LoadGenerateResource(&GenericResource{}),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"credhub_generic":     LoadDataSource(&GenericDataSource{}),
			"credhub_value":       LoadDataSource(&ValueDataSource{}),
			"credhub_json":        LoadDataSource(&JsonDataSource{}),
			"credhub_password":    LoadDataSource(&PasswordDataSource{}),
			"credhub_certificate": LoadDataSource(&CertificateDataSource{}),
			"credhub_rsa":         LoadDataSource(&RSADataSource{}),
			"credhub_ssh":         LoadDataSource(&SSHDataSource{}),
			"credhub_user":        LoadDataSource(&UserDataSource{}),
		},

		ConfigureFunc: providerConfigure,
	}
}
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	apiEndpoint := strings.TrimPrefix(d.Get("credhub_server").(string), "http://")
	if !strings.HasPrefix(apiEndpoint, "https://") {
		apiEndpoint = "https://" + apiEndpoint
	}
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	clientId := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)
	if (username == "" || password == "") && (clientId == "" || clientSecret == "") {
		return nil, fmt.Errorf("One of pair Username/Password or Client_id/client_secret must be set.")
	}
	tokens, err := retrieveTokens()
	if err != nil {
		return nil, err
	}
	options := make([]credhub.Option, 0)
	if username != "" && password != "" {
		options = append(options, credhub.Auth(auth.Uaa(clientId, clientSecret, username, password, tokens.AccessToken, tokens.RefreshToken, false)))
	} else {
		options = append(options, credhub.Auth(auth.Uaa(clientId, clientSecret, username, password, tokens.AccessToken, tokens.RefreshToken, true)))
	}
	if d.Get("skip_ssl_validation").(bool) {
		options = append(options, credhub.SkipTLSValidation(true))
	}
	caCert := d.Get("ca_cert").(string)
	if caCert != "" {
		options = append(options, credhub.CaCerts(caCert))
	}
	client, err := credhub.New(apiEndpoint, options...)
	if err != nil {
		return nil, err
	}

	oauthStrategy := client.Auth.(*auth.OAuthStrategy)
	err = uaaLogin(client, oauthStrategy)
	if err != nil {
		return nil, err
	}

	tokens.AccessToken = oauthStrategy.AccessToken()
	tokens.RefreshToken = oauthStrategy.RefreshToken()
	err = storeTokens(tokens)
	if err != nil {
		return nil, err
	}
	return client, nil
}
func uaaLogin(client *credhub.CredHub, oauthStrat *auth.OAuthStrategy) error {
	_, err := client.GetById("fake")
	if err == nil || !strings.Contains(err.Error(), "invalid_token") {
		return nil
	}
	oauthStrat.SetTokens("", "")
	return oauthStrat.Login()
}
func retrieveTokens() (Tokens, error) {
	tokenPath := filepath.Join(os.TempDir(), TOKENS_FILENAME)

	if _, err := os.Stat(tokenPath); os.IsNotExist(err) {
		return Tokens{}, nil
	}
	b, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return Tokens{}, err
	}
	var tokens Tokens
	err = json.Unmarshal(b, &tokens)
	if err != nil {
		return Tokens{}, err
	}
	return tokens, nil
}
func storeTokens(tokens Tokens, fail ...bool) error {
	b, _ := json.Marshal(tokens)
	err := ioutil.WriteFile(filepath.Join(os.TempDir(), TOKENS_FILENAME), b, 0644)
	if err != nil && len(fail) == 0 {
		time.Sleep(time.Millisecond * 5)
		return storeTokens(tokens, true)
	}
	return err
}
