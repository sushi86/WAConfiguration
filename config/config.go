package config

import (
	"io/ioutil"
	"fmt"
)

var filePath = "config.py"

func IsValidConfigFile() bool {
	_, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false;
	} else {
		return true
	}
}

func ReadFile() {
	dat, err := ioutil.ReadFile("config.py")
	if err != nil {
		fmt.Println("file does not exist")
	} else {
		fmt.Println(string(dat))
	}

}
