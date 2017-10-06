package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/SAP/cf-mta-plugin/clients/models"
)

// GetOperationReader is a Reader for the GetOperation structure.
type GetOperationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOperationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOperationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOperationOK creates a GetOperationOK with default headers values
func NewGetOperationOK() *GetOperationOK {
	return &GetOperationOK{}
}

/*GetOperationOK handles this case with default header values.

OK
*/
type GetOperationOK struct {
	Payload *models.Operation
}

func (o *GetOperationOK) Error() string {
	return fmt.Sprintf("[GET /operations/{processId}][%d] getOperationOK  %+v", 200, o.Payload)
}

func (o *GetOperationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Operation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
