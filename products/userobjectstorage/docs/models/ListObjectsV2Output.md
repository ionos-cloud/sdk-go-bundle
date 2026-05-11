# ListObjectsV2Output

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The bucket name. | |
|**Prefix** | **string** | Object key prefix that identifies one or more objects to which this rule applies. Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. | |
|**MaxKeys** | **int32** | The maximum number of keys returned in the response. By default the operation returns up to 1000 key names. The response might contain fewer keys but will never contain more. | [default to 1000]|
|**IsTruncated** | **bool** | A flag that indicates whether IONOS Object Storage returned all of the results that satisfied the search criteria. If your results were truncated, you can make a follow-up paginated request using the NextKeyMarker and NextVersionIdMarker response parameters as a starting place in another request to return the rest of the results. | |
|**KeyCount** | **int32** |  | |
|**Contents** | [**[]Object**](Object.md) | Metadata about each object returned. | |
|**Delimiter** | Pointer to **string** |  | [optional] |
|**CommonPrefixes** | Pointer to [**[]CommonPrefix**](CommonPrefix.md) | All of the keys rolled up into a common prefix count as a single return when calculating the number of returns. | [optional] |
|**EncodingType** | Pointer to [**EncodingType**](EncodingType.md) |  | [optional] |
|**ContinuationToken** | Pointer to **string** | If ContinuationToken was sent with the request, it is included in the response. | [optional] |
|**NextContinuationToken** | Pointer to **string** | &#x60;NextContinuationToken&#x60; is sent when &#x60;isTruncated&#x60; is true. Use it to continue the listing.  | [optional] |
|**StartAfter** | Pointer to **string** | If StartAfter was sent with the request, it is included in the response. | [optional] |

## Methods

### NewListObjectsV2Output

`func NewListObjectsV2Output(name string, prefix string, maxKeys int32, isTruncated bool, keyCount int32, contents []Object, ) *ListObjectsV2Output`

NewListObjectsV2Output instantiates a new ListObjectsV2Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListObjectsV2OutputWithDefaults

`func NewListObjectsV2OutputWithDefaults() *ListObjectsV2Output`

NewListObjectsV2OutputWithDefaults instantiates a new ListObjectsV2Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListObjectsV2Output) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListObjectsV2Output) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListObjectsV2Output) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *ListObjectsV2Output) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ListObjectsV2Output) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ListObjectsV2Output) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetMaxKeys

`func (o *ListObjectsV2Output) GetMaxKeys() int32`

GetMaxKeys returns the MaxKeys field if non-nil, zero value otherwise.

### GetMaxKeysOk

`func (o *ListObjectsV2Output) GetMaxKeysOk() (*int32, bool)`

GetMaxKeysOk returns a tuple with the MaxKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxKeys

`func (o *ListObjectsV2Output) SetMaxKeys(v int32)`

SetMaxKeys sets MaxKeys field to given value.


### GetIsTruncated

`func (o *ListObjectsV2Output) GetIsTruncated() bool`

GetIsTruncated returns the IsTruncated field if non-nil, zero value otherwise.

### GetIsTruncatedOk

`func (o *ListObjectsV2Output) GetIsTruncatedOk() (*bool, bool)`

GetIsTruncatedOk returns a tuple with the IsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTruncated

`func (o *ListObjectsV2Output) SetIsTruncated(v bool)`

SetIsTruncated sets IsTruncated field to given value.


### GetKeyCount

`func (o *ListObjectsV2Output) GetKeyCount() int32`

GetKeyCount returns the KeyCount field if non-nil, zero value otherwise.

### GetKeyCountOk

`func (o *ListObjectsV2Output) GetKeyCountOk() (*int32, bool)`

GetKeyCountOk returns a tuple with the KeyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyCount

`func (o *ListObjectsV2Output) SetKeyCount(v int32)`

SetKeyCount sets KeyCount field to given value.


### GetContents

`func (o *ListObjectsV2Output) GetContents() []Object`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ListObjectsV2Output) GetContentsOk() (*[]Object, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ListObjectsV2Output) SetContents(v []Object)`

SetContents sets Contents field to given value.


### GetDelimiter

`func (o *ListObjectsV2Output) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *ListObjectsV2Output) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *ListObjectsV2Output) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *ListObjectsV2Output) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### GetCommonPrefixes

`func (o *ListObjectsV2Output) GetCommonPrefixes() []CommonPrefix`

GetCommonPrefixes returns the CommonPrefixes field if non-nil, zero value otherwise.

### GetCommonPrefixesOk

`func (o *ListObjectsV2Output) GetCommonPrefixesOk() (*[]CommonPrefix, bool)`

GetCommonPrefixesOk returns a tuple with the CommonPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonPrefixes

`func (o *ListObjectsV2Output) SetCommonPrefixes(v []CommonPrefix)`

SetCommonPrefixes sets CommonPrefixes field to given value.

### HasCommonPrefixes

`func (o *ListObjectsV2Output) HasCommonPrefixes() bool`

HasCommonPrefixes returns a boolean if a field has been set.

### GetEncodingType

`func (o *ListObjectsV2Output) GetEncodingType() EncodingType`

GetEncodingType returns the EncodingType field if non-nil, zero value otherwise.

### GetEncodingTypeOk

`func (o *ListObjectsV2Output) GetEncodingTypeOk() (*EncodingType, bool)`

GetEncodingTypeOk returns a tuple with the EncodingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingType

`func (o *ListObjectsV2Output) SetEncodingType(v EncodingType)`

SetEncodingType sets EncodingType field to given value.

### HasEncodingType

`func (o *ListObjectsV2Output) HasEncodingType() bool`

HasEncodingType returns a boolean if a field has been set.

### GetContinuationToken

`func (o *ListObjectsV2Output) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *ListObjectsV2Output) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *ListObjectsV2Output) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *ListObjectsV2Output) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### GetNextContinuationToken

`func (o *ListObjectsV2Output) GetNextContinuationToken() string`

GetNextContinuationToken returns the NextContinuationToken field if non-nil, zero value otherwise.

### GetNextContinuationTokenOk

`func (o *ListObjectsV2Output) GetNextContinuationTokenOk() (*string, bool)`

GetNextContinuationTokenOk returns a tuple with the NextContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextContinuationToken

`func (o *ListObjectsV2Output) SetNextContinuationToken(v string)`

SetNextContinuationToken sets NextContinuationToken field to given value.

### HasNextContinuationToken

`func (o *ListObjectsV2Output) HasNextContinuationToken() bool`

HasNextContinuationToken returns a boolean if a field has been set.

### GetStartAfter

`func (o *ListObjectsV2Output) GetStartAfter() string`

GetStartAfter returns the StartAfter field if non-nil, zero value otherwise.

### GetStartAfterOk

`func (o *ListObjectsV2Output) GetStartAfterOk() (*string, bool)`

GetStartAfterOk returns a tuple with the StartAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAfter

`func (o *ListObjectsV2Output) SetStartAfter(v string)`

SetStartAfter sets StartAfter field to given value.

### HasStartAfter

`func (o *ListObjectsV2Output) HasStartAfter() bool`

HasStartAfter returns a boolean if a field has been set.


