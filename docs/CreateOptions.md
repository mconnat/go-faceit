# CreateOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterVolumeSpec** | Pointer to [**ClusterVolumeSpec**](ClusterVolumeSpec.md) |  | [optional] 
**Driver** | Pointer to **string** | Name of the volume driver to use. | [optional] 
**DriverOpts** | Pointer to **map[string]string** | A mapping of driver options and values. These options are passed directly to the driver and are driver specific. | [optional] 
**Labels** | Pointer to **map[string]string** | User-defined key/value metadata. | [optional] 
**Name** | Pointer to **string** | The new volume&#39;s name. If not specified, Docker generates a name. | [optional] 

## Methods

### NewCreateOptions

`func NewCreateOptions() *CreateOptions`

NewCreateOptions instantiates a new CreateOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOptionsWithDefaults

`func NewCreateOptionsWithDefaults() *CreateOptions`

NewCreateOptionsWithDefaults instantiates a new CreateOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterVolumeSpec

`func (o *CreateOptions) GetClusterVolumeSpec() ClusterVolumeSpec`

GetClusterVolumeSpec returns the ClusterVolumeSpec field if non-nil, zero value otherwise.

### GetClusterVolumeSpecOk

`func (o *CreateOptions) GetClusterVolumeSpecOk() (*ClusterVolumeSpec, bool)`

GetClusterVolumeSpecOk returns a tuple with the ClusterVolumeSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterVolumeSpec

`func (o *CreateOptions) SetClusterVolumeSpec(v ClusterVolumeSpec)`

SetClusterVolumeSpec sets ClusterVolumeSpec field to given value.

### HasClusterVolumeSpec

`func (o *CreateOptions) HasClusterVolumeSpec() bool`

HasClusterVolumeSpec returns a boolean if a field has been set.

### GetDriver

`func (o *CreateOptions) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *CreateOptions) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *CreateOptions) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *CreateOptions) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetDriverOpts

`func (o *CreateOptions) GetDriverOpts() map[string]string`

GetDriverOpts returns the DriverOpts field if non-nil, zero value otherwise.

### GetDriverOptsOk

`func (o *CreateOptions) GetDriverOptsOk() (*map[string]string, bool)`

GetDriverOptsOk returns a tuple with the DriverOpts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverOpts

`func (o *CreateOptions) SetDriverOpts(v map[string]string)`

SetDriverOpts sets DriverOpts field to given value.

### HasDriverOpts

`func (o *CreateOptions) HasDriverOpts() bool`

HasDriverOpts returns a boolean if a field has been set.

### GetLabels

`func (o *CreateOptions) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateOptions) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateOptions) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateOptions) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *CreateOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOptions) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


