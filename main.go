package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {

	http.HandleFunc("/hello/sendudp", echoPayload)
	log.Printf("Go Backend: { HTTPVersion = 1 }; serving on http://localhost:9191/hello/sendudp")
	log.Fatal(http.ListenAndServe(":9191", nil))
}

func echoPayload(w http.ResponseWriter, req *http.Request) {

	log.Printf("Request connection: %s, path: %s", req.Proto, req.URL.Path[1:])
	defer req.Body.Close()

	q := req.URL.Query()

	if (!q.Has("ip")) || (!q.Has("text")) {
		fmt.Fprintf(w, "Query string dont have ?ip=... param %s\n", string(req.URL.RawPath))
		//fmt.Fprintln(w, "Request connections http://urlservice:9191/hello/sendudp?ip=ipadress(hexdecimal format)&text=text(hexdecimalformat)")
		return
	}

	qiphex := q.Get("ip")
	qtexthex := q.Get("text")

	qip, err := hex.DecodeString(qiphex)
	if err != nil {
		log.Printf("error: %s", err)
		log.Printf("Send qip: %s", string(qiphex))
		return
	}

	qtext, err := hex.DecodeString(qtexthex)
	if err != nil {
		log.Printf("error: %s", err)
		log.Printf("Send qtext: %s", string(qtexthex))
		return
	}
	log.Printf("Send: %s, path: %s", string(qtext), string(qip))

	go sendToUDP(qip, qtext)

	//	fmt.Fprintf(w, "%s\n", string(contents))
	//	fmt.Fprintf(w, "%s\n", string(qip))
	//	fmt.Fprintf(w, "%s\n", string(qtext))

}

func sendToUDP(ip []byte, text []byte) {
	conn, err := net.Dial("udp", string(ip))
	if err != nil {
		fmt.Printf("Some error %v", err)
		log.Printf("Some error %v", err)
		return
	}
	//	sbyte := []byte(s)
	_, err = conn.Write(text)
	if err != nil {
		fmt.Printf("Some error %v", err)
		log.Printf("Some error %v", err)
		return
	}

	conn.Close()
}
