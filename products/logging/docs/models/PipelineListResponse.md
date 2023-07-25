# PipelineListResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** |  | [optional] |
|**Type** | Pointer to **string** |  | [optional] |
|**Items** | Pointer to [**[]Pipeline**](Pipeline.md) |  | [optional] |

## Methods

### NewPipelineListResponse

`func NewPipelineListResponse() *PipelineListResponse`

NewPipelineListResponse instantiates a new PipelineListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineListResponseWithDefaults

`func NewPipelineListResponseWithDefaults() *PipelineListResponse`

NewPipelineListResponseWithDefaults instantiates a new PipelineListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PipelineListResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PipelineListResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PipelineListResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PipelineListResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PipelineListResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PipelineListResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PipelineListResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PipelineListResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItems

`func (o *PipelineListResponse) GetItems() []Pipeline`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PipelineListResponse) GetItemsOk() (*[]Pipeline, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PipelineListResponse) SetItems(v []Pipeline)`

SetItems sets Items field to given value.

### HasItems

`func (o *PipelineListResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.


