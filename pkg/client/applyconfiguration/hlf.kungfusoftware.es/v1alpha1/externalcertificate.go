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

// ExternalCertificateApplyConfiguration represents an declarative configuration of the ExternalCertificate type for use
// with apply.
type ExternalCertificateApplyConfiguration struct {
	SecretName         *string `json:"secretName,omitempty"`
	SecretNamespace    *string `json:"secretNamespace,omitempty"`
	RootCertificateKey *string `json:"rootCertificateKey,omitempty"`
	CertificateKey     *string `json:"certificateKey,omitempty"`
	PrivateKeyKey      *string `json:"privateKeyKey,omitempty"`
}

// ExternalCertificateApplyConfiguration constructs an declarative configuration of the ExternalCertificate type for use with
// apply.
func ExternalCertificate() *ExternalCertificateApplyConfiguration {
	return &ExternalCertificateApplyConfiguration{}
}

// WithSecretName sets the SecretName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretName field is set to the value of the last call.
func (b *ExternalCertificateApplyConfiguration) WithSecretName(value string) *ExternalCertificateApplyConfiguration {
	b.SecretName = &value
	return b
}

// WithSecretNamespace sets the SecretNamespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretNamespace field is set to the value of the last call.
func (b *ExternalCertificateApplyConfiguration) WithSecretNamespace(value string) *ExternalCertificateApplyConfiguration {
	b.SecretNamespace = &value
	return b
}

// WithRootCertificateKey sets the RootCertificateKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RootCertificateKey field is set to the value of the last call.
func (b *ExternalCertificateApplyConfiguration) WithRootCertificateKey(value string) *ExternalCertificateApplyConfiguration {
	b.RootCertificateKey = &value
	return b
}

// WithCertificateKey sets the CertificateKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CertificateKey field is set to the value of the last call.
func (b *ExternalCertificateApplyConfiguration) WithCertificateKey(value string) *ExternalCertificateApplyConfiguration {
	b.CertificateKey = &value
	return b
}

// WithPrivateKeyKey sets the PrivateKeyKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PrivateKeyKey field is set to the value of the last call.
func (b *ExternalCertificateApplyConfiguration) WithPrivateKeyKey(value string) *ExternalCertificateApplyConfiguration {
	b.PrivateKeyKey = &value
	return b
}
