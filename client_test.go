package versions_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/cappyzawa/go-versions"
)

func TestATCLIRun(t *testing.T) {
	t.Parallel()
	isATStr := os.Getenv("AT")
	if isATStr != "true" {
		isATStr = "false"
	}
	isAT, _ := strconv.ParseBool(isATStr)
	if !isAT {
		t.Skip("skip, non acceptance test")
		return
	}
	cases := map[string]struct {
		expect   []string
		existErr bool
	}{
		"basic": {
			expect: []string{},
		},
	}

	for name, test := range cases {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			client := versions.NewClient(&versions.Config{})
			_, err := client.List()
			if err != nil && !test.existErr {
				t.Errorf("error should not been occurred: %v", err)
			}
			if err == nil && test.existErr {
				t.Errorf("error should been occurred")
			}
		})
	}
}
