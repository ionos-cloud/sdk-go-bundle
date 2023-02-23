# NodePoolListResponseData

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The ID of a list of resources.  | [optional] |
|**Type** | Pointer to **string** |  | [optional] |
|**Href** | Pointer to **string** |  | [optional] |
|**Items** | [**[]NodePoolResponseData**](NodePoolResponseData.md) |  | |

## Methods

### NewNodePoolListResponseData

`func NewNodePoolListResponseData(items []NodePoolResponseData, ) *NodePoolListResponseData`

NewNodePoolListResponseData instantiates a new NodePoolListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePoolListResponseDataWithDefaults

`func NewNodePoolListResponseDataWithDefaults() *NodePoolListResponseData`

NewNodePoolListResponseDataWithDefaults instantiates a new NodePoolListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodePoolListResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodePoolListResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodePoolListResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NodePoolListResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NodePoolListResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodePoolListResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodePoolListResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NodePoolListResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *NodePoolListResponseData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NodePoolListResponseData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NodePoolListResponseData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *NodePoolListResponseData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *NodePoolListResponseData) GetItems() []NodePoolResponseData`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NodePoolListResponseData) GetItemsOk() (*[]NodePoolResponseData, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NodePoolListResponseData) SetItems(v []NodePoolResponseData)`

SetItems sets Items field to given value.



