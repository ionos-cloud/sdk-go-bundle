# ReplicaSet

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DisplayName** | **string** | The human readable name of your replica set. | |
|**Version** | **string** | The In-Memory DB version of your replica set. | |
|**Replicas** | **int32** | The total number of replicas in the replica set (one active and n-1 passive). In case of a standalone instance, the value is 1. In all other cases, the value is &gt;1. The replicas will not be available as read replicas, they are only standby for a failure of the active instance.  | |
|**Resources** | [**Resources**](Resources.md) |  | |
|**PersistenceMode** | [**PersistenceMode**](PersistenceMode.md) |  | [default to PERSISTENCEMODE_NONE]|
|**EvictionPolicy** | [**EvictionPolicy**](EvictionPolicy.md) |  | [default to EVICTIONPOLICY_ALLKEYS_LRU]|
|**Connections** | [**[]Connection**](Connection.md) | The network connection for your replica set. Only one connection is allowed.  | |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |
|**Backup** | Pointer to [**BackupProperties**](BackupProperties.md) |  | [optional] |
|**Credentials** | [**User**](User.md) |  | |
|**InitialSnapshotId** | Pointer to **string** | The ID of a snapshot to restore the replica set from. If set, the replica set will be created from the snapshot.  | [optional] |

## Methods

### NewReplicaSet

`func NewReplicaSet(displayName string, version string, replicas int32, resources Resources, persistenceMode PersistenceMode, evictionPolicy EvictionPolicy, connections []Connection, credentials User, ) *ReplicaSet`

NewReplicaSet instantiates a new ReplicaSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaSetWithDefaults

`func NewReplicaSetWithDefaults() *ReplicaSet`

NewReplicaSetWithDefaults instantiates a new ReplicaSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ReplicaSet) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ReplicaSet) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ReplicaSet) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetVersion

`func (o *ReplicaSet) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReplicaSet) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReplicaSet) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetReplicas

`func (o *ReplicaSet) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *ReplicaSet) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *ReplicaSet) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetResources

`func (o *ReplicaSet) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ReplicaSet) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ReplicaSet) SetResources(v Resources)`

SetResources sets Resources field to given value.


### GetPersistenceMode

`func (o *ReplicaSet) GetPersistenceMode() PersistenceMode`

GetPersistenceMode returns the PersistenceMode field if non-nil, zero value otherwise.

### GetPersistenceModeOk

`func (o *ReplicaSet) GetPersistenceModeOk() (*PersistenceMode, bool)`

GetPersistenceModeOk returns a tuple with the PersistenceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceMode

`func (o *ReplicaSet) SetPersistenceMode(v PersistenceMode)`

SetPersistenceMode sets PersistenceMode field to given value.


### GetEvictionPolicy

`func (o *ReplicaSet) GetEvictionPolicy() EvictionPolicy`

GetEvictionPolicy returns the EvictionPolicy field if non-nil, zero value otherwise.

### GetEvictionPolicyOk

`func (o *ReplicaSet) GetEvictionPolicyOk() (*EvictionPolicy, bool)`

GetEvictionPolicyOk returns a tuple with the EvictionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvictionPolicy

`func (o *ReplicaSet) SetEvictionPolicy(v EvictionPolicy)`

SetEvictionPolicy sets EvictionPolicy field to given value.


### GetConnections

`func (o *ReplicaSet) GetConnections() []Connection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ReplicaSet) GetConnectionsOk() (*[]Connection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ReplicaSet) SetConnections(v []Connection)`

SetConnections sets Connections field to given value.


### GetMaintenanceWindow

`func (o *ReplicaSet) GetMaintenanceWindow() MaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *ReplicaSet) GetMaintenanceWindowOk() (*MaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *ReplicaSet) SetMaintenanceWindow(v MaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *ReplicaSet) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.

### GetBackup

`func (o *ReplicaSet) GetBackup() BackupProperties`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *ReplicaSet) GetBackupOk() (*BackupProperties, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *ReplicaSet) SetBackup(v BackupProperties)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *ReplicaSet) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetCredentials

`func (o *ReplicaSet) GetCredentials() User`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ReplicaSet) GetCredentialsOk() (*User, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ReplicaSet) SetCredentials(v User)`

SetCredentials sets Credentials field to given value.


### GetInitialSnapshotId

`func (o *ReplicaSet) GetInitialSnapshotId() string`

GetInitialSnapshotId returns the InitialSnapshotId field if non-nil, zero value otherwise.

### GetInitialSnapshotIdOk

`func (o *ReplicaSet) GetInitialSnapshotIdOk() (*string, bool)`

GetInitialSnapshotIdOk returns a tuple with the InitialSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSnapshotId

`func (o *ReplicaSet) SetInitialSnapshotId(v string)`

SetInitialSnapshotId sets InitialSnapshotId field to given value.

### HasInitialSnapshotId

`func (o *ReplicaSet) HasInitialSnapshotId() bool`

HasInitialSnapshotId returns a boolean if a field has been set.


