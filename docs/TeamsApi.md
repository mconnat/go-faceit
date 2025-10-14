# \TeamsAPI

All URIs are relative to *https://open.faceit.com/data/v4*

| Method                                                   | HTTP request                             | Description                    |
| -------------------------------------------------------- | ---------------------------------------- | ------------------------------ |
| [**GetTeam**](TeamsAPI.md#GetTeam)                       | **Get** /teams/{team_id}                 | Retrieve team details          |
| [**GetTeamStats**](TeamsAPI.md#GetTeamStats)             | **Get** /teams/{team_id}/stats/{game_id} | Retrieve statistics of a team  |
| [**GetTeamTournaments**](TeamsAPI.md#GetTeamTournaments) | **Get** /teams/{team_id}/tournaments     | Retrieve tournaments of a team |



## GetTeam

> Team GetTeam(ctx, teamId).Execute()

Retrieve team details



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
	teamId := "teamId_example" // string | The id of the team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetTeam(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeam`: Team
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetTeam`: %v\n", resp)
}
```

### Path Parameters


| Name       | Type                | Description                                                                 | Notes |
| ---------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **teamId** | **string**          | The id of the team                                                          |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


### Return type

[**Team**](Team.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamStats

> TeamStats GetTeamStats(ctx, teamId, gameId).Execute()

Retrieve statistics of a team



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
	teamId := "teamId_example" // string | The id of the team
	gameId := "gameId_example" // string | A game on FACEIT

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetTeamStats(context.Background(), teamId, gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetTeamStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamStats`: TeamStats
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetTeamStats`: %v\n", resp)
}
```

### Path Parameters


| Name       | Type                | Description                                                                 | Notes |
| ---------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **teamId** | **string**          | The id of the team                                                          |
| **gameId** | **string**          | A game on FACEIT                                                            |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamStatsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**TeamStats**](TeamStats.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamTournaments

> TournamentsList GetTeamTournaments(ctx, teamId).Offset(offset).Limit(limit).Execute()

Retrieve tournaments of a team



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
	teamId := "teamId_example" // string | The id of the team
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetTeamTournaments(context.Background(), teamId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetTeamTournaments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamTournaments`: TournamentsList
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetTeamTournaments`: %v\n", resp)
}
```

### Path Parameters


| Name       | Type                | Description                                                                 | Notes |
| ---------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **teamId** | **string**          | The id of the team                                                          |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamTournamentsRequest struct via the builder pattern


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

