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

// GetRoadmapTaskReader is a Reader for the GetRoadmapTask structure.
type GetRoadmapTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoadmapTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRoadmapTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRoadmapTaskOK creates a GetRoadmapTaskOK with default headers values
func NewGetRoadmapTaskOK() *GetRoadmapTaskOK {
	return &GetRoadmapTaskOK{}
}

/*GetRoadmapTaskOK handles this case with default header values.

OK
*/
type GetRoadmapTaskOK struct {
	Payload *models.Task
}

func (o *GetRoadmapTaskOK) Error() string {
	return fmt.Sprintf("[GET /roadmap/{taskId}][%d] getRoadmapTaskOK  %+v", 200, o.Payload)
}

func (o *GetRoadmapTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}