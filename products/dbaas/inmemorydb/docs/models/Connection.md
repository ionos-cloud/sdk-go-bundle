# Connection

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DatacenterId** | **string** | The datacenter to connect your instance to. | |
|**LanId** | **string** | The numeric LAN ID to connect your instance to. | |
|**Cidr** | **string** | The IP and subnet for your instance. Note the following unavailable IP ranges: 10.210.0.0/16 10.212.0.0/14  | |

## Methods

### NewConnection

`func NewConnection(datacenterId string, lanId string, cidr string, ) *Connection`

NewConnection instantiates a new Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionWithDefaults

`func NewConnectionWithDefaults() *Connection`

NewConnectionWithDefaults instantiates a new Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterId

`func (o *Connection) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *Connection) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *Connection) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetLanId

`func (o *Connection) GetLanId() string`

GetLanId returns the LanId field if non-nil, zero value otherwise.

### GetLanIdOk

`func (o *Connection) GetLanIdOk() (*string, bool)`

GetLanIdOk returns a tuple with the LanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanId

`func (o *Connection) SetLanId(v string)`

SetLanId sets LanId field to given value.


### GetCidr

`func (o *Connection) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *Connection) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *Connection) SetCidr(v string)`

SetCidr sets Cidr field to given value.



