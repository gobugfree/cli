# \SkillsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSkills**](SkillsAPI.md#GetSkills) | **Get** /skills | The agent skill as an installable SKILL.md: markdown by default, a JSON wrapper with the version stamp when application/json is accepted; public, no authentication required



## GetSkills

> Skill GetSkills(ctx).Execute()

The agent skill as an installable SKILL.md: markdown by default, a JSON wrapper with the version stamp when application/json is accepted; public, no authentication required

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
	resp, r, err := apiClient.SkillsAPI.GetSkills(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.GetSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSkills`: Skill
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.GetSkills`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSkillsRequest struct via the builder pattern


### Return type

[**Skill**](Skill.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

