# EnterpriseCluster

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
|**Sharding** | Pointer to [**MongoSharding**](MongoSharding.md) |  | [optional] |
|**Instances** | [**InstanceConfiguration**](InstanceConfiguration.md) |  | |
|**BiConnectorEnabled** | Pointer to **bool** | A flag to enable or disable the MongoDB BI Connector for this cluster. Can only be enabled if the MongoDB edition is ENTERPRISE. When enabled, the BI Connector is reachable over the MySQL wire protocol at host &#x60;bi-connect.&lt;mongo-host&gt;&#x60; on the fixed port &#x60;27015&#x60; (for example, &#x60;bi-connect.m-12abd343.mongodb.gb-bhx.ionoscloud.com:27015&#x60;).  | [optional] [default to false]|
|**LogsEnabled** | Pointer to **bool** | Allows or disallows the collection and reporting of logs for this cluster&#39;s observability. If the observability service is not activated on the contract, this setting is accepted but has no effect; log collection will not be enabled until the observability service is activated.  | [optional] [default to false]|
|**MetricsEnabled** | Pointer to **bool** | Allows or disallows the collection and reporting of metrics for this cluster&#39;s observability. If the observability service is not activated on the contract, this setting is accepted but has no effect; metric collection will not be enabled until the observability service is activated.  | [optional] [default to false]|
|**Snapshot** | [**ClusterSnapshot**](ClusterSnapshot.md) |  | |

## Methods

### NewEnterpriseCluster

`func NewEnterpriseCluster(name string, version string, connection MongoClusterConnection, maintenanceWindow MaintenanceWindow, edition MongoEdition, instances InstanceConfiguration, snapshot ClusterSnapshot, ) *EnterpriseCluster`

NewEnterpriseCluster instantiates a new EnterpriseCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseClusterWithDefaults

`func NewEnterpriseClusterWithDefaults() *EnterpriseCluster`

NewEnterpriseClusterWithDefaults instantiates a new EnterpriseCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnterpriseCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnterpriseCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnterpriseCluster) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EnterpriseCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnterpriseCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnterpriseCluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnterpriseCluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *EnterpriseCluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EnterpriseCluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EnterpriseCluster) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetConnection

`func (o *EnterpriseCluster) GetConnection() MongoClusterConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *EnterpriseCluster) GetConnectionOk() (*MongoClusterConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *EnterpriseCluster) SetConnection(v MongoClusterConnection)`

SetConnection sets Connection field to given value.


### GetMaintenanceWindow

`func (o *EnterpriseCluster) GetMaintenanceWindow() MaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *EnterpriseCluster) GetMaintenanceWindowOk() (*MaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *EnterpriseCluster) SetMaintenanceWindow(v MaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.


### GetCredentials

`func (o *EnterpriseCluster) GetCredentials() MongoUser`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *EnterpriseCluster) GetCredentialsOk() (*MongoUser, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *EnterpriseCluster) SetCredentials(v MongoUser)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *EnterpriseCluster) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetRestoreFromSnapshot

`func (o *EnterpriseCluster) GetRestoreFromSnapshot() MongoRestoreFromSnapshot`

GetRestoreFromSnapshot returns the RestoreFromSnapshot field if non-nil, zero value otherwise.

### GetRestoreFromSnapshotOk

`func (o *EnterpriseCluster) GetRestoreFromSnapshotOk() (*MongoRestoreFromSnapshot, bool)`

GetRestoreFromSnapshotOk returns a tuple with the RestoreFromSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreFromSnapshot

`func (o *EnterpriseCluster) SetRestoreFromSnapshot(v MongoRestoreFromSnapshot)`

SetRestoreFromSnapshot sets RestoreFromSnapshot field to given value.

### HasRestoreFromSnapshot

`func (o *EnterpriseCluster) HasRestoreFromSnapshot() bool`

HasRestoreFromSnapshot returns a boolean if a field has been set.

### GetEdition

`func (o *EnterpriseCluster) GetEdition() MongoEdition`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *EnterpriseCluster) GetEditionOk() (*MongoEdition, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *EnterpriseCluster) SetEdition(v MongoEdition)`

SetEdition sets Edition field to given value.


### GetSharding

`func (o *EnterpriseCluster) GetSharding() MongoSharding`

GetSharding returns the Sharding field if non-nil, zero value otherwise.

### GetShardingOk

`func (o *EnterpriseCluster) GetShardingOk() (*MongoSharding, bool)`

GetShardingOk returns a tuple with the Sharding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharding

`func (o *EnterpriseCluster) SetSharding(v MongoSharding)`

SetSharding sets Sharding field to given value.

### HasSharding

`func (o *EnterpriseCluster) HasSharding() bool`

HasSharding returns a boolean if a field has been set.

### GetInstances

`func (o *EnterpriseCluster) GetInstances() InstanceConfiguration`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *EnterpriseCluster) GetInstancesOk() (*InstanceConfiguration, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *EnterpriseCluster) SetInstances(v InstanceConfiguration)`

SetInstances sets Instances field to given value.


### GetBiConnectorEnabled

`func (o *EnterpriseCluster) GetBiConnectorEnabled() bool`

GetBiConnectorEnabled returns the BiConnectorEnabled field if non-nil, zero value otherwise.

### GetBiConnectorEnabledOk

`func (o *EnterpriseCluster) GetBiConnectorEnabledOk() (*bool, bool)`

GetBiConnectorEnabledOk returns a tuple with the BiConnectorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiConnectorEnabled

`func (o *EnterpriseCluster) SetBiConnectorEnabled(v bool)`

SetBiConnectorEnabled sets BiConnectorEnabled field to given value.

### HasBiConnectorEnabled

`func (o *EnterpriseCluster) HasBiConnectorEnabled() bool`

HasBiConnectorEnabled returns a boolean if a field has been set.

### GetLogsEnabled

`func (o *EnterpriseCluster) GetLogsEnabled() bool`

GetLogsEnabled returns the LogsEnabled field if non-nil, zero value otherwise.

### GetLogsEnabledOk

`func (o *EnterpriseCluster) GetLogsEnabledOk() (*bool, bool)`

GetLogsEnabledOk returns a tuple with the LogsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsEnabled

`func (o *EnterpriseCluster) SetLogsEnabled(v bool)`

SetLogsEnabled sets LogsEnabled field to given value.

### HasLogsEnabled

`func (o *EnterpriseCluster) HasLogsEnabled() bool`

HasLogsEnabled returns a boolean if a field has been set.

### GetMetricsEnabled

`func (o *EnterpriseCluster) GetMetricsEnabled() bool`

GetMetricsEnabled returns the MetricsEnabled field if non-nil, zero value otherwise.

### GetMetricsEnabledOk

`func (o *EnterpriseCluster) GetMetricsEnabledOk() (*bool, bool)`

GetMetricsEnabledOk returns a tuple with the MetricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsEnabled

`func (o *EnterpriseCluster) SetMetricsEnabled(v bool)`

SetMetricsEnabled sets MetricsEnabled field to given value.

### HasMetricsEnabled

`func (o *EnterpriseCluster) HasMetricsEnabled() bool`

HasMetricsEnabled returns a boolean if a field has been set.

### GetSnapshot

`func (o *EnterpriseCluster) GetSnapshot() ClusterSnapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *EnterpriseCluster) GetSnapshotOk() (*ClusterSnapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *EnterpriseCluster) SetSnapshot(v ClusterSnapshot)`

SetSnapshot sets Snapshot field to given value.



