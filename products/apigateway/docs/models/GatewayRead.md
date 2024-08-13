# GatewayRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the Gateway. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Gateway. | |
|**Metadata** | [**MetadataWithEndpoint**](MetadataWithEndpoint.md) |  | |
|**Properties** | [**Gateway**](Gateway.md) |  | |

## Methods

### NewGatewayRead

`func NewGatewayRead(id string, type_ string, href string, metadata MetadataWithEndpoint, properties Gateway, ) *GatewayRead`

NewGatewayRead instantiates a new GatewayRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayReadWithDefaults

`func NewGatewayReadWithDefaults() *GatewayRead`

NewGatewayReadWithDefaults instantiates a new GatewayRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *GatewayRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *GatewayRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GatewayRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GatewayRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *GatewayRead) GetMetadata() MetadataWithEndpoint`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GatewayRead) GetMetadataOk() (*MetadataWithEndpoint, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GatewayRead) SetMetadata(v MetadataWithEndpoint)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *GatewayRead) GetProperties() Gateway`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GatewayRead) GetPropertiesOk() (*Gateway, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GatewayRead) SetProperties(v Gateway)`

SetProperties sets Properties field to given value.



