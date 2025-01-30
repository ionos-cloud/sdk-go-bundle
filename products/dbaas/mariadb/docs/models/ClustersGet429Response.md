# ClustersGet429Response

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**HttpStatus** | **int32** | The HTTP status code of the operation. | |
|**Messages** | [**[]ErrorMessage**](ErrorMessage.md) |  | |

## Methods

### NewClustersGet429Response

`func NewClustersGet429Response(httpStatus int32, messages []ErrorMessage, ) *ClustersGet429Response`

NewClustersGet429Response instantiates a new ClustersGet429Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClustersGet429ResponseWithDefaults

`func NewClustersGet429ResponseWithDefaults() *ClustersGet429Response`

NewClustersGet429ResponseWithDefaults instantiates a new ClustersGet429Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpStatus

`func (o *ClustersGet429Response) GetHttpStatus() int32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *ClustersGet429Response) GetHttpStatusOk() (*int32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *ClustersGet429Response) SetHttpStatus(v int32)`

SetHttpStatus sets HttpStatus field to given value.


### GetMessages

`func (o *ClustersGet429Response) GetMessages() []ErrorMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ClustersGet429Response) GetMessagesOk() (*[]ErrorMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ClustersGet429Response) SetMessages(v []ErrorMessage)`

SetMessages sets Messages field to given value.



