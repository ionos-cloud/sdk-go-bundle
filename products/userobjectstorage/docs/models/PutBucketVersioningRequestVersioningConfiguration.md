# PutBucketVersioningRequestVersioningConfiguration

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Status** | Pointer to [**BucketVersioningStatus**](BucketVersioningStatus.md) |  | [optional] |
|**MfaDelete** | Pointer to [**MfaDeleteStatus**](MfaDeleteStatus.md) |  | [optional] |

## Methods

### NewPutBucketVersioningRequestVersioningConfiguration

`func NewPutBucketVersioningRequestVersioningConfiguration() *PutBucketVersioningRequestVersioningConfiguration`

NewPutBucketVersioningRequestVersioningConfiguration instantiates a new PutBucketVersioningRequestVersioningConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBucketVersioningRequestVersioningConfigurationWithDefaults

`func NewPutBucketVersioningRequestVersioningConfigurationWithDefaults() *PutBucketVersioningRequestVersioningConfiguration`

NewPutBucketVersioningRequestVersioningConfigurationWithDefaults instantiates a new PutBucketVersioningRequestVersioningConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PutBucketVersioningRequestVersioningConfiguration) GetStatus() BucketVersioningStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PutBucketVersioningRequestVersioningConfiguration) GetStatusOk() (*BucketVersioningStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PutBucketVersioningRequestVersioningConfiguration) SetStatus(v BucketVersioningStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PutBucketVersioningRequestVersioningConfiguration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMfaDelete

`func (o *PutBucketVersioningRequestVersioningConfiguration) GetMfaDelete() MfaDeleteStatus`

GetMfaDelete returns the MfaDelete field if non-nil, zero value otherwise.

### GetMfaDeleteOk

`func (o *PutBucketVersioningRequestVersioningConfiguration) GetMfaDeleteOk() (*MfaDeleteStatus, bool)`

GetMfaDeleteOk returns a tuple with the MfaDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaDelete

`func (o *PutBucketVersioningRequestVersioningConfiguration) SetMfaDelete(v MfaDeleteStatus)`

SetMfaDelete sets MfaDelete field to given value.

### HasMfaDelete

`func (o *PutBucketVersioningRequestVersioningConfiguration) HasMfaDelete() bool`

HasMfaDelete returns a boolean if a field has been set.


