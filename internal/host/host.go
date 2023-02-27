package host

import (
	"log"
	"net"
)

// FindCNAME finds the CNAMES associated for a host
func FindCNAME(host string) error {
	cname, err := net.LookupCNAME(host)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(cname)
	return nil
}

// FindTXTRecords finds the TXT records associated for a host
func FindTXTRecords(host string) error {
	txtRecords, err := net.LookupTXT(host)
	if err != nil {
		log.Println(err)
		return err
	}
	for _, txtRecord := range txtRecords {
		log.Println(txtRecord)
	}
	return nil
}

// FindIPAddresses finds the IP addresses associated for a host
func FindIPAddresses(host string) error {
	ip, err := net.LookupIP(host)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, ipAddr := range ip {
		log.Println(ipAddr)
	}
	return nil
}

// FindNameServers finds the nameservers associated with a host
func FindNameServers(host string) error {
	nameServers, err := net.LookupNS(host)
	if err != nil {
		log.Println(err)
		return err
	}
	for _, nameServer := range nameServers {
		log.Println(nameServer.Host)
	}
	return nil
}

// FindMXRecords finds the MX records associated with a host
func FindMXRecords(host string) error {
	mx, err := net.LookupMX(host)
	if err != nil {
		log.Println(err)
		return err
	}
	for _, mxRecord := range mx {
		log.Println(mxRecord.Host, mxRecord.Pref)
	}
	return nil
}
