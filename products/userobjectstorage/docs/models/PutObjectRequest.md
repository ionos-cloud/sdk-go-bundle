# PutObjectRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Body** | Pointer to **string** | Object data. | [optional] |
|**XAmzMeta** | Pointer to **map[string]string** | A map of metadata to store with the object. | [optional] |

## Methods

### NewPutObjectRequest

`func NewPutObjectRequest() *PutObjectRequest`

NewPutObjectRequest instantiates a new PutObjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutObjectRequestWithDefaults

`func NewPutObjectRequestWithDefaults() *PutObjectRequest`

NewPutObjectRequestWithDefaults instantiates a new PutObjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *PutObjectRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PutObjectRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PutObjectRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PutObjectRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetXAmzMeta

`func (o *PutObjectRequest) GetXAmzMeta() map[string]string`

GetXAmzMeta returns the XAmzMeta field if non-nil, zero value otherwise.

### GetXAmzMetaOk

`func (o *PutObjectRequest) GetXAmzMetaOk() (*map[string]string, bool)`

GetXAmzMetaOk returns a tuple with the XAmzMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAmzMeta

`func (o *PutObjectRequest) SetXAmzMeta(v map[string]string)`

SetXAmzMeta sets XAmzMeta field to given value.

### HasXAmzMeta

`func (o *PutObjectRequest) HasXAmzMeta() bool`

HasXAmzMeta returns a boolean if a field has been set.


