# Cluster

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of your PostgreSQL cluster. Must be 63 characters or less and must begin and end with an alphanumeric character (&#x60;[a-z0-9A-Z]&#x60;) with dashes (&#x60;-&#x60;), underscores (&#x60;_&#x60;), dots (&#x60;.&#x60;), and alphanumerics between.  | |
|**Description** | Pointer to **string** | Human-readable description for the cluster. | [optional] |
|**Version** | **string** | The PostgreSQL version for the cluster. | |
|**Instances** | [**InstanceConfiguration**](InstanceConfiguration.md) |  | |
|**Connection** | [**PostgresClusterConnection**](PostgresClusterConnection.md) |  | |
|**MaintenanceWindow** | [**MaintenanceWindow**](MaintenanceWindow.md) |  | |
|**ReplicationMode** | [**PostgresClusterReplicationMode**](PostgresClusterReplicationMode.md) |  | |
|**Credentials** | Pointer to [**PostgresUser**](PostgresUser.md) |  | [optional] |
|**ConnectionPooler** | Pointer to **string** | Defines how database connections are managed and reused. Default value is DISABLED. DISABLED: No connection pooling is used. Each request opens a new connection, which is closed immediately after use. It ensures isolation but may impact performance due to frequent connection setup and teardown. TRANSACTION: Connections are pooled and reused for the duration of a transaction. Once the transaction completes, the connection is returned to the pool. This mode balances efficiency with transactional integrity. SESSION: Connections are retained for the entire session and reused across multiple transactions. Offers the highest performance by minimizing connection overhead, but may tie up resources longer.  | [optional] |
|**RestoreFromBackup** | Pointer to [**ClusterRestoreFromBackup**](ClusterRestoreFromBackup.md) |  | [optional] |
|**LogsEnabled** | Pointer to **bool** | Allows or disallows the collection and reporting of logs for this cluster&#39;s observability. If the observability service is not activated on the contract, this setting is accepted but has no effect; log collection will not be enabled until the observability service is activated.  | [optional] [default to false]|
|**MetricsEnabled** | Pointer to **bool** | Allows or disallows the collection and reporting of metrics for this cluster&#39;s observability. If the observability service is not activated on the contract, this setting is accepted but has no effect; metric collection will not be enabled until the observability service is activated.  | [optional] [default to false]|
|**Backup** | [**ClusterBackup**](ClusterBackup.md) |  | |

## Methods

### NewCluster

`func NewCluster(name string, version string, instances InstanceConfiguration, connection PostgresClusterConnection, maintenanceWindow MaintenanceWindow, replicationMode PostgresClusterReplicationMode, backup ClusterBackup, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Cluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Cluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Cluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Cluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *Cluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Cluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Cluster) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetInstances

`func (o *Cluster) GetInstances() InstanceConfiguration`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *Cluster) GetInstancesOk() (*InstanceConfiguration, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *Cluster) SetInstances(v InstanceConfiguration)`

SetInstances sets Instances field to given value.


### GetConnection

`func (o *Cluster) GetConnection() PostgresClusterConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *Cluster) GetConnectionOk() (*PostgresClusterConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *Cluster) SetConnection(v PostgresClusterConnection)`

SetConnection sets Connection field to given value.


### GetMaintenanceWindow

`func (o *Cluster) GetMaintenanceWindow() MaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *Cluster) GetMaintenanceWindowOk() (*MaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *Cluster) SetMaintenanceWindow(v MaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.


### GetReplicationMode

`func (o *Cluster) GetReplicationMode() PostgresClusterReplicationMode`

GetReplicationMode returns the ReplicationMode field if non-nil, zero value otherwise.

### GetReplicationModeOk

`func (o *Cluster) GetReplicationModeOk() (*PostgresClusterReplicationMode, bool)`

GetReplicationModeOk returns a tuple with the ReplicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMode

`func (o *Cluster) SetReplicationMode(v PostgresClusterReplicationMode)`

SetReplicationMode sets ReplicationMode field to given value.


### GetCredentials

`func (o *Cluster) GetCredentials() PostgresUser`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Cluster) GetCredentialsOk() (*PostgresUser, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Cluster) SetCredentials(v PostgresUser)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *Cluster) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetConnectionPooler

`func (o *Cluster) GetConnectionPooler() string`

GetConnectionPooler returns the ConnectionPooler field if non-nil, zero value otherwise.

### GetConnectionPoolerOk

`func (o *Cluster) GetConnectionPoolerOk() (*string, bool)`

GetConnectionPoolerOk returns a tuple with the ConnectionPooler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPooler

`func (o *Cluster) SetConnectionPooler(v string)`

SetConnectionPooler sets ConnectionPooler field to given value.

### HasConnectionPooler

`func (o *Cluster) HasConnectionPooler() bool`

HasConnectionPooler returns a boolean if a field has been set.

### GetRestoreFromBackup

`func (o *Cluster) GetRestoreFromBackup() ClusterRestoreFromBackup`

GetRestoreFromBackup returns the RestoreFromBackup field if non-nil, zero value otherwise.

### GetRestoreFromBackupOk

`func (o *Cluster) GetRestoreFromBackupOk() (*ClusterRestoreFromBackup, bool)`

GetRestoreFromBackupOk returns a tuple with the RestoreFromBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreFromBackup

`func (o *Cluster) SetRestoreFromBackup(v ClusterRestoreFromBackup)`

SetRestoreFromBackup sets RestoreFromBackup field to given value.

### HasRestoreFromBackup

`func (o *Cluster) HasRestoreFromBackup() bool`

HasRestoreFromBackup returns a boolean if a field has been set.

### GetLogsEnabled

`func (o *Cluster) GetLogsEnabled() bool`

GetLogsEnabled returns the LogsEnabled field if non-nil, zero value otherwise.

### GetLogsEnabledOk

`func (o *Cluster) GetLogsEnabledOk() (*bool, bool)`

GetLogsEnabledOk returns a tuple with the LogsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsEnabled

`func (o *Cluster) SetLogsEnabled(v bool)`

SetLogsEnabled sets LogsEnabled field to given value.

### HasLogsEnabled

`func (o *Cluster) HasLogsEnabled() bool`

HasLogsEnabled returns a boolean if a field has been set.

### GetMetricsEnabled

`func (o *Cluster) GetMetricsEnabled() bool`

GetMetricsEnabled returns the MetricsEnabled field if non-nil, zero value otherwise.

### GetMetricsEnabledOk

`func (o *Cluster) GetMetricsEnabledOk() (*bool, bool)`

GetMetricsEnabledOk returns a tuple with the MetricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsEnabled

`func (o *Cluster) SetMetricsEnabled(v bool)`

SetMetricsEnabled sets MetricsEnabled field to given value.

### HasMetricsEnabled

`func (o *Cluster) HasMetricsEnabled() bool`

HasMetricsEnabled returns a boolean if a field has been set.

### GetBackup

`func (o *Cluster) GetBackup() ClusterBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *Cluster) GetBackupOk() (*ClusterBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *Cluster) SetBackup(v ClusterBackup)`

SetBackup sets Backup field to given value.



