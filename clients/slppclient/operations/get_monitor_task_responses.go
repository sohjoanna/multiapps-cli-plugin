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

// GetMonitorTaskReader is a Reader for the GetMonitorTask structure.
type GetMonitorTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMonitorTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMonitorTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMonitorTaskOK creates a GetMonitorTaskOK with default headers values
func NewGetMonitorTaskOK() *GetMonitorTaskOK {
	return &GetMonitorTaskOK{}
}

/*GetMonitorTaskOK handles this case with default header values.

OK
*/
type GetMonitorTaskOK struct {
	Payload *models.Task
}

func (o *GetMonitorTaskOK) Error() string {
	return fmt.Sprintf("[GET /monitor/{taskId}][%d] getMonitorTaskOK  %+v", 200, o.Payload)
}

func (o *GetMonitorTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
