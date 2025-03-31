package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

var version string

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "Nieznany"
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "Nieznany"
}

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	ip := getIP()

	fmt.Fprintf(w, "Adres IP serwera: %s\n", ip)
	fmt.Fprintf(w, "Nazwa serwera: %s\n", hostname)
	fmt.Fprintf(w, "Wersja aplikacji: %s\n", version)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Serwer działa na porcie 8080...")
	http.ListenAndServe(":8080", nil)
}
