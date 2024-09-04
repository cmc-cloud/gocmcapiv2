package gocmcapiv2

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

// Kubernetesv2Service interface
type Kubernetesv2Service interface {
	Get(id string) (Kubernetesv2, error)
	GetStatus(id string) (Kubernetesv2Status, error)
	List(params map[string]string) ([]Kubernetesv2ListItem, error)
	Create(params map[string]interface{}) (Kubernetesv2CreatedResponse, error)
	Delete(id string) (ActionResponse, error)

	GetNodeGroups(id string, show_nodes bool) ([]Kubernetesv2NodeGroup, error)
	GetNodeGroup(id string, nodegroup_id string) (Kubernetesv2NodeGroup, error)
	CreateNodeGroup(id string, params map[string]interface{}) (Kubernetesv2NodeGroup, error)
	DeleteNodeGroup(id string, nodegroup_id string) (ActionResponse, error)
	UpdateNodeGroup(id string, params map[string]interface{}) (ActionResponse, error)
	UpdateAddon(id string, params map[string]interface{}) (ActionResponse, error)
}
type Kubernetesv2StatusWrapper struct {
	Data Kubernetesv2Status `json:"data"`
}
type Kubernetesv2Status struct {
	EtcdClusterReady  BoolFromString `json:"etcdClusterReady"`
	MachinesReady     BoolFromString `json:"machinesReady"`
	ControlPlaneReady BoolFromString `json:"controlPlaneReady"`
	EnableAutoScale   BoolFromString `json:"enableAutoScale"`
	EnableAutoHealing BoolFromString `json:"enableAutoHealing"`
	EnableMonitor     BoolFromString `json:"enableMonitor"`
}

type Kubernetesv2NodeGroup struct {
	ID                        string                        `json:"id"`
	Name                      string                        `json:"nodeGroupName"`
	KeyName                   string                        `json:"sshKeyName"`
	NodeCount                 int                           `json:"replicas"`
	Status                    string                        `json:"status"`
	MetadataMachineDeployment AutoScalev2MetadataFromString `json:"metadataMachineDeployment"`
	ExternalProviders         []struct {
		Name      string                      `json:"name"`
		Config    AutoScalev2ConfigFromString `json:"config"`
		ClusterID string                      `json:"clusterId"`
		Status    string                      `json:"status"`
	} `json:"externalProviders"`
	// Nodes []struct {
	// 	Name              string    `json:"name"`
	// 	Status            string    `json:"status"`
	// 	Message           string    `json:"message"`
	// 	ProviderID        string    `json:"providerId"`
	// 	CreationTimestamp time.Time `json:"creationTimestamp"`
	// } `json:"nodes"`
	// Version                   string                        `json:"version"`
	// SSHKeyName                string                        `json:"sshKeyName"`
}

type AutoScalev2Metadata struct {
	Image      string `json:"image"`
	FlavorName string `json:"flavor"`
}
type AutoScalev2Config struct {
	// MemoryThreshold   string `json:"memoryThreshold"`
	// CPUThreshold      string `json:"cpuThreshold"`
	// DiskThreshold     string `json:"diskThreshold"`
	MaxPods           int `json:"maxPods"`
	MinNode           int `json:"minNode"`
	MaxNode           int `json:"maxNode"`
	MetaDataAutoScale struct {
		PercentCPU    int `json:"percentCpu"`
		PercentMemory int `json:"percentMemory"`
		PercentDisk   int `json:"percentDisk"`
	} `json:"metaDataAutoScale"`
	MaxUnhealthy       ExtractIntFromString `json:"maxUnhealthy"`
	NodeStartupTimeout ExtractIntFromString `json:"nodeStartupTimeout"`
}

type AutoScalev2MetadataFromString AutoScalev2Metadata
type AutoScalev2ConfigFromString AutoScalev2Config
type ExtractIntFromString int

func (b *AutoScalev2ConfigFromString) UnmarshalJSON(data []byte) error {
	var val AutoScalev2Config
	input := string(data)
	input = strings.Trim(input, `"`)
	input = strings.ReplaceAll(input, `\`, ``)
	if err := json.Unmarshal([]byte(input), &val); err != nil {
		Logo("AutoScalev2Config Unmarshal err =", err)
		return err
	}
	// Logo("AutoScalev2Config after Unmarshal = ", val)
	*b = AutoScalev2ConfigFromString(val)
	return nil
}
func (b *AutoScalev2MetadataFromString) UnmarshalJSON(data []byte) error {
	var val AutoScalev2Metadata
	input := string(data)
	input = strings.Trim(input, `"`)
	input = strings.ReplaceAll(input, `\`, ``)
	if err := json.Unmarshal([]byte(input), &val); err != nil {
		Logo("AutoScalev2Metadata Unmarshal err =", err)
		return err
	}
	// Logo("AutoScalev2MetadataFromString after Unmarshal = ", val)
	*b = AutoScalev2MetadataFromString(val)
	return nil
}

// dau vao la 40%, 10m => dau ra la 40, 10
func (b *ExtractIntFromString) UnmarshalJSON(data []byte) error {
	// xoa tat ca chi de lai so
	input := string(data)
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(input, -1)[0]
	// Parse the string value to an integer
	intValue, err := strconv.Atoi(matches)
	if err != nil {
		return err
	}
	// Set the integer value to the custom type
	*b = ExtractIntFromString(intValue)
	return nil
}

type Kubernetesv2NodeGroupWrapper struct {
	Data Kubernetesv2NodeGroup `json:"data"`
}
type Kubernetesv2NodeGroupListWrapper struct {
	Data []Kubernetesv2NodeGroup `json:"data"`
}
type Kubernetesv2Wrapper struct {
	Data Kubernetesv2 `json:"data"`
}
type Kubernetesv2ListWrapper struct {
	Data struct {
		Page      int                    `json:"page"`
		Size      int                    `json:"size"`
		Total     int                    `json:"total"`
		TotalPage int                    `json:"totalPage"`
		Docs      []Kubernetesv2ListItem `json:"docs"`
	} `json:"data"`
}

type Kubernetesv2ListItem struct {
	ClusterID        string `json:"clusterId"`
	ClusterName      string `json:"clusterName"`
	State            string `json:"state"`
	NumberMasterNode int    `json:"numberMasterNode"`
	NumberWorkerNode int    `json:"numberWorkerNode"`
	CreatedAt        string `json:"createdAt"`
	Namespace        string `json:"namespace"`
}
type Kubernetesv2 struct {
	ClusterID        string `json:"clusterId"`
	ClusterName      string `json:"clusterName"`
	State            string `json:"state"`
	NumberMasterNode int    `json:"numberMasterNode"`
	NumberWorkerNode int    `json:"numberWorkerNode"`
	CreatedAt        string `json:"createdAt"`
	KubeletVersion   string `json:"kubeletVersion"`
	VpcID            string `json:"vpcId"`
	SubnetID         string `json:"subnetId"`
	CidrBlockPod     string `json:"cidrBlockPod"`
	ServiceDomain    string `json:"serviceDomain"`
	SecurityGroupID  string `json:"securityGroupId"`
	MasterURL        string `json:"masterUrl"`
	CidrBlockService string `json:"cidrBlockService"`
}
type Kubernetesv2Addon struct {
	Name      string `json:"name"`
	ClusterID string `json:"cluster_id"`
	Status    string `json:"status"`
	Config    string `json:"config"`
}
type kubernetesv2 struct {
	client *Client
}

type Kubernetesv2CreatedResponse struct {
	Data struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

// Get kubernetesv2 detail
func (v *kubernetesv2) Get(id string) (Kubernetesv2, error) {
	jsonStr, err := v.client.Get("cloudops-core/api/v1/k8s/clusters/"+id, map[string]string{})
	var obj Kubernetesv2Wrapper
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &obj)
	}
	return obj.Data, err
}

// Get kubernetesv2 detail
func (v *kubernetesv2) GetStatus(id string) (Kubernetesv2Status, error) {
	jsonStr, err := v.client.Get("cloudops-core/api/v1/k8s/clusters/"+id+"/status", map[string]string{})
	var obj Kubernetesv2StatusWrapper
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &obj)
	}
	return obj.Data, err
}

func (s *kubernetesv2) List(params map[string]string) ([]Kubernetesv2ListItem, error) {
	jsonStr, err := s.client.Get("cloudops-core/api/v1/k8s/clusters", params)
	var obj Kubernetesv2ListWrapper
	if err != nil {
		return []Kubernetesv2ListItem{}, err
	}
	err = json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		return []Kubernetesv2ListItem{}, err
	}
	return obj.Data.Docs, err
}
func (s *kubernetesv2) GetNodeGroups(id string, show_nodes bool) ([]Kubernetesv2NodeGroup, error) {
	jsonStr, err := s.client.Get("cloudops-core/api/v1/k8s/clusters/"+id+"/node-groups", map[string]string{})
	var obj Kubernetesv2NodeGroupListWrapper
	if err != nil {
		var nilres []Kubernetesv2NodeGroup
		return nilres, err
	}
	err = json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		return []Kubernetesv2NodeGroup{}, err
	}
	return obj.Data, err
}
func (v *kubernetesv2) GetNodeGroup(id string, nodegroup_id string) (Kubernetesv2NodeGroup, error) {
	jsonStr, err := v.client.Get("cloudops-core/api/v1/k8s/clusters/"+id+"/node-groups/"+nodegroup_id, map[string]string{})
	var response Kubernetesv2NodeGroupWrapper
	var nilres Kubernetesv2NodeGroup
	if err != nil {
		return nilres, err
	}
	err = json.Unmarshal([]byte(jsonStr), &response)
	if err != nil {
		return nilres, err
	}
	return response.Data, err
}

// Delete a kubernetesv2
func (v *kubernetesv2) Delete(id string) (ActionResponse, error) {
	return v.client.PerformDelete("cloudops-core/api/v1/k8s/clusters/" + id)
}
func (v *kubernetesv2) DeleteNodeGroup(id string, nodegroup_id string) (ActionResponse, error) {
	return v.client.PerformDelete("cloudops-core/api/v1/k8s/clusters/" + id + "/node-groups/" + nodegroup_id)
}
func (s *kubernetesv2) Create(params map[string]interface{}) (Kubernetesv2CreatedResponse, error) {
	jsonStr, err := s.client.Post("cloudops-core/api/v1/k8s/clusters", params)
	var response Kubernetesv2CreatedResponse
	if err != nil {
		return response, err
	}
	err = json.Unmarshal([]byte(jsonStr), &response)
	return response, err
}

func (s *kubernetesv2) CreateNodeGroup(id string, params map[string]interface{}) (Kubernetesv2NodeGroup, error) {
	jsonStr, err := s.client.Post("cloudops-core/api/v1/k8s/clusters/"+id+"/node-groups", params)
	var response Kubernetesv2NodeGroupWrapper
	var nilres Kubernetesv2NodeGroup
	if err != nil {
		return nilres, err
	}
	err = json.Unmarshal([]byte(jsonStr), &response)
	if err != nil {
		return nilres, err
	}
	return response.Data, err
}

func (s *kubernetesv2) UpdateNodeGroup(id string, params map[string]interface{}) (ActionResponse, error) {
	return s.client.PerformAction("cloudops-core/api/v1/k8s/clusters/"+id+"/addons", params)
}

func (s *kubernetesv2) UpdateAddon(id string, params map[string]interface{}) (ActionResponse, error) {
	return s.client.PerformAction("cloudops-core/api/v1/k8s/clusters/"+id+"/addons", params)
}
