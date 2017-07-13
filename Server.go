// Live Screen Capture project main.go

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var imageDat string = ""

func Server() {
	http.HandleFunc("/connws/", ConnWs)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func ConnWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}
	res := map[string]interface{}{}
	for {
		if err = ws.ReadJSON(&res); err != nil {
			if err.Error() == "EOF" {
				return
			}
			if err.Error() == "unexpected EOF" {
				return
			}
			fmt.Println("Read : " + err.Error())
			return
		}
		res["a"] = "a"
		log.Println(res)
		for {
			res["img64"] = imageDat
			if err = ws.WriteJSON(&res); err != nil {
				fmt.Println("watch dir - Write : " + err.Error())
			}
			time.Sleep(5 * time.Millisecond)
		}
	}
}

func main() {
	listen := flag.Int("listen", 8080, "Port you want to listen on.")
	flag.Parse()
	fmt.Println("Live Screen Capture Server")
	go Server()
	ln, _ := net.Listen("tcp", ":"+strconv.Itoa(*listen))
	fmt.Println("Listening on port: " + strconv.Itoa(*listen))
	conn, _ := ln.Accept()
	fmt.Println("Connected to", conn.LocalAddr().String())
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if len(message) >= 1 {
			if string(message) != imageDat {
				imageDat = string(message)
			}
		} else {
			fmt.Println("Connection to client lost.")
			os.Exit(0)
		}
	}
}
