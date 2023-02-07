# sdk-go-bundle
Enables users of IONOS Cloud sdks to use one repo for all the GO SDKs released.

# Go API client for ionoscloud

IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.

Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.

## Overview
The IONOS Cloud SDK for GO bundle provides you with access to the IONOS Cloud API. The client library supports both simple and complex requests.
It is designed for developers who are building applications in GO . The SDK for GO wraps the IONOS Cloud API. All API operations are performed over SSL and authenticated using your IONOS Cloud portal credentials.
The API can be accessed within an instance running in IONOS Cloud or directly over the Internet from any application that can send an HTTPS request and receive an HTTPS response.

## Environment Variables

| Environment Variable | Description                                                                                                                                                                                                                    |
|----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `IONOS_USERNAME`     | Specify the username used to login, to authenticate against the IONOS Cloud API                                                                                                                                                |
| `IONOS_PASSWORD`     | Specify the password used to login, to authenticate against the IONOS Cloud API                                                                                                                                                |
| `IONOS_TOKEN`        | Specify the token used to login, if a token is being used instead of username and password                                                                                                                                     |
| `IONOS_API_URL`      | Specify the API URL. It will overwrite the API endpoint default value `api.ionos.com`. Note: the host URL does not contain the `/cloudapi/v6` path, so it should _not_ be included in the `IONOS_API_URL` environment variable |
| `IONOS_LOG_LEVEL`    | Specify the Log Level used to log messages. Possible values: Off, Debug, Trace |
| `IONOS_PINNED_CERT`  | Specify the SHA-256 public fingerprint here, enables certificate pinning                                                                                                                                                       |

⚠️ **_Note: To overwrite the api endpoint - `api.ionos.com`, the environment variable `$IONOS_API_URL` can be set, and used with `NewConfigurationFromEnv()` function._**

## Getting started

To start working with the SDK, setup your project for Go modules, and retrieve the SDK dependencies with `go get`. This example shows how to use the SDK
to make an API request to IONOS Cloud.

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

Using your IDE of choice, add the following code:

### Token Authentication
There are 2 ways to generate your token:

### Generate token using [SDK go auth](https://github.com/ionos-cloud/sdk-go-bundle/products/auth):
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
Run commands to log in and generate your token.
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
There are now log levels that you can set: `Off`, `Debug` and `Trace`.

| Log level | Description                                                                                   |
|-----------|-----------------------------------------------------------------------------------------------|
| `Off`     | Does not show any logs                                                                        |
| `Debug`   | Regular logs, no sensitive information                                                        |
| `Trace`   | Logs everything, the full request(along with authentication) and response without encryption. |

⚠️ **_Note: We recommend you only set this `TRACE` for debugging purposes. Disable it in your production environments because it can log sensitive data. <br>
It logs the full request and response without encryption, even for an HTTPS call. <br>
Verbose request and response logging can also significantly impact your application's performance._**

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

All SDKs are found in the folder `https://github.com/ionos-cloud/sdk-go-bundle/tree/master/products`.
- Structs moved to common package: `GenericOpenAPIError`, `Configuration`, `APIResponse`, utility functions, logging interface,  environment variables.
- Import `github.com/ionos-cloud/sdk-go-bundle/common` and replace the structs as required.
- IONOS_DEBUG removed, debugging now set with `IONOS_LOG_LEVEL`, as described [here](https://github.com/ionos-cloud/sdk-go-bundle#debugging)
- Replaced `PtrBool`, `PtrInt`, `PtrInt32`, `PtrInt64`, `PtrFloat`, `PtrFloat32`, `PtrFloat64`, `PtrString`, `PtrTime` with `ToPtr` generic functions

Example migration for compute `github.com/ionos-cloud/sdk-go/v6` to `github.com/ionos-cloud/sdk-go-bundle/products/compute`

1. Replace the import ionoscloud `github.com/ionos-cloud/sdk-go/v6` with `github.com/ionos-cloud/sdk-go-bundle/products/compute` You will get errors, as some structs have been moved the the `common` package
2. Import `github.com/ionos-cloud/sdk-go-bundle/common` and replace the structs as required.
   Example replacements:
      - var apiResponse *ionoscloud.APIResponse with var apiResponse *common.APIResponse
      - ionoscloud.NewConfiguration(username, password, token, url) with common.NewConfiguration(username, password, token, url).
3. Replace the structs that were in the old repo with the ones in the new repo.

All replacements work as-is, no other changes are required.

### List of all SDK products in `sdk-go-bundle`, along with their equivalents in the separate repos

| SDK Name          | Sdk-go-bundle module                             | Full repo equivalent                                               | Description                                                                                                                                                                                                                                                                                                                                                                       |
|-------------------|--------------------------------------------------|--------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Authentication    | [readme](./products/auth/README.md)              | [readme](https://github.com/ionos-cloud/sdk-go-auth)               | Use the Auth API to manage tokens for secure access to IONOS Cloud  APIs (Auth API, Cloud API, Reseller API, Activity Log API, and others).                                                                                                                                                                                                                                       |
| Compute           | [readme](./products/compute/README.md)           | [readme](https://github.com/ionos-cloud/sdk-go)                    | Use the Auth API to manage resources in the compute engine: datacenters, locations, servers, kubernetes resources, images, volumes, lans, ip blocks, network interfaces, private cross connects, firewall rules, flow logs, load balancers, nat gateways, network load balancers, application load balancers, target groups, s3 keys, snapshots, users, groups, labels, templates |
| DBaaS Mongo      | [readme](./products/dbaas/mongo/README.md)       | [readme](https://github.com/ionos-cloud/sdk-go-dbaas-mongo)        | You have the ability to quickly set up and manage a MongoDB database. You can also delete clusters, manage backups and users via the API. The MongoDB API allows you to create additional database clusters or modify existing ones.                                                                                                                                              |
| DBaaS Postgres SQL | [readme](./products/dbaas/psql/README.md)        | [readme](https://github.com/ionos-cloud/sdk-go-dbaas-postgres)     | The API allows you to create additional PostgreSQL database clusters or modify existing ones.                                                                                                                                                                                                                                                                                     |
| Certificate Manager | [readme](./products/cert/README.md)              | [readme](https://github.com/ionos-cloud/sdk-go-cert-manager)       | Using the Certificate Manager product, you can conveniently provision and manage SSL certificates with IONOS services and your internal connected resources.                                                                                                                                                                                                                      |
| Container Registry | [readme](./products/containerregistry/README.md) | [readme](https://github.com/ionos-cloud/sdk-go-container-registry) | Container Registry product enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls.                                                                                                                        |
