# ResourceMetadata

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
|**State** | **string** | State of the resource. Resource states: &#x60;AVAILABLE&#x60;: There are no pending modification requests for this item. &#x60;BUSY&#x60;: There is at least one modification request pending and all following requests will be queued. &#x60;DEPLOYING&#x60;: The resource is being created. &#x60;FAILED&#x60;: The creation of the resource failed. &#x60;UPDATING&#x60;: The resource is being updated. &#x60;FAILED_UPDATING&#x60;: An update to the resource was not successful. &#x60;DESTROYING&#x60;: A delete command was issued, and the resource is being deleted.  | |
|**Message** | Pointer to **string** | A human readable message describing the current state. In case of an error, the message will contain a detailed error message.  | [optional] |

## Methods

### NewResourceMetadata

`func NewResourceMetadata(state string, ) *ResourceMetadata`

NewResourceMetadata instantiates a new ResourceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMetadataWithDefaults

`func NewResourceMetadataWithDefaults() *ResourceMetadata`

NewResourceMetadataWithDefaults instantiates a new ResourceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *ResourceMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ResourceMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ResourceMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ResourceMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ResourceMetadata) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ResourceMetadata) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ResourceMetadata) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ResourceMetadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ResourceMetadata) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ResourceMetadata) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ResourceMetadata) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ResourceMetadata) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *ResourceMetadata) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *ResourceMetadata) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *ResourceMetadata) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *ResourceMetadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *ResourceMetadata) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *ResourceMetadata) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *ResourceMetadata) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *ResourceMetadata) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *ResourceMetadata) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *ResourceMetadata) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *ResourceMetadata) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *ResourceMetadata) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetResourceURN

`func (o *ResourceMetadata) GetResourceURN() string`

GetResourceURN returns the ResourceURN field if non-nil, zero value otherwise.

### GetResourceURNOk

`func (o *ResourceMetadata) GetResourceURNOk() (*string, bool)`

GetResourceURNOk returns a tuple with the ResourceURN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceURN

`func (o *ResourceMetadata) SetResourceURN(v string)`

SetResourceURN sets ResourceURN field to given value.

### HasResourceURN

`func (o *ResourceMetadata) HasResourceURN() bool`

HasResourceURN returns a boolean if a field has been set.

### GetState

`func (o *ResourceMetadata) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResourceMetadata) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResourceMetadata) SetState(v string)`

SetState sets State field to given value.


### GetMessage

`func (o *ResourceMetadata) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResourceMetadata) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResourceMetadata) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResourceMetadata) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


