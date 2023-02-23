# ClusterListResponseData

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The ID of a list of resources.  | [optional] |
|**Type** | Pointer to **string** |  | [optional] |
|**Href** | Pointer to **string** |  | [optional] |
|**Items** | [**[]ClusterResponseData**](ClusterResponseData.md) |  | |

## Methods

### NewClusterListResponseData

`func NewClusterListResponseData(items []ClusterResponseData, ) *ClusterListResponseData`

NewClusterListResponseData instantiates a new ClusterListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterListResponseDataWithDefaults

`func NewClusterListResponseDataWithDefaults() *ClusterListResponseData`

NewClusterListResponseDataWithDefaults instantiates a new ClusterListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterListResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterListResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterListResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterListResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ClusterListResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterListResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterListResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterListResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *ClusterListResponseData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ClusterListResponseData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ClusterListResponseData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ClusterListResponseData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *ClusterListResponseData) GetItems() []ClusterResponseData`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ClusterListResponseData) GetItemsOk() (*[]ClusterResponseData, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ClusterListResponseData) SetItems(v []ClusterResponseData)`

SetItems sets Items field to given value.



