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

// checks if the PostTokenProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostTokenProperties{}

// PostTokenProperties struct for PostTokenProperties
type PostTokenProperties struct {
	ExpiryDate *IonosTime `json:"expiryDate,omitempty"`
	Name       string     `json:"name"`
	Scopes     []Scope    `json:"scopes,omitempty"`
	Status     *string    `json:"status,omitempty"`
}

// NewPostTokenProperties instantiates a new PostTokenProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTokenProperties(name string) *PostTokenProperties {
	this := PostTokenProperties{}

	this.Name = name

	return &this
}

// NewPostTokenPropertiesWithDefaults instantiates a new PostTokenProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTokenPropertiesWithDefaults() *PostTokenProperties {
	this := PostTokenProperties{}
	return &this
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *PostTokenProperties) GetExpiryDate() time.Time {
	if o == nil || IsNil(o.ExpiryDate) {
		var ret time.Time
		return ret
	}
	return o.ExpiryDate.Time
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTokenProperties) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiryDate) {
		return nil, false
	}
	return &o.ExpiryDate.Time, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *PostTokenProperties) HasExpiryDate() bool {
	if o != nil && !IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *PostTokenProperties) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &IonosTime{v}
}

// GetName returns the Name field value
func (o *PostTokenProperties) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostTokenProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostTokenProperties) SetName(v string) {
	o.Name = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *PostTokenProperties) GetScopes() []Scope {
	if o == nil || IsNil(o.Scopes) {
		var ret []Scope
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTokenProperties) GetScopesOk() ([]Scope, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *PostTokenProperties) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []Scope and assigns it to the Scopes field.
func (o *PostTokenProperties) SetScopes(v []Scope) {
	o.Scopes = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PostTokenProperties) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTokenProperties) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PostTokenProperties) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PostTokenProperties) SetStatus(v string) {
	o.Status = &v
}

func (o PostTokenProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiryDate) {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullablePostTokenProperties struct {
	value *PostTokenProperties
	isSet bool
}

func (v NullablePostTokenProperties) Get() *PostTokenProperties {
	return v.value
}

func (v *NullablePostTokenProperties) Set(val *PostTokenProperties) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTokenProperties) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTokenProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTokenProperties(val *PostTokenProperties) *NullablePostTokenProperties {
	return &NullablePostTokenProperties{value: val, isSet: true}
}

func (v NullablePostTokenProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTokenProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
