# CHANGELOG

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