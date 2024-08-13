# Route

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the route. | |
|**Type** | **string** | This field specifies the protocol used by the ingress to route traffic to the backend service. | [default to "http"]|
|**Paths** | **[]string** | The paths that the route should match.  | |
|**Methods** | **[]string** | The HTTP methods that the route should match.  | |
|**Websocket** | Pointer to **bool** | To enable websocket support. | [optional] [default to false]|
|**Upstreams** | [**[]RouteUpstreams**](RouteUpstreams.md) |  | |

## Methods

### NewRoute

`func NewRoute(name string, type_ string, paths []string, methods []string, upstreams []RouteUpstreams, ) *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Route) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Route) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Route) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Route) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Route) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Route) SetType(v string)`

SetType sets Type field to given value.


### GetPaths

`func (o *Route) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *Route) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *Route) SetPaths(v []string)`

SetPaths sets Paths field to given value.


### GetMethods

`func (o *Route) GetMethods() []string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *Route) GetMethodsOk() (*[]string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *Route) SetMethods(v []string)`

SetMethods sets Methods field to given value.


### GetWebsocket

`func (o *Route) GetWebsocket() bool`

GetWebsocket returns the Websocket field if non-nil, zero value otherwise.

### GetWebsocketOk

`func (o *Route) GetWebsocketOk() (*bool, bool)`

GetWebsocketOk returns a tuple with the Websocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocket

`func (o *Route) SetWebsocket(v bool)`

SetWebsocket sets Websocket field to given value.

### HasWebsocket

`func (o *Route) HasWebsocket() bool`

HasWebsocket returns a boolean if a field has been set.

### GetUpstreams

`func (o *Route) GetUpstreams() []RouteUpstreams`

GetUpstreams returns the Upstreams field if non-nil, zero value otherwise.

### GetUpstreamsOk

`func (o *Route) GetUpstreamsOk() (*[]RouteUpstreams, bool)`

GetUpstreamsOk returns a tuple with the Upstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreams

`func (o *Route) SetUpstreams(v []RouteUpstreams)`

SetUpstreams sets Upstreams field to given value.



