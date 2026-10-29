# \DocsAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDocs**](DocsAPI.md#GetDocs) | **Get** /docs | Interactive API reference: the caller&#39;s OpenAPI document rendered for people, with try-it requests; JSON callers get the document locations
[**GetDocsAgents**](DocsAPI.md#GetDocsAgents) | **Get** /docs/agents | How to connect an agent: the MCP endpoint and its OAuth discovery document, the client setup, the command line, and the rules agents follow
[**GetHelp**](DocsAPI.md#GetHelp) | **Get** /help | Help for one screen: what the page does and the agent and gbf operations reachable from it, scoped to the caller



## GetDocs

> ApidocReference GetDocs(ctx).Execute()

Interactive API reference: the caller's OpenAPI document rendered for people, with try-it requests; JSON callers get the document locations

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
	resp, r, err := apiClient.DocsAPI.GetDocs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocsAPI.GetDocs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDocs`: ApidocReference
	fmt.Fprintf(os.Stdout, "Response from `DocsAPI.GetDocs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocsRequest struct via the builder pattern


### Return type

[**ApidocReference**](ApidocReference.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocsAgents

> McpAgents GetDocsAgents(ctx).Execute()

How to connect an agent: the MCP endpoint and its OAuth discovery document, the client setup, the command line, and the rules agents follow

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
	resp, r, err := apiClient.DocsAPI.GetDocsAgents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocsAPI.GetDocsAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDocsAgents`: McpAgents
	fmt.Fprintf(os.Stdout, "Response from `DocsAPI.GetDocsAgents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocsAgentsRequest struct via the builder pattern


### Return type

[**McpAgents**](McpAgents.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHelp

> HelpTopic GetHelp(ctx).Page(page).Execute()

Help for one screen: what the page does and the agent and gbf operations reachable from it, scoped to the caller

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
	page := "page_example" // string | The screen's path, for example /organizations/abacus/engagements/internal-api (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocsAPI.GetHelp(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocsAPI.GetHelp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHelp`: HelpTopic
	fmt.Fprintf(os.Stdout, "Response from `DocsAPI.GetHelp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHelpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | The screen&#39;s path, for example /organizations/abacus/engagements/internal-api | 

### Return type

[**HelpTopic**](HelpTopic.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

