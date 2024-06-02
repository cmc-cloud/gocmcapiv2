package gocmcapiv2

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Logo log object
func Logo(pre string, object interface{}) {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	// logger.Println(object)
	jsonString, err := json.Marshal(object)
	if err != nil {
		logger.Println("Error:", err)
		return
	}

	// Print JSON string
	logger.Println(pre + string(jsonString))
}

func convert2JsonString(object interface{}) string {
	jsonString, err := json.Marshal(object)
	if err != nil {
		fmt.Printf("err %v", err)
		return ""
	}
	return string(jsonString)
}

// Logs log string
func Logs(message string) {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Println(message)
}
