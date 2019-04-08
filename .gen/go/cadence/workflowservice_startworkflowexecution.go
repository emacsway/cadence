// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by thriftrw v1.18.0. DO NOT EDIT.
// @generated

package cadence

import (
	errors "errors"
	fmt "fmt"
	shared "github.com/uber/cadence/.gen/go/shared"
	multierr "go.uber.org/multierr"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

// WorkflowService_StartWorkflowExecution_Args represents the arguments for the WorkflowService.StartWorkflowExecution function.
//
// The arguments for StartWorkflowExecution are sent and received over the wire as this struct.
type WorkflowService_StartWorkflowExecution_Args struct {
	StartRequest *shared.StartWorkflowExecutionRequest `json:"startRequest,omitempty"`
}

// ToWire translates a WorkflowService_StartWorkflowExecution_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_StartWorkflowExecution_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.StartRequest != nil {
		w, err = v.StartRequest.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _StartWorkflowExecutionRequest_Read(w wire.Value) (*shared.StartWorkflowExecutionRequest, error) {
	var v shared.StartWorkflowExecutionRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a WorkflowService_StartWorkflowExecution_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_StartWorkflowExecution_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_StartWorkflowExecution_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_StartWorkflowExecution_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.StartRequest, err = _StartWorkflowExecutionRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_StartWorkflowExecution_Args
// struct.
func (v *WorkflowService_StartWorkflowExecution_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.StartRequest != nil {
		fields[i] = fmt.Sprintf("StartRequest: %v", v.StartRequest)
		i++
	}

	return fmt.Sprintf("WorkflowService_StartWorkflowExecution_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_StartWorkflowExecution_Args match the
// provided WorkflowService_StartWorkflowExecution_Args.
//
// This function performs a deep comparison.
func (v *WorkflowService_StartWorkflowExecution_Args) Equals(rhs *WorkflowService_StartWorkflowExecution_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.StartRequest == nil && rhs.StartRequest == nil) || (v.StartRequest != nil && rhs.StartRequest != nil && v.StartRequest.Equals(rhs.StartRequest))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of WorkflowService_StartWorkflowExecution_Args.
func (v *WorkflowService_StartWorkflowExecution_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.StartRequest != nil {
		err = multierr.Append(err, enc.AddObject("startRequest", v.StartRequest))
	}
	return err
}

// GetStartRequest returns the value of StartRequest if it is set or its
// zero value if it is unset.
func (v *WorkflowService_StartWorkflowExecution_Args) GetStartRequest() (o *shared.StartWorkflowExecutionRequest) {
	if v != nil && v.StartRequest != nil {
		return v.StartRequest
	}

	return
}

// IsSetStartRequest returns true if StartRequest is not nil.
func (v *WorkflowService_StartWorkflowExecution_Args) IsSetStartRequest() bool {
	return v != nil && v.StartRequest != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "StartWorkflowExecution" for this struct.
func (v *WorkflowService_StartWorkflowExecution_Args) MethodName() string {
	return "StartWorkflowExecution"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *WorkflowService_StartWorkflowExecution_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// WorkflowService_StartWorkflowExecution_Helper provides functions that aid in handling the
// parameters and return values of the WorkflowService.StartWorkflowExecution
// function.
var WorkflowService_StartWorkflowExecution_Helper = struct {
	// Args accepts the parameters of StartWorkflowExecution in-order and returns
	// the arguments struct for the function.
	Args func(
		startRequest *shared.StartWorkflowExecutionRequest,
	) *WorkflowService_StartWorkflowExecution_Args

	// IsException returns true if the given error can be thrown
	// by StartWorkflowExecution.
	//
	// An error can be thrown by StartWorkflowExecution only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for StartWorkflowExecution
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// StartWorkflowExecution into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by StartWorkflowExecution
	//
	//   value, err := StartWorkflowExecution(args)
	//   result, err := WorkflowService_StartWorkflowExecution_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from StartWorkflowExecution: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*shared.StartWorkflowExecutionResponse, error) (*WorkflowService_StartWorkflowExecution_Result, error)

	// UnwrapResponse takes the result struct for StartWorkflowExecution
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if StartWorkflowExecution threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := WorkflowService_StartWorkflowExecution_Helper.UnwrapResponse(result)
	UnwrapResponse func(*WorkflowService_StartWorkflowExecution_Result) (*shared.StartWorkflowExecutionResponse, error)
}{}

func init() {
	WorkflowService_StartWorkflowExecution_Helper.Args = func(
		startRequest *shared.StartWorkflowExecutionRequest,
	) *WorkflowService_StartWorkflowExecution_Args {
		return &WorkflowService_StartWorkflowExecution_Args{
			StartRequest: startRequest,
		}
	}

	WorkflowService_StartWorkflowExecution_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.BadRequestError:
			return true
		case *shared.InternalServiceError:
			return true
		case *shared.WorkflowExecutionAlreadyStartedError:
			return true
		case *shared.ServiceBusyError:
			return true
		case *shared.DomainNotActiveError:
			return true
		case *shared.LimitExceededError:
			return true
		case *shared.EntityNotExistsError:
			return true
		default:
			return false
		}
	}

	WorkflowService_StartWorkflowExecution_Helper.WrapResponse = func(success *shared.StartWorkflowExecutionResponse, err error) (*WorkflowService_StartWorkflowExecution_Result, error) {
		if err == nil {
			return &WorkflowService_StartWorkflowExecution_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_StartWorkflowExecution_Result.BadRequestError")
			}
			return &WorkflowService_StartWorkflowExecution_Result{BadRequestError: e}, nil
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_StartWorkflowExecution_Result.InternalServiceError")
			}
			return &WorkflowService_StartWorkflowExecution_Result{InternalServiceError: e}, nil
		case *shared.WorkflowExecutionAlreadyStartedError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_StartWorkflowExecution_Result.SessionAlreadyExistError")
			}
			return &WorkflowService_StartWorkflowExecution_Result{SessionAlreadyExistError: e}, nil
		case *shared.ServiceBusyError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_StartWorkflowExecution_Result.ServiceBusyError")
			}
			return &WorkflowService_StartWorkflowExecution_Result{ServiceBusyError: e}, nil
		case *shared.DomainNotActiveError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_StartWorkflowExecution_Result.DomainNotActiveError")
			}
			return &WorkflowService_StartWorkflowExecution_Result{DomainNotActiveError: e}, nil
		case *shared.LimitExceededError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_StartWorkflowExecution_Result.LimitExceededError")
			}
			return &WorkflowService_StartWorkflowExecution_Result{LimitExceededError: e}, nil
		case *shared.EntityNotExistsError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_StartWorkflowExecution_Result.EntityNotExistError")
			}
			return &WorkflowService_StartWorkflowExecution_Result{EntityNotExistError: e}, nil
		}

		return nil, err
	}
	WorkflowService_StartWorkflowExecution_Helper.UnwrapResponse = func(result *WorkflowService_StartWorkflowExecution_Result) (success *shared.StartWorkflowExecutionResponse, err error) {
		if result.BadRequestError != nil {
			err = result.BadRequestError
			return
		}
		if result.InternalServiceError != nil {
			err = result.InternalServiceError
			return
		}
		if result.SessionAlreadyExistError != nil {
			err = result.SessionAlreadyExistError
			return
		}
		if result.ServiceBusyError != nil {
			err = result.ServiceBusyError
			return
		}
		if result.DomainNotActiveError != nil {
			err = result.DomainNotActiveError
			return
		}
		if result.LimitExceededError != nil {
			err = result.LimitExceededError
			return
		}
		if result.EntityNotExistError != nil {
			err = result.EntityNotExistError
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// WorkflowService_StartWorkflowExecution_Result represents the result of a WorkflowService.StartWorkflowExecution function call.
//
// The result of a StartWorkflowExecution execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type WorkflowService_StartWorkflowExecution_Result struct {
	// Value returned by StartWorkflowExecution after a successful execution.
	Success                  *shared.StartWorkflowExecutionResponse       `json:"success,omitempty"`
	BadRequestError          *shared.BadRequestError                      `json:"badRequestError,omitempty"`
	InternalServiceError     *shared.InternalServiceError                 `json:"internalServiceError,omitempty"`
	SessionAlreadyExistError *shared.WorkflowExecutionAlreadyStartedError `json:"sessionAlreadyExistError,omitempty"`
	ServiceBusyError         *shared.ServiceBusyError                     `json:"serviceBusyError,omitempty"`
	DomainNotActiveError     *shared.DomainNotActiveError                 `json:"domainNotActiveError,omitempty"`
	LimitExceededError       *shared.LimitExceededError                   `json:"limitExceededError,omitempty"`
	EntityNotExistError      *shared.EntityNotExistsError                 `json:"entityNotExistError,omitempty"`
}

// ToWire translates a WorkflowService_StartWorkflowExecution_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_StartWorkflowExecution_Result) ToWire() (wire.Value, error) {
	var (
		fields [8]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.BadRequestError != nil {
		w, err = v.BadRequestError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.InternalServiceError != nil {
		w, err = v.InternalServiceError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.SessionAlreadyExistError != nil {
		w, err = v.SessionAlreadyExistError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.ServiceBusyError != nil {
		w, err = v.ServiceBusyError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	if v.DomainNotActiveError != nil {
		w, err = v.DomainNotActiveError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 5, Value: w}
		i++
	}
	if v.LimitExceededError != nil {
		w, err = v.LimitExceededError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 6, Value: w}
		i++
	}
	if v.EntityNotExistError != nil {
		w, err = v.EntityNotExistError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 7, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("WorkflowService_StartWorkflowExecution_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a WorkflowService_StartWorkflowExecution_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_StartWorkflowExecution_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_StartWorkflowExecution_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_StartWorkflowExecution_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _StartWorkflowExecutionResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BadRequestError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.InternalServiceError, err = _InternalServiceError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TStruct {
				v.SessionAlreadyExistError, err = _WorkflowExecutionAlreadyStartedError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 4:
			if field.Value.Type() == wire.TStruct {
				v.ServiceBusyError, err = _ServiceBusyError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 5:
			if field.Value.Type() == wire.TStruct {
				v.DomainNotActiveError, err = _DomainNotActiveError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 6:
			if field.Value.Type() == wire.TStruct {
				v.LimitExceededError, err = _LimitExceededError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 7:
			if field.Value.Type() == wire.TStruct {
				v.EntityNotExistError, err = _EntityNotExistsError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.BadRequestError != nil {
		count++
	}
	if v.InternalServiceError != nil {
		count++
	}
	if v.SessionAlreadyExistError != nil {
		count++
	}
	if v.ServiceBusyError != nil {
		count++
	}
	if v.DomainNotActiveError != nil {
		count++
	}
	if v.LimitExceededError != nil {
		count++
	}
	if v.EntityNotExistError != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("WorkflowService_StartWorkflowExecution_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_StartWorkflowExecution_Result
// struct.
func (v *WorkflowService_StartWorkflowExecution_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [8]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.BadRequestError != nil {
		fields[i] = fmt.Sprintf("BadRequestError: %v", v.BadRequestError)
		i++
	}
	if v.InternalServiceError != nil {
		fields[i] = fmt.Sprintf("InternalServiceError: %v", v.InternalServiceError)
		i++
	}
	if v.SessionAlreadyExistError != nil {
		fields[i] = fmt.Sprintf("SessionAlreadyExistError: %v", v.SessionAlreadyExistError)
		i++
	}
	if v.ServiceBusyError != nil {
		fields[i] = fmt.Sprintf("ServiceBusyError: %v", v.ServiceBusyError)
		i++
	}
	if v.DomainNotActiveError != nil {
		fields[i] = fmt.Sprintf("DomainNotActiveError: %v", v.DomainNotActiveError)
		i++
	}
	if v.LimitExceededError != nil {
		fields[i] = fmt.Sprintf("LimitExceededError: %v", v.LimitExceededError)
		i++
	}
	if v.EntityNotExistError != nil {
		fields[i] = fmt.Sprintf("EntityNotExistError: %v", v.EntityNotExistError)
		i++
	}

	return fmt.Sprintf("WorkflowService_StartWorkflowExecution_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_StartWorkflowExecution_Result match the
// provided WorkflowService_StartWorkflowExecution_Result.
//
// This function performs a deep comparison.
func (v *WorkflowService_StartWorkflowExecution_Result) Equals(rhs *WorkflowService_StartWorkflowExecution_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.BadRequestError == nil && rhs.BadRequestError == nil) || (v.BadRequestError != nil && rhs.BadRequestError != nil && v.BadRequestError.Equals(rhs.BadRequestError))) {
		return false
	}
	if !((v.InternalServiceError == nil && rhs.InternalServiceError == nil) || (v.InternalServiceError != nil && rhs.InternalServiceError != nil && v.InternalServiceError.Equals(rhs.InternalServiceError))) {
		return false
	}
	if !((v.SessionAlreadyExistError == nil && rhs.SessionAlreadyExistError == nil) || (v.SessionAlreadyExistError != nil && rhs.SessionAlreadyExistError != nil && v.SessionAlreadyExistError.Equals(rhs.SessionAlreadyExistError))) {
		return false
	}
	if !((v.ServiceBusyError == nil && rhs.ServiceBusyError == nil) || (v.ServiceBusyError != nil && rhs.ServiceBusyError != nil && v.ServiceBusyError.Equals(rhs.ServiceBusyError))) {
		return false
	}
	if !((v.DomainNotActiveError == nil && rhs.DomainNotActiveError == nil) || (v.DomainNotActiveError != nil && rhs.DomainNotActiveError != nil && v.DomainNotActiveError.Equals(rhs.DomainNotActiveError))) {
		return false
	}
	if !((v.LimitExceededError == nil && rhs.LimitExceededError == nil) || (v.LimitExceededError != nil && rhs.LimitExceededError != nil && v.LimitExceededError.Equals(rhs.LimitExceededError))) {
		return false
	}
	if !((v.EntityNotExistError == nil && rhs.EntityNotExistError == nil) || (v.EntityNotExistError != nil && rhs.EntityNotExistError != nil && v.EntityNotExistError.Equals(rhs.EntityNotExistError))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of WorkflowService_StartWorkflowExecution_Result.
func (v *WorkflowService_StartWorkflowExecution_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	if v.BadRequestError != nil {
		err = multierr.Append(err, enc.AddObject("badRequestError", v.BadRequestError))
	}
	if v.InternalServiceError != nil {
		err = multierr.Append(err, enc.AddObject("internalServiceError", v.InternalServiceError))
	}
	if v.SessionAlreadyExistError != nil {
		err = multierr.Append(err, enc.AddObject("sessionAlreadyExistError", v.SessionAlreadyExistError))
	}
	if v.ServiceBusyError != nil {
		err = multierr.Append(err, enc.AddObject("serviceBusyError", v.ServiceBusyError))
	}
	if v.DomainNotActiveError != nil {
		err = multierr.Append(err, enc.AddObject("domainNotActiveError", v.DomainNotActiveError))
	}
	if v.LimitExceededError != nil {
		err = multierr.Append(err, enc.AddObject("limitExceededError", v.LimitExceededError))
	}
	if v.EntityNotExistError != nil {
		err = multierr.Append(err, enc.AddObject("entityNotExistError", v.EntityNotExistError))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *WorkflowService_StartWorkflowExecution_Result) GetSuccess() (o *shared.StartWorkflowExecutionResponse) {
	if v != nil && v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *WorkflowService_StartWorkflowExecution_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// GetBadRequestError returns the value of BadRequestError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_StartWorkflowExecution_Result) GetBadRequestError() (o *shared.BadRequestError) {
	if v != nil && v.BadRequestError != nil {
		return v.BadRequestError
	}

	return
}

// IsSetBadRequestError returns true if BadRequestError is not nil.
func (v *WorkflowService_StartWorkflowExecution_Result) IsSetBadRequestError() bool {
	return v != nil && v.BadRequestError != nil
}

// GetInternalServiceError returns the value of InternalServiceError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_StartWorkflowExecution_Result) GetInternalServiceError() (o *shared.InternalServiceError) {
	if v != nil && v.InternalServiceError != nil {
		return v.InternalServiceError
	}

	return
}

// IsSetInternalServiceError returns true if InternalServiceError is not nil.
func (v *WorkflowService_StartWorkflowExecution_Result) IsSetInternalServiceError() bool {
	return v != nil && v.InternalServiceError != nil
}

// GetSessionAlreadyExistError returns the value of SessionAlreadyExistError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_StartWorkflowExecution_Result) GetSessionAlreadyExistError() (o *shared.WorkflowExecutionAlreadyStartedError) {
	if v != nil && v.SessionAlreadyExistError != nil {
		return v.SessionAlreadyExistError
	}

	return
}

// IsSetSessionAlreadyExistError returns true if SessionAlreadyExistError is not nil.
func (v *WorkflowService_StartWorkflowExecution_Result) IsSetSessionAlreadyExistError() bool {
	return v != nil && v.SessionAlreadyExistError != nil
}

// GetServiceBusyError returns the value of ServiceBusyError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_StartWorkflowExecution_Result) GetServiceBusyError() (o *shared.ServiceBusyError) {
	if v != nil && v.ServiceBusyError != nil {
		return v.ServiceBusyError
	}

	return
}

// IsSetServiceBusyError returns true if ServiceBusyError is not nil.
func (v *WorkflowService_StartWorkflowExecution_Result) IsSetServiceBusyError() bool {
	return v != nil && v.ServiceBusyError != nil
}

// GetDomainNotActiveError returns the value of DomainNotActiveError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_StartWorkflowExecution_Result) GetDomainNotActiveError() (o *shared.DomainNotActiveError) {
	if v != nil && v.DomainNotActiveError != nil {
		return v.DomainNotActiveError
	}

	return
}

// IsSetDomainNotActiveError returns true if DomainNotActiveError is not nil.
func (v *WorkflowService_StartWorkflowExecution_Result) IsSetDomainNotActiveError() bool {
	return v != nil && v.DomainNotActiveError != nil
}

// GetLimitExceededError returns the value of LimitExceededError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_StartWorkflowExecution_Result) GetLimitExceededError() (o *shared.LimitExceededError) {
	if v != nil && v.LimitExceededError != nil {
		return v.LimitExceededError
	}

	return
}

// IsSetLimitExceededError returns true if LimitExceededError is not nil.
func (v *WorkflowService_StartWorkflowExecution_Result) IsSetLimitExceededError() bool {
	return v != nil && v.LimitExceededError != nil
}

// GetEntityNotExistError returns the value of EntityNotExistError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_StartWorkflowExecution_Result) GetEntityNotExistError() (o *shared.EntityNotExistsError) {
	if v != nil && v.EntityNotExistError != nil {
		return v.EntityNotExistError
	}

	return
}

// IsSetEntityNotExistError returns true if EntityNotExistError is not nil.
func (v *WorkflowService_StartWorkflowExecution_Result) IsSetEntityNotExistError() bool {
	return v != nil && v.EntityNotExistError != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "StartWorkflowExecution" for this struct.
func (v *WorkflowService_StartWorkflowExecution_Result) MethodName() string {
	return "StartWorkflowExecution"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *WorkflowService_StartWorkflowExecution_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}