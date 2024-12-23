# PatchNodePoolProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**NodeCount** | Pointer to **int32** | The number of nodes that make up the node pool.  | [optional] |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |
|**Labels** | Pointer to **map[string]interface{}** | Key-value pairs attached to the node pool resource as [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/).  | [optional] |
|**Annotations** | Pointer to **map[string]interface{}** | Key-value pairs attached to node pool resource as [Kubernetes annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/).  | [optional] |
|**AutoScaling** | Pointer to [**NullableAutoScaling**](AutoScaling.md) |  | [optional] |

## Methods

### NewPatchNodePoolProperties

`func NewPatchNodePoolProperties() *PatchNodePoolProperties`

NewPatchNodePoolProperties instantiates a new PatchNodePoolProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchNodePoolPropertiesWithDefaults

`func NewPatchNodePoolPropertiesWithDefaults() *PatchNodePoolProperties`

NewPatchNodePoolPropertiesWithDefaults instantiates a new PatchNodePoolProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeCount

`func (o *PatchNodePoolProperties) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *PatchNodePoolProperties) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *PatchNodePoolProperties) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *PatchNodePoolProperties) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetMaintenanceWindow

`func (o *PatchNodePoolProperties) GetMaintenanceWindow() MaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *PatchNodePoolProperties) GetMaintenanceWindowOk() (*MaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *PatchNodePoolProperties) SetMaintenanceWindow(v MaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *PatchNodePoolProperties) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.

### GetLabels

`func (o *PatchNodePoolProperties) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PatchNodePoolProperties) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PatchNodePoolProperties) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PatchNodePoolProperties) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *PatchNodePoolProperties) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *PatchNodePoolProperties) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *PatchNodePoolProperties) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *PatchNodePoolProperties) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetAutoScaling

`func (o *PatchNodePoolProperties) GetAutoScaling() AutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *PatchNodePoolProperties) GetAutoScalingOk() (*AutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *PatchNodePoolProperties) SetAutoScaling(v AutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *PatchNodePoolProperties) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### SetAutoScalingNil

`func (o *PatchNodePoolProperties) SetAutoScalingNil(b bool)`

 SetAutoScalingNil sets the value for AutoScaling to be an explicit nil

### UnsetAutoScaling
`func (o *PatchNodePoolProperties) UnsetAutoScaling()`

UnsetAutoScaling ensures that no value is present for AutoScaling, not even an explicit nil

