# EvnItem

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ResourceType** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] |
|**ResourceUUID** | Pointer to **string** |  | [optional] |
|**IntervalMin** | Pointer to **int32** |  | [optional] |
|**IntervalDivisor** | Pointer to **int32** |  | [optional] |
|**From** | Pointer to [**time.Time**](time.Time.md) |  | [optional] |
|**To** | Pointer to [**time.Time**](time.Time.md) |  | [optional] |
|**ItemStub** | Pointer to **string** |  | [optional] |
|**Value** | Pointer to **float32** |  | [optional] |
|**ValueDivisor** | Pointer to **int32** |  | [optional] |
|**AdditionalParameters** | Pointer to **string** |  | [optional] |

## Methods

### NewEvnItem

`func NewEvnItem() *EvnItem`

NewEvnItem instantiates a new EvnItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvnItemWithDefaults

`func NewEvnItemWithDefaults() *EvnItem`

NewEvnItemWithDefaults instantiates a new EvnItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *EvnItem) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *EvnItem) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *EvnItem) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *EvnItem) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceUUID

`func (o *EvnItem) GetResourceUUID() string`

GetResourceUUID returns the ResourceUUID field if non-nil, zero value otherwise.

### GetResourceUUIDOk

`func (o *EvnItem) GetResourceUUIDOk() (*string, bool)`

GetResourceUUIDOk returns a tuple with the ResourceUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUUID

`func (o *EvnItem) SetResourceUUID(v string)`

SetResourceUUID sets ResourceUUID field to given value.

### HasResourceUUID

`func (o *EvnItem) HasResourceUUID() bool`

HasResourceUUID returns a boolean if a field has been set.

### GetIntervalMin

`func (o *EvnItem) GetIntervalMin() int32`

GetIntervalMin returns the IntervalMin field if non-nil, zero value otherwise.

### GetIntervalMinOk

`func (o *EvnItem) GetIntervalMinOk() (*int32, bool)`

GetIntervalMinOk returns a tuple with the IntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMin

`func (o *EvnItem) SetIntervalMin(v int32)`

SetIntervalMin sets IntervalMin field to given value.

### HasIntervalMin

`func (o *EvnItem) HasIntervalMin() bool`

HasIntervalMin returns a boolean if a field has been set.

### GetIntervalDivisor

`func (o *EvnItem) GetIntervalDivisor() int32`

GetIntervalDivisor returns the IntervalDivisor field if non-nil, zero value otherwise.

### GetIntervalDivisorOk

`func (o *EvnItem) GetIntervalDivisorOk() (*int32, bool)`

GetIntervalDivisorOk returns a tuple with the IntervalDivisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalDivisor

`func (o *EvnItem) SetIntervalDivisor(v int32)`

SetIntervalDivisor sets IntervalDivisor field to given value.

### HasIntervalDivisor

`func (o *EvnItem) HasIntervalDivisor() bool`

HasIntervalDivisor returns a boolean if a field has been set.

### GetFrom

`func (o *EvnItem) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EvnItem) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EvnItem) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EvnItem) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *EvnItem) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EvnItem) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EvnItem) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *EvnItem) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetItemStub

`func (o *EvnItem) GetItemStub() string`

GetItemStub returns the ItemStub field if non-nil, zero value otherwise.

### GetItemStubOk

`func (o *EvnItem) GetItemStubOk() (*string, bool)`

GetItemStubOk returns a tuple with the ItemStub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemStub

`func (o *EvnItem) SetItemStub(v string)`

SetItemStub sets ItemStub field to given value.

### HasItemStub

`func (o *EvnItem) HasItemStub() bool`

HasItemStub returns a boolean if a field has been set.

### GetValue

`func (o *EvnItem) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EvnItem) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EvnItem) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *EvnItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueDivisor

`func (o *EvnItem) GetValueDivisor() int32`

GetValueDivisor returns the ValueDivisor field if non-nil, zero value otherwise.

### GetValueDivisorOk

`func (o *EvnItem) GetValueDivisorOk() (*int32, bool)`

GetValueDivisorOk returns a tuple with the ValueDivisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDivisor

`func (o *EvnItem) SetValueDivisor(v int32)`

SetValueDivisor sets ValueDivisor field to given value.

### HasValueDivisor

`func (o *EvnItem) HasValueDivisor() bool`

HasValueDivisor returns a boolean if a field has been set.

### GetAdditionalParameters

`func (o *EvnItem) GetAdditionalParameters() string`

GetAdditionalParameters returns the AdditionalParameters field if non-nil, zero value otherwise.

### GetAdditionalParametersOk

`func (o *EvnItem) GetAdditionalParametersOk() (*string, bool)`

GetAdditionalParametersOk returns a tuple with the AdditionalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParameters

`func (o *EvnItem) SetAdditionalParameters(v string)`

SetAdditionalParameters sets AdditionalParameters field to given value.

### HasAdditionalParameters

`func (o *EvnItem) HasAdditionalParameters() bool`

HasAdditionalParameters returns a boolean if a field has been set.


