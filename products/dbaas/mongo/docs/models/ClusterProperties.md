# ClusterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | Pointer to **string** | The cluster type, either &#x60;replicaset&#x60; or &#x60;sharded-cluster&#x60;.  | [optional] |
|**DisplayName** | Pointer to **string** | The name of your cluster. | [optional] |
|**MongoDBVersion** | Pointer to **string** | The MongoDB version of your cluster. | [optional] |
|**Location** | Pointer to **string** | The physical location where the cluster will be created. This is the location where all your instances will be located. This property is immutable.  | [optional] |
|**Backup** | Pointer to [**BackupProperties**](BackupProperties.md) |  | [optional] |
|**Instances** | Pointer to **int32** | The total number of instances in the cluster (one primary and n-1 secondaries).  | [optional] |
|**Shards** | Pointer to **int32** | The total number of shards in the cluster.  | [optional] |
|**Connections** | Pointer to [**[]Connection**](Connection.md) |  | [optional] |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |
|**TemplateID** | Pointer to **string** | The unique ID of the template, which specifies the number of cores, storage size, and memory. You cannot downgrade to a smaller template or minor edition (e.g. from business to playground). To get a list of all templates to confirm the changes use the /templates endpoint.  | [optional] |
|**ConnectionString** | Pointer to **string** | The connection string for your cluster. | [optional] |
|**BiConnector** | Pointer to [**BiConnectorProperties**](BiConnectorProperties.md) |  | [optional] |
|**Edition** | Pointer to **string** | The cluster edition. | [optional] |
|**Cores** | Pointer to **int32** | The number of CPU cores per instance. | [optional] |
|**Ram** | Pointer to **int32** | The amount of memory per instance in megabytes. Has to be a multiple of 1024. | [optional] |
|**StorageSize** | Pointer to **int32** | The amount of storage per instance in megabytes. | [optional] |
|**StorageType** | Pointer to [**StorageType**](StorageType.md) |  | [optional] |

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

### GetType

`func (o *ClusterProperties) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterProperties) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterProperties) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterProperties) HasType() bool`

HasType returns a boolean if a field has been set.

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

### GetMongoDBVersion

`func (o *ClusterProperties) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *ClusterProperties) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *ClusterProperties) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *ClusterProperties) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.

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

### GetBackup

`func (o *ClusterProperties) GetBackup() BackupProperties`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *ClusterProperties) GetBackupOk() (*BackupProperties, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *ClusterProperties) SetBackup(v BackupProperties)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *ClusterProperties) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

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

### GetShards

`func (o *ClusterProperties) GetShards() int32`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *ClusterProperties) GetShardsOk() (*int32, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *ClusterProperties) SetShards(v int32)`

SetShards sets Shards field to given value.

### HasShards

`func (o *ClusterProperties) HasShards() bool`

HasShards returns a boolean if a field has been set.

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

### GetTemplateID

`func (o *ClusterProperties) GetTemplateID() string`

GetTemplateID returns the TemplateID field if non-nil, zero value otherwise.

### GetTemplateIDOk

`func (o *ClusterProperties) GetTemplateIDOk() (*string, bool)`

GetTemplateIDOk returns a tuple with the TemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateID

`func (o *ClusterProperties) SetTemplateID(v string)`

SetTemplateID sets TemplateID field to given value.

### HasTemplateID

`func (o *ClusterProperties) HasTemplateID() bool`

HasTemplateID returns a boolean if a field has been set.

### GetConnectionString

`func (o *ClusterProperties) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *ClusterProperties) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *ClusterProperties) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.

### HasConnectionString

`func (o *ClusterProperties) HasConnectionString() bool`

HasConnectionString returns a boolean if a field has been set.

### GetBiConnector

`func (o *ClusterProperties) GetBiConnector() BiConnectorProperties`

GetBiConnector returns the BiConnector field if non-nil, zero value otherwise.

### GetBiConnectorOk

`func (o *ClusterProperties) GetBiConnectorOk() (*BiConnectorProperties, bool)`

GetBiConnectorOk returns a tuple with the BiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiConnector

`func (o *ClusterProperties) SetBiConnector(v BiConnectorProperties)`

SetBiConnector sets BiConnector field to given value.

### HasBiConnector

`func (o *ClusterProperties) HasBiConnector() bool`

HasBiConnector returns a boolean if a field has been set.

### GetEdition

`func (o *ClusterProperties) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *ClusterProperties) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *ClusterProperties) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *ClusterProperties) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

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


