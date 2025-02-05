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

// ZskBits Zone signing key length in bits.
type ZskBits int32

// List of zskBits
const (
	ZSKBITS__1024 ZskBits = 1024
	ZSKBITS__2048 ZskBits = 2048
	ZSKBITS__4096 ZskBits = 4096
)

func (v *ZskBits) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ZskBits(value)
	for _, existing := range []ZskBits{1024, 2048, 4096} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ZskBits", value)
}

// Ptr returns reference to zskBits value
func (v ZskBits) Ptr() *ZskBits {
	return &v
}

type NullableZskBits struct {
	value *ZskBits
	isSet bool
}

func (v NullableZskBits) Get() *ZskBits {
	return v.value
}

func (v *NullableZskBits) Set(val *ZskBits) {
	v.value = val
	v.isSet = true
}

func (v NullableZskBits) IsSet() bool {
	return v.isSet
}

func (v *NullableZskBits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZskBits(val *ZskBits) *NullableZskBits {
	return &NullableZskBits{value: val, isSet: true}
}

func (v NullableZskBits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZskBits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
