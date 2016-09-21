package main

import (
	"net/http"
	"log"
	"os/exec"
	
	"github.com/caarlos0/env"
	"github.com/seth-paxton/nmap_api/router"
)

// ENV variables for server configuration
type Config struct {
	Port string `env:"API_PORT" envDefault:"8080"`
}

// Start the web service
func main() {
	LookPath()
	cfg := Config{}
	env.Parse(&cfg)
	//Load up all the URLs
	router.ScanApi()
	err := http.ListenAndServe(":" + cfg.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Checks to see if Nmap is installed
func LookPath() {
	path, err := exec.LookPath("nmap")
	if err != nil {
		log.Fatal("Please install Nmap using your package manager")
	}
	log.Printf("Nmap executable found in %s", path)

}



