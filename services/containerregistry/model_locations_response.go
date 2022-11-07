/*
 * Container Registry service
 *
 * Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls.
 *
 * API version: 1.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerregistry

import (
	"encoding/json"
)

// LocationsResponse struct for LocationsResponse
type LocationsResponse struct {
	Href  *string     `json:"href,omitempty"`
	Id    *string     `json:"id,omitempty"`
	Items *[]Location `json:"items,omitempty"`
	Type  *string     `json:"type,omitempty"`
}

// NewLocationsResponse instantiates a new LocationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationsResponse() *LocationsResponse {
	this := LocationsResponse{}

	return &this
}

// NewLocationsResponseWithDefaults instantiates a new LocationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationsResponseWithDefaults() *LocationsResponse {
	this := LocationsResponse{}
	return &this
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LocationsResponse) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationsResponse) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *LocationsResponse) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *LocationsResponse) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LocationsResponse) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationsResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *LocationsResponse) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *LocationsResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []Location will be returned
func (o *LocationsResponse) GetItems() *[]Location {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationsResponse) GetItemsOk() (*[]Location, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *LocationsResponse) SetItems(v []Location) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *LocationsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LocationsResponse) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationsResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *LocationsResponse) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *LocationsResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

func (o LocationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	toSerialize["items"] = o.Items

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	return json.Marshal(toSerialize)
}

type NullableLocationsResponse struct {
	value *LocationsResponse
	isSet bool
}

func (v NullableLocationsResponse) Get() *LocationsResponse {
	return v.value
}

func (v *NullableLocationsResponse) Set(val *LocationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationsResponse(val *LocationsResponse) *NullableLocationsResponse {
	return &NullableLocationsResponse{value: val, isSet: true}
}

func (v NullableLocationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
