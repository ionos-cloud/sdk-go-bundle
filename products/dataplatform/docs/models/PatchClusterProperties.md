# PatchClusterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of your cluster. Must be 63 characters or less and must begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.  | [optional] |
|**DataPlatformVersion** | Pointer to **string** | The version of the data platform.  | [optional] |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |

## Methods

### NewPatchClusterProperties

`func NewPatchClusterProperties() *PatchClusterProperties`

NewPatchClusterProperties instantiates a new PatchClusterProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchClusterPropertiesWithDefaults

`func NewPatchClusterPropertiesWithDefaults() *PatchClusterProperties`

NewPatchClusterPropertiesWithDefaults instantiates a new PatchClusterProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchClusterProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchClusterProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchClusterProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchClusterProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataPlatformVersion

`func (o *PatchClusterProperties) GetDataPlatformVersion() string`

GetDataPlatformVersion returns the DataPlatformVersion field if non-nil, zero value otherwise.

### GetDataPlatformVersionOk

`func (o *PatchClusterProperties) GetDataPlatformVersionOk() (*string, bool)`

GetDataPlatformVersionOk returns a tuple with the DataPlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPlatformVersion

`func (o *PatchClusterProperties) SetDataPlatformVersion(v string)`

SetDataPlatformVersion sets DataPlatformVersion field to given value.

### HasDataPlatformVersion

`func (o *PatchClusterProperties) HasDataPlatformVersion() bool`

HasDataPlatformVersion returns a boolean if a field has been set.

### GetMaintenanceWindow

`func (o *PatchClusterProperties) GetMaintenanceWindow() MaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *PatchClusterProperties) GetMaintenanceWindowOk() (*MaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *PatchClusterProperties) SetMaintenanceWindow(v MaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *PatchClusterProperties) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.


