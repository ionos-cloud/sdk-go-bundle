# EvnMetadata

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ContractId** | Pointer to **string** |  | [optional] |
|**CustomerId** | Pointer to **string** |  | [optional] |
|**Reference** | Pointer to **NullableString** |  | [optional] |
|**Period** | Pointer to **string** |  | [optional] |

## Methods

### NewEvnMetadata

`func NewEvnMetadata() *EvnMetadata`

NewEvnMetadata instantiates a new EvnMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvnMetadataWithDefaults

`func NewEvnMetadataWithDefaults() *EvnMetadata`

NewEvnMetadataWithDefaults instantiates a new EvnMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *EvnMetadata) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *EvnMetadata) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *EvnMetadata) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *EvnMetadata) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetCustomerId

`func (o *EvnMetadata) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *EvnMetadata) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *EvnMetadata) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *EvnMetadata) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetReference

`func (o *EvnMetadata) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *EvnMetadata) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *EvnMetadata) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *EvnMetadata) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *EvnMetadata) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *EvnMetadata) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetPeriod

`func (o *EvnMetadata) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EvnMetadata) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EvnMetadata) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *EvnMetadata) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


