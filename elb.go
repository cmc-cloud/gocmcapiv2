package gocmcapiv2

import (
	"encoding/json"
)

// ELBService interface
type ELBService interface {
	Get(id string) (ELB, error)
	List(params map[string]string) ([]ELB, error)
	Create(params map[string]interface{}) (ELB, error)
	Update(id string, params map[string]interface{}) (ActionResponse, error)
	Resize(id string, params map[string]interface{}) (ActionResponse, error)
	Delete(id string) (ActionResponse, error)
}

// ELB object
type ELB struct {
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

type elb struct {
	client *Client
}

// Get ELB detail
func (v *elb) Get(id string) (ELB, error) {
	jsonStr, err := v.client.Get("lbaas/"+id, map[string]string{})
	var elb ELB
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &elb)
	}
	return elb, err
}

func (s *elb) List(params map[string]string) ([]ELB, error) {
	restext, err := s.client.Get("lbaas", params)
	items := make([]ELB, 0)
	err = json.Unmarshal([]byte(restext), &items)
	return items, err
}

// Delete a ELB
func (v *elb) Delete(id string) (ActionResponse, error) {
	return v.client.PerformDelete("lbaas/" + id)
}
func (v *elb) Update(id string, params map[string]interface{}) (ActionResponse, error) {
	return v.client.PerformUpdate("lbaas/"+id, params)
}
func (v *elb) Resize(id string, params map[string]interface{}) (ActionResponse, error) {
	return v.client.PerformAction("lbaas/"+id+"/resize", params)
}
func (v *elb) Create(params map[string]interface{}) (ELB, error) {
	jsonStr, err := v.client.Post("lbaas", params)
	var response ELB
	if err != nil {
		return response, err
	}
	json.Unmarshal([]byte(jsonStr), &response)
	return response, nil
}
