package constants

import "crypto/tls"

// CipherSuites is a constant maps that displays the string based on the uint16 id of the tls.crypto package
var CipherSuites = map[uint16]string{
	tls.TLS_RSA_WITH_RC4_128_SHA:                      `TLS_RSA_WITH_RC4_128_SHA`,
	tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA:                 `TLS_RSA_WITH_3DES_EDE_CBC_SHA`,
	tls.TLS_RSA_WITH_AES_128_CBC_SHA:                  `TLS_RSA_WITH_AES_128_CBC_SHA`,
	tls.TLS_RSA_WITH_AES_256_CBC_SHA:                  `TLS_RSA_WITH_AES_256_CBC_SHA`,
	tls.TLS_RSA_WITH_AES_128_CBC_SHA256:               `TLS_RSA_WITH_AES_128_CBC_SHA256`,
	tls.TLS_RSA_WITH_AES_128_GCM_SHA256:               `TLS_RSA_WITH_AES_128_GCM_SHA256`,
	tls.TLS_RSA_WITH_AES_256_GCM_SHA384:               `TLS_RSA_WITH_AES_256_GCM_SHA384`,
	tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA:              `TLS_ECDHE_ECDSA_WITH_RC4_128_SHA`,
	tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA:          `TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA`,
	tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA:          `TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA`,
	tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA:                `TLS_ECDHE_RSA_WITH_RC4_128_SHA`,
	tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA:           `TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA`,
	tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA:            `TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA`,
	tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA:            `TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA`,
	tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256:       `TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256`,
	tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256:         `TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256`,
	tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:         `TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256`,
	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384:         `TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384`,
	tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384:       `TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384`,
	tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256:   `TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256`,
	tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256: `TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256`,

	// TLS 1.3 cipher suites.
	tls.TLS_AES_128_GCM_SHA256:       `TLS_AES_128_GCM_SHA256`,
	tls.TLS_AES_256_GCM_SHA384:       `TLS_AES_256_GCM_SHA384`,
	tls.TLS_CHACHA20_POLY1305_SHA256: `TLS_CHACHA20_POLY1305_SHA256`,

	// TLS_FALLBACK_SCSV isn't a standard cipher suite but an indicator
	// that the client is doing version fallback. See RFC 7507.
	tls.TLS_FALLBACK_SCSV: `TLS_FALLBACK_SCSV`,
}

// TLSVersions is a constant map that displays the TLS version based on the uint16 id of the tls.crypto package
var TLSVersions = map[uint16]string{
	tls.VersionTLS10: `VersionTLS10`,
	tls.VersionTLS11: `VersionTLS11`,
	tls.VersionTLS12: `VersionTLS12`,
	tls.VersionTLS13: `VersionTLS13`,

	// Deprecated: SSLv3 is cryptographically broken, and is no longer
	// supported by this package. See golang.org/issue/32716.
	tls.VersionSSL30: `VersionSSL30`,
}
