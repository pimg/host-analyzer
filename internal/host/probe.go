package host

import (
	"crypto/tls"
	"log"
	"net/http"
	"strconv"

	"github.com/pimg/host-analyzer/internal/constants"
	probing "github.com/prometheus-community/pro-bing"
)

// PingHost performs an ICMP ping on a particular host
func PingHost(host string) error {
	pinger, err := probing.NewPinger(host)
	if err != nil {
		log.Println(err)
		return err
	}
	pinger.Count = 3
	pinger.Run()
	stats := pinger.Statistics()
	if stats.PacketsRecv > 0 {
		log.Println("ICMP enabled")
	}

	return nil
}

// ProbeTLS probes a host for information regarding the TLS configuration, like TLS version and supported Cipher suites.
func ProbeTLS(host string, port int) error {

	for tlsVersion, tlsVersionString := range constants.TLSVersions {

		conf := &tls.Config{
			MaxVersion:               tlsVersion,
			PreferServerCipherSuites: false,
		}
		conn, err := tls.Dial("tcp", host+":"+strconv.Itoa(port), conf)
		if err != nil {
			log.Println(tlsVersionString + " is unsupported.")
		} else {
			log.Println(tlsVersionString + " is supported")
			conn.Close()
		}

	}
	return nil
}

// ProbeHTTP probes a host for information regarding the HTTP protocol
func ProbeHTTP(host string) error {

	//TODO fix hardcoded protocol prefix.
	resp, err := http.Get("https://" + host)
	if err != nil {
		log.Println(err)
		return err
	}

	defer resp.Body.Close()

	log.Printf("HTTP version: %s\n", resp.Proto)
	log.Printf("HTTP response status: %s\n", resp.Status)
	log.Printf("The default Cipher suite is: %s\n", constants.CipherSuites[resp.TLS.CipherSuite])
	log.Printf("The default TLS version is: %s\n", constants.TLSVersions[resp.TLS.Version])
	return nil
}
