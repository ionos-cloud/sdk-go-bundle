/*
 * Container Registry service
 *
 * ## Overview Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls. ## Changelog ### 1.1.0  - Added new endpoints for Repositories  - Added new endpoints for Artifacts  - Added new endpoints for Vulnerabilities  - Added registry vulnerabilityScanning feature ### 1.2.0  - Added registry `apiSubnetAllowList` ### 1.2.1  - Amended `apiSubnetAllowList` Regex
 *
 * API version: 1.2.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerregistry

import (
	"encoding/json"

	"time"
)

// checks if the PatchTokenInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchTokenInput{}

// PatchTokenInput struct for PatchTokenInput
type PatchTokenInput struct {
	ExpiryDate *NullableIonosTime `json:"expiryDate,omitempty"`
	Scopes     []Scope            `json:"scopes,omitempty"`
	Status     *string            `json:"status,omitempty"`
}

// NewPatchTokenInput instantiates a new PatchTokenInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchTokenInput() *PatchTokenInput {
	this := PatchTokenInput{}

	return &this
}

// NewPatchTokenInputWithDefaults instantiates a new PatchTokenInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchTokenInputWithDefaults() *PatchTokenInput {
	this := PatchTokenInput{}
	return &this
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchTokenInput) GetExpiryDate() time.Time {
	if o == nil || IsNil(o.ExpiryDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate.Get()
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchTokenInput) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiryDate.Get(), o.ExpiryDate.IsSet()
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *PatchTokenInput) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate.IsSet() {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given NullableTime and assigns it to the ExpiryDate field.
func (o *PatchTokenInput) SetExpiryDate(v time.Time) {
	o.ExpiryDate.Set(&v)
}

// SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil
func (o *PatchTokenInput) SetExpiryDateNil() {
	o.ExpiryDate.Set(nil)
}

// UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
func (o *PatchTokenInput) UnsetExpiryDate() {
	o.ExpiryDate.Unset()
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *PatchTokenInput) GetScopes() []Scope {
	if o == nil || IsNil(o.Scopes) {
		var ret []Scope
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTokenInput) GetScopesOk() ([]Scope, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *PatchTokenInput) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []Scope and assigns it to the Scopes field.
func (o *PatchTokenInput) SetScopes(v []Scope) {
	o.Scopes = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchTokenInput) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTokenInput) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchTokenInput) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PatchTokenInput) SetStatus(v string) {
	o.Status = &v
}

func (o PatchTokenInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchTokenInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiryDate != nil && o.ExpiryDate.IsSet() {
		toSerialize["expiryDate"] = o.ExpiryDate.Get()
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullablePatchTokenInput struct {
	value *PatchTokenInput
	isSet bool
}

func (v NullablePatchTokenInput) Get() *PatchTokenInput {
	return v.value
}

func (v *NullablePatchTokenInput) Set(val *PatchTokenInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchTokenInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchTokenInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchTokenInput(val *PatchTokenInput) *NullablePatchTokenInput {
	return &NullablePatchTokenInput{value: val, isSet: true}
}

func (v NullablePatchTokenInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchTokenInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
