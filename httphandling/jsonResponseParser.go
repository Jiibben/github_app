package httphandling

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Jiibben/GITHUB_APP/apistructures"
)

func ReadResponse(resp *http.Response) []byte {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error occured when reading body response", err)

	}
	return body

}

func ParseUserReposJson(body []byte) (decodedJson []apistructures.Repo) {
	decodedJson = []apistructures.Repo{}
	err := json.Unmarshal(body, &decodedJson)
	if err != nil {
		fmt.Println("error while parsing users repo json", err)
	}
	return

}

func ParseUniqueRepo(body []byte) (dJson apistructures.Repo) {
	dJson = apistructures.Repo{}
	err := json.Unmarshal(body, &dJson)
	if err != nil {
		fmt.Println("error while parsing repo")
	}
	return
}
