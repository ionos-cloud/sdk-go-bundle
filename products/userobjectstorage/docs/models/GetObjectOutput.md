# GetObjectOutput

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Body** | Pointer to [***os.File**](*os.File.md) | Object data. | [optional] |
|**Metadata** | Pointer to [**map[string]Metadata1**](Metadata1.md) | A map of metadata to store with the object. Each key must start with  &#x60;x-amz-meta-&#x60; prefix.  | [optional] |

## Methods

### NewGetObjectOutput

`func NewGetObjectOutput() *GetObjectOutput`

NewGetObjectOutput instantiates a new GetObjectOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectOutputWithDefaults

`func NewGetObjectOutputWithDefaults() *GetObjectOutput`

NewGetObjectOutputWithDefaults instantiates a new GetObjectOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *GetObjectOutput) GetBody() *os.File`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GetObjectOutput) GetBodyOk() (**os.File, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GetObjectOutput) SetBody(v *os.File)`

SetBody sets Body field to given value.

### HasBody

`func (o *GetObjectOutput) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetMetadata

`func (o *GetObjectOutput) GetMetadata() map[string]Metadata1`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetObjectOutput) GetMetadataOk() (*map[string]Metadata1, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetObjectOutput) SetMetadata(v map[string]Metadata1)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetObjectOutput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


