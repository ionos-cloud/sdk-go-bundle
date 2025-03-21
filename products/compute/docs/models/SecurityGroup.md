# SecurityGroup

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | The URL to the object representation (absolute path). | [optional] [readonly] |
|**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] |
|**Properties** | [**SecurityGroupProperties**](SecurityGroupProperties.md) |  | |
|**Entities** | Pointer to [**SecurityGroupEntities**](SecurityGroupEntities.md) |  | [optional] |

## Methods

### NewSecurityGroup

`func NewSecurityGroup(properties SecurityGroupProperties, ) *SecurityGroup`

NewSecurityGroup instantiates a new SecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupWithDefaults

`func NewSecurityGroupWithDefaults() *SecurityGroup`

NewSecurityGroupWithDefaults instantiates a new SecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SecurityGroup) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityGroup) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityGroup) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *SecurityGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *SecurityGroup) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SecurityGroup) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SecurityGroup) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SecurityGroup) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *SecurityGroup) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecurityGroup) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecurityGroup) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SecurityGroup) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *SecurityGroup) GetProperties() SecurityGroupProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SecurityGroup) GetPropertiesOk() (*SecurityGroupProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SecurityGroup) SetProperties(v SecurityGroupProperties)`

SetProperties sets Properties field to given value.


### GetEntities

`func (o *SecurityGroup) GetEntities() SecurityGroupEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *SecurityGroup) GetEntitiesOk() (*SecurityGroupEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *SecurityGroup) SetEntities(v SecurityGroupEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *SecurityGroup) HasEntities() bool`

HasEntities returns a boolean if a field has been set.


