# TrafficMetadata

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ContractId** | Pointer to **string** |  | [optional] |
|**CustomerId** | Pointer to **string** |  | [optional] |
|**Reference** | Pointer to **NullableString** |  | [optional] |
|**Period** | Pointer to **string** |  | [optional] |
|**Unit** | Pointer to **string** |  | [optional] |

## Methods

### NewTrafficMetadata

`func NewTrafficMetadata() *TrafficMetadata`

NewTrafficMetadata instantiates a new TrafficMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficMetadataWithDefaults

`func NewTrafficMetadataWithDefaults() *TrafficMetadata`

NewTrafficMetadataWithDefaults instantiates a new TrafficMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *TrafficMetadata) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *TrafficMetadata) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *TrafficMetadata) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *TrafficMetadata) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetCustomerId

`func (o *TrafficMetadata) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TrafficMetadata) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TrafficMetadata) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *TrafficMetadata) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetReference

`func (o *TrafficMetadata) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *TrafficMetadata) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *TrafficMetadata) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *TrafficMetadata) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *TrafficMetadata) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *TrafficMetadata) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetPeriod

`func (o *TrafficMetadata) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TrafficMetadata) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TrafficMetadata) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *TrafficMetadata) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetUnit

`func (o *TrafficMetadata) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TrafficMetadata) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TrafficMetadata) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *TrafficMetadata) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


