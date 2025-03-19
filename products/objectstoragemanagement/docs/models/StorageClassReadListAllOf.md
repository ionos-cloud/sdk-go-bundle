# StorageClassReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of StorageClass resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of StorageClass resources. | |
|**Items** | Pointer to [**[]StorageClassRead**](StorageClassRead.md) | The list of StorageClass resources. | [optional] |

## Methods

### NewStorageClassReadListAllOf

`func NewStorageClassReadListAllOf(id string, type_ string, href string, ) *StorageClassReadListAllOf`

NewStorageClassReadListAllOf instantiates a new StorageClassReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageClassReadListAllOfWithDefaults

`func NewStorageClassReadListAllOfWithDefaults() *StorageClassReadListAllOf`

NewStorageClassReadListAllOfWithDefaults instantiates a new StorageClassReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageClassReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageClassReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageClassReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *StorageClassReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageClassReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageClassReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *StorageClassReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *StorageClassReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *StorageClassReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *StorageClassReadListAllOf) GetItems() []StorageClassRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *StorageClassReadListAllOf) GetItemsOk() (*[]StorageClassRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *StorageClassReadListAllOf) SetItems(v []StorageClassRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *StorageClassReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


