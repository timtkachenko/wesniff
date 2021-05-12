// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EventBase event base
//
// swagger:discriminator EventBase EventBase
type EventBase interface {
	runtime.Validatable

	// status
	Status() string
	SetStatus(string)

	// verification
	Verification() *Verification
	SetVerification(*Verification)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type eventBase struct {
	statusField string

	verificationField *Verification
}

// Status gets the status of this polymorphic type
func (m *eventBase) Status() string {
	return m.statusField
}

// SetStatus sets the status of this polymorphic type
func (m *eventBase) SetStatus(val string) {
	m.statusField = val
}

// Verification gets the verification of this polymorphic type
func (m *eventBase) Verification() *Verification {
	return m.verificationField
}

// SetVerification sets the verification of this polymorphic type
func (m *eventBase) SetVerification(val *Verification) {
	m.verificationField = val
}

// UnmarshalEventBaseSlice unmarshals polymorphic slices of EventBase
func UnmarshalEventBaseSlice(reader io.Reader, consumer runtime.Consumer) ([]EventBase, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []EventBase
	for _, element := range elements {
		obj, err := unmarshalEventBase(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalEventBase unmarshals polymorphic EventBase
func UnmarshalEventBase(reader io.Reader, consumer runtime.Consumer) (EventBase, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalEventBase(data, consumer)
}

func unmarshalEventBase(data []byte, consumer runtime.Consumer) (EventBase, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the EventBase property.
	var getType struct {
		EventBase string `json:"EventBase"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("EventBase", "body", getType.EventBase); err != nil {
		return nil, err
	}

	// The value of EventBase is used to determine which type to create and unmarshal the data into
	switch getType.EventBase {
	case "EventBase":
		var result eventBase
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "EventBaseDto":
		var result EventBaseDto
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid EventBase value: %q", getType.EventBase)
}

// Validate validates this event base
func (m *eventBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVerification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *eventBase) validateVerification(formats strfmt.Registry) error {

	if swag.IsZero(m.Verification()) { // not required
		return nil
	}

	if m.Verification() != nil {
		if err := m.Verification().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("verification")
			}
			return err
		}
	}

	return nil
}