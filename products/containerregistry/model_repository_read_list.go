/*
 * Container Registry service
 *
 * ## Overview Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls. ## Changelog ### 1.1.0  - Added new endpoints for Repositories  - Added new endpoints for Artifacts  - Added new endpoints for Vulnerabilities  - Added registry vulnerabilityScanning feature ### 1.2.0  - Added registry `apiSubnetAllowList` ### 1.2.1  - Amended `apiSubnetAllowList` Regex
 *
 * API version: 1.2.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerregistry

import (
	"encoding/json"
)

// checks if the RepositoryReadList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryReadList{}

// RepositoryReadList struct for RepositoryReadList
type RepositoryReadList struct {
	Id    string           `json:"id"`
	Type  string           `json:"type"`
	Href  string           `json:"href"`
	Items []RepositoryRead `json:"items,omitempty"`
	// The offset specified in the request (if none was specified, the default offset is 0).
	Offset int32 `json:"offset"`
	// The limit specified in the request (if none was specified, use the endpoint's default pagination limit).
	Limit int32 `json:"limit"`
	Links Links `json:"_links"`
}

// NewRepositoryReadList instantiates a new RepositoryReadList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryReadList(id string, type_ string, href string, offset int32, limit int32, links Links) *RepositoryReadList {
	this := RepositoryReadList{}

	this.Id = id
	this.Type = type_
	this.Href = href
	this.Offset = offset
	this.Limit = limit
	this.Links = links

	return &this
}

// NewRepositoryReadListWithDefaults instantiates a new RepositoryReadList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryReadListWithDefaults() *RepositoryReadList {
	this := RepositoryReadList{}
	return &this
}

// GetId returns the Id field value
func (o *RepositoryReadList) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RepositoryReadList) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RepositoryReadList) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *RepositoryReadList) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RepositoryReadList) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RepositoryReadList) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *RepositoryReadList) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *RepositoryReadList) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *RepositoryReadList) SetHref(v string) {
	o.Href = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *RepositoryReadList) GetItems() []RepositoryRead {
	if o == nil || IsNil(o.Items) {
		var ret []RepositoryRead
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryReadList) GetItemsOk() ([]RepositoryRead, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *RepositoryReadList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []RepositoryRead and assigns it to the Items field.
func (o *RepositoryReadList) SetItems(v []RepositoryRead) {
	o.Items = v
}

// GetOffset returns the Offset field value
func (o *RepositoryReadList) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *RepositoryReadList) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *RepositoryReadList) SetOffset(v int32) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *RepositoryReadList) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *RepositoryReadList) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *RepositoryReadList) SetLimit(v int32) {
	o.Limit = v
}

// GetLinks returns the Links field value
func (o *RepositoryReadList) GetLinks() Links {
	if o == nil {
		var ret Links
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RepositoryReadList) GetLinksOk() (*Links, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RepositoryReadList) SetLinks(v Links) {
	o.Links = v
}

func (o RepositoryReadList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryReadList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["href"] = o.Href
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	toSerialize["offset"] = o.Offset
	toSerialize["limit"] = o.Limit
	toSerialize["_links"] = o.Links
	return toSerialize, nil
}

type NullableRepositoryReadList struct {
	value *RepositoryReadList
	isSet bool
}

func (v NullableRepositoryReadList) Get() *RepositoryReadList {
	return v.value
}

func (v *NullableRepositoryReadList) Set(val *RepositoryReadList) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryReadList) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryReadList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryReadList(val *RepositoryReadList) *NullableRepositoryReadList {
	return &NullableRepositoryReadList{value: val, isSet: true}
}

func (v NullableRepositoryReadList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryReadList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
