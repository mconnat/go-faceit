# \HubsAPI

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHub**](HubsAPI.md#GetHub) | **Get** /hubs/{hub_id} | Retrieve hub details
[**GetHubMatches**](HubsAPI.md#GetHubMatches) | **Get** /hubs/{hub_id}/matches | Retrieve all matches of a hub
[**GetHubMembers**](HubsAPI.md#GetHubMembers) | **Get** /hubs/{hub_id}/members | Retrieve all members of a hub
[**GetHubRoles**](HubsAPI.md#GetHubRoles) | **Get** /hubs/{hub_id}/roles | Retrieve all roles members can have in a hub
[**GetHubRules**](HubsAPI.md#GetHubRules) | **Get** /hubs/{hub_id}/rules | Retrieve rules of a hub
[**GetHubStats**](HubsAPI.md#GetHubStats) | **Get** /hubs/{hub_id}/stats | Retrieve statistics of a hub



## GetHub

> Hub GetHub(ctx, hubId).Expanded(expanded).Execute()

Retrieve hub details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	hubId := "hubId_example" // string | The id of the hub
	expanded := []string{"Expanded_example"} // []string | List of entity names to expand in request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HubsAPI.GetHub(context.Background(), hubId).Expanded(expanded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HubsAPI.GetHub``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHub`: Hub
	fmt.Fprintf(os.Stdout, "Response from `HubsAPI.GetHub`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hubId** | **string** | The id of the hub | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expanded** | **[]string** | List of entity names to expand in request | 

### Return type

[**Hub**](Hub.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHubMatches

> MatchList GetHubMatches(ctx, hubId).Type_(type_).Offset(offset).Limit(limit).Execute()

Retrieve all matches of a hub



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	hubId := "hubId_example" // string | The id of the hub
	type_ := "type__example" // string | Kind of matches to return. Can be all(default), upcoming, ongoing or past (optional)
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HubsAPI.GetHubMatches(context.Background(), hubId).Type_(type_).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HubsAPI.GetHubMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHubMatches`: MatchList
	fmt.Fprintf(os.Stdout, "Response from `HubsAPI.GetHubMatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hubId** | **string** | The id of the hub | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHubMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Kind of matches to return. Can be all(default), upcoming, ongoing or past | 
 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**MatchList**](MatchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHubMembers

> HubMembers GetHubMembers(ctx, hubId).Offset(offset).Limit(limit).Execute()

Retrieve all members of a hub



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	hubId := "hubId_example" // string | The id of the hub
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HubsAPI.GetHubMembers(context.Background(), hubId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HubsAPI.GetHubMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHubMembers`: HubMembers
	fmt.Fprintf(os.Stdout, "Response from `HubsAPI.GetHubMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hubId** | **string** | The id of the hub | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHubMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 50]

### Return type

[**HubMembers**](HubMembers.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHubRoles

> RolesList GetHubRoles(ctx, hubId).Offset(offset).Limit(limit).Execute()

Retrieve all roles members can have in a hub



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	hubId := "hubId_example" // string | The id of the hub
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HubsAPI.GetHubRoles(context.Background(), hubId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HubsAPI.GetHubRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHubRoles`: RolesList
	fmt.Fprintf(os.Stdout, "Response from `HubsAPI.GetHubRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hubId** | **string** | The id of the hub | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHubRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 50]

### Return type

[**RolesList**](RolesList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHubRules

> Rules GetHubRules(ctx, hubId).Execute()

Retrieve rules of a hub



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	hubId := "hubId_example" // string | The id of the hub

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HubsAPI.GetHubRules(context.Background(), hubId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HubsAPI.GetHubRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHubRules`: Rules
	fmt.Fprintf(os.Stdout, "Response from `HubsAPI.GetHubRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hubId** | **string** | The id of the hub | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHubRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Rules**](Rules.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHubStats

> HubStats GetHubStats(ctx, hubId).Offset(offset).Limit(limit).Execute()

Retrieve statistics of a hub



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	hubId := "hubId_example" // string | The id of the hub
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HubsAPI.GetHubStats(context.Background(), hubId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HubsAPI.GetHubStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHubStats`: HubStats
	fmt.Fprintf(os.Stdout, "Response from `HubsAPI.GetHubStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hubId** | **string** | The id of the hub | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHubStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**HubStats**](HubStats.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

