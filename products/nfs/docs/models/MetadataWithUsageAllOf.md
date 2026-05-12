# MetadataWithUsageAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Used** | **int64** | The storage capacity currently taken up by the customer&#39;s files in bytes, or 0 on error.  | [readonly] |

## Methods

### NewMetadataWithUsageAllOf

`func NewMetadataWithUsageAllOf(used int64, ) *MetadataWithUsageAllOf`

NewMetadataWithUsageAllOf instantiates a new MetadataWithUsageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithUsageAllOfWithDefaults

`func NewMetadataWithUsageAllOfWithDefaults() *MetadataWithUsageAllOf`

NewMetadataWithUsageAllOfWithDefaults instantiates a new MetadataWithUsageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsed

`func (o *MetadataWithUsageAllOf) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *MetadataWithUsageAllOf) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *MetadataWithUsageAllOf) SetUsed(v int64)`

SetUsed sets Used field to given value.



