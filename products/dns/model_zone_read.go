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

// checks if the ZoneRead type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneRead{}

// ZoneRead struct for ZoneRead
type ZoneRead struct {
	// The zone ID (UUID).
	Id         string                       `json:"id"`
	Type       string                       `json:"type"`
	Href       string                       `json:"href"`
	Metadata   MetadataWithStateNameservers `json:"metadata"`
	Properties Zone                         `json:"properties"`
}

// NewZoneRead instantiates a new ZoneRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneRead(id string, type_ string, href string, metadata MetadataWithStateNameservers, properties Zone) *ZoneRead {
	this := ZoneRead{}

	this.Id = id
	this.Type = type_
	this.Href = href
	this.Metadata = metadata
	this.Properties = properties

	return &this
}

// NewZoneReadWithDefaults instantiates a new ZoneRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneReadWithDefaults() *ZoneRead {
	this := ZoneRead{}
	return &this
}

// GetId returns the Id field value
func (o *ZoneRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ZoneRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ZoneRead) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ZoneRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ZoneRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ZoneRead) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *ZoneRead) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ZoneRead) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ZoneRead) SetHref(v string) {
	o.Href = v
}

// GetMetadata returns the Metadata field value
func (o *ZoneRead) GetMetadata() MetadataWithStateNameservers {
	if o == nil {
		var ret MetadataWithStateNameservers
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ZoneRead) GetMetadataOk() (*MetadataWithStateNameservers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ZoneRead) SetMetadata(v MetadataWithStateNameservers) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *ZoneRead) GetProperties() Zone {
	if o == nil {
		var ret Zone
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ZoneRead) GetPropertiesOk() (*Zone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *ZoneRead) SetProperties(v Zone) {
	o.Properties = v
}

func (o ZoneRead) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneRead) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["href"] = o.Href
	toSerialize["metadata"] = o.Metadata
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableZoneRead struct {
	value *ZoneRead
	isSet bool
}

func (v NullableZoneRead) Get() *ZoneRead {
	return v.value
}

func (v *NullableZoneRead) Set(val *ZoneRead) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneRead) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneRead(val *ZoneRead) *NullableZoneRead {
	return &NullableZoneRead{value: val, isSet: true}
}

func (v NullableZoneRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
