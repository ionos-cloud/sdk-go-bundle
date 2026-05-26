# GetByContractResponseHitsHitsSourceEvent

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Param** | Pointer to [**GetByContractResponseHitsHitsSourceEventParam**](GetByContractResponseHitsHitsSourceEventParam.md) |  | [optional] |
|**Resources** | Pointer to [**[]GetByContractResponseHitsHitsSourceEventResources**](GetByContractResponseHitsHitsSourceEventResources.md) | An array of resources affected by the given activity. | [optional] |
|**Message** | Pointer to **string** | Message explaining the current status of the request. | [optional] |
|**Status** | Pointer to **string** | The current status of the request. | [optional] |
|**Type** | Pointer to **string** | Type of the activity event. The exact contents of the &#x60;event&#x60; object will vary depending on this value.  There are a large number of possible values, examples are * &#x60;Error&#x60; for a request that couldn&#39;t be completed, * &#x60;RequestAccepted&#x60; for an incoming request (e.g. via HTTP), * &#x60;RequestStatusUpdate&#x60; for a later update about a request which was   logged earlier, * &#x60;Provision&#x60; for resource provisionings running in the background or  * &#x60;DCDUserEvent&#x60; for events about a DCD user.  | [optional] |

## Methods

### NewGetByContractResponseHitsHitsSourceEvent

`func NewGetByContractResponseHitsHitsSourceEvent() *GetByContractResponseHitsHitsSourceEvent`

NewGetByContractResponseHitsHitsSourceEvent instantiates a new GetByContractResponseHitsHitsSourceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetByContractResponseHitsHitsSourceEventWithDefaults

`func NewGetByContractResponseHitsHitsSourceEventWithDefaults() *GetByContractResponseHitsHitsSourceEvent`

NewGetByContractResponseHitsHitsSourceEventWithDefaults instantiates a new GetByContractResponseHitsHitsSourceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParam

`func (o *GetByContractResponseHitsHitsSourceEvent) GetParam() GetByContractResponseHitsHitsSourceEventParam`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *GetByContractResponseHitsHitsSourceEvent) GetParamOk() (*GetByContractResponseHitsHitsSourceEventParam, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *GetByContractResponseHitsHitsSourceEvent) SetParam(v GetByContractResponseHitsHitsSourceEventParam)`

SetParam sets Param field to given value.

### HasParam

`func (o *GetByContractResponseHitsHitsSourceEvent) HasParam() bool`

HasParam returns a boolean if a field has been set.

### GetResources

`func (o *GetByContractResponseHitsHitsSourceEvent) GetResources() []GetByContractResponseHitsHitsSourceEventResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GetByContractResponseHitsHitsSourceEvent) GetResourcesOk() (*[]GetByContractResponseHitsHitsSourceEventResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GetByContractResponseHitsHitsSourceEvent) SetResources(v []GetByContractResponseHitsHitsSourceEventResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *GetByContractResponseHitsHitsSourceEvent) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetMessage

`func (o *GetByContractResponseHitsHitsSourceEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetByContractResponseHitsHitsSourceEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetByContractResponseHitsHitsSourceEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetByContractResponseHitsHitsSourceEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *GetByContractResponseHitsHitsSourceEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetByContractResponseHitsHitsSourceEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetByContractResponseHitsHitsSourceEvent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetByContractResponseHitsHitsSourceEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *GetByContractResponseHitsHitsSourceEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetByContractResponseHitsHitsSourceEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetByContractResponseHitsHitsSourceEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetByContractResponseHitsHitsSourceEvent) HasType() bool`

HasType returns a boolean if a field has been set.


