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
)

// checks if the SecondaryZoneRead type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecondaryZoneRead{}

// SecondaryZoneRead struct for SecondaryZoneRead
type SecondaryZoneRead struct {
	// The zone ID (UUID).
	Id         string                       `json:"id"`
	Type       string                       `json:"type"`
	Href       string                       `json:"href"`
	Metadata   MetadataWithStateNameservers `json:"metadata"`
	Properties SecondaryZone                `json:"properties"`
}

// NewSecondaryZoneRead instantiates a new SecondaryZoneRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecondaryZoneRead(id string, type_ string, href string, metadata MetadataWithStateNameservers, properties SecondaryZone) *SecondaryZoneRead {
	this := SecondaryZoneRead{}

	this.Id = id
	this.Type = type_
	this.Href = href
	this.Metadata = metadata
	this.Properties = properties

	return &this
}

// NewSecondaryZoneReadWithDefaults instantiates a new SecondaryZoneRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecondaryZoneReadWithDefaults() *SecondaryZoneRead {
	this := SecondaryZoneRead{}
	return &this
}

// GetId returns the Id field value
func (o *SecondaryZoneRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecondaryZoneRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecondaryZoneRead) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *SecondaryZoneRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecondaryZoneRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SecondaryZoneRead) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *SecondaryZoneRead) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *SecondaryZoneRead) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *SecondaryZoneRead) SetHref(v string) {
	o.Href = v
}

// GetMetadata returns the Metadata field value
func (o *SecondaryZoneRead) GetMetadata() MetadataWithStateNameservers {
	if o == nil {
		var ret MetadataWithStateNameservers
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *SecondaryZoneRead) GetMetadataOk() (*MetadataWithStateNameservers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *SecondaryZoneRead) SetMetadata(v MetadataWithStateNameservers) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *SecondaryZoneRead) GetProperties() SecondaryZone {
	if o == nil {
		var ret SecondaryZone
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *SecondaryZoneRead) GetPropertiesOk() (*SecondaryZone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *SecondaryZoneRead) SetProperties(v SecondaryZone) {
	o.Properties = v
}

func (o SecondaryZoneRead) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecondaryZoneRead) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsZero(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsZero(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsZero(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsZero(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsZero(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableSecondaryZoneRead struct {
	value *SecondaryZoneRead
	isSet bool
}

func (v NullableSecondaryZoneRead) Get() *SecondaryZoneRead {
	return v.value
}

func (v *NullableSecondaryZoneRead) Set(val *SecondaryZoneRead) {
	v.value = val
	v.isSet = true
}

func (v NullableSecondaryZoneRead) IsSet() bool {
	return v.isSet
}

func (v *NullableSecondaryZoneRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecondaryZoneRead(val *SecondaryZoneRead) *NullableSecondaryZoneRead {
	return &NullableSecondaryZoneRead{value: val, isSet: true}
}

func (v NullableSecondaryZoneRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecondaryZoneRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
