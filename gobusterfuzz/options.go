package gobusterfuzz

import (
	"github.com/sm4rtshr1mp/gobuster/v3/libgobuster"
)

// OptionsFuzz is the struct to hold all options for this plugin
type OptionsFuzz struct {
	libgobuster.HTTPOptions
	ExcludedStatusCodes       string
	ExcludedStatusCodesParsed libgobuster.IntSet
	WildcardForced            bool
	ExcludeLength             []int
}

// NewOptionsFuzz returns a new initialized OptionsFuzz
func NewOptionsFuzz() *OptionsFuzz {
	return &OptionsFuzz{
		ExcludedStatusCodesParsed: libgobuster.NewIntSet(),
	}
}
