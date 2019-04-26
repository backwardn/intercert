package server

import (
	"strings"

	"github.com/bobesa/go-domain-util/domainutil"
	"github.com/go-acme/lego/log"
	"github.com/pkg/errors"
)

// Whitelist holds valid domains that certificates can be
// issued under.
type Whitelist struct {
	// The top level domains to allow
	domains []string
}

// NewWhitelist creates a new instance of the Whitelist struct
func NewWhitelist(domains []string) Whitelist {
	lowercasedDomains := make([]string, len(domains))

	for i, v := range domains {
		lowercasedDomains[i] = strings.ToLower(v)
	}

	if len(lowercasedDomains) == 0 {
		log.Warnf("No domains in whitelist - every domain is allowed")
	} else {
		log.Infof("Loaded whitelist: %v", lowercasedDomains)
	}

	return Whitelist{domains: lowercasedDomains}
}

// Checks whether a DNS name (e.g. foo.bar.com) is allowed. If no domains is configured,
// every DNS name will be allowed.
func (w Whitelist) isDNSNameAllowed(dnsName string) error {
	// Empty whitelists allows everything
	if len(w.domains) == 0 {
		return nil
	}

	topLevel := domainutil.Domain(dnsName)

	if topLevel == "" {
		return errors.New("Could not check whether " + dnsName + " is allowed")

	}

	for _, domain := range w.domains {
		if domain == topLevel {
			return nil
		}
	}

	return errors.New(topLevel + " did not match any entries in the whitelist")
}
