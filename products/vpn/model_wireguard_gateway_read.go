/*
 * IONOS Cloud VPN Gateway API
 *
 * The Managed VPN Gateway service provides secure and scalable connectivity, enabling encrypted communication between your IONOS cloud resources in a VDC and remote networks (on-premises, multi-cloud, private LANs in other VDCs etc).
 *
 * API version: 1.0.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vpn

import (
	"encoding/json"
)

// checks if the WireguardGatewayRead type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WireguardGatewayRead{}

// WireguardGatewayRead struct for WireguardGatewayRead
type WireguardGatewayRead struct {
	// The ID (UUID) of the WireguardGateway.
	Id string `json:"id"`
	// The type of the resource.
	Type string `json:"type"`
	// The URL of the WireguardGateway.
	Href       string                   `json:"href"`
	Metadata   WireguardGatewayMetadata `json:"metadata"`
	Properties WireguardGateway         `json:"properties"`
}

// NewWireguardGatewayRead instantiates a new WireguardGatewayRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireguardGatewayRead(id string, type_ string, href string, metadata WireguardGatewayMetadata, properties WireguardGateway) *WireguardGatewayRead {
	this := WireguardGatewayRead{}

	this.Id = id
	this.Type = type_
	this.Href = href
	this.Metadata = metadata
	this.Properties = properties

	return &this
}

// NewWireguardGatewayReadWithDefaults instantiates a new WireguardGatewayRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireguardGatewayReadWithDefaults() *WireguardGatewayRead {
	this := WireguardGatewayRead{}
	return &this
}

// GetId returns the Id field value
func (o *WireguardGatewayRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WireguardGatewayRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WireguardGatewayRead) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *WireguardGatewayRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WireguardGatewayRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WireguardGatewayRead) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *WireguardGatewayRead) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *WireguardGatewayRead) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *WireguardGatewayRead) SetHref(v string) {
	o.Href = v
}

// GetMetadata returns the Metadata field value
func (o *WireguardGatewayRead) GetMetadata() WireguardGatewayMetadata {
	if o == nil {
		var ret WireguardGatewayMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *WireguardGatewayRead) GetMetadataOk() (*WireguardGatewayMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *WireguardGatewayRead) SetMetadata(v WireguardGatewayMetadata) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *WireguardGatewayRead) GetProperties() WireguardGateway {
	if o == nil {
		var ret WireguardGateway
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *WireguardGatewayRead) GetPropertiesOk() (*WireguardGateway, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *WireguardGatewayRead) SetProperties(v WireguardGateway) {
	o.Properties = v
}

func (o WireguardGatewayRead) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["href"] = o.Href
	toSerialize["metadata"] = o.Metadata
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableWireguardGatewayRead struct {
	value *WireguardGatewayRead
	isSet bool
}

func (v NullableWireguardGatewayRead) Get() *WireguardGatewayRead {
	return v.value
}

func (v *NullableWireguardGatewayRead) Set(val *WireguardGatewayRead) {
	v.value = val
	v.isSet = true
}

func (v NullableWireguardGatewayRead) IsSet() bool {
	return v.isSet
}

func (v *NullableWireguardGatewayRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireguardGatewayRead(val *WireguardGatewayRead) *NullableWireguardGatewayRead {
	return &NullableWireguardGatewayRead{value: val, isSet: true}
}

func (v NullableWireguardGatewayRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireguardGatewayRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
