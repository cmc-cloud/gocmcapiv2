package gocmcapiv2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

// const (
// 	defaultAPIURL = "https://apiv2.cloud.cmctelecom.vn"
// )

var (
	// ErrNotFound for resource not found status
	ErrNotFound = errors.New("Resource not found")
	// ErrPermissionDenied for permission denied
	ErrPermissionDenied = errors.New("You are not allowed to do this action")
	// ErrCommon for common error
	ErrCommon = errors.New("Error")
)

// OrderResponse response when create a Server
type OrderResponse struct {
	//Success    bool   `json:"success"`
	// ID     string `json:"id"`
	TaskID string `json:"jobid"`
	Price  int    `json:"price"`
	Paid   bool   `json:"paid"`
}

type ClientConfigs struct {
	APIKey      string
	APIEndpoint string
	ProjectId   string
	RegionId    string
	// apiURL    string
	// apiKey    string
	// projectId string
	// regionId  string
}

// Client represents CMC Cloud API client.
type Client struct {
	Configs                  ClientConfigs
	Account                  AccountService
	Image                    ImageService
	Flavor                   FlavorService
	Server                   ServerService
	EcsGroup                 EcsGroupService
	Task                     TaskService
	Volume                   VolumeService
	VolumeType               VolumeTypeService
	NetworkInterface         NetworkInterfaceService
	VolumeAutoBackup         VolumeAutoBackupService
	VPC                      VPCService
	Subnet                   SubnetService
	EIP                      EIPService
	ELB                      ELBService
	EFS                      EFSService
	VA                       VAService
	Waf                      WafService
	WafIP                    WafIPService
	WafRule                  WafRuleService
	WafCert                  WafCertService
	Dns                      DnsZoneService
	DnsRecord                DnsRecordService
	DnsAcl                   DnsAclService
	WafWhitelist             WafWhitelistService
	Certificate              CertificateService
	DatabaseInstance         DatabaseInstanceService
	DatabaseBackup           DatabaseBackupService
	DatabaseAutoBackup       DatabaseAutoBackupService
	ContainerRegistry        ContainerRegistryService
	DevopsProject            DevopsProjectService
	RedisConfiguration       RedisConfigurationService
	RedisInstance            RedisInstanceService
	Kubernetes               KubernetesService
	Kubernetesv2             Kubernetesv2Service
	DatabaseConfiguration    DatabaseConfigurationService
	SecurityGroup            SecurityGroupService
	Keypair                  KeypairService
	KeyManagement            KeyManagementService
	Snapshot                 SnapshotService
	Backup                   BackupService
	AutoScalingConfiguration AutoScalingConfigurationService
	AutoScalingGroup         AutoScalingGroupService
	AutoScalingPolicy        AutoScalingPolicyService
	BillingMode              BillingModeService
}

// APIError is return when there are an error when call api
type APIError struct {
	Success bool `json:"success"`
	Error   struct {
		ErrorCode int    `json:"code"`
		ErrorText string `json:"message"`
	} `json:"error"`
	// ErrorCode int    `json:"error_code"`
	// ErrorText string `json:"error_text"`
}
type SecurityAPIError struct {
	ErrorCode int    `json:"code"`
	ErrorText string `json:"error"`
}
type DnsAPIError struct {
	ErrorText string `json:"error"`
}

// Timeout is timeout info for a long task
// type Timeout struct {
// 	Delay      time.Duration `default:"geek"` // Wait this time before starting checks
// 	Timeout    time.Duration // The amount of time to wait before timeout
// 	MinTimeout time.Duration // Smallest time to wait before refreshes
// }

// NewClient creates new CMC Cloud Api client.
func NewClient(configs ClientConfigs) (*Client, error) {
	c := &Client{}
	c.Configs = configs
	c.Account = &account{client: c}
	c.Image = &image{client: c}
	c.Flavor = &flavor{client: c}
	c.Server = &server{client: c}
	c.EcsGroup = &ecsgroup{client: c}
	c.Task = &task{client: c}
	c.Volume = &volume{client: c}
	c.VolumeType = &volumetype{client: c}
	c.VolumeAutoBackup = &volumeautobackup{client: c}
	c.VPC = &vpc{client: c}
	c.Subnet = &subnet{client: c}
	c.NetworkInterface = &networkinterface{client: c}
	c.EIP = &eip{client: c}
	c.ELB = &elb{client: c}
	c.EFS = &efs{client: c}
	c.VA = &va{client: c}
	c.Waf = &waf{client: c}
	c.WafIP = &wafip{client: c}
	c.WafRule = &wafrule{client: c}
	c.WafCert = &wafcert{client: c}
	c.Dns = &dns{client: c}
	c.DnsRecord = &dnsrecord{client: c}
	c.DnsAcl = &dnsacl{client: c}
	c.WafWhitelist = &wafwhitelist{client: c}
	c.Certificate = &certificate{client: c}
	c.Kubernetes = &kubernetes{client: c}
	c.Kubernetesv2 = &kubernetesv2{client: c}
	c.DevopsProject = &devopsproject{client: c}
	c.ContainerRegistry = &containerregistry{client: c}
	c.RedisConfiguration = &redisconfiguration{client: c}
	c.RedisInstance = &redisinstance{client: c}
	c.DatabaseInstance = &databaseinstance{client: c}
	c.DatabaseAutoBackup = &databaseautobackup{client: c}
	c.DatabaseConfiguration = &databaseconfiguration{client: c}
	c.SecurityGroup = &securitygroup{client: c}
	c.Keypair = &keypair{client: c}
	c.KeyManagement = &keymanagement{client: c}
	c.Snapshot = &snapshot{client: c}
	c.Backup = &backup{client: c}
	c.AutoScalingConfiguration = &asconfiguration{client: c}
	c.AutoScalingGroup = &autoscalinggroup{client: c}
	c.AutoScalingPolicy = &autoscalingpolicy{client: c}
	c.BillingMode = &billingmode{client: c}
	return c, nil
}

func (c *Client) createRequest(params map[string]string, ctx context.Context) *resty.Request {
	client := resty.New()

	if params == nil {
		params = make(map[string]string)
	}

	//c.apiKey = "vTMSG7F9mFKnNRYIz8eA9N9XrHJ4zP"
	params["api_key"] = c.Configs.APIKey

	//var obj interface{}
	request := client.R().
		SetContext(ctx).
		SetHeader("Accept", "application/json").
		SetHeader("Project-Id", c.Configs.ProjectId).
		SetHeader("Region-Id", c.Configs.RegionId).
		// SetAuthToken(c.Configs.apiKey).
		SetError(&APIError{}).
		SetQueryParams(params)

	return request
}
func (c *Client) parseResponse(response *resty.Response, err error) (string, error) {
	restext := response.String() // fmt.Sprint(response)
	if err != nil {
		return restext, err
	}
	if response.Error() != nil {
		apiError := response.Error().(*APIError)
		if apiError != nil {
			if apiError.Error.ErrorCode == 0 {
				apiError.Error.ErrorCode = response.StatusCode()
			}

			code := apiError.Error.ErrorCode
			if code == http.StatusNotFound {
				return restext, fmt.Errorf("%s: %w", apiError.Error.ErrorText, ErrNotFound)
			}

			// sua lai ma loi voi api cloudops-core
			requestURL := response.Request.URL
			if strings.Contains(requestURL, "cloudops-core") && apiError.Error.ErrorCode == 500 && strings.Contains(apiError.Error.ErrorText, "not found") {
				return restext, fmt.Errorf("%s: %w", apiError.Error.ErrorText, ErrNotFound)
			}
			return restext, fmt.Errorf("Error %d: %s", apiError.Error.ErrorCode, apiError.Error.ErrorText)
		}
	}

	if strings.Contains(restext, "error_code") && strings.Contains(restext, "error_text") {
		var apiError APIError
		json.Unmarshal([]byte(restext), &apiError)
		// return restext, fmt.Errorf("Error %d: %s", apiError.Error.ErrorCode, apiError.Error.ErrorText)
		if apiError.Error.ErrorCode == 0 {
			apiError.Error.ErrorCode = response.StatusCode()
		}

		code := apiError.Error.ErrorCode
		if code == http.StatusNotFound {
			return restext, fmt.Errorf("%s: %w", apiError.Error.ErrorText, ErrNotFound)
		}
		return restext, fmt.Errorf("Error %d: %s", apiError.Error.ErrorCode, apiError.Error.ErrorText)
	}

	// security api
	if strings.Contains(restext, "code") && strings.Contains(restext, "error") {
		var apiError SecurityAPIError
		err := json.Unmarshal([]byte(restext), &apiError)
		if err == nil {
			if apiError.ErrorCode == 0 {
				apiError.ErrorCode = response.StatusCode()
			}

			code := apiError.ErrorCode
			if code == http.StatusNotFound {
				return restext, fmt.Errorf("%s: %w", apiError.ErrorText, ErrNotFound)
			}
			return restext, fmt.Errorf("Error %d: %s", apiError.ErrorCode, apiError.ErrorText)
		}
	}

	// {"success":false,"error":"Zone field is missing or invalid","messages":[],"request_id":"d348a1d6-4f02-4c86-886b-0ef69fcbf696"}
	// dns api
	if strings.Contains(restext, "\"success\":false") && strings.Contains(restext, "error") {
		var apiError DnsAPIError
		err := json.Unmarshal([]byte(restext), &apiError)
		if err == nil {
			return restext, fmt.Errorf("Erro: %s", apiError.ErrorText)
		}
	}
	return restext, err
}

func (c *Client) Get(path string, params map[string]string) (string, error) {
	resp, err := c.createRequest(params, context.Background()).Get(c.Configs.APIEndpoint + "/" + path)
	c._logRequest(path, params, "GET", resp)
	return c.parseResponse(resp, err)
}

func (c *Client) Post(path string, params map[string]interface{}) (string, error) {
	resp, err := c.createRequest(nil, context.Background()).SetBody(params).Post(c.Configs.APIEndpoint + "/" + path)
	c._logRequest2(path, params, "POST", resp)
	return c.parseResponse(resp, err)
}

func (c *Client) Put(path string, params map[string]interface{}) (string, error) {
	resp, err := c.createRequest(nil, context.Background()).SetBody(params).Put(c.Configs.APIEndpoint + "/" + path)
	c._logRequest2(path, params, "PUT", resp)
	return c.parseResponse(resp, err)
}

func (c *Client) Patch(path string, params map[string]interface{}) (string, error) {
	resp, err := c.createRequest(nil, context.Background()).SetBody(params).Patch(c.Configs.APIEndpoint + "/" + path)
	c._logRequest2(path, params, "PATCH", resp)
	return c.parseResponse(resp, err)
}
func (c *Client) Delete(path string, params map[string]interface{}) (string, error) {
	resp, err := c.createRequest(nil, context.Background()).SetBody(params).Delete(c.Configs.APIEndpoint + "/" + path)
	c._logRequest2(path, params, "DELETE", resp)
	return c.parseResponse(resp, err)
}

func (c *Client) _logRequest(path string, params map[string]string, method string, response *resty.Response) {
	delete(params, "api_key")
	Logs(fmt.Sprintf("call api %s %s/%s params = %v, res = %s", method, c.Configs.APIEndpoint, path, convert2JsonString(params), response.String()))
}
func (c *Client) _logRequest2(path string, params map[string]interface{}, method string, response *resty.Response) {
	delete(params, "api_key")
	Logs(fmt.Sprintf("call api %s %s/%s params = %s, res = %s", method, c.Configs.APIEndpoint, path, convert2JsonString(params), response.String()))
}

type ActionResponse struct {
	Success bool `json:"success"`
}

func (c *Client) PerformDeleteWithBody(path string, params map[string]interface{}) (ActionResponse, error) {
	jsonStr, err := c.Delete(path, params)
	var res ActionResponse
	json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		return res, err
	}
	return res, err
}
func (c *Client) PerformDelete(path string) (ActionResponse, error) {
	jsonStr, err := c.Delete(path, map[string]interface{}{})
	var res ActionResponse
	json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *Client) PerformAction(path string, params map[string]interface{}) (ActionResponse, error) {
	jsonStr, err := c.Post(path, params)
	var res ActionResponse
	json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *Client) PerformUpdate(path string, params map[string]interface{}) (ActionResponse, error) {
	jsonStr, err := c.Put(path, params)
	var res ActionResponse
	json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *Client) PerformPatch(path string, params map[string]interface{}) (ActionResponse, error) {
	jsonStr, err := c.Patch(path, params)
	var res ActionResponse
	json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		return res, err
	}
	return res, err
}
func (c *Client) SimplePost(path string) (string, error) {
	return c.Post(path, map[string]interface{}{})
}

// TimeSettings object
// type TimeSettings struct {
// 	Delay    int
// 	Interval int
// 	Timeout  int
// }

// // ShortTimeSettings predefined TimeSettings for short task
// var ShortTimeSettings = TimeSettings{Delay: 1, Interval: 1, Timeout: 60}

// // MediumTimeSettings predefined TimeSettings for medium task
// var MediumTimeSettings = TimeSettings{Delay: 3, Interval: 3, Timeout: 5 * 60}

// // LongTimeSettings predefined TimeSettings for long task
// var LongTimeSettings = TimeSettings{Delay: 10, Interval: 20, Timeout: 20 * 60}

// // SuperLongTimeSettings predefined TimeSettings for long task
// var SuperLongTimeSettings = TimeSettings{Delay: 20, Interval: 20, Timeout: 5 * 60 * 60}

// // HalfDayTimeSettings for long task like take snapshot
// var HalfDayTimeSettings = TimeSettings{Delay: 60, Interval: 60, Timeout: 12 * 60 * 60}

// // OneDayTimeSettings for long task like take snapshot
// var OneDayTimeSettings = TimeSettings{Delay: 60, Interval: 60, Timeout: 24 * 60 * 60}

// func (c *Client) waitForTaskFinished(taskID string, timeSettings TimeSettings) (TaskStatus, error) {
// 	log.Printf("[INFO] Waiting for server with task id (%s) to be created", taskID)
// 	stateConf := &StateChangeConf{
// 		Pending:    []string{"WAIT", "PROCESSING"},
// 		Target:     []string{"DONE"},
// 		Refresh:    c.taskStateRefreshfunc(taskID),
// 		Timeout:    time.Duration(timeSettings.Timeout) * time.Second,
// 		Delay:      time.Duration(timeSettings.Delay) * time.Second,
// 		MinTimeout: time.Duration(timeSettings.Interval) * time.Second,
// 	}
// 	res, err := stateConf.WaitForState()
// 	if err != nil {
// 		return TaskStatus{}, err
// 	}
// 	return res.(TaskStatus), err
// }

// func (c *Client) taskStateRefreshfunc(taskID string) StateRefreshFunc {
// 	return func() (interface{}, string, error) {
// 		// Get task result from cloud server API
// 		resp, err := c.Task.Get(taskID)
// 		if err != nil {
// 			return nil, "", err
// 		}
// 		// if the task is not ready, we need to wait for a moment
// 		if resp.Status == "ERROR" {
// 			log.Println("[DEBUG] Task is failed")
// 			return nil, "", errors.New(fmt.Sprint(resp))
// 		}

// 		if resp.Status == "DONE" {
// 			return resp, "DONE", nil
// 		}

// 		log.Println("[DEBUG] Task is not done")
// 		return nil, "", nil
// 	}
// }
