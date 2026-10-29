# \DevAPI

All URIs are relative to *https://app.gobugfree.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDevJiraRestApi3IssueLinkById**](DevAPI.md#DeleteDevJiraRestApi3IssueLinkById) | **Delete** /dev/jira/rest/api/3/issueLink/{id} | Development-only stub of GObugfree&#39;s Jira site: remove a link
[**GetDevAthenaBlobByBucketAndKey**](DevAPI.md#GetDevAthenaBlobByBucketAndKey) | **Get** /dev/athena/blob/{bucket}/{key} | Development-only stub of athena&#39;s private file storage: one blob, read with the stub&#39;s signature
[**GetDevAthenaJiraRestApi3AttachmentContentById**](DevAPI.md#GetDevAthenaJiraRestApi3AttachmentContentById) | **Get** /dev/athena/jira/rest/api/3/attachment/content/{id} | Development-only stub of athena&#39;s Jira: the content of one attachment
[**GetDevAthenaJiraRestApi3Field**](DevAPI.md#GetDevAthenaJiraRestApi3Field) | **Get** /dev/athena/jira/rest/api/3/field | Development-only stub of athena&#39;s Jira: the field catalog
[**GetDevAthenaJiraRestApi3IssueByKey**](DevAPI.md#GetDevAthenaJiraRestApi3IssueByKey) | **Get** /dev/athena/jira/rest/api/3/issue/{key} | Development-only stub of athena&#39;s Jira: one issue with every field
[**GetDevAthenaJiraRestApi3IssueChangelog**](DevAPI.md#GetDevAthenaJiraRestApi3IssueChangelog) | **Get** /dev/athena/jira/rest/api/3/issue/{key}/changelog | Development-only stub of athena&#39;s Jira: the changelog of an issue
[**GetDevAthenaJiraRestApi3IssueComment**](DevAPI.md#GetDevAthenaJiraRestApi3IssueComment) | **Get** /dev/athena/jira/rest/api/3/issue/{key}/comment | Development-only stub of athena&#39;s Jira: the comments of an issue
[**GetDevAthenaStrapiApiOrganisations**](DevAPI.md#GetDevAthenaStrapiApiOrganisations) | **Get** /dev/athena/strapi/api/organisations | Development-only stub of athena&#39;s Strapi: the organisations with their programs
[**GetDevAthenaStrapiApiPrograms**](DevAPI.md#GetDevAthenaStrapiApiPrograms) | **Get** /dev/athena/strapi/api/programs | Development-only stub of athena&#39;s Strapi: one program by slug
[**GetDevAthenaStrapiApiScanningReports**](DevAPI.md#GetDevAthenaStrapiApiScanningReports) | **Get** /dev/athena/strapi/api/scanning-reports | Development-only stub of athena&#39;s Strapi: the scanning reports of an organisation
[**GetDevAthenaStrapiUploadsByName**](DevAPI.md#GetDevAthenaStrapiUploadsByName) | **Get** /dev/athena/strapi/uploads/{name} | Development-only stub of athena&#39;s Strapi: one uploaded file
[**GetDevBrandSiteBySite**](DevAPI.md#GetDevBrandSiteBySite) | **Get** /dev/brand-site/{site} | Development-only stub brand site homepage (acme, plain, ico-only) the website brand extraction can fetch locally
[**GetDevBrandSiteBySiteAndFile**](DevAPI.md#GetDevBrandSiteBySiteAndFile) | **Get** /dev/brand-site/{site}/{file} | Development-only stub brand site asset, addressed by its file stem
[**GetDevJiraIssuesByKey**](DevAPI.md#GetDevJiraIssuesByKey) | **Get** /dev/jira/_issues/{key} | Development-only stub of GObugfree&#39;s Jira site: the stub&#39;s view of an issue, for tests
[**GetDevJiraRestApi3IssueByKey**](DevAPI.md#GetDevJiraRestApi3IssueByKey) | **Get** /dev/jira/rest/api/3/issue/{key} | Development-only stub of GObugfree&#39;s Jira site: one issue
[**GetDevJiraRestApi3IssueComment**](DevAPI.md#GetDevJiraRestApi3IssueComment) | **Get** /dev/jira/rest/api/3/issue/{key}/comment | Development-only stub of GObugfree&#39;s Jira site: the comments of an issue
[**GetDevJiraRestApi3IssueCommentById**](DevAPI.md#GetDevJiraRestApi3IssueCommentById) | **Get** /dev/jira/rest/api/3/issue/{key}/comment/{id} | Development-only stub of GObugfree&#39;s Jira site: one comment
[**GetDevJiraRestApi3IssueTransitions**](DevAPI.md#GetDevJiraRestApi3IssueTransitions) | **Get** /dev/jira/rest/api/3/issue/{key}/transitions | Development-only stub of GObugfree&#39;s Jira site: the moves of an issue
[**GetDevJiraRestApi3Myself**](DevAPI.md#GetDevJiraRestApi3Myself) | **Get** /dev/jira/rest/api/3/myself | Development-only stub of GObugfree&#39;s Jira site: the app&#39;s own account
[**GetDevJiraRestApi3ProjectByKey**](DevAPI.md#GetDevJiraRestApi3ProjectByKey) | **Get** /dev/jira/rest/api/3/project/{key} | Development-only stub of GObugfree&#39;s Jira site: one project
[**GetDevJiraRestApi3ProjectRole**](DevAPI.md#GetDevJiraRestApi3ProjectRole) | **Get** /dev/jira/rest/api/3/project/{key}/role | Development-only stub of GObugfree&#39;s Jira site: the roles of a project
[**GetDevJiraRestApi3SearchJql**](DevAPI.md#GetDevJiraRestApi3SearchJql) | **Get** /dev/jira/rest/api/3/search/jql | Development-only stub of GObugfree&#39;s Jira site: the issues of a project
[**GetDevLlmV1Models**](DevAPI.md#GetDevLlmV1Models) | **Get** /dev/llm/v1/models | Development-only stub worker: the models it answers to, in the OpenAI format
[**PostDevJiraAct**](DevAPI.md#PostDevJiraAct) | **Post** /dev/jira/_act | Development-only stub of GObugfree&#39;s Jira site: act as a Jira user and deliver the webhook, for tests
[**PostDevJiraRestApi3Issue**](DevAPI.md#PostDevJiraRestApi3Issue) | **Post** /dev/jira/rest/api/3/issue | Development-only stub of GObugfree&#39;s Jira site: create an issue
[**PostDevJiraRestApi3IssueAttachments**](DevAPI.md#PostDevJiraRestApi3IssueAttachments) | **Post** /dev/jira/rest/api/3/issue/{key}/attachments | Development-only stub of GObugfree&#39;s Jira site: attach a file
[**PostDevJiraRestApi3IssueComment**](DevAPI.md#PostDevJiraRestApi3IssueComment) | **Post** /dev/jira/rest/api/3/issue/{key}/comment | Development-only stub of GObugfree&#39;s Jira site: comment on an issue
[**PostDevJiraRestApi3IssueLink**](DevAPI.md#PostDevJiraRestApi3IssueLink) | **Post** /dev/jira/rest/api/3/issueLink | Development-only stub of GObugfree&#39;s Jira site: link two issues
[**PostDevJiraRestApi3IssueTransitions**](DevAPI.md#PostDevJiraRestApi3IssueTransitions) | **Post** /dev/jira/rest/api/3/issue/{key}/transitions | Development-only stub of GObugfree&#39;s Jira site: move an issue
[**PostDevLlmV1ChatCompletions**](DevAPI.md#PostDevLlmV1ChatCompletions) | **Post** /dev/llm/v1/chat/completions | Development-only stub worker: streams the start of the last user message back in the OpenAI format
[**PostDevLlmV1Embeddings**](DevAPI.md#PostDevLlmV1Embeddings) | **Post** /dev/llm/v1/embeddings | Development-only stub worker: hashed bag-of-words embeddings in the OpenAI format
[**PutDevJiraRestApi3IssueByKey**](DevAPI.md#PutDevJiraRestApi3IssueByKey) | **Put** /dev/jira/rest/api/3/issue/{key} | Development-only stub of GObugfree&#39;s Jira site: set fields of an issue



## DeleteDevJiraRestApi3IssueLinkById

> DeleteDevJiraRestApi3IssueLinkById(ctx, id).Execute()

Development-only stub of GObugfree's Jira site: remove a link

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
	r, err := apiClient.DevAPI.DeleteDevJiraRestApi3IssueLinkById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.DeleteDevJiraRestApi3IssueLinkById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDevJiraRestApi3IssueLinkByIdRequest struct via the builder pattern


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


## GetDevAthenaBlobByBucketAndKey

> GetDevAthenaBlobByBucketAndKey(ctx, bucket, key).Execute()

Development-only stub of athena's private file storage: one blob, read with the stub's signature

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
	bucket := "bucket_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevAthenaBlobByBucketAndKey(context.Background(), bucket, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevAthenaBlobByBucketAndKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevAthenaBlobByBucketAndKeyRequest struct via the builder pattern


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


## GetDevAthenaJiraRestApi3AttachmentContentById

> GetDevAthenaJiraRestApi3AttachmentContentById(ctx, id).Execute()

Development-only stub of athena's Jira: the content of one attachment

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
	r, err := apiClient.DevAPI.GetDevAthenaJiraRestApi3AttachmentContentById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevAthenaJiraRestApi3AttachmentContentById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetDevAthenaJiraRestApi3AttachmentContentByIdRequest struct via the builder pattern


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


## GetDevAthenaJiraRestApi3Field

> GetDevAthenaJiraRestApi3Field(ctx).Execute()

Development-only stub of athena's Jira: the field catalog

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
	r, err := apiClient.DevAPI.GetDevAthenaJiraRestApi3Field(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevAthenaJiraRestApi3Field``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevAthenaJiraRestApi3FieldRequest struct via the builder pattern


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


## GetDevAthenaJiraRestApi3IssueByKey

> GetDevAthenaJiraRestApi3IssueByKey(ctx, key).Execute()

Development-only stub of athena's Jira: one issue with every field

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevAthenaJiraRestApi3IssueByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevAthenaJiraRestApi3IssueByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevAthenaJiraRestApi3IssueByKeyRequest struct via the builder pattern


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


## GetDevAthenaJiraRestApi3IssueChangelog

> GetDevAthenaJiraRestApi3IssueChangelog(ctx, key).Execute()

Development-only stub of athena's Jira: the changelog of an issue

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevAthenaJiraRestApi3IssueChangelog(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevAthenaJiraRestApi3IssueChangelog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevAthenaJiraRestApi3IssueChangelogRequest struct via the builder pattern


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


## GetDevAthenaJiraRestApi3IssueComment

> GetDevAthenaJiraRestApi3IssueComment(ctx, key).Execute()

Development-only stub of athena's Jira: the comments of an issue

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevAthenaJiraRestApi3IssueComment(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevAthenaJiraRestApi3IssueComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevAthenaJiraRestApi3IssueCommentRequest struct via the builder pattern


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


## GetDevAthenaStrapiApiOrganisations

> GetDevAthenaStrapiApiOrganisations(ctx).Execute()

Development-only stub of athena's Strapi: the organisations with their programs

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
	r, err := apiClient.DevAPI.GetDevAthenaStrapiApiOrganisations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevAthenaStrapiApiOrganisations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevAthenaStrapiApiOrganisationsRequest struct via the builder pattern


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


## GetDevAthenaStrapiApiPrograms

> GetDevAthenaStrapiApiPrograms(ctx).Execute()

Development-only stub of athena's Strapi: one program by slug

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
	r, err := apiClient.DevAPI.GetDevAthenaStrapiApiPrograms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevAthenaStrapiApiPrograms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevAthenaStrapiApiProgramsRequest struct via the builder pattern


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


## GetDevAthenaStrapiApiScanningReports

> GetDevAthenaStrapiApiScanningReports(ctx).Execute()

Development-only stub of athena's Strapi: the scanning reports of an organisation

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
	r, err := apiClient.DevAPI.GetDevAthenaStrapiApiScanningReports(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevAthenaStrapiApiScanningReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevAthenaStrapiApiScanningReportsRequest struct via the builder pattern


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


## GetDevAthenaStrapiUploadsByName

> GetDevAthenaStrapiUploadsByName(ctx, name).Execute()

Development-only stub of athena's Strapi: one uploaded file

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevAthenaStrapiUploadsByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevAthenaStrapiUploadsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevAthenaStrapiUploadsByNameRequest struct via the builder pattern


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


## GetDevBrandSiteBySite

> GetDevBrandSiteBySite(ctx, site).Execute()

Development-only stub brand site homepage (acme, plain, ico-only) the website brand extraction can fetch locally

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
	site := "site_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevBrandSiteBySite(context.Background(), site).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevBrandSiteBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**site** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevBrandSiteBySiteRequest struct via the builder pattern


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


## GetDevBrandSiteBySiteAndFile

> GetDevBrandSiteBySiteAndFile(ctx, site, file).Execute()

Development-only stub brand site asset, addressed by its file stem

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
	site := "site_example" // string | 
	file := "file_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevBrandSiteBySiteAndFile(context.Background(), site, file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevBrandSiteBySiteAndFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**site** | **string** |  | 
**file** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevBrandSiteBySiteAndFileRequest struct via the builder pattern


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


## GetDevJiraIssuesByKey

> GetDevJiraIssuesByKey(ctx, key).Execute()

Development-only stub of GObugfree's Jira site: the stub's view of an issue, for tests

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevJiraIssuesByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevJiraIssuesByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevJiraIssuesByKeyRequest struct via the builder pattern


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


## GetDevJiraRestApi3IssueByKey

> GetDevJiraRestApi3IssueByKey(ctx, key).Execute()

Development-only stub of GObugfree's Jira site: one issue

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevJiraRestApi3IssueByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevJiraRestApi3IssueByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevJiraRestApi3IssueByKeyRequest struct via the builder pattern


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


## GetDevJiraRestApi3IssueComment

> GetDevJiraRestApi3IssueComment(ctx, key).Execute()

Development-only stub of GObugfree's Jira site: the comments of an issue

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevJiraRestApi3IssueComment(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevJiraRestApi3IssueComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevJiraRestApi3IssueCommentRequest struct via the builder pattern


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


## GetDevJiraRestApi3IssueCommentById

> GetDevJiraRestApi3IssueCommentById(ctx, key, id).Execute()

Development-only stub of GObugfree's Jira site: one comment

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
	key := "key_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevJiraRestApi3IssueCommentById(context.Background(), key, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevJiraRestApi3IssueCommentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevJiraRestApi3IssueCommentByIdRequest struct via the builder pattern


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


## GetDevJiraRestApi3IssueTransitions

> GetDevJiraRestApi3IssueTransitions(ctx, key).Execute()

Development-only stub of GObugfree's Jira site: the moves of an issue

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevJiraRestApi3IssueTransitions(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevJiraRestApi3IssueTransitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevJiraRestApi3IssueTransitionsRequest struct via the builder pattern


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


## GetDevJiraRestApi3Myself

> GetDevJiraRestApi3Myself(ctx).Execute()

Development-only stub of GObugfree's Jira site: the app's own account

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
	r, err := apiClient.DevAPI.GetDevJiraRestApi3Myself(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevJiraRestApi3Myself``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevJiraRestApi3MyselfRequest struct via the builder pattern


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


## GetDevJiraRestApi3ProjectByKey

> GetDevJiraRestApi3ProjectByKey(ctx, key).Execute()

Development-only stub of GObugfree's Jira site: one project

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevJiraRestApi3ProjectByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevJiraRestApi3ProjectByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevJiraRestApi3ProjectByKeyRequest struct via the builder pattern


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


## GetDevJiraRestApi3ProjectRole

> GetDevJiraRestApi3ProjectRole(ctx, key).Execute()

Development-only stub of GObugfree's Jira site: the roles of a project

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.GetDevJiraRestApi3ProjectRole(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevJiraRestApi3ProjectRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevJiraRestApi3ProjectRoleRequest struct via the builder pattern


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


## GetDevJiraRestApi3SearchJql

> GetDevJiraRestApi3SearchJql(ctx).Execute()

Development-only stub of GObugfree's Jira site: the issues of a project

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
	r, err := apiClient.DevAPI.GetDevJiraRestApi3SearchJql(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevJiraRestApi3SearchJql``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevJiraRestApi3SearchJqlRequest struct via the builder pattern


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


## GetDevLlmV1Models

> GetDevLlmV1Models(ctx).Execute()

Development-only stub worker: the models it answers to, in the OpenAI format

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
	r, err := apiClient.DevAPI.GetDevLlmV1Models(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.GetDevLlmV1Models``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevLlmV1ModelsRequest struct via the builder pattern


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


## PostDevJiraAct

> PostDevJiraAct(ctx).Execute()

Development-only stub of GObugfree's Jira site: act as a Jira user and deliver the webhook, for tests

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
	r, err := apiClient.DevAPI.PostDevJiraAct(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.PostDevJiraAct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostDevJiraActRequest struct via the builder pattern


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


## PostDevJiraRestApi3Issue

> PostDevJiraRestApi3Issue(ctx).Execute()

Development-only stub of GObugfree's Jira site: create an issue

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
	r, err := apiClient.DevAPI.PostDevJiraRestApi3Issue(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.PostDevJiraRestApi3Issue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostDevJiraRestApi3IssueRequest struct via the builder pattern


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


## PostDevJiraRestApi3IssueAttachments

> PostDevJiraRestApi3IssueAttachments(ctx, key).Execute()

Development-only stub of GObugfree's Jira site: attach a file

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.PostDevJiraRestApi3IssueAttachments(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.PostDevJiraRestApi3IssueAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDevJiraRestApi3IssueAttachmentsRequest struct via the builder pattern


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


## PostDevJiraRestApi3IssueComment

> PostDevJiraRestApi3IssueComment(ctx, key).Execute()

Development-only stub of GObugfree's Jira site: comment on an issue

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.PostDevJiraRestApi3IssueComment(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.PostDevJiraRestApi3IssueComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDevJiraRestApi3IssueCommentRequest struct via the builder pattern


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


## PostDevJiraRestApi3IssueLink

> PostDevJiraRestApi3IssueLink(ctx).Execute()

Development-only stub of GObugfree's Jira site: link two issues

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
	r, err := apiClient.DevAPI.PostDevJiraRestApi3IssueLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.PostDevJiraRestApi3IssueLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostDevJiraRestApi3IssueLinkRequest struct via the builder pattern


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


## PostDevJiraRestApi3IssueTransitions

> PostDevJiraRestApi3IssueTransitions(ctx, key).Execute()

Development-only stub of GObugfree's Jira site: move an issue

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.PostDevJiraRestApi3IssueTransitions(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.PostDevJiraRestApi3IssueTransitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDevJiraRestApi3IssueTransitionsRequest struct via the builder pattern


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


## PostDevLlmV1ChatCompletions

> PostDevLlmV1ChatCompletions(ctx).Execute()

Development-only stub worker: streams the start of the last user message back in the OpenAI format

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
	r, err := apiClient.DevAPI.PostDevLlmV1ChatCompletions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.PostDevLlmV1ChatCompletions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostDevLlmV1ChatCompletionsRequest struct via the builder pattern


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


## PostDevLlmV1Embeddings

> PostDevLlmV1Embeddings(ctx).Execute()

Development-only stub worker: hashed bag-of-words embeddings in the OpenAI format

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
	r, err := apiClient.DevAPI.PostDevLlmV1Embeddings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.PostDevLlmV1Embeddings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostDevLlmV1EmbeddingsRequest struct via the builder pattern


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


## PutDevJiraRestApi3IssueByKey

> PutDevJiraRestApi3IssueByKey(ctx, key).Execute()

Development-only stub of GObugfree's Jira site: set fields of an issue

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevAPI.PutDevJiraRestApi3IssueByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevAPI.PutDevJiraRestApi3IssueByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDevJiraRestApi3IssueByKeyRequest struct via the builder pattern


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

