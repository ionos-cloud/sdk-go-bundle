# MariadbClusterConnection

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DatacenterId** | **string** | The data center to connect your instance to. | |
|**LanId** | **string** | The numeric LAN ID to connect your instance to. | |
|**PrimaryInstanceAddress** | **string** | Assigns the IP address and netmask to the cluster&#39;s primary instance, in CIDR notation. Note the following unavailable IP ranges: 10.208.0.0/12 10.233.0.0/18 10.233.64.0/18 192.168.230.0/24  | |

## Methods

### NewMariadbClusterConnection

`func NewMariadbClusterConnection(datacenterId string, lanId string, primaryInstanceAddress string, ) *MariadbClusterConnection`

NewMariadbClusterConnection instantiates a new MariadbClusterConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMariadbClusterConnectionWithDefaults

`func NewMariadbClusterConnectionWithDefaults() *MariadbClusterConnection`

NewMariadbClusterConnectionWithDefaults instantiates a new MariadbClusterConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterId

`func (o *MariadbClusterConnection) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *MariadbClusterConnection) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *MariadbClusterConnection) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetLanId

`func (o *MariadbClusterConnection) GetLanId() string`

GetLanId returns the LanId field if non-nil, zero value otherwise.

### GetLanIdOk

`func (o *MariadbClusterConnection) GetLanIdOk() (*string, bool)`

GetLanIdOk returns a tuple with the LanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanId

`func (o *MariadbClusterConnection) SetLanId(v string)`

SetLanId sets LanId field to given value.


### GetPrimaryInstanceAddress

`func (o *MariadbClusterConnection) GetPrimaryInstanceAddress() string`

GetPrimaryInstanceAddress returns the PrimaryInstanceAddress field if non-nil, zero value otherwise.

### GetPrimaryInstanceAddressOk

`func (o *MariadbClusterConnection) GetPrimaryInstanceAddressOk() (*string, bool)`

GetPrimaryInstanceAddressOk returns a tuple with the PrimaryInstanceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInstanceAddress

`func (o *MariadbClusterConnection) SetPrimaryInstanceAddress(v string)`

SetPrimaryInstanceAddress sets PrimaryInstanceAddress field to given value.



