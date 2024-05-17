package extend

import (
	"code.cloudfoundry.org/credhub-cli/credhub"
	"code.cloudfoundry.org/credhub-cli/credhub/permissions"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Permission struct {
	client *credhub.CredHub
}

func NewPermission(client *credhub.CredHub) *Permission {
	return &Permission{client}
}

func (p *Permission) updateV1Permission(uuid string, credName string, perms []permissions.V1_Permission) (*http.Response, error) {
	ch := p.client
	requestBody := map[string]interface{}{}
	requestBody["credential_name"] = credName
	requestBody["permissions"] = perms

	resp, err := ch.Request(http.MethodPut, fmt.Sprintf("/api/v1/permissions/%s", uuid), nil, requestBody, true)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *Permission) updateV2Permission(uuid string, path string, actor string, ops []string) (*http.Response, error) {
	ch := p.client
	requestBody := map[string]interface{}{}
	requestBody["path"] = path
	requestBody["actor"] = actor
	requestBody["operations"] = ops

	resp, err := ch.Request(http.MethodPut, fmt.Sprintf("/api/v2/permissions/%s", uuid), nil, requestBody, true)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *Permission) UpdatePermission(uuid string, path string, actor string, ops []string) (*permissions.Permission, error) {
	ch := p.client
	serverVersion, err := ch.ServerVersion()
	if err != nil {
		return nil, err
	}

	var resp *http.Response
	isOlderVersion := serverVersion.Segments()[0] < 2

	if isOlderVersion {
		resp, err = p.updateV1Permission(uuid, path, []permissions.V1_Permission{{Actor: actor, Operations: ops}})
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
	} else {
		resp, err = p.updateV2Permission(uuid, path, actor, ops)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
	}

	if isOlderVersion {
		return nil, nil
	}
	//nolint:errcheck
	defer io.Copy(io.Discard, resp.Body)
	var response permissions.Permission

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *Permission) DeletePermission(uuid string) error {
	ch := p.client
	resp, err := ch.Request(http.MethodDelete, fmt.Sprintf("/api/v2/permissions/%s", uuid), nil, nil, true)

	if err != nil {
		return err
	}
	return resp.Body.Close()
}
