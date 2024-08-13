# GatewayReadList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the Gateway. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Gateway. | |
|**Items** | Pointer to [**[]GatewayRead**](GatewayRead.md) | The list of Gateway resources. | [optional] |
|**Offset** | **int32** | The offset specified in the request (if none was specified, the default offset is 0).  | [readonly] |
|**Limit** | **int32** | The limit specified in the request (if none was specified, use the endpoint&#39;s default pagination limit).  | [readonly] |
|**Links** | [**Links**](Links.md) |  | |

## Methods

### NewGatewayReadList

`func NewGatewayReadList(id string, type_ string, href string, offset int32, limit int32, links Links, ) *GatewayReadList`

NewGatewayReadList instantiates a new GatewayReadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayReadListWithDefaults

`func NewGatewayReadListWithDefaults() *GatewayReadList`

NewGatewayReadListWithDefaults instantiates a new GatewayReadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayReadList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayReadList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayReadList) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *GatewayReadList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayReadList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayReadList) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *GatewayReadList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GatewayReadList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GatewayReadList) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *GatewayReadList) GetItems() []GatewayRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GatewayReadList) GetItemsOk() (*[]GatewayRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GatewayReadList) SetItems(v []GatewayRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *GatewayReadList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *GatewayReadList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GatewayReadList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GatewayReadList) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *GatewayReadList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GatewayReadList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GatewayReadList) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetLinks

`func (o *GatewayReadList) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GatewayReadList) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GatewayReadList) SetLinks(v Links)`

SetLinks sets Links field to given value.



