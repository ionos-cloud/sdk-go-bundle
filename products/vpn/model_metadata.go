/*
 * VPN Gateways
 *
 * POC Docs for VPN gateway as service
 *
 * API version: 0.0.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vpn

import (
	"encoding/json"

	"time"
)

// checks if the Metadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Metadata{}

// Metadata Metadata of the resource.
type Metadata struct {
	// The ISO 8601 creation timestamp.
	CreatedDate *IonosTime `json:"createdDate,omitempty"`
	// Unique name of the identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Unique id of the identity that created the resource.
	CreatedByUserId *string `json:"createdByUserId,omitempty"`
	// The ISO 8601 modified timestamp.
	LastModifiedDate *IonosTime `json:"lastModifiedDate,omitempty"`
	// Unique name of the identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// Unique id of the identity that last modified the resource.
	LastModifiedByUserId *string `json:"lastModifiedByUserId,omitempty"`
	// Unique name of the resource.
	ResourceURN *string `json:"resourceURN,omitempty"`
}

// NewMetadata instantiates a new Metadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadata() *Metadata {
	this := Metadata{}

	return &this
}

// NewMetadataWithDefaults instantiates a new Metadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataWithDefaults() *Metadata {
	this := Metadata{}
	return &this
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *Metadata) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return o.CreatedDate.Time
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return &o.CreatedDate.Time, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *Metadata) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *Metadata) SetCreatedDate(v time.Time) {
	o.CreatedDate = &IonosTime{v}
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Metadata) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Metadata) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Metadata) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *Metadata) GetCreatedByUserId() string {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret string
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetCreatedByUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *Metadata) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given string and assigns it to the CreatedByUserId field.
func (o *Metadata) SetCreatedByUserId(v string) {
	o.CreatedByUserId = &v
}

// GetLastModifiedDate returns the LastModifiedDate field value if set, zero value otherwise.
func (o *Metadata) GetLastModifiedDate() time.Time {
	if o == nil || IsNil(o.LastModifiedDate) {
		var ret time.Time
		return ret
	}
	return o.LastModifiedDate.Time
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedDate) {
		return nil, false
	}
	return &o.LastModifiedDate.Time, true
}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *Metadata) HasLastModifiedDate() bool {
	if o != nil && !IsNil(o.LastModifiedDate) {
		return true
	}

	return false
}

// SetLastModifiedDate gets a reference to the given time.Time and assigns it to the LastModifiedDate field.
func (o *Metadata) SetLastModifiedDate(v time.Time) {
	o.LastModifiedDate = &IonosTime{v}
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *Metadata) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetLastModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedBy) {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *Metadata) HasLastModifiedBy() bool {
	if o != nil && !IsNil(o.LastModifiedBy) {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given string and assigns it to the LastModifiedBy field.
func (o *Metadata) SetLastModifiedBy(v string) {
	o.LastModifiedBy = &v
}

// GetLastModifiedByUserId returns the LastModifiedByUserId field value if set, zero value otherwise.
func (o *Metadata) GetLastModifiedByUserId() string {
	if o == nil || IsNil(o.LastModifiedByUserId) {
		var ret string
		return ret
	}
	return *o.LastModifiedByUserId
}

// GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetLastModifiedByUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedByUserId) {
		return nil, false
	}
	return o.LastModifiedByUserId, true
}

// HasLastModifiedByUserId returns a boolean if a field has been set.
func (o *Metadata) HasLastModifiedByUserId() bool {
	if o != nil && !IsNil(o.LastModifiedByUserId) {
		return true
	}

	return false
}

// SetLastModifiedByUserId gets a reference to the given string and assigns it to the LastModifiedByUserId field.
func (o *Metadata) SetLastModifiedByUserId(v string) {
	o.LastModifiedByUserId = &v
}

// GetResourceURN returns the ResourceURN field value if set, zero value otherwise.
func (o *Metadata) GetResourceURN() string {
	if o == nil || IsNil(o.ResourceURN) {
		var ret string
		return ret
	}
	return *o.ResourceURN
}

// GetResourceURNOk returns a tuple with the ResourceURN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetResourceURNOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceURN) {
		return nil, false
	}
	return o.ResourceURN, true
}

// HasResourceURN returns a boolean if a field has been set.
func (o *Metadata) HasResourceURN() bool {
	if o != nil && !IsNil(o.ResourceURN) {
		return true
	}

	return false
}

// SetResourceURN gets a reference to the given string and assigns it to the ResourceURN field.
func (o *Metadata) SetResourceURN(v string) {
	o.ResourceURN = &v
}

func (o Metadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Metadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.CreatedByUserId) {
		toSerialize["createdByUserId"] = o.CreatedByUserId
	}
	if !IsNil(o.LastModifiedDate) {
		toSerialize["lastModifiedDate"] = o.LastModifiedDate
	}
	if !IsNil(o.LastModifiedBy) {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if !IsNil(o.LastModifiedByUserId) {
		toSerialize["lastModifiedByUserId"] = o.LastModifiedByUserId
	}
	if !IsNil(o.ResourceURN) {
		toSerialize["resourceURN"] = o.ResourceURN
	}
	return toSerialize, nil
}

type NullableMetadata struct {
	value *Metadata
	isSet bool
}

func (v NullableMetadata) Get() *Metadata {
	return v.value
}

func (v *NullableMetadata) Set(val *Metadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadata(val *Metadata) *NullableMetadata {
	return &NullableMetadata{value: val, isSet: true}
}

func (v NullableMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
