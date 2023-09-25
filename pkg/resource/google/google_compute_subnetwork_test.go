package google_test

import (
	"testing"
	"time"

	"github.com/snyk/driftctl/test"
	"github.com/snyk/driftctl/test/acceptance"
)

func TestAcc_Google_ComputeSubnetwork(t *testing.T) {
	t.Skip("flake")

	acceptance.Run(t, acceptance.AccTestCase{
		TerraformVersion: "0.15.5",
		Paths:            []string{"./testdata/acc/google_compute_subnetwork"},
		Args: []string{
			"scan",
			"--to", "gcp+tf",
			"--filter", "contains(Id, 'driftctl-acc-subnet-')",
		},
		Checks: []acceptance.AccCheck{
			{
				// New resources are not visible immediately through GCP API after an apply operation.
				ShouldRetry: acceptance.LinearBackoff(15 * time.Minute),
				Check: func(result *test.ScanResult, stdout string, err error) {
					if err != nil {
						t.Fatal(err)
					}
					result.AssertInfrastructureIsInSync()
					result.AssertManagedCount(1)
				},
			},
		},
	})
}
