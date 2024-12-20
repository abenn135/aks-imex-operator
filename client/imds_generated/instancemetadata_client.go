// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.10.3, generator: @autorest/go@4.0.0-preview.69)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package imds_generated

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// InstanceMetadataClient contains the methods for the InstanceMetadataClient group.
// Don't use this type directly, use a constructor function instead.
type InstanceMetadataClient struct {
	internal *azcore.Client
}

// GetVersions - Get supported API versions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - options - InstanceMetadataClientGetVersionsOptions contains the optional parameters for the InstanceMetadataClient.GetVersions
//     method.
func (client *InstanceMetadataClient) GetVersions(ctx context.Context, options *InstanceMetadataClientGetVersionsOptions) (InstanceMetadataClientGetVersionsResponse, error) {
	var err error
	req, err := client.getVersionsCreateRequest(ctx, options)
	if err != nil {
		return InstanceMetadataClientGetVersionsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InstanceMetadataClientGetVersionsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InstanceMetadataClientGetVersionsResponse{}, err
	}
	resp, err := client.getVersionsHandleResponse(httpResp)
	return resp, err
}

// getVersionsCreateRequest creates the GetVersions request.
func (client *InstanceMetadataClient) getVersionsCreateRequest(ctx context.Context, _ *InstanceMetadataClientGetVersionsOptions) (*policy.Request, error) {
	urlPath := "/versions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getVersionsHandleResponse handles the GetVersions response.
func (client *InstanceMetadataClient) getVersionsHandleResponse(resp *http.Response) (InstanceMetadataClientGetVersionsResponse, error) {
	result := InstanceMetadataClientGetVersionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Versions); err != nil {
		return InstanceMetadataClientGetVersionsResponse{}, err
	}
	return result, nil
}
