package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/cappyzawa/go-versions"
)

const (
	statusOK = iota
	statusParseFlagErr
	statusVersionsErr
	statusNoVersionsErr
)

type cli struct {
	err    io.Writer
	out    io.Writer
	client versions.Client
}

func (c *cli) Run(args []string) int {
	var (
		goos   string
		goarch string
	)
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.StringVar(&goos, "os", "", "select os")
	flags.StringVar(&goarch, "arch", "", "select arch")
	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintf(c.err, "failed to parse flags: %v\n", err)
		return statusParseFlagErr
	}

	versions, err := c.client.List()
	if err != nil {
		fmt.Fprintf(c.err, "failed to get go versions: %v\n", err)
		return statusVersionsErr
	}
	if len(versions) == 0 {
		fmt.Fprintf(c.err, "there is no go versions\n")
		return statusNoVersionsErr
	}

	for _, v := range versions {
		if strings.Contains(v, goos) && strings.Contains(v, goarch) {
			fmt.Fprintf(c.out, "%s\n", v)
		}
	}
	return statusOK
}

func main() {
	c := &cli{
		out:    os.Stdout,
		err:    os.Stderr,
		client: versions.NewClient(&versions.Config{}),
	}
	os.Exit(c.Run(os.Args))
}
