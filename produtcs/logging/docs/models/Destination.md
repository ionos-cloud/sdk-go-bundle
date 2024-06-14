# Destination

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | Pointer to **string** | The internal output stream to send logs to | [optional] [default to "loki"]|
|**RetentionInDays** | Pointer to **int32** | defines the number of days a log record should be kept in loki. Works with loki destination type only. | [optional] [default to 30]|

## Methods

### NewDestination

`func NewDestination() *Destination`

NewDestination instantiates a new Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationWithDefaults

`func NewDestinationWithDefaults() *Destination`

NewDestinationWithDefaults instantiates a new Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Destination) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Destination) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Destination) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Destination) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRetentionInDays

`func (o *Destination) GetRetentionInDays() int32`

GetRetentionInDays returns the RetentionInDays field if non-nil, zero value otherwise.

### GetRetentionInDaysOk

`func (o *Destination) GetRetentionInDaysOk() (*int32, bool)`

GetRetentionInDaysOk returns a tuple with the RetentionInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionInDays

`func (o *Destination) SetRetentionInDays(v int32)`

SetRetentionInDays sets RetentionInDays field to given value.

### HasRetentionInDays

`func (o *Destination) HasRetentionInDays() bool`

HasRetentionInDays returns a boolean if a field has been set.


