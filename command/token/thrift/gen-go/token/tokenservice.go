// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package token

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type TokenService interface {
	Ping() (r string, err error)
	// Parameters:
	//  - Key
	//  - Ttype
	CreateToken(key string, ttype int64) (r string, err error)
	// Parameters:
	//  - Key
	//  - Ttype
	DeleteToken(key string, ttype int64) (r bool, err error)
}

type TokenServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewTokenServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TokenServiceClient {
	return &TokenServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewTokenServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TokenServiceClient {
	return &TokenServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

func (p *TokenServiceClient) Ping() (r string, err error) {
	if err = p.sendPing(); err != nil {
		return
	}
	return p.recvPing()
}

func (p *TokenServiceClient) sendPing() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("Ping", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := TokenServicePingArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TokenServiceClient) recvPing() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "Ping" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "Ping failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Ping failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "Ping failed: invalid message type")
		return
	}
	result := TokenServicePingResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Key
//  - Ttype
func (p *TokenServiceClient) CreateToken(key string, ttype int64) (r string, err error) {
	if err = p.sendCreateToken(key, ttype); err != nil {
		return
	}
	return p.recvCreateToken()
}

func (p *TokenServiceClient) sendCreateToken(key string, ttype int64) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("CreateToken", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := TokenServiceCreateTokenArgs{
		Key:   key,
		Ttype: ttype,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TokenServiceClient) recvCreateToken() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "CreateToken" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "CreateToken failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "CreateToken failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "CreateToken failed: invalid message type")
		return
	}
	result := TokenServiceCreateTokenResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Key
//  - Ttype
func (p *TokenServiceClient) DeleteToken(key string, ttype int64) (r bool, err error) {
	if err = p.sendDeleteToken(key, ttype); err != nil {
		return
	}
	return p.recvDeleteToken()
}

func (p *TokenServiceClient) sendDeleteToken(key string, ttype int64) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("DeleteToken", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := TokenServiceDeleteTokenArgs{
		Key:   key,
		Ttype: ttype,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TokenServiceClient) recvDeleteToken() (value bool, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "DeleteToken" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "DeleteToken failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "DeleteToken failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error5 error
		error5, err = error4.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error5
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "DeleteToken failed: invalid message type")
		return
	}
	result := TokenServiceDeleteTokenResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type TokenServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      TokenService
}

func (p *TokenServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *TokenServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *TokenServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewTokenServiceProcessor(handler TokenService) *TokenServiceProcessor {

	self6 := &TokenServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self6.processorMap["Ping"] = &tokenServiceProcessorPing{handler: handler}
	self6.processorMap["CreateToken"] = &tokenServiceProcessorCreateToken{handler: handler}
	self6.processorMap["DeleteToken"] = &tokenServiceProcessorDeleteToken{handler: handler}
	return self6
}

func (p *TokenServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x7.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x7

}

type tokenServiceProcessorPing struct {
	handler TokenService
}

func (p *tokenServiceProcessorPing) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TokenServicePingArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Ping", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TokenServicePingResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.Ping(); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Ping: "+err2.Error())
		oprot.WriteMessageBegin("Ping", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("Ping", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tokenServiceProcessorCreateToken struct {
	handler TokenService
}

func (p *tokenServiceProcessorCreateToken) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TokenServiceCreateTokenArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("CreateToken", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TokenServiceCreateTokenResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.CreateToken(args.Key, args.Ttype); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing CreateToken: "+err2.Error())
		oprot.WriteMessageBegin("CreateToken", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("CreateToken", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tokenServiceProcessorDeleteToken struct {
	handler TokenService
}

func (p *tokenServiceProcessorDeleteToken) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TokenServiceDeleteTokenArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("DeleteToken", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TokenServiceDeleteTokenResult{}
	var retval bool
	var err2 error
	if retval, err2 = p.handler.DeleteToken(args.Key, args.Ttype); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing DeleteToken: "+err2.Error())
		oprot.WriteMessageBegin("DeleteToken", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("DeleteToken", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type TokenServicePingArgs struct {
}

func NewTokenServicePingArgs() *TokenServicePingArgs {
	return &TokenServicePingArgs{}
}

func (p *TokenServicePingArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TokenServicePingArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Ping_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TokenServicePingArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TokenServicePingArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TokenServicePingResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewTokenServicePingResult() *TokenServicePingResult {
	return &TokenServicePingResult{}
}

var TokenServicePingResult_Success_DEFAULT string

func (p *TokenServicePingResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return TokenServicePingResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *TokenServicePingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TokenServicePingResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TokenServicePingResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *TokenServicePingResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Ping_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TokenServicePingResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TokenServicePingResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TokenServicePingResult(%+v)", *p)
}

// Attributes:
//  - Key
//  - Ttype
type TokenServiceCreateTokenArgs struct {
	Key   string `thrift:"key,1" json:"key"`
	Ttype int64  `thrift:"ttype,2" json:"ttype"`
}

func NewTokenServiceCreateTokenArgs() *TokenServiceCreateTokenArgs {
	return &TokenServiceCreateTokenArgs{}
}

func (p *TokenServiceCreateTokenArgs) GetKey() string {
	return p.Key
}

func (p *TokenServiceCreateTokenArgs) GetTtype() int64 {
	return p.Ttype
}
func (p *TokenServiceCreateTokenArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TokenServiceCreateTokenArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Key = v
	}
	return nil
}

func (p *TokenServiceCreateTokenArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Ttype = v
	}
	return nil
}

func (p *TokenServiceCreateTokenArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CreateToken_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TokenServiceCreateTokenArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err)
	}
	if err := oprot.WriteString(string(p.Key)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err)
	}
	return err
}

func (p *TokenServiceCreateTokenArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ttype", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:ttype: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Ttype)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ttype (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:ttype: ", p), err)
	}
	return err
}

func (p *TokenServiceCreateTokenArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TokenServiceCreateTokenArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TokenServiceCreateTokenResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewTokenServiceCreateTokenResult() *TokenServiceCreateTokenResult {
	return &TokenServiceCreateTokenResult{}
}

var TokenServiceCreateTokenResult_Success_DEFAULT string

func (p *TokenServiceCreateTokenResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return TokenServiceCreateTokenResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *TokenServiceCreateTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TokenServiceCreateTokenResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TokenServiceCreateTokenResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *TokenServiceCreateTokenResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CreateToken_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TokenServiceCreateTokenResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TokenServiceCreateTokenResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TokenServiceCreateTokenResult(%+v)", *p)
}

// Attributes:
//  - Key
//  - Ttype
type TokenServiceDeleteTokenArgs struct {
	Key   string `thrift:"key,1" json:"key"`
	Ttype int64  `thrift:"ttype,2" json:"ttype"`
}

func NewTokenServiceDeleteTokenArgs() *TokenServiceDeleteTokenArgs {
	return &TokenServiceDeleteTokenArgs{}
}

func (p *TokenServiceDeleteTokenArgs) GetKey() string {
	return p.Key
}

func (p *TokenServiceDeleteTokenArgs) GetTtype() int64 {
	return p.Ttype
}
func (p *TokenServiceDeleteTokenArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TokenServiceDeleteTokenArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Key = v
	}
	return nil
}

func (p *TokenServiceDeleteTokenArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Ttype = v
	}
	return nil
}

func (p *TokenServiceDeleteTokenArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("DeleteToken_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TokenServiceDeleteTokenArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err)
	}
	if err := oprot.WriteString(string(p.Key)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err)
	}
	return err
}

func (p *TokenServiceDeleteTokenArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ttype", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:ttype: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Ttype)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ttype (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:ttype: ", p), err)
	}
	return err
}

func (p *TokenServiceDeleteTokenArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TokenServiceDeleteTokenArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TokenServiceDeleteTokenResult struct {
	Success *bool `thrift:"success,0" json:"success,omitempty"`
}

func NewTokenServiceDeleteTokenResult() *TokenServiceDeleteTokenResult {
	return &TokenServiceDeleteTokenResult{}
}

var TokenServiceDeleteTokenResult_Success_DEFAULT bool

func (p *TokenServiceDeleteTokenResult) GetSuccess() bool {
	if !p.IsSetSuccess() {
		return TokenServiceDeleteTokenResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *TokenServiceDeleteTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TokenServiceDeleteTokenResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TokenServiceDeleteTokenResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *TokenServiceDeleteTokenResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("DeleteToken_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TokenServiceDeleteTokenResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TokenServiceDeleteTokenResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TokenServiceDeleteTokenResult(%+v)", *p)
}
