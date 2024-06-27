package gocmcapiv2

import (
	"encoding/json"
)

// CDNCertService interface
type CDNCertService interface {
	Get(id string) (CDNCert, error)
	List(params map[string]string) ([]CDNCert, error)
	Create(params map[string]interface{}) (CDNCert, error)
	Update(id string, params map[string]interface{}) (ActionResponse, error)
	Delete(id string) (ActionResponse, error)
}

type CDNCert struct {
	ID              string `json:"id"`
	CertificateType string `json:"certificate_type"`
	CommonName      string `json:"common_name"`
	ExpirationDate  string `json:"expiration_date"`
	Status          string `json:"status"`
}
type CDNCertListWrapper struct {
	Data     []CDNCert `json:"data"`
	PageInfo struct {
		Page       string `json:"page"`
		PerPage    int    `json:"per_page"`
		TotalCount int    `json:"total_count"`
		TotalPages int    `json:"total_pages"`
	} `json:"page_info"`
}

type cdncert struct {
	client *Client
}

// Get CDNCert detail
func (v *cdncert) Get(id string) (CDNCert, error) {
	jsonStr, err := v.client.Get("cdn/ssl/"+id, map[string]string{})
	var cdncert CDNCert
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &cdncert)
	}
	return cdncert, err
}
func (v *cdncert) List(params map[string]string) ([]CDNCert, error) {
	jsonStr, err := v.client.Get("cdn/ssl", map[string]string{})
	var cdncert CDNCertListWrapper
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &cdncert)
	}
	return cdncert.Data, err
}
func (v *cdncert) Delete(id string) (ActionResponse, error) {
	return v.client.PerformDelete("cdn/ssl/" + id)
}
func (v *cdncert) Create(params map[string]interface{}) (CDNCert, error) {
	jsonStr, err := v.client.Post("cdn/ssl/customssl", params)
	var response CDNCert
	if err != nil {
		return response, err
	}
	json.Unmarshal([]byte(jsonStr), &response)
	return response, nil
}
func (s *cdncert) Update(id string, params map[string]interface{}) (ActionResponse, error) {
	return s.client.PerformUpdate("cdn/ssl/"+id, params)
}
