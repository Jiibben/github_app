package apistructures

type Repo struct {
	ID         int    `json:"id"`
	Visibility string `json:"visibility"`
	Name       string `json:"name"`
	Full_name  string `json:"full_name`
	Ssh_url    string `json:"ssh_url`
	Owner      User   `json:"owner`
}
