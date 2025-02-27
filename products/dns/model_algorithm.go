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

	"fmt"
)

// Algorithm Algorithm used to generate signing keys (both Key Signing Keys and Zone Signing Keys).
type Algorithm string

// List of algorithm
const (
	ALGORITHM_RSASHA256 Algorithm = "RSASHA256"
)

func (v *Algorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Algorithm(value)
	for _, existing := range []Algorithm{"RSASHA256"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Algorithm", value)
}

// Ptr returns reference to algorithm value
func (v Algorithm) Ptr() *Algorithm {
	return &v
}

type NullableAlgorithm struct {
	value *Algorithm
	isSet bool
}

func (v NullableAlgorithm) Get() *Algorithm {
	return v.value
}

func (v *NullableAlgorithm) Set(val *Algorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlgorithm(val *Algorithm) *NullableAlgorithm {
	return &NullableAlgorithm{value: val, isSet: true}
}

func (v NullableAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
