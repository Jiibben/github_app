package main

import (
	"net/http"

	"github.com/Jiibben/GITHUB_APP/interactiveMode"
)

func main() {

	client := &http.Client{}

	interactiveMode.Run(client)

}
