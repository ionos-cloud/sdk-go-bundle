# ClusterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DisplayName** | Pointer to **string** | The friendly name of your cluster. | [optional] |
|**PostgresVersion** | Pointer to **string** | The PostgreSQL version of your cluster. | [optional] |
|**Location** | Pointer to **string** | The physical location where the cluster will be created. This will be where all of your instances live. Property cannot be modified after datacenter creation.  | [optional] |
|**BackupLocation** | Pointer to **string** | The S3 location where the backups will be stored. | [optional] |
|**Instances** | Pointer to **int32** | The total number of instances in the cluster (one master and n-1 standbys).  | [optional] |
|**Ram** | Pointer to **int32** | The amount of memory per instance in megabytes. Has to be a multiple of 1024. | [optional] |
|**Cores** | Pointer to **int32** | The number of CPU cores per instance. | [optional] |
|**StorageSize** | Pointer to **int32** | The amount of storage per instance in megabytes. | [optional] |
|**StorageType** | Pointer to [**StorageType**](StorageType.md) |  | [optional] |
|**Connections** | Pointer to [**[]Connection**](Connection.md) |  | [optional] |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |
|**SynchronizationMode** | Pointer to [**SynchronizationMode**](SynchronizationMode.md) |  | [optional] |

## Methods

### NewClusterProperties

`func NewClusterProperties() *ClusterProperties`

NewClusterProperties instantiates a new ClusterProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPropertiesWithDefaults

`func NewClusterPropertiesWithDefaults() *ClusterProperties`

NewClusterPropertiesWithDefaults instantiates a new ClusterProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ClusterProperties) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ClusterProperties) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ClusterProperties) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ClusterProperties) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPostgresVersion

`func (o *ClusterProperties) GetPostgresVersion() string`

GetPostgresVersion returns the PostgresVersion field if non-nil, zero value otherwise.

### GetPostgresVersionOk

`func (o *ClusterProperties) GetPostgresVersionOk() (*string, bool)`

GetPostgresVersionOk returns a tuple with the PostgresVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresVersion

`func (o *ClusterProperties) SetPostgresVersion(v string)`

SetPostgresVersion sets PostgresVersion field to given value.

### HasPostgresVersion

`func (o *ClusterProperties) HasPostgresVersion() bool`

HasPostgresVersion returns a boolean if a field has been set.

### GetLocation

`func (o *ClusterProperties) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ClusterProperties) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ClusterProperties) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ClusterProperties) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetBackupLocation

`func (o *ClusterProperties) GetBackupLocation() string`

GetBackupLocation returns the BackupLocation field if non-nil, zero value otherwise.

### GetBackupLocationOk

`func (o *ClusterProperties) GetBackupLocationOk() (*string, bool)`

GetBackupLocationOk returns a tuple with the BackupLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocation

`func (o *ClusterProperties) SetBackupLocation(v string)`

SetBackupLocation sets BackupLocation field to given value.

### HasBackupLocation

`func (o *ClusterProperties) HasBackupLocation() bool`

HasBackupLocation returns a boolean if a field has been set.

### GetInstances

`func (o *ClusterProperties) GetInstances() int32`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ClusterProperties) GetInstancesOk() (*int32, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ClusterProperties) SetInstances(v int32)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ClusterProperties) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetRam

`func (o *ClusterProperties) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ClusterProperties) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ClusterProperties) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *ClusterProperties) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetCores

`func (o *ClusterProperties) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *ClusterProperties) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *ClusterProperties) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *ClusterProperties) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetStorageSize

`func (o *ClusterProperties) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *ClusterProperties) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *ClusterProperties) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *ClusterProperties) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetStorageType

`func (o *ClusterProperties) GetStorageType() StorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *ClusterProperties) GetStorageTypeOk() (*StorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *ClusterProperties) SetStorageType(v StorageType)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *ClusterProperties) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetConnections

`func (o *ClusterProperties) GetConnections() []Connection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ClusterProperties) GetConnectionsOk() (*[]Connection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ClusterProperties) SetConnections(v []Connection)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *ClusterProperties) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetMaintenanceWindow

`func (o *ClusterProperties) GetMaintenanceWindow() MaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *ClusterProperties) GetMaintenanceWindowOk() (*MaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *ClusterProperties) SetMaintenanceWindow(v MaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *ClusterProperties) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.

### GetSynchronizationMode

`func (o *ClusterProperties) GetSynchronizationMode() SynchronizationMode`

GetSynchronizationMode returns the SynchronizationMode field if non-nil, zero value otherwise.

### GetSynchronizationModeOk

`func (o *ClusterProperties) GetSynchronizationModeOk() (*SynchronizationMode, bool)`

GetSynchronizationModeOk returns a tuple with the SynchronizationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizationMode

`func (o *ClusterProperties) SetSynchronizationMode(v SynchronizationMode)`

SetSynchronizationMode sets SynchronizationMode field to given value.

### HasSynchronizationMode

`func (o *ClusterProperties) HasSynchronizationMode() bool`

HasSynchronizationMode returns a boolean if a field has been set.


