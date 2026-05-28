# GetByContractResponseHitsHitsSourceEventParamError

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**HttpStatus** | Pointer to **int32** | HTTP status code for the given activity. | [optional] |
|**Messages** | Pointer to [**[]GetByContractResponseHitsHitsSourceEventParamErrorMessages**](GetByContractResponseHitsHitsSourceEventParamErrorMessages.md) | An array of error messages corresponding to the given activity. | [optional] |

## Methods

### NewGetByContractResponseHitsHitsSourceEventParamError

`func NewGetByContractResponseHitsHitsSourceEventParamError() *GetByContractResponseHitsHitsSourceEventParamError`

NewGetByContractResponseHitsHitsSourceEventParamError instantiates a new GetByContractResponseHitsHitsSourceEventParamError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetByContractResponseHitsHitsSourceEventParamErrorWithDefaults

`func NewGetByContractResponseHitsHitsSourceEventParamErrorWithDefaults() *GetByContractResponseHitsHitsSourceEventParamError`

NewGetByContractResponseHitsHitsSourceEventParamErrorWithDefaults instantiates a new GetByContractResponseHitsHitsSourceEventParamError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpStatus

`func (o *GetByContractResponseHitsHitsSourceEventParamError) GetHttpStatus() int32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *GetByContractResponseHitsHitsSourceEventParamError) GetHttpStatusOk() (*int32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *GetByContractResponseHitsHitsSourceEventParamError) SetHttpStatus(v int32)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *GetByContractResponseHitsHitsSourceEventParamError) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### GetMessages

`func (o *GetByContractResponseHitsHitsSourceEventParamError) GetMessages() []GetByContractResponseHitsHitsSourceEventParamErrorMessages`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *GetByContractResponseHitsHitsSourceEventParamError) GetMessagesOk() (*[]GetByContractResponseHitsHitsSourceEventParamErrorMessages, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *GetByContractResponseHitsHitsSourceEventParamError) SetMessages(v []GetByContractResponseHitsHitsSourceEventParamErrorMessages)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *GetByContractResponseHitsHitsSourceEventParamError) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


