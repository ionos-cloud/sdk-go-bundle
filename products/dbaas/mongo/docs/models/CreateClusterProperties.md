# CreateClusterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | Pointer to **string** | The cluster type, either &#x60;replicaset&#x60; or &#x60;sharded-cluster&#x60;.  | [optional] |
|**TemplateID** | Pointer to **string** | The unique ID of the template, which specifies the number of cores, storage size, and memory. You cannot downgrade to a smaller template or minor edition (e.g. from business to playground). To get a list of all templates to confirm the changes use the /templates endpoint.  | [optional] |
|**MongoDBVersion** | Pointer to **string** | The MongoDB version of your cluster. | [optional] |
|**Instances** | **int32** | The total number of instances in the cluster (one primary and n-1 secondaries).  | |
|**Shards** | Pointer to **int32** | The total number of shards in the cluster.  | [optional] |
|**Connections** | [**[]Connection**](Connection.md) |  | |
|**Location** | **string** | The physical location where the cluster will be created. This is the location where all your instances will be located. This property is immutable.  | |
|**Backup** | Pointer to [**BackupProperties**](BackupProperties.md) |  | [optional] |
|**DisplayName** | **string** | The name of your cluster. | |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |
|**BiConnector** | Pointer to [**BiConnectorProperties**](BiConnectorProperties.md) |  | [optional] |
|**FromBackup** | Pointer to [**CreateRestoreRequest**](CreateRestoreRequest.md) |  | [optional] |
|**Edition** | Pointer to **string** | The cluster edition. | [optional] |
|**Cores** | Pointer to **int32** | The number of CPU cores per instance. | [optional] |
|**Ram** | Pointer to **int32** | The amount of memory per instance in megabytes. Has to be a multiple of 1024. | [optional] |
|**StorageSize** | Pointer to **int32** | The amount of storage per instance in megabytes. | [optional] |
|**StorageType** | Pointer to [**StorageType**](StorageType.md) |  | [optional] |

## Methods

### NewCreateClusterProperties

`func NewCreateClusterProperties(instances int32, connections []Connection, location string, displayName string, ) *CreateClusterProperties`

NewCreateClusterProperties instantiates a new CreateClusterProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterPropertiesWithDefaults

`func NewCreateClusterPropertiesWithDefaults() *CreateClusterProperties`

NewCreateClusterPropertiesWithDefaults instantiates a new CreateClusterProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateClusterProperties) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateClusterProperties) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateClusterProperties) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateClusterProperties) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTemplateID

`func (o *CreateClusterProperties) GetTemplateID() string`

GetTemplateID returns the TemplateID field if non-nil, zero value otherwise.

### GetTemplateIDOk

`func (o *CreateClusterProperties) GetTemplateIDOk() (*string, bool)`

GetTemplateIDOk returns a tuple with the TemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateID

`func (o *CreateClusterProperties) SetTemplateID(v string)`

SetTemplateID sets TemplateID field to given value.

### HasTemplateID

`func (o *CreateClusterProperties) HasTemplateID() bool`

HasTemplateID returns a boolean if a field has been set.

### GetMongoDBVersion

`func (o *CreateClusterProperties) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *CreateClusterProperties) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *CreateClusterProperties) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *CreateClusterProperties) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.

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


### GetShards

`func (o *CreateClusterProperties) GetShards() int32`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *CreateClusterProperties) GetShardsOk() (*int32, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *CreateClusterProperties) SetShards(v int32)`

SetShards sets Shards field to given value.

### HasShards

`func (o *CreateClusterProperties) HasShards() bool`

HasShards returns a boolean if a field has been set.

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


### GetBackup

`func (o *CreateClusterProperties) GetBackup() BackupProperties`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *CreateClusterProperties) GetBackupOk() (*BackupProperties, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *CreateClusterProperties) SetBackup(v BackupProperties)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *CreateClusterProperties) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

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

### GetBiConnector

`func (o *CreateClusterProperties) GetBiConnector() BiConnectorProperties`

GetBiConnector returns the BiConnector field if non-nil, zero value otherwise.

### GetBiConnectorOk

`func (o *CreateClusterProperties) GetBiConnectorOk() (*BiConnectorProperties, bool)`

GetBiConnectorOk returns a tuple with the BiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiConnector

`func (o *CreateClusterProperties) SetBiConnector(v BiConnectorProperties)`

SetBiConnector sets BiConnector field to given value.

### HasBiConnector

`func (o *CreateClusterProperties) HasBiConnector() bool`

HasBiConnector returns a boolean if a field has been set.

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

### GetEdition

`func (o *CreateClusterProperties) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *CreateClusterProperties) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *CreateClusterProperties) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *CreateClusterProperties) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

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

### HasCores

`func (o *CreateClusterProperties) HasCores() bool`

HasCores returns a boolean if a field has been set.

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

### HasRam

`func (o *CreateClusterProperties) HasRam() bool`

HasRam returns a boolean if a field has been set.

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

### HasStorageSize

`func (o *CreateClusterProperties) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

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

### HasStorageType

`func (o *CreateClusterProperties) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.


