# DistributionMetadata

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**CreatedDate** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 creation timestamp. | [optional] [readonly] |
|**CreatedBy** | Pointer to **string** | Unique name of the identity that created the resource. | [optional] [readonly] |
|**CreatedByUserId** | Pointer to **string** | Unique id of the identity that created the resource. | [optional] [readonly] |
|**LastModifiedDate** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 modified timestamp. | [optional] [readonly] |
|**LastModifiedBy** | Pointer to **string** | Unique name of the identity that last modified the resource. | [optional] [readonly] |
|**LastModifiedByUserId** | Pointer to **string** | Unique id of the identity that last modified the resource. | [optional] [readonly] |
|**ResourceURN** | Pointer to **string** | Unique name of the resource. | [optional] [readonly] |
|**PublicEndpointIpv4** | Pointer to **string** | IP of the distribution. It has to be included on the domain DNS Zone as A record. | [optional] [readonly] |
|**PublicEndpointIpv6** | Pointer to **string** | IP of the distribution, it has to be included on the domain DNS Zone as AAAA record. | [optional] [readonly] |
|**State** | **string** | Represents one of the possible states of the resource. | [readonly] |
|**Message** | Pointer to **string** | A human readable message describing the current state. In case of an error, the message will contain a detailed error message.  | [optional] [readonly] |

## Methods

### NewDistributionMetadata

`func NewDistributionMetadata(state string, ) *DistributionMetadata`

NewDistributionMetadata instantiates a new DistributionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionMetadataWithDefaults

`func NewDistributionMetadataWithDefaults() *DistributionMetadata`

NewDistributionMetadataWithDefaults instantiates a new DistributionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *DistributionMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *DistributionMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *DistributionMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *DistributionMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DistributionMetadata) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DistributionMetadata) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DistributionMetadata) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DistributionMetadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *DistributionMetadata) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *DistributionMetadata) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *DistributionMetadata) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *DistributionMetadata) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *DistributionMetadata) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *DistributionMetadata) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *DistributionMetadata) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *DistributionMetadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *DistributionMetadata) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *DistributionMetadata) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *DistributionMetadata) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *DistributionMetadata) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *DistributionMetadata) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *DistributionMetadata) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *DistributionMetadata) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *DistributionMetadata) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetResourceURN

`func (o *DistributionMetadata) GetResourceURN() string`

GetResourceURN returns the ResourceURN field if non-nil, zero value otherwise.

### GetResourceURNOk

`func (o *DistributionMetadata) GetResourceURNOk() (*string, bool)`

GetResourceURNOk returns a tuple with the ResourceURN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceURN

`func (o *DistributionMetadata) SetResourceURN(v string)`

SetResourceURN sets ResourceURN field to given value.

### HasResourceURN

`func (o *DistributionMetadata) HasResourceURN() bool`

HasResourceURN returns a boolean if a field has been set.

### GetPublicEndpointIpv4

`func (o *DistributionMetadata) GetPublicEndpointIpv4() string`

GetPublicEndpointIpv4 returns the PublicEndpointIpv4 field if non-nil, zero value otherwise.

### GetPublicEndpointIpv4Ok

`func (o *DistributionMetadata) GetPublicEndpointIpv4Ok() (*string, bool)`

GetPublicEndpointIpv4Ok returns a tuple with the PublicEndpointIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicEndpointIpv4

`func (o *DistributionMetadata) SetPublicEndpointIpv4(v string)`

SetPublicEndpointIpv4 sets PublicEndpointIpv4 field to given value.

### HasPublicEndpointIpv4

`func (o *DistributionMetadata) HasPublicEndpointIpv4() bool`

HasPublicEndpointIpv4 returns a boolean if a field has been set.

### GetPublicEndpointIpv6

`func (o *DistributionMetadata) GetPublicEndpointIpv6() string`

GetPublicEndpointIpv6 returns the PublicEndpointIpv6 field if non-nil, zero value otherwise.

### GetPublicEndpointIpv6Ok

`func (o *DistributionMetadata) GetPublicEndpointIpv6Ok() (*string, bool)`

GetPublicEndpointIpv6Ok returns a tuple with the PublicEndpointIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicEndpointIpv6

`func (o *DistributionMetadata) SetPublicEndpointIpv6(v string)`

SetPublicEndpointIpv6 sets PublicEndpointIpv6 field to given value.

### HasPublicEndpointIpv6

`func (o *DistributionMetadata) HasPublicEndpointIpv6() bool`

HasPublicEndpointIpv6 returns a boolean if a field has been set.

### GetState

`func (o *DistributionMetadata) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DistributionMetadata) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DistributionMetadata) SetState(v string)`

SetState sets State field to given value.


### GetMessage

`func (o *DistributionMetadata) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DistributionMetadata) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DistributionMetadata) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DistributionMetadata) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


