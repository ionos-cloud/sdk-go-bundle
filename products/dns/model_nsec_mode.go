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

	"fmt"
)

// NsecMode NSEC mode.
type NsecMode string

// List of nsecMode
const (
	NSECMODE_NSEC  NsecMode = "NSEC"
	NSECMODE_NSEC3 NsecMode = "NSEC3"
)

func (v *NsecMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NsecMode(value)
	for _, existing := range []NsecMode{"NSEC", "NSEC3"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NsecMode", value)
}

// Ptr returns reference to nsecMode value
func (v NsecMode) Ptr() *NsecMode {
	return &v
}

type NullableNsecMode struct {
	value *NsecMode
	isSet bool
}

func (v NullableNsecMode) Get() *NsecMode {
	return v.value
}

func (v *NullableNsecMode) Set(val *NsecMode) {
	v.value = val
	v.isSet = true
}

func (v NullableNsecMode) IsSet() bool {
	return v.isSet
}

func (v *NullableNsecMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsecMode(val *NsecMode) *NullableNsecMode {
	return &NullableNsecMode{value: val, isSet: true}
}

func (v NullableNsecMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsecMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
