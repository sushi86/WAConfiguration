package server

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/sushi86/WAConfiguration/command"
)

type message struct {
	Step string `json:"step"`
	Data string `json:"data"`
}

func Start() {
	fs := http.FileServer(http.Dir("./server/static"))
	http.Handle("/", fs)
	http.HandleFunc("/api/", api)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("unable to start server process")
	}
}

func api(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg message

	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if msg.Step == "step1" {
		command.Register(msg.Data, "sms")
	}
	
}
