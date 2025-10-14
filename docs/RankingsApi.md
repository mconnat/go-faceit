# \RankingsAPI

All URIs are relative to *https://open.faceit.com/data/v4*

| Method                                                  | HTTP request                                                           | Description                                            |
| ------------------------------------------------------- | ---------------------------------------------------------------------- | ------------------------------------------------------ |
| [**GetGlobalRanking**](RankingsAPI.md#GetGlobalRanking) | **Get** /rankings/games/{game_id}/regions/{region}                     | Retrieve global ranking of a game                      |
| [**GetPlayerRanking**](RankingsAPI.md#GetPlayerRanking) | **Get** /rankings/games/{game_id}/regions/{region}/players/{player_id} | Retrieve user position in the global ranking of a game |



## GetGlobalRanking

> GlobalRankingList GetGlobalRanking(ctx, gameId, region).Country(country).Offset(offset).Limit(limit).Execute()

Retrieve global ranking of a game



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
	gameId := "gameId_example" // string | The id of the game
	region := "region_example" // string | A region of a game
	country := "country_example" // string | A country code (ISO 3166-1) (optional)
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RankingsAPI.GetGlobalRanking(context.Background(), gameId, region).Country(country).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RankingsAPI.GetGlobalRanking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalRanking`: GlobalRankingList
	fmt.Fprintf(os.Stdout, "Response from `RankingsAPI.GetGlobalRanking`: %v\n", resp)
}
```

### Path Parameters


| Name       | Type                | Description                                                                 | Notes |
| ---------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **gameId** | **string**          | The id of the game                                                          |
| **region** | **string**          | A region of a game                                                          |

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalRankingRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **country** | **string** | A country code (ISO 3166-1) | 
 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**GlobalRankingList**](GlobalRankingList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayerRanking

> PlayerGlobalRanking GetPlayerRanking(ctx, gameId, region, playerId).Country(country).Limit(limit).Execute()

Retrieve user position in the global ranking of a game



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
	gameId := "gameId_example" // string | The id of the game
	region := "region_example" // string | A region of a game
	playerId := "playerId_example" // string | The id of a player
	country := "country_example" // string | A country code (ISO 3166-1) (optional)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RankingsAPI.GetPlayerRanking(context.Background(), gameId, region, playerId).Country(country).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RankingsAPI.GetPlayerRanking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayerRanking`: PlayerGlobalRanking
	fmt.Fprintf(os.Stdout, "Response from `RankingsAPI.GetPlayerRanking`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **gameId**   | **string**          | The id of the game                                                          |
| **region**   | **string**          | A region of a game                                                          |
| **playerId** | **string**          | The id of a player                                                          |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerRankingRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



 **country** | **string** | A country code (ISO 3166-1) | 
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**PlayerGlobalRanking**](PlayerGlobalRanking.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

