package gobustervhost

import (
	"github.com/sm4rtshr1mp/gobuster/v3/libgobuster"
)

// OptionsVhost is the struct to hold all options for this plugin
type OptionsVhost struct {
	libgobuster.HTTPOptions
	AppendDomain bool
}
