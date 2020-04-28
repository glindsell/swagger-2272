// Code generated by go-swagger; DO NOT EDIT.

package hello_tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"hello/models"
)

// IDOfHelloEndpointReader is a Reader for the IDOfHelloEndpoint structure.
type IDOfHelloEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IDOfHelloEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIDOfHelloEndpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIDOfHelloEndpointOK creates a IDOfHelloEndpointOK with default headers values
func NewIDOfHelloEndpointOK() *IDOfHelloEndpointOK {
	return &IDOfHelloEndpointOK{}
}

/*IDOfHelloEndpointOK handles this case with default header values.

This text will appear as description of your response body.
*/
type IDOfHelloEndpointOK struct {
	Payload []*models.Hello
}

func (o *IDOfHelloEndpointOK) Error() string {
	return fmt.Sprintf("[POST /hello][%d] idOfHelloEndpointOK  %+v", 200, o.Payload)
}

func (o *IDOfHelloEndpointOK) GetPayload() []*models.Hello {
	return o.Payload
}

func (o *IDOfHelloEndpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
