// (c) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package precompilebind

// tmplSourcePrecompileConfigGo is the Go precompiled config source template.
const tmplSourcePrecompileConfigGo = `
// Code generated
// This file is a generated precompile contract config with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package {{.Package}}

import (
	"github.com/MetalBlockchain/subnet-evm/precompile/precompileconfig"
	{{- if .Contract.AllowList}}
	"github.com/MetalBlockchain/subnet-evm/precompile/allowlist"

	"github.com/ethereum/go-ethereum/common"
	{{- end}}

)

var _ precompileconfig.Config = &Config{}

// Config implements the precompileconfig.Config interface and
// adds specific configuration for {{.Contract.Type}}.
type Config struct {
	{{- if .Contract.AllowList}}
	allowlist.AllowListConfig
	{{- end}}
	precompileconfig.Upgrade
	// CUSTOM CODE STARTS HERE
	// Add your own custom fields for Config here
}

// NewConfig returns a config for a network upgrade at [blockTimestamp] that enables
// {{.Contract.Type}} {{- if .Contract.AllowList}} with the given [admins], [enableds] and [managers] members of the allowlist {{end}}.
func NewConfig(blockTimestamp *uint64{{if .Contract.AllowList}}, admins []common.Address, enableds []common.Address, managers []common.Address{{end}}) *Config {
	return &Config{
		{{- if .Contract.AllowList}}
		AllowListConfig: allowlist.AllowListConfig{
			AdminAddresses: admins,
			EnabledAddresses: enableds,
			ManagerAddresses: managers,
		},
		{{- end}}
		Upgrade: precompileconfig.Upgrade{BlockTimestamp: blockTimestamp},
	}
}

// NewDisableConfig returns config for a network upgrade at [blockTimestamp]
// that disables {{.Contract.Type}}.
func NewDisableConfig(blockTimestamp *uint64) *Config {
	return &Config{
		Upgrade: precompileconfig.Upgrade{
			BlockTimestamp: blockTimestamp,
			Disable:        true,
		},
	}
}

// Key returns the key for the {{.Contract.Type}} precompileconfig.
// This should be the same key as used in the precompile module.
func (*Config) Key() string { return ConfigKey }

// Verify tries to verify Config and returns an error accordingly.
func (c *Config) Verify(chainConfig precompileconfig.ChainConfig) error {
	{{- if .Contract.AllowList}}
	// Verify AllowList first
	if err := c.AllowListConfig.Verify(chainConfig, c.Upgrade); err != nil {
		return err
	}
	{{- end}}
	// CUSTOM CODE STARTS HERE
	// Add your own custom verify code for Config here
	// and return an error accordingly
	return nil
}

// Equal returns true if [s] is a [*Config] and it has been configured identical to [c].
func (c *Config) Equal(s precompileconfig.Config) bool {
	// typecast before comparison
	other, ok := (s).(*Config)
	if !ok {
		return false
	}
	// CUSTOM CODE STARTS HERE
	// modify this boolean accordingly with your custom Config, to check if [other] and the current [c] are equal
	// if Config contains only Upgrade {{- if .Contract.AllowList}} and AllowListConfig {{end}} you can skip modifying it.
	equals := c.Upgrade.Equal(&other.Upgrade) {{- if .Contract.AllowList}} && c.AllowListConfig.Equal(&other.AllowListConfig) {{end}}
	return equals
}
`
