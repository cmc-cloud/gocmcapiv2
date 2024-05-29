package gocmcapiv2

import (
	"encoding/json"
	"fmt"
)

// TaskService interface
type TaskService interface {
	Get(uuid string) (TaskStatus, error)
}

type task struct {
	client *Client
}

// Task response when get task status
type Task struct {
	TaskID string `json:"jobid"`
}

// TaskStatus response when get task status
type TaskStatus struct {
	Command   string `json:"command"`
	Status    string `json:"status"`
	ResultID  string `json:"id"`
	ErrorText string `json:"error_text"`
}

// Get task status
func (t *task) Get(uuid string) (TaskStatus, error) {
	jsonStr, err := t.client.Get("job/status", map[string]string{"id": uuid})
	fmt.Println("jsonStr, err", jsonStr, err)
	var task TaskStatus
	err = json.Unmarshal([]byte(jsonStr), &task)
	return task, err
}
