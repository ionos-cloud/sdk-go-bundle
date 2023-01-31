# PatchClusterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DisplayName** | Pointer to **string** | The name of your cluster. | [optional] |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |
|**Instances** | Pointer to **int32** | The total number of instances in the cluster (one primary and n-1 secondaries).  | [optional] |
|**Connections** | Pointer to [**[]Connection**](Connection.md) |  | [optional] |
|**TemplateID** | Pointer to **string** | The unique ID of the template, which specifies the number of cores, storage size, and memory. You cannot downgrade to a smaller template or minor edition (e.g. from business to playground). To get a list of all templates to confirm the changes use the /templates endpoint.  | [optional] |

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

### GetDisplayName

`func (o *PatchClusterProperties) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PatchClusterProperties) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PatchClusterProperties) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PatchClusterProperties) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

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

### GetInstances

`func (o *PatchClusterProperties) GetInstances() int32`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *PatchClusterProperties) GetInstancesOk() (*int32, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *PatchClusterProperties) SetInstances(v int32)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *PatchClusterProperties) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetConnections

`func (o *PatchClusterProperties) GetConnections() []Connection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *PatchClusterProperties) GetConnectionsOk() (*[]Connection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *PatchClusterProperties) SetConnections(v []Connection)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *PatchClusterProperties) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetTemplateID

`func (o *PatchClusterProperties) GetTemplateID() string`

GetTemplateID returns the TemplateID field if non-nil, zero value otherwise.

### GetTemplateIDOk

`func (o *PatchClusterProperties) GetTemplateIDOk() (*string, bool)`

GetTemplateIDOk returns a tuple with the TemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateID

`func (o *PatchClusterProperties) SetTemplateID(v string)`

SetTemplateID sets TemplateID field to given value.

### HasTemplateID

`func (o *PatchClusterProperties) HasTemplateID() bool`

HasTemplateID returns a boolean if a field has been set.


