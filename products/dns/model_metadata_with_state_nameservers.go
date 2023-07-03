/*
 * IONOS Cloud - DNS API
 *
 * DNS API Specification
 *
 * API version: 1.2.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"

	"time"
)

// MetadataWithStateNameservers struct for MetadataWithStateNameservers
type MetadataWithStateNameservers struct {
	// The date of the last change formatted as yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	LastModifiedDate *IonosTime `json:"lastModifiedDate,omitempty"`
	// The date of creation of the zone formatted as yyyy-MM-dd'T'HH:mm:ss.SSS'Z'.
	CreatedDate *IonosTime         `json:"createdDate,omitempty"`
	State       *ProvisioningState `json:"state"`
	// The list of nameservers associated to the zone
	Nameservers *[]string `json:"nameservers"`
}

// NewMetadataWithStateNameservers instantiates a new MetadataWithStateNameservers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataWithStateNameservers(state ProvisioningState, nameservers []string) *MetadataWithStateNameservers {
	this := MetadataWithStateNameservers{}

	this.State = &state
	this.Nameservers = &nameservers

	return &this
}

// NewMetadataWithStateNameserversWithDefaults instantiates a new MetadataWithStateNameservers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataWithStateNameserversWithDefaults() *MetadataWithStateNameservers {
	this := MetadataWithStateNameservers{}
	return &this
}

// GetLastModifiedDate returns the LastModifiedDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *MetadataWithStateNameservers) GetLastModifiedDate() *time.Time {
	if o == nil {
		return nil
	}

	if o.LastModifiedDate == nil {
		return nil
	}
	return &o.LastModifiedDate.Time

}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataWithStateNameservers) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.LastModifiedDate == nil {
		return nil, false
	}
	return &o.LastModifiedDate.Time, true

}

// SetLastModifiedDate sets field value
func (o *MetadataWithStateNameservers) SetLastModifiedDate(v time.Time) {

	o.LastModifiedDate = &IonosTime{v}

}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *MetadataWithStateNameservers) HasLastModifiedDate() bool {
	if o != nil && o.LastModifiedDate != nil {
		return true
	}

	return false
}

// GetCreatedDate returns the CreatedDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *MetadataWithStateNameservers) GetCreatedDate() *time.Time {
	if o == nil {
		return nil
	}

	if o.CreatedDate == nil {
		return nil
	}
	return &o.CreatedDate.Time

}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataWithStateNameservers) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.CreatedDate == nil {
		return nil, false
	}
	return &o.CreatedDate.Time, true

}

// SetCreatedDate sets field value
func (o *MetadataWithStateNameservers) SetCreatedDate(v time.Time) {

	o.CreatedDate = &IonosTime{v}

}

// HasCreatedDate returns a boolean if a field has been set.
func (o *MetadataWithStateNameservers) HasCreatedDate() bool {
	if o != nil && o.CreatedDate != nil {
		return true
	}

	return false
}

// GetState returns the State field value
// If the value is explicit nil, the zero value for ProvisioningState will be returned
func (o *MetadataWithStateNameservers) GetState() *ProvisioningState {
	if o == nil {
		return nil
	}

	return o.State

}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataWithStateNameservers) GetStateOk() (*ProvisioningState, bool) {
	if o == nil {
		return nil, false
	}

	return o.State, true
}

// SetState sets field value
func (o *MetadataWithStateNameservers) SetState(v ProvisioningState) {

	o.State = &v

}

// HasState returns a boolean if a field has been set.
func (o *MetadataWithStateNameservers) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// GetNameservers returns the Nameservers field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *MetadataWithStateNameservers) GetNameservers() *[]string {
	if o == nil {
		return nil
	}

	return o.Nameservers

}

// GetNameserversOk returns a tuple with the Nameservers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataWithStateNameservers) GetNameserversOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Nameservers, true
}

// SetNameservers sets field value
func (o *MetadataWithStateNameservers) SetNameservers(v []string) {

	o.Nameservers = &v

}

// HasNameservers returns a boolean if a field has been set.
func (o *MetadataWithStateNameservers) HasNameservers() bool {
	if o != nil && o.Nameservers != nil {
		return true
	}

	return false
}

func (o MetadataWithStateNameservers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastModifiedDate != nil {
		toSerialize["lastModifiedDate"] = o.LastModifiedDate
	}

	if o.CreatedDate != nil {
		toSerialize["createdDate"] = o.CreatedDate
	}

	if o.State != nil {
		toSerialize["state"] = o.State
	}

	if o.Nameservers != nil {
		toSerialize["nameservers"] = o.Nameservers
	}

	return json.Marshal(toSerialize)
}

type NullableMetadataWithStateNameservers struct {
	value *MetadataWithStateNameservers
	isSet bool
}

func (v NullableMetadataWithStateNameservers) Get() *MetadataWithStateNameservers {
	return v.value
}

func (v *NullableMetadataWithStateNameservers) Set(val *MetadataWithStateNameservers) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataWithStateNameservers) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataWithStateNameservers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataWithStateNameservers(val *MetadataWithStateNameservers) *NullableMetadataWithStateNameservers {
	return &NullableMetadataWithStateNameservers{value: val, isSet: true}
}

func (v NullableMetadataWithStateNameservers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataWithStateNameservers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
