# Secret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key is the name of the key of the key-value pair passed to the plugin. | [optional] 
**Secret** | Pointer to **string** | Secret is the swarm Secret object from which to read data. This can be a Secret name or ID. The Secret data is retrieved by Swarm and used as the value of the key-value pair passed to the plugin. | [optional] 

## Methods

### NewSecret

`func NewSecret() *Secret`

NewSecret instantiates a new Secret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretWithDefaults

`func NewSecretWithDefaults() *Secret`

NewSecretWithDefaults instantiates a new Secret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Secret) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Secret) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Secret) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Secret) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSecret

`func (o *Secret) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Secret) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Secret) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *Secret) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


