# CreateClusterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**PostgresVersion** | **string** | The PostgreSQL version of your cluster. | |
|**Instances** | **int32** | The total number of instances in the cluster (one master and n-1 standbys).  | |
|**Cores** | **int32** | The number of CPU cores per instance. | |
|**Ram** | **int32** | The amount of memory per instance in megabytes. Has to be a multiple of 1024. | |
|**StorageSize** | **int32** | The amount of storage per instance in megabytes. | |
|**StorageType** | [**StorageType**](StorageType.md) |  | |
|**Connections** | [**[]Connection**](Connection.md) |  | |
|**Location** | **string** | The physical location where the cluster will be created. This will be where all of your instances live. Property cannot be modified after datacenter creation.  | |
|**BackupLocation** | Pointer to **string** | The S3 location where the backups will be stored. | [optional] |
|**DisplayName** | **string** | The friendly name of your cluster. | |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |
|**Credentials** | [**DBUser**](DBUser.md) |  | |
|**SynchronizationMode** | [**SynchronizationMode**](SynchronizationMode.md) |  | |
|**FromBackup** | Pointer to [**CreateRestoreRequest**](CreateRestoreRequest.md) |  | [optional] |

## Methods

### NewCreateClusterProperties

`func NewCreateClusterProperties(postgresVersion string, instances int32, cores int32, ram int32, storageSize int32, storageType StorageType, connections []Connection, location string, displayName string, credentials DBUser, synchronizationMode SynchronizationMode, ) *CreateClusterProperties`

NewCreateClusterProperties instantiates a new CreateClusterProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterPropertiesWithDefaults

`func NewCreateClusterPropertiesWithDefaults() *CreateClusterProperties`

NewCreateClusterPropertiesWithDefaults instantiates a new CreateClusterProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostgresVersion

`func (o *CreateClusterProperties) GetPostgresVersion() string`

GetPostgresVersion returns the PostgresVersion field if non-nil, zero value otherwise.

### GetPostgresVersionOk

`func (o *CreateClusterProperties) GetPostgresVersionOk() (*string, bool)`

GetPostgresVersionOk returns a tuple with the PostgresVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresVersion

`func (o *CreateClusterProperties) SetPostgresVersion(v string)`

SetPostgresVersion sets PostgresVersion field to given value.


### GetInstances

`func (o *CreateClusterProperties) GetInstances() int32`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *CreateClusterProperties) GetInstancesOk() (*int32, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *CreateClusterProperties) SetInstances(v int32)`

SetInstances sets Instances field to given value.


### GetCores

`func (o *CreateClusterProperties) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *CreateClusterProperties) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *CreateClusterProperties) SetCores(v int32)`

SetCores sets Cores field to given value.


### GetRam

`func (o *CreateClusterProperties) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *CreateClusterProperties) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *CreateClusterProperties) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetStorageSize

`func (o *CreateClusterProperties) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *CreateClusterProperties) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *CreateClusterProperties) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.


### GetStorageType

`func (o *CreateClusterProperties) GetStorageType() StorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *CreateClusterProperties) GetStorageTypeOk() (*StorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *CreateClusterProperties) SetStorageType(v StorageType)`

SetStorageType sets StorageType field to given value.


### GetConnections

`func (o *CreateClusterProperties) GetConnections() []Connection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *CreateClusterProperties) GetConnectionsOk() (*[]Connection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *CreateClusterProperties) SetConnections(v []Connection)`

SetConnections sets Connections field to given value.


### GetLocation

`func (o *CreateClusterProperties) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateClusterProperties) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateClusterProperties) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetBackupLocation

`func (o *CreateClusterProperties) GetBackupLocation() string`

GetBackupLocation returns the BackupLocation field if non-nil, zero value otherwise.

### GetBackupLocationOk

`func (o *CreateClusterProperties) GetBackupLocationOk() (*string, bool)`

GetBackupLocationOk returns a tuple with the BackupLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocation

`func (o *CreateClusterProperties) SetBackupLocation(v string)`

SetBackupLocation sets BackupLocation field to given value.

### HasBackupLocation

`func (o *CreateClusterProperties) HasBackupLocation() bool`

HasBackupLocation returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateClusterProperties) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateClusterProperties) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateClusterProperties) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMaintenanceWindow

`func (o *CreateClusterProperties) GetMaintenanceWindow() MaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *CreateClusterProperties) GetMaintenanceWindowOk() (*MaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *CreateClusterProperties) SetMaintenanceWindow(v MaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *CreateClusterProperties) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.

### GetCredentials

`func (o *CreateClusterProperties) GetCredentials() DBUser`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CreateClusterProperties) GetCredentialsOk() (*DBUser, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CreateClusterProperties) SetCredentials(v DBUser)`

SetCredentials sets Credentials field to given value.


### GetSynchronizationMode

`func (o *CreateClusterProperties) GetSynchronizationMode() SynchronizationMode`

GetSynchronizationMode returns the SynchronizationMode field if non-nil, zero value otherwise.

### GetSynchronizationModeOk

`func (o *CreateClusterProperties) GetSynchronizationModeOk() (*SynchronizationMode, bool)`

GetSynchronizationModeOk returns a tuple with the SynchronizationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizationMode

`func (o *CreateClusterProperties) SetSynchronizationMode(v SynchronizationMode)`

SetSynchronizationMode sets SynchronizationMode field to given value.


### GetFromBackup

`func (o *CreateClusterProperties) GetFromBackup() CreateRestoreRequest`

GetFromBackup returns the FromBackup field if non-nil, zero value otherwise.

### GetFromBackupOk

`func (o *CreateClusterProperties) GetFromBackupOk() (*CreateRestoreRequest, bool)`

GetFromBackupOk returns a tuple with the FromBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromBackup

`func (o *CreateClusterProperties) SetFromBackup(v CreateRestoreRequest)`

SetFromBackup sets FromBackup field to given value.

### HasFromBackup

`func (o *CreateClusterProperties) HasFromBackup() bool`

HasFromBackup returns a boolean if a field has been set.


