// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.10.3, generator: @autorest/go@4.0.0-preview.69)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package imds_generated

// AttestedClientGetDocumentResponse contains the response from method AttestedClient.GetDocument.
type AttestedClientGetDocumentResponse struct {
// This is the response from the Attested_GetDocument operation.
	AttestedData
}

// IdentityClientGetInfoResponse contains the response from method IdentityClient.GetInfo.
type IdentityClientGetInfoResponse struct {
// This is the response from the Identity_GetInfo operation.
	IdentityInfoResponse
}

// IdentityClientGetTokenResponse contains the response from method IdentityClient.GetToken.
type IdentityClientGetTokenResponse struct {
// This is the response from the Identity_GetToken operation.
	IdentityTokenResponse
}

// InstanceMetadataClientGetVersionsResponse contains the response from method InstanceMetadataClient.GetVersions.
type InstanceMetadataClientGetVersionsResponse struct {
// This is the response from the GetVersions operation.
	Versions
}

// InstancesClientGetMetadataResponse contains the response from method InstancesClient.GetMetadata.
type InstancesClientGetMetadataResponse struct {
// This is the response from the Instance_GetMetadata operation.
	Instance
}
