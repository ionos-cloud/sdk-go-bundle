# GetByContractResponseHitsHitsSourceMeta

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**AuditVersion** | Pointer to **float32** | The values returned is currently 0.1. | [optional] |
|**RequestId** | Pointer to **string** | Identifier of the individual request which triggered the activity. | [optional] |
|**QueueRefId** | Pointer to **int32** | Identifier of the provisioning queue reference. Use this value to group different events that were part of the same provisioning request. | [optional] |
|**Time** | Pointer to **string** | Combined date and time of the activity event in UTC. | [optional] |
|**TransactionId** | Pointer to **string** | Identifier of the transaction which triggered the activity. | [optional] |

## Methods

### NewGetByContractResponseHitsHitsSourceMeta

`func NewGetByContractResponseHitsHitsSourceMeta() *GetByContractResponseHitsHitsSourceMeta`

NewGetByContractResponseHitsHitsSourceMeta instantiates a new GetByContractResponseHitsHitsSourceMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetByContractResponseHitsHitsSourceMetaWithDefaults

`func NewGetByContractResponseHitsHitsSourceMetaWithDefaults() *GetByContractResponseHitsHitsSourceMeta`

NewGetByContractResponseHitsHitsSourceMetaWithDefaults instantiates a new GetByContractResponseHitsHitsSourceMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditVersion

`func (o *GetByContractResponseHitsHitsSourceMeta) GetAuditVersion() float32`

GetAuditVersion returns the AuditVersion field if non-nil, zero value otherwise.

### GetAuditVersionOk

`func (o *GetByContractResponseHitsHitsSourceMeta) GetAuditVersionOk() (*float32, bool)`

GetAuditVersionOk returns a tuple with the AuditVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditVersion

`func (o *GetByContractResponseHitsHitsSourceMeta) SetAuditVersion(v float32)`

SetAuditVersion sets AuditVersion field to given value.

### HasAuditVersion

`func (o *GetByContractResponseHitsHitsSourceMeta) HasAuditVersion() bool`

HasAuditVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *GetByContractResponseHitsHitsSourceMeta) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetByContractResponseHitsHitsSourceMeta) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetByContractResponseHitsHitsSourceMeta) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *GetByContractResponseHitsHitsSourceMeta) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetQueueRefId

`func (o *GetByContractResponseHitsHitsSourceMeta) GetQueueRefId() int32`

GetQueueRefId returns the QueueRefId field if non-nil, zero value otherwise.

### GetQueueRefIdOk

`func (o *GetByContractResponseHitsHitsSourceMeta) GetQueueRefIdOk() (*int32, bool)`

GetQueueRefIdOk returns a tuple with the QueueRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueRefId

`func (o *GetByContractResponseHitsHitsSourceMeta) SetQueueRefId(v int32)`

SetQueueRefId sets QueueRefId field to given value.

### HasQueueRefId

`func (o *GetByContractResponseHitsHitsSourceMeta) HasQueueRefId() bool`

HasQueueRefId returns a boolean if a field has been set.

### GetTime

`func (o *GetByContractResponseHitsHitsSourceMeta) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetByContractResponseHitsHitsSourceMeta) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetByContractResponseHitsHitsSourceMeta) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetByContractResponseHitsHitsSourceMeta) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTransactionId

`func (o *GetByContractResponseHitsHitsSourceMeta) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GetByContractResponseHitsHitsSourceMeta) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GetByContractResponseHitsHitsSourceMeta) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *GetByContractResponseHitsHitsSourceMeta) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


