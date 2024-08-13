# RouteReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the Route. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Route. | |
|**Items** | Pointer to [**[]RouteRead**](RouteRead.md) | The list of Route resources. | [optional] |

## Methods

### NewRouteReadListAllOf

`func NewRouteReadListAllOf(id string, type_ string, href string, ) *RouteReadListAllOf`

NewRouteReadListAllOf instantiates a new RouteReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteReadListAllOfWithDefaults

`func NewRouteReadListAllOfWithDefaults() *RouteReadListAllOf`

NewRouteReadListAllOfWithDefaults instantiates a new RouteReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RouteReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouteReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouteReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RouteReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *RouteReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouteReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouteReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *RouteReadListAllOf) GetItems() []RouteRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RouteReadListAllOf) GetItemsOk() (*[]RouteRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RouteReadListAllOf) SetItems(v []RouteRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *RouteReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


