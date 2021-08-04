package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"
)

type configuration struct {
	ServerInfo serverInfo
	RedisInfo  redisInfo
}

type serverInfo struct {
	Host string
}

type redisInfo struct {
	Host        string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var Configuration = configuration{}

func init() {
	filePath := path.Join(os.Getenv("GOPATH"), "src/apiproject/config/config.json")
	fmt.Println("filePath",filePath)
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Printf("Open file error: %v\n", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Configuration)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
