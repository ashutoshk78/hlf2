// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// FabricMainChannelExternalOrdererNodeApplyConfiguration represents an declarative configuration of the FabricMainChannelExternalOrdererNode type for use
// with apply.
type FabricMainChannelExternalOrdererNodeApplyConfiguration struct {
	Host      *string `json:"host,omitempty"`
	AdminPort *int    `json:"port,omitempty"`
}

// FabricMainChannelExternalOrdererNodeApplyConfiguration constructs an declarative configuration of the FabricMainChannelExternalOrdererNode type for use with
// apply.
func FabricMainChannelExternalOrdererNode() *FabricMainChannelExternalOrdererNodeApplyConfiguration {
	return &FabricMainChannelExternalOrdererNodeApplyConfiguration{}
}

// WithHost sets the Host field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Host field is set to the value of the last call.
func (b *FabricMainChannelExternalOrdererNodeApplyConfiguration) WithHost(value string) *FabricMainChannelExternalOrdererNodeApplyConfiguration {
	b.Host = &value
	return b
}

// WithAdminPort sets the AdminPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AdminPort field is set to the value of the last call.
func (b *FabricMainChannelExternalOrdererNodeApplyConfiguration) WithAdminPort(value int) *FabricMainChannelExternalOrdererNodeApplyConfiguration {
	b.AdminPort = &value
	return b
}
