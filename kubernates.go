package gocmcapiv2

import (
	"encoding/json"
)

// KubernatesService interface
type KubernatesService interface {
	Get(id string) (Kubernates, error)
	List(params map[string]string) ([]Kubernates, error)
	Create(params map[string]interface{}) (KubernatesCreatedResponse, error)
	Delete(id string) (ActionResponse, error)
	UpdateNodeCount(id string, node_count int) (ActionResponse, error)

	GetNodeGroup(id string, nodegroup_id string) (KubernatesNodeGroup, error)
	CreateNodeGroup(id string, params map[string]interface{}) (KubernatesNodeGroup, error)
	DeleteNodeGroup(id string, nodegroup_id string) (ActionResponse, error)
	ResizeNodeGroup(id string, params map[string]interface{}) (ActionResponse, error)
	UpdateNodeGroup(id string, nodegroup_id string, min_node_count int, max_node_count int) (ActionResponse, error)
}

// Kubernates object
type Kubernates struct {
	ID               string `json:"uuid"`
	Name             string `json:"name"`
	Keypair          string `json:"keypair"`
	NodeCount        int    `json:"node_count"`
	MasterCount      int    `json:"master_count"`
	DockerVolumeSize int    `json:"docker_volume_size"`
	Labels           struct {
		AvailabilityZone     string         `json:"availability_zone"`
		KubeTag              string         `json:"kube_tag"`
		DockerVolumeType     string         `json:"docker_volume_type"`
		KubeDashboardEnabled BoolFromString `json:"kube_dashboard_enabled"`
		MetricsServerEnabled BoolFromString `json:"metrics_server_enabled"`
		NpdEnabled           BoolFromString `json:"npd_enabled"`
		AutoScalingEnabled   BoolFromString `json:"auto_scaling_enabled"`
		AutoHealingEnabled   BoolFromString `json:"auto_healing_enabled"`
		MinNodeCount         IntFromString  `json:"min_node_count"`
		MaxNodeCount         IntFromString  `json:"max_node_count"`
		CalicoIpv4Pool       string         `json:"calico_ipv4pool"`
		CalicoIpv4PoolIpip   string         `json:"calico_ipv4pool_ipip"`
		NetworkDriver        string         `json:"network-driver"`
	} `json:"labels"`
	SubnetID          string `json:"fixed_subnet"`
	MasterFlavorID    string `json:"master_flavor_id"`
	MasterBillingMode string `json:"master_billing_mode"`
	NodeBillingMode   string `json:"node_billing_mode"`
	NodeFlavorID      string `json:"flavor_id"`
	Status            string `json:"status"`
	HealthStatus      string `json:"health_status"`
	CreateTimeout     int    `json:"create_timeout"`
	CreatedAt         string `json:"created_at"`
}
type KubernatesNodeGroup struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ClusterID        string `json:"cluster_id"`
	DockerVolumeSize int    `json:"docker_volume_size"`
	Labels           struct {
		DockerVolumeType string `json:"docker_volume_type"`
		AvailabilityZone string `json:"availability_zone"`
	} `json:"labels"`
	FlavorID      string   `json:"flavor_id"`
	NodeAddresses []string `json:"node_addresses"`
	NodeCount     int      `json:"node_count"`
	Role          string   `json:"role"`
	MinNodeCount  int      `json:"min_node_count"`
	MaxNodeCount  int      `json:"max_node_count"`
	Status        string   `json:"status"`
	StatusReason  string   `json:"status_reason"`
	CreatedAt     string   `json:"created_at"`
	Nodes         []struct {
		NodegroupID int    `json:"nodegroup_id"`
		Name        string `json:"name"`
		ID          string `json:"id"`
		Created     string `json:"created"`
		Status      string `json:"status"`
	} `json:"nodes"`
	BillingMode string `json:"billing_mode"`
}
type kubernetes struct {
	client *Client
}

type KubernatesCreatedResponse struct {
	ID      string `json:"uuid"`
	Success bool   `json:"success"`
}

// Get kubernetes detail
func (v *kubernetes) Get(id string) (Kubernates, error) {
	jsonStr, err := v.client.Get("kubernetes/cluster/"+id, map[string]string{})
	var obj Kubernates
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &obj)
	}
	return obj, err
}

func (s *kubernetes) List(params map[string]string) ([]Kubernates, error) {
	restext, err := s.client.Get("kubernetes/cluster", params)
	items := make([]Kubernates, 0)
	err = json.Unmarshal([]byte(restext), &items)
	return items, err
}
func (v *kubernetes) GetNodeGroup(id string, nodegroup_id string) (KubernatesNodeGroup, error) {
	jsonStr, err := v.client.Get("kubernetes/cluster/"+id+"/nodegroup/"+nodegroup_id, map[string]string{})
	var obj KubernatesNodeGroup
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &obj)
	}
	return obj, err
}

// Delete a kubernetes
func (v *kubernetes) Delete(id string) (ActionResponse, error) {
	return v.client.PerformDelete("kubernetes/cluster/" + id)
}
func (v *kubernetes) DeleteNodeGroup(id string, nodegroup_id string) (ActionResponse, error) {
	return v.client.PerformDelete("kubernetes/cluster/" + id + "/nodegroup/" + nodegroup_id)
}
func (v *kubernetes) UpdateNodeCount(id string, node_count int) (ActionResponse, error) {
	return v.client.PerformUpdate("kubernetes/cluster/"+id, map[string]interface{}{
		"node_count": node_count,
	})
}
func (s *kubernetes) Create(params map[string]interface{}) (KubernatesCreatedResponse, error) {
	jsonStr, err := s.client.Post("kubernetes/cluster", params)
	var response KubernatesCreatedResponse
	if err != nil {
		return response, err
	}
	json.Unmarshal([]byte(jsonStr), &response)
	return response, nil
}

func (s *kubernetes) CreateNodeGroup(id string, params map[string]interface{}) (KubernatesNodeGroup, error) {
	jsonStr, err := s.client.Post("kubernetes/cluster/"+id+"/nodegroup", params)
	var response KubernatesNodeGroup
	if err != nil {
		return response, err
	}
	json.Unmarshal([]byte(jsonStr), &response)
	return response, nil
}

func (s *kubernetes) ResizeNodeGroup(id string, params map[string]interface{}) (ActionResponse, error) {
	return s.client.PerformAction("kubernetes/cluster/"+id+"/resize", params)
}

func (s *kubernetes) UpdateNodeGroup(id string, nodegroup_id string, min_node_count int, max_node_count int) (ActionResponse, error) {
	return s.client.PerformUpdate("kubernetes/cluster/"+id+"/nodegroup/"+nodegroup_id, map[string]interface{}{
		"min_node_count": min_node_count,
		"max_node_count": max_node_count,
	})
}
