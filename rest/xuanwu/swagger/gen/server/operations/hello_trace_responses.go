// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ExcitingFrog/xuanwu/swagger/gen/models"
)

// HelloTraceOKCode is the HTTP code returned for type HelloTraceOK
const HelloTraceOKCode int = 200

/*
HelloTraceOK Hello

swagger:response helloTraceOK
*/
type HelloTraceOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewHelloTraceOK creates HelloTraceOK with default headers values
func NewHelloTraceOK() *HelloTraceOK {

	return &HelloTraceOK{}
}

// WithPayload adds the payload to the hello trace o k response
func (o *HelloTraceOK) WithPayload(payload string) *HelloTraceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the hello trace o k response
func (o *HelloTraceOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HelloTraceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// HelloTraceBadRequestCode is the HTTP code returned for type HelloTraceBadRequest
const HelloTraceBadRequestCode int = 400

/*
HelloTraceBadRequest BadRequest

swagger:response helloTraceBadRequest
*/
type HelloTraceBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewHelloTraceBadRequest creates HelloTraceBadRequest with default headers values
func NewHelloTraceBadRequest() *HelloTraceBadRequest {

	return &HelloTraceBadRequest{}
}

// WithPayload adds the payload to the hello trace bad request response
func (o *HelloTraceBadRequest) WithPayload(payload *models.ErrorResponse) *HelloTraceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the hello trace bad request response
func (o *HelloTraceBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HelloTraceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}