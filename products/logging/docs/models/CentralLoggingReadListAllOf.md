# CentralLoggingReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of CentralLogging resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of CentralLogging resources. | |
|**Items** | Pointer to [**[]CentralLoggingRead**](CentralLoggingRead.md) | The list of CentralLogging resources. | [optional] |

## Methods

### NewCentralLoggingReadListAllOf

`func NewCentralLoggingReadListAllOf(id string, type_ string, href string, ) *CentralLoggingReadListAllOf`

NewCentralLoggingReadListAllOf instantiates a new CentralLoggingReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCentralLoggingReadListAllOfWithDefaults

`func NewCentralLoggingReadListAllOfWithDefaults() *CentralLoggingReadListAllOf`

NewCentralLoggingReadListAllOfWithDefaults instantiates a new CentralLoggingReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CentralLoggingReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CentralLoggingReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CentralLoggingReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CentralLoggingReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CentralLoggingReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CentralLoggingReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *CentralLoggingReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CentralLoggingReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CentralLoggingReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *CentralLoggingReadListAllOf) GetItems() []CentralLoggingRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CentralLoggingReadListAllOf) GetItemsOk() (*[]CentralLoggingRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CentralLoggingReadListAllOf) SetItems(v []CentralLoggingRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *CentralLoggingReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


