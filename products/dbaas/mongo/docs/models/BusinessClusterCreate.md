# BusinessClusterCreate

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the MongoDB cluster. Must be 63 characters or less and must begin and end with an alphanumeric character (&#x60;[a-z0-9A-Z]&#x60;) with dashes (&#x60;-&#x60;), underscores (&#x60;_&#x60;), dots (&#x60;.&#x60;), and alphanumerics between.  | |
|**Description** | Pointer to **string** | Human-readable description for the cluster. | [optional] |
|**Version** | **string** | The MongoDB version for the cluster. Use GET /versions to retrieve the list of supported versions and their available upgrade paths.  | |
|**Connection** | [**MongoClusterConnection**](MongoClusterConnection.md) |  | |
|**MaintenanceWindow** | [**MaintenanceWindow**](MaintenanceWindow.md) |  | |
|**Credentials** | [**MongoUser**](MongoUser.md) |  | |
|**RestoreFromSnapshot** | Pointer to [**MongoRestoreFromSnapshot**](MongoRestoreFromSnapshot.md) |  | [optional] |
|**Edition** | [**MongoEdition**](MongoEdition.md) |  | |
|**InstanceTemplateId** | **string** | A unique identifier for the instance template. Check /instance-templates for a list. | |
|**InstanceCount** | **int32** | The total number of instances in the cluster (one primary and n-1 secondary).  | |
|**LogsEnabled** | Pointer to **bool** | Allows or disallows the collection and reporting of logs for this cluster&#39;s observability. If the observability service is not activated on the contract, this setting is accepted but has no effect; log collection will not be enabled until the observability service is activated.  | [optional] [default to false]|
|**MetricsEnabled** | Pointer to **bool** | Allows or disallows the collection and reporting of metrics for this cluster&#39;s observability. If the observability service is not activated on the contract, this setting is accepted but has no effect; metric collection will not be enabled until the observability service is activated.  | [optional] [default to false]|
|**Snapshot** | [**ClusterSnapshot**](ClusterSnapshot.md) |  | |

## Methods

### NewBusinessClusterCreate

`func NewBusinessClusterCreate(name string, version string, connection MongoClusterConnection, maintenanceWindow MaintenanceWindow, credentials MongoUser, edition MongoEdition, instanceTemplateId string, instanceCount int32, snapshot ClusterSnapshot, ) *BusinessClusterCreate`

NewBusinessClusterCreate instantiates a new BusinessClusterCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessClusterCreateWithDefaults

`func NewBusinessClusterCreateWithDefaults() *BusinessClusterCreate`

NewBusinessClusterCreateWithDefaults instantiates a new BusinessClusterCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BusinessClusterCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BusinessClusterCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BusinessClusterCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BusinessClusterCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BusinessClusterCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BusinessClusterCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BusinessClusterCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *BusinessClusterCreate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BusinessClusterCreate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BusinessClusterCreate) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetConnection

`func (o *BusinessClusterCreate) GetConnection() MongoClusterConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *BusinessClusterCreate) GetConnectionOk() (*MongoClusterConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *BusinessClusterCreate) SetConnection(v MongoClusterConnection)`

SetConnection sets Connection field to given value.


### GetMaintenanceWindow

`func (o *BusinessClusterCreate) GetMaintenanceWindow() MaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *BusinessClusterCreate) GetMaintenanceWindowOk() (*MaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *BusinessClusterCreate) SetMaintenanceWindow(v MaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.


### GetCredentials

`func (o *BusinessClusterCreate) GetCredentials() MongoUser`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *BusinessClusterCreate) GetCredentialsOk() (*MongoUser, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *BusinessClusterCreate) SetCredentials(v MongoUser)`

SetCredentials sets Credentials field to given value.


### GetRestoreFromSnapshot

`func (o *BusinessClusterCreate) GetRestoreFromSnapshot() MongoRestoreFromSnapshot`

GetRestoreFromSnapshot returns the RestoreFromSnapshot field if non-nil, zero value otherwise.

### GetRestoreFromSnapshotOk

`func (o *BusinessClusterCreate) GetRestoreFromSnapshotOk() (*MongoRestoreFromSnapshot, bool)`

GetRestoreFromSnapshotOk returns a tuple with the RestoreFromSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreFromSnapshot

`func (o *BusinessClusterCreate) SetRestoreFromSnapshot(v MongoRestoreFromSnapshot)`

SetRestoreFromSnapshot sets RestoreFromSnapshot field to given value.

### HasRestoreFromSnapshot

`func (o *BusinessClusterCreate) HasRestoreFromSnapshot() bool`

HasRestoreFromSnapshot returns a boolean if a field has been set.

### GetEdition

`func (o *BusinessClusterCreate) GetEdition() MongoEdition`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *BusinessClusterCreate) GetEditionOk() (*MongoEdition, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *BusinessClusterCreate) SetEdition(v MongoEdition)`

SetEdition sets Edition field to given value.


### GetInstanceTemplateId

`func (o *BusinessClusterCreate) GetInstanceTemplateId() string`

GetInstanceTemplateId returns the InstanceTemplateId field if non-nil, zero value otherwise.

### GetInstanceTemplateIdOk

`func (o *BusinessClusterCreate) GetInstanceTemplateIdOk() (*string, bool)`

GetInstanceTemplateIdOk returns a tuple with the InstanceTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTemplateId

`func (o *BusinessClusterCreate) SetInstanceTemplateId(v string)`

SetInstanceTemplateId sets InstanceTemplateId field to given value.


### GetInstanceCount

`func (o *BusinessClusterCreate) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *BusinessClusterCreate) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *BusinessClusterCreate) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.


### GetLogsEnabled

`func (o *BusinessClusterCreate) GetLogsEnabled() bool`

GetLogsEnabled returns the LogsEnabled field if non-nil, zero value otherwise.

### GetLogsEnabledOk

`func (o *BusinessClusterCreate) GetLogsEnabledOk() (*bool, bool)`

GetLogsEnabledOk returns a tuple with the LogsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsEnabled

`func (o *BusinessClusterCreate) SetLogsEnabled(v bool)`

SetLogsEnabled sets LogsEnabled field to given value.

### HasLogsEnabled

`func (o *BusinessClusterCreate) HasLogsEnabled() bool`

HasLogsEnabled returns a boolean if a field has been set.

### GetMetricsEnabled

`func (o *BusinessClusterCreate) GetMetricsEnabled() bool`

GetMetricsEnabled returns the MetricsEnabled field if non-nil, zero value otherwise.

### GetMetricsEnabledOk

`func (o *BusinessClusterCreate) GetMetricsEnabledOk() (*bool, bool)`

GetMetricsEnabledOk returns a tuple with the MetricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsEnabled

`func (o *BusinessClusterCreate) SetMetricsEnabled(v bool)`

SetMetricsEnabled sets MetricsEnabled field to given value.

### HasMetricsEnabled

`func (o *BusinessClusterCreate) HasMetricsEnabled() bool`

HasMetricsEnabled returns a boolean if a field has been set.

### GetSnapshot

`func (o *BusinessClusterCreate) GetSnapshot() ClusterSnapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *BusinessClusterCreate) GetSnapshotOk() (*ClusterSnapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *BusinessClusterCreate) SetSnapshot(v ClusterSnapshot)`

SetSnapshot sets Snapshot field to given value.



