# IPSecGatewayReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of IPSecGateway resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of IPSecGateway resources. | |
|**Items** | Pointer to [**[]IPSecGatewayRead**](IPSecGatewayRead.md) | The list of IPSecGateway resources. | [optional] |

## Methods

### NewIPSecGatewayReadListAllOf

`func NewIPSecGatewayReadListAllOf(id string, type_ string, href string, ) *IPSecGatewayReadListAllOf`

NewIPSecGatewayReadListAllOf instantiates a new IPSecGatewayReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecGatewayReadListAllOfWithDefaults

`func NewIPSecGatewayReadListAllOfWithDefaults() *IPSecGatewayReadListAllOf`

NewIPSecGatewayReadListAllOfWithDefaults instantiates a new IPSecGatewayReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPSecGatewayReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPSecGatewayReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPSecGatewayReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *IPSecGatewayReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPSecGatewayReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPSecGatewayReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *IPSecGatewayReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IPSecGatewayReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IPSecGatewayReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *IPSecGatewayReadListAllOf) GetItems() []IPSecGatewayRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IPSecGatewayReadListAllOf) GetItemsOk() (*[]IPSecGatewayRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IPSecGatewayReadListAllOf) SetItems(v []IPSecGatewayRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *IPSecGatewayReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


