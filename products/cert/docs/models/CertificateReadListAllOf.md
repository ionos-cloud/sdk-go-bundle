# CertificateReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of Certificate resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of Certificate resources. | |
|**Items** | Pointer to [**[]CertificateRead**](CertificateRead.md) | The list of Certificate resources. | [optional] |

## Methods

### NewCertificateReadListAllOf

`func NewCertificateReadListAllOf(id string, type_ string, href string, ) *CertificateReadListAllOf`

NewCertificateReadListAllOf instantiates a new CertificateReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateReadListAllOfWithDefaults

`func NewCertificateReadListAllOfWithDefaults() *CertificateReadListAllOf`

NewCertificateReadListAllOfWithDefaults instantiates a new CertificateReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CertificateReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *CertificateReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CertificateReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CertificateReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *CertificateReadListAllOf) GetItems() []CertificateRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CertificateReadListAllOf) GetItemsOk() (*[]CertificateRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CertificateReadListAllOf) SetItems(v []CertificateRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *CertificateReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


