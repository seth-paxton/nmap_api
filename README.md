# nmap_api

## Description
Rest API around NMAP written in Go. Wanted to learn how to write restful services in GoLang. This is a work in progress. 

## Usage
To run an NMAP scan, run the following command.

```
curl -H "Content-Type: application/json" -X POST -d '{"network": "192.168.1.2","Scantype": "ping"}' http://localhost:8080/scan
```

## Architecture

[High Level Architecture](docs/architecture/README.md)