package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateServiceFiles Upload multiple files for a specific service
*/
func (a *Client) CreateServiceFiles(params *CreateServiceFilesParams, authInfo runtime.ClientAuthInfoWriter) (*CreateServiceFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateServiceFilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateServiceFiles",
		Method:             "POST",
		PathPattern:        "/services/{serviceId}/files",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateServiceFilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateServiceFilesOK), nil

}

/*
CreateServiceProcess create service process API
*/
func (a *Client) CreateServiceProcess(params *CreateServiceProcessParams, authInfo runtime.ClientAuthInfoWriter) (*CreateServiceProcessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateServiceProcessParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateServiceProcess",
		Method:             "POST",
		PathPattern:        "/services/{serviceId}/processes",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateServiceProcessReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateServiceProcessOK), nil

}

/*
DeleteProcess The deletion of an SL process results in the clean-up of the resources used by the process (e.g. deleting downloaded files, logs, parameter files, etc.)
*/
func (a *Client) DeleteProcess(params *DeleteProcessParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteProcessNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProcessParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteProcess",
		Method:             "DELETE",
		PathPattern:        "/processes/{processId}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteProcessReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteProcessNoContent), nil

}

/*
DeleteServiceFile Delete a single file for a specific service
*/
func (a *Client) DeleteServiceFile(params *DeleteServiceFileParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServiceFileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServiceFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteServiceFile",
		Method:             "DELETE",
		PathPattern:        "/services/{serviceId}/files/{fileId}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteServiceFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteServiceFileNoContent), nil

}

/*
DeleteServiceFiles Delete all files for a specific service
*/
func (a *Client) DeleteServiceFiles(params *DeleteServiceFilesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServiceFilesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServiceFilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteServiceFiles",
		Method:             "DELETE",
		PathPattern:        "/services/{serviceId}/files",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteServiceFilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteServiceFilesNoContent), nil

}

/*
DeleteServiceProcess The deletion of an SL process results in the clean-up of the resources used by the process (e.g. deleting downloaded files, logs, parameter files, etc.)
*/
func (a *Client) DeleteServiceProcess(params *DeleteServiceProcessParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServiceProcessNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServiceProcessParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteServiceProcess",
		Method:             "DELETE",
		PathPattern:        "/services/{serviceId}/processes/{processId}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteServiceProcessReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteServiceProcessNoContent), nil

}

/*
GetMetadata Metadata, relevant to this SL process manager (Mandatory)
*/
func (a *Client) GetMetadata(params *GetMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*GetMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMetadata",
		Method:             "GET",
		PathPattern:        "/metadata",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMetadataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMetadataOK), nil

}

/*
GetProcess get process API
*/
func (a *Client) GetProcess(params *GetProcessParams, authInfo runtime.ClientAuthInfoWriter) (*GetProcessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProcessParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetProcess",
		Method:             "GET",
		PathPattern:        "/processes/{processId}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProcessReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProcessOK), nil

}

/*
GetProcessService get process service API
*/
func (a *Client) GetProcessService(params *GetProcessServiceParams, authInfo runtime.ClientAuthInfoWriter) (*GetProcessServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProcessServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetProcessService",
		Method:             "GET",
		PathPattern:        "/processes/{processId}/service",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProcessServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProcessServiceOK), nil

}

/*
GetProcesses Processes that are running in the particular process manager
*/
func (a *Client) GetProcesses(params *GetProcessesParams, authInfo runtime.ClientAuthInfoWriter) (*GetProcessesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProcessesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetProcesses",
		Method:             "GET",
		PathPattern:        "/processes",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProcessesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProcessesOK), nil

}

/*
GetService get service API
*/
func (a *Client) GetService(params *GetServiceParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetService",
		Method:             "GET",
		PathPattern:        "/services/{serviceId}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetServiceOK), nil

}

/*
GetServiceFiles get service files API
*/
func (a *Client) GetServiceFiles(params *GetServiceFilesParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceFilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetServiceFiles",
		Method:             "GET",
		PathPattern:        "/services/{serviceId}/files",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServiceFilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetServiceFilesOK), nil

}

/*
GetServiceParameters get service parameters API
*/
func (a *Client) GetServiceParameters(params *GetServiceParametersParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceParametersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceParametersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetServiceParameters",
		Method:             "GET",
		PathPattern:        "/services/{serviceId}/parameters",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServiceParametersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetServiceParametersOK), nil

}

/*
GetServiceProcesses Retrieves the currently running and recently finished processes
*/
func (a *Client) GetServiceProcesses(params *GetServiceProcessesParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceProcessesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceProcessesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetServiceProcesses",
		Method:             "GET",
		PathPattern:        "/services/{serviceId}/processes",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServiceProcessesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetServiceProcessesOK), nil

}

/*
GetServiceVersions get service versions API
*/
func (a *Client) GetServiceVersions(params *GetServiceVersionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceVersionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetServiceVersions",
		Method:             "GET",
		PathPattern:        "/services/{serviceId}/versions",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServiceVersionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetServiceVersionsOK), nil

}

/*
GetServices Collection of metadata elements describing definitions of SL services known to a process manager (Mandatory)
*/
func (a *Client) GetServices(params *GetServicesParams, authInfo runtime.ClientAuthInfoWriter) (*GetServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetServices",
		Method:             "GET",
		PathPattern:        "/services",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServicesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetServicesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
