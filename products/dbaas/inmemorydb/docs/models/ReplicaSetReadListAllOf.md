# ReplicaSetReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of ReplicaSet resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of ReplicaSet resources. | |
|**Items** | Pointer to [**[]ReplicaSetRead**](ReplicaSetRead.md) | The list of ReplicaSet resources. | [optional] |

## Methods

### NewReplicaSetReadListAllOf

`func NewReplicaSetReadListAllOf(id string, type_ string, href string, ) *ReplicaSetReadListAllOf`

NewReplicaSetReadListAllOf instantiates a new ReplicaSetReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaSetReadListAllOfWithDefaults

`func NewReplicaSetReadListAllOfWithDefaults() *ReplicaSetReadListAllOf`

NewReplicaSetReadListAllOfWithDefaults instantiates a new ReplicaSetReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReplicaSetReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReplicaSetReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReplicaSetReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ReplicaSetReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReplicaSetReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReplicaSetReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *ReplicaSetReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ReplicaSetReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ReplicaSetReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *ReplicaSetReadListAllOf) GetItems() []ReplicaSetRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ReplicaSetReadListAllOf) GetItemsOk() (*[]ReplicaSetRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ReplicaSetReadListAllOf) SetItems(v []ReplicaSetRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *ReplicaSetReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


