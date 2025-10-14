# ServiceCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | The ID of the created service. | [optional] 
**Warnings** | Pointer to **[]string** | Optional warning message.  FIXME(thaJeztah): this should have \&quot;omitempty\&quot; in the generated type. | [optional] 

## Methods

### NewServiceCreateResponse

`func NewServiceCreateResponse() *ServiceCreateResponse`

NewServiceCreateResponse instantiates a new ServiceCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCreateResponseWithDefaults

`func NewServiceCreateResponseWithDefaults() *ServiceCreateResponse`

NewServiceCreateResponseWithDefaults instantiates a new ServiceCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *ServiceCreateResponse) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ServiceCreateResponse) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ServiceCreateResponse) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *ServiceCreateResponse) HasID() bool`

HasID returns a boolean if a field has been set.

### GetWarnings

`func (o *ServiceCreateResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ServiceCreateResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ServiceCreateResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ServiceCreateResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


