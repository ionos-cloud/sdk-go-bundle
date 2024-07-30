/*
 * IONOS DBaaS MongoDB REST API
 *
 * With IONOS Cloud Database as a Service, you have the ability to quickly set up and manage a MongoDB database. You can also delete clusters, manage backups and users via the API.  MongoDB is an open source, cross-platform, document-oriented database program. Classified as a NoSQL database program, it uses JSON-like documents with optional schemas.  The MongoDB API allows you to create additional database clusters or modify existing ones. Both tools, the Data Center Designer (DCD) and the API use the same concepts consistently and are well suited for smooth and intuitive use.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongo

import (
	"encoding/json"
)

// checks if the TemplateProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateProperties{}

// TemplateProperties The properties of a MongoDB template.
type TemplateProperties struct {
	// The name of the template.
	Name *string `json:"name,omitempty"`
	// The edition of the template (e.g. enterprise)
	Edition *string `json:"edition,omitempty"`
	// The number of CPU cores.
	Cores *int32 `json:"cores,omitempty"`
	// The amount of memory in MB.
	Ram *int32 `json:"ram,omitempty"`
	// The amount of storage size in GB.
	StorageSize *int32 `json:"storageSize,omitempty"`
}

// NewTemplateProperties instantiates a new TemplateProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateProperties() *TemplateProperties {
	this := TemplateProperties{}

	return &this
}

// NewTemplatePropertiesWithDefaults instantiates a new TemplateProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplatePropertiesWithDefaults() *TemplateProperties {
	this := TemplateProperties{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateProperties) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateProperties) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateProperties) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateProperties) SetName(v string) {
	o.Name = &v
}

// GetEdition returns the Edition field value if set, zero value otherwise.
func (o *TemplateProperties) GetEdition() string {
	if o == nil || IsNil(o.Edition) {
		var ret string
		return ret
	}
	return *o.Edition
}

// GetEditionOk returns a tuple with the Edition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateProperties) GetEditionOk() (*string, bool) {
	if o == nil || IsNil(o.Edition) {
		return nil, false
	}
	return o.Edition, true
}

// HasEdition returns a boolean if a field has been set.
func (o *TemplateProperties) HasEdition() bool {
	if o != nil && !IsNil(o.Edition) {
		return true
	}

	return false
}

// SetEdition gets a reference to the given string and assigns it to the Edition field.
func (o *TemplateProperties) SetEdition(v string) {
	o.Edition = &v
}

// GetCores returns the Cores field value if set, zero value otherwise.
func (o *TemplateProperties) GetCores() int32 {
	if o == nil || IsNil(o.Cores) {
		var ret int32
		return ret
	}
	return *o.Cores
}

// GetCoresOk returns a tuple with the Cores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateProperties) GetCoresOk() (*int32, bool) {
	if o == nil || IsNil(o.Cores) {
		return nil, false
	}
	return o.Cores, true
}

// HasCores returns a boolean if a field has been set.
func (o *TemplateProperties) HasCores() bool {
	if o != nil && !IsNil(o.Cores) {
		return true
	}

	return false
}

// SetCores gets a reference to the given int32 and assigns it to the Cores field.
func (o *TemplateProperties) SetCores(v int32) {
	o.Cores = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *TemplateProperties) GetRam() int32 {
	if o == nil || IsNil(o.Ram) {
		var ret int32
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateProperties) GetRamOk() (*int32, bool) {
	if o == nil || IsNil(o.Ram) {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *TemplateProperties) HasRam() bool {
	if o != nil && !IsNil(o.Ram) {
		return true
	}

	return false
}

// SetRam gets a reference to the given int32 and assigns it to the Ram field.
func (o *TemplateProperties) SetRam(v int32) {
	o.Ram = &v
}

// GetStorageSize returns the StorageSize field value if set, zero value otherwise.
func (o *TemplateProperties) GetStorageSize() int32 {
	if o == nil || IsNil(o.StorageSize) {
		var ret int32
		return ret
	}
	return *o.StorageSize
}

// GetStorageSizeOk returns a tuple with the StorageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateProperties) GetStorageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.StorageSize) {
		return nil, false
	}
	return o.StorageSize, true
}

// HasStorageSize returns a boolean if a field has been set.
func (o *TemplateProperties) HasStorageSize() bool {
	if o != nil && !IsNil(o.StorageSize) {
		return true
	}

	return false
}

// SetStorageSize gets a reference to the given int32 and assigns it to the StorageSize field.
func (o *TemplateProperties) SetStorageSize(v int32) {
	o.StorageSize = &v
}

func (o TemplateProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Edition) {
		toSerialize["edition"] = o.Edition
	}
	if !IsNil(o.Cores) {
		toSerialize["cores"] = o.Cores
	}
	if !IsNil(o.Ram) {
		toSerialize["ram"] = o.Ram
	}
	if !IsNil(o.StorageSize) {
		toSerialize["storageSize"] = o.StorageSize
	}
	return toSerialize, nil
}

type NullableTemplateProperties struct {
	value *TemplateProperties
	isSet bool
}

func (v NullableTemplateProperties) Get() *TemplateProperties {
	return v.value
}

func (v *NullableTemplateProperties) Set(val *TemplateProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateProperties(val *TemplateProperties) *NullableTemplateProperties {
	return &NullableTemplateProperties{value: val, isSet: true}
}

func (v NullableTemplateProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
