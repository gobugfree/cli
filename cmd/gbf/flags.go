package main

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"strings"
)

func newFlags(name string, out io.Writer) *flag.FlagSet {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.SetOutput(out)
	return flags
}

// parseInterleaved parses flags that may appear before or after positional
// arguments (the flag package stops at the first positional) and returns
// the positionals in order.
func parseInterleaved(flags *flag.FlagSet, args []string) ([]string, error) {
	var positional []string
	for {
		if err := flags.Parse(args); err != nil {
			return nil, err
		}
		if flags.NArg() == 0 {
			return positional, nil
		}
		positional = append(positional, flags.Arg(0))
		args = flags.Args()[1:]
	}
}

// keyValues collects repeatable KEY=VALUE flags (--form, --param, --query).
type keyValues url.Values

func (v *keyValues) String() string { return url.Values(*v).Encode() }

func (v *keyValues) Set(value string) error {
	key, val, ok := strings.Cut(value, "=")
	if !ok || key == "" {
		return fmt.Errorf("%q is not KEY=VALUE", value)
	}
	if *v == nil {
		*v = keyValues{}
	}
	url.Values(*v).Add(key, val)
	return nil
}
