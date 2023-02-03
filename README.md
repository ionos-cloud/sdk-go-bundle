# sdk-go-bundle
Enables users of ionocloud sdks to use one repo for all the GO sdks released

## Getting started

To start working with the SDK, setup your project for Go modules, and retrieve the SDK dependencies with `go get`. This example shows how to use the SDK
to make an API request to IONOS-CLOUD.

### Create directory

```bash
$ mkdir ~/firstbundleproj
$ cd ~/firstbundleproj
$ go mod init firstbundleproj
```
### Add SDK dependencies

```bash
$ go get github.com/ionos-cloud/sdk-go-bundle/common
$ go get github.com/ionos-cloud/sdk-go-bundle/products/compute
```

### Add the code: 
```golang
package main

import (
   "context"
   "fmt"
   "github.com/ionos-cloud/sdk-go-bundle/common"
   "github.com/ionos-cloud/sdk-go-bundle/products/compute"
   "log"
)

func main() {
   // NewConfigurationFromEnv looks for the following variables in the environment: IONOS_USERNAME, IONOS_PASSWORD, IONOS_TOKEN, IONOS_API_URL
   // You can either export IONOS_USERNAME and IONOS_PASSWORD, or IONOS_TOKEN
   // IONOS_API_URL - should be set only if you with to overwrite the default ionoscloud.DefaultIonosServerUrl
   cfg := common.NewConfigurationFromEnv()

   computeClient := compute.NewAPIClient(cfg)
   //setting Depth to 1 here makes sure we get the properties of the object(eg name)
   locations, apiResponse, err := computeClient.LocationsApi.LocationsGet(context.Background()).Depth(1).Execute()
   if err != nil {
      log.Fatalf("failed to list locations: %v", err)
   }
   apiResponse.LogInfo()
   if !locations.HasItems() || len(*locations.Items) == 0 {
      log.Fatalf("no locations found")
   }
   for _, location := range *locations.Items {
      if location.HasProperties() {
         fmt.Printf("location id is %s and location name %s\n", *location.Id, *location.Properties.Name)
      }
   }
}
```
### Run the code
```bash
go run .
[DEBUG] Request time : 601.664832ms for operation : LocationsGet
[DEBUG] response status code : 200
location id is de/fra and location name frankfurt
location id is us/las and location name lasvegas
location id is us/ewr and location name newark
location id is de/txl and location name berlin
location id is gb/lhr and location name london
location id is es/vit and location name logrono
location id is fr/par and location name paris

```

Using your IDE of choice, add the following code to `main.go`:

### Token Authentication
There are 2 ways to generate your token:

### Generate token using [sdk-go-auth](https://github.com/ionos-cloud/sdk-go-bundle/products/auth):
```golang
package main

import (
   "context"
   "github.com/ionos-cloud/sdk-go-bundle/common"
   "github.com/ionos-cloud/sdk-go-bundle/products/auth"
   "github.com/ionos-cloud/sdk-go-bundle/products/compute"
   "log"
)

func main() {
   //note: to use NewConfigurationFromEnv(), you need to previously set IONOS_USERNAME and IONOS_PASSWORD as env variables
   authClient := auth.NewAPIClient(common.NewConfigurationFromEnv())
   jwt, _, err := authClient.TokensApi.TokensGenerate(context.Background()).Execute()
   if err != nil {
      log.Fatalf("error occurred while generating token (%v)", err)
   }
   if !jwt.HasToken() {
      log.Fatalf("could not generate token")
   }
   cfg := common.NewConfiguration("", "", *jwt.GetToken(), "")
   apiClient := compute.NewAPIClient(cfg)
   datacenters, apiResponse, err := apiClient.DataCentersApi.DatacentersGet(context.Background()).Depth(1).Execute()
   if err != nil {
      log.Fatalf("error retrieving datacenters (%v)", err)
   }
   apiResponse.LogInfo()
   if !datacenters.HasItems() || len(*datacenters.Items) == 0 {
      log.Fatalf("no datacenters found, please create one first using SDK or DCD")
   }
   for _, datacenter := range *datacenters.Items {
      if datacenter.HasProperties() {
         log.Printf("datacenter name %s in location %s \n", *datacenter.Properties.Name, *datacenter.Properties.Location)
      }
   }
}

```
### Generate token using ionosctl:
Install ionosctl as explained [here](https://github.com/ionos-cloud/ionosctl)
Run commands to login and generate your token.
```golang
ionosctl login
ionosctl token generate
export IONOS_TOKEN="insert_here_token_saved_from_generate_command"
```
Save the generated token and use it to authenticate:
```golang
    import (
        "context"
        "fmt"
        "log"
		
        "github.com/ionos-cloud/sdk-go-bundle/common"
        "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    )

    func TokenAuthExample() error {
        //note: to use NewConfigurationFromEnv(), you need to previously set IONOS_TOKEN as env variables
        authClient := common.NewAPIClient(authApi.NewConfigurationFromEnv())
        apiClient := compute.NewAPIClient(cfg)
        datacenters, _, err := apiClient.DataCentersApi.DatacentersGet(context.Background()).Depth(1).Execute()
        if err != nil {
            return fmt.Errorf("error retrieving datacenters (%w)", err)
        }
        return nil
    }
```

## Certificate pinning:

You can enable certificate pinning if you want to bypass the normal certificate checking procedure,
by doing the following:

Set env variable IONOS_PINNED_CERT=<insert_sha256_public_fingerprint_here>

You can get the sha256 fingerprint most easily from the browser by inspecting the certificate.

## Debugging

You can now inject any logger that implements Printf as a logger
instead of using the default sdk logger.
There are now Loglevels that you can set: `Off`, `Debug` and `Trace`.
`Off` - does not show any logs
`Debug` - regular logs, no sensitive information
`Trace` - we recommend you only set this field for debugging purposes. Disable it in your production environments because it can log sensitive data.
It logs the full request and response without encryption, even for an HTTPS call. Verbose request and response logging can also significantly impact your application's performance.


```golang
package main

import (
   "github.com/ionos-cloud/sdk-go-bundle/common"
   "github.com/sirupsen/logrus"
)

func main() {
   // create your configuration. replace username, password, token and url with correct values, or use NewConfigurationFromEnv()
   // if you have set your env variables as explained above
   cfg := common.NewConfiguration("username", "password", "token", "hostUrl")
   // enable request and response logging. this is the most verbose loglevel
   cfg.LogLevel = common.Trace
   // inject your own logger that implements Printf
   cfg.Logger = logrus.New()
}
```

## Migration Guide

All sdks are found in the folder `https://github.com/ionos-cloud/sdk-go-bundle/tree/master/products`.
Structs moved to common package: `GenericOpenAPIError`, `Configuration`, `APIResponse`, utility functions, logging interface,  environment variables.
Import `github.com/ionos-cloud/sdk-go-bundle/common` and replace the structs as required.
IONOS_DEBUG removed, debugging now set with `IONOS_LOG_LEVEL` like described here: https://github.com/ionos-cloud/sdk-go#debugging
Replaced `PtrBool`, `PtrInt`, `PtrInt32`, `PtrInt64`, `PtrFloat`, `PtrFloat32`, `PtrFloat64`, `PtrString`, `PtrTime` with `ToPtr` generic functions

Example migration from  compute `github.com/ionos-cloud/sdk-go/v6` to `github.com/ionos-cloud/sdk-go-bundle/products/compute'

1. Replace the import ionoscloud `github.com/ionos-cloud/sdk-go/v6` with `github.com/ionos-cloud/sdk-go-bundle/products/compute` You will get errors, as some structs have been moved the the `common` package
2. Import `github.com/ionos-cloud/sdk-go-bundle/common` and replace the structs as required.
   Example replacements:
- var apiResponse *ionoscloud.APIResponse with var apiResponse *common.APIResponse
- ionoscloud.NewConfiguration(username, password, token, url) with common.NewConfiguration(username, password, token, url)
3. Replace the structs that were in the old repo with the ones in the new repo.

All replacements work as-is, no other changes are required.

  