# \MatchmakingsAPI

All URIs are relative to *https://open.faceit.com/data/v4*

| Method                                                  | HTTP request                           | Description                                           |
| ------------------------------------------------------- | -------------------------------------- | ----------------------------------------------------- |
| [**GetMatchmaking**](MatchmakingsAPI.md#GetMatchmaking) | **Get** /matchmakings/{matchmaking_id} | Retrieve details of a matchmaking of a game on FACEIT |



## GetMatchmaking

> Matchmaking GetMatchmaking(ctx, matchmakingId).Execute()

Retrieve details of a matchmaking of a game on FACEIT



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mconnat/go-faceit"
)

func main() {
	matchmakingId := "matchmakingId_example" // string | The id of the matchmaking

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MatchmakingsAPI.GetMatchmaking(context.Background(), matchmakingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MatchmakingsAPI.GetMatchmaking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMatchmaking`: Matchmaking
	fmt.Fprintf(os.Stdout, "Response from `MatchmakingsAPI.GetMatchmaking`: %v\n", resp)
}
```

### Path Parameters


| Name              | Type                | Description                                                                 | Notes |
| ----------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**           | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **matchmakingId** | **string**          | The id of the matchmaking                                                   |

### Other Parameters

Other parameters are passed through a pointer to a apiGetMatchmakingRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


### Return type

[**Matchmaking**](Matchmaking.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

