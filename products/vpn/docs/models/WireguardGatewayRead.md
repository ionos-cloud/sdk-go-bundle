# WireguardGatewayRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the WireguardGateway. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the WireguardGateway. | |
|**Metadata** | [**WireguardGatewayMetadata**](WireguardGatewayMetadata.md) |  | |
|**Properties** | [**WireguardGateway**](WireguardGateway.md) |  | |

## Methods

### NewWireguardGatewayRead

`func NewWireguardGatewayRead(id string, type_ string, href string, metadata WireguardGatewayMetadata, properties WireguardGateway, ) *WireguardGatewayRead`

NewWireguardGatewayRead instantiates a new WireguardGatewayRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardGatewayReadWithDefaults

`func NewWireguardGatewayReadWithDefaults() *WireguardGatewayRead`

NewWireguardGatewayReadWithDefaults instantiates a new WireguardGatewayRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WireguardGatewayRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireguardGatewayRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireguardGatewayRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *WireguardGatewayRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WireguardGatewayRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WireguardGatewayRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *WireguardGatewayRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *WireguardGatewayRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *WireguardGatewayRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *WireguardGatewayRead) GetMetadata() WireguardGatewayMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WireguardGatewayRead) GetMetadataOk() (*WireguardGatewayMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WireguardGatewayRead) SetMetadata(v WireguardGatewayMetadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *WireguardGatewayRead) GetProperties() WireguardGateway`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WireguardGatewayRead) GetPropertiesOk() (*WireguardGateway, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WireguardGatewayRead) SetProperties(v WireguardGateway)`

SetProperties sets Properties field to given value.



