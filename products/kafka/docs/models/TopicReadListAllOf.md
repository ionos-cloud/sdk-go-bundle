# TopicReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of Topic resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of Topic resources. | |
|**Items** | Pointer to [**[]TopicRead**](TopicRead.md) | The list of Topic resources. | [optional] |

## Methods

### NewTopicReadListAllOf

`func NewTopicReadListAllOf(id string, type_ string, href string, ) *TopicReadListAllOf`

NewTopicReadListAllOf instantiates a new TopicReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicReadListAllOfWithDefaults

`func NewTopicReadListAllOfWithDefaults() *TopicReadListAllOf`

NewTopicReadListAllOfWithDefaults instantiates a new TopicReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TopicReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopicReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopicReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *TopicReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopicReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopicReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *TopicReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *TopicReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *TopicReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *TopicReadListAllOf) GetItems() []TopicRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TopicReadListAllOf) GetItemsOk() (*[]TopicRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TopicReadListAllOf) SetItems(v []TopicRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *TopicReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


