# ClusterConnection

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DatacenterId** | **string** | The ID of the Virtual Data Center to connect the cluster to. | |
|**LanId** | **string** | The numeric LAN ID within the data center to connect the cluster to. | |
|**PrimaryInstanceAddress** | **string** | The IP address and subnet mask assigned to the primary instance in CIDR notation. Note the following unavailable IP range: 10.210.0.0/16, 10.212.0.0/14.  | |

## Methods

### NewClusterConnection

`func NewClusterConnection(datacenterId string, lanId string, primaryInstanceAddress string, ) *ClusterConnection`

NewClusterConnection instantiates a new ClusterConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConnectionWithDefaults

`func NewClusterConnectionWithDefaults() *ClusterConnection`

NewClusterConnectionWithDefaults instantiates a new ClusterConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterId

`func (o *ClusterConnection) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *ClusterConnection) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *ClusterConnection) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetLanId

`func (o *ClusterConnection) GetLanId() string`

GetLanId returns the LanId field if non-nil, zero value otherwise.

### GetLanIdOk

`func (o *ClusterConnection) GetLanIdOk() (*string, bool)`

GetLanIdOk returns a tuple with the LanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanId

`func (o *ClusterConnection) SetLanId(v string)`

SetLanId sets LanId field to given value.


### GetPrimaryInstanceAddress

`func (o *ClusterConnection) GetPrimaryInstanceAddress() string`

GetPrimaryInstanceAddress returns the PrimaryInstanceAddress field if non-nil, zero value otherwise.

### GetPrimaryInstanceAddressOk

`func (o *ClusterConnection) GetPrimaryInstanceAddressOk() (*string, bool)`

GetPrimaryInstanceAddressOk returns a tuple with the PrimaryInstanceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInstanceAddress

`func (o *ClusterConnection) SetPrimaryInstanceAddress(v string)`

SetPrimaryInstanceAddress sets PrimaryInstanceAddress field to given value.



