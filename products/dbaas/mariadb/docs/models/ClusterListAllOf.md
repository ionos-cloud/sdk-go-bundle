# ClusterListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**Items** | Pointer to [**[]ClusterResponse**](ClusterResponse.md) |  | [optional] |

## Methods

### NewClusterListAllOf

`func NewClusterListAllOf() *ClusterListAllOf`

NewClusterListAllOf instantiates a new ClusterListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterListAllOfWithDefaults

`func NewClusterListAllOfWithDefaults() *ClusterListAllOf`

NewClusterListAllOfWithDefaults instantiates a new ClusterListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterListAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterListAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *ClusterListAllOf) GetItems() []ClusterResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ClusterListAllOf) GetItemsOk() (*[]ClusterResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ClusterListAllOf) SetItems(v []ClusterResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *ClusterListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


