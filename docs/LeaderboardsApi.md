# \LeaderboardsAPI

All URIs are relative to *https://open.faceit.com/data/v4*

| Method                                                                                | HTTP request                                                         | Description                                 |
| ------------------------------------------------------------------------------------- | -------------------------------------------------------------------- | ------------------------------------------- |
| [**GetChampionshipGroupRanking**](LeaderboardsAPI.md#GetChampionshipGroupRanking)     | **Get** /leaderboards/championships/{championship_id}/groups/{group} | Retrieve group ranking of a championship    |
| [**GetChampionshipLeaderboards**](LeaderboardsAPI.md#GetChampionshipLeaderboards)     | **Get** /leaderboards/championships/{championship_id}                | Retrieve all leaderboards of a championship |
| [**GetHubLeaderboards**](LeaderboardsAPI.md#GetHubLeaderboards)                       | **Get** /leaderboards/hubs/{hub_id}                                  | Retrieve all leaderboards of a hub          |
| [**GetHubRanking**](LeaderboardsAPI.md#GetHubRanking)                                 | **Get** /leaderboards/hubs/{hub_id}/general                          | Retrieve all time ranking of a hub          |
| [**GetHubSeasonRanking**](LeaderboardsAPI.md#GetHubSeasonRanking)                     | **Get** /leaderboards/hubs/{hub_id}/seasons/{season}                 | Retrieve seasonal ranking of a hub          |
| [**GetLeaderboard**](LeaderboardsAPI.md#GetLeaderboard)                               | **Get** /leaderboards/{leaderboard_id}                               | Retrieve ranking from a leaderboard id      |
| [**GetPlayerRankingInLeaderboard**](LeaderboardsAPI.md#GetPlayerRankingInLeaderboard) | **Get** /leaderboards/{leaderboard_id}/players/{player_id}           | Retrieve a players ranking in a leaderboard |



## GetChampionshipGroupRanking

> EntityRanking GetChampionshipGroupRanking(ctx, championshipId, group).Offset(offset).Limit(limit).Execute()

Retrieve group ranking of a championship



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
	championshipId := "championshipId_example" // string | The id of the championship
	group := int32(56) // int32 | A group of the championship
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardsAPI.GetChampionshipGroupRanking(context.Background(), championshipId, group).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardsAPI.GetChampionshipGroupRanking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChampionshipGroupRanking`: EntityRanking
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardsAPI.GetChampionshipGroupRanking`: %v\n", resp)
}
```

### Path Parameters


| Name               | Type                | Description                                                                 | Notes |
| ------------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**            | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **championshipId** | **string**          | The id of the championship                                                  |
| **group**          | **int32**           | A group of the championship                                                 |

### Other Parameters

Other parameters are passed through a pointer to a apiGetChampionshipGroupRankingRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**EntityRanking**](EntityRanking.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChampionshipLeaderboards

> LeaderboardsList GetChampionshipLeaderboards(ctx, championshipId).Offset(offset).Limit(limit).Execute()

Retrieve all leaderboards of a championship



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
	championshipId := "championshipId_example" // string | The id of the championship
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardsAPI.GetChampionshipLeaderboards(context.Background(), championshipId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardsAPI.GetChampionshipLeaderboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChampionshipLeaderboards`: LeaderboardsList
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardsAPI.GetChampionshipLeaderboards`: %v\n", resp)
}
```

### Path Parameters


| Name               | Type                | Description                                                                 | Notes |
| ------------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**            | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **championshipId** | **string**          | The id of the championship                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetChampionshipLeaderboardsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**LeaderboardsList**](LeaderboardsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHubLeaderboards

> LeaderboardsList GetHubLeaderboards(ctx, hubId).Offset(offset).Limit(limit).Execute()

Retrieve all leaderboards of a hub



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
	hubId := "hubId_example" // string | The id of the hub
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardsAPI.GetHubLeaderboards(context.Background(), hubId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardsAPI.GetHubLeaderboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHubLeaderboards`: LeaderboardsList
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardsAPI.GetHubLeaderboards`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **hubId** | **string**          | The id of the hub                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiGetHubLeaderboardsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**LeaderboardsList**](LeaderboardsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHubRanking

> EntityRanking GetHubRanking(ctx, hubId).Offset(offset).Limit(limit).Execute()

Retrieve all time ranking of a hub



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
	hubId := "hubId_example" // string | The id of the hub
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardsAPI.GetHubRanking(context.Background(), hubId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardsAPI.GetHubRanking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHubRanking`: EntityRanking
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardsAPI.GetHubRanking`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **hubId** | **string**          | The id of the hub                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiGetHubRankingRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**EntityRanking**](EntityRanking.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHubSeasonRanking

> EntityRanking GetHubSeasonRanking(ctx, hubId, season).Offset(offset).Limit(limit).Execute()

Retrieve seasonal ranking of a hub



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
	hubId := "hubId_example" // string | The id of the hub
	season := int32(56) // int32 | A season of the hub
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardsAPI.GetHubSeasonRanking(context.Background(), hubId, season).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardsAPI.GetHubSeasonRanking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHubSeasonRanking`: EntityRanking
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardsAPI.GetHubSeasonRanking`: %v\n", resp)
}
```

### Path Parameters


| Name       | Type                | Description                                                                 | Notes |
| ---------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **hubId**  | **string**          | The id of the hub                                                           |
| **season** | **int32**           | A season of the hub                                                         |

### Other Parameters

Other parameters are passed through a pointer to a apiGetHubSeasonRankingRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**EntityRanking**](EntityRanking.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaderboard

> EntityRanking GetLeaderboard(ctx, leaderboardId).Offset(offset).Limit(limit).Execute()

Retrieve ranking from a leaderboard id



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
	leaderboardId := "leaderboardId_example" // string | The id of the leaderboard
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardsAPI.GetLeaderboard(context.Background(), leaderboardId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardsAPI.GetLeaderboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaderboard`: EntityRanking
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardsAPI.GetLeaderboard`: %v\n", resp)
}
```

### Path Parameters


| Name              | Type                | Description                                                                 | Notes |
| ----------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**           | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **leaderboardId** | **string**          | The id of the leaderboard                                                   |

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaderboardRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**EntityRanking**](EntityRanking.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayerRankingInLeaderboard

> Ranking GetPlayerRankingInLeaderboard(ctx, leaderboardId, playerId).Execute()

Retrieve a players ranking in a leaderboard



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
	leaderboardId := "leaderboardId_example" // string | The id of the leaderboard
	playerId := "playerId_example" // string | The id of the player

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardsAPI.GetPlayerRankingInLeaderboard(context.Background(), leaderboardId, playerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardsAPI.GetPlayerRankingInLeaderboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayerRankingInLeaderboard`: Ranking
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardsAPI.GetPlayerRankingInLeaderboard`: %v\n", resp)
}
```

### Path Parameters


| Name              | Type                | Description                                                                 | Notes |
| ----------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**           | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **leaderboardId** | **string**          | The id of the leaderboard                                                   |
| **playerId**      | **string**          | The id of the player                                                        |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerRankingInLeaderboardRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**Ranking**](Ranking.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

