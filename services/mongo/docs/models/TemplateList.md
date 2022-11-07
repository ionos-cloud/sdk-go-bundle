# TemplateList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Items** | Pointer to [**[]TemplateResponse**](TemplateResponse.md) |  | [optional] |

## Methods

### NewTemplateList

`func NewTemplateList() *TemplateList`

NewTemplateList instantiates a new TemplateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateListWithDefaults

`func NewTemplateListWithDefaults() *TemplateList`

NewTemplateListWithDefaults instantiates a new TemplateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TemplateList) GetItems() []TemplateResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TemplateList) GetItemsOk() (*[]TemplateResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TemplateList) SetItems(v []TemplateResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *TemplateList) HasItems() bool`

HasItems returns a boolean if a field has been set.


