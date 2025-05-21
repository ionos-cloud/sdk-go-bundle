# RestoreReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of Restore resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of Restore resources. | |
|**Items** | Pointer to [**[]RestoreRead**](RestoreRead.md) | The list of Restore resources. | [optional] |

## Methods

### NewRestoreReadListAllOf

`func NewRestoreReadListAllOf(id string, type_ string, href string, ) *RestoreReadListAllOf`

NewRestoreReadListAllOf instantiates a new RestoreReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreReadListAllOfWithDefaults

`func NewRestoreReadListAllOfWithDefaults() *RestoreReadListAllOf`

NewRestoreReadListAllOfWithDefaults instantiates a new RestoreReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestoreReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestoreReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestoreReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RestoreReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *RestoreReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RestoreReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RestoreReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *RestoreReadListAllOf) GetItems() []RestoreRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RestoreReadListAllOf) GetItemsOk() (*[]RestoreRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RestoreReadListAllOf) SetItems(v []RestoreRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *RestoreReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


