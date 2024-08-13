# DnssecKeyReadListProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**KeyParameters** | [**DnssecKeyReadListPropertiesKeyParameters**](DnssecKeyReadListPropertiesKeyParameters.md) |  | |
|**NsecParameters** | [**DnssecKeyReadListPropertiesNsecParameters**](DnssecKeyReadListPropertiesNsecParameters.md) |  | |

## Methods

### NewDnssecKeyReadListProperties

`func NewDnssecKeyReadListProperties(keyParameters DnssecKeyReadListPropertiesKeyParameters, nsecParameters DnssecKeyReadListPropertiesNsecParameters, ) *DnssecKeyReadListProperties`

NewDnssecKeyReadListProperties instantiates a new DnssecKeyReadListProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnssecKeyReadListPropertiesWithDefaults

`func NewDnssecKeyReadListPropertiesWithDefaults() *DnssecKeyReadListProperties`

NewDnssecKeyReadListPropertiesWithDefaults instantiates a new DnssecKeyReadListProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyParameters

`func (o *DnssecKeyReadListProperties) GetKeyParameters() DnssecKeyReadListPropertiesKeyParameters`

GetKeyParameters returns the KeyParameters field if non-nil, zero value otherwise.

### GetKeyParametersOk

`func (o *DnssecKeyReadListProperties) GetKeyParametersOk() (*DnssecKeyReadListPropertiesKeyParameters, bool)`

GetKeyParametersOk returns a tuple with the KeyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyParameters

`func (o *DnssecKeyReadListProperties) SetKeyParameters(v DnssecKeyReadListPropertiesKeyParameters)`

SetKeyParameters sets KeyParameters field to given value.


### GetNsecParameters

`func (o *DnssecKeyReadListProperties) GetNsecParameters() DnssecKeyReadListPropertiesNsecParameters`

GetNsecParameters returns the NsecParameters field if non-nil, zero value otherwise.

### GetNsecParametersOk

`func (o *DnssecKeyReadListProperties) GetNsecParametersOk() (*DnssecKeyReadListPropertiesNsecParameters, bool)`

GetNsecParametersOk returns a tuple with the NsecParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsecParameters

`func (o *DnssecKeyReadListProperties) SetNsecParameters(v DnssecKeyReadListPropertiesNsecParameters)`

SetNsecParameters sets NsecParameters field to given value.



