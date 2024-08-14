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

// checks if the NatGateways type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NatGateways{}

// NatGateways struct for NatGateways
type NatGateways struct {
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path).
	Href *string `json:"href,omitempty"`
	// Array of items in the collection.
	Items []NatGateway `json:"items,omitempty"`
	// The offset (if specified in the request).
	Offset *float32 `json:"offset,omitempty"`
	// The limit (if specified in the request).
	Limit *float32         `json:"limit,omitempty"`
	Links *PaginationLinks `json:"_links,omitempty"`
}

// NewNatGateways instantiates a new NatGateways object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNatGateways() *NatGateways {
	this := NatGateways{}

	return &this
}

// NewNatGatewaysWithDefaults instantiates a new NatGateways object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNatGatewaysWithDefaults() *NatGateways {
	this := NatGateways{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NatGateways) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NatGateways) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NatGateways) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NatGateways) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NatGateways) GetType() Type {
	if o == nil || IsNil(o.Type) {
		var ret Type
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NatGateways) GetTypeOk() (*Type, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NatGateways) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given Type and assigns it to the Type field.
func (o *NatGateways) SetType(v Type) {
	o.Type = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *NatGateways) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NatGateways) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *NatGateways) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *NatGateways) SetHref(v string) {
	o.Href = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *NatGateways) GetItems() []NatGateway {
	if o == nil || IsNil(o.Items) {
		var ret []NatGateway
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NatGateways) GetItemsOk() ([]NatGateway, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *NatGateways) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []NatGateway and assigns it to the Items field.
func (o *NatGateways) SetItems(v []NatGateway) {
	o.Items = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *NatGateways) GetOffset() float32 {
	if o == nil || IsNil(o.Offset) {
		var ret float32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NatGateways) GetOffsetOk() (*float32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *NatGateways) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given float32 and assigns it to the Offset field.
func (o *NatGateways) SetOffset(v float32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *NatGateways) GetLimit() float32 {
	if o == nil || IsNil(o.Limit) {
		var ret float32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NatGateways) GetLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *NatGateways) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given float32 and assigns it to the Limit field.
func (o *NatGateways) SetLimit(v float32) {
	o.Limit = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *NatGateways) GetLinks() PaginationLinks {
	if o == nil || IsNil(o.Links) {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NatGateways) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *NatGateways) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *NatGateways) SetLinks(v PaginationLinks) {
	o.Links = &v
}

func (o NatGateways) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NatGateways) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableNatGateways struct {
	value *NatGateways
	isSet bool
}

func (v NullableNatGateways) Get() *NatGateways {
	return v.value
}

func (v *NullableNatGateways) Set(val *NatGateways) {
	v.value = val
	v.isSet = true
}

func (v NullableNatGateways) IsSet() bool {
	return v.isSet
}

func (v *NullableNatGateways) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNatGateways(val *NatGateways) *NullableNatGateways {
	return &NullableNatGateways{value: val, isSet: true}
}

func (v NullableNatGateways) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNatGateways) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
