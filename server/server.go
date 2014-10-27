package server

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

var listener net.Listener

func Run() error {
	http.HandleFunc("/", handler)
	var err error
	listener, err = net.Listen("tcp", "localhost:0")
	if err != nil {
		return err
	}
	fmt.Println("Port", listener.Addr().(*net.TCPAddr).Port)
	if err := http.Serve(listener, nil); err != nil {
		if netOpErr, ok := err.(*net.OpError); ok && netOpErr.Op == "accept" {
			return nil
		}
		return err
	}
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ok.")
	fmt.Println("Ok.")
	go killServer()
}

var killServerMutex = new(sync.Mutex)
var killServerClosed = false

func killServer() {
	time.Sleep(3 * time.Second)
	killServerMutex.Lock()
	defer killServerMutex.Unlock()
	if killServerClosed {
		return
	}
	killServerClosed = true
	listener.Close()
}
