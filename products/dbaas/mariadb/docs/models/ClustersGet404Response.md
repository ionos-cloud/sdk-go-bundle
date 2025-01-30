# ClustersGet404Response

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**HttpStatus** | **int32** | The HTTP status code of the operation. | |
|**Messages** | [**[]ErrorMessage**](ErrorMessage.md) |  | |

## Methods

### NewClustersGet404Response

`func NewClustersGet404Response(httpStatus int32, messages []ErrorMessage, ) *ClustersGet404Response`

NewClustersGet404Response instantiates a new ClustersGet404Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClustersGet404ResponseWithDefaults

`func NewClustersGet404ResponseWithDefaults() *ClustersGet404Response`

NewClustersGet404ResponseWithDefaults instantiates a new ClustersGet404Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpStatus

`func (o *ClustersGet404Response) GetHttpStatus() int32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *ClustersGet404Response) GetHttpStatusOk() (*int32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *ClustersGet404Response) SetHttpStatus(v int32)`

SetHttpStatus sets HttpStatus field to given value.


### GetMessages

`func (o *ClustersGet404Response) GetMessages() []ErrorMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ClustersGet404Response) GetMessagesOk() (*[]ErrorMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ClustersGet404Response) SetMessages(v []ErrorMessage)`

SetMessages sets Messages field to given value.



