# ClusterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DisplayName** | Pointer to **string** | The name of your cluster. | [optional] |
|**MongoDBVersion** | Pointer to **string** | The MongoDB version of your cluster. | [optional] |
|**Location** | Pointer to **string** | The physical location where the cluster will be created. This is the location where all your instances will be located. This property is immutable.  | [optional] |
|**Instances** | Pointer to **int32** | The total number of instances in the cluster (one primary and n-1 secondaries).  | [optional] |
|**Connections** | Pointer to [**[]Connection**](Connection.md) |  | [optional] |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |
|**TemplateID** | Pointer to **string** | The unique ID of the template, which specifies the number of cores, storage size, and memory. You cannot downgrade to a smaller template or minor edition (e.g. from business to playground). To get a list of all templates to confirm the changes use the /templates endpoint.  | [optional] |
|**ConnectionString** | Pointer to **string** | The connection string for your cluster. | [optional] |

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

### GetMongoDBVersion

`func (o *ClusterProperties) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *ClusterProperties) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *ClusterProperties) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *ClusterProperties) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.

### GetLocation

`func (o *ClusterProperties) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ClusterProperties) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ClusterProperties) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ClusterProperties) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

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

### GetTemplateID

`func (o *ClusterProperties) GetTemplateID() string`

GetTemplateID returns the TemplateID field if non-nil, zero value otherwise.

### GetTemplateIDOk

`func (o *ClusterProperties) GetTemplateIDOk() (*string, bool)`

GetTemplateIDOk returns a tuple with the TemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateID

`func (o *ClusterProperties) SetTemplateID(v string)`

SetTemplateID sets TemplateID field to given value.

### HasTemplateID

`func (o *ClusterProperties) HasTemplateID() bool`

HasTemplateID returns a boolean if a field has been set.

### GetConnectionString

`func (o *ClusterProperties) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *ClusterProperties) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *ClusterProperties) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.

### HasConnectionString

`func (o *ClusterProperties) HasConnectionString() bool`

HasConnectionString returns a boolean if a field has been set.


