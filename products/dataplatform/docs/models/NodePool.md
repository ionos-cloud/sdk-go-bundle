# NodePool

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of your node pool. Must be 63 characters or less and must begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.  | [optional] |
|**DataPlatformVersion** | Pointer to **string** | The version of the data platform.  | [optional] |
|**DatacenterId** | Pointer to **string** | The UUID of the virtual data center (VDC) the cluster is provisioned.  | [optional] |
|**NodeCount** | Pointer to **int32** | The number of nodes that make up the node pool.  | [optional] |
|**CpuFamily** | Pointer to **string** | A valid CPU family name or &#x60;AUTO&#x60; if the platform shall choose the best fitting option. Available CPU architectures can be retrieved from the data center resource.  | [optional] [default to "AUTO"]|
|**CoresCount** | Pointer to **int32** | The number of CPU cores per node.  | [optional] [default to 4]|
|**RamSize** | Pointer to **int32** | The RAM size for one node in MB. Must be set in multiples of 1024 MB, with a minimum size is of 2048 MB. | [optional] [default to 4096]|
|**AvailabilityZone** | Pointer to [**AvailabilityZone**](AvailabilityZone.md) |  | [optional] [default to AVAILABILITYZONE_AUTO]|
|**StorageType** | Pointer to [**StorageType**](StorageType.md) |  | [optional] [default to STORAGETYPE_SSD]|
|**StorageSize** | Pointer to **int32** | The size of the volume in GB. The size must be greater than 10 GB. | [optional] [default to 20]|
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |
|**Labels** | Pointer to **map[string]interface{}** | Key-value pairs attached to the node pool resource as [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)  | [optional] |
|**Annotations** | Pointer to **map[string]interface{}** | Key-value pairs attached to node pool resource as [Kubernetes annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/)  | [optional] |

## Methods

### NewNodePool

`func NewNodePool() *NodePool`

NewNodePool instantiates a new NodePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePoolWithDefaults

`func NewNodePoolWithDefaults() *NodePool`

NewNodePoolWithDefaults instantiates a new NodePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodePool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodePool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodePool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodePool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataPlatformVersion

`func (o *NodePool) GetDataPlatformVersion() string`

GetDataPlatformVersion returns the DataPlatformVersion field if non-nil, zero value otherwise.

### GetDataPlatformVersionOk

`func (o *NodePool) GetDataPlatformVersionOk() (*string, bool)`

GetDataPlatformVersionOk returns a tuple with the DataPlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPlatformVersion

`func (o *NodePool) SetDataPlatformVersion(v string)`

SetDataPlatformVersion sets DataPlatformVersion field to given value.

### HasDataPlatformVersion

`func (o *NodePool) HasDataPlatformVersion() bool`

HasDataPlatformVersion returns a boolean if a field has been set.

### GetDatacenterId

`func (o *NodePool) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *NodePool) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *NodePool) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *NodePool) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetNodeCount

`func (o *NodePool) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *NodePool) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *NodePool) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *NodePool) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCpuFamily

`func (o *NodePool) GetCpuFamily() string`

GetCpuFamily returns the CpuFamily field if non-nil, zero value otherwise.

### GetCpuFamilyOk

`func (o *NodePool) GetCpuFamilyOk() (*string, bool)`

GetCpuFamilyOk returns a tuple with the CpuFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFamily

`func (o *NodePool) SetCpuFamily(v string)`

SetCpuFamily sets CpuFamily field to given value.

### HasCpuFamily

`func (o *NodePool) HasCpuFamily() bool`

HasCpuFamily returns a boolean if a field has been set.

### GetCoresCount

`func (o *NodePool) GetCoresCount() int32`

GetCoresCount returns the CoresCount field if non-nil, zero value otherwise.

### GetCoresCountOk

`func (o *NodePool) GetCoresCountOk() (*int32, bool)`

GetCoresCountOk returns a tuple with the CoresCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresCount

`func (o *NodePool) SetCoresCount(v int32)`

SetCoresCount sets CoresCount field to given value.

### HasCoresCount

`func (o *NodePool) HasCoresCount() bool`

HasCoresCount returns a boolean if a field has been set.

### GetRamSize

`func (o *NodePool) GetRamSize() int32`

GetRamSize returns the RamSize field if non-nil, zero value otherwise.

### GetRamSizeOk

`func (o *NodePool) GetRamSizeOk() (*int32, bool)`

GetRamSizeOk returns a tuple with the RamSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamSize

`func (o *NodePool) SetRamSize(v int32)`

SetRamSize sets RamSize field to given value.

### HasRamSize

`func (o *NodePool) HasRamSize() bool`

HasRamSize returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *NodePool) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *NodePool) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *NodePool) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *NodePool) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetStorageType

`func (o *NodePool) GetStorageType() StorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *NodePool) GetStorageTypeOk() (*StorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *NodePool) SetStorageType(v StorageType)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *NodePool) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetStorageSize

`func (o *NodePool) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *NodePool) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *NodePool) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *NodePool) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetMaintenanceWindow

`func (o *NodePool) GetMaintenanceWindow() MaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *NodePool) GetMaintenanceWindowOk() (*MaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *NodePool) SetMaintenanceWindow(v MaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *NodePool) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.

### GetLabels

`func (o *NodePool) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NodePool) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NodePool) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NodePool) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *NodePool) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *NodePool) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *NodePool) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *NodePool) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


