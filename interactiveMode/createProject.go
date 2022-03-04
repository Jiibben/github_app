package interactiveMode

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/Jiibben/GITHUB_APP/httphandling"
)

func createDirectory(path string, name string) {
	os.Chdir(path)
	os.Mkdir(name, os.FileMode(0777))
	os.Chdir(path + string(os.PathSeparator) + name)

}

func CreateEmptyProject(client *http.Client, path string, name string) {

	os.Chdir(path)
	repo, err := httphandling.CreateRepo(client, name)
	if err != nil {
		fmt.Println(err)
	}

	exec.Command("git", "clone", repo.Ssh_url).Run()
}

func MakeProjectFromCurrentDir(client *http.Client, path string, name string) {
	os.Chdir(path)
	repo, _ := httphandling.CreateRepo(client, name)
	exec.Command("git", "init").Run()
	exec.Command("git", "add", ".").Run()
	exec.Command("git", "commit", "-m", "first commit").Run()
	exec.Command("git", "branch", "-M", "main").Run()
	exec.Command("git", "remote", "add", "origin", repo.Ssh_url).Run()
	exec.Command("git", "push", "-u", "origin", "main").Run()
}
