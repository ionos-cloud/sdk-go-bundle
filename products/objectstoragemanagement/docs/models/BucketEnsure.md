# BucketEnsure

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The Bucket of the Bucket. | |
|**Metadata** | Pointer to **map[string]interface{}** | Metadata | [optional] |
|**Properties** | [**Bucket**](Bucket.md) |  | |

## Methods

### NewBucketEnsure

`func NewBucketEnsure(id string, properties Bucket, ) *BucketEnsure`

NewBucketEnsure instantiates a new BucketEnsure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketEnsureWithDefaults

`func NewBucketEnsureWithDefaults() *BucketEnsure`

NewBucketEnsureWithDefaults instantiates a new BucketEnsure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BucketEnsure) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BucketEnsure) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BucketEnsure) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *BucketEnsure) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BucketEnsure) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BucketEnsure) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BucketEnsure) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *BucketEnsure) GetProperties() Bucket`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BucketEnsure) GetPropertiesOk() (*Bucket, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BucketEnsure) SetProperties(v Bucket)`

SetProperties sets Properties field to given value.



