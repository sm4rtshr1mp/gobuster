package gobusters3

import (
	"github.com/sm4rtshr1mp/gobuster/v3/libgobuster"
)

// OptionsS3 is the struct to hold all options for this plugin
type OptionsS3 struct {
	libgobuster.BasicHTTPOptions
	MaxFilesToList  int
	NoTLSValidation bool
}

// NewOptionsS3 returns a new initialized OptionsS3
func NewOptionsS3() *OptionsS3 {
	return &OptionsS3{}
}
