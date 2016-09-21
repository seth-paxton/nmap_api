package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/seth-paxton/nmap_api/scanner"
)

type scanOptions struct {
	Scantype string `json:"Scantype"`
	Network  string `json:"Network"`
}

// URLs exposed to the user. 
func ScanApi() {
    http.HandleFunc("/scan", scan)
}

func scan(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t scanOptions
	scantypes := []string{"ping", "udp"}
	err := decoder.Decode(&t)
	if err != nil {
		log.Println("error")
	}
	for _, stype := range scantypes {
		if stype == t.Scantype {
			data, _ := json.Marshal("{'success':'Scan has been submitted'}")
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Write(data)
			log.Printf("%s scan has been submitted against %s", t.Scantype, t.Network)
			scanner.Run(t.Network, stype)
			break
		} else {
			data, _ := json.Marshal("{'error':'Invalid scan type'}")
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Write(data)
			log.Printf("%s scan submitted as invalid", t.Scantype)
		}
	}
}