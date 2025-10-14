# \MatchesAPI

All URIs are relative to *https://open.faceit.com/data/v4*

| Method                                           | HTTP request                      | Description                    |
| ------------------------------------------------ | --------------------------------- | ------------------------------ |
| [**GetMatch**](MatchesAPI.md#GetMatch)           | **Get** /matches/{match_id}       | Retrieve match details         |
| [**GetMatchStats**](MatchesAPI.md#GetMatchStats) | **Get** /matches/{match_id}/stats | Retrieve statistics of a match |



## GetMatch

> Match GetMatch(ctx, matchId).Execute()

Retrieve match details



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
	matchId := "matchId_example" // string | The id of the match

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MatchesAPI.GetMatch(context.Background(), matchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MatchesAPI.GetMatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMatch`: Match
	fmt.Fprintf(os.Stdout, "Response from `MatchesAPI.GetMatch`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **matchId** | **string**          | The id of the match                                                         |

### Other Parameters

Other parameters are passed through a pointer to a apiGetMatchRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


### Return type

[**Match**](Match.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMatchStats

> MatchStats GetMatchStats(ctx, matchId).Execute()

Retrieve statistics of a match



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
	matchId := "matchId_example" // string | The id of the match

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MatchesAPI.GetMatchStats(context.Background(), matchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MatchesAPI.GetMatchStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMatchStats`: MatchStats
	fmt.Fprintf(os.Stdout, "Response from `MatchesAPI.GetMatchStats`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **matchId** | **string**          | The id of the match                                                         |

### Other Parameters

Other parameters are passed through a pointer to a apiGetMatchStatsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


### Return type

[**MatchStats**](MatchStats.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

