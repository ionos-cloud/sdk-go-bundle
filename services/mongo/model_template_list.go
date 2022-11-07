/*
 * IONOS DBaaS MongoDB REST API
 *
 * With IONOS Cloud Database as a Service, you have the ability to quickly set up and manage a MongoDB database. You can also delete clusters, manage backups and users via the API.   MongoDB is an open source, cross-platform, document-oriented database program. Classified as a NoSQL database program, it uses JSON-like documents with optional schemas.  The MongoDB API allows you to create additional database clusters or modify existing ones. Both tools, the Data Center Designer (DCD) and the API use the same concepts consistently and are well suited for smooth and intuitive use.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongo

import (
	"encoding/json"
)

// TemplateList The list of MongoDB templates.
type TemplateList struct {
	Items *[]TemplateResponse `json:"items,omitempty"`
}

// NewTemplateList instantiates a new TemplateList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateList() *TemplateList {
	this := TemplateList{}

	return &this
}

// NewTemplateListWithDefaults instantiates a new TemplateList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateListWithDefaults() *TemplateList {
	this := TemplateList{}
	return &this
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []TemplateResponse will be returned
func (o *TemplateList) GetItems() *[]TemplateResponse {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateList) GetItemsOk() (*[]TemplateResponse, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *TemplateList) SetItems(v []TemplateResponse) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *TemplateList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o TemplateList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	return json.Marshal(toSerialize)
}

type NullableTemplateList struct {
	value *TemplateList
	isSet bool
}

func (v NullableTemplateList) Get() *TemplateList {
	return v.value
}

func (v *NullableTemplateList) Set(val *TemplateList) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateList) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateList(val *TemplateList) *NullableTemplateList {
	return &NullableTemplateList{value: val, isSet: true}
}

func (v NullableTemplateList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
