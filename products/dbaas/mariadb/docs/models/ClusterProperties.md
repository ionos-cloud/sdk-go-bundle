# ClusterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DisplayName** | Pointer to **string** | The friendly name of your cluster. | [optional] |
|**MariadbVersion** | Pointer to [**MariadbVersion**](MariadbVersion.md) |  | [optional] |
|**DnsName** | Pointer to **string** | The DNS name pointing to your cluster. | [optional] |
|**Instances** | Pointer to **int32** | The total number of instances in the cluster (one primary and n-1 secondary).  | [optional] |
|**Ram** | Pointer to **int32** | The amount of memory per instance in gigabytes (GB). | [optional] |
|**Cores** | Pointer to **int32** | The number of CPU cores per instance. | [optional] |
|**StorageSize** | Pointer to **int32** | The amount of storage per instance in gigabytes (GB). | [optional] |
|**Connections** | Pointer to [**[]Connection**](Connection.md) |  | [optional] |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |

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

### GetMariadbVersion

`func (o *ClusterProperties) GetMariadbVersion() MariadbVersion`

GetMariadbVersion returns the MariadbVersion field if non-nil, zero value otherwise.

### GetMariadbVersionOk

`func (o *ClusterProperties) GetMariadbVersionOk() (*MariadbVersion, bool)`

GetMariadbVersionOk returns a tuple with the MariadbVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMariadbVersion

`func (o *ClusterProperties) SetMariadbVersion(v MariadbVersion)`

SetMariadbVersion sets MariadbVersion field to given value.

### HasMariadbVersion

`func (o *ClusterProperties) HasMariadbVersion() bool`

HasMariadbVersion returns a boolean if a field has been set.

### GetDnsName

`func (o *ClusterProperties) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *ClusterProperties) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *ClusterProperties) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *ClusterProperties) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

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


