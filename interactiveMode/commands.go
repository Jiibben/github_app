package interactiveMode

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Jiibben/GITHUB_APP/httphandling"
)

func LS(client *http.Client) {
	repos, _ := httphandling.ListRepos(client)
	for _, v := range repos {
		fmt.Println("-> ", v.Name)
	}
}

func RM(client *http.Client) {
	var repoName string
	fmt.Print("\t Enter name of repository to delete : ")
	fmt.Scanln(&repoName)
	a, b := httphandling.FindDir(client, repoName)
	if b != nil {
		fmt.Println("Something happened couldn't delete repo with name ", repoName)
	} else {

		var accept string
		fmt.Print("\t type -> Yes <- if you wish to delete repo ", repoName, ": ")
		fmt.Scanln(&accept)

		if accept == "Yes" {
			err := httphandling.DeleteRepo(client, a.Full_name)
			if err != nil {
				fmt.Println("Couldn't delete given repository ", repoName)
			} else {
				fmt.Println("Successfully deleted repository ", repoName)
			}
		} else {
			fmt.Println("aborting ...")
		}
	}
}

func MKEP(client *http.Client) {
	var newRepoName string
	var path string
	fmt.Print("\t Enter a name for the new repository: ")
	fmt.Scanln(&newRepoName)
	fmt.Print("\t Enter a path for the new repository (leave blank for cwd): ")
	fmt.Scanln(&path)

	if path == "" {
		path, _ = os.Getwd()
	}

	CreateEmptyProject(client, path, newRepoName)
}

func MKCW(client *http.Client) {
	var newRepoName string
	var path string
	fmt.Print("\t Enter a name for the new repository: ")
	fmt.Scanln(&newRepoName)
	fmt.Print("\t Enter the path to commit to the new repository (leave blank for cwd): K")
	fmt.Scanln(&path)

	if path == "" {
		path, _ = os.Getwd()
	}
	MakeProjectFromCurrentDir(client, path, newRepoName)

}
