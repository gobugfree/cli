# \PayoutsAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPayouts**](PayoutsAPI.md#GetPayouts) | **Get** /payouts | The researcher&#39;s payout method, payout preferences, and payout history
[**PostPayoutsMethod**](PayoutsAPI.md#PostPayoutsMethod) | **Post** /payouts/method | Save the researcher&#39;s payout details: bank account, postal address, company fields, and account currency (sealed; an empty IBAN keeps the stored one)
[**PostPayoutsPreferences**](PayoutsAPI.md#PostPayoutsPreferences) | **Post** /payouts/preferences | Save the researcher&#39;s payout hold and minimum-threshold preferences



## GetPayouts

> GetPayouts(ctx).Execute()

The researcher's payout method, payout preferences, and payout history

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
	r, err := apiClient.PayoutsAPI.GetPayouts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PayoutsAPI.GetPayouts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutsRequest struct via the builder pattern


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


## PostPayoutsMethod

> PostPayoutsMethod(ctx).AccountHolder(accountHolder).Bic(bic).City(city).Country(country).Currency(currency).PostalCode(postalCode).Street(street).BankAddress(bankAddress).CompanyName(companyName).CompanyVatOrId(companyVatOrId).Iban(iban).UseCompanyAddress(useCompanyAddress).Execute()

Save the researcher's payout details: bank account, postal address, company fields, and account currency (sealed; an empty IBAN keeps the stored one)

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
	accountHolder := "accountHolder_example" // string | Name of the account holder
	bic := "bic_example" // string | BIC of the bank
	city := "city_example" // string | City
	country := "country_example" // string | Country
	currency := "currency_example" // string | Currency the account receives
	postalCode := "postalCode_example" // string | Postal code
	street := "street_example" // string | Street and number
	bankAddress := "bankAddress_example" // string | Address of the bank (optional)
	companyName := "companyName_example" // string | Company name; required with the company address (optional)
	companyVatOrId := "companyVatOrId_example" // string | VAT or company id; required with the company address (optional)
	iban := "iban_example" // string | IBAN; empty keeps the stored one (optional)
	useCompanyAddress := "useCompanyAddress_example" // string | Use the address as a company address, checkbox value 'on' (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PayoutsAPI.PostPayoutsMethod(context.Background()).AccountHolder(accountHolder).Bic(bic).City(city).Country(country).Currency(currency).PostalCode(postalCode).Street(street).BankAddress(bankAddress).CompanyName(companyName).CompanyVatOrId(companyVatOrId).Iban(iban).UseCompanyAddress(useCompanyAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PayoutsAPI.PostPayoutsMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPayoutsMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountHolder** | **string** | Name of the account holder | 
 **bic** | **string** | BIC of the bank | 
 **city** | **string** | City | 
 **country** | **string** | Country | 
 **currency** | **string** | Currency the account receives | 
 **postalCode** | **string** | Postal code | 
 **street** | **string** | Street and number | 
 **bankAddress** | **string** | Address of the bank | 
 **companyName** | **string** | Company name; required with the company address | 
 **companyVatOrId** | **string** | VAT or company id; required with the company address | 
 **iban** | **string** | IBAN; empty keeps the stored one | 
 **useCompanyAddress** | **string** | Use the address as a company address, checkbox value &#39;on&#39; | 

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


## PostPayoutsPreferences

> PostPayoutsPreferences(ctx).Hold(hold).MinAmount(minAmount).Execute()

Save the researcher's payout hold and minimum-threshold preferences

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
	hold := "hold_example" // string | Hold payouts, checkbox value 'on' (optional)
	minAmount := "minAmount_example" // string | Minimum payout amount in whole CHF (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PayoutsAPI.PostPayoutsPreferences(context.Background()).Hold(hold).MinAmount(minAmount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PayoutsAPI.PostPayoutsPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPayoutsPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hold** | **string** | Hold payouts, checkbox value &#39;on&#39; | 
 **minAmount** | **string** | Minimum payout amount in whole CHF | 

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

