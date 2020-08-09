package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/cappyzawa/go-versions"
)

func TestCliRun(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		args        []string
		versions    []string
		versionsErr error
		expect      int
		expectOut   string
		expectErr   string
	}{
		"basic": {
			args:      nil,
			versions:  []string{"go1", "go2"},
			expect:    statusOK,
			expectOut: "go1\ngo2\n",
		},
		"failed to get go versions": {
			args:        nil,
			versionsErr: fmt.Errorf("some error"),
			expect:      statusVersionsErr,
			expectErr:   "failed to get go versions: some error\n",
		},
	}

	for name, test := range cases {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			outBuf := new(bytes.Buffer)
			errBuf := new(bytes.Buffer)
			c := &cli{
				out: outBuf,
				err: errBuf,
				client: &versions.MockClient{
					MockList: func() ([]string, error) {
						return test.versions, test.versionsErr
					},
				},
			}
			actual := c.Run(test.args)
			if actual != test.expect {
				t.Errorf("code should be %d, but it is %d", test.expect, actual)
			}
			if outBuf.String() != test.expectOut {
				t.Errorf("output should be %s, but it is %s", test.expectOut, outBuf.String())
			}
			if errBuf.String() != test.expectErr {
				t.Errorf("errput should be %s, but it is %s", test.expectErr, errBuf.String())
			}
		})
	}
}
