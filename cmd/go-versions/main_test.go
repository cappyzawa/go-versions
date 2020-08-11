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
			args:      []string{"command"},
			versions:  []string{"go1", "go2"},
			expect:    statusOK,
			expectOut: "go1\ngo2\n",
		},
		"os is specified as linux": {
			args: []string{"command", "-os", "linux"},
			versions: []string{
				"https://golang.org/dl/go1.14.7.linux-amd64.tar.gz",
				"https://golang.org/dl/go1.14.7.windows-386.zip",
				"https://golang.org/dl/go1.9.linux-386.tar.gz",
			},
			expect:    statusOK,
			expectOut: "https://golang.org/dl/go1.14.7.linux-amd64.tar.gz\nhttps://golang.org/dl/go1.9.linux-386.tar.gz\n",
		},
		"arch is specified as amd64": {
			args: []string{"command", "-arch", "amd64"},
			versions: []string{
				"https://golang.org/dl/go1.14.7.linux-amd64.tar.gz",
				"https://golang.org/dl/go1.14.7.windows-386.zip",
				"https://golang.org/dl/go1.2.2.windows-amd64.zip",
			},
			expect:    statusOK,
			expectOut: "https://golang.org/dl/go1.14.7.linux-amd64.tar.gz\nhttps://golang.org/dl/go1.2.2.windows-amd64.zip\n",
		},
		"os and arch are specified as linux and amd64": {
			args: []string{"command", "-os", "linux", "-arch", "amd64"},
			versions: []string{
				"https://golang.org/dl/go1.14.7.linux-amd64.tar.gz",
				"https://golang.org/dl/go1.14.7.windows-386.zip",
				"https://golang.org/dl/go1.2.2.windows-amd64.zip",
			},
			expect:    statusOK,
			expectOut: "https://golang.org/dl/go1.14.7.linux-amd64.tar.gz\n",
		},
		"invalid args": {
			args:      []string{"command", "-invalid", "args"},
			expect:    statusParseFlagErr,
			expectErr: "failed to parse flags: flag provided but not defined: -invalid\n",
		},
		"failed to get go versions": {
			args:        []string{"command"},
			versionsErr: fmt.Errorf("some error"),
			expect:      statusVersionsErr,
			expectErr:   "failed to get go versions: some error\n",
		},
		"there is no go versions": {
			args:      []string{"command"},
			versions:  []string{},
			expect:    statusNoVersionsErr,
			expectErr: "there is no go versions\n",
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
				t.Errorf("error should be %s, but it is %s", test.expectErr, errBuf.String())
			}
		})
	}
}
