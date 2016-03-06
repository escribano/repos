package commands

import (
	"github.com/escribano/clif"
)

var (
	Commands = make([]func() *clif.Command, 0)
)

func stringsToMap(s []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}
	return m
}