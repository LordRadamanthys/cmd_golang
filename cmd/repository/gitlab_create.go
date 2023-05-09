package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github/LordRadamanthys/cmd_golang/cmd/repository/model/request"
	"github/LordRadamanthys/cmd_golang/cmd/repository/model/response"
	"io/ioutil"
	"net/http"
	"os"
)

func CreateRepository(name, description string) *response.CreateResponse {
	var TOKEN = os.Getenv("TOKEN")

	var URL = os.Getenv("GITLAB_URL")
	fmt.Println(TOKEN, URL)
	postBody, _ := json.Marshal(createRequest(name, description))
	responseBody := bytes.NewBuffer(postBody)

	req, _ := http.NewRequest(http.MethodPost, URL, responseBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("PRIVATE-TOKEN", TOKEN)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %v \n", err)
		os.Exit(1)
	}

	if res.StatusCode != 201 {
		fmt.Printf("client: error to create repo, code: %v \n", res.StatusCode)
		os.Exit(1)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("readAll: error to read response: %s\n", err)
		os.Exit(1)
	}
	var responseCreate response.CreateResponse
	json.Unmarshal(resBody, &responseCreate)
	fmt.Printf("CreateRepository - status: success, response: %v \n", &responseCreate)
	return &responseCreate
}

func createRequest(name, description string) *request.CreateRequest {
	return &request.CreateRequest{
		Name:        name,
		Description: description,
		Readme:      true,
	}
}
