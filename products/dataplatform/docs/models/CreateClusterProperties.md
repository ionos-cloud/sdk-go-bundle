# CreateClusterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of your cluster. Must be 63 characters or less and must begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.  | |
|**DataPlatformVersion** | Pointer to **string** | The version of the data platform.  | [optional] |
|**DatacenterId** | **string** | The UUID of the virtual data center (VDC) the cluster is provisioned.  | |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |

## Methods

### NewCreateClusterProperties

`func NewCreateClusterProperties(name string, datacenterId string, ) *CreateClusterProperties`

NewCreateClusterProperties instantiates a new CreateClusterProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterPropertiesWithDefaults

`func NewCreateClusterPropertiesWithDefaults() *CreateClusterProperties`

NewCreateClusterPropertiesWithDefaults instantiates a new CreateClusterProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateClusterProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClusterProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClusterProperties) SetName(v string)`

SetName sets Name field to given value.


### GetDataPlatformVersion

`func (o *CreateClusterProperties) GetDataPlatformVersion() string`

GetDataPlatformVersion returns the DataPlatformVersion field if non-nil, zero value otherwise.

### GetDataPlatformVersionOk

`func (o *CreateClusterProperties) GetDataPlatformVersionOk() (*string, bool)`

GetDataPlatformVersionOk returns a tuple with the DataPlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPlatformVersion

`func (o *CreateClusterProperties) SetDataPlatformVersion(v string)`

SetDataPlatformVersion sets DataPlatformVersion field to given value.

### HasDataPlatformVersion

`func (o *CreateClusterProperties) HasDataPlatformVersion() bool`

HasDataPlatformVersion returns a boolean if a field has been set.

### GetDatacenterId

`func (o *CreateClusterProperties) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *CreateClusterProperties) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *CreateClusterProperties) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetMaintenanceWindow

`func (o *CreateClusterProperties) GetMaintenanceWindow() MaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *CreateClusterProperties) GetMaintenanceWindowOk() (*MaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *CreateClusterProperties) SetMaintenanceWindow(v MaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *CreateClusterProperties) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.


