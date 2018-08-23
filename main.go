package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

var port string

func init() {
	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		log.Fatal(fmt.Errorf("$PORT not set"))
	}
	port = ":" + portEnv
}

func main() {
	http.HandleFunc("/", sayIP)
	log.Printf("Listening on %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

func GetIP(req *http.Request) string {
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		addresses := strings.Split(req.Header.Get(h), ",")
		for i := len(addresses) - 1; i >= 0; i-- {
			ip := strings.TrimSpace(addresses[i])
			realIP := net.ParseIP(ip)
			if !realIP.IsGlobalUnicast() || IsPrivateSubnet(realIP) {
				continue
			}
			return ip
		}
	}
	return ""
}

func sayIP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, GetIP(req))
}
