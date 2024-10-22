/*
 * Certificate Manager Service API
 *
 * Using the Certificate Manager Service, you can conveniently provision and manage SSL certificates  with IONOS services and your internal connected resources.   For the [Application Load Balancer](https://api.ionos.com/docs/cloud/v6/#Application-Load-Balancers-get-datacenters-datacenterId-applicationloadbalancers), you usually need a certificate to encrypt your HTTPS traffic. The service provides the basic functions of uploading and deleting your certificates for this purpose.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cert

import (
	"encoding/json"
)

// checks if the Provider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Provider{}

// Provider Certificate provider used to renew certificates before their expiry.
type Provider struct {
	// The name of the certificate provider.
	Name string `json:"name"`
	// The email address of the certificate requester.
	Email string `json:"email"`
	// The URL of the certificate provider.
	Server                 string                          `json:"server"`
	ExternalAccountBinding *ProviderExternalAccountBinding `json:"externalAccountBinding,omitempty"`
}

// NewProvider instantiates a new Provider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvider(name string, email string, server string) *Provider {
	this := Provider{}

	this.Name = name
	this.Email = email
	this.Server = server

	return &this
}

// NewProviderWithDefaults instantiates a new Provider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderWithDefaults() *Provider {
	this := Provider{}
	return &this
}

// GetName returns the Name field value
func (o *Provider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Provider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Provider) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *Provider) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Provider) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Provider) SetEmail(v string) {
	o.Email = v
}

// GetServer returns the Server field value
func (o *Provider) GetServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *Provider) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *Provider) SetServer(v string) {
	o.Server = v
}

// GetExternalAccountBinding returns the ExternalAccountBinding field value if set, zero value otherwise.
func (o *Provider) GetExternalAccountBinding() ProviderExternalAccountBinding {
	if o == nil || IsNil(o.ExternalAccountBinding) {
		var ret ProviderExternalAccountBinding
		return ret
	}
	return *o.ExternalAccountBinding
}

// GetExternalAccountBindingOk returns a tuple with the ExternalAccountBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetExternalAccountBindingOk() (*ProviderExternalAccountBinding, bool) {
	if o == nil || IsNil(o.ExternalAccountBinding) {
		return nil, false
	}
	return o.ExternalAccountBinding, true
}

// HasExternalAccountBinding returns a boolean if a field has been set.
func (o *Provider) HasExternalAccountBinding() bool {
	if o != nil && !IsNil(o.ExternalAccountBinding) {
		return true
	}

	return false
}

// SetExternalAccountBinding gets a reference to the given ProviderExternalAccountBinding and assigns it to the ExternalAccountBinding field.
func (o *Provider) SetExternalAccountBinding(v ProviderExternalAccountBinding) {
	o.ExternalAccountBinding = &v
}

func (o Provider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	toSerialize["server"] = o.Server
	if !IsNil(o.ExternalAccountBinding) {
		toSerialize["externalAccountBinding"] = o.ExternalAccountBinding
	}
	return toSerialize, nil
}

type NullableProvider struct {
	value *Provider
	isSet bool
}

func (v NullableProvider) Get() *Provider {
	return v.value
}

func (v *NullableProvider) Set(val *Provider) {
	v.value = val
	v.isSet = true
}

func (v NullableProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvider(val *Provider) *NullableProvider {
	return &NullableProvider{value: val, isSet: true}
}

func (v NullableProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}