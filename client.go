package gocmcapiv2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

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
	Certificate              CertificateService
	DatabaseInstance         DatabaseInstanceService
	DatabaseBackup           DatabaseBackupService
	DatabaseAutoBackup       DatabaseAutoBackupService
	Kubernates               KubernatesService
	DatabaseConfiguration    DatabaseConfigurationService
	SecurityGroup            SecurityGroupService
	Keypair                  KeypairService
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
	c.Certificate = &certificate{client: c}
	c.Kubernates = &kubernetes{client: c}
	c.DatabaseInstance = &databaseinstance{client: c}
	c.DatabaseAutoBackup = &databaseautobackup{client: c}
	c.DatabaseConfiguration = &databaseconfiguration{client: c}
	c.SecurityGroup = &securitygroup{client: c}
	c.Keypair = &keypair{client: c}
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
	return restext, err
}

// Get Request, return resty Response
func (c *Client) Get(path string, params map[string]string) (string, error) {
	resp, err := c.createRequest(params, context.Background()).Get(c.Configs.APIEndpoint + "/" + path + ".json")
	c._logRequest(path, params, "GET", resp)
	return c.parseResponse(resp, err)
}

func (c *Client) Post(path string, params map[string]interface{}) (string, error) {
	resp, err := c.createRequest(nil, context.Background()).SetBody(params).Post(c.Configs.APIEndpoint + "/" + path + ".json")
	c._logRequest2(path, params, "POST", resp)
	return c.parseResponse(resp, err)
}

// Put request
func (c *Client) Put(path string, params map[string]interface{}) (string, error) {
	resp, err := c.createRequest(nil, context.Background()).SetBody(params).Put(c.Configs.APIEndpoint + "/" + path + ".json")
	c._logRequest2(path, params, "PUT", resp)
	return c.parseResponse(resp, err)
}

// Delete request

func (c *Client) Delete(path string, params map[string]string) (string, error) {
	resp, err := c.createRequest(params, context.Background()).Delete(c.Configs.APIEndpoint + "/" + path + ".json")
	c._logRequest(path, params, "DELETE", resp)
	return c.parseResponse(resp, err)
}

func (c *Client) _logRequest(path string, params map[string]string, method string, response *resty.Response) {
	// if len(params) > 1 {
	// 	Logo("call api GET " + c.Configs.APIEndpoint + "/" + path + " params = ", params)
	// }else{
	// 	Logs("call api GET " + c.Configs.APIEndpoint + "/" + path + " params = ", params)
	// }
	delete(params, "api_key")
	Logs(fmt.Sprintf("call api %s %s/%s params = %v, res = %s", method, c.Configs.APIEndpoint, path, convert2JsonString(params), response.String()))
	// Logs("res = " + response.String())
}
func (c *Client) _logRequest2(path string, params map[string]interface{}, method string, response *resty.Response) {
	// Logs("call api GET " + c.Configs.APIEndpoint + "/" + path + ".json")
	// if len(params) > 1 {
	// Logo("params = ", params)
	// }
	// Logs("res = " + response.String())
	delete(params, "api_key")
	Logs(fmt.Sprintf("call api %s %s/%s params = %s, res = %s", method, c.Configs.APIEndpoint, path, convert2JsonString(params), response.String()))
}

type ActionResponse struct {
	Success bool `json:"success"`
}

func (c *Client) PerformDelete(path string) (ActionResponse, error) {
	jsonStr, err := c.Delete(path, map[string]string{})
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

func (c *Client) SimplePost(path string) (string, error) {
	return c.Post(path, map[string]interface{}{})
}

// LongTask execute a action that return a task
func (c *Client) LongTask(action string, id string, params map[string]interface{}, timeSettings TimeSettings) (TaskStatus, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	if id != "" {
		params["id"] = id
	}

	jsonStr, err := c.Post(action, params)
	var task Task
	var taskResponse TaskStatus
	json.Unmarshal([]byte(jsonStr), &task)

	if err != nil {
		return taskResponse, err
	}
	taskResponse, err = c.waitForTaskFinished(task.TaskID, timeSettings)
	if err != nil {
		return taskResponse, fmt.Errorf("Error perform action %s: %s, params: %+v", action, err, params)
	}
	return taskResponse, err
}

// LongDeleteTask execute a action that return a task
func (c *Client) LongDeleteTask(action string, id string, params map[string]string, timeSettings TimeSettings) (TaskStatus, error) {
	if params == nil {
		params = make(map[string]string)
	}
	if id != "" {
		params["id"] = id
	}

	jsonStr, err := c.Delete(action, params)
	var task Task
	var taskResponse TaskStatus
	json.Unmarshal([]byte(jsonStr), &task)

	if err != nil {
		return taskResponse, err
	}
	taskResponse, err = c.waitForTaskFinished(task.TaskID, timeSettings)
	if err != nil {
		return taskResponse, fmt.Errorf("Error perform action %s: %s, params: %+v", action, err, params)
	}
	return taskResponse, err
}

// Order create an resource order
func (c *Client) Order(action string, id string, params map[string]interface{}, timeSettings TimeSettings) (OrderResponse, TaskStatus, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	if id != "" {
		params["id"] = id
	}

	jsonStr, err := c.Post(action, params)
	var order OrderResponse
	var taskStatus TaskStatus

	if err != nil {
		return order, taskStatus, fmt.Errorf("Error perform action %s: %s, params: %+v", action, err, params)
	}

	json.Unmarshal([]byte(jsonStr), &order)
	if !order.Paid {
		return order, taskStatus, fmt.Errorf("Error perform action %s cause order is not paid, input = %+v, response = %s", action, params, jsonStr)
		//errors.New("Can not perform this action cause of payment failed, connect to CMC administrator for your advice")
	}

	taskStatus, err = c.waitForTaskFinished(order.TaskID, timeSettings)
	if err != nil {
		return order, taskStatus, fmt.Errorf("Error perform action %s with task id (%s): %s", action, order.TaskID, err)
	}

	return order, taskStatus, err
}

// TimeSettings object
type TimeSettings struct {
	Delay    int
	Interval int
	Timeout  int
}

// ShortTimeSettings predefined TimeSettings for short task
var ShortTimeSettings = TimeSettings{Delay: 1, Interval: 1, Timeout: 60}

// MediumTimeSettings predefined TimeSettings for medium task
var MediumTimeSettings = TimeSettings{Delay: 3, Interval: 3, Timeout: 5 * 60}

// LongTimeSettings predefined TimeSettings for long task
var LongTimeSettings = TimeSettings{Delay: 10, Interval: 20, Timeout: 20 * 60}

// SuperLongTimeSettings predefined TimeSettings for long task
var SuperLongTimeSettings = TimeSettings{Delay: 20, Interval: 20, Timeout: 5 * 60 * 60}

// HalfDayTimeSettings for long task like take snapshot
var HalfDayTimeSettings = TimeSettings{Delay: 60, Interval: 60, Timeout: 12 * 60 * 60}

// OneDayTimeSettings for long task like take snapshot
var OneDayTimeSettings = TimeSettings{Delay: 60, Interval: 60, Timeout: 24 * 60 * 60}

func (c *Client) waitForTaskFinished(taskID string, timeSettings TimeSettings) (TaskStatus, error) {
	log.Printf("[INFO] Waiting for server with task id (%s) to be created", taskID)
	stateConf := &StateChangeConf{
		Pending:    []string{"WAIT", "PROCESSING"},
		Target:     []string{"DONE"},
		Refresh:    c.taskStateRefreshfunc(taskID),
		Timeout:    time.Duration(timeSettings.Timeout) * time.Second,
		Delay:      time.Duration(timeSettings.Delay) * time.Second,
		MinTimeout: time.Duration(timeSettings.Interval) * time.Second,
	}
	res, err := stateConf.WaitForState()
	if err != nil {
		return TaskStatus{}, err
	}
	return res.(TaskStatus), err
}

func (c *Client) taskStateRefreshfunc(taskID string) StateRefreshFunc {
	return func() (interface{}, string, error) {
		// Get task result from cloud server API
		resp, err := c.Task.Get(taskID)
		if err != nil {
			return nil, "", err
		}
		// if the task is not ready, we need to wait for a moment
		if resp.Status == "ERROR" {
			log.Println("[DEBUG] Task is failed")
			return nil, "", errors.New(fmt.Sprint(resp))
		}

		if resp.Status == "DONE" {
			return resp, "DONE", nil
		}

		log.Println("[DEBUG] Task is not done")
		return nil, "", nil
	}
}
