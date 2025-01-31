# KafkaClusterConnection

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DatacenterId** | **string** | The datacenter to connect your instance to. | |
|**LanId** | **string** | The numeric LAN ID to connect your instance to. | |
|**BrokerAddresses** | **[]string** | IP addresses and subnet of cluster brokers. Note the following unavailable IP range: 10.224.0.0/11  | |

## Methods

### NewKafkaClusterConnection

`func NewKafkaClusterConnection(datacenterId string, lanId string, brokerAddresses []string, ) *KafkaClusterConnection`

NewKafkaClusterConnection instantiates a new KafkaClusterConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaClusterConnectionWithDefaults

`func NewKafkaClusterConnectionWithDefaults() *KafkaClusterConnection`

NewKafkaClusterConnectionWithDefaults instantiates a new KafkaClusterConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterId

`func (o *KafkaClusterConnection) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *KafkaClusterConnection) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *KafkaClusterConnection) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetLanId

`func (o *KafkaClusterConnection) GetLanId() string`

GetLanId returns the LanId field if non-nil, zero value otherwise.

### GetLanIdOk

`func (o *KafkaClusterConnection) GetLanIdOk() (*string, bool)`

GetLanIdOk returns a tuple with the LanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanId

`func (o *KafkaClusterConnection) SetLanId(v string)`

SetLanId sets LanId field to given value.


### GetBrokerAddresses

`func (o *KafkaClusterConnection) GetBrokerAddresses() []string`

GetBrokerAddresses returns the BrokerAddresses field if non-nil, zero value otherwise.

### GetBrokerAddressesOk

`func (o *KafkaClusterConnection) GetBrokerAddressesOk() (*[]string, bool)`

GetBrokerAddressesOk returns a tuple with the BrokerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerAddresses

`func (o *KafkaClusterConnection) SetBrokerAddresses(v []string)`

SetBrokerAddresses sets BrokerAddresses field to given value.



