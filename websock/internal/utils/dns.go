package utils

import (
	"errors"
	"log"
	"net"
)

func ResolveHeadlessServiceDNS(address string, service string) []string {
	if len(address) == 0 {
		log.Fatalln(errors.New("invalid source address"))
	}
	ips, err := net.LookupIP(address)
	var res []string
	if err != nil {
		log.Fatalln(errors.New("cant resolve service IPs"))
	}
	for _, val := range ips {
		if service == "kafka" {
			res = append(res, val.String()+":9092")
		} else if service == "redis" {
			res = append(res, val.String()+":6379")
		} else {
			log.Fatalln(errors.New("unidentified service"))
		}
	}
	return res
}
