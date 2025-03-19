# BucketReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of Bucket resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of Bucket resources. | |
|**Items** | Pointer to [**[]BucketRead**](BucketRead.md) | The list of Bucket resources. | [optional] |

## Methods

### NewBucketReadListAllOf

`func NewBucketReadListAllOf(id string, type_ string, href string, ) *BucketReadListAllOf`

NewBucketReadListAllOf instantiates a new BucketReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketReadListAllOfWithDefaults

`func NewBucketReadListAllOfWithDefaults() *BucketReadListAllOf`

NewBucketReadListAllOfWithDefaults instantiates a new BucketReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BucketReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BucketReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BucketReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *BucketReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BucketReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BucketReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *BucketReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BucketReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BucketReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *BucketReadListAllOf) GetItems() []BucketRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BucketReadListAllOf) GetItemsOk() (*[]BucketRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BucketReadListAllOf) SetItems(v []BucketRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *BucketReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


