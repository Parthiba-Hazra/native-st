// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UploadFileCreatedCode is the HTTP code returned for type UploadFileCreated
const UploadFileCreatedCode int = 201

/*
UploadFileCreated File uploaded successfully

swagger:response uploadFileCreated
*/
type UploadFileCreated struct {

	/*
	  In: Body
	*/
	Payload *UploadFileCreatedBody `json:"body,omitempty"`
}

// NewUploadFileCreated creates UploadFileCreated with default headers values
func NewUploadFileCreated() *UploadFileCreated {

	return &UploadFileCreated{}
}

// WithPayload adds the payload to the upload file created response
func (o *UploadFileCreated) WithPayload(payload *UploadFileCreatedBody) *UploadFileCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload file created response
func (o *UploadFileCreated) SetPayload(payload *UploadFileCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadFileCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}