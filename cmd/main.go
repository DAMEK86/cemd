package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/DerAndereAndi/eebus-go-cem/cem"
)

// main app
func usage() {
	fmt.Println("Usage: go run /cmd/main.go <serverport> <evse-ski> <crtfile> <keyfile> <iface>")
}

func main() {
	if len(os.Args) < 5 {
		usage()
		return
	}

	h := cem.NewCEM("Demo", "HEMS", "123456789", "Demo-HEMS-0123456789")
	ifaces := []string{os.Args[5]}
	if err := h.Setup(os.Args[1], os.Args[2], os.Args[3], os.Args[4], ifaces); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Clean exit to make sure mdns shutdown is invoked
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	// User exit
}
