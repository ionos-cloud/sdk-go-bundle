# Route

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Network** | **string** | IPv4 or IPv6 CIDR to be routed via the interface.  | |
|**Gateway** | **string** | IPv4 or IPv6 gateway IP for the route.  | |

## Methods

### NewRoute

`func NewRoute(network string, gateway string, ) *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *Route) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Route) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Route) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetGateway

`func (o *Route) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Route) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Route) SetGateway(v string)`

SetGateway sets Gateway field to given value.



