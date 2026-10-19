# MongoVersionReadList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of MongoVersion resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of MongoVersion resources. | |
|**Items** | Pointer to [**[]MongoVersionRead**](MongoVersionRead.md) | The list of MongoVersion resources. | [optional] |
|**Offset** | **int32** | The offset specified in the request (if none was specified, the default offset is 0).  | [readonly] |
|**Limit** | **int32** | The limit specified in the request (if none was specified, use the endpoint&#39;s default pagination limit).  | [readonly] |
|**Links** | [**Links**](Links.md) |  | |

## Methods

### NewMongoVersionReadList

`func NewMongoVersionReadList(id string, type_ string, href string, offset int32, limit int32, links Links, ) *MongoVersionReadList`

NewMongoVersionReadList instantiates a new MongoVersionReadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoVersionReadListWithDefaults

`func NewMongoVersionReadListWithDefaults() *MongoVersionReadList`

NewMongoVersionReadListWithDefaults instantiates a new MongoVersionReadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MongoVersionReadList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MongoVersionReadList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MongoVersionReadList) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *MongoVersionReadList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MongoVersionReadList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MongoVersionReadList) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *MongoVersionReadList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MongoVersionReadList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MongoVersionReadList) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *MongoVersionReadList) GetItems() []MongoVersionRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MongoVersionReadList) GetItemsOk() (*[]MongoVersionRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MongoVersionReadList) SetItems(v []MongoVersionRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *MongoVersionReadList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *MongoVersionReadList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *MongoVersionReadList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *MongoVersionReadList) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *MongoVersionReadList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MongoVersionReadList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MongoVersionReadList) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetLinks

`func (o *MongoVersionReadList) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MongoVersionReadList) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MongoVersionReadList) SetLinks(v Links)`

SetLinks sets Links field to given value.



