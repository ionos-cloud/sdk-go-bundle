# DeleteObjectsRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Delete** | [**DeleteObjectsRequestDelete**](DeleteObjectsRequestDelete.md) |  | |

## Methods

### NewDeleteObjectsRequest

`func NewDeleteObjectsRequest(delete DeleteObjectsRequestDelete, ) *DeleteObjectsRequest`

NewDeleteObjectsRequest instantiates a new DeleteObjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteObjectsRequestWithDefaults

`func NewDeleteObjectsRequestWithDefaults() *DeleteObjectsRequest`

NewDeleteObjectsRequestWithDefaults instantiates a new DeleteObjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelete

`func (o *DeleteObjectsRequest) GetDelete() DeleteObjectsRequestDelete`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *DeleteObjectsRequest) GetDeleteOk() (*DeleteObjectsRequestDelete, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *DeleteObjectsRequest) SetDelete(v DeleteObjectsRequestDelete)`

SetDelete sets Delete field to given value.



