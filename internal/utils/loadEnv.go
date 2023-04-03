package utils

import (
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirectory = "task_manager_api"

func LoadEnv() {
	projectname := regexp.MustCompile(`^(.*` + projectDirectory + `)`)
	curentWorkDir, _ := os.Getwd()
	rootPath := projectname.Find([]byte(curentWorkDir))

	err := godotenv.Load(string(rootPath) + `/ .env`)
	if err != nil {
		panic(err)
	}
}
