# GetByContractResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Hits** | Pointer to [**GetByContractResponseHits**](GetByContractResponseHits.md) |  | [optional] |

## Methods

### NewGetByContractResponse

`func NewGetByContractResponse() *GetByContractResponse`

NewGetByContractResponse instantiates a new GetByContractResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetByContractResponseWithDefaults

`func NewGetByContractResponseWithDefaults() *GetByContractResponse`

NewGetByContractResponseWithDefaults instantiates a new GetByContractResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *GetByContractResponse) GetHits() GetByContractResponseHits`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *GetByContractResponse) GetHitsOk() (*GetByContractResponseHits, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *GetByContractResponse) SetHits(v GetByContractResponseHits)`

SetHits sets Hits field to given value.

### HasHits

`func (o *GetByContractResponse) HasHits() bool`

HasHits returns a boolean if a field has been set.


