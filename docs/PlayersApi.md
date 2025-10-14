# \PlayersAPI

All URIs are relative to *https://open.faceit.com/data/v4*

| Method                                                         | HTTP request                                       | Description                                                   |
| -------------------------------------------------------------- | -------------------------------------------------- | ------------------------------------------------------------- |
| [**GetPlayer**](PlayersAPI.md#GetPlayer)                       | **Get** /players/{player_id}                       | Retrieve player details                                       |
| [**GetPlayerBans**](PlayersAPI.md#GetPlayerBans)               | **Get** /players/{player_id}/bans                  | Retrieve all bans of a player                                 |
| [**GetPlayerFromLookup**](PlayersAPI.md#GetPlayerFromLookup)   | **Get** /players                                   | Retrieve player details                                       |
| [**GetPlayerHistory**](PlayersAPI.md#GetPlayerHistory)         | **Get** /players/{player_id}/history               | Retrieve all matches of a player                              |
| [**GetPlayerHubs**](PlayersAPI.md#GetPlayerHubs)               | **Get** /players/{player_id}/hubs                  | Retrieve all hubs of a player                                 |
| [**GetPlayerStats**](PlayersAPI.md#GetPlayerStats)             | **Get** /players/{player_id}/games/{game_id}/stats | Retrieve statistics of a player for a given amount of matches |
| [**GetPlayerStats_0**](PlayersAPI.md#GetPlayerStats_0)         | **Get** /players/{player_id}/stats/{game_id}       | Retrieve statistics of a player                               |
| [**GetPlayerTeams**](PlayersAPI.md#GetPlayerTeams)             | **Get** /players/{player_id}/teams                 | Retrieve all teams of a player                                |
| [**GetPlayerTournaments**](PlayersAPI.md#GetPlayerTournaments) | **Get** /players/{player_id}/tournaments           | Retrieve all tournaments of a player                          |



## GetPlayer

> Player GetPlayer(ctx, playerId).Execute()

Retrieve player details



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
	playerId := "playerId_example" // string | The id of the player

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayer(context.Background(), playerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayer`: Player
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayer`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **playerId** | **string**          | The id of the player                                                        |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


### Return type

[**Player**](Player.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayerBans

> PlayerBansList GetPlayerBans(ctx, playerId).Offset(offset).Limit(limit).Execute()

Retrieve all bans of a player



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
	playerId := "playerId_example" // string | The id of the player
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayerBans(context.Background(), playerId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayerBans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayerBans`: PlayerBansList
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayerBans`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **playerId** | **string**          | The id of the player                                                        |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerBansRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**PlayerBansList**](PlayerBansList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayerFromLookup

> Player GetPlayerFromLookup(ctx).Nickname(nickname).Game(game).GamePlayerId(gamePlayerId).Execute()

Retrieve player details



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
	nickname := "nickname_example" // string | The nickname of the player on FACEIT (optional)
	game := "game_example" // string | A game on FACEIT (optional)
	gamePlayerId := "gamePlayerId_example" // string | The ID of a player on game's platform (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayerFromLookup(context.Background()).Nickname(nickname).Game(game).GamePlayerId(gamePlayerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayerFromLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayerFromLookup`: Player
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayerFromLookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerFromLookupRequest struct via the builder pattern


| Name             | Type       | Description                               | Notes |
| ---------------- | ---------- | ----------------------------------------- | ----- |
| **nickname**     | **string** | The nickname of the player on FACEIT      |
| **game**         | **string** | A game on FACEIT                          |
| **gamePlayerId** | **string** | The ID of a player on game&#39;s platform |

### Return type

[**Player**](Player.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayerHistory

> MatchHistoryList GetPlayerHistory(ctx, playerId).Game(game).From(from).To(to).Offset(offset).Limit(limit).Execute()

Retrieve all matches of a player



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
	playerId := "playerId_example" // string | The id of the player
	game := "game_example" // string | A game on FACEIT
	from := int32(56) // int32 | The timestamp (Unix time) as lower bound of the query. 1 month ago if not specified (optional)
	to := int32(56) // int32 | The timestamp (Unix time) as higher bound of the query. Current timestamp if not specified (optional)
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayerHistory(context.Background(), playerId).Game(game).From(from).To(to).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayerHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayerHistory`: MatchHistoryList
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayerHistory`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **playerId** | **string**          | The id of the player                                                        |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerHistoryRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **game** | **string** | A game on FACEIT | 
 **from** | **int32** | The timestamp (Unix time) as lower bound of the query. 1 month ago if not specified | 
 **to** | **int32** | The timestamp (Unix time) as higher bound of the query. Current timestamp if not specified | 
 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**MatchHistoryList**](MatchHistoryList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayerHubs

> HubsList GetPlayerHubs(ctx, playerId).Offset(offset).Limit(limit).Execute()

Retrieve all hubs of a player



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
	playerId := "playerId_example" // string | The id of the player
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayerHubs(context.Background(), playerId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayerHubs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayerHubs`: HubsList
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayerHubs`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **playerId** | **string**          | The id of the player                                                        |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerHubsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 50]

### Return type

[**HubsList**](HubsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayerStats

> PlayerStatsForMatchesList GetPlayerStats(ctx, playerId, gameId).Offset(offset).Limit(limit).From(from).To(to).Execute()

Retrieve statistics of a player for a given amount of matches



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
	playerId := "playerId_example" // string | The id of the player
	gameId := "gameId_example" // string | A game on FACEIT
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)
	from := int32(56) // int32 | Used to filter the dataset by date (minimum). Expected value is date (\"items.stats.Match Finished At\") in epoch milliseconds.  (optional)
	to := int32(56) // int32 | Used to filter the dataset by date (maximum). Expected value is date (\"items.stats.Match Finished At\") in epoch milliseconds.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayerStats(context.Background(), playerId, gameId).Offset(offset).Limit(limit).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayerStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayerStats`: PlayerStatsForMatchesList
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayerStats`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **playerId** | **string**          | The id of the player                                                        |
| **gameId**   | **string**          | A game on FACEIT                                                            |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerStatsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]
 **from** | **int32** | Used to filter the dataset by date (minimum). Expected value is date (\&quot;items.stats.Match Finished At\&quot;) in epoch milliseconds.  | 
 **to** | **int32** | Used to filter the dataset by date (maximum). Expected value is date (\&quot;items.stats.Match Finished At\&quot;) in epoch milliseconds.  | 

### Return type

[**PlayerStatsForMatchesList**](PlayerStatsForMatchesList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayerStats_0

> PlayerStats GetPlayerStats_0(ctx, playerId, gameId).Execute()

Retrieve statistics of a player



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
	playerId := "playerId_example" // string | The id of the player
	gameId := "gameId_example" // string | A game on FACEIT

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayerStats_0(context.Background(), playerId, gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayerStats_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayerStats_0`: PlayerStats
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayerStats_0`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **playerId** | **string**          | The id of the player                                                        |
| **gameId**   | **string**          | A game on FACEIT                                                            |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerStats_1Request struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**PlayerStats**](PlayerStats.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayerTeams

> TeamList GetPlayerTeams(ctx, playerId).Offset(offset).Limit(limit).Execute()

Retrieve all teams of a player



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
	playerId := "playerId_example" // string | The id of the player
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayerTeams(context.Background(), playerId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayerTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayerTeams`: TeamList
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayerTeams`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **playerId** | **string**          | The id of the player                                                        |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerTeamsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**TeamList**](TeamList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayerTournaments

> TournamentsList GetPlayerTournaments(ctx, playerId).Offset(offset).Limit(limit).Execute()

Retrieve all tournaments of a player



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
	playerId := "playerId_example" // string | The id of the player
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayersAPI.GetPlayerTournaments(context.Background(), playerId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayersAPI.GetPlayerTournaments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayerTournaments`: TournamentsList
	fmt.Fprintf(os.Stdout, "Response from `PlayersAPI.GetPlayerTournaments`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **playerId** | **string**          | The id of the player                                                        |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerTournamentsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**TournamentsList**](TournamentsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

