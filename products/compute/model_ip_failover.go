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
)

// checks if the IPFailover type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPFailover{}

// IPFailover struct for IPFailover
type IPFailover struct {
	Ip      *string `json:"ip,omitempty"`
	NicUuid *string `json:"nicUuid,omitempty"`
}

// NewIPFailover instantiates a new IPFailover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPFailover() *IPFailover {
	this := IPFailover{}

	return &this
}

// NewIPFailoverWithDefaults instantiates a new IPFailover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPFailoverWithDefaults() *IPFailover {
	this := IPFailover{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *IPFailover) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPFailover) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *IPFailover) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *IPFailover) SetIp(v string) {
	o.Ip = &v
}

// GetNicUuid returns the NicUuid field value if set, zero value otherwise.
func (o *IPFailover) GetNicUuid() string {
	if o == nil || IsNil(o.NicUuid) {
		var ret string
		return ret
	}
	return *o.NicUuid
}

// GetNicUuidOk returns a tuple with the NicUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPFailover) GetNicUuidOk() (*string, bool) {
	if o == nil || IsNil(o.NicUuid) {
		return nil, false
	}
	return o.NicUuid, true
}

// HasNicUuid returns a boolean if a field has been set.
func (o *IPFailover) HasNicUuid() bool {
	if o != nil && !IsNil(o.NicUuid) {
		return true
	}

	return false
}

// SetNicUuid gets a reference to the given string and assigns it to the NicUuid field.
func (o *IPFailover) SetNicUuid(v string) {
	o.NicUuid = &v
}

func (o IPFailover) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.NicUuid) {
		toSerialize["nicUuid"] = o.NicUuid
	}
	return toSerialize, nil
}

type NullableIPFailover struct {
	value *IPFailover
	isSet bool
}

func (v NullableIPFailover) Get() *IPFailover {
	return v.value
}

func (v *NullableIPFailover) Set(val *IPFailover) {
	v.value = val
	v.isSet = true
}

func (v NullableIPFailover) IsSet() bool {
	return v.isSet
}

func (v *NullableIPFailover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPFailover(val *IPFailover) *NullableIPFailover {
	return &NullableIPFailover{value: val, isSet: true}
}

func (v NullableIPFailover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPFailover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
