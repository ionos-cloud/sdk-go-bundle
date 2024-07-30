# IPSecTunnelRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the IPSecTunnel. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the IPSecTunnel. | |
|**Metadata** | [**IPSecTunnelMetadata**](IPSecTunnelMetadata.md) |  | |
|**Properties** | [**IPSecTunnel**](IPSecTunnel.md) |  | |

## Methods

### NewIPSecTunnelRead

`func NewIPSecTunnelRead(id string, type_ string, href string, metadata IPSecTunnelMetadata, properties IPSecTunnel, ) *IPSecTunnelRead`

NewIPSecTunnelRead instantiates a new IPSecTunnelRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecTunnelReadWithDefaults

`func NewIPSecTunnelReadWithDefaults() *IPSecTunnelRead`

NewIPSecTunnelReadWithDefaults instantiates a new IPSecTunnelRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPSecTunnelRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPSecTunnelRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPSecTunnelRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *IPSecTunnelRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPSecTunnelRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPSecTunnelRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *IPSecTunnelRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IPSecTunnelRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IPSecTunnelRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *IPSecTunnelRead) GetMetadata() IPSecTunnelMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IPSecTunnelRead) GetMetadataOk() (*IPSecTunnelMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IPSecTunnelRead) SetMetadata(v IPSecTunnelMetadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *IPSecTunnelRead) GetProperties() IPSecTunnel`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IPSecTunnelRead) GetPropertiesOk() (*IPSecTunnel, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IPSecTunnelRead) SetProperties(v IPSecTunnel)`

SetProperties sets Properties field to given value.



