package httphandling

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Jiibben/GITHUB_APP/apistructures"
)

func FindDir(client *http.Client, name string) (apistructures.Repo, error) {
	request := GetRequestRepo()

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range ParseUserReposJson(ReadResponse(resp)) {
		if strings.EqualFold(name, v.Name) {
			return v, nil
		}
	}
	return apistructures.Repo{}, errors.New("Directory of name " + name + " not found")
}

func DeleteRepo(client *http.Client, repoFullName string) error {
	request := DeleteRequestDeleteRepo(repoFullName)
	_, err := client.Do(request)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func CreateRepo(client *http.Client, name string) (apistructures.Repo, error) {
	resp, err := client.Do(PostRequestCreateRepo(map[string]string{
		"name":        name,
		"description": "This is my project",
		"homepage":    "https://github.com",
		"private":     "true",
	}))
	if err != nil {
		return apistructures.Repo{}, err
	} else {
		return ParseUniqueRepo(ReadResponse(resp)), nil
	}

}

func ListRepos(client *http.Client) ([]apistructures.Repo, error) {
	request := GetRequestRepo()
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return []apistructures.Repo{}, errors.New("didn't find any repo")
	} else {
		return ParseUserReposJson(ReadResponse(resp)), nil
	}
}
