# IPSecPSK

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Key** | **string** | The Pre-Shared Key used for IPSec Authentication. | |

## Methods

### NewIPSecPSK

`func NewIPSecPSK(key string, ) *IPSecPSK`

NewIPSecPSK instantiates a new IPSecPSK object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecPSKWithDefaults

`func NewIPSecPSKWithDefaults() *IPSecPSK`

NewIPSecPSKWithDefaults instantiates a new IPSecPSK object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *IPSecPSK) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IPSecPSK) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IPSecPSK) SetKey(v string)`

SetKey sets Key field to given value.



