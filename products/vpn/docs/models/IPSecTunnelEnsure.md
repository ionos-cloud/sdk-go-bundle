# IPSecTunnelEnsure

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the IPSecTunnel. | |
|**Metadata** | Pointer to **map[string]interface{}** | Metadata | [optional] |
|**Properties** | [**IPSecTunnel**](IPSecTunnel.md) |  | |

## Methods

### NewIPSecTunnelEnsure

`func NewIPSecTunnelEnsure(id string, properties IPSecTunnel, ) *IPSecTunnelEnsure`

NewIPSecTunnelEnsure instantiates a new IPSecTunnelEnsure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecTunnelEnsureWithDefaults

`func NewIPSecTunnelEnsureWithDefaults() *IPSecTunnelEnsure`

NewIPSecTunnelEnsureWithDefaults instantiates a new IPSecTunnelEnsure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPSecTunnelEnsure) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPSecTunnelEnsure) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPSecTunnelEnsure) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *IPSecTunnelEnsure) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IPSecTunnelEnsure) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IPSecTunnelEnsure) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IPSecTunnelEnsure) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *IPSecTunnelEnsure) GetProperties() IPSecTunnel`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IPSecTunnelEnsure) GetPropertiesOk() (*IPSecTunnel, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IPSecTunnelEnsure) SetProperties(v IPSecTunnel)`

SetProperties sets Properties field to given value.



