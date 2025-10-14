# \LeaguesAPI

All URIs are relative to *https://open.faceit.com/data/v4*

| Method                                                                 | HTTP request                                                         | Description                                                          |
| ---------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| [**GetLeagueById**](LeaguesAPI.md#GetLeagueById)                       | **Get** /leagues/{league_id}                                         | Retrieve details of a league of a matchmaking on FACEIT              |
| [**GetLeagueSeason**](LeaguesAPI.md#GetLeagueSeason)                   | **Get** /leagues/{league_id}/seasons/{season_id}                     | Retrieve details of a season of a league on FACEIT                   |
| [**GetPlayerForLeagueSeason**](LeaguesAPI.md#GetPlayerForLeagueSeason) | **Get** /leagues/{league_id}/seasons/{season_id}/players/{player_id} | Retrieve details of a player for a given league and season on FACEIT |



## GetLeagueById

> League GetLeagueById(ctx, leagueId).Execute()

Retrieve details of a league of a matchmaking on FACEIT



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
	leagueId := "leagueId_example" // string | The id of the league

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaguesAPI.GetLeagueById(context.Background(), leagueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaguesAPI.GetLeagueById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeagueById`: League
	fmt.Fprintf(os.Stdout, "Response from `LeaguesAPI.GetLeagueById`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **leagueId** | **string**          | The id of the league                                                        |

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeagueByIdRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


### Return type

[**League**](League.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeagueSeason

> SeasonDetailed GetLeagueSeason(ctx, leagueId, seasonId).Execute()

Retrieve details of a season of a league on FACEIT



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
	leagueId := "leagueId_example" // string | The id of the league
	seasonId := int32(56) // int32 | The id of the season

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaguesAPI.GetLeagueSeason(context.Background(), leagueId, seasonId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaguesAPI.GetLeagueSeason``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeagueSeason`: SeasonDetailed
	fmt.Fprintf(os.Stdout, "Response from `LeaguesAPI.GetLeagueSeason`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **leagueId** | **string**          | The id of the league                                                        |
| **seasonId** | **int32**           | The id of the season                                                        |

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeagueSeasonRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**SeasonDetailed**](SeasonDetailed.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayerForLeagueSeason

> PlayerInLeague GetPlayerForLeagueSeason(ctx, leagueId, seasonId, playerId).Execute()

Retrieve details of a player for a given league and season on FACEIT



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
	leagueId := "leagueId_example" // string | The id of the league
	seasonId := int32(56) // int32 | The id of the season
	playerId := "playerId_example" // string | The id of the player

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaguesAPI.GetPlayerForLeagueSeason(context.Background(), leagueId, seasonId, playerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaguesAPI.GetPlayerForLeagueSeason``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayerForLeagueSeason`: PlayerInLeague
	fmt.Fprintf(os.Stdout, "Response from `LeaguesAPI.GetPlayerForLeagueSeason`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **leagueId** | **string**          | The id of the league                                                        |
| **seasonId** | **int32**           | The id of the season                                                        |
| **playerId** | **string**          | The id of the player                                                        |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerForLeagueSeasonRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |




### Return type

[**PlayerInLeague**](PlayerInLeague.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

