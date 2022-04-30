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
	v1 "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
)

// VenafiTPPApplyConfiguration represents an declarative configuration of the VenafiTPP type for use
// with apply.
type VenafiTPPApplyConfiguration struct {
	URL            *string                  `json:"url,omitempty"`
	CredentialsRef *v1.LocalObjectReference `json:"credentialsRef,omitempty"`
	CABundle       []byte                   `json:"caBundle,omitempty"`
}

// VenafiTPPApplyConfiguration constructs an declarative configuration of the VenafiTPP type for use with
// apply.
func VenafiTPP() *VenafiTPPApplyConfiguration {
	return &VenafiTPPApplyConfiguration{}
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *VenafiTPPApplyConfiguration) WithURL(value string) *VenafiTPPApplyConfiguration {
	b.URL = &value
	return b
}

// WithCredentialsRef sets the CredentialsRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CredentialsRef field is set to the value of the last call.
func (b *VenafiTPPApplyConfiguration) WithCredentialsRef(value v1.LocalObjectReference) *VenafiTPPApplyConfiguration {
	b.CredentialsRef = &value
	return b
}

// WithCABundle adds the given value to the CABundle field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the CABundle field.
func (b *VenafiTPPApplyConfiguration) WithCABundle(values ...byte) *VenafiTPPApplyConfiguration {
	for i := range values {
		b.CABundle = append(b.CABundle, values[i])
	}
	return b
}
