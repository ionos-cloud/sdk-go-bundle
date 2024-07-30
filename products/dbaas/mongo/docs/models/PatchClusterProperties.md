# PatchClusterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | Pointer to **string** | The cluster type, either &#x60;replicaset&#x60; or &#x60;sharded-cluster&#x60;.  | [optional] |
|**DisplayName** | Pointer to **string** | The name of your cluster. | [optional] |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |
|**Instances** | Pointer to **int32** | The total number of instances in the cluster (one primary and n-1 secondaries).  | [optional] |
|**Shards** | Pointer to **int32** | The total number of shards in the cluster.  | [optional] |
|**Backup** | Pointer to [**BackupProperties**](BackupProperties.md) |  | [optional] |
|**BiConnector** | Pointer to [**BiConnectorProperties**](BiConnectorProperties.md) |  | [optional] |
|**Connections** | Pointer to [**[]Connection**](Connection.md) |  | [optional] |
|**TemplateID** | Pointer to **string** | The unique ID of the template, which specifies the number of cores, storage size, and memory. You cannot downgrade to a smaller template or minor edition (e.g. from business to playground). To get a list of all templates to confirm the changes use the /templates endpoint.  | [optional] |
|**Edition** | Pointer to **string** | The cluster edition. | [optional] |
|**Cores** | Pointer to **int32** | The number of CPU cores per instance. | [optional] |
|**Ram** | Pointer to **int32** | The amount of memory per instance in megabytes. Has to be a multiple of 1024. | [optional] |
|**StorageSize** | Pointer to **int32** | The amount of storage per instance in megabytes. | [optional] |
|**StorageType** | Pointer to [**StorageType**](StorageType.md) |  | [optional] |

## Methods

### NewPatchClusterProperties

`func NewPatchClusterProperties() *PatchClusterProperties`

NewPatchClusterProperties instantiates a new PatchClusterProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchClusterPropertiesWithDefaults

`func NewPatchClusterPropertiesWithDefaults() *PatchClusterProperties`

NewPatchClusterPropertiesWithDefaults instantiates a new PatchClusterProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PatchClusterProperties) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchClusterProperties) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchClusterProperties) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchClusterProperties) HasType() bool`

HasType returns a boolean if a field has been set.

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

### GetShards

`func (o *PatchClusterProperties) GetShards() int32`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *PatchClusterProperties) GetShardsOk() (*int32, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *PatchClusterProperties) SetShards(v int32)`

SetShards sets Shards field to given value.

### HasShards

`func (o *PatchClusterProperties) HasShards() bool`

HasShards returns a boolean if a field has been set.

### GetBackup

`func (o *PatchClusterProperties) GetBackup() BackupProperties`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *PatchClusterProperties) GetBackupOk() (*BackupProperties, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *PatchClusterProperties) SetBackup(v BackupProperties)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *PatchClusterProperties) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetBiConnector

`func (o *PatchClusterProperties) GetBiConnector() BiConnectorProperties`

GetBiConnector returns the BiConnector field if non-nil, zero value otherwise.

### GetBiConnectorOk

`func (o *PatchClusterProperties) GetBiConnectorOk() (*BiConnectorProperties, bool)`

GetBiConnectorOk returns a tuple with the BiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiConnector

`func (o *PatchClusterProperties) SetBiConnector(v BiConnectorProperties)`

SetBiConnector sets BiConnector field to given value.

### HasBiConnector

`func (o *PatchClusterProperties) HasBiConnector() bool`

HasBiConnector returns a boolean if a field has been set.

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

### HasConnections

`func (o *PatchClusterProperties) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetTemplateID

`func (o *PatchClusterProperties) GetTemplateID() string`

GetTemplateID returns the TemplateID field if non-nil, zero value otherwise.

### GetTemplateIDOk

`func (o *PatchClusterProperties) GetTemplateIDOk() (*string, bool)`

GetTemplateIDOk returns a tuple with the TemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateID

`func (o *PatchClusterProperties) SetTemplateID(v string)`

SetTemplateID sets TemplateID field to given value.

### HasTemplateID

`func (o *PatchClusterProperties) HasTemplateID() bool`

HasTemplateID returns a boolean if a field has been set.

### GetEdition

`func (o *PatchClusterProperties) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *PatchClusterProperties) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *PatchClusterProperties) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *PatchClusterProperties) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

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

### GetStorageType

`func (o *PatchClusterProperties) GetStorageType() StorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *PatchClusterProperties) GetStorageTypeOk() (*StorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *PatchClusterProperties) SetStorageType(v StorageType)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *PatchClusterProperties) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.


