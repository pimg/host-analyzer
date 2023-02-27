package host

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProbeTLS_known(t *testing.T) {
	host := "apily.dev"
	output := captureLog(func() {
		ProbeTLS(host, 443)
	})
	assert.Contains(t, output, "VersionTLS13 is supported\n")
}

func TestProbeTLS_unknown(t *testing.T) {
	host := "idonotexist.iamunknown"
	output := captureLog(func() {
		ProbeTLS(host, 443)
	})
	assert.Contains(t, output, "no such host")
}

func TestProbeHttp_known(t *testing.T) {
	host := "apily.dev"
	expected := "HTTP version: HTTP/2.0\nHTTP response status: 200 OK\nThe default Cipher suite is: TLS_AES_256_GCM_SHA384\nThe default TLS version is: VersionTLS13\n"
	output := captureLog(func() {
		ProbeHTTP(host)
	})
	assert.Contains(t, output, expected)
}

func TestProbeHttp_unknown(t *testing.T) {
	host := "idonotexist.iamunknown"
	output := captureLog(func() {
		ProbeHTTP(host)
	})
	assert.Contains(t, output, "no such host")
}
