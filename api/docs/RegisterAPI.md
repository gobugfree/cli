# \RegisterAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRegisterConfirmByToken**](RegisterAPI.md#GetRegisterConfirmByToken) | **Get** /register/confirm/{token} | Confirm the company-domain email token; the registration enters the operations legitimacy queue
[**PostRegister**](RegisterAPI.md#PostRegister) | **Post** /register | Register an organization: provisions the restricted workspace with its first admin, records the terms acceptance, and mails the company-domain confirmation link



## GetRegisterConfirmByToken

> GetRegisterConfirmByToken(ctx, token).Execute()

Confirm the company-domain email token; the registration enters the operations legitimacy queue

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
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RegisterAPI.GetRegisterConfirmByToken(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegisterAPI.GetRegisterConfirmByToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisterConfirmByTokenRequest struct via the builder pattern


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


## PostRegister

> RegisterRegistrationDetail PostRegister(ctx).AcceptTerms(acceptTerms).AdminEmail(adminEmail).AdminName(adminName).CompanyName(companyName).ContactEmail(contactEmail).Country(country).Domain(domain).LegalName(legalName).RegistryNumber(registryNumber).Website(website).Execute()

Register an organization: provisions the restricted workspace with its first admin, records the terms acceptance, and mails the company-domain confirmation link

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
	acceptTerms := "acceptTerms_example" // string | Platform terms acceptance, checkbox value 'on'
	adminEmail := "adminEmail_example" // string | Primary admin login email
	adminName := "adminName_example" // string | Primary admin full name
	companyName := "companyName_example" // string | Company display name
	contactEmail := "contactEmail_example" // string | Verification contact on the company domain
	country := "country_example" // string | Country of incorporation
	domain := "domain_example" // string | Company email domain used for ownership verification (must match the contact email's domain)
	legalName := "legalName_example" // string | Registered legal entity name
	registryNumber := "registryNumber_example" // string | Commercial-register number (optional)
	website := "website_example" // string | Company website (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegisterAPI.PostRegister(context.Background()).AcceptTerms(acceptTerms).AdminEmail(adminEmail).AdminName(adminName).CompanyName(companyName).ContactEmail(contactEmail).Country(country).Domain(domain).LegalName(legalName).RegistryNumber(registryNumber).Website(website).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegisterAPI.PostRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRegister`: RegisterRegistrationDetail
	fmt.Fprintf(os.Stdout, "Response from `RegisterAPI.PostRegister`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptTerms** | **string** | Platform terms acceptance, checkbox value &#39;on&#39; | 
 **adminEmail** | **string** | Primary admin login email | 
 **adminName** | **string** | Primary admin full name | 
 **companyName** | **string** | Company display name | 
 **contactEmail** | **string** | Verification contact on the company domain | 
 **country** | **string** | Country of incorporation | 
 **domain** | **string** | Company email domain used for ownership verification (must match the contact email&#39;s domain) | 
 **legalName** | **string** | Registered legal entity name | 
 **registryNumber** | **string** | Commercial-register number | 
 **website** | **string** | Company website | 

### Return type

[**RegisterRegistrationDetail**](RegisterRegistrationDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

