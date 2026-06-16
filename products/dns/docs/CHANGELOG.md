# CHANGELOG

## 2.2.3 (June, 2026)

### Features:

- Updated to DNS API version 1.18.0
- Added 409 Conflict error handling for record create (`POST`) and update (`PUT`) operations
- Added debug logging for endpoint resolution, request failures, retry logic, and rate limiting

### Fixes:

- Fixed version and user-agent strings (removed path prefix, e.g. `products/dns/v2.2.2` → `2.2.3`)

## 2.2.2 (January, 2026)

### Fixes:

- Fixed `filter.recordIp` query parameter serialization in reverse records list endpoint

## 2.2.1 (October, 2025)

### Fixes:

- Updated dependencies (`shared` v0.1.6, `golang.org/x/oauth2` v0.30.0)

## 2.2.0 (March, 2025)

### Features:

- Added DNSSEC support:
    * New models: `Algorithm`, `KskBits`, `ZskBits`, `NsecMode`, `NsecParameters`, `KeyData`, `KeyParameters`, `DnssecKey`, `DnssecKeyCreate`, `DnssecKeyParameters`, `DnssecKeyReadCreation`, `DnssecKeyReadList`, `DnssecKeyReadListMetadata`, `DnssecKeyReadListProperties`, `DnssecKeyReadListPropertiesKeyParameters`, `DnssecKeyReadListPropertiesNsecParameters`
    * New models: `ZoneTransferPrimaryIpStatus`, `ZoneTransferPrimaryIpsStatus`, `MetadataForSecondaryZoneRecords`

## 2.1.0 (February, 2025)

### Features:

- Added `RecordType` model for strongly-typed DNS record type values

## 2.0.1 (December, 2024)

### Fixes:

- Bug fixes and stability improvements

## 2.0.0 (August, 2024)

### Features:

- New major version with updated models and APIs
- Full support for Zones, Records, Reverse Records, Secondary Zones, Zone Files, DNSSEC, and Quota APIs

## 1.1.1 (June, 2023)

### Fixes:

- Fixed date parser error

### Documentation:

- Fixed the number of returned params for delete functions

## 1.1.0 (June, 2023)

### Features:

- Renamed some models and methods:
    * `provisioningState` to `ProvisioningState`
    * `zonesResponse` to `ZoneReadList`
    * `zoneCreateRequest` to `ZoneCreate`
    * `zoneResponse` to `ZoneRead`
    * `zoneUpdateRequest` to `ZoneEnsure`
    * `recordsResponse` to `RecordReadList`
    * `recordCreateRequest` to `RecordCreate`
    * `recordResponse` to `RecordRead`
    * `recordUpdateRequest` to `RecordEnsure`
    * `RecordMetadata` to `MetadataWithStateFqdnZoneId`
    * `ErrorResponse` to `Error`
    * `ErrorMessage` to `ErrorMessages`
    * `ZoneResponseMetadata` to `MetadataWithStateNameservers`
    * `RecordProperties` to `Record`
    * `ZoneResponseProperties` to `Zone`
- Added new models:
    * `Links`
    * `Metadata`
    * `MetadataWithStateFqdnZoneIdAllOf`

### Documentation:

- Added documentation for Links, Metadata

**NOTE:** Cloud DNS is currently in the **Early Access (EA)** phase. We recommend keeping usage and testing to non-production critical applications. Please contact your sales representative or support for more information.
