package command

import (
	"os/exec"
	"log"
)

func Register(number string, channel string) {
	out, err := exec.Command("yowsup-cli", "").Output()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(out))
}

func Activate(number string, password string) {

}

func SendMessage(number string, receipient string, message string) {

}
