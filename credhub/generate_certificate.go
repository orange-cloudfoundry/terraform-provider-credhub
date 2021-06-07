package credhub

import (
	"code.cloudfoundry.org/credhub-cli/credhub"
	"code.cloudfoundry.org/credhub-cli/credhub/credentials/generate"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var validKeyUsage map[string]bool = map[string]bool{
	"digital_signature": true,
	"non_repudiation":   true,
	"key_encipherment":  true,
	"data_encipherment": true,
	"key_agreement":     true,
	"key_cert_sign":     true,
	"crl_sign":          true,
	"encipher_only":     true,
	"decipher_only":     true,
}

var validExtendKeyUsage map[string]bool = map[string]bool{
	"client_auth":      true,
	"server_auth":      true,
	"code_signing":     true,
	"email_protection": true,
	"timestamping":     true,
}

type GenerateCertificateResource struct {
}

func (GenerateCertificateResource) Create(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*credhub.CredHub)
	creds, err := client.GenerateCertificate(Name(d), generate.Certificate{
		KeyLength:        d.Get("key_length").(int),
		Duration:         d.Get("duration").(int),
		CommonName:       d.Get("common_name").(string),
		Organization:     d.Get("organization").(string),
		OrganizationUnit: d.Get("organization_unit").(string),
		Locality:         d.Get("locality").(string),
		State:            d.Get("state").(string),
		Country:          d.Get("country").(string),
		AlternativeNames: SchemaSetToStringList(d.Get("alternative_names").(*schema.Set)),
		KeyUsage:         SchemaSetToStringList(d.Get("key_usage").(*schema.Set)),
		ExtendedKeyUsage: SchemaSetToStringList(d.Get("extended_key_usage").(*schema.Set)),
		Ca:               d.Get("ca").(string),
		IsCA:             d.Get("is_ca").(bool),
		SelfSign:         d.Get("self_sign").(bool),
	}, credhub.Overwrite)
	if err != nil {
		return err
	}
	d.SetId(creds.Id)
	return nil
}

func (r GenerateCertificateResource) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key_length": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validateKeyLength,
		},
		"duration": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"common_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"organization": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"organization_unit": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"locality": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"state": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"country": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"alternative_names": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
		},
		"key_usage": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validateFromMap(validKeyUsage, "key usage"),
			},
			Set: schema.HashString,
		},
		"extended_key_usage": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validateFromMap(validExtendKeyUsage, "extended key usage"),
			},
			Set: schema.HashString,
		},
		"ca": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"is_ca": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"self_sign": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}
