# CentralMonitoringReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of CentralMonitoring resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of CentralMonitoring resources. | |
|**Items** | Pointer to [**[]CentralMonitoringRead**](CentralMonitoringRead.md) | The list of CentralMonitoring resources. | [optional] |

## Methods

### NewCentralMonitoringReadListAllOf

`func NewCentralMonitoringReadListAllOf(id string, type_ string, href string, ) *CentralMonitoringReadListAllOf`

NewCentralMonitoringReadListAllOf instantiates a new CentralMonitoringReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCentralMonitoringReadListAllOfWithDefaults

`func NewCentralMonitoringReadListAllOfWithDefaults() *CentralMonitoringReadListAllOf`

NewCentralMonitoringReadListAllOfWithDefaults instantiates a new CentralMonitoringReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CentralMonitoringReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CentralMonitoringReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CentralMonitoringReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CentralMonitoringReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CentralMonitoringReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CentralMonitoringReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *CentralMonitoringReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CentralMonitoringReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CentralMonitoringReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *CentralMonitoringReadListAllOf) GetItems() []CentralMonitoringRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CentralMonitoringReadListAllOf) GetItemsOk() (*[]CentralMonitoringRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CentralMonitoringReadListAllOf) SetItems(v []CentralMonitoringRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *CentralMonitoringReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


