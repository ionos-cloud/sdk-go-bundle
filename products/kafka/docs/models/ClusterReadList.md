# ClusterReadList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of Cluster resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of Cluster resources. | |
|**Items** | Pointer to [**[]ClusterRead**](ClusterRead.md) | The list of Cluster resources. | [optional] |

## Methods

### NewClusterReadList

`func NewClusterReadList(id string, type_ string, href string, ) *ClusterReadList`

NewClusterReadList instantiates a new ClusterReadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterReadListWithDefaults

`func NewClusterReadListWithDefaults() *ClusterReadList`

NewClusterReadListWithDefaults instantiates a new ClusterReadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterReadList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterReadList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterReadList) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ClusterReadList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterReadList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterReadList) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *ClusterReadList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ClusterReadList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ClusterReadList) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *ClusterReadList) GetItems() []ClusterRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ClusterReadList) GetItemsOk() (*[]ClusterRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ClusterReadList) SetItems(v []ClusterRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *ClusterReadList) HasItems() bool`

HasItems returns a boolean if a field has been set.


