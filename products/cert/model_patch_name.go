/*
 * Certificate Manager Service API
 *
 * Using the Certificate Manager Service, you can conveniently provision and manage SSL certificates  with IONOS services and your internal connected resources.   For the [Application Load Balancer](https://api.ionos.com/docs/cloud/v6/#Application-Load-Balancers-get-datacenters-datacenterId-applicationloadbalancers), you usually need a certificate to encrypt your HTTPS traffic. The service provides the basic functions of uploading and deleting your certificates for this purpose.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cert

import (
	"encoding/json"
)

// checks if the PatchName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchName{}

// PatchName struct for PatchName
type PatchName struct {
	// The new name.
	Name string `json:"name"`
}

// NewPatchName instantiates a new PatchName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchName(name string) *PatchName {
	this := PatchName{}

	this.Name = name

	return &this
}

// NewPatchNameWithDefaults instantiates a new PatchName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchNameWithDefaults() *PatchName {
	this := PatchName{}
	return &this
}

// GetName returns the Name field value
func (o *PatchName) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PatchName) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PatchName) SetName(v string) {
	o.Name = v
}

func (o PatchName) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullablePatchName struct {
	value *PatchName
	isSet bool
}

func (v NullablePatchName) Get() *PatchName {
	return v.value
}

func (v *NullablePatchName) Set(val *PatchName) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchName) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchName(val *PatchName) *NullablePatchName {
	return &NullablePatchName{value: val, isSet: true}
}

func (v NullablePatchName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
