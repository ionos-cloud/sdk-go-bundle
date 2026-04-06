# UsageMeterQuantity

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Quantity** | Pointer to **string** |  | [optional] |
|**Unit** | Pointer to **string** |  | [optional] |

## Methods

### NewUsageMeterQuantity

`func NewUsageMeterQuantity() *UsageMeterQuantity`

NewUsageMeterQuantity instantiates a new UsageMeterQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMeterQuantityWithDefaults

`func NewUsageMeterQuantityWithDefaults() *UsageMeterQuantity`

NewUsageMeterQuantityWithDefaults instantiates a new UsageMeterQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *UsageMeterQuantity) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UsageMeterQuantity) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UsageMeterQuantity) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UsageMeterQuantity) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnit

`func (o *UsageMeterQuantity) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *UsageMeterQuantity) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *UsageMeterQuantity) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *UsageMeterQuantity) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


