# TemplateListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] |
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**Items** | Pointer to [**[]TemplateResponse**](TemplateResponse.md) |  | [optional] |

## Methods

### NewTemplateListAllOf

`func NewTemplateListAllOf() *TemplateListAllOf`

NewTemplateListAllOf instantiates a new TemplateListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateListAllOfWithDefaults

`func NewTemplateListAllOfWithDefaults() *TemplateListAllOf`

NewTemplateListAllOfWithDefaults instantiates a new TemplateListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TemplateListAllOf) GetType() ResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateListAllOf) GetTypeOk() (*ResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateListAllOf) SetType(v ResourceType)`

SetType sets Type field to given value.

### HasType

`func (o *TemplateListAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *TemplateListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateListAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateListAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *TemplateListAllOf) GetItems() []TemplateResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TemplateListAllOf) GetItemsOk() (*[]TemplateResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TemplateListAllOf) SetItems(v []TemplateResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *TemplateListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


