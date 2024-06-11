package gocmcapiv2

import (
	"encoding/json"
)

// RedisBackupService interface
type RedisBackupService interface {
	Get(id string) (RedisBackup, error)
	List(params map[string]string) ([]RedisBackup, error)
	Delete(id string) (ActionResponse, error)
	Update(id string, params map[string]interface{}) (ActionResponse, error)
}

type RedisSnapshot struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	InstanceID string  `json:"instance_id"`
	Created    string  `json:"created"`
	Size       float64 `json:"size"`
	Status     string  `json:"status"`
	RealSize   int     `json:"real_size"`
}
type RedisBackup struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	InstanceID string  `json:"instance_id"`
	ParentID   string  `json:"parent_id"`
	Created    string  `json:"created"`
	Size       float64 `json:"size"`
	Status     string  `json:"status"`
	RealSize   int     `json:"real_size"`
	RealSizeGB float64 `json:"real_size_gb"`
}

type redisbackup struct {
	client *Client
}

// Get redisbackup detail
func (v *redisbackup) Get(id string) (RedisBackup, error) {
	jsonStr, err := v.client.Get("dbaas/instance/"+id, map[string]string{})
	var obj RedisBackup
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), &obj)
	}
	return obj, err
}

func (s *redisbackup) List(params map[string]string) ([]RedisBackup, error) {
	restext, err := s.client.Get("dbaas/backup", params)
	items := make([]RedisBackup, 0)
	err = json.Unmarshal([]byte(restext), &items)
	return items, err
}

// Delete a redisbackup
func (v *redisbackup) Delete(id string) (ActionResponse, error) {
	return v.client.PerformDelete("dbaas/instance/" + id)
}
func (v *redisbackup) Update(id string, params map[string]interface{}) (ActionResponse, error) {
	return v.client.PerformUpdate("dbaas/instance/"+id, params)
}
func (v *redisbackup) Resize(id string, flavor_id string) (ActionResponse, error) {
	return v.client.PerformAction("dbaas/instance/"+id+"/resize", map[string]interface{}{"flavor_id": flavor_id})
}
func (v *redisbackup) ResizeVolume(id string, volume_size int) (ActionResponse, error) {
	return v.client.PerformAction("dbaas/instance/"+id+"/resize_volume", map[string]interface{}{"size": volume_size})
}
func (v *redisbackup) UpdateInstanceAccessbility(id string, params map[string]interface{}) (ActionResponse, error) {
	return v.client.PerformUpdate("dbaas/instance/"+id+"/accessbility", params)
}
func (v *redisbackup) UpgradeDatastoreVersion(id string, datastore_version string) (ActionResponse, error) {
	return v.client.PerformUpdate("dbaas/instance/"+id+"/upgrade_datastore_version", map[string]interface{}{"datastore_version": datastore_version})
}

func (s *redisbackup) Create(params map[string]interface{}) (RedisBackup, error) {
	jsonStr, err := s.client.Post("dbaas/instance", params)
	var response RedisBackup
	if err != nil {
		return response, err
	}
	json.Unmarshal([]byte(jsonStr), &response)
	return response, nil
}
