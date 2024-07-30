# Cluster

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of your cluster. Must be 63 characters or less and must begin and end with an alphanumeric character (&#x60;[a-z0-9A-Z]&#x60;) with dashes (&#x60;-&#x60;), underscores (&#x60;_&#x60;), dots (&#x60;.&#x60;), and alphanumerics between.  | [optional] |
|**DataPlatformVersion** | Pointer to **string** | The version of the data platform.  | [optional] |
|**DatacenterId** | Pointer to **string** | The UUID of the virtual data center (VDC) the cluster is provisioned.  | [optional] |
|**MaintenanceWindow** | Pointer to [**MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] |
|**Lans** | Pointer to [**[]Lan**](Lan.md) | A list of LANs you want this node pool to be part of.  | [optional] |

## Methods

### NewCluster

`func NewCluster() *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataPlatformVersion

`func (o *Cluster) GetDataPlatformVersion() string`

GetDataPlatformVersion returns the DataPlatformVersion field if non-nil, zero value otherwise.

### GetDataPlatformVersionOk

`func (o *Cluster) GetDataPlatformVersionOk() (*string, bool)`

GetDataPlatformVersionOk returns a tuple with the DataPlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPlatformVersion

`func (o *Cluster) SetDataPlatformVersion(v string)`

SetDataPlatformVersion sets DataPlatformVersion field to given value.

### HasDataPlatformVersion

`func (o *Cluster) HasDataPlatformVersion() bool`

HasDataPlatformVersion returns a boolean if a field has been set.

### GetDatacenterId

`func (o *Cluster) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *Cluster) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *Cluster) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *Cluster) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetMaintenanceWindow

`func (o *Cluster) GetMaintenanceWindow() MaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *Cluster) GetMaintenanceWindowOk() (*MaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *Cluster) SetMaintenanceWindow(v MaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *Cluster) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.

### GetLans

`func (o *Cluster) GetLans() []Lan`

GetLans returns the Lans field if non-nil, zero value otherwise.

### GetLansOk

`func (o *Cluster) GetLansOk() (*[]Lan, bool)`

GetLansOk returns a tuple with the Lans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLans

`func (o *Cluster) SetLans(v []Lan)`

SetLans sets Lans field to given value.

### HasLans

`func (o *Cluster) HasLans() bool`

HasLans returns a boolean if a field has been set.


