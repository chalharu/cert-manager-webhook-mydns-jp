package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/cert-manager/cert-manager/test/acme/dns"
)

var (
	zone = os.Getenv("TEST_ZONE_NAME")
)

func TestRunsSuite(t *testing.T) {
	// The manifest path should contain a file named config.json that is a
	// snippet of valid configuration that should be included on the
	// ChallengeRequest passed as part of the test cases.
	//

	// Uncomment the below fixture when implementing your custom DNS provider
	//fixture := dns.NewFixture(&customDNSProviderSolver{},
	//	dns.SetResolvedZone(zone),
	//	dns.SetAllowAmbientCredentials(false),
	//	dns.SetManifestPath("testdata/my-custom-solver"),
	//	dns.SetBinariesPath("_test/kubebuilder/bin"),
	//)
	fmt.Printf("Zone: %v\n", zone)
	solver := &customDNSProviderSolver{}
	fixture := dns.NewFixture(solver,
		// dns.SetDNSName(zone),
		dns.SetResolvedZone(zone),
		dns.SetResolvedFQDN(fmt.Sprintf("_acme-challenge.%s", zone)),
		dns.SetManifestPath("testdata/mydns-solver"),
		dns.SetDNSServer("210.197.74.200:53"),
		dns.SetUseAuthoritative(false),
	)
	//need to uncomment and  RunConformance delete runBasic and runExtended once https://github.com/cert-manager/cert-manager/pull/4835 is merged
	//fixture.RunConformance(t)
	fixture.RunBasic(t)
	fixture.RunExtended(t)
}
