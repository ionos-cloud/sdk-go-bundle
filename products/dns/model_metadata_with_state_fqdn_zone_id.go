/*
 * IONOS Cloud - DNS API
 *
 * Cloud DNS service helps IONOS Cloud customers to automate DNS Zone and Record management.
 *
 * API version: 1.16.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"

	"time"
)

// checks if the MetadataWithStateFqdnZoneId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataWithStateFqdnZoneId{}

// MetadataWithStateFqdnZoneId struct for MetadataWithStateFqdnZoneId
type MetadataWithStateFqdnZoneId struct {
	// The creation date formatted as yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	CreatedDate *IonosTime `json:"createdDate,omitempty"`
	// Unique name of the identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The unique ID of the user who created the resource.
	CreatedByUserId *string `json:"createdByUserId,omitempty"`
	// The date of the last change formatted as yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	LastModifiedDate *IonosTime `json:"lastModifiedDate,omitempty"`
	// Unique name of the identity that created the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// The unique ID of the user who last modified the resource.
	LastModifiedByUserId *string           `json:"lastModifiedByUserId,omitempty"`
	State                ProvisioningState `json:"state"`
	// A fully qualified domain name. FQDN consists of two parts - the hostname and the domain name.
	Fqdn string `json:"fqdn"`
	// The ID (UUID) of the DNS zone of which record belongs to.
	ZoneId string `json:"zoneId"`
}

// NewMetadataWithStateFqdnZoneId instantiates a new MetadataWithStateFqdnZoneId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataWithStateFqdnZoneId(state ProvisioningState, fqdn string, zoneId string) *MetadataWithStateFqdnZoneId {
	this := MetadataWithStateFqdnZoneId{}

	this.State = state
	this.Fqdn = fqdn
	this.ZoneId = zoneId

	return &this
}

// NewMetadataWithStateFqdnZoneIdWithDefaults instantiates a new MetadataWithStateFqdnZoneId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataWithStateFqdnZoneIdWithDefaults() *MetadataWithStateFqdnZoneId {
	this := MetadataWithStateFqdnZoneId{}
	return &this
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *MetadataWithStateFqdnZoneId) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return o.CreatedDate.Time
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataWithStateFqdnZoneId) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return &o.CreatedDate.Time, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *MetadataWithStateFqdnZoneId) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *MetadataWithStateFqdnZoneId) SetCreatedDate(v time.Time) {
	o.CreatedDate = &IonosTime{v}
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *MetadataWithStateFqdnZoneId) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataWithStateFqdnZoneId) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MetadataWithStateFqdnZoneId) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *MetadataWithStateFqdnZoneId) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *MetadataWithStateFqdnZoneId) GetCreatedByUserId() string {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret string
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataWithStateFqdnZoneId) GetCreatedByUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *MetadataWithStateFqdnZoneId) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given string and assigns it to the CreatedByUserId field.
func (o *MetadataWithStateFqdnZoneId) SetCreatedByUserId(v string) {
	o.CreatedByUserId = &v
}

// GetLastModifiedDate returns the LastModifiedDate field value if set, zero value otherwise.
func (o *MetadataWithStateFqdnZoneId) GetLastModifiedDate() time.Time {
	if o == nil || IsNil(o.LastModifiedDate) {
		var ret time.Time
		return ret
	}
	return o.LastModifiedDate.Time
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataWithStateFqdnZoneId) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedDate) {
		return nil, false
	}
	return &o.LastModifiedDate.Time, true
}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *MetadataWithStateFqdnZoneId) HasLastModifiedDate() bool {
	if o != nil && !IsNil(o.LastModifiedDate) {
		return true
	}

	return false
}

// SetLastModifiedDate gets a reference to the given time.Time and assigns it to the LastModifiedDate field.
func (o *MetadataWithStateFqdnZoneId) SetLastModifiedDate(v time.Time) {
	o.LastModifiedDate = &IonosTime{v}
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *MetadataWithStateFqdnZoneId) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataWithStateFqdnZoneId) GetLastModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedBy) {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *MetadataWithStateFqdnZoneId) HasLastModifiedBy() bool {
	if o != nil && !IsNil(o.LastModifiedBy) {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given string and assigns it to the LastModifiedBy field.
func (o *MetadataWithStateFqdnZoneId) SetLastModifiedBy(v string) {
	o.LastModifiedBy = &v
}

// GetLastModifiedByUserId returns the LastModifiedByUserId field value if set, zero value otherwise.
func (o *MetadataWithStateFqdnZoneId) GetLastModifiedByUserId() string {
	if o == nil || IsNil(o.LastModifiedByUserId) {
		var ret string
		return ret
	}
	return *o.LastModifiedByUserId
}

// GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataWithStateFqdnZoneId) GetLastModifiedByUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedByUserId) {
		return nil, false
	}
	return o.LastModifiedByUserId, true
}

// HasLastModifiedByUserId returns a boolean if a field has been set.
func (o *MetadataWithStateFqdnZoneId) HasLastModifiedByUserId() bool {
	if o != nil && !IsNil(o.LastModifiedByUserId) {
		return true
	}

	return false
}

// SetLastModifiedByUserId gets a reference to the given string and assigns it to the LastModifiedByUserId field.
func (o *MetadataWithStateFqdnZoneId) SetLastModifiedByUserId(v string) {
	o.LastModifiedByUserId = &v
}

// GetState returns the State field value
func (o *MetadataWithStateFqdnZoneId) GetState() ProvisioningState {
	if o == nil {
		var ret ProvisioningState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *MetadataWithStateFqdnZoneId) GetStateOk() (*ProvisioningState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *MetadataWithStateFqdnZoneId) SetState(v ProvisioningState) {
	o.State = v
}

// GetFqdn returns the Fqdn field value
func (o *MetadataWithStateFqdnZoneId) GetFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value
// and a boolean to check if the value has been set.
func (o *MetadataWithStateFqdnZoneId) GetFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqdn, true
}

// SetFqdn sets field value
func (o *MetadataWithStateFqdnZoneId) SetFqdn(v string) {
	o.Fqdn = v
}

// GetZoneId returns the ZoneId field value
func (o *MetadataWithStateFqdnZoneId) GetZoneId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value
// and a boolean to check if the value has been set.
func (o *MetadataWithStateFqdnZoneId) GetZoneIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneId, true
}

// SetZoneId sets field value
func (o *MetadataWithStateFqdnZoneId) SetZoneId(v string) {
	o.ZoneId = v
}

func (o MetadataWithStateFqdnZoneId) ToMap() (map[string]interface{}, error) {
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
	toSerialize["state"] = o.State
	toSerialize["fqdn"] = o.Fqdn
	toSerialize["zoneId"] = o.ZoneId
	return toSerialize, nil
}

type NullableMetadataWithStateFqdnZoneId struct {
	value *MetadataWithStateFqdnZoneId
	isSet bool
}

func (v NullableMetadataWithStateFqdnZoneId) Get() *MetadataWithStateFqdnZoneId {
	return v.value
}

func (v *NullableMetadataWithStateFqdnZoneId) Set(val *MetadataWithStateFqdnZoneId) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataWithStateFqdnZoneId) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataWithStateFqdnZoneId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataWithStateFqdnZoneId(val *MetadataWithStateFqdnZoneId) *NullableMetadataWithStateFqdnZoneId {
	return &NullableMetadataWithStateFqdnZoneId{value: val, isSet: true}
}

func (v NullableMetadataWithStateFqdnZoneId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataWithStateFqdnZoneId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
