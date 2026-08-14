# ClusterMetadataAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**State** | Pointer to [**ClusterState**](ClusterState.md) |  | [optional] |
|**StatusMessage** | Pointer to **string** | A human-readable message describing the current state. Populated when state is FAILED. | [optional] [readonly] |
|**DnsName** | Pointer to **string** | The DNS name used to connect to the cluster&#39;s primary instance. Pattern: &lt;clusterId&gt;.in-memory-db.&lt;location&gt;.ionos.com  | [optional] [readonly] |

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

### GetState

`func (o *ClusterMetadataAllOf) GetState() ClusterState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClusterMetadataAllOf) GetStateOk() (*ClusterState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClusterMetadataAllOf) SetState(v ClusterState)`

SetState sets State field to given value.

### HasState

`func (o *ClusterMetadataAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ClusterMetadataAllOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ClusterMetadataAllOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ClusterMetadataAllOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ClusterMetadataAllOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetDnsName

`func (o *ClusterMetadataAllOf) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *ClusterMetadataAllOf) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *ClusterMetadataAllOf) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *ClusterMetadataAllOf) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.


