# MongoVersionReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of MongoVersion resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of MongoVersion resources. | |
|**Items** | Pointer to [**[]MongoVersionRead**](MongoVersionRead.md) | The list of MongoVersion resources. | [optional] |

## Methods

### NewMongoVersionReadListAllOf

`func NewMongoVersionReadListAllOf(id string, type_ string, href string, ) *MongoVersionReadListAllOf`

NewMongoVersionReadListAllOf instantiates a new MongoVersionReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoVersionReadListAllOfWithDefaults

`func NewMongoVersionReadListAllOfWithDefaults() *MongoVersionReadListAllOf`

NewMongoVersionReadListAllOfWithDefaults instantiates a new MongoVersionReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MongoVersionReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MongoVersionReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MongoVersionReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *MongoVersionReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MongoVersionReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MongoVersionReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *MongoVersionReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MongoVersionReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MongoVersionReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *MongoVersionReadListAllOf) GetItems() []MongoVersionRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MongoVersionReadListAllOf) GetItemsOk() (*[]MongoVersionRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MongoVersionReadListAllOf) SetItems(v []MongoVersionRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *MongoVersionReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


