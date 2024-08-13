# DnssecKeyReadCreation

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**DnssecKeyParameters**](DnssecKeyParameters.md) |  | |
|**Id** | **string** |  | |
|**Type** | **string** |  | |
|**Href** | **string** |  | [readonly] |

## Methods

### NewDnssecKeyReadCreation

`func NewDnssecKeyReadCreation(properties DnssecKeyParameters, id string, type_ string, href string, ) *DnssecKeyReadCreation`

NewDnssecKeyReadCreation instantiates a new DnssecKeyReadCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnssecKeyReadCreationWithDefaults

`func NewDnssecKeyReadCreationWithDefaults() *DnssecKeyReadCreation`

NewDnssecKeyReadCreationWithDefaults instantiates a new DnssecKeyReadCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *DnssecKeyReadCreation) GetProperties() DnssecKeyParameters`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DnssecKeyReadCreation) GetPropertiesOk() (*DnssecKeyParameters, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DnssecKeyReadCreation) SetProperties(v DnssecKeyParameters)`

SetProperties sets Properties field to given value.


### GetId

`func (o *DnssecKeyReadCreation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnssecKeyReadCreation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnssecKeyReadCreation) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *DnssecKeyReadCreation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnssecKeyReadCreation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnssecKeyReadCreation) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *DnssecKeyReadCreation) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DnssecKeyReadCreation) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DnssecKeyReadCreation) SetHref(v string)`

SetHref sets Href field to given value.



