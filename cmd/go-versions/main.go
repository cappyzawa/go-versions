package main

import (
	"fmt"
	"io"
	"os"

	"github.com/cappyzawa/go-versions"
)

const (
	statusOK = iota
	statusVersionsErr
	statusNoVersionsErr
)

type cli struct {
	err    io.Writer
	out    io.Writer
	client versions.Client
}

func (c *cli) Run(args []string) int {
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
		fmt.Fprintf(c.out, "%s\n", v)
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
