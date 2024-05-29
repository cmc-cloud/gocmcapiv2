package gocmcapiv2

import (
	"encoding/json"
)

// AutoScalingGroupService interface
type AutoScalingGroupService interface {
	Get(id string) (AutoScalingGroup, error)
	List(params map[string]string) ([]AutoScalingGroup, error)
	GetAction(id string) (AutoScalingAction, error)
	Create(params map[string]interface{}) (AutoScalingGroup, error)
	Update(id string, params map[string]interface{}) (ActionResponse, error)
	UpdateCapacity(id string, params map[string]interface{}) (ActionResponse, error)
	Delete(id string) (ActionResponse, error)
}

/*
	{
	    "created_at": "2022-09-18T16:19:27.000000",
	    "data": [],
	    "domain": null,
	    "id": "ddeb634e-42ec-4936-b3b3-87255209eaa2",
	    "name": "delete-policy-as-group-ggd8",
	    "project": "b613e96b87ce425f8cba5a0b549439ee",
	    "spec": {
	        "description": "41f88a92-4fe4-46a3-a724-b7fb51e8cf25",
	        "properties": {
	            "criteria": "OLDEST_FIRST",
	            "destroy_after_deletion": true,
	            "grace_period": 60,
	            "hooks": {
	                "params": {
	                    "url": "https://test.com"
	                },
	                "timeout": "3600",
	                "type": "webhook"
	            },
	            "reduce_desired_capacity": false,
	            "zones": [
	                {
	                    "name": "AZ1",
	                    "weight": 100
	                }
	            ],
	            "adjustment": {
	                "best_effort": true,
	                "cooldown": 300,
	                "min_step": 1,
	                "number": 1,
	                "type": "CHANGE_IN_CAPACITY"
	            },
	            "event": "CLUSTER_SCALE_IN",
	            "detection": {
	                "detection_modes": [
	                    {
	                        "type": "NODE_STATUS_POLLING"
	                    }
	                ],
	                "interval": 300,
	                "node_update_timeout": 60
	            },
	            "recovery": {
	                "actions": [
	                    {
	                        "name": "RECREATE"
	                    }
	                ],
	                "node_delete_timeout": 600,
	                "node_force_recreate": true
	            },
	            "health_monitor": {
	                "id": "66cca881-3b45-490f-93bf-f05e9f2938c1"
	            },
	            "lb_status_timeout": 300,
	            "loadbalancer": "9a38cdd9-bc27-42fe-aa02-19e31e427554",
	            "pool": {
	                "id": "bbc42901-0417-4e0d-a386-b9614f852d32",
	                "protocol_port": 443,
	                "subnet": "458731a9-9be4-4aa9-8dc7-41287aec76df"
	            },
	            "vip": {
	                "subnet": "458731a9-9be4-4aa9-8dc7-41287aec76df"
	            }
	        },
	        "type": "senlin.policy.deletion",
	        "version": "1.1"
	    },
	    "type": "senlin.policy.deletion-1.1",
	    "updated_at": null,
	    "user": "4e41dd85e4624341ba41e046c9654d2c"
	}
*/
type AutoScalingAction struct {
	Action       string  `json:"action"`
	Cause        string  `json:"cause"`
	CreatedAt    string  `json:"created_at"`
	EndTime      float64 `json:"end_time"`
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	StartTime    float64 `json:"start_time"`
	Status       string  `json:"status"`
	StatusReason string  `json:"status_reason"`
	Target       string  `json:"target"`
	Timeout      int     `json:"timeout"`
}

// AutoScalingGroup object
type AutoScalingGroup struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	CreatedAt       string   `json:"created_at"`
	DesiredCapacity int      `json:"desired_capacity"`
	MaxSize         int      `json:"max_size"`
	MinSize         int      `json:"min_size"`
	Nodes           []string `json:"nodes"`
	Policies        []string `json:"policies"`
	ProfileID       string   `json:"profile_id"`
	ProfileName     string   `json:"profile_name"`
	Status          string   `json:"status"`
	StatusReason    string   `json:"status_reason"`
	Timeout         int      `json:"timeout"`
	// Config          []string    `json:"config"`
	// Domain          string      `json:"domain"`
	// Metadata        []any     `json:"metadata"`
}
type autoscalinggroup struct {
	client *Client
}

// Get autoscalinggroup detail
func (s *autoscalinggroup) Get(id string) (AutoScalingGroup, error) {
	jsonStr, err := s.client.Get("as/group/"+id, map[string]string{})
	var obj AutoScalingGroup
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &obj)
	}
	return obj, err
}

func (s *autoscalinggroup) List(params map[string]string) ([]AutoScalingGroup, error) {
	restext, err := s.client.Get("as/group", params)
	items := make([]AutoScalingGroup, 0)
	err = json.Unmarshal([]byte(restext), &items)
	return items, err
}

func (s *autoscalinggroup) GetAction(id string) (AutoScalingAction, error) {
	jsonStr, err := s.client.Get("as/group/action/"+id, map[string]string{})
	var obj AutoScalingAction
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &obj)
	}
	return obj, err
}

// Delete a autoscalinggroup
func (s *autoscalinggroup) Delete(id string) (ActionResponse, error) {
	return s.client.PerformDelete("as/group/" + id)
}
func (s *autoscalinggroup) Update(id string, params map[string]interface{}) (ActionResponse, error) {
	return s.client.PerformUpdate("as/group/"+id, params)
}
func (s *autoscalinggroup) UpdateCapacity(id string, params map[string]interface{}) (ActionResponse, error) {
	return s.client.PerformUpdate("as/group/"+id+"/capacity", params)
}

// Create a new autoscalinggroup
func (s *autoscalinggroup) Create(params map[string]interface{}) (AutoScalingGroup, error) {
	jsonStr, err := s.client.Post("as/group", params)
	var response AutoScalingGroup
	if err != nil {
		return response, err
	}
	json.Unmarshal([]byte(jsonStr), &response)
	return response, nil
}
