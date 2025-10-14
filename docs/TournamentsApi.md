# \TournamentsAPI

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTournament**](TournamentsAPI.md#GetTournament) | **Get** /tournaments/{tournament_id} | Retrieve tournament details
[**GetTournamentBrackets**](TournamentsAPI.md#GetTournamentBrackets) | **Get** /tournaments/{tournament_id}/brackets | Retrieve brackets of a tournament
[**GetTournamentMatches**](TournamentsAPI.md#GetTournamentMatches) | **Get** /tournaments/{tournament_id}/matches | Retrieve all matches of a tournament
[**GetTournamentTeams**](TournamentsAPI.md#GetTournamentTeams) | **Get** /tournaments/{tournament_id}/teams | Retrieve all teams of a tournament
[**GetTournamentsList**](TournamentsAPI.md#GetTournamentsList) | **Get** /tournaments | Retrieve tournaments v1 (no longer used)



## GetTournament

> Tournament GetTournament(ctx, tournamentId).Expanded(expanded).Execute()

Retrieve tournament details



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
	tournamentId := "tournamentId_example" // string | The id of the tournament
	expanded := []string{"Expanded_example"} // []string | List of entity names to expand in request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsAPI.GetTournament(context.Background(), tournamentId).Expanded(expanded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsAPI.GetTournament``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTournament`: Tournament
	fmt.Fprintf(os.Stdout, "Response from `TournamentsAPI.GetTournament`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tournamentId** | **string** | The id of the tournament | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTournamentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expanded** | **[]string** | List of entity names to expand in request | 

### Return type

[**Tournament**](Tournament.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTournamentBrackets

> Brackets GetTournamentBrackets(ctx, tournamentId).Execute()

Retrieve brackets of a tournament



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
	tournamentId := "tournamentId_example" // string | The id of the tournament

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsAPI.GetTournamentBrackets(context.Background(), tournamentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsAPI.GetTournamentBrackets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTournamentBrackets`: Brackets
	fmt.Fprintf(os.Stdout, "Response from `TournamentsAPI.GetTournamentBrackets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tournamentId** | **string** | The id of the tournament | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTournamentBracketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Brackets**](Brackets.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTournamentMatches

> MatchList GetTournamentMatches(ctx, tournamentId).Offset(offset).Limit(limit).Execute()

Retrieve all matches of a tournament



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
	tournamentId := "tournamentId_example" // string | The id of the tournament
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsAPI.GetTournamentMatches(context.Background(), tournamentId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsAPI.GetTournamentMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTournamentMatches`: MatchList
	fmt.Fprintf(os.Stdout, "Response from `TournamentsAPI.GetTournamentMatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tournamentId** | **string** | The id of the tournament | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTournamentMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetTournamentTeams

> TournamentTeams GetTournamentTeams(ctx, tournamentId).Offset(offset).Limit(limit).Execute()

Retrieve all teams of a tournament



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
	tournamentId := "tournamentId_example" // string | The id of the tournament
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsAPI.GetTournamentTeams(context.Background(), tournamentId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsAPI.GetTournamentTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTournamentTeams`: TournamentTeams
	fmt.Fprintf(os.Stdout, "Response from `TournamentsAPI.GetTournamentTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tournamentId** | **string** | The id of the tournament | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTournamentTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**TournamentTeams**](TournamentTeams.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTournamentsList

> TournamentsList GetTournamentsList(ctx).Game(game).Region(region).Offset(offset).Limit(limit).Execute()

Retrieve tournaments v1 (no longer used)



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
	game := "game_example" // string | A game on FACEIT (optional)
	region := "region_example" // string | A region of the game (optional)
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsAPI.GetTournamentsList(context.Background()).Game(game).Region(region).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsAPI.GetTournamentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTournamentsList`: TournamentsList
	fmt.Fprintf(os.Stdout, "Response from `TournamentsAPI.GetTournamentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTournamentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **game** | **string** | A game on FACEIT | 
 **region** | **string** | A region of the game | 
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

