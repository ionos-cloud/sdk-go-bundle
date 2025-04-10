# PutObjectRetentionRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Mode** | Pointer to **string** | Indicates the Retention mode for the specified object. | [optional] |
|**RetainUntilDate** | Pointer to **string** | The date on which this Object Lock Retention will expire. | [optional] |

## Methods

### NewPutObjectRetentionRequest

`func NewPutObjectRetentionRequest() *PutObjectRetentionRequest`

NewPutObjectRetentionRequest instantiates a new PutObjectRetentionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutObjectRetentionRequestWithDefaults

`func NewPutObjectRetentionRequestWithDefaults() *PutObjectRetentionRequest`

NewPutObjectRetentionRequestWithDefaults instantiates a new PutObjectRetentionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *PutObjectRetentionRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PutObjectRetentionRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PutObjectRetentionRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PutObjectRetentionRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRetainUntilDate

`func (o *PutObjectRetentionRequest) GetRetainUntilDate() string`

GetRetainUntilDate returns the RetainUntilDate field if non-nil, zero value otherwise.

### GetRetainUntilDateOk

`func (o *PutObjectRetentionRequest) GetRetainUntilDateOk() (*string, bool)`

GetRetainUntilDateOk returns a tuple with the RetainUntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainUntilDate

`func (o *PutObjectRetentionRequest) SetRetainUntilDate(v string)`

SetRetainUntilDate sets RetainUntilDate field to given value.

### HasRetainUntilDate

`func (o *PutObjectRetentionRequest) HasRetainUntilDate() bool`

HasRetainUntilDate returns a boolean if a field has been set.


