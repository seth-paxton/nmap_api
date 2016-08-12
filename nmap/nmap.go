package nmap

import (
	"bytes"
	"log"
	"os/exec"
)

// LookPath looks for nmap in the local path
func LookPath() {
	path, err := exec.LookPath("nmap")
	if err != nil {
		log.Fatal("Please install Nmap using your package manager")
	}
	log.Printf("Nmap executable found in %s", path)

}

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