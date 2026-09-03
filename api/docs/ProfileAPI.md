# \ProfileAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProfileCertifications**](ProfileAPI.md#GetProfileCertifications) | **Get** /profile/certifications | Own certifications with their verification status
[**GetProfileExternal**](ProfileAPI.md#GetProfileExternal) | **Get** /profile/external | Own external profiles with proof-of-control status; pending challenges include the proof token and placement instructions
[**PostProfileCertifications**](ProfileAPI.md#PostProfileCertifications) | **Post** /profile/certifications | Submit a certification; issuers with a verification URL are checked automatically, others go to staff review
[**PostProfileCertificationsDelete**](ProfileAPI.md#PostProfileCertificationsDelete) | **Post** /profile/certifications/{id}/delete | Remove an own certification
[**PostProfileExternal**](ProfileAPI.md#PostProfileExternal) | **Post** /profile/external | Start (or restart) a proof-of-control challenge: place the returned proof token on the external profile, then check
[**PostProfileExternalCheck**](ProfileAPI.md#PostProfileExternalCheck) | **Post** /profile/external/{id}/check | Enqueue an asynchronous proof check for an own open challenge; a verified profile is returned unchanged
[**PostProfileExternalDelete**](ProfileAPI.md#PostProfileExternalDelete) | **Post** /profile/external/{id}/delete | Remove an own external profile
[**PostProfileExternalRevoke**](ProfileAPI.md#PostProfileExternalRevoke) | **Post** /profile/external/{id}/revoke | Revoke an own verified external profile



## GetProfileCertifications

> UserCertificationList GetProfileCertifications(ctx).Execute()

Own certifications with their verification status

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gobugfree/cli/v3/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfileAPI.GetProfileCertifications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.GetProfileCertifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfileCertifications`: UserCertificationList
	fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.GetProfileCertifications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileCertificationsRequest struct via the builder pattern


### Return type

[**UserCertificationList**](UserCertificationList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileExternal

> ProofExternalProfileList GetProfileExternal(ctx).Execute()

Own external profiles with proof-of-control status; pending challenges include the proof token and placement instructions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gobugfree/cli/v3/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfileAPI.GetProfileExternal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.GetProfileExternal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfileExternal`: ProofExternalProfileList
	fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.GetProfileExternal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileExternalRequest struct via the builder pattern


### Return type

[**ProofExternalProfileList**](ProofExternalProfileList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProfileCertifications

> PostProfileCertifications(ctx).CredentialId(credentialId).Issuer(issuer).Name(name).ExpiresOn(expiresOn).IssuedOn(issuedOn).Url(url).Execute()

Submit a certification; issuers with a verification URL are checked automatically, others go to staff review

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/gobugfree/cli/v3/api"
)

func main() {
	credentialId := "credentialId_example" // string | Issuer-side credential ID
	issuer := "issuer_example" // string | Certification issuer
	name := "name_example" // string | Certification name
	expiresOn := time.Now() // string | Expiry date (YYYY-MM-DD); an expired certification loses its verified mark (optional)
	issuedOn := time.Now() // string | Issue date (YYYY-MM-DD) (optional)
	url := "url_example" // string | Issuer verification URL on the issuer's domain; required for issuers with automatic checks (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProfileAPI.PostProfileCertifications(context.Background()).CredentialId(credentialId).Issuer(issuer).Name(name).ExpiresOn(expiresOn).IssuedOn(issuedOn).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.PostProfileCertifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProfileCertificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialId** | **string** | Issuer-side credential ID | 
 **issuer** | **string** | Certification issuer | 
 **name** | **string** | Certification name | 
 **expiresOn** | **string** | Expiry date (YYYY-MM-DD); an expired certification loses its verified mark | 
 **issuedOn** | **string** | Issue date (YYYY-MM-DD) | 
 **url** | **string** | Issuer verification URL on the issuer&#39;s domain; required for issuers with automatic checks | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProfileCertificationsDelete

> PostProfileCertificationsDelete(ctx, id).Execute()

Remove an own certification

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gobugfree/cli/v3/api"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProfileAPI.PostProfileCertificationsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.PostProfileCertificationsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProfileCertificationsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProfileExternal

> ProofExternalProfile PostProfileExternal(ctx).Handle(handle).Provider(provider).Execute()

Start (or restart) a proof-of-control challenge: place the returned proof token on the external profile, then check

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gobugfree/cli/v3/api"
)

func main() {
	handle := "handle_example" // string | Handle or profile URL on the provider; normalized server-side
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfileAPI.PostProfileExternal(context.Background()).Handle(handle).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.PostProfileExternal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProfileExternal`: ProofExternalProfile
	fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.PostProfileExternal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProfileExternalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **handle** | **string** | Handle or profile URL on the provider; normalized server-side | 
 **provider** | **string** |  | 

### Return type

[**ProofExternalProfile**](ProofExternalProfile.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProfileExternalCheck

> ProofExternalProfile PostProfileExternalCheck(ctx, id).Execute()

Enqueue an asynchronous proof check for an own open challenge; a verified profile is returned unchanged

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gobugfree/cli/v3/api"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfileAPI.PostProfileExternalCheck(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.PostProfileExternalCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProfileExternalCheck`: ProofExternalProfile
	fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.PostProfileExternalCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProfileExternalCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProofExternalProfile**](ProofExternalProfile.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProfileExternalDelete

> PostProfileExternalDelete(ctx, id).Execute()

Remove an own external profile

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gobugfree/cli/v3/api"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProfileAPI.PostProfileExternalDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.PostProfileExternalDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProfileExternalDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProfileExternalRevoke

> ProofExternalProfile PostProfileExternalRevoke(ctx, id).Execute()

Revoke an own verified external profile

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gobugfree/cli/v3/api"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfileAPI.PostProfileExternalRevoke(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.PostProfileExternalRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProfileExternalRevoke`: ProofExternalProfile
	fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.PostProfileExternalRevoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProfileExternalRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProofExternalProfile**](ProofExternalProfile.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

