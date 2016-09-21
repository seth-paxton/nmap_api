package scanner

import (
	"bytes"
	"log"
    "os/exec"
)

// Run is the logic around the scan itself
func Run(net string, scan string) {
	var out bytes.Buffer
	if scan == "ping" {
		cmd := exec.Command("nmap", "-sn", net)
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		
	} else if scan == "udp" {
		cmd := exec.Command("nmap", "-sU", net)
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
		log.Printf("scan results %q\n", out.String())
}
