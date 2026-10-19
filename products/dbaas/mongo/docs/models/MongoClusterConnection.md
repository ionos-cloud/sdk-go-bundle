# MongoClusterConnection

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DatacenterId** | **string** | The datacenter to connect your instance to. Immutable. Cannot be changed after cluster creation. | |
|**LanId** | **string** | The numeric LAN ID to connect your instance to. Immutable. Cannot be changed after cluster creation. | |
|**InstanceAddresses** | **[]string** | A list of IP addresses assigned to the cluster instances. The number of IP addresses must match the number of instances in the cluster.  | |

## Methods

### NewMongoClusterConnection

`func NewMongoClusterConnection(datacenterId string, lanId string, instanceAddresses []string, ) *MongoClusterConnection`

NewMongoClusterConnection instantiates a new MongoClusterConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoClusterConnectionWithDefaults

`func NewMongoClusterConnectionWithDefaults() *MongoClusterConnection`

NewMongoClusterConnectionWithDefaults instantiates a new MongoClusterConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterId

`func (o *MongoClusterConnection) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *MongoClusterConnection) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *MongoClusterConnection) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetLanId

`func (o *MongoClusterConnection) GetLanId() string`

GetLanId returns the LanId field if non-nil, zero value otherwise.

### GetLanIdOk

`func (o *MongoClusterConnection) GetLanIdOk() (*string, bool)`

GetLanIdOk returns a tuple with the LanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanId

`func (o *MongoClusterConnection) SetLanId(v string)`

SetLanId sets LanId field to given value.


### GetInstanceAddresses

`func (o *MongoClusterConnection) GetInstanceAddresses() []string`

GetInstanceAddresses returns the InstanceAddresses field if non-nil, zero value otherwise.

### GetInstanceAddressesOk

`func (o *MongoClusterConnection) GetInstanceAddressesOk() (*[]string, bool)`

GetInstanceAddressesOk returns a tuple with the InstanceAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAddresses

`func (o *MongoClusterConnection) SetInstanceAddresses(v []string)`

SetInstanceAddresses sets InstanceAddresses field to given value.



