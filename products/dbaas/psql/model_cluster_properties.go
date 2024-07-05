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

// checks if the ClusterProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterProperties{}

// ClusterProperties Properties of a database cluster.
type ClusterProperties struct {
	// The friendly name of your cluster.
	DisplayName *string `json:"displayName,omitempty"`
	// The PostgreSQL version of your cluster.
	PostgresVersion *string `json:"postgresVersion,omitempty"`
	// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot be modified after datacenter creation.
	Location *string `json:"location,omitempty"`
	// The DNS name pointing to your cluster.
	DnsName *string `json:"dnsName,omitempty"`
	// The S3 location where the backups will be stored.
	BackupLocation *string `json:"backupLocation,omitempty"`
	// The total number of instances in the cluster (one master and n-1 standbys).
	Instances *int32 `json:"instances,omitempty"`
	// The amount of memory per instance in megabytes. Has to be a multiple of 1024.
	Ram *int32 `json:"ram,omitempty"`
	// The number of CPU cores per instance.
	Cores *int32 `json:"cores,omitempty"`
	// The amount of storage per instance in megabytes.
	StorageSize         *int32               `json:"storageSize,omitempty"`
	StorageType         *StorageType         `json:"storageType,omitempty"`
	Connections         []Connection         `json:"connections,omitempty"`
	MaintenanceWindow   *MaintenanceWindow   `json:"maintenanceWindow,omitempty"`
	SynchronizationMode *SynchronizationMode `json:"synchronizationMode,omitempty"`
	ConnectionPooler    *ConnectionPooler    `json:"connectionPooler,omitempty"`
}

// NewClusterProperties instantiates a new ClusterProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterProperties() *ClusterProperties {
	this := ClusterProperties{}

	return &this
}

// NewClusterPropertiesWithDefaults instantiates a new ClusterProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterPropertiesWithDefaults() *ClusterProperties {
	this := ClusterProperties{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ClusterProperties) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ClusterProperties) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ClusterProperties) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetPostgresVersion returns the PostgresVersion field value if set, zero value otherwise.
func (o *ClusterProperties) GetPostgresVersion() string {
	if o == nil || IsNil(o.PostgresVersion) {
		var ret string
		return ret
	}
	return *o.PostgresVersion
}

// GetPostgresVersionOk returns a tuple with the PostgresVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetPostgresVersionOk() (*string, bool) {
	if o == nil || IsNil(o.PostgresVersion) {
		return nil, false
	}
	return o.PostgresVersion, true
}

// HasPostgresVersion returns a boolean if a field has been set.
func (o *ClusterProperties) HasPostgresVersion() bool {
	if o != nil && !IsNil(o.PostgresVersion) {
		return true
	}

	return false
}

// SetPostgresVersion gets a reference to the given string and assigns it to the PostgresVersion field.
func (o *ClusterProperties) SetPostgresVersion(v string) {
	o.PostgresVersion = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ClusterProperties) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ClusterProperties) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ClusterProperties) SetLocation(v string) {
	o.Location = &v
}

// GetDnsName returns the DnsName field value if set, zero value otherwise.
func (o *ClusterProperties) GetDnsName() string {
	if o == nil || IsNil(o.DnsName) {
		var ret string
		return ret
	}
	return *o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetDnsNameOk() (*string, bool) {
	if o == nil || IsNil(o.DnsName) {
		return nil, false
	}
	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *ClusterProperties) HasDnsName() bool {
	if o != nil && !IsNil(o.DnsName) {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given string and assigns it to the DnsName field.
func (o *ClusterProperties) SetDnsName(v string) {
	o.DnsName = &v
}

// GetBackupLocation returns the BackupLocation field value if set, zero value otherwise.
func (o *ClusterProperties) GetBackupLocation() string {
	if o == nil || IsNil(o.BackupLocation) {
		var ret string
		return ret
	}
	return *o.BackupLocation
}

// GetBackupLocationOk returns a tuple with the BackupLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetBackupLocationOk() (*string, bool) {
	if o == nil || IsNil(o.BackupLocation) {
		return nil, false
	}
	return o.BackupLocation, true
}

// HasBackupLocation returns a boolean if a field has been set.
func (o *ClusterProperties) HasBackupLocation() bool {
	if o != nil && !IsNil(o.BackupLocation) {
		return true
	}

	return false
}

// SetBackupLocation gets a reference to the given string and assigns it to the BackupLocation field.
func (o *ClusterProperties) SetBackupLocation(v string) {
	o.BackupLocation = &v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *ClusterProperties) GetInstances() int32 {
	if o == nil || IsNil(o.Instances) {
		var ret int32
		return ret
	}
	return *o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetInstancesOk() (*int32, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *ClusterProperties) HasInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given int32 and assigns it to the Instances field.
func (o *ClusterProperties) SetInstances(v int32) {
	o.Instances = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *ClusterProperties) GetRam() int32 {
	if o == nil || IsNil(o.Ram) {
		var ret int32
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetRamOk() (*int32, bool) {
	if o == nil || IsNil(o.Ram) {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *ClusterProperties) HasRam() bool {
	if o != nil && !IsNil(o.Ram) {
		return true
	}

	return false
}

// SetRam gets a reference to the given int32 and assigns it to the Ram field.
func (o *ClusterProperties) SetRam(v int32) {
	o.Ram = &v
}

// GetCores returns the Cores field value if set, zero value otherwise.
func (o *ClusterProperties) GetCores() int32 {
	if o == nil || IsNil(o.Cores) {
		var ret int32
		return ret
	}
	return *o.Cores
}

// GetCoresOk returns a tuple with the Cores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetCoresOk() (*int32, bool) {
	if o == nil || IsNil(o.Cores) {
		return nil, false
	}
	return o.Cores, true
}

// HasCores returns a boolean if a field has been set.
func (o *ClusterProperties) HasCores() bool {
	if o != nil && !IsNil(o.Cores) {
		return true
	}

	return false
}

// SetCores gets a reference to the given int32 and assigns it to the Cores field.
func (o *ClusterProperties) SetCores(v int32) {
	o.Cores = &v
}

// GetStorageSize returns the StorageSize field value if set, zero value otherwise.
func (o *ClusterProperties) GetStorageSize() int32 {
	if o == nil || IsNil(o.StorageSize) {
		var ret int32
		return ret
	}
	return *o.StorageSize
}

// GetStorageSizeOk returns a tuple with the StorageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetStorageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.StorageSize) {
		return nil, false
	}
	return o.StorageSize, true
}

// HasStorageSize returns a boolean if a field has been set.
func (o *ClusterProperties) HasStorageSize() bool {
	if o != nil && !IsNil(o.StorageSize) {
		return true
	}

	return false
}

// SetStorageSize gets a reference to the given int32 and assigns it to the StorageSize field.
func (o *ClusterProperties) SetStorageSize(v int32) {
	o.StorageSize = &v
}

// GetStorageType returns the StorageType field value if set, zero value otherwise.
func (o *ClusterProperties) GetStorageType() StorageType {
	if o == nil || IsNil(o.StorageType) {
		var ret StorageType
		return ret
	}
	return *o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetStorageTypeOk() (*StorageType, bool) {
	if o == nil || IsNil(o.StorageType) {
		return nil, false
	}
	return o.StorageType, true
}

// HasStorageType returns a boolean if a field has been set.
func (o *ClusterProperties) HasStorageType() bool {
	if o != nil && !IsNil(o.StorageType) {
		return true
	}

	return false
}

// SetStorageType gets a reference to the given StorageType and assigns it to the StorageType field.
func (o *ClusterProperties) SetStorageType(v StorageType) {
	o.StorageType = &v
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *ClusterProperties) GetConnections() []Connection {
	if o == nil || IsNil(o.Connections) {
		var ret []Connection
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetConnectionsOk() ([]Connection, bool) {
	if o == nil || IsNil(o.Connections) {
		return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *ClusterProperties) HasConnections() bool {
	if o != nil && !IsNil(o.Connections) {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []Connection and assigns it to the Connections field.
func (o *ClusterProperties) SetConnections(v []Connection) {
	o.Connections = v
}

// GetMaintenanceWindow returns the MaintenanceWindow field value if set, zero value otherwise.
func (o *ClusterProperties) GetMaintenanceWindow() MaintenanceWindow {
	if o == nil || IsNil(o.MaintenanceWindow) {
		var ret MaintenanceWindow
		return ret
	}
	return *o.MaintenanceWindow
}

// GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetMaintenanceWindowOk() (*MaintenanceWindow, bool) {
	if o == nil || IsNil(o.MaintenanceWindow) {
		return nil, false
	}
	return o.MaintenanceWindow, true
}

// HasMaintenanceWindow returns a boolean if a field has been set.
func (o *ClusterProperties) HasMaintenanceWindow() bool {
	if o != nil && !IsNil(o.MaintenanceWindow) {
		return true
	}

	return false
}

// SetMaintenanceWindow gets a reference to the given MaintenanceWindow and assigns it to the MaintenanceWindow field.
func (o *ClusterProperties) SetMaintenanceWindow(v MaintenanceWindow) {
	o.MaintenanceWindow = &v
}

// GetSynchronizationMode returns the SynchronizationMode field value if set, zero value otherwise.
func (o *ClusterProperties) GetSynchronizationMode() SynchronizationMode {
	if o == nil || IsNil(o.SynchronizationMode) {
		var ret SynchronizationMode
		return ret
	}
	return *o.SynchronizationMode
}

// GetSynchronizationModeOk returns a tuple with the SynchronizationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetSynchronizationModeOk() (*SynchronizationMode, bool) {
	if o == nil || IsNil(o.SynchronizationMode) {
		return nil, false
	}
	return o.SynchronizationMode, true
}

// HasSynchronizationMode returns a boolean if a field has been set.
func (o *ClusterProperties) HasSynchronizationMode() bool {
	if o != nil && !IsNil(o.SynchronizationMode) {
		return true
	}

	return false
}

// SetSynchronizationMode gets a reference to the given SynchronizationMode and assigns it to the SynchronizationMode field.
func (o *ClusterProperties) SetSynchronizationMode(v SynchronizationMode) {
	o.SynchronizationMode = &v
}

// GetConnectionPooler returns the ConnectionPooler field value if set, zero value otherwise.
func (o *ClusterProperties) GetConnectionPooler() ConnectionPooler {
	if o == nil || IsNil(o.ConnectionPooler) {
		var ret ConnectionPooler
		return ret
	}
	return *o.ConnectionPooler
}

// GetConnectionPoolerOk returns a tuple with the ConnectionPooler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProperties) GetConnectionPoolerOk() (*ConnectionPooler, bool) {
	if o == nil || IsNil(o.ConnectionPooler) {
		return nil, false
	}
	return o.ConnectionPooler, true
}

// HasConnectionPooler returns a boolean if a field has been set.
func (o *ClusterProperties) HasConnectionPooler() bool {
	if o != nil && !IsNil(o.ConnectionPooler) {
		return true
	}

	return false
}

// SetConnectionPooler gets a reference to the given ConnectionPooler and assigns it to the ConnectionPooler field.
func (o *ClusterProperties) SetConnectionPooler(v ConnectionPooler) {
	o.ConnectionPooler = &v
}

func (o ClusterProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.PostgresVersion) {
		toSerialize["postgresVersion"] = o.PostgresVersion
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.DnsName) {
		toSerialize["dnsName"] = o.DnsName
	}
	if !IsNil(o.BackupLocation) {
		toSerialize["backupLocation"] = o.BackupLocation
	}
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	if !IsNil(o.Ram) {
		toSerialize["ram"] = o.Ram
	}
	if !IsNil(o.Cores) {
		toSerialize["cores"] = o.Cores
	}
	if !IsNil(o.StorageSize) {
		toSerialize["storageSize"] = o.StorageSize
	}
	if !IsNil(o.StorageType) {
		toSerialize["storageType"] = o.StorageType
	}
	if !IsNil(o.Connections) {
		toSerialize["connections"] = o.Connections
	}
	if !IsNil(o.MaintenanceWindow) {
		toSerialize["maintenanceWindow"] = o.MaintenanceWindow
	}
	if !IsNil(o.SynchronizationMode) {
		toSerialize["synchronizationMode"] = o.SynchronizationMode
	}
	if !IsNil(o.ConnectionPooler) {
		toSerialize["connectionPooler"] = o.ConnectionPooler
	}
	return toSerialize, nil
}

type NullableClusterProperties struct {
	value *ClusterProperties
	isSet bool
}

func (v NullableClusterProperties) Get() *ClusterProperties {
	return v.value
}

func (v *NullableClusterProperties) Set(val *ClusterProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterProperties(val *ClusterProperties) *NullableClusterProperties {
	return &NullableClusterProperties{value: val, isSet: true}
}

func (v NullableClusterProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
