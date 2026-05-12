# Cluster

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | Name of the cluster | |
|**Connections** | [**[]ClusterConnections**](ClusterConnections.md) |  | |
|**Nfs** | Pointer to [**ClusterNfs**](ClusterNfs.md) |  | [optional] |
|**Size** | Pointer to **int32** | The size of the Network File Storage cluster in TiB or GiB. Note that the cluster size cannot be reduced after provisioning. This value determines the billing fees. - When &#x60;sizeUnit&#x60; is &#x60;TiB&#x60;, the allowed values are between 2 and 42. - When &#x60;sizeUnit&#x60; is &#x60;GiB&#x60;, the allowed values are between 2048 and 43008.  | [optional] [default to 2]|
|**SizeUnit** | Pointer to **string** | The type of the size for the Network File Storage cluster which can be TiB or GiB. | [optional] [default to "TiB"]|

## Methods

### NewCluster

`func NewCluster(name string, connections []ClusterConnections, ) *Cluster`

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


### GetConnections

`func (o *Cluster) GetConnections() []ClusterConnections`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *Cluster) GetConnectionsOk() (*[]ClusterConnections, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *Cluster) SetConnections(v []ClusterConnections)`

SetConnections sets Connections field to given value.


### GetNfs

`func (o *Cluster) GetNfs() ClusterNfs`

GetNfs returns the Nfs field if non-nil, zero value otherwise.

### GetNfsOk

`func (o *Cluster) GetNfsOk() (*ClusterNfs, bool)`

GetNfsOk returns a tuple with the Nfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfs

`func (o *Cluster) SetNfs(v ClusterNfs)`

SetNfs sets Nfs field to given value.

### HasNfs

`func (o *Cluster) HasNfs() bool`

HasNfs returns a boolean if a field has been set.

### GetSize

`func (o *Cluster) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Cluster) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Cluster) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Cluster) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeUnit

`func (o *Cluster) GetSizeUnit() string`

GetSizeUnit returns the SizeUnit field if non-nil, zero value otherwise.

### GetSizeUnitOk

`func (o *Cluster) GetSizeUnitOk() (*string, bool)`

GetSizeUnitOk returns a tuple with the SizeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeUnit

`func (o *Cluster) SetSizeUnit(v string)`

SetSizeUnit sets SizeUnit field to given value.

### HasSizeUnit

`func (o *Cluster) HasSizeUnit() bool`

HasSizeUnit returns a boolean if a field has been set.


