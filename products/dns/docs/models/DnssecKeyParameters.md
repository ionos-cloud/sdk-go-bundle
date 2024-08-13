# DnssecKeyParameters

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**KeyParameters** | [**KeyParameters**](KeyParameters.md) |  | |
|**NsecParameters** | [**NsecParameters**](NsecParameters.md) |  | |
|**Validity** | **int32** | Signature validity in days  | |

## Methods

### NewDnssecKeyParameters

`func NewDnssecKeyParameters(keyParameters KeyParameters, nsecParameters NsecParameters, validity int32, ) *DnssecKeyParameters`

NewDnssecKeyParameters instantiates a new DnssecKeyParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnssecKeyParametersWithDefaults

`func NewDnssecKeyParametersWithDefaults() *DnssecKeyParameters`

NewDnssecKeyParametersWithDefaults instantiates a new DnssecKeyParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyParameters

`func (o *DnssecKeyParameters) GetKeyParameters() KeyParameters`

GetKeyParameters returns the KeyParameters field if non-nil, zero value otherwise.

### GetKeyParametersOk

`func (o *DnssecKeyParameters) GetKeyParametersOk() (*KeyParameters, bool)`

GetKeyParametersOk returns a tuple with the KeyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyParameters

`func (o *DnssecKeyParameters) SetKeyParameters(v KeyParameters)`

SetKeyParameters sets KeyParameters field to given value.


### GetNsecParameters

`func (o *DnssecKeyParameters) GetNsecParameters() NsecParameters`

GetNsecParameters returns the NsecParameters field if non-nil, zero value otherwise.

### GetNsecParametersOk

`func (o *DnssecKeyParameters) GetNsecParametersOk() (*NsecParameters, bool)`

GetNsecParametersOk returns a tuple with the NsecParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsecParameters

`func (o *DnssecKeyParameters) SetNsecParameters(v NsecParameters)`

SetNsecParameters sets NsecParameters field to given value.


### GetValidity

`func (o *DnssecKeyParameters) GetValidity() int32`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *DnssecKeyParameters) GetValidityOk() (*int32, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *DnssecKeyParameters) SetValidity(v int32)`

SetValidity sets Validity field to given value.



