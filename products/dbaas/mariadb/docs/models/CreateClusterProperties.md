# CreateClusterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**MariadbVersion** | [**MariadbVersion**](MariadbVersion.md) |  | |
|**Instances** | **int32** | The total number of instances in the cluster (one primary and n-1 secondary).  | |
|**Cores** | **int32** | The number of CPU cores per instance. | |
|**Ram** | **int32** | The amount of memory per instance in gigabytes (GB). | |
|**StorageSize** | **int32** | The amount of storage per instance in gigabytes (GB). | |
|**Connections** | [**[]Connection**](Connection.md) | The network connection for your cluster. Only one connection is allowed.  | |
|**DisplayName** | **string** | The friendly name of your cluster. | |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |
|**Credentials** | [**DBUser**](DBUser.md) |  | |
|**FromBackup** | Pointer to [**RestoreRequest**](RestoreRequest.md) |  | [optional] |

## Methods

### NewCreateClusterProperties

`func NewCreateClusterProperties(mariadbVersion MariadbVersion, instances int32, cores int32, ram int32, storageSize int32, connections []Connection, displayName string, credentials DBUser, ) *CreateClusterProperties`

NewCreateClusterProperties instantiates a new CreateClusterProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterPropertiesWithDefaults

`func NewCreateClusterPropertiesWithDefaults() *CreateClusterProperties`

NewCreateClusterPropertiesWithDefaults instantiates a new CreateClusterProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMariadbVersion

`func (o *CreateClusterProperties) GetMariadbVersion() MariadbVersion`

GetMariadbVersion returns the MariadbVersion field if non-nil, zero value otherwise.

### GetMariadbVersionOk

`func (o *CreateClusterProperties) GetMariadbVersionOk() (*MariadbVersion, bool)`

GetMariadbVersionOk returns a tuple with the MariadbVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMariadbVersion

`func (o *CreateClusterProperties) SetMariadbVersion(v MariadbVersion)`

SetMariadbVersion sets MariadbVersion field to given value.


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


### GetFromBackup

`func (o *CreateClusterProperties) GetFromBackup() RestoreRequest`

GetFromBackup returns the FromBackup field if non-nil, zero value otherwise.

### GetFromBackupOk

`func (o *CreateClusterProperties) GetFromBackupOk() (*RestoreRequest, bool)`

GetFromBackupOk returns a tuple with the FromBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromBackup

`func (o *CreateClusterProperties) SetFromBackup(v RestoreRequest)`

SetFromBackup sets FromBackup field to given value.

### HasFromBackup

`func (o *CreateClusterProperties) HasFromBackup() bool`

HasFromBackup returns a boolean if a field has been set.


