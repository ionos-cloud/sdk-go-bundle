# ApplicationLoadBalancer

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | The URL to the object representation (absolute path). | [optional] [readonly] |
|**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] |
|**Properties** | [**ApplicationLoadBalancerProperties**](ApplicationLoadBalancerProperties.md) |  | |
|**Entities** | Pointer to [**ApplicationLoadBalancerEntities**](ApplicationLoadBalancerEntities.md) |  | [optional] |

## Methods

### NewApplicationLoadBalancer

`func NewApplicationLoadBalancer(properties ApplicationLoadBalancerProperties, ) *ApplicationLoadBalancer`

NewApplicationLoadBalancer instantiates a new ApplicationLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationLoadBalancerWithDefaults

`func NewApplicationLoadBalancerWithDefaults() *ApplicationLoadBalancer`

NewApplicationLoadBalancerWithDefaults instantiates a new ApplicationLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationLoadBalancer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationLoadBalancer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationLoadBalancer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationLoadBalancer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ApplicationLoadBalancer) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationLoadBalancer) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationLoadBalancer) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationLoadBalancer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *ApplicationLoadBalancer) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ApplicationLoadBalancer) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ApplicationLoadBalancer) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ApplicationLoadBalancer) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *ApplicationLoadBalancer) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationLoadBalancer) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationLoadBalancer) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationLoadBalancer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *ApplicationLoadBalancer) GetProperties() ApplicationLoadBalancerProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ApplicationLoadBalancer) GetPropertiesOk() (*ApplicationLoadBalancerProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ApplicationLoadBalancer) SetProperties(v ApplicationLoadBalancerProperties)`

SetProperties sets Properties field to given value.


### GetEntities

`func (o *ApplicationLoadBalancer) GetEntities() ApplicationLoadBalancerEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ApplicationLoadBalancer) GetEntitiesOk() (*ApplicationLoadBalancerEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ApplicationLoadBalancer) SetEntities(v ApplicationLoadBalancerEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ApplicationLoadBalancer) HasEntities() bool`

HasEntities returns a boolean if a field has been set.


