/*
 * IONOS DBaaS PostgreSQL REST API
 *
 * An enterprise-grade Database is provided as a Service (DBaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.  The API allows you to create additional PostgreSQL database clusters or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psql

import (
	"encoding/json"
)

// checks if the DatabaseResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseResource{}

// DatabaseResource struct for DatabaseResource
type DatabaseResource struct {
	Type ResourceType `json:"type"`
	// The unique ID of the resource.
	Id string `json:"id"`
	// Absolute URL of the resource.
	Href       string             `json:"href"`
	Metadata   *Metadata          `json:"metadata,omitempty"`
	Properties DatabaseProperties `json:"properties"`
}

// NewDatabaseResource instantiates a new DatabaseResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseResource(type_ ResourceType, id string, href string, properties DatabaseProperties) *DatabaseResource {
	this := DatabaseResource{}

	this.Type = type_
	this.Id = id
	this.Href = href
	this.Properties = properties

	return &this
}

// NewDatabaseResourceWithDefaults instantiates a new DatabaseResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseResourceWithDefaults() *DatabaseResource {
	this := DatabaseResource{}
	return &this
}

// GetType returns the Type field value
func (o *DatabaseResource) GetType() ResourceType {
	if o == nil {
		var ret ResourceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DatabaseResource) GetTypeOk() (*ResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DatabaseResource) SetType(v ResourceType) {
	o.Type = v
}

// GetId returns the Id field value
func (o *DatabaseResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DatabaseResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DatabaseResource) SetId(v string) {
	o.Id = v
}

// GetHref returns the Href field value
func (o *DatabaseResource) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *DatabaseResource) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *DatabaseResource) SetHref(v string) {
	o.Href = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DatabaseResource) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseResource) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DatabaseResource) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *DatabaseResource) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetProperties returns the Properties field value
func (o *DatabaseResource) GetProperties() DatabaseProperties {
	if o == nil {
		var ret DatabaseProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *DatabaseResource) GetPropertiesOk() (*DatabaseProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *DatabaseResource) SetProperties(v DatabaseProperties) {
	o.Properties = v
}

func (o DatabaseResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["href"] = o.Href
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableDatabaseResource struct {
	value *DatabaseResource
	isSet bool
}

func (v NullableDatabaseResource) Get() *DatabaseResource {
	return v.value
}

func (v *NullableDatabaseResource) Set(val *DatabaseResource) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseResource) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseResource(val *DatabaseResource) *NullableDatabaseResource {
	return &NullableDatabaseResource{value: val, isSet: true}
}

func (v NullableDatabaseResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}