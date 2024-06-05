package test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name        string
	FixturesDir string
	ModuleDir   string
}

func TestPeeringActive(t *testing.T) {
	testCases := []TestCase{
		{"SingleAccountSingleRegion", "./fixtures/single-account-single-region", "../examples/single-account-single-region"},
		{"SingleAccountSingleRegionOneDualStack", "./fixtures/single-account-single-region-one-dualstack", "../examples/single-account-single-region-one-dualstack"},
		{"SingleAccountSingleRegionWithOptions", "./fixtures/single-account-single-region-with-options", "../examples/single-account-single-region-with-options"},
		{"SingleAccountMultiRegion", "./fixtures/single-account-multi-region", "../examples/single-account-multi-region"},
		{"MultiAccountSingleRegion", "./fixtures/multi-account-single-region", "../examples/multi-account-single-region"},
		{"MultiAccountSingleRegionBothDualStack", "./fixtures/multi-account-single-region-both-dualstack", "../examples/multi-account-single-region-both-dualstack"},
		{"MultiAccountMultiRegion", "./fixtures/multi-account-multi-region", "../examples/multi-account-multi-region"},
		// {"ModuleDependsOn", "", "../examples/module-depends-on"},
		{"AssociatedCIDRs", "./fixtures/associated-cidr", "../examples/associated-cidrs"},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) { terratestRun(t, tc) })
	}
}

func terratestRun(t *testing.T, tc TestCase) {
	t.Parallel()

	testRand := rand.Intn(int(math.Pow10(11)))
	testId := fmt.Sprintf("%011d", testRand)
	tfVars := make(map[string]interface{})
	tfVars["test_id"] = testId

	// Assertions
	expectedPeeringStatus := "active"

	// Check if we need to apply fixtures first
	if tc.FixturesDir != "" {
		// Terraform Options for fixtures
		fixturesTerraformOptions := &terraform.Options{
			TerraformDir: tc.FixturesDir,
			Vars: map[string]interface{}{
				"test_id": testId,
			},
		}

		// Remove the fixtures resources in the end of the test
		defer terraform.Destroy(t, fixturesTerraformOptions)
		// Install Prerequisites
		terraform.InitAndApply(t, fixturesTerraformOptions)
		// Get the outputs from fixtures
		thisVpcID := terraform.Output(t, fixturesTerraformOptions, "this_vpc_id")
		peerVpcID := terraform.Output(t, fixturesTerraformOptions, "peer_vpc_id")

		tfVars["this_vpc_id"] = thisVpcID
		tfVars["peer_vpc_id"] = peerVpcID
	}

	// Terraform Options for module
	moduleTerraformOptions := &terraform.Options{
		TerraformDir: tc.ModuleDir,
		Vars:         tfVars,
	}

	// Remove the module resources in the end of the test
	defer terraform.Destroy(t, moduleTerraformOptions)
	// Create module resources
	terraform.InitAndApply(t, moduleTerraformOptions)
	// Retrieve information with `terraform output`
	actualPeeringStatus := terraform.Output(t, moduleTerraformOptions, "vpc_peering_accept_status")
	// Verify results
	assert.Equal(t, expectedPeeringStatus, actualPeeringStatus)
}
