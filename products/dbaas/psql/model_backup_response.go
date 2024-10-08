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

// checks if the BackupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupResponse{}

// BackupResponse A database backup.
type BackupResponse struct {
	Type *ResourceType `json:"type,omitempty"`
	// The unique ID of the resource.
	Id         *string         `json:"id,omitempty"`
	Metadata   *BackupMetadata `json:"metadata,omitempty"`
	Properties *ClusterBackup  `json:"properties,omitempty"`
}

// NewBackupResponse instantiates a new BackupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupResponse() *BackupResponse {
	this := BackupResponse{}

	return &this
}

// NewBackupResponseWithDefaults instantiates a new BackupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupResponseWithDefaults() *BackupResponse {
	this := BackupResponse{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BackupResponse) GetType() ResourceType {
	if o == nil || IsNil(o.Type) {
		var ret ResourceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetTypeOk() (*ResourceType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BackupResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ResourceType and assigns it to the Type field.
func (o *BackupResponse) SetType(v ResourceType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BackupResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BackupResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BackupResponse) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *BackupResponse) GetMetadata() BackupMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret BackupMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetMetadataOk() (*BackupMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BackupResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given BackupMetadata and assigns it to the Metadata field.
func (o *BackupResponse) SetMetadata(v BackupMetadata) {
	o.Metadata = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BackupResponse) GetProperties() ClusterBackup {
	if o == nil || IsNil(o.Properties) {
		var ret ClusterBackup
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetPropertiesOk() (*ClusterBackup, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BackupResponse) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given ClusterBackup and assigns it to the Properties field.
func (o *BackupResponse) SetProperties(v ClusterBackup) {
	o.Properties = &v
}

func (o BackupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableBackupResponse struct {
	value *BackupResponse
	isSet bool
}

func (v NullableBackupResponse) Get() *BackupResponse {
	return v.value
}

func (v *NullableBackupResponse) Set(val *BackupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupResponse(val *BackupResponse) *NullableBackupResponse {
	return &NullableBackupResponse{value: val, isSet: true}
}

func (v NullableBackupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
