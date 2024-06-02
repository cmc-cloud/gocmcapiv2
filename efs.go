package gocmcapiv2

import (
	"encoding/json"
)

// EFSService interface
type EFSService interface {
	Get(id string) (EFS, error)
	List(params map[string]string) ([]EFS, error)
	Create(params map[string]interface{}) (EFS, error)
	Update(id string, params map[string]interface{}) (ActionResponse, error)
	Resize(id string, params map[string]interface{}) (ActionResponse, error)
	Delete(id string) (ActionResponse, error)
}

// EFS object
type EFS struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	ProvisioningStatus string `json:"provisioning_status"`
	OperatingStatus    string `json:"operating_status"`
	CreatedAt          string `json:"created_at"`
	VipAddress         string `json:"vip_address"`
	VipPortID          string `json:"vip_port_id"`
	VipSubnetID        string `json:"vip_subnet_id"`
	VipNetworkID       string `json:"vip_network_id"`
	Listeners          []struct {
		ID string `json:"id"`
	} `json:"listeners"`
	Pools []struct {
		ID string `json:"id"`
	} `json:"pools"`
	FlavorID              string   `json:"flavor_id"`
	VipQosPolicyID        string   `json:"vip_qos_policy_id"`
	Tags                  []string `json:"tags"`
	BillingMode           string   `json:"billing_mode"`
	AvailabilityZone      string   `json:"availability_zone"`
	DomesticBandwidthMbps int      `json:"domestic_bandwidth_mbps"`
}

type efs struct {
	client *Client
}

// Get EFS detail
func (v *efs) Get(id string) (EFS, error) {
	jsonStr, err := v.client.Get("efs/"+id, map[string]string{})
	var efs EFS
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &efs)
	}
	return efs, err
}
func (s *efs) List(params map[string]string) ([]EFS, error) {
	restext, err := s.client.Get("efs", params)
	items := make([]EFS, 0)
	err = json.Unmarshal([]byte(restext), &items)
	return items, err
}
func (v *efs) Delete(id string) (ActionResponse, error) {
	return v.client.PerformDelete("efs/" + id)
}
func (v *efs) Update(id string, params map[string]interface{}) (ActionResponse, error) {
	return v.client.PerformUpdate("efs/"+id, params)
}
func (v *efs) Resize(id string, params map[string]interface{}) (ActionResponse, error) {
	return v.client.PerformAction("efs/"+id+"/resize", params)
}
func (v *efs) Create(params map[string]interface{}) (EFS, error) {
	jsonStr, err := v.client.Post("efs", params)
	var response EFS
	if err != nil {
		return response, err
	}
	json.Unmarshal([]byte(jsonStr), &response)
	return response, nil
}
