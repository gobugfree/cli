# \SettingsAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettings**](SettingsAPI.md#GetSettings) | **Get** /settings | Own researcher profile and visibility settings, with the caller&#39;s numeric userId
[**GetSettingsSecurity**](SettingsAPI.md#GetSettingsSecurity) | **Get** /settings/security | Own sign-in methods as Keycloak holds them (password, passkeys, security keys, authenticator app, recovery codes) and the linked external accounts; never a secret
[**GetSettingsSkills**](SettingsAPI.md#GetSettingsSkills) | **Get** /settings/skills | Own skills with their levels
[**PostSettingsLanguage**](SettingsAPI.md#PostSettingsLanguage) | **Post** /settings/language | Switch the interface language
[**PostSettingsLook**](SettingsAPI.md#PostSettingsLook) | **Post** /settings/look | Switch between the new look and the classic look of the former portal
[**PostSettingsPrivacy**](SettingsAPI.md#PostSettingsPrivacy) | **Post** /settings/privacy | Update the per-section public visibility
[**PostSettingsProfile**](SettingsAPI.md#PostSettingsProfile) | **Post** /settings/profile | Update the researcher profile
[**PostSettingsSecurityActions**](SettingsAPI.md#PostSettingsSecurityActions) | **Post** /settings/security/actions | Start a sign-in method action in Keycloak (add a passkey, set up recovery codes, change the password, link Microsoft, remove a method) and send the browser there; browser sessions only, a JSON client gets 400
[**PostSettingsSkills**](SettingsAPI.md#PostSettingsSkills) | **Post** /settings/skills | Add a skill from the shared taxonomy, or change its level
[**PostSettingsSkillsDelete**](SettingsAPI.md#PostSettingsSkillsDelete) | **Post** /settings/skills/{id}/delete | Remove an own skill



## GetSettings

> UserSettings GetSettings(ctx).Execute()

Own researcher profile and visibility settings, with the caller's numeric userId

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
	resp, r, err := apiClient.SettingsAPI.GetSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettings`: UserSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsRequest struct via the builder pattern


### Return type

[**UserSettings**](UserSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingsSecurity

> UserSecurity GetSettingsSecurity(ctx).Execute()

Own sign-in methods as Keycloak holds them (password, passkeys, security keys, authenticator app, recovery codes) and the linked external accounts; never a secret

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
	resp, r, err := apiClient.SettingsAPI.GetSettingsSecurity(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetSettingsSecurity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingsSecurity`: UserSecurity
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetSettingsSecurity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsSecurityRequest struct via the builder pattern


### Return type

[**UserSecurity**](UserSecurity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingsSkills

> GetSettingsSkills(ctx).Execute()

Own skills with their levels

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
	r, err := apiClient.SettingsAPI.GetSettingsSkills(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetSettingsSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsSkillsRequest struct via the builder pattern


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


## PostSettingsLanguage

> PostSettingsLanguage(ctx).Locale(locale).Redirect(redirect).Execute()

Switch the interface language

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
	locale := "locale_example" // string | 
	redirect := "redirect_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.PostSettingsLanguage(context.Background()).Locale(locale).Redirect(redirect).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PostSettingsLanguage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingsLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locale** | **string** |  | 
 **redirect** | **string** |  | 

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


## PostSettingsLook

> PostSettingsLook(ctx).LegacyLook(legacyLook).Redirect(redirect).Execute()

Switch between the new look and the classic look of the former portal

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
	legacyLook := "legacyLook_example" // string | Send \\\"on\\\" for the classic look, omit for the new look (optional)
	redirect := "redirect_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.PostSettingsLook(context.Background()).LegacyLook(legacyLook).Redirect(redirect).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PostSettingsLook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingsLookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **legacyLook** | **string** | Send \\\&quot;on\\\&quot; for the classic look, omit for the new look | 
 **redirect** | **string** |  | 

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


## PostSettingsPrivacy

> PostSettingsPrivacy(ctx).IsPublic(isPublic).ShowEarnings(showEarnings).ShowOnLeaderboard(showOnLeaderboard).ShowVulnTypes(showVulnTypes).Execute()

Update the per-section public visibility

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
	isPublic := "isPublic_example" // string | Send \\\"on\\\" to show the public profile, omit to hide it (optional)
	showEarnings := "showEarnings_example" // string | Send \\\"on\\\" to enable, omit to disable (optional)
	showOnLeaderboard := "showOnLeaderboard_example" // string | Send \\\"on\\\" to enable, omit to disable (optional)
	showVulnTypes := "showVulnTypes_example" // string | Send \\\"on\\\" to enable, omit to disable (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.PostSettingsPrivacy(context.Background()).IsPublic(isPublic).ShowEarnings(showEarnings).ShowOnLeaderboard(showOnLeaderboard).ShowVulnTypes(showVulnTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PostSettingsPrivacy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingsPrivacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isPublic** | **string** | Send \\\&quot;on\\\&quot; to show the public profile, omit to hide it | 
 **showEarnings** | **string** | Send \\\&quot;on\\\&quot; to enable, omit to disable | 
 **showOnLeaderboard** | **string** | Send \\\&quot;on\\\&quot; to enable, omit to disable | 
 **showVulnTypes** | **string** | Send \\\&quot;on\\\&quot; to enable, omit to disable | 

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


## PostSettingsProfile

> UserSettings PostSettingsProfile(ctx).DisplayName(displayName).Handle(handle).Bio(bio).BountyExperience(bountyExperience).Country(country).ExperienceLevel(experienceLevel).Website(website).Execute()

Update the researcher profile

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
	displayName := "displayName_example" // string | Public display name
	handle := "handle_example" // string | Public researcher handle: lowercase letters, numbers, and dashes
	bio := "bio_example" // string | Profile bio (optional)
	bountyExperience := "bountyExperience_example" // string | Empty means not specified (optional)
	country := "country_example" // string | Country (optional)
	experienceLevel := "experienceLevel_example" // string | Empty means not specified (optional)
	website := "website_example" // string | Website URL (http or https) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.PostSettingsProfile(context.Background()).DisplayName(displayName).Handle(handle).Bio(bio).BountyExperience(bountyExperience).Country(country).ExperienceLevel(experienceLevel).Website(website).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PostSettingsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSettingsProfile`: UserSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.PostSettingsProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **displayName** | **string** | Public display name | 
 **handle** | **string** | Public researcher handle: lowercase letters, numbers, and dashes | 
 **bio** | **string** | Profile bio | 
 **bountyExperience** | **string** | Empty means not specified | 
 **country** | **string** | Country | 
 **experienceLevel** | **string** | Empty means not specified | 
 **website** | **string** | Website URL (http or https) | 

### Return type

[**UserSettings**](UserSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSettingsSecurityActions

> PostSettingsSecurityActions(ctx).Action(action).CredentialId(credentialId).Execute()

Start a sign-in method action in Keycloak (add a passkey, set up recovery codes, change the password, link Microsoft, remove a method) and send the browser there; browser sessions only, a JSON client gets 400

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
	action := "action_example" // string | The sign-in method action to start in Keycloak
	credentialId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The sign-in method to remove, required for remove-credential (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.PostSettingsSecurityActions(context.Background()).Action(action).CredentialId(credentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PostSettingsSecurityActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingsSecurityActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | The sign-in method action to start in Keycloak | 
 **credentialId** | **string** | The sign-in method to remove, required for remove-credential | 

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


## PostSettingsSkills

> PostSettingsSkills(ctx).Level(level).Name(name).Execute()

Add a skill from the shared taxonomy, or change its level

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
	level := "level_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.PostSettingsSkills(context.Background()).Level(level).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PostSettingsSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingsSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **level** | **string** |  | 
 **name** | **string** |  | 

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


## PostSettingsSkillsDelete

> PostSettingsSkillsDelete(ctx, id).Execute()

Remove an own skill

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
	r, err := apiClient.SettingsAPI.PostSettingsSkillsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PostSettingsSkillsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostSettingsSkillsDeleteRequest struct via the builder pattern


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

