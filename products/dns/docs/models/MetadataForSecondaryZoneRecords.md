# MetadataForSecondaryZoneRecords

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Fqdn** | **string** | A fully qualified domain name. FQDN consists of two parts - the hostname and the domain name. | [readonly] |
|**ZoneId** | **string** | The ID (UUID) of the DNS zone of which record belongs to. | [readonly] |
|**RootName** | **string** | Indicates the root name (from the primary zone) for the record | |

## Methods

### NewMetadataForSecondaryZoneRecords

`func NewMetadataForSecondaryZoneRecords(fqdn string, zoneId string, rootName string, ) *MetadataForSecondaryZoneRecords`

NewMetadataForSecondaryZoneRecords instantiates a new MetadataForSecondaryZoneRecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataForSecondaryZoneRecordsWithDefaults

`func NewMetadataForSecondaryZoneRecordsWithDefaults() *MetadataForSecondaryZoneRecords`

NewMetadataForSecondaryZoneRecordsWithDefaults instantiates a new MetadataForSecondaryZoneRecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *MetadataForSecondaryZoneRecords) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *MetadataForSecondaryZoneRecords) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *MetadataForSecondaryZoneRecords) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetZoneId

`func (o *MetadataForSecondaryZoneRecords) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *MetadataForSecondaryZoneRecords) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *MetadataForSecondaryZoneRecords) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.


### GetRootName

`func (o *MetadataForSecondaryZoneRecords) GetRootName() string`

GetRootName returns the RootName field if non-nil, zero value otherwise.

### GetRootNameOk

`func (o *MetadataForSecondaryZoneRecords) GetRootNameOk() (*string, bool)`

GetRootNameOk returns a tuple with the RootName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootName

`func (o *MetadataForSecondaryZoneRecords) SetRootName(v string)`

SetRootName sets RootName field to given value.



