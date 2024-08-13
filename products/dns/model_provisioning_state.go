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

// ProvisioningState The list of possible provisioning states in which DNS resource could be at the specific time. * AVAILABLE - resource exists and is healthy. * PROVISIONING - resource is being created or updated. * DESTROYING - delete command was issued, the resource is being deleted. * FAILED - creation of the resource failed.
type ProvisioningState string

// List of ProvisioningState
const (
	PROVISIONINGSTATE_PROVISIONING ProvisioningState = "PROVISIONING"
	PROVISIONINGSTATE_DESTROYING   ProvisioningState = "DESTROYING"
	PROVISIONINGSTATE_AVAILABLE    ProvisioningState = "AVAILABLE"
	PROVISIONINGSTATE_FAILED       ProvisioningState = "FAILED"
)

func (v *ProvisioningState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProvisioningState(value)
	for _, existing := range []ProvisioningState{"PROVISIONING", "DESTROYING", "AVAILABLE", "FAILED"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProvisioningState", value)
}

// Ptr returns reference to ProvisioningState value
func (v ProvisioningState) Ptr() *ProvisioningState {
	return &v
}

type NullableProvisioningState struct {
	value *ProvisioningState
	isSet bool
}

func (v NullableProvisioningState) Get() *ProvisioningState {
	return v.value
}

func (v *NullableProvisioningState) Set(val *ProvisioningState) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningState) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningState(val *ProvisioningState) *NullableProvisioningState {
	return &NullableProvisioningState{value: val, isSet: true}
}

func (v NullableProvisioningState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
