# WireguardGatewayEnsure

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the WireguardGateway. | |
|**Metadata** | Pointer to **map[string]interface{}** | Metadata | [optional] |
|**Properties** | [**WireguardGateway**](WireguardGateway.md) |  | |

## Methods

### NewWireguardGatewayEnsure

`func NewWireguardGatewayEnsure(id string, properties WireguardGateway, ) *WireguardGatewayEnsure`

NewWireguardGatewayEnsure instantiates a new WireguardGatewayEnsure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardGatewayEnsureWithDefaults

`func NewWireguardGatewayEnsureWithDefaults() *WireguardGatewayEnsure`

NewWireguardGatewayEnsureWithDefaults instantiates a new WireguardGatewayEnsure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WireguardGatewayEnsure) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireguardGatewayEnsure) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireguardGatewayEnsure) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *WireguardGatewayEnsure) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WireguardGatewayEnsure) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WireguardGatewayEnsure) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WireguardGatewayEnsure) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *WireguardGatewayEnsure) GetProperties() WireguardGateway`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WireguardGatewayEnsure) GetPropertiesOk() (*WireguardGateway, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WireguardGatewayEnsure) SetProperties(v WireguardGateway)`

SetProperties sets Properties field to given value.



