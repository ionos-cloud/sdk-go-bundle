# CreateNodePoolProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of your node pool. Must be 63 characters or less and must begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.  | |
|**NodeCount** | **int32** | The number of nodes that make up the node pool.  | |
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

### NewCreateNodePoolProperties

`func NewCreateNodePoolProperties(name string, nodeCount int32, ) *CreateNodePoolProperties`

NewCreateNodePoolProperties instantiates a new CreateNodePoolProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNodePoolPropertiesWithDefaults

`func NewCreateNodePoolPropertiesWithDefaults() *CreateNodePoolProperties`

NewCreateNodePoolPropertiesWithDefaults instantiates a new CreateNodePoolProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNodePoolProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNodePoolProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNodePoolProperties) SetName(v string)`

SetName sets Name field to given value.


### GetNodeCount

`func (o *CreateNodePoolProperties) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *CreateNodePoolProperties) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *CreateNodePoolProperties) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetCpuFamily

`func (o *CreateNodePoolProperties) GetCpuFamily() string`

GetCpuFamily returns the CpuFamily field if non-nil, zero value otherwise.

### GetCpuFamilyOk

`func (o *CreateNodePoolProperties) GetCpuFamilyOk() (*string, bool)`

GetCpuFamilyOk returns a tuple with the CpuFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFamily

`func (o *CreateNodePoolProperties) SetCpuFamily(v string)`

SetCpuFamily sets CpuFamily field to given value.

### HasCpuFamily

`func (o *CreateNodePoolProperties) HasCpuFamily() bool`

HasCpuFamily returns a boolean if a field has been set.

### GetCoresCount

`func (o *CreateNodePoolProperties) GetCoresCount() int32`

GetCoresCount returns the CoresCount field if non-nil, zero value otherwise.

### GetCoresCountOk

`func (o *CreateNodePoolProperties) GetCoresCountOk() (*int32, bool)`

GetCoresCountOk returns a tuple with the CoresCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresCount

`func (o *CreateNodePoolProperties) SetCoresCount(v int32)`

SetCoresCount sets CoresCount field to given value.

### HasCoresCount

`func (o *CreateNodePoolProperties) HasCoresCount() bool`

HasCoresCount returns a boolean if a field has been set.

### GetRamSize

`func (o *CreateNodePoolProperties) GetRamSize() int32`

GetRamSize returns the RamSize field if non-nil, zero value otherwise.

### GetRamSizeOk

`func (o *CreateNodePoolProperties) GetRamSizeOk() (*int32, bool)`

GetRamSizeOk returns a tuple with the RamSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamSize

`func (o *CreateNodePoolProperties) SetRamSize(v int32)`

SetRamSize sets RamSize field to given value.

### HasRamSize

`func (o *CreateNodePoolProperties) HasRamSize() bool`

HasRamSize returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *CreateNodePoolProperties) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateNodePoolProperties) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateNodePoolProperties) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *CreateNodePoolProperties) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetStorageType

`func (o *CreateNodePoolProperties) GetStorageType() StorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *CreateNodePoolProperties) GetStorageTypeOk() (*StorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *CreateNodePoolProperties) SetStorageType(v StorageType)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *CreateNodePoolProperties) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetStorageSize

`func (o *CreateNodePoolProperties) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *CreateNodePoolProperties) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *CreateNodePoolProperties) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *CreateNodePoolProperties) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetMaintenanceWindow

`func (o *CreateNodePoolProperties) GetMaintenanceWindow() MaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *CreateNodePoolProperties) GetMaintenanceWindowOk() (*MaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *CreateNodePoolProperties) SetMaintenanceWindow(v MaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *CreateNodePoolProperties) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.

### GetLabels

`func (o *CreateNodePoolProperties) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateNodePoolProperties) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateNodePoolProperties) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateNodePoolProperties) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateNodePoolProperties) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateNodePoolProperties) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateNodePoolProperties) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateNodePoolProperties) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


