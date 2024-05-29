package gocmcapiv2

import (
	"encoding/json"
)

// NetworkService interface
type NetworkService interface {
	Get(id string) (Network, error)
	Update(id, name, description string) error
	Delete(id string) (TaskStatus, error)
	ChangeFirewall(id, firewallID string) (TaskStatus, error)
	CreateVPCNetwork(vpcID, name, description, gateway, netmask, firewallID string) (ResultResponse, error)
}

// ResultResponse return from CreateVPCNetwork
type ResultResponse struct {
	ResultID string `json:"result_id"`
}

// Network object
type Network struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Gateway     string   `json:"gateway"`
	Netmask     string   `json:"netmask"`
	Cidr        string   `json:"cidr"`
	State       string   `json:"state"`
	Type        string   `json:"type"`
	FirewallID  string   `json:"firewall_id"`
	VPCID       string   `json:"vpc_id"`
	ServerIDs   []string `json:"server_ids"`
}
type Nic struct {
	Id        string `json:"id"`
	NetworkID string `json:"net_id"`
	// VpcID      string `json:"vpc_id"`
	MacAddress string `json:"mac_addr"`
	FixedIps   []struct {
		SubnetID  string `json:"subnet_id"`
		IPAddress string `json:"ip_address"`
	} `json:"fixed_ips"`
	SecurityGroups []string `json:"security_groups"`
}
type network struct {
	client *Client
}

// Get Network detail
func (v *network) Get(id string) (Network, error) {
	jsonStr, err := v.client.Get("network/info", map[string]string{"id": id})
	var network Network
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &network)
	}
	return network, err
}

// Delete a Network
func (v *network) Delete(id string) (TaskStatus, error) {
	return v.client.LongDeleteTask("network/delete", id, nil, MediumTimeSettings)
}
func (v *network) Update(id, name, description string) error {
	_, err := v.client.LongTask("network/update", id, map[string]interface{}{"name": name, "description": description}, ShortTimeSettings)
	return err
}
func (v *network) ChangeFirewall(id, firewallID string) (TaskStatus, error) {
	return v.client.LongTask("network/change_firewall", id, map[string]interface{}{"firewall_id": firewallID}, MediumTimeSettings)
}
func (v *network) CreateVPCNetwork(vpcID, name, description, gateway, netmask, firewallID string) (ResultResponse, error) {
	jsonStr, err := v.client.Post("network/create_vpc_network", map[string]interface{}{
		"vpc_id":      vpcID,
		"name":        name,
		"description": description,
		"gateway":     gateway,
		"netmask":     netmask,
		"firewall_id": firewallID,
	})
	var result ResultResponse
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &result)
	}
	return result, err
}
