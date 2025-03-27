package gocmcapiv2

import (
	"encoding/json"
	"fmt"
)

// ELBService interface
type ELBService interface {
	Get(id string) (ELB, error)
	List(params map[string]string) ([]ELB, error)

	GetFlavor(id string) (ELBFlavor, error)
	ListFlavors() ([]ELBFlavor, error)

	Create(params map[string]interface{}) (ELB, error)
	Update(id string, params map[string]interface{}) (ActionResponse, error)
	Resize(id string, params map[string]interface{}) (ActionResponse, error)
	Delete(id string) (ActionResponse, error)

	GetListener(id string) (ELBListener, error)
	GetPool(id string) (ELBPool, error)
	GetHealthMonitor(id string) (ELBHealthMonitor, error)
	GetPoolMember(id string, member_id string) (ELBPoolMember, error)

	CreateListener(id string, params map[string]interface{}) (ELBListener, error)
	CreatePool(id string, params map[string]interface{}) (ELBPool, error)
	CreateHealthMonitor(params map[string]interface{}) (ELBHealthMonitor, error)
	CreatePoolMember(id string, params map[string]interface{}) (ELBPoolMember, error)

	UpdateListener(id string, params map[string]interface{}) (ActionResponse, error)
	UpdatePool(id string, params map[string]interface{}) (ActionResponse, error)
	UpdateHealthMonitor(id string, params map[string]interface{}) (ActionResponse, error)
	UpdatePoolMember(id string, member_id string, params map[string]interface{}) (ActionResponse, error)

	DeleteListener(id string) (ActionResponse, error)
	DeletePool(id string) (ActionResponse, error)
	DeleteHealthMonitor(id string) (ActionResponse, error)
	DeletePoolMember(id string, member_id string) (ActionResponse, error)
}

type ELBFlavor struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

type ELBHealthMonitor struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Delay          int    `json:"delay"`
	Timeout        int    `json:"timeout"`
	MaxRetries     int    `json:"max_retries"`
	MaxRetriesDown int    `json:"max_retries_down"`
	HTTPMethod     string `json:"http_method"`
	URLPath        string `json:"url_path"`
	DomainName     string `json:"domain_name"`
	ExpectedCodes  string `json:"expected_codes"`
	ProjectID      string `json:"project_id"`
	Pools          []struct {
		ID string `json:"id"`
	} `json:"pools"`
	ProvisioningStatus string `json:"provisioning_status"`
	OperatingStatus    string `json:"operating_status"`
	CreatedAt          string `json:"created_at"`
	HTTPVersion        string `json:"http_version"`
}
type ELBPool struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	ProvisioningStatus string `json:"provisioning_status"`
	OperatingStatus    string `json:"operating_status"`
	AdminStateUp       bool   `json:"admin_state_up"`
	Protocol           string `json:"protocol"`
	LbAlgorithm        string `json:"lb_algorithm"`
	SessionPersistence struct {
		Type                   string `json:"type"`
		CookieName             string `json:"cookie_name"`
		PersistenceTimeout     any    `json:"persistence_timeout"`
		PersistenceGranularity any    `json:"persistence_granularity"`
	} `json:"session_persistence"`
	Loadbalancers []struct {
		ID string `json:"id"`
	} `json:"loadbalancers"`
	Listeners []struct {
		ID string `json:"id"`
	} `json:"listeners"`
	CreatedAt       string   `json:"created_at"`
	HealthmonitorID string   `json:"healthmonitor_id"`
	Members         []any    `json:"members"`
	Tags            []Tag    `json:"tags"`
	TLSContainerRef any      `json:"tls_container_ref"`
	TLSEnabled      bool     `json:"tls_enabled"`
	TLSCiphers      string   `json:"tls_ciphers"`
	TLSVersions     []string `json:"tls_versions"`
	AlpnProtocols   []string `json:"alpn_protocols"`
}

/* insert_headers có thể là một mảng rỗng [] hoặc một struct */
type InsertHeaders struct {
	XForwardedFor   string `json:"X-Forwarded-For"`
	XForwardedPort  string `json:"X-Forwarded-Port"`
	XForwardedProto string `json:"X-Forwarded-Proto"`
}
type InsertHeadersOrArray struct {
	Headers *InsertHeaders
	IsArray bool
}

func (i *InsertHeadersOrArray) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as InsertHeaders
	var headers InsertHeaders
	if err := json.Unmarshal(data, &headers); err == nil {
		i.Headers = &headers
		i.IsArray = false
		return nil
	}

	// Try to unmarshal as empty array
	var emptyArray []interface{}
	if err := json.Unmarshal(data, &emptyArray); err == nil && len(emptyArray) == 0 {
		i.Headers = nil
		i.IsArray = true
		return nil
	}

	return fmt.Errorf("insert_headers must be either an empty array or an object")
}

type ELBListener struct {
	ID                     string               `json:"id"`
	Name                   string               `json:"name"`
	Description            string               `json:"description"`
	ProvisioningStatus     string               `json:"provisioning_status"`
	OperatingStatus        string               `json:"operating_status"`
	Protocol               string               `json:"protocol"`
	ProtocolPort           int                  `json:"protocol_port"`
	DefaultTLSContainerRef string               `json:"default_tls_container_ref"`
	SniContainerRefs       []string             `json:"sni_container_refs"`
	DefaultPoolID          string               `json:"default_pool_id"`
	L7Policies             []string             `json:"l7policies"`
	InsertHeaders          InsertHeadersOrArray `json:"insert_headers"`
	CreatedAt              string               `json:"created_at"`
	Loadbalancers          []struct {
		ID string `json:"id"`
	} `json:"loadbalancers"`
	TimeoutClientData       int      `json:"timeout_client_data"`
	TimeoutMemberConnect    int      `json:"timeout_member_connect"`
	TimeoutMemberData       int      `json:"timeout_member_data"`
	TimeoutTCPInspect       int      `json:"timeout_tcp_inspect"`
	ConnectionLimit         int      `json:"connection_limit"`
	AllowedCidrs            []string `json:"allowed_cidrs"`
	Tags                    []Tag    `json:"tags"`
	ClientCaTLSContainerRef string   `json:"client_ca_tls_container_ref"`
	ClientCrlContainerRef   string   `json:"client_crl_container_ref"`
	TLSCiphers              string   `json:"tls_ciphers"`
	TLSVersions             string   `json:"tls_versions"`
	// AllowedCidrs            string   `json:"allowed_cidrs"`
}
type ELBPoolMember struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Address            string `json:"address"`
	ProtocolPort       int    `json:"protocol_port"`
	Weight             int    `json:"weight"`
	SubnetID           string `json:"subnet_id"`
	CreatedAt          string `json:"created_at"`
	MonitorAddress     string `json:"monitor_address"`
	MonitorPort        int    `json:"monitor_port"`
	OperatingStatus    string `json:"operating_status"`
	ProvisioningStatus string `json:"provisioning_status"`
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
	FlavorID              string `json:"flavor_id"`
	VipQosPolicyID        string `json:"vip_qos_policy_id"`
	Tags                  []Tag  `json:"tags"`
	BillingMode           string `json:"billing_mode"`
	AvailabilityZone      string `json:"availability_zone"`
	DomesticBandwidthMbps int    `json:"domestic_bandwidth_mbps"`
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

func (v *elb) GetListener(id string) (ELBListener, error) {
	jsonStr, err := v.client.Get("lbaas/listener/"+id, map[string]string{})
	var elb ELBListener
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &elb)
	}
	return elb, err
}

func (v *elb) GetPool(id string) (ELBPool, error) {
	jsonStr, err := v.client.Get("lbaas/pool/"+id, map[string]string{})
	var elb ELBPool
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &elb)
	}
	return elb, err
}

func (v *elb) GetHealthMonitor(id string) (ELBHealthMonitor, error) {
	jsonStr, err := v.client.Get("lbaas/healthmonitor/"+id, map[string]string{})
	var elb ELBHealthMonitor
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &elb)
	}
	return elb, err
}

func (v *elb) GetPoolMember(id string, member_id string) (ELBPoolMember, error) {
	jsonStr, err := v.client.Get("lbaas/pool/"+id+"/member/"+member_id, map[string]string{})
	var elb ELBPoolMember
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
func (v *elb) DeleteListener(id string) (ActionResponse, error) {
	return v.client.PerformDelete("lbaas/listener/" + id)
}
func (v *elb) DeletePool(id string) (ActionResponse, error) {
	return v.client.PerformDelete("lbaas/pool/" + id)
}
func (v *elb) DeleteHealthMonitor(id string) (ActionResponse, error) {
	return v.client.PerformDelete("lbaas/healthmonitor/" + id)
}
func (v *elb) DeletePoolMember(id string, member_id string) (ActionResponse, error) {
	return v.client.PerformDelete("lbaas/pool/" + id + "/member/" + member_id)
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
	err = json.Unmarshal([]byte(jsonStr), &response)
	return response, err
}

func (v *elb) CreateListener(id string, params map[string]interface{}) (ELBListener, error) {
	jsonStr, err := v.client.Post("lbaas/"+id+"/listener", params)
	var response ELBListener
	if err != nil {
		return response, err
	}
	err = json.Unmarshal([]byte(jsonStr), &response)
	return response, err
}

func (v *elb) CreatePool(id string, params map[string]interface{}) (ELBPool, error) {
	jsonStr, err := v.client.Post("lbaas/"+id+"/pool", params)
	var response ELBPool
	if err != nil {
		return response, err
	}
	err = json.Unmarshal([]byte(jsonStr), &response)
	return response, err
}

func (v *elb) CreateHealthMonitor(params map[string]interface{}) (ELBHealthMonitor, error) {
	jsonStr, err := v.client.Post("lbaas/healthmonitor", params)
	var response ELBHealthMonitor
	if err != nil {
		return response, err
	}
	err = json.Unmarshal([]byte(jsonStr), &response)
	return response, err
}
func (v *elb) CreatePoolMember(id string, params map[string]interface{}) (ELBPoolMember, error) {
	jsonStr, err := v.client.Post("lbaas/pool/"+id+"/member", params)
	var response ELBPoolMember
	if err != nil {
		return response, err
	}
	err = json.Unmarshal([]byte(jsonStr), &response)
	return response, err
}

func (v *elb) UpdateListener(id string, params map[string]interface{}) (ActionResponse, error) {
	return v.client.PerformUpdate("lbaas/listener/"+id, params)
}

func (v *elb) UpdatePool(id string, params map[string]interface{}) (ActionResponse, error) {
	return v.client.PerformUpdate("lbaas/pool/"+id, params)
}

func (v *elb) UpdateHealthMonitor(id string, params map[string]interface{}) (ActionResponse, error) {
	return v.client.PerformUpdate("lbaas/healthmonitor/"+id, params)
}
func (v *elb) UpdatePoolMember(id string, member_id string, params map[string]interface{}) (ActionResponse, error) {
	return v.client.PerformUpdate("lbaas/pool/"+id+"/member/"+member_id, params)
}

func (v *elb) GetFlavor(id string) (ELBFlavor, error) {
	jsonStr, err := v.client.Get("lbaas/flavor/"+id, map[string]string{})
	var vpc ELBFlavor
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &vpc)
	}
	return vpc, err
}

// List lists all flavors. []*Flavor
func (s *elb) ListFlavors() ([]ELBFlavor, error) {
	restext, err := s.client.Get("lbaas/flavor", map[string]string{})

	flavors := make([]ELBFlavor, 0)
	err = json.Unmarshal([]byte(restext), &flavors)
	return flavors, err
}
