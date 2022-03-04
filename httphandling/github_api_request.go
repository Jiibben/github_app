package httphandling

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const TOKEN = "token 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"

var baseHeader map[string]string = map[string]string{
	"Authorization": TOKEN,
	"Accept":        "application/json"}

func addHeaders(headers map[string]string, req *http.Request) {
	for k, v := range headers {
		req.Header.Add(k, v)
	}
}

func GetRequestRepo() *http.Request {
	req, err := http.NewRequest("GET", "https://api.github.com/user/repos", nil)
	if err != nil {
		fmt.Println("error while creating request to get repo", err)
	}
	addHeaders(baseHeader, req)
	return req
}

func DeleteRequestDeleteRepo(fullname string) *http.Request {

	req, err := http.NewRequest("DELETE", "https://api.github.com/repos/"+fullname, nil)
	if err != nil {
		fmt.Println("error", err)
	}
	addHeaders(baseHeader, req)
	return req
}

func PostRequestCreateRepo(repoDetails map[string]string) *http.Request {
	repoDetailJson, ok := json.Marshal(repoDetails)
	if ok != nil {
		fmt.Println(ok)
	}
	req, err := http.NewRequest("POST", "https://api.github.com/user/repos", bytes.NewBuffer(repoDetailJson))
	if err != nil {
		fmt.Println(req)
	}
	addHeaders(baseHeader, req)
	return req
}
