package repository

import (
	"encoding/json"
	"fmt"
	"github/LordRadamanthys/cmd_golang/cmd/repository/model/response"
	"io/ioutil"
	"net/http"
	"os"
)

func CreateBranch(id int) {
	var TOKEN = os.Getenv("TOKEN")
	var URL = os.Getenv("GITLAB_URL")

	path := fmt.Sprintf("%s/%v%s", URL, id, "/repository/branches?branch=develop&ref=main")
	fmt.Println(path)
	req, _ := http.NewRequest(http.MethodPost, path, nil)
	req.Header.Add("PRIVATE-TOKEN", TOKEN)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client branch: error making http request: %v \n", err)
		os.Exit(1)
	}

	if res.StatusCode != 201 {
		fmt.Printf("client branch: error to create branch, code: %v \n", res.StatusCode)
		os.Exit(1)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("readAll branch: error to read response: %s\n", err)
		os.Exit(1)
	}
	var responseBranch response.CreateBranch
	json.Unmarshal(resBody, &responseBranch)
	fmt.Printf("CreateBranch - status: success, response: %v \n", &responseBranch)
}
