# ReverseRecord

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The reverse DNS record name | |
|**Description** | Pointer to **string** | Description stored along with the reverse DNS record to describe its usage. | [optional] |
|**Ip** | **string** | Specifies for which IP address the reverse record should be created. The IP addresses needs to be owned by the contract. Accepts IPv4 and IPv6 addresses. | |

## Methods

### NewReverseRecord

`func NewReverseRecord(name string, ip string, ) *ReverseRecord`

NewReverseRecord instantiates a new ReverseRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseRecordWithDefaults

`func NewReverseRecordWithDefaults() *ReverseRecord`

NewReverseRecordWithDefaults instantiates a new ReverseRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReverseRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReverseRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReverseRecord) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ReverseRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReverseRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReverseRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReverseRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIp

`func (o *ReverseRecord) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ReverseRecord) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ReverseRecord) SetIp(v string)`

SetIp sets Ip field to given value.



