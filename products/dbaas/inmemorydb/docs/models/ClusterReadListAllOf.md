# ClusterReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of Cluster resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of Cluster resources. | |
|**Items** | Pointer to [**[]ClusterRead**](ClusterRead.md) | The list of Cluster resources. | [optional] |

## Methods

### NewClusterReadListAllOf

`func NewClusterReadListAllOf(id string, type_ string, href string, ) *ClusterReadListAllOf`

NewClusterReadListAllOf instantiates a new ClusterReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterReadListAllOfWithDefaults

`func NewClusterReadListAllOfWithDefaults() *ClusterReadListAllOf`

NewClusterReadListAllOfWithDefaults instantiates a new ClusterReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ClusterReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *ClusterReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ClusterReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ClusterReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *ClusterReadListAllOf) GetItems() []ClusterRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ClusterReadListAllOf) GetItemsOk() (*[]ClusterRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ClusterReadListAllOf) SetItems(v []ClusterRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *ClusterReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


