# ClusterCreateProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the In-Memory DB cluster. Must be 2-63 characters and must begin and end with an alphanumeric character ([A-Za-z0-9]), with dashes (-), underscores (_), and dots (.) allowed in between.  | |
|**Description** | Pointer to **string** | Human-readable description for the cluster. | [optional] |
|**Version** | **string** | The version for the cluster. Use GET /versions to retrieve the list of supported versions. To upgrade, provide a version listed in canUpgradeTo for the current version. Downgrades are not supported.  | |
|**Instances** | [**InstanceConfiguration**](InstanceConfiguration.md) |  | |
|**Connection** | [**ClusterConnection**](ClusterConnection.md) |  | |
|**PersistenceMode** | Pointer to [**PersistenceMode**](PersistenceMode.md) |  | [optional] [default to PERSISTENCEMODE_NONE]|
|**EvictionPolicy** | [**EvictionPolicy**](EvictionPolicy.md) |  | [default to EVICTIONPOLICY_ALLKEYS_LRU]|
|**Snapshot** | [**SnapshotConfiguration**](SnapshotConfiguration.md) |  | |
|**MaintenanceWindow** | [**MaintenanceWindow**](MaintenanceWindow.md) |  | |
|**Credentials** | [**ClusterCredentials**](ClusterCredentials.md) |  | |
|**RestoreFromSnapshot** | Pointer to [**ClusterRestoreFromSnapshot**](ClusterRestoreFromSnapshot.md) |  | [optional] |
|**LogsEnabled** | Pointer to **bool** | Activates or deactivates log collection and reporting for this cluster&#39;s observability. If the observability service is not activated on the contract, this setting is accepted but has no effect until the service is activated.  | [optional] [default to false]|
|**MetricsEnabled** | Pointer to **bool** | Activates or deactivates metrics collection and reporting for this cluster&#39;s observability. If the observability service is not activated on the contract, this setting is accepted but has no effect until the service is activated.  | [optional] [default to false]|

## Methods

### NewClusterCreateProperties

`func NewClusterCreateProperties(name string, version string, instances InstanceConfiguration, connection ClusterConnection, evictionPolicy EvictionPolicy, snapshot SnapshotConfiguration, maintenanceWindow MaintenanceWindow, credentials ClusterCredentials, ) *ClusterCreateProperties`

NewClusterCreateProperties instantiates a new ClusterCreateProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreatePropertiesWithDefaults

`func NewClusterCreatePropertiesWithDefaults() *ClusterCreateProperties`

NewClusterCreatePropertiesWithDefaults instantiates a new ClusterCreateProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterCreateProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterCreateProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterCreateProperties) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ClusterCreateProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterCreateProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterCreateProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterCreateProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *ClusterCreateProperties) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterCreateProperties) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterCreateProperties) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetInstances

`func (o *ClusterCreateProperties) GetInstances() InstanceConfiguration`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ClusterCreateProperties) GetInstancesOk() (*InstanceConfiguration, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ClusterCreateProperties) SetInstances(v InstanceConfiguration)`

SetInstances sets Instances field to given value.


### GetConnection

`func (o *ClusterCreateProperties) GetConnection() ClusterConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ClusterCreateProperties) GetConnectionOk() (*ClusterConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ClusterCreateProperties) SetConnection(v ClusterConnection)`

SetConnection sets Connection field to given value.


### GetPersistenceMode

`func (o *ClusterCreateProperties) GetPersistenceMode() PersistenceMode`

GetPersistenceMode returns the PersistenceMode field if non-nil, zero value otherwise.

### GetPersistenceModeOk

`func (o *ClusterCreateProperties) GetPersistenceModeOk() (*PersistenceMode, bool)`

GetPersistenceModeOk returns a tuple with the PersistenceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceMode

`func (o *ClusterCreateProperties) SetPersistenceMode(v PersistenceMode)`

SetPersistenceMode sets PersistenceMode field to given value.

### HasPersistenceMode

`func (o *ClusterCreateProperties) HasPersistenceMode() bool`

HasPersistenceMode returns a boolean if a field has been set.

### GetEvictionPolicy

`func (o *ClusterCreateProperties) GetEvictionPolicy() EvictionPolicy`

GetEvictionPolicy returns the EvictionPolicy field if non-nil, zero value otherwise.

### GetEvictionPolicyOk

`func (o *ClusterCreateProperties) GetEvictionPolicyOk() (*EvictionPolicy, bool)`

GetEvictionPolicyOk returns a tuple with the EvictionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvictionPolicy

`func (o *ClusterCreateProperties) SetEvictionPolicy(v EvictionPolicy)`

SetEvictionPolicy sets EvictionPolicy field to given value.


### GetSnapshot

`func (o *ClusterCreateProperties) GetSnapshot() SnapshotConfiguration`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *ClusterCreateProperties) GetSnapshotOk() (*SnapshotConfiguration, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *ClusterCreateProperties) SetSnapshot(v SnapshotConfiguration)`

SetSnapshot sets Snapshot field to given value.


### GetMaintenanceWindow

`func (o *ClusterCreateProperties) GetMaintenanceWindow() MaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *ClusterCreateProperties) GetMaintenanceWindowOk() (*MaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *ClusterCreateProperties) SetMaintenanceWindow(v MaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.


### GetCredentials

`func (o *ClusterCreateProperties) GetCredentials() ClusterCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ClusterCreateProperties) GetCredentialsOk() (*ClusterCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ClusterCreateProperties) SetCredentials(v ClusterCredentials)`

SetCredentials sets Credentials field to given value.


### GetRestoreFromSnapshot

`func (o *ClusterCreateProperties) GetRestoreFromSnapshot() ClusterRestoreFromSnapshot`

GetRestoreFromSnapshot returns the RestoreFromSnapshot field if non-nil, zero value otherwise.

### GetRestoreFromSnapshotOk

`func (o *ClusterCreateProperties) GetRestoreFromSnapshotOk() (*ClusterRestoreFromSnapshot, bool)`

GetRestoreFromSnapshotOk returns a tuple with the RestoreFromSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreFromSnapshot

`func (o *ClusterCreateProperties) SetRestoreFromSnapshot(v ClusterRestoreFromSnapshot)`

SetRestoreFromSnapshot sets RestoreFromSnapshot field to given value.

### HasRestoreFromSnapshot

`func (o *ClusterCreateProperties) HasRestoreFromSnapshot() bool`

HasRestoreFromSnapshot returns a boolean if a field has been set.

### GetLogsEnabled

`func (o *ClusterCreateProperties) GetLogsEnabled() bool`

GetLogsEnabled returns the LogsEnabled field if non-nil, zero value otherwise.

### GetLogsEnabledOk

`func (o *ClusterCreateProperties) GetLogsEnabledOk() (*bool, bool)`

GetLogsEnabledOk returns a tuple with the LogsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsEnabled

`func (o *ClusterCreateProperties) SetLogsEnabled(v bool)`

SetLogsEnabled sets LogsEnabled field to given value.

### HasLogsEnabled

`func (o *ClusterCreateProperties) HasLogsEnabled() bool`

HasLogsEnabled returns a boolean if a field has been set.

### GetMetricsEnabled

`func (o *ClusterCreateProperties) GetMetricsEnabled() bool`

GetMetricsEnabled returns the MetricsEnabled field if non-nil, zero value otherwise.

### GetMetricsEnabledOk

`func (o *ClusterCreateProperties) GetMetricsEnabledOk() (*bool, bool)`

GetMetricsEnabledOk returns a tuple with the MetricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsEnabled

`func (o *ClusterCreateProperties) SetMetricsEnabled(v bool)`

SetMetricsEnabled sets MetricsEnabled field to given value.

### HasMetricsEnabled

`func (o *ClusterCreateProperties) HasMetricsEnabled() bool`

HasMetricsEnabled returns a boolean if a field has been set.


