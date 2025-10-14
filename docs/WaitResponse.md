# WaitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**WaitExitError**](WaitExitError.md) |  | [optional] 
**StatusCode** | **int64** | Exit code of the container | 

## Methods

### NewWaitResponse

`func NewWaitResponse(statusCode int64, ) *WaitResponse`

NewWaitResponse instantiates a new WaitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaitResponseWithDefaults

`func NewWaitResponseWithDefaults() *WaitResponse`

NewWaitResponseWithDefaults instantiates a new WaitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *WaitResponse) GetError() WaitExitError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WaitResponse) GetErrorOk() (*WaitExitError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WaitResponse) SetError(v WaitExitError)`

SetError sets Error field to given value.

### HasError

`func (o *WaitResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStatusCode

`func (o *WaitResponse) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *WaitResponse) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *WaitResponse) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


