# NsecParameters

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**NsecMode** | [**NsecMode**](NsecMode.md) |  | |
|**Nsec3Iterations** | **int32** | Number of iterations for NSEC3. (between 0 and 50)  | |
|**Nsec3SaltBits** | **int32** | Salt length in bits for NSEC3. (between 64 and 128, multiples of 8)  | |

## Methods

### NewNsecParameters

`func NewNsecParameters(nsecMode NsecMode, nsec3Iterations int32, nsec3SaltBits int32, ) *NsecParameters`

NewNsecParameters instantiates a new NsecParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsecParametersWithDefaults

`func NewNsecParametersWithDefaults() *NsecParameters`

NewNsecParametersWithDefaults instantiates a new NsecParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNsecMode

`func (o *NsecParameters) GetNsecMode() NsecMode`

GetNsecMode returns the NsecMode field if non-nil, zero value otherwise.

### GetNsecModeOk

`func (o *NsecParameters) GetNsecModeOk() (*NsecMode, bool)`

GetNsecModeOk returns a tuple with the NsecMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsecMode

`func (o *NsecParameters) SetNsecMode(v NsecMode)`

SetNsecMode sets NsecMode field to given value.


### GetNsec3Iterations

`func (o *NsecParameters) GetNsec3Iterations() int32`

GetNsec3Iterations returns the Nsec3Iterations field if non-nil, zero value otherwise.

### GetNsec3IterationsOk

`func (o *NsecParameters) GetNsec3IterationsOk() (*int32, bool)`

GetNsec3IterationsOk returns a tuple with the Nsec3Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsec3Iterations

`func (o *NsecParameters) SetNsec3Iterations(v int32)`

SetNsec3Iterations sets Nsec3Iterations field to given value.


### GetNsec3SaltBits

`func (o *NsecParameters) GetNsec3SaltBits() int32`

GetNsec3SaltBits returns the Nsec3SaltBits field if non-nil, zero value otherwise.

### GetNsec3SaltBitsOk

`func (o *NsecParameters) GetNsec3SaltBitsOk() (*int32, bool)`

GetNsec3SaltBitsOk returns a tuple with the Nsec3SaltBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsec3SaltBits

`func (o *NsecParameters) SetNsec3SaltBits(v int32)`

SetNsec3SaltBits sets Nsec3SaltBits field to given value.



