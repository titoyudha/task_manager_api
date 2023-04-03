package main

import (
	"github.com/titoyudha/task_manager_api/internal/config"
	"github.com/titoyudha/task_manager_api/internal/server"
)

func main() {
	config.ConnectDB()
	server.Run()
}
