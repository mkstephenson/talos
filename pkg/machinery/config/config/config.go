// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package config provides interfaces to consume machine configuration values.
package config

// Config defines the interface to access contents of the machine configuration.
type Config interface {
	Debug() bool
	Machine() MachineConfig
	Cluster() ClusterConfig
	SideroLink() SideroLinkConfig
	ExtensionServicesConfig() ExtensionServicesConfigConfig
	Runtime() RuntimeConfig
	NetworkRules() NetworkRuleConfig
}
