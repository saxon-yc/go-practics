package main

import (
	"go-practics/config"
	structer "go-practics/internal/structer"
)

func init() {
	config.New("config/proxy.yaml")
}

func main() {
	// structer.Base()
	// structer.Array()
	structer.Pointer()
	structer.Mapp()
}
