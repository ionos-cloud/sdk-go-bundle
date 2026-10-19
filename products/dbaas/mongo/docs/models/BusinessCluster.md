# BusinessCluster

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the MongoDB cluster. Must be 63 characters or less and must begin and end with an alphanumeric character (&#x60;[a-z0-9A-Z]&#x60;) with dashes (&#x60;-&#x60;), underscores (&#x60;_&#x60;), dots (&#x60;.&#x60;), and alphanumerics between.  | |
|**Description** | Pointer to **string** | Human-readable description for the cluster. | [optional] |
|**Version** | **string** | The MongoDB version for the cluster. Use GET /versions to retrieve the list of supported versions and their available upgrade paths.  | |
|**Connection** | [**MongoClusterConnection**](MongoClusterConnection.md) |  | |
|**MaintenanceWindow** | [**MaintenanceWindow**](MaintenanceWindow.md) |  | |
|**Credentials** | Pointer to [**MongoUser**](MongoUser.md) |  | [optional] |
|**RestoreFromSnapshot** | Pointer to [**MongoRestoreFromSnapshot**](MongoRestoreFromSnapshot.md) |  | [optional] |
|**Edition** | [**MongoEdition**](MongoEdition.md) |  | |
|**InstanceTemplateId** | **string** | A unique identifier for the instance template. Check /instance-templates for a list. | |
|**InstanceCount** | **int32** | The total number of instances in the cluster (one primary and n-1 secondary).  | |
|**LogsEnabled** | Pointer to **bool** | Allows or disallows the collection and reporting of logs for this cluster&#39;s observability. If the observability service is not activated on the contract, this setting is accepted but has no effect; log collection will not be enabled until the observability service is activated.  | [optional] [default to false]|
|**MetricsEnabled** | Pointer to **bool** | Allows or disallows the collection and reporting of metrics for this cluster&#39;s observability. If the observability service is not activated on the contract, this setting is accepted but has no effect; metric collection will not be enabled until the observability service is activated.  | [optional] [default to false]|
|**Snapshot** | [**ClusterSnapshot**](ClusterSnapshot.md) |  | |

## Methods

### NewBusinessCluster

`func NewBusinessCluster(name string, version string, connection MongoClusterConnection, maintenanceWindow MaintenanceWindow, edition MongoEdition, instanceTemplateId string, instanceCount int32, snapshot ClusterSnapshot, ) *BusinessCluster`

NewBusinessCluster instantiates a new BusinessCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessClusterWithDefaults

`func NewBusinessClusterWithDefaults() *BusinessCluster`

NewBusinessClusterWithDefaults instantiates a new BusinessCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BusinessCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BusinessCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BusinessCluster) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BusinessCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BusinessCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BusinessCluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BusinessCluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *BusinessCluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BusinessCluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BusinessCluster) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetConnection

`func (o *BusinessCluster) GetConnection() MongoClusterConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *BusinessCluster) GetConnectionOk() (*MongoClusterConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *BusinessCluster) SetConnection(v MongoClusterConnection)`

SetConnection sets Connection field to given value.


### GetMaintenanceWindow

`func (o *BusinessCluster) GetMaintenanceWindow() MaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *BusinessCluster) GetMaintenanceWindowOk() (*MaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *BusinessCluster) SetMaintenanceWindow(v MaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.


### GetCredentials

`func (o *BusinessCluster) GetCredentials() MongoUser`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *BusinessCluster) GetCredentialsOk() (*MongoUser, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *BusinessCluster) SetCredentials(v MongoUser)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *BusinessCluster) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetRestoreFromSnapshot

`func (o *BusinessCluster) GetRestoreFromSnapshot() MongoRestoreFromSnapshot`

GetRestoreFromSnapshot returns the RestoreFromSnapshot field if non-nil, zero value otherwise.

### GetRestoreFromSnapshotOk

`func (o *BusinessCluster) GetRestoreFromSnapshotOk() (*MongoRestoreFromSnapshot, bool)`

GetRestoreFromSnapshotOk returns a tuple with the RestoreFromSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreFromSnapshot

`func (o *BusinessCluster) SetRestoreFromSnapshot(v MongoRestoreFromSnapshot)`

SetRestoreFromSnapshot sets RestoreFromSnapshot field to given value.

### HasRestoreFromSnapshot

`func (o *BusinessCluster) HasRestoreFromSnapshot() bool`

HasRestoreFromSnapshot returns a boolean if a field has been set.

### GetEdition

`func (o *BusinessCluster) GetEdition() MongoEdition`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *BusinessCluster) GetEditionOk() (*MongoEdition, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *BusinessCluster) SetEdition(v MongoEdition)`

SetEdition sets Edition field to given value.


### GetInstanceTemplateId

`func (o *BusinessCluster) GetInstanceTemplateId() string`

GetInstanceTemplateId returns the InstanceTemplateId field if non-nil, zero value otherwise.

### GetInstanceTemplateIdOk

`func (o *BusinessCluster) GetInstanceTemplateIdOk() (*string, bool)`

GetInstanceTemplateIdOk returns a tuple with the InstanceTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTemplateId

`func (o *BusinessCluster) SetInstanceTemplateId(v string)`

SetInstanceTemplateId sets InstanceTemplateId field to given value.


### GetInstanceCount

`func (o *BusinessCluster) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *BusinessCluster) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *BusinessCluster) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.


### GetLogsEnabled

`func (o *BusinessCluster) GetLogsEnabled() bool`

GetLogsEnabled returns the LogsEnabled field if non-nil, zero value otherwise.

### GetLogsEnabledOk

`func (o *BusinessCluster) GetLogsEnabledOk() (*bool, bool)`

GetLogsEnabledOk returns a tuple with the LogsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsEnabled

`func (o *BusinessCluster) SetLogsEnabled(v bool)`

SetLogsEnabled sets LogsEnabled field to given value.

### HasLogsEnabled

`func (o *BusinessCluster) HasLogsEnabled() bool`

HasLogsEnabled returns a boolean if a field has been set.

### GetMetricsEnabled

`func (o *BusinessCluster) GetMetricsEnabled() bool`

GetMetricsEnabled returns the MetricsEnabled field if non-nil, zero value otherwise.

### GetMetricsEnabledOk

`func (o *BusinessCluster) GetMetricsEnabledOk() (*bool, bool)`

GetMetricsEnabledOk returns a tuple with the MetricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsEnabled

`func (o *BusinessCluster) SetMetricsEnabled(v bool)`

SetMetricsEnabled sets MetricsEnabled field to given value.

### HasMetricsEnabled

`func (o *BusinessCluster) HasMetricsEnabled() bool`

HasMetricsEnabled returns a boolean if a field has been set.

### GetSnapshot

`func (o *BusinessCluster) GetSnapshot() ClusterSnapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *BusinessCluster) GetSnapshotOk() (*ClusterSnapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *BusinessCluster) SetSnapshot(v ClusterSnapshot)`

SetSnapshot sets Snapshot field to given value.



