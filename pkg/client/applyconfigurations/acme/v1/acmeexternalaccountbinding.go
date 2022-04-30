/*
Copyright The cert-manager Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	acmev1 "github.com/cert-manager/cert-manager/pkg/apis/acme/v1"
	v1 "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
)

// ACMEExternalAccountBindingApplyConfiguration represents an declarative configuration of the ACMEExternalAccountBinding type for use
// with apply.
type ACMEExternalAccountBindingApplyConfiguration struct {
	KeyID        *string                  `json:"keyID,omitempty"`
	Key          *v1.SecretKeySelector    `json:"keySecretRef,omitempty"`
	KeyAlgorithm *acmev1.HMACKeyAlgorithm `json:"keyAlgorithm,omitempty"`
}

// ACMEExternalAccountBindingApplyConfiguration constructs an declarative configuration of the ACMEExternalAccountBinding type for use with
// apply.
func ACMEExternalAccountBinding() *ACMEExternalAccountBindingApplyConfiguration {
	return &ACMEExternalAccountBindingApplyConfiguration{}
}

// WithKeyID sets the KeyID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeyID field is set to the value of the last call.
func (b *ACMEExternalAccountBindingApplyConfiguration) WithKeyID(value string) *ACMEExternalAccountBindingApplyConfiguration {
	b.KeyID = &value
	return b
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *ACMEExternalAccountBindingApplyConfiguration) WithKey(value v1.SecretKeySelector) *ACMEExternalAccountBindingApplyConfiguration {
	b.Key = &value
	return b
}

// WithKeyAlgorithm sets the KeyAlgorithm field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeyAlgorithm field is set to the value of the last call.
func (b *ACMEExternalAccountBindingApplyConfiguration) WithKeyAlgorithm(value acmev1.HMACKeyAlgorithm) *ACMEExternalAccountBindingApplyConfiguration {
	b.KeyAlgorithm = &value
	return b
}
