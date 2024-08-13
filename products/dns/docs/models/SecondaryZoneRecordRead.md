# SecondaryZoneRecordRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | **string** |  | |
|**Metadata** | [**MetadataForSecondaryZoneRecords**](MetadataForSecondaryZoneRecords.md) |  | |
|**Properties** | [**Record**](Record.md) |  | |

## Methods

### NewSecondaryZoneRecordRead

`func NewSecondaryZoneRecordRead(type_ string, metadata MetadataForSecondaryZoneRecords, properties Record, ) *SecondaryZoneRecordRead`

NewSecondaryZoneRecordRead instantiates a new SecondaryZoneRecordRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecondaryZoneRecordReadWithDefaults

`func NewSecondaryZoneRecordReadWithDefaults() *SecondaryZoneRecordRead`

NewSecondaryZoneRecordReadWithDefaults instantiates a new SecondaryZoneRecordRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SecondaryZoneRecordRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecondaryZoneRecordRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecondaryZoneRecordRead) SetType(v string)`

SetType sets Type field to given value.


### GetMetadata

`func (o *SecondaryZoneRecordRead) GetMetadata() MetadataForSecondaryZoneRecords`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecondaryZoneRecordRead) GetMetadataOk() (*MetadataForSecondaryZoneRecords, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecondaryZoneRecordRead) SetMetadata(v MetadataForSecondaryZoneRecords)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *SecondaryZoneRecordRead) GetProperties() Record`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SecondaryZoneRecordRead) GetPropertiesOk() (*Record, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SecondaryZoneRecordRead) SetProperties(v Record)`

SetProperties sets Properties field to given value.



