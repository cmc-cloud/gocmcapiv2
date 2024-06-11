package gocmcapiv2

import (
	"encoding/json"
)

// KeyManagementService interface
type KeyManagementService interface {
	Get(id string) (KeyManagementContainer, error)
	List(params map[string]string) ([]KeyManagementContainer, error)
	Create(params map[string]interface{}) (KeyManagementContainer, error)
	Delete(id string) (ActionResponse, error)

	GetSecrets(id string, show_nodes bool) ([]KeyManagementSecret, error)
	GetSecret(id string, secret_id string) (KeyManagementSecret, error)
	CreateSecret(id string, params map[string]interface{}) (KeyManagementSecret, error)
	DeleteSecret(id string, secret_id string) (ActionResponse, error)
	UpdateSecret(id string, params map[string]interface{}) (ActionResponse, error)
}

type KeyManagementContainerWrapper struct {
	Data KeyManagementContainer `json:"data"`
}

type KeyManagementSecretWrapper struct {
	Data KeyManagementSecret `json:"data"`
}

type KeyManagementContainerListWrapper struct {
	Data struct {
		Docs      []KeyManagementContainer `json:"docs"`
		Page      int                      `json:"page"`
		Size      int                      `json:"size"`
		Total     int                      `json:"total"`
		TotalPage int                      `json:"totalPage"`
	} `json:"data"`
}

type KeyManagementSecretListWrapper struct {
	Data struct {
		Docs      []KeyManagementSecret `json:"docs"`
		Page      int                   `json:"page"`
		Size      int                   `json:"size"`
		Total     int                   `json:"total"`
		TotalPage int                   `json:"totalPage"`
	} `json:"data"`
}

type KeyManagementContainer struct {
	ID      string `json:"containerUuid"`
	Name    string `json:"containerName"`
	Type    string `json:"containerType"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

type KeyManagementSecret struct {
	Algorithm  string `json:"algorithm"`
	BitLength  string `json:"bit_length"`
	Created    string `json:"created"`
	CreatorID  string `json:"creator_id"`
	Expiration string `json:"expiration"`
	Mode       string `json:"mode"`
	Name       string `json:"name"`
	SecretRef  string `json:"secret_ref"`
	SecretType string `json:"secret_type"`
	Status     string `json:"status"`
	Updated    string `json:"updated"`
	// ContentTypes struct {
	// 	Default string `json:"default"`
	// } `json:"content_types"`
}
type keymanagement struct {
	client *Client
}

type KeyManagementCreatedResponse struct {
	Data struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

// Get keymanagement detail
func (v *keymanagement) Get(id string) (KeyManagementContainer, error) {
	jsonStr, err := v.client.Get("cloudops-core/api/v1/bbc-keys/containers/"+id, map[string]string{})
	var obj KeyManagementContainerWrapper
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &obj)
	}
	return obj.Data, err
}

func (s *keymanagement) List(params map[string]string) ([]KeyManagementContainer, error) {
	jsonStr, err := s.client.Get("cloudops-core/api/v1/bbc-keys/containers", params)
	var obj KeyManagementContainerListWrapper
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &obj)
	}
	return obj.Data.Docs, err
}
func (s *keymanagement) GetSecrets(id string, show_nodes bool) ([]KeyManagementSecret, error) {
	jsonStr, err := s.client.Get("cloudops-core/api/v1/bbc-keys/containers/"+id+"/node-groups", map[string]string{})
	var obj KeyManagementSecretListWrapper
	if err != nil {
		var nilres []KeyManagementSecret
		return nilres, err
	}
	err = json.Unmarshal([]byte(jsonStr), &obj)
	return obj.Data.Docs, err
}
func (v *keymanagement) GetSecret(id string, secret_id string) (KeyManagementSecret, error) {
	jsonStr, err := v.client.Get("cloudops-core/api/v1/bbc-keys/containers/"+id+"/node-groups/"+secret_id, map[string]string{})
	var response KeyManagementSecretWrapper
	if err != nil {
		var nilres KeyManagementSecret
		return nilres, err
	}
	json.Unmarshal([]byte(jsonStr), &response)
	// Logo("get nodegroup ", response)
	return response.Data, nil
}

// Delete a keymanagement
func (v *keymanagement) Delete(id string) (ActionResponse, error) {
	return v.client.PerformDelete("cloudops-core/api/v1/bbc-keys/containers/" + id)
}
func (v *keymanagement) DeleteSecret(id string, secret_id string) (ActionResponse, error) {
	return v.client.PerformDelete("cloudops-core/api/v1/bbc-keys/containers/" + id + "/secrets/" + secret_id)
}
func (s *keymanagement) Create(params map[string]interface{}) (KeyManagementContainer, error) {
	jsonStr, err := s.client.Post("cloudops-core/api/v1/bbc-keys/containers", params)
	var response KeyManagementContainer
	if err != nil {
		return response, err
	}
	json.Unmarshal([]byte(jsonStr), &response)
	return response, nil
}

func (s *keymanagement) CreateSecret(id string, params map[string]interface{}) (KeyManagementSecret, error) {
	jsonStr, err := s.client.Post("cloudops-core/api/v1/bbc-keys/containers/"+id+"/secrets", params)
	var response KeyManagementSecretWrapper
	if err != nil {
		var nilres KeyManagementSecret
		return nilres, err
	}
	json.Unmarshal([]byte(jsonStr), &response)
	return response.Data, nil
}

func (s *keymanagement) UpdateSecret(id string, params map[string]interface{}) (ActionResponse, error) {
	return s.client.PerformAction("cloudops-core/api/v1/bbc-keys/containers/"+id+"/secrets", params)
}
