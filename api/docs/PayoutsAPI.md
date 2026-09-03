# \PayoutsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPayouts**](PayoutsAPI.md#GetPayouts) | **Get** /payouts | The researcher&#39;s payout method, payout compliance, and payout history
[**PostPayoutsMethod**](PayoutsAPI.md#PostPayoutsMethod) | **Post** /payouts/method | Save the researcher&#39;s default payout method (account details are sealed)
[**PostPayoutsPreferences**](PayoutsAPI.md#PostPayoutsPreferences) | **Post** /payouts/preferences | Save the researcher&#39;s payout hold and minimum-threshold preferences



## GetPayouts

> GetPayouts(ctx).Execute()

The researcher's payout method, payout compliance, and payout history

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

> PostPayoutsMethod(ctx).Iban(iban).Execute()

Save the researcher's default payout method (account details are sealed)

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
	iban := "iban_example" // string | IBAN for the bank transfer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PayoutsAPI.PostPayoutsMethod(context.Background()).Iban(iban).Execute()
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
 **iban** | **string** | IBAN for the bank transfer | 

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

