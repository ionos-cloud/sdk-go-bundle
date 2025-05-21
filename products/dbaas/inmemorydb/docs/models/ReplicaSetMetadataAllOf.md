# ReplicaSetMetadataAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DnsName** | **string** | The DNS name pointing to your replica set. Will be used to connect to the active/standalone instance.  | |

## Methods

### NewReplicaSetMetadataAllOf

`func NewReplicaSetMetadataAllOf(dnsName string, ) *ReplicaSetMetadataAllOf`

NewReplicaSetMetadataAllOf instantiates a new ReplicaSetMetadataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaSetMetadataAllOfWithDefaults

`func NewReplicaSetMetadataAllOfWithDefaults() *ReplicaSetMetadataAllOf`

NewReplicaSetMetadataAllOfWithDefaults instantiates a new ReplicaSetMetadataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsName

`func (o *ReplicaSetMetadataAllOf) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *ReplicaSetMetadataAllOf) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *ReplicaSetMetadataAllOf) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.



