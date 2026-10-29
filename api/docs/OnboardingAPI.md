# \OnboardingAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOnboarding**](OnboardingAPI.md#GetOnboarding) | **Get** /onboarding | Researcher onboarding state: the current wizard step, conduct acceptance, and completion
[**PostOnboardingConduct**](OnboardingAPI.md#PostOnboardingConduct) | **Post** /onboarding/conduct | Onboarding step 4: accept the community guidelines and complete registration
[**PostOnboardingExperience**](OnboardingAPI.md#PostOnboardingExperience) | **Post** /onboarding/experience | Onboarding step 2: experience
[**PostOnboardingInterests**](OnboardingAPI.md#PostOnboardingInterests) | **Post** /onboarding/interests | Onboarding step 3: skills from the shared taxonomy



## GetOnboarding

> UserOnboarding GetOnboarding(ctx).Execute()

Researcher onboarding state: the current wizard step, conduct acceptance, and completion

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
	resp, r, err := apiClient.OnboardingAPI.GetOnboarding(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnboardingAPI.GetOnboarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnboarding`: UserOnboarding
	fmt.Fprintf(os.Stdout, "Response from `OnboardingAPI.GetOnboarding`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnboardingRequest struct via the builder pattern


### Return type

[**UserOnboarding**](UserOnboarding.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOnboardingConduct

> PostOnboardingConduct(ctx).Accept(accept).Execute()

Onboarding step 4: accept the community guidelines and complete registration

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
	accept := "accept_example" // string | Must be \\\"on\\\"; registration cannot complete without acceptance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OnboardingAPI.PostOnboardingConduct(context.Background()).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnboardingAPI.PostOnboardingConduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOnboardingConductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Must be \\\&quot;on\\\&quot;; registration cannot complete without acceptance | 

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


## PostOnboardingExperience

> PostOnboardingExperience(ctx).ExperienceLevel(experienceLevel).BountyExperience(bountyExperience).Execute()

Onboarding step 2: experience

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
	experienceLevel := "experienceLevel_example" // string | 
	bountyExperience := "bountyExperience_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OnboardingAPI.PostOnboardingExperience(context.Background()).ExperienceLevel(experienceLevel).BountyExperience(bountyExperience).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnboardingAPI.PostOnboardingExperience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOnboardingExperienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **experienceLevel** | **string** |  | 
 **bountyExperience** | **string** |  | 

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


## PostOnboardingInterests

> PostOnboardingInterests(ctx).Skills(skills).Execute()

Onboarding step 3: skills from the shared taxonomy

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
	skills := []string{"Inner_example"} // []string | Repeat the parameter per selected skill (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OnboardingAPI.PostOnboardingInterests(context.Background()).Skills(skills).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnboardingAPI.PostOnboardingInterests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOnboardingInterestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skills** | **[]string** | Repeat the parameter per selected skill | 

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

