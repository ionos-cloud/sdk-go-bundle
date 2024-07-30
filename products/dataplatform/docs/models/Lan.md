# Lan

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**LanId** | **string** | The LAN ID of an existing LAN at the related data center.  | |
|**Dhcp** | Pointer to **bool** | Indicates if the Kubernetes node pool LAN will reserve an IP using DHCP. The default value is &#x60;true&#x60;.  | [optional] |
|**Routes** | Pointer to [**[]Route**](Route.md) | An array of additional LANs attached to worker nodes.  | [optional] |

## Methods

### NewLan

`func NewLan(lanId string, ) *Lan`

NewLan instantiates a new Lan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanWithDefaults

`func NewLanWithDefaults() *Lan`

NewLanWithDefaults instantiates a new Lan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanId

`func (o *Lan) GetLanId() string`

GetLanId returns the LanId field if non-nil, zero value otherwise.

### GetLanIdOk

`func (o *Lan) GetLanIdOk() (*string, bool)`

GetLanIdOk returns a tuple with the LanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanId

`func (o *Lan) SetLanId(v string)`

SetLanId sets LanId field to given value.


### GetDhcp

`func (o *Lan) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *Lan) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *Lan) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *Lan) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetRoutes

`func (o *Lan) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *Lan) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *Lan) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *Lan) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


