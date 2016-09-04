package main

import (
	"fmt"
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
)

func bob(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bob\n%s", r.UserAgent())
}

func main() {
	conf := rice.Config{
		LocateOrder: []rice.LocateMethod{rice.LocateAppended},
	}
	box, err := conf.FindBox("example-files")
	if err != nil {
		log.Fatalf("error opening rice.Box: %s\n", err)
	}

	http.HandleFunc("/bob", bob)
	http.Handle("/", http.FileServer(box.HTTPBox()))
	go func() {
		fmt.Println("Serving files on :8080, press ctrl-C to exit")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatalf("error serving files: %v", err)
		}
	}()

	select {}
}
