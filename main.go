package main

import (
	fmt "fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/sendudp", echoPayload)
	log.Printf("Go Backend: { HTTPVersion = 1 }; serving on https://localhost:9191/hello/sendudp")
	log.Fatal(http.ListenAndServe(":9191", nil))
}

func echoPayload(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request connection: %s, path: %s", req.Proto, req.URL.Path[1:])
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf(
				"Oops! Failed close connection.\n %s",
				err,
			)
		}
	}(req.Body)

	var contents, err = io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalf("Oops! Failed reading body of the request.\n %s", err)
	}
	_, err = fmt.Fprintf(w, "%s\n", string(contents))
	if err != nil {
		return
	}
}
