/*
 * Certificate Manager Service API
 *
 * Using the Certificate Manager Service, you can conveniently provision and manage SSL certificates  with IONOS services and your internal connected resources.   For the [Application Load Balancer](https://api.ionos.com/docs/cloud/v6/#Application-Load-Balancers-get-datacenters-datacenterId-applicationloadbalancers), you usually need a certificate to encrypt your HTTPS traffic. The service provides the basic functions of uploading and deleting your certificates for this purpose.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cert

import (
	"encoding/json"
)

// checks if the ProviderRead type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderRead{}

// ProviderRead struct for ProviderRead
type ProviderRead struct {
	// The ID (UUID) of the Provider.
	Id string `json:"id"`
	// The type of the resource.
	Type string `json:"type"`
	// The URL of the Provider.
	Href       string             `json:"href"`
	Metadata   MetadataWithStatus `json:"metadata"`
	Properties Provider           `json:"properties"`
}

// NewProviderRead instantiates a new ProviderRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderRead(id string, type_ string, href string, metadata MetadataWithStatus, properties Provider) *ProviderRead {
	this := ProviderRead{}

	this.Id = id
	this.Type = type_
	this.Href = href
	this.Metadata = metadata
	this.Properties = properties

	return &this
}

// NewProviderReadWithDefaults instantiates a new ProviderRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderReadWithDefaults() *ProviderRead {
	this := ProviderRead{}
	return &this
}

// GetId returns the Id field value
func (o *ProviderRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProviderRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProviderRead) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ProviderRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProviderRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProviderRead) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *ProviderRead) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ProviderRead) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ProviderRead) SetHref(v string) {
	o.Href = v
}

// GetMetadata returns the Metadata field value
func (o *ProviderRead) GetMetadata() MetadataWithStatus {
	if o == nil {
		var ret MetadataWithStatus
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ProviderRead) GetMetadataOk() (*MetadataWithStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ProviderRead) SetMetadata(v MetadataWithStatus) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *ProviderRead) GetProperties() Provider {
	if o == nil {
		var ret Provider
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ProviderRead) GetPropertiesOk() (*Provider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *ProviderRead) SetProperties(v Provider) {
	o.Properties = v
}

func (o ProviderRead) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["href"] = o.Href
	toSerialize["metadata"] = o.Metadata
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableProviderRead struct {
	value *ProviderRead
	isSet bool
}

func (v NullableProviderRead) Get() *ProviderRead {
	return v.value
}

func (v *NullableProviderRead) Set(val *ProviderRead) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderRead) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderRead(val *ProviderRead) *NullableProviderRead {
	return &NullableProviderRead{value: val, isSet: true}
}

func (v NullableProviderRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}