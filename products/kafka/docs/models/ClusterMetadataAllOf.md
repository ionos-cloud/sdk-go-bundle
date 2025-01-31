# ClusterMetadataAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**BrokerAddresses** | Pointer to **[]string** | IP addresses and ports of cluster brokers. | [optional] |

## Methods

### NewClusterMetadataAllOf

`func NewClusterMetadataAllOf() *ClusterMetadataAllOf`

NewClusterMetadataAllOf instantiates a new ClusterMetadataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterMetadataAllOfWithDefaults

`func NewClusterMetadataAllOfWithDefaults() *ClusterMetadataAllOf`

NewClusterMetadataAllOfWithDefaults instantiates a new ClusterMetadataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokerAddresses

`func (o *ClusterMetadataAllOf) GetBrokerAddresses() []string`

GetBrokerAddresses returns the BrokerAddresses field if non-nil, zero value otherwise.

### GetBrokerAddressesOk

`func (o *ClusterMetadataAllOf) GetBrokerAddressesOk() (*[]string, bool)`

GetBrokerAddressesOk returns a tuple with the BrokerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerAddresses

`func (o *ClusterMetadataAllOf) SetBrokerAddresses(v []string)`

SetBrokerAddresses sets BrokerAddresses field to given value.

### HasBrokerAddresses

`func (o *ClusterMetadataAllOf) HasBrokerAddresses() bool`

HasBrokerAddresses returns a boolean if a field has been set.


