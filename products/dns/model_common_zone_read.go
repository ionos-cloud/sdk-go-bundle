/*
 * IONOS Cloud - DNS API
 *
 * Cloud DNS service helps IONOS Cloud customers to automate DNS Zone and Record management.
 *
 * API version: 1.17.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the CommonZoneRead type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonZoneRead{}

// CommonZoneRead struct for CommonZoneRead
type CommonZoneRead struct {
	// The zone ID (UUID).
	Id       string                       `json:"id"`
	Type     string                       `json:"type"`
	Href     string                       `json:"href"`
	Metadata MetadataWithStateNameservers `json:"metadata"`
}

// NewCommonZoneRead instantiates a new CommonZoneRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonZoneRead(id string, type_ string, href string, metadata MetadataWithStateNameservers) *CommonZoneRead {
	this := CommonZoneRead{}

	this.Id = id
	this.Type = type_
	this.Href = href
	this.Metadata = metadata

	return &this
}

// NewCommonZoneReadWithDefaults instantiates a new CommonZoneRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonZoneReadWithDefaults() *CommonZoneRead {
	this := CommonZoneRead{}
	return &this
}

// GetId returns the Id field value
func (o *CommonZoneRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CommonZoneRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CommonZoneRead) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *CommonZoneRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CommonZoneRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CommonZoneRead) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *CommonZoneRead) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *CommonZoneRead) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *CommonZoneRead) SetHref(v string) {
	o.Href = v
}

// GetMetadata returns the Metadata field value
func (o *CommonZoneRead) GetMetadata() MetadataWithStateNameservers {
	if o == nil {
		var ret MetadataWithStateNameservers
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *CommonZoneRead) GetMetadataOk() (*MetadataWithStateNameservers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *CommonZoneRead) SetMetadata(v MetadataWithStateNameservers) {
	o.Metadata = v
}

func (o CommonZoneRead) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["href"] = o.Href
	toSerialize["metadata"] = o.Metadata
	return toSerialize, nil
}

type NullableCommonZoneRead struct {
	value *CommonZoneRead
	isSet bool
}

func (v NullableCommonZoneRead) Get() *CommonZoneRead {
	return v.value
}

func (v *NullableCommonZoneRead) Set(val *CommonZoneRead) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonZoneRead) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonZoneRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonZoneRead(val *CommonZoneRead) *NullableCommonZoneRead {
	return &NullableCommonZoneRead{value: val, isSet: true}
}

func (v NullableCommonZoneRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonZoneRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
