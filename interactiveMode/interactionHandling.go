package interactiveMode

import (
	"fmt"
	"net/http"
	"strings"
)

func Run(client *http.Client) {
	var userCommand string
	running := true
	for running {
		fmt.Print("MYSUPERGIT >>> ")
		fmt.Scanln(&userCommand)

		switch strings.ToLower(userCommand) {
		case "ls":
			LS(client)
		case "rm":
			RM(client)
		case "quit":
			fmt.Println("Quitting ... good bye <3")
			running = false
		case "mkep":
			MKEP(client)
		case "mkcw":
			MKCW(client)
		default:
			fmt.Println("Unrecognised command  ", userCommand)
		}

	}
}
