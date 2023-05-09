package service

import (
	"github/LordRadamanthys/cmd_golang/cmd/repository"
)

func CreateRepositoryGitlabService(name, description string) {

	respCreate := repository.CreateRepository(name, description)

	repository.CreateBranch(respCreate.ID)
}
