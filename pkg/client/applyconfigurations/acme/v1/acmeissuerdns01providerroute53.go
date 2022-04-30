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

// ACMEIssuerDNS01ProviderRoute53ApplyConfiguration represents an declarative configuration of the ACMEIssuerDNS01ProviderRoute53 type for use
// with apply.
type ACMEIssuerDNS01ProviderRoute53ApplyConfiguration struct {
	AccessKeyID     *string               `json:"accessKeyID,omitempty"`
	SecretAccessKey *v1.SecretKeySelector `json:"secretAccessKeySecretRef,omitempty"`
	Role            *string               `json:"role,omitempty"`
	HostedZoneID    *string               `json:"hostedZoneID,omitempty"`
	Region          *string               `json:"region,omitempty"`
}

// ACMEIssuerDNS01ProviderRoute53ApplyConfiguration constructs an declarative configuration of the ACMEIssuerDNS01ProviderRoute53 type for use with
// apply.
func ACMEIssuerDNS01ProviderRoute53() *ACMEIssuerDNS01ProviderRoute53ApplyConfiguration {
	return &ACMEIssuerDNS01ProviderRoute53ApplyConfiguration{}
}

// WithAccessKeyID sets the AccessKeyID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AccessKeyID field is set to the value of the last call.
func (b *ACMEIssuerDNS01ProviderRoute53ApplyConfiguration) WithAccessKeyID(value string) *ACMEIssuerDNS01ProviderRoute53ApplyConfiguration {
	b.AccessKeyID = &value
	return b
}

// WithSecretAccessKey sets the SecretAccessKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretAccessKey field is set to the value of the last call.
func (b *ACMEIssuerDNS01ProviderRoute53ApplyConfiguration) WithSecretAccessKey(value v1.SecretKeySelector) *ACMEIssuerDNS01ProviderRoute53ApplyConfiguration {
	b.SecretAccessKey = &value
	return b
}

// WithRole sets the Role field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Role field is set to the value of the last call.
func (b *ACMEIssuerDNS01ProviderRoute53ApplyConfiguration) WithRole(value string) *ACMEIssuerDNS01ProviderRoute53ApplyConfiguration {
	b.Role = &value
	return b
}

// WithHostedZoneID sets the HostedZoneID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostedZoneID field is set to the value of the last call.
func (b *ACMEIssuerDNS01ProviderRoute53ApplyConfiguration) WithHostedZoneID(value string) *ACMEIssuerDNS01ProviderRoute53ApplyConfiguration {
	b.HostedZoneID = &value
	return b
}

// WithRegion sets the Region field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Region field is set to the value of the last call.
func (b *ACMEIssuerDNS01ProviderRoute53ApplyConfiguration) WithRegion(value string) *ACMEIssuerDNS01ProviderRoute53ApplyConfiguration {
	b.Region = &value
	return b
}
