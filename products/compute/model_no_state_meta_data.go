/*
 * CLOUD API
 *
 *  IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package compute

import (
	"encoding/json"

	"time"
)

// checks if the NoStateMetaData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NoStateMetaData{}

// NoStateMetaData struct for NoStateMetaData
type NoStateMetaData struct {
	// Resource's Entity Tag as defined in http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.11  Entity Tag is also added as an 'ETag response header to requests which don't use 'depth' parameter.
	Etag *string `json:"etag,omitempty"`
	// The time when the resource was created.
	CreatedDate *IonosTime `json:"createdDate,omitempty"`
	// The user who has created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The unique ID of the user who created the resource.
	CreatedByUserId *string `json:"createdByUserId,omitempty"`
	// The last time the resource was modified.
	LastModifiedDate *IonosTime `json:"lastModifiedDate,omitempty"`
	// The user who last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// The unique ID of the user who last modified the resource.
	LastModifiedByUserId *string `json:"lastModifiedByUserId,omitempty"`
}

// NewNoStateMetaData instantiates a new NoStateMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoStateMetaData() *NoStateMetaData {
	this := NoStateMetaData{}

	return &this
}

// NewNoStateMetaDataWithDefaults instantiates a new NoStateMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoStateMetaDataWithDefaults() *NoStateMetaData {
	this := NoStateMetaData{}
	return &this
}

// GetEtag returns the Etag field value if set, zero value otherwise.
func (o *NoStateMetaData) GetEtag() string {
	if o == nil || IsNil(o.Etag) {
		var ret string
		return ret
	}
	return *o.Etag
}

// GetEtagOk returns a tuple with the Etag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoStateMetaData) GetEtagOk() (*string, bool) {
	if o == nil || IsNil(o.Etag) {
		return nil, false
	}
	return o.Etag, true
}

// HasEtag returns a boolean if a field has been set.
func (o *NoStateMetaData) HasEtag() bool {
	if o != nil && !IsNil(o.Etag) {
		return true
	}

	return false
}

// SetEtag gets a reference to the given string and assigns it to the Etag field.
func (o *NoStateMetaData) SetEtag(v string) {
	o.Etag = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *NoStateMetaData) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return o.CreatedDate.Time
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoStateMetaData) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return &o.CreatedDate.Time, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *NoStateMetaData) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *NoStateMetaData) SetCreatedDate(v time.Time) {
	o.CreatedDate = &IonosTime{v}
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *NoStateMetaData) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoStateMetaData) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *NoStateMetaData) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *NoStateMetaData) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *NoStateMetaData) GetCreatedByUserId() string {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret string
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoStateMetaData) GetCreatedByUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *NoStateMetaData) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given string and assigns it to the CreatedByUserId field.
func (o *NoStateMetaData) SetCreatedByUserId(v string) {
	o.CreatedByUserId = &v
}

// GetLastModifiedDate returns the LastModifiedDate field value if set, zero value otherwise.
func (o *NoStateMetaData) GetLastModifiedDate() time.Time {
	if o == nil || IsNil(o.LastModifiedDate) {
		var ret time.Time
		return ret
	}
	return o.LastModifiedDate.Time
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoStateMetaData) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedDate) {
		return nil, false
	}
	return &o.LastModifiedDate.Time, true
}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *NoStateMetaData) HasLastModifiedDate() bool {
	if o != nil && !IsNil(o.LastModifiedDate) {
		return true
	}

	return false
}

// SetLastModifiedDate gets a reference to the given time.Time and assigns it to the LastModifiedDate field.
func (o *NoStateMetaData) SetLastModifiedDate(v time.Time) {
	o.LastModifiedDate = &IonosTime{v}
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *NoStateMetaData) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoStateMetaData) GetLastModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedBy) {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *NoStateMetaData) HasLastModifiedBy() bool {
	if o != nil && !IsNil(o.LastModifiedBy) {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given string and assigns it to the LastModifiedBy field.
func (o *NoStateMetaData) SetLastModifiedBy(v string) {
	o.LastModifiedBy = &v
}

// GetLastModifiedByUserId returns the LastModifiedByUserId field value if set, zero value otherwise.
func (o *NoStateMetaData) GetLastModifiedByUserId() string {
	if o == nil || IsNil(o.LastModifiedByUserId) {
		var ret string
		return ret
	}
	return *o.LastModifiedByUserId
}

// GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoStateMetaData) GetLastModifiedByUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedByUserId) {
		return nil, false
	}
	return o.LastModifiedByUserId, true
}

// HasLastModifiedByUserId returns a boolean if a field has been set.
func (o *NoStateMetaData) HasLastModifiedByUserId() bool {
	if o != nil && !IsNil(o.LastModifiedByUserId) {
		return true
	}

	return false
}

// SetLastModifiedByUserId gets a reference to the given string and assigns it to the LastModifiedByUserId field.
func (o *NoStateMetaData) SetLastModifiedByUserId(v string) {
	o.LastModifiedByUserId = &v
}

func (o NoStateMetaData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NoStateMetaData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Etag) {
		toSerialize["etag"] = o.Etag
	}
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
	return toSerialize, nil
}

type NullableNoStateMetaData struct {
	value *NoStateMetaData
	isSet bool
}

func (v NullableNoStateMetaData) Get() *NoStateMetaData {
	return v.value
}

func (v *NullableNoStateMetaData) Set(val *NoStateMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableNoStateMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableNoStateMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoStateMetaData(val *NoStateMetaData) *NullableNoStateMetaData {
	return &NullableNoStateMetaData{value: val, isSet: true}
}

func (v NullableNoStateMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoStateMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
