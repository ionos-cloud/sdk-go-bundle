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

// checks if the TemplateProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateProperties{}

// TemplateProperties struct for TemplateProperties
type TemplateProperties struct {
	// The resource name.
	Name string `json:"name"`
	// The CPU cores count.
	Cores float32 `json:"cores"`
	// The RAM size in MB.
	Ram float32 `json:"ram"`
	// The storage size in GB.
	StorageSize float32 `json:"storageSize"`
}

// NewTemplateProperties instantiates a new TemplateProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateProperties(name string, cores float32, ram float32, storageSize float32) *TemplateProperties {
	this := TemplateProperties{}

	this.Name = name
	this.Cores = cores
	this.Ram = ram
	this.StorageSize = storageSize

	return &this
}

// NewTemplatePropertiesWithDefaults instantiates a new TemplateProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplatePropertiesWithDefaults() *TemplateProperties {
	this := TemplateProperties{}
	return &this
}

// GetName returns the Name field value
func (o *TemplateProperties) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateProperties) SetName(v string) {
	o.Name = v
}

// GetCores returns the Cores field value
func (o *TemplateProperties) GetCores() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Cores
}

// GetCoresOk returns a tuple with the Cores field value
// and a boolean to check if the value has been set.
func (o *TemplateProperties) GetCoresOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cores, true
}

// SetCores sets field value
func (o *TemplateProperties) SetCores(v float32) {
	o.Cores = v
}

// GetRam returns the Ram field value
func (o *TemplateProperties) GetRam() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Ram
}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
func (o *TemplateProperties) GetRamOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ram, true
}

// SetRam sets field value
func (o *TemplateProperties) SetRam(v float32) {
	o.Ram = v
}

// GetStorageSize returns the StorageSize field value
func (o *TemplateProperties) GetStorageSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StorageSize
}

// GetStorageSizeOk returns a tuple with the StorageSize field value
// and a boolean to check if the value has been set.
func (o *TemplateProperties) GetStorageSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageSize, true
}

// SetStorageSize sets field value
func (o *TemplateProperties) SetStorageSize(v float32) {
	o.StorageSize = v
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
	if !IsZero(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsZero(o.Cores) {
		toSerialize["cores"] = o.Cores
	}
	if !IsZero(o.Ram) {
		toSerialize["ram"] = o.Ram
	}
	if !IsZero(o.StorageSize) {
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
