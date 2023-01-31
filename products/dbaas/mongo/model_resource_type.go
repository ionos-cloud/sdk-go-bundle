/*
 * IONOS DBaaS MongoDB REST API
 *
 * With IONOS Cloud Database as a Service, you have the ability to quickly set up and manage a MongoDB database. You can also delete clusters, manage backups and users via the API.   MongoDB is an open source, cross-platform, document-oriented database program. Classified as a NoSQL database program, it uses JSON-like documents with optional schemas.  The MongoDB API allows you to create additional database clusters or modify existing ones. Both tools, the Data Center Designer (DCD) and the API use the same concepts consistently and are well suited for smooth and intuitive use.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongo

import (
	"encoding/json"

	"fmt"
)

// ResourceType The resource type.
type ResourceType string

// List of ResourceType
const (
	RESOURCETYPE_COLLECTION ResourceType = "collection"
	RESOURCETYPE_CLUSTER    ResourceType = "cluster"
	RESOURCETYPE_USER       ResourceType = "user"
	RESOURCETYPE_SNAPSHOT   ResourceType = "snapshot"
)

func (v *ResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceType(value)
	for _, existing := range []ResourceType{"collection", "cluster", "user", "snapshot"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceType", value)
}

// Ptr returns reference to ResourceType value
func (v ResourceType) Ptr() *ResourceType {
	return &v
}

type NullableResourceType struct {
	value *ResourceType
	isSet bool
}

func (v NullableResourceType) Get() *ResourceType {
	return v.value
}

func (v *NullableResourceType) Set(val *ResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceType(val *ResourceType) *NullableResourceType {
	return &NullableResourceType{value: val, isSet: true}
}

func (v NullableResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
