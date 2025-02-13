# AccessKeyReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of AccessKey resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of AccessKey resources. | |
|**Items** | Pointer to [**[]AccessKeyRead**](AccessKeyRead.md) | The list of AccessKey resources. | [optional] |

## Methods

### NewAccessKeyReadListAllOf

`func NewAccessKeyReadListAllOf(id string, type_ string, href string, ) *AccessKeyReadListAllOf`

NewAccessKeyReadListAllOf instantiates a new AccessKeyReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyReadListAllOfWithDefaults

`func NewAccessKeyReadListAllOfWithDefaults() *AccessKeyReadListAllOf`

NewAccessKeyReadListAllOfWithDefaults instantiates a new AccessKeyReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessKeyReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessKeyReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessKeyReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *AccessKeyReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessKeyReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessKeyReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *AccessKeyReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AccessKeyReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AccessKeyReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *AccessKeyReadListAllOf) GetItems() []AccessKeyRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AccessKeyReadListAllOf) GetItemsOk() (*[]AccessKeyRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AccessKeyReadListAllOf) SetItems(v []AccessKeyRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *AccessKeyReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


