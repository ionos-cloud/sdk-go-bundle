# ClusterConnections

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DatacenterId** | **string** | The ID of the datacenter where the Network File Storage cluster is located.  | |
|**Lan** | **string** | The LAN to which the Network File Storage cluster must be connected.  | |
|**IpAddress** | **string** | The IP address and prefix of the Network File Storage cluster. The IP address can be either IPv4 or IPv6. The IP address has to be given with CIDR notation.  | |

## Methods

### NewClusterConnections

`func NewClusterConnections(datacenterId string, lan string, ipAddress string, ) *ClusterConnections`

NewClusterConnections instantiates a new ClusterConnections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConnectionsWithDefaults

`func NewClusterConnectionsWithDefaults() *ClusterConnections`

NewClusterConnectionsWithDefaults instantiates a new ClusterConnections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterId

`func (o *ClusterConnections) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *ClusterConnections) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *ClusterConnections) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetLan

`func (o *ClusterConnections) GetLan() string`

GetLan returns the Lan field if non-nil, zero value otherwise.

### GetLanOk

`func (o *ClusterConnections) GetLanOk() (*string, bool)`

GetLanOk returns a tuple with the Lan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan

`func (o *ClusterConnections) SetLan(v string)`

SetLan sets Lan field to given value.


### GetIpAddress

`func (o *ClusterConnections) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ClusterConnections) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ClusterConnections) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.



