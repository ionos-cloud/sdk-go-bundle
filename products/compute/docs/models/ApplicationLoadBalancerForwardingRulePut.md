# ApplicationLoadBalancerForwardingRulePut

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | The URL to the object representation (absolute path). | [optional] [readonly] |
|**Properties** | [**ApplicationLoadBalancerForwardingRuleProperties**](ApplicationLoadBalancerForwardingRuleProperties.md) |  | |

## Methods

### NewApplicationLoadBalancerForwardingRulePut

`func NewApplicationLoadBalancerForwardingRulePut(properties ApplicationLoadBalancerForwardingRuleProperties, ) *ApplicationLoadBalancerForwardingRulePut`

NewApplicationLoadBalancerForwardingRulePut instantiates a new ApplicationLoadBalancerForwardingRulePut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationLoadBalancerForwardingRulePutWithDefaults

`func NewApplicationLoadBalancerForwardingRulePutWithDefaults() *ApplicationLoadBalancerForwardingRulePut`

NewApplicationLoadBalancerForwardingRulePutWithDefaults instantiates a new ApplicationLoadBalancerForwardingRulePut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationLoadBalancerForwardingRulePut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationLoadBalancerForwardingRulePut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationLoadBalancerForwardingRulePut) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationLoadBalancerForwardingRulePut) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ApplicationLoadBalancerForwardingRulePut) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationLoadBalancerForwardingRulePut) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationLoadBalancerForwardingRulePut) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationLoadBalancerForwardingRulePut) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *ApplicationLoadBalancerForwardingRulePut) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ApplicationLoadBalancerForwardingRulePut) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ApplicationLoadBalancerForwardingRulePut) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ApplicationLoadBalancerForwardingRulePut) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetProperties

`func (o *ApplicationLoadBalancerForwardingRulePut) GetProperties() ApplicationLoadBalancerForwardingRuleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ApplicationLoadBalancerForwardingRulePut) GetPropertiesOk() (*ApplicationLoadBalancerForwardingRuleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ApplicationLoadBalancerForwardingRulePut) SetProperties(v ApplicationLoadBalancerForwardingRuleProperties)`

SetProperties sets Properties field to given value.



