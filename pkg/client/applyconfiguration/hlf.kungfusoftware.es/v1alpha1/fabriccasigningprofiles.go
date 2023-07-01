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

// FabricCASigningProfilesApplyConfiguration represents an declarative configuration of the FabricCASigningProfiles type for use
// with apply.
type FabricCASigningProfilesApplyConfiguration struct {
	CA  *FabricCASigningSignProfileApplyConfiguration `json:"ca,omitempty"`
	TLS *FabricCASigningTLSProfileApplyConfiguration  `json:"tls,omitempty"`
}

// FabricCASigningProfilesApplyConfiguration constructs an declarative configuration of the FabricCASigningProfiles type for use with
// apply.
func FabricCASigningProfiles() *FabricCASigningProfilesApplyConfiguration {
	return &FabricCASigningProfilesApplyConfiguration{}
}

// WithCA sets the CA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CA field is set to the value of the last call.
func (b *FabricCASigningProfilesApplyConfiguration) WithCA(value *FabricCASigningSignProfileApplyConfiguration) *FabricCASigningProfilesApplyConfiguration {
	b.CA = value
	return b
}

// WithTLS sets the TLS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLS field is set to the value of the last call.
func (b *FabricCASigningProfilesApplyConfiguration) WithTLS(value *FabricCASigningTLSProfileApplyConfiguration) *FabricCASigningProfilesApplyConfiguration {
	b.TLS = value
	return b
}
