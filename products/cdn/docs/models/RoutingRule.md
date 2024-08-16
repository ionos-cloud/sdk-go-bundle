# RoutingRule

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Scheme** | **string** | The scheme of the routing rule. | |
|**Prefix** | **string** | The prefix of the routing rule. | |
|**Upstream** | [**Upstream**](Upstream.md) |  | |

## Methods

### NewRoutingRule

`func NewRoutingRule(scheme string, prefix string, upstream Upstream, ) *RoutingRule`

NewRoutingRule instantiates a new RoutingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingRuleWithDefaults

`func NewRoutingRuleWithDefaults() *RoutingRule`

NewRoutingRuleWithDefaults instantiates a new RoutingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheme

`func (o *RoutingRule) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *RoutingRule) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *RoutingRule) SetScheme(v string)`

SetScheme sets Scheme field to given value.


### GetPrefix

`func (o *RoutingRule) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RoutingRule) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RoutingRule) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetUpstream

`func (o *RoutingRule) GetUpstream() Upstream`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *RoutingRule) GetUpstreamOk() (*Upstream, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *RoutingRule) SetUpstream(v Upstream)`

SetUpstream sets Upstream field to given value.



