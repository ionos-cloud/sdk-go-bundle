# IPSecGatewayRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the IPSecGateway. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the IPSecGateway. | |
|**Metadata** | [**IPSecGatewayMetadata**](IPSecGatewayMetadata.md) |  | |
|**Properties** | [**IPSecGateway**](IPSecGateway.md) |  | |

## Methods

### NewIPSecGatewayRead

`func NewIPSecGatewayRead(id string, type_ string, href string, metadata IPSecGatewayMetadata, properties IPSecGateway, ) *IPSecGatewayRead`

NewIPSecGatewayRead instantiates a new IPSecGatewayRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecGatewayReadWithDefaults

`func NewIPSecGatewayReadWithDefaults() *IPSecGatewayRead`

NewIPSecGatewayReadWithDefaults instantiates a new IPSecGatewayRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPSecGatewayRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPSecGatewayRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPSecGatewayRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *IPSecGatewayRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPSecGatewayRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPSecGatewayRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *IPSecGatewayRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IPSecGatewayRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IPSecGatewayRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *IPSecGatewayRead) GetMetadata() IPSecGatewayMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IPSecGatewayRead) GetMetadataOk() (*IPSecGatewayMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IPSecGatewayRead) SetMetadata(v IPSecGatewayMetadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *IPSecGatewayRead) GetProperties() IPSecGateway`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IPSecGatewayRead) GetPropertiesOk() (*IPSecGateway, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IPSecGatewayRead) SetProperties(v IPSecGateway)`

SetProperties sets Properties field to given value.



