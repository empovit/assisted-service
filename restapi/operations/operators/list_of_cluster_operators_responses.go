// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// ListOfClusterOperatorsOKCode is the HTTP code returned for type ListOfClusterOperatorsOK
const ListOfClusterOperatorsOKCode int = 200

/*ListOfClusterOperatorsOK Success.

swagger:response listOfClusterOperatorsOK
*/
type ListOfClusterOperatorsOK struct {

	/*
	  In: Body
	*/
	Payload models.MonitoredOperatorsList `json:"body,omitempty"`
}

// NewListOfClusterOperatorsOK creates ListOfClusterOperatorsOK with default headers values
func NewListOfClusterOperatorsOK() *ListOfClusterOperatorsOK {

	return &ListOfClusterOperatorsOK{}
}

// WithPayload adds the payload to the list of cluster operators o k response
func (o *ListOfClusterOperatorsOK) WithPayload(payload models.MonitoredOperatorsList) *ListOfClusterOperatorsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list of cluster operators o k response
func (o *ListOfClusterOperatorsOK) SetPayload(payload models.MonitoredOperatorsList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListOfClusterOperatorsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.MonitoredOperatorsList{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListOfClusterOperatorsUnauthorizedCode is the HTTP code returned for type ListOfClusterOperatorsUnauthorized
const ListOfClusterOperatorsUnauthorizedCode int = 401

/*ListOfClusterOperatorsUnauthorized Unauthorized.

swagger:response listOfClusterOperatorsUnauthorized
*/
type ListOfClusterOperatorsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewListOfClusterOperatorsUnauthorized creates ListOfClusterOperatorsUnauthorized with default headers values
func NewListOfClusterOperatorsUnauthorized() *ListOfClusterOperatorsUnauthorized {

	return &ListOfClusterOperatorsUnauthorized{}
}

// WithPayload adds the payload to the list of cluster operators unauthorized response
func (o *ListOfClusterOperatorsUnauthorized) WithPayload(payload *models.InfraError) *ListOfClusterOperatorsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list of cluster operators unauthorized response
func (o *ListOfClusterOperatorsUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListOfClusterOperatorsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListOfClusterOperatorsForbiddenCode is the HTTP code returned for type ListOfClusterOperatorsForbidden
const ListOfClusterOperatorsForbiddenCode int = 403

/*ListOfClusterOperatorsForbidden Forbidden.

swagger:response listOfClusterOperatorsForbidden
*/
type ListOfClusterOperatorsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewListOfClusterOperatorsForbidden creates ListOfClusterOperatorsForbidden with default headers values
func NewListOfClusterOperatorsForbidden() *ListOfClusterOperatorsForbidden {

	return &ListOfClusterOperatorsForbidden{}
}

// WithPayload adds the payload to the list of cluster operators forbidden response
func (o *ListOfClusterOperatorsForbidden) WithPayload(payload *models.InfraError) *ListOfClusterOperatorsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list of cluster operators forbidden response
func (o *ListOfClusterOperatorsForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListOfClusterOperatorsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListOfClusterOperatorsNotFoundCode is the HTTP code returned for type ListOfClusterOperatorsNotFound
const ListOfClusterOperatorsNotFoundCode int = 404

/*ListOfClusterOperatorsNotFound Error.

swagger:response listOfClusterOperatorsNotFound
*/
type ListOfClusterOperatorsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListOfClusterOperatorsNotFound creates ListOfClusterOperatorsNotFound with default headers values
func NewListOfClusterOperatorsNotFound() *ListOfClusterOperatorsNotFound {

	return &ListOfClusterOperatorsNotFound{}
}

// WithPayload adds the payload to the list of cluster operators not found response
func (o *ListOfClusterOperatorsNotFound) WithPayload(payload *models.Error) *ListOfClusterOperatorsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list of cluster operators not found response
func (o *ListOfClusterOperatorsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListOfClusterOperatorsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListOfClusterOperatorsMethodNotAllowedCode is the HTTP code returned for type ListOfClusterOperatorsMethodNotAllowed
const ListOfClusterOperatorsMethodNotAllowedCode int = 405

/*ListOfClusterOperatorsMethodNotAllowed Method Not Allowed.

swagger:response listOfClusterOperatorsMethodNotAllowed
*/
type ListOfClusterOperatorsMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListOfClusterOperatorsMethodNotAllowed creates ListOfClusterOperatorsMethodNotAllowed with default headers values
func NewListOfClusterOperatorsMethodNotAllowed() *ListOfClusterOperatorsMethodNotAllowed {

	return &ListOfClusterOperatorsMethodNotAllowed{}
}

// WithPayload adds the payload to the list of cluster operators method not allowed response
func (o *ListOfClusterOperatorsMethodNotAllowed) WithPayload(payload *models.Error) *ListOfClusterOperatorsMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list of cluster operators method not allowed response
func (o *ListOfClusterOperatorsMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListOfClusterOperatorsMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListOfClusterOperatorsInternalServerErrorCode is the HTTP code returned for type ListOfClusterOperatorsInternalServerError
const ListOfClusterOperatorsInternalServerErrorCode int = 500

/*ListOfClusterOperatorsInternalServerError Error.

swagger:response listOfClusterOperatorsInternalServerError
*/
type ListOfClusterOperatorsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListOfClusterOperatorsInternalServerError creates ListOfClusterOperatorsInternalServerError with default headers values
func NewListOfClusterOperatorsInternalServerError() *ListOfClusterOperatorsInternalServerError {

	return &ListOfClusterOperatorsInternalServerError{}
}

// WithPayload adds the payload to the list of cluster operators internal server error response
func (o *ListOfClusterOperatorsInternalServerError) WithPayload(payload *models.Error) *ListOfClusterOperatorsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list of cluster operators internal server error response
func (o *ListOfClusterOperatorsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListOfClusterOperatorsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
