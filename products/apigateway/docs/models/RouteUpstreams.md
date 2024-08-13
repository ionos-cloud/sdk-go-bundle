# RouteUpstreams

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Scheme** | **string** | The target URL of the upstream. | [default to "http"]|
|**Loadbalancer** | **string** | The load balancer algorithm. | [default to "roundrobin"]|
|**Host** | **string** | The host of the upstream. | |
|**Port** | **int32** | The port of the upstream. | [default to 80]|
|**Weight** | Pointer to **int32** | Weight with which to split traffic to the upstream. | [optional] [default to 100]|

## Methods

### NewRouteUpstreams

`func NewRouteUpstreams(scheme string, loadbalancer string, host string, port int32, ) *RouteUpstreams`

NewRouteUpstreams instantiates a new RouteUpstreams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteUpstreamsWithDefaults

`func NewRouteUpstreamsWithDefaults() *RouteUpstreams`

NewRouteUpstreamsWithDefaults instantiates a new RouteUpstreams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheme

`func (o *RouteUpstreams) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *RouteUpstreams) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *RouteUpstreams) SetScheme(v string)`

SetScheme sets Scheme field to given value.


### GetLoadbalancer

`func (o *RouteUpstreams) GetLoadbalancer() string`

GetLoadbalancer returns the Loadbalancer field if non-nil, zero value otherwise.

### GetLoadbalancerOk

`func (o *RouteUpstreams) GetLoadbalancerOk() (*string, bool)`

GetLoadbalancerOk returns a tuple with the Loadbalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadbalancer

`func (o *RouteUpstreams) SetLoadbalancer(v string)`

SetLoadbalancer sets Loadbalancer field to given value.


### GetHost

`func (o *RouteUpstreams) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RouteUpstreams) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RouteUpstreams) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *RouteUpstreams) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RouteUpstreams) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RouteUpstreams) SetPort(v int32)`

SetPort sets Port field to given value.


### GetWeight

`func (o *RouteUpstreams) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RouteUpstreams) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RouteUpstreams) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *RouteUpstreams) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


