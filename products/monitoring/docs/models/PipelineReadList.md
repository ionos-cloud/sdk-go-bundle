# PipelineReadList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of Pipeline resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of Pipeline resources. | |
|**Items** | Pointer to [**[]PipelineRead**](PipelineRead.md) | The list of Pipeline resources. | [optional] |

## Methods

### NewPipelineReadList

`func NewPipelineReadList(id string, type_ string, href string, ) *PipelineReadList`

NewPipelineReadList instantiates a new PipelineReadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineReadListWithDefaults

`func NewPipelineReadListWithDefaults() *PipelineReadList`

NewPipelineReadListWithDefaults instantiates a new PipelineReadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PipelineReadList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PipelineReadList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PipelineReadList) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PipelineReadList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PipelineReadList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PipelineReadList) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *PipelineReadList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PipelineReadList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PipelineReadList) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *PipelineReadList) GetItems() []PipelineRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PipelineReadList) GetItemsOk() (*[]PipelineRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PipelineReadList) SetItems(v []PipelineRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *PipelineReadList) HasItems() bool`

HasItems returns a boolean if a field has been set.


