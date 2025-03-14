package host

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureLog(f func()) string {
	log.SetFlags(0)
	var buf bytes.Buffer

	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func TestFindCNAME_known(t *testing.T) {
	host := "apily.dev"
	output := captureLog(func() {
		FindCNAME(host)
	})
	assert.Equal(t, output, "apily.dev.\n")
}

func TestFindCNAME_localhost(t *testing.T) {
	host := "localhost"
	output := captureLog(func() {
		FindCNAME(host)
	})
	assert.Contains(t, output, "localhost")
}

func TestFindCNAME_unknown(t *testing.T) {
	host := "idonotexist.iamunknown"
	output := captureLog(func() {
		FindCNAME(host)
	})
	assert.Contains(t, output, "no such host")
}

func TestFindTXTRecords_known(t *testing.T) {
	host := "google.com"
	output := captureLog(func() {
		FindTXTRecords(host)
	})
	assert.Contains(t, output, "MS=")
}

func TestFindIPAddresses_known(t *testing.T) {
	host := "apily.dev"
	output := captureLog(func() {
		FindIPAddresses(host)
	})
	assert.Regexp(t, regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\n`), output)
}

func TestFindIPAddresses_unknown(t *testing.T) {
	host := "idonotexist.iamunknown"
	output := captureLog(func() {
		FindIPAddresses(host)
	})
	assert.Contains(t, output, "no such host")
}

func TestFindMXRecords_known(t *testing.T) {
	host := "example.com"
	output := captureLog(func() {
		FindMXRecords(host)
	})
	assert.Equal(t, ". 0\n", output)
}

func TestFindMXRecords_unknown(t *testing.T) {
	host := "idonotexist.iamunknown"
	output := captureLog(func() {
		FindMXRecords(host)
	})
	assert.Contains(t, output, "no such host")
}

func TestFindNameServers_known(t *testing.T) {
	host := "apily.dev"
	output := captureLog(func() {
		FindNameServers(host)
	})
	assert.Contains(t, output, "transip")
}

func TestFindNameServers_unknown(t *testing.T) {
	host := "idonotexist.iamunknown"
	output := captureLog(func() {
		FindNameServers(host)
	})
	assert.Contains(t, output, "no such host")
}
