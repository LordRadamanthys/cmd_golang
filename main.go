package main

import (
	createrepository "github/LordRadamanthys/cmd_golang/cmd/create_repository"
	"github/LordRadamanthys/cmd_golang/config"
)

func main() {
	config.LoadConfig()
	createrepository.Execute()
}
