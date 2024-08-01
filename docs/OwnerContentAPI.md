# \OwnerContentAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContent**](OwnerContentAPI.md#CreateContent) | **Post** /owners/{owner_key}/content | 
[**CreateContentBatch**](OwnerContentAPI.md#CreateContentBatch) | **Post** /owners/{owner_key}/content/batch | 
[**GetContentById**](OwnerContentAPI.md#GetContentById) | **Get** /owners/{owner_key}/content/{content_id} | 
[**GetContentsByOwner**](OwnerContentAPI.md#GetContentsByOwner) | **Get** /owners/{owner_key}/content | 
[**RemoveContent**](OwnerContentAPI.md#RemoveContent) | **Delete** /owners/{owner_key}/content/{content_id} | 
[**UpdateContent**](OwnerContentAPI.md#UpdateContent) | **Put** /owners/{owner_key}/content/{content_id} | 



## CreateContent

> ContentDTO CreateContent(ctx, ownerKey).ContentDTO(contentDTO).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/joshmalbrecht/candlepin-go-client/pkg"
)

func main() {
	ownerKey := "ownerKey_example" // string | The key of the owner
	contentDTO := *openapiclient.NewContentDTO() // ContentDTO | Content to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerContentAPI.CreateContent(context.Background(), ownerKey).ContentDTO(contentDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerContentAPI.CreateContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContent`: ContentDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerContentAPI.CreateContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentDTO** | [**ContentDTO**](ContentDTO.md) | Content to be created | 

### Return type

[**ContentDTO**](ContentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContentBatch

> []ContentDTO CreateContentBatch(ctx, ownerKey).ContentDTO(contentDTO).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/joshmalbrecht/candlepin-go-client/pkg"
)

func main() {
	ownerKey := "ownerKey_example" // string | key of the owner whose content has to be created
	contentDTO := []openapiclient.ContentDTO{*openapiclient.NewContentDTO()} // []ContentDTO | List of contents to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerContentAPI.CreateContentBatch(context.Background(), ownerKey).ContentDTO(contentDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerContentAPI.CreateContentBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContentBatch`: []ContentDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerContentAPI.CreateContentBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | key of the owner whose content has to be created | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContentBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentDTO** | [**[]ContentDTO**](ContentDTO.md) | List of contents to be created | 

### Return type

[**[]ContentDTO**](ContentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentById

> ContentDTO GetContentById(ctx, ownerKey, contentId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/joshmalbrecht/candlepin-go-client/pkg"
)

func main() {
	ownerKey := "ownerKey_example" // string | The key of the owner whose content has to be retrieved
	contentId := "contentId_example" // string | The id of the content to be retrieved

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerContentAPI.GetContentById(context.Background(), ownerKey, contentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerContentAPI.GetContentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentById`: ContentDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerContentAPI.GetContentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner whose content has to be retrieved | 
**contentId** | **string** | The id of the content to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContentDTO**](ContentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentsByOwner

> []ContentDTO GetContentsByOwner(ctx, ownerKey).Content(content).OmitGlobal(omitGlobal).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/joshmalbrecht/candlepin-go-client/pkg"
)

func main() {
	ownerKey := "ownerKey_example" // string | Owner key
	content := []string{"Inner_example"} // []string | The ID of a content to fetch. If defined, the list of contents returned by this method will only include those matching the given ID. May be specified multiple times to filter on multiple content IDs.  (optional)
	omitGlobal := true // bool | whether or not to limit the lookup to only contents defined within the organization's namespace, excluding any globally defined contents  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerContentAPI.GetContentsByOwner(context.Background(), ownerKey).Content(content).OmitGlobal(omitGlobal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerContentAPI.GetContentsByOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentsByOwner`: []ContentDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerContentAPI.GetContentsByOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | Owner key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentsByOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | **[]string** | The ID of a content to fetch. If defined, the list of contents returned by this method will only include those matching the given ID. May be specified multiple times to filter on multiple content IDs.  | 
 **omitGlobal** | **bool** | whether or not to limit the lookup to only contents defined within the organization&#39;s namespace, excluding any globally defined contents  | [default to false]

### Return type

[**[]ContentDTO**](ContentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveContent

> RemoveContent(ctx, ownerKey, contentId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/joshmalbrecht/candlepin-go-client/pkg"
)

func main() {
	ownerKey := "ownerKey_example" // string | The key of the owner whose content has to be deleted
	contentId := "contentId_example" // string | The ID of the content to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OwnerContentAPI.RemoveContent(context.Background(), ownerKey, contentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerContentAPI.RemoveContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner whose content has to be deleted | 
**contentId** | **string** | The ID of the content to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContent

> ContentDTO UpdateContent(ctx, ownerKey, contentId).ContentDTO(contentDTO).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/joshmalbrecht/candlepin-go-client/pkg"
)

func main() {
	ownerKey := "ownerKey_example" // string | The key of the owner to be updated
	contentId := "contentId_example" // string | The ID of the Content to be updated
	contentDTO := *openapiclient.NewContentDTO() // ContentDTO | Content to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerContentAPI.UpdateContent(context.Background(), ownerKey, contentId).ContentDTO(contentDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerContentAPI.UpdateContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContent`: ContentDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerContentAPI.UpdateContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner to be updated | 
**contentId** | **string** | The ID of the Content to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentDTO** | [**ContentDTO**](ContentDTO.md) | Content to be updated | 

### Return type

[**ContentDTO**](ContentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

