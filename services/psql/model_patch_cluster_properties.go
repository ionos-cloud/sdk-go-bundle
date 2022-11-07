/*
 * IONOS DBaaS REST API
 *
 * An enterprise-grade Database is provided as a Service (DBaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.  The API allows you to create additional database clusters or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psql

import (
	"encoding/json"
)

// PatchClusterProperties Properties of the payload to change a cluster.
type PatchClusterProperties struct {
	// The number of CPU cores per instance.
	Cores *int32 `json:"cores,omitempty"`
	// The amount of memory per instance in megabytes. Has to be a multiple of 1024.
	Ram *int32 `json:"ram,omitempty"`
	// The amount of storage per instance in megabytes.
	StorageSize *int32        `json:"storageSize,omitempty"`
	Connections *[]Connection `json:"connections,omitempty"`
	// The friendly name of your cluster.
	DisplayName       *string            `json:"displayName,omitempty"`
	MaintenanceWindow *MaintenanceWindow `json:"maintenanceWindow,omitempty"`
	// The PostgreSQL version of your cluster.
	PostgresVersion *string `json:"postgresVersion,omitempty"`
	// The total number of instances in the cluster (one master and n-1 standbys).
	Instances *int32 `json:"instances,omitempty"`
}

// NewPatchClusterProperties instantiates a new PatchClusterProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchClusterProperties() *PatchClusterProperties {
	this := PatchClusterProperties{}

	return &this
}

// NewPatchClusterPropertiesWithDefaults instantiates a new PatchClusterProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchClusterPropertiesWithDefaults() *PatchClusterProperties {
	this := PatchClusterProperties{}
	return &this
}

// GetCores returns the Cores field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PatchClusterProperties) GetCores() *int32 {
	if o == nil {
		return nil
	}

	return o.Cores

}

// GetCoresOk returns a tuple with the Cores field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchClusterProperties) GetCoresOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Cores, true
}

// SetCores sets field value
func (o *PatchClusterProperties) SetCores(v int32) {

	o.Cores = &v

}

// HasCores returns a boolean if a field has been set.
func (o *PatchClusterProperties) HasCores() bool {
	if o != nil && o.Cores != nil {
		return true
	}

	return false
}

// GetRam returns the Ram field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PatchClusterProperties) GetRam() *int32 {
	if o == nil {
		return nil
	}

	return o.Ram

}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchClusterProperties) GetRamOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Ram, true
}

// SetRam sets field value
func (o *PatchClusterProperties) SetRam(v int32) {

	o.Ram = &v

}

// HasRam returns a boolean if a field has been set.
func (o *PatchClusterProperties) HasRam() bool {
	if o != nil && o.Ram != nil {
		return true
	}

	return false
}

// GetStorageSize returns the StorageSize field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PatchClusterProperties) GetStorageSize() *int32 {
	if o == nil {
		return nil
	}

	return o.StorageSize

}

// GetStorageSizeOk returns a tuple with the StorageSize field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchClusterProperties) GetStorageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.StorageSize, true
}

// SetStorageSize sets field value
func (o *PatchClusterProperties) SetStorageSize(v int32) {

	o.StorageSize = &v

}

// HasStorageSize returns a boolean if a field has been set.
func (o *PatchClusterProperties) HasStorageSize() bool {
	if o != nil && o.StorageSize != nil {
		return true
	}

	return false
}

// GetConnections returns the Connections field value
// If the value is explicit nil, the zero value for []Connection will be returned
func (o *PatchClusterProperties) GetConnections() *[]Connection {
	if o == nil {
		return nil
	}

	return o.Connections

}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchClusterProperties) GetConnectionsOk() (*[]Connection, bool) {
	if o == nil {
		return nil, false
	}

	return o.Connections, true
}

// SetConnections sets field value
func (o *PatchClusterProperties) SetConnections(v []Connection) {

	o.Connections = &v

}

// HasConnections returns a boolean if a field has been set.
func (o *PatchClusterProperties) HasConnections() bool {
	if o != nil && o.Connections != nil {
		return true
	}

	return false
}

// GetDisplayName returns the DisplayName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PatchClusterProperties) GetDisplayName() *string {
	if o == nil {
		return nil
	}

	return o.DisplayName

}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchClusterProperties) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.DisplayName, true
}

// SetDisplayName sets field value
func (o *PatchClusterProperties) SetDisplayName(v string) {

	o.DisplayName = &v

}

// HasDisplayName returns a boolean if a field has been set.
func (o *PatchClusterProperties) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// GetMaintenanceWindow returns the MaintenanceWindow field value
// If the value is explicit nil, the zero value for MaintenanceWindow will be returned
func (o *PatchClusterProperties) GetMaintenanceWindow() *MaintenanceWindow {
	if o == nil {
		return nil
	}

	return o.MaintenanceWindow

}

// GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchClusterProperties) GetMaintenanceWindowOk() (*MaintenanceWindow, bool) {
	if o == nil {
		return nil, false
	}

	return o.MaintenanceWindow, true
}

// SetMaintenanceWindow sets field value
func (o *PatchClusterProperties) SetMaintenanceWindow(v MaintenanceWindow) {

	o.MaintenanceWindow = &v

}

// HasMaintenanceWindow returns a boolean if a field has been set.
func (o *PatchClusterProperties) HasMaintenanceWindow() bool {
	if o != nil && o.MaintenanceWindow != nil {
		return true
	}

	return false
}

// GetPostgresVersion returns the PostgresVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PatchClusterProperties) GetPostgresVersion() *string {
	if o == nil {
		return nil
	}

	return o.PostgresVersion

}

// GetPostgresVersionOk returns a tuple with the PostgresVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchClusterProperties) GetPostgresVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.PostgresVersion, true
}

// SetPostgresVersion sets field value
func (o *PatchClusterProperties) SetPostgresVersion(v string) {

	o.PostgresVersion = &v

}

// HasPostgresVersion returns a boolean if a field has been set.
func (o *PatchClusterProperties) HasPostgresVersion() bool {
	if o != nil && o.PostgresVersion != nil {
		return true
	}

	return false
}

// GetInstances returns the Instances field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PatchClusterProperties) GetInstances() *int32 {
	if o == nil {
		return nil
	}

	return o.Instances

}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchClusterProperties) GetInstancesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Instances, true
}

// SetInstances sets field value
func (o *PatchClusterProperties) SetInstances(v int32) {

	o.Instances = &v

}

// HasInstances returns a boolean if a field has been set.
func (o *PatchClusterProperties) HasInstances() bool {
	if o != nil && o.Instances != nil {
		return true
	}

	return false
}

func (o PatchClusterProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cores != nil {
		toSerialize["cores"] = o.Cores
	}

	if o.Ram != nil {
		toSerialize["ram"] = o.Ram
	}

	if o.StorageSize != nil {
		toSerialize["storageSize"] = o.StorageSize
	}

	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}

	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}

	if o.MaintenanceWindow != nil {
		toSerialize["maintenanceWindow"] = o.MaintenanceWindow
	}

	if o.PostgresVersion != nil {
		toSerialize["postgresVersion"] = o.PostgresVersion
	}

	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
	}

	return json.Marshal(toSerialize)
}

type NullablePatchClusterProperties struct {
	value *PatchClusterProperties
	isSet bool
}

func (v NullablePatchClusterProperties) Get() *PatchClusterProperties {
	return v.value
}

func (v *NullablePatchClusterProperties) Set(val *PatchClusterProperties) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchClusterProperties) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchClusterProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchClusterProperties(val *PatchClusterProperties) *NullablePatchClusterProperties {
	return &NullablePatchClusterProperties{value: val, isSet: true}
}

func (v NullablePatchClusterProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchClusterProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
