# Images

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | The URL to the object representation (absolute path). | [optional] [readonly] |
|**Items** | Pointer to [**[]Image**](Image.md) | Array of items in the collection. | [optional] [readonly] |

## Methods

### NewImages

`func NewImages() *Images`

NewImages instantiates a new Images object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagesWithDefaults

`func NewImagesWithDefaults() *Images`

NewImagesWithDefaults instantiates a new Images object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Images) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Images) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Images) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Images) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Images) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Images) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Images) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *Images) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Images) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Images) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Images) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Images) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *Images) GetItems() []Image`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Images) GetItemsOk() (*[]Image, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Images) SetItems(v []Image)`

SetItems sets Items field to given value.

### HasItems

`func (o *Images) HasItems() bool`

HasItems returns a boolean if a field has been set.


