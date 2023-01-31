# CreateClusterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**TemplateID** | **string** | The unique ID of the template, which specifies the number of cores, storage size, and memory. You cannot downgrade to a smaller template or minor edition (e.g. from business to playground). To get a list of all templates to confirm the changes use the /templates endpoint.  | |
|**MongoDBVersion** | Pointer to **string** | The MongoDB version of your cluster. | [optional] |
|**Instances** | **int32** | The total number of instances in the cluster (one primary and n-1 secondaries).  | |
|**Connections** | [**[]Connection**](Connection.md) |  | |
|**Location** | **string** | The physical location where the cluster will be created. This is the location where all your instances will be located. This property is immutable.  | |
|**DisplayName** | **string** | The name of your cluster. | |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |

## Methods

### NewCreateClusterProperties

`func NewCreateClusterProperties(templateID string, instances int32, connections []Connection, location string, displayName string, ) *CreateClusterProperties`

NewCreateClusterProperties instantiates a new CreateClusterProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterPropertiesWithDefaults

`func NewCreateClusterPropertiesWithDefaults() *CreateClusterProperties`

NewCreateClusterPropertiesWithDefaults instantiates a new CreateClusterProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateID

`func (o *CreateClusterProperties) GetTemplateID() string`

GetTemplateID returns the TemplateID field if non-nil, zero value otherwise.

### GetTemplateIDOk

`func (o *CreateClusterProperties) GetTemplateIDOk() (*string, bool)`

GetTemplateIDOk returns a tuple with the TemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateID

`func (o *CreateClusterProperties) SetTemplateID(v string)`

SetTemplateID sets TemplateID field to given value.


### GetMongoDBVersion

`func (o *CreateClusterProperties) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *CreateClusterProperties) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *CreateClusterProperties) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *CreateClusterProperties) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.

### GetInstances

`func (o *CreateClusterProperties) GetInstances() int32`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *CreateClusterProperties) GetInstancesOk() (*int32, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *CreateClusterProperties) SetInstances(v int32)`

SetInstances sets Instances field to given value.


### GetConnections

`func (o *CreateClusterProperties) GetConnections() []Connection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *CreateClusterProperties) GetConnectionsOk() (*[]Connection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *CreateClusterProperties) SetConnections(v []Connection)`

SetConnections sets Connections field to given value.


### GetLocation

`func (o *CreateClusterProperties) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateClusterProperties) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateClusterProperties) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetDisplayName

`func (o *CreateClusterProperties) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateClusterProperties) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateClusterProperties) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


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


