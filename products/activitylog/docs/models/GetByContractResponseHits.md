# GetByContractResponseHits

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Total** | Pointer to **int32** | A count of total number of available activity events. | [optional] |
|**Hits** | Pointer to [**[]GetByContractResponseHitsHits**](GetByContractResponseHitsHits.md) |  | [optional] |

## Methods

### NewGetByContractResponseHits

`func NewGetByContractResponseHits() *GetByContractResponseHits`

NewGetByContractResponseHits instantiates a new GetByContractResponseHits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetByContractResponseHitsWithDefaults

`func NewGetByContractResponseHitsWithDefaults() *GetByContractResponseHits`

NewGetByContractResponseHitsWithDefaults instantiates a new GetByContractResponseHits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetByContractResponseHits) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetByContractResponseHits) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetByContractResponseHits) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetByContractResponseHits) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetHits

`func (o *GetByContractResponseHits) GetHits() []GetByContractResponseHitsHits`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *GetByContractResponseHits) GetHitsOk() (*[]GetByContractResponseHitsHits, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *GetByContractResponseHits) SetHits(v []GetByContractResponseHitsHits)`

SetHits sets Hits field to given value.

### HasHits

`func (o *GetByContractResponseHits) HasHits() bool`

HasHits returns a boolean if a field has been set.


