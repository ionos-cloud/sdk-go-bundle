# KeyParameters

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Algorithm** | [**Algorithm**](Algorithm.md) |  | |
|**KskBits** | [**KskBits**](KskBits.md) |  | |
|**ZskBits** | [**ZskBits**](ZskBits.md) |  | |

## Methods

### NewKeyParameters

`func NewKeyParameters(algorithm Algorithm, kskBits KskBits, zskBits ZskBits, ) *KeyParameters`

NewKeyParameters instantiates a new KeyParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyParametersWithDefaults

`func NewKeyParametersWithDefaults() *KeyParameters`

NewKeyParametersWithDefaults instantiates a new KeyParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *KeyParameters) GetAlgorithm() Algorithm`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *KeyParameters) GetAlgorithmOk() (*Algorithm, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *KeyParameters) SetAlgorithm(v Algorithm)`

SetAlgorithm sets Algorithm field to given value.


### GetKskBits

`func (o *KeyParameters) GetKskBits() KskBits`

GetKskBits returns the KskBits field if non-nil, zero value otherwise.

### GetKskBitsOk

`func (o *KeyParameters) GetKskBitsOk() (*KskBits, bool)`

GetKskBitsOk returns a tuple with the KskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskBits

`func (o *KeyParameters) SetKskBits(v KskBits)`

SetKskBits sets KskBits field to given value.


### GetZskBits

`func (o *KeyParameters) GetZskBits() ZskBits`

GetZskBits returns the ZskBits field if non-nil, zero value otherwise.

### GetZskBitsOk

`func (o *KeyParameters) GetZskBitsOk() (*ZskBits, bool)`

GetZskBitsOk returns a tuple with the ZskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZskBits

`func (o *KeyParameters) SetZskBits(v ZskBits)`

SetZskBits sets ZskBits field to given value.



