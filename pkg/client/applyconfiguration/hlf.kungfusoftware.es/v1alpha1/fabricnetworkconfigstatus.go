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

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.pkuidlabs.com/v1alpha1"
	status "github.com/kfsoftware/hlf-operator/pkg/status"
)

// FabricNetworkConfigStatusApplyConfiguration represents an declarative configuration of the FabricNetworkConfigStatus type for use
// with apply.
type FabricNetworkConfigStatusApplyConfiguration struct {
	Conditions *status.Conditions         `json:"conditions,omitempty"`
	Message    *string                    `json:"message,omitempty"`
	Status     *v1alpha1.DeploymentStatus `json:"status,omitempty"`
}

// FabricNetworkConfigStatusApplyConfiguration constructs an declarative configuration of the FabricNetworkConfigStatus type for use with
// apply.
func FabricNetworkConfigStatus() *FabricNetworkConfigStatusApplyConfiguration {
	return &FabricNetworkConfigStatusApplyConfiguration{}
}

// WithConditions sets the Conditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conditions field is set to the value of the last call.
func (b *FabricNetworkConfigStatusApplyConfiguration) WithConditions(value status.Conditions) *FabricNetworkConfigStatusApplyConfiguration {
	b.Conditions = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *FabricNetworkConfigStatusApplyConfiguration) WithMessage(value string) *FabricNetworkConfigStatusApplyConfiguration {
	b.Message = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *FabricNetworkConfigStatusApplyConfiguration) WithStatus(value v1alpha1.DeploymentStatus) *FabricNetworkConfigStatusApplyConfiguration {
	b.Status = &value
	return b
}
