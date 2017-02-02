package main

import (
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/matthewmueller/store"
)

// App struct
type App struct {
	Name     string
	User     string
	Settings map[string]interface{}
}

func main() {
	log.SetHandler(text.New(os.Stderr))

	db, err := store.New("app")
	if err != nil {
		log.WithError(err).Fatal("unable to open config")
	}
	defer db.Close()

	err = db.Put("user", "matt")
	if err != nil {
		log.WithError(err).Fatal("unable to put config")
	}

	var v string
	err = db.Get("user", &v)
	if err != nil {
		log.WithError(err).Fatal("unable to get the value")
	}

	fmt.Println("got value", v)
}
