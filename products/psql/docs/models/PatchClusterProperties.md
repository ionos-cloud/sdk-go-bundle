# PatchClusterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Cores** | Pointer to **int32** | The number of CPU cores per instance. | [optional] |
|**Ram** | Pointer to **int32** | The amount of memory per instance in megabytes. Has to be a multiple of 1024. | [optional] |
|**StorageSize** | Pointer to **int32** | The amount of storage per instance in megabytes. | [optional] |
|**Connections** | [**[]Connection**](Connection.md) |  | |
|**DisplayName** | Pointer to **string** | The friendly name of your cluster. | [optional] |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |
|**PostgresVersion** | Pointer to **string** | The PostgreSQL version of your cluster. | [optional] |
|**Instances** | Pointer to **int32** | The total number of instances in the cluster (one master and n-1 standbys).  | [optional] |
|**ConnectionPooler** | Pointer to [**ConnectionPooler**](ConnectionPooler.md) |  | [optional] |

## Methods

### NewPatchClusterProperties

`func NewPatchClusterProperties(connections []Connection, ) *PatchClusterProperties`

NewPatchClusterProperties instantiates a new PatchClusterProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchClusterPropertiesWithDefaults

`func NewPatchClusterPropertiesWithDefaults() *PatchClusterProperties`

NewPatchClusterPropertiesWithDefaults instantiates a new PatchClusterProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCores

`func (o *PatchClusterProperties) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *PatchClusterProperties) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *PatchClusterProperties) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *PatchClusterProperties) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetRam

`func (o *PatchClusterProperties) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *PatchClusterProperties) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *PatchClusterProperties) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *PatchClusterProperties) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetStorageSize

`func (o *PatchClusterProperties) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *PatchClusterProperties) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *PatchClusterProperties) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *PatchClusterProperties) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetConnections

`func (o *PatchClusterProperties) GetConnections() []Connection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *PatchClusterProperties) GetConnectionsOk() (*[]Connection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *PatchClusterProperties) SetConnections(v []Connection)`

SetConnections sets Connections field to given value.


### GetDisplayName

`func (o *PatchClusterProperties) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PatchClusterProperties) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PatchClusterProperties) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PatchClusterProperties) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMaintenanceWindow

`func (o *PatchClusterProperties) GetMaintenanceWindow() MaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *PatchClusterProperties) GetMaintenanceWindowOk() (*MaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *PatchClusterProperties) SetMaintenanceWindow(v MaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *PatchClusterProperties) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.

### GetPostgresVersion

`func (o *PatchClusterProperties) GetPostgresVersion() string`

GetPostgresVersion returns the PostgresVersion field if non-nil, zero value otherwise.

### GetPostgresVersionOk

`func (o *PatchClusterProperties) GetPostgresVersionOk() (*string, bool)`

GetPostgresVersionOk returns a tuple with the PostgresVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresVersion

`func (o *PatchClusterProperties) SetPostgresVersion(v string)`

SetPostgresVersion sets PostgresVersion field to given value.

### HasPostgresVersion

`func (o *PatchClusterProperties) HasPostgresVersion() bool`

HasPostgresVersion returns a boolean if a field has been set.

### GetInstances

`func (o *PatchClusterProperties) GetInstances() int32`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *PatchClusterProperties) GetInstancesOk() (*int32, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *PatchClusterProperties) SetInstances(v int32)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *PatchClusterProperties) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetConnectionPooler

`func (o *PatchClusterProperties) GetConnectionPooler() ConnectionPooler`

GetConnectionPooler returns the ConnectionPooler field if non-nil, zero value otherwise.

### GetConnectionPoolerOk

`func (o *PatchClusterProperties) GetConnectionPoolerOk() (*ConnectionPooler, bool)`

GetConnectionPoolerOk returns a tuple with the ConnectionPooler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPooler

`func (o *PatchClusterProperties) SetConnectionPooler(v ConnectionPooler)`

SetConnectionPooler sets ConnectionPooler field to given value.

### HasConnectionPooler

`func (o *PatchClusterProperties) HasConnectionPooler() bool`

HasConnectionPooler returns a boolean if a field has been set.


