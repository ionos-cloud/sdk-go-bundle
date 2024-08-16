# Distribution

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the Distribution. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Distribution. | |
|**Metadata** | [**DistributionMetadata**](DistributionMetadata.md) |  | |
|**Properties** | [**DistributionProperties**](DistributionProperties.md) |  | |

## Methods

### NewDistribution

`func NewDistribution(id string, type_ string, href string, metadata DistributionMetadata, properties DistributionProperties, ) *Distribution`

NewDistribution instantiates a new Distribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionWithDefaults

`func NewDistributionWithDefaults() *Distribution`

NewDistributionWithDefaults instantiates a new Distribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Distribution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Distribution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Distribution) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Distribution) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Distribution) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Distribution) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *Distribution) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Distribution) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Distribution) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *Distribution) GetMetadata() DistributionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Distribution) GetMetadataOk() (*DistributionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Distribution) SetMetadata(v DistributionMetadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *Distribution) GetProperties() DistributionProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Distribution) GetPropertiesOk() (*DistributionProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Distribution) SetProperties(v DistributionProperties)`

SetProperties sets Properties field to given value.



