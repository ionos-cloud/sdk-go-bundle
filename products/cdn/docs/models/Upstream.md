# Upstream

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Host** | **string** | The upstream host that handles the requests if not already cached. This host will be protected by the WAF if the option is enabled.  | |
|**Caching** | **bool** | Enable or disable caching. If enabled, the CDN will cache the responses from the upstream host. Subsequent requests for the same resource will be served from the cache.  | |
|**Waf** | **bool** | Enable or disable WAF to protect the upstream host. | |
|**GeoRestrictions** | Pointer to [**UpstreamGeoRestrictions**](UpstreamGeoRestrictions.md) |  | [optional] |
|**RateLimitClass** | **string** | Rate limit class that will be applied to limit the number of incoming requests per IP. | |
|**SniMode** | **string** | The SNI (Server Name Indication) mode of the upstream host. It supports two modes: - distribution: for outgoing connections to the upstream host, the CDN requires the upstream host to present a valid certificate that matches the configured domain of the CDN distribution. - origin: for outgoing connections to the upstream host, the CDN requires the upstream host to present a valid certificate that matches the configured upstream/origin hostname.  | |

## Methods

### NewUpstream

`func NewUpstream(host string, caching bool, waf bool, rateLimitClass string, sniMode string, ) *Upstream`

NewUpstream instantiates a new Upstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamWithDefaults

`func NewUpstreamWithDefaults() *Upstream`

NewUpstreamWithDefaults instantiates a new Upstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *Upstream) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Upstream) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Upstream) SetHost(v string)`

SetHost sets Host field to given value.


### GetCaching

`func (o *Upstream) GetCaching() bool`

GetCaching returns the Caching field if non-nil, zero value otherwise.

### GetCachingOk

`func (o *Upstream) GetCachingOk() (*bool, bool)`

GetCachingOk returns a tuple with the Caching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaching

`func (o *Upstream) SetCaching(v bool)`

SetCaching sets Caching field to given value.


### GetWaf

`func (o *Upstream) GetWaf() bool`

GetWaf returns the Waf field if non-nil, zero value otherwise.

### GetWafOk

`func (o *Upstream) GetWafOk() (*bool, bool)`

GetWafOk returns a tuple with the Waf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaf

`func (o *Upstream) SetWaf(v bool)`

SetWaf sets Waf field to given value.


### GetGeoRestrictions

`func (o *Upstream) GetGeoRestrictions() UpstreamGeoRestrictions`

GetGeoRestrictions returns the GeoRestrictions field if non-nil, zero value otherwise.

### GetGeoRestrictionsOk

`func (o *Upstream) GetGeoRestrictionsOk() (*UpstreamGeoRestrictions, bool)`

GetGeoRestrictionsOk returns a tuple with the GeoRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoRestrictions

`func (o *Upstream) SetGeoRestrictions(v UpstreamGeoRestrictions)`

SetGeoRestrictions sets GeoRestrictions field to given value.

### HasGeoRestrictions

`func (o *Upstream) HasGeoRestrictions() bool`

HasGeoRestrictions returns a boolean if a field has been set.

### GetRateLimitClass

`func (o *Upstream) GetRateLimitClass() string`

GetRateLimitClass returns the RateLimitClass field if non-nil, zero value otherwise.

### GetRateLimitClassOk

`func (o *Upstream) GetRateLimitClassOk() (*string, bool)`

GetRateLimitClassOk returns a tuple with the RateLimitClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitClass

`func (o *Upstream) SetRateLimitClass(v string)`

SetRateLimitClass sets RateLimitClass field to given value.


### GetSniMode

`func (o *Upstream) GetSniMode() string`

GetSniMode returns the SniMode field if non-nil, zero value otherwise.

### GetSniModeOk

`func (o *Upstream) GetSniModeOk() (*string, bool)`

GetSniModeOk returns a tuple with the SniMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSniMode

`func (o *Upstream) SetSniMode(v string)`

SetSniMode sets SniMode field to given value.



