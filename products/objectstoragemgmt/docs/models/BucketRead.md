# BucketRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The Bucket of the Bucket. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Bucket. | |
|**Metadata** | **map[string]interface{}** |  | |
|**Properties** | [**Bucket**](Bucket.md) |  | |

## Methods

### NewBucketRead

`func NewBucketRead(id string, type_ string, href string, metadata map[string]interface{}, properties Bucket, ) *BucketRead`

NewBucketRead instantiates a new BucketRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketReadWithDefaults

`func NewBucketReadWithDefaults() *BucketRead`

NewBucketReadWithDefaults instantiates a new BucketRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BucketRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BucketRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BucketRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *BucketRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BucketRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BucketRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *BucketRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BucketRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BucketRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *BucketRead) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BucketRead) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BucketRead) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *BucketRead) GetProperties() Bucket`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BucketRead) GetPropertiesOk() (*Bucket, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BucketRead) SetProperties(v Bucket)`

SetProperties sets Properties field to given value.



