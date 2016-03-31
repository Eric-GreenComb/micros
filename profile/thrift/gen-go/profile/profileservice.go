// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package profile

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type ProfileService interface {
	// Parameters:
	//  - ProfileID
	GetProfile(profile_id string) (r string, err error)
	// Parameters:
	//  - ProfileSearchCondition
	SearchProfiles(profile_search_condition *ProfileSearchCondition) (r string, err error)
}

type ProfileServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewProfileServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ProfileServiceClient {
	return &ProfileServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewProfileServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ProfileServiceClient {
	return &ProfileServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - ProfileID
func (p *ProfileServiceClient) GetProfile(profile_id string) (r string, err error) {
	if err = p.sendGetProfile(profile_id); err != nil {
		return
	}
	return p.recvGetProfile()
}

func (p *ProfileServiceClient) sendGetProfile(profile_id string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("GetProfile", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ProfileServiceGetProfileArgs{
		ProfileID: profile_id,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ProfileServiceClient) recvGetProfile() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "GetProfile" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "GetProfile failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "GetProfile failed: out of sequence response")
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
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "GetProfile failed: invalid message type")
		return
	}
	result := ProfileServiceGetProfileResult{}
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
//  - ProfileSearchCondition
func (p *ProfileServiceClient) SearchProfiles(profile_search_condition *ProfileSearchCondition) (r string, err error) {
	if err = p.sendSearchProfiles(profile_search_condition); err != nil {
		return
	}
	return p.recvSearchProfiles()
}

func (p *ProfileServiceClient) sendSearchProfiles(profile_search_condition *ProfileSearchCondition) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("SearchProfiles", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ProfileServiceSearchProfilesArgs{
		ProfileSearchCondition: profile_search_condition,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ProfileServiceClient) recvSearchProfiles() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "SearchProfiles" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "SearchProfiles failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "SearchProfiles failed: out of sequence response")
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
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "SearchProfiles failed: invalid message type")
		return
	}
	result := ProfileServiceSearchProfilesResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type ProfileServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      ProfileService
}

func (p *ProfileServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ProfileServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ProfileServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewProfileServiceProcessor(handler ProfileService) *ProfileServiceProcessor {

	self4 := &ProfileServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self4.processorMap["GetProfile"] = &profileServiceProcessorGetProfile{handler: handler}
	self4.processorMap["SearchProfiles"] = &profileServiceProcessorSearchProfiles{handler: handler}
	return self4
}

func (p *ProfileServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x5.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x5

}

type profileServiceProcessorGetProfile struct {
	handler ProfileService
}

func (p *profileServiceProcessorGetProfile) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ProfileServiceGetProfileArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetProfile", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ProfileServiceGetProfileResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.GetProfile(args.ProfileID); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetProfile: "+err2.Error())
		oprot.WriteMessageBegin("GetProfile", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("GetProfile", thrift.REPLY, seqId); err2 != nil {
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

type profileServiceProcessorSearchProfiles struct {
	handler ProfileService
}

func (p *profileServiceProcessorSearchProfiles) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ProfileServiceSearchProfilesArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("SearchProfiles", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ProfileServiceSearchProfilesResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.SearchProfiles(args.ProfileSearchCondition); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing SearchProfiles: "+err2.Error())
		oprot.WriteMessageBegin("SearchProfiles", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("SearchProfiles", thrift.REPLY, seqId); err2 != nil {
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

// Attributes:
//  - ProfileID
type ProfileServiceGetProfileArgs struct {
	ProfileID string `thrift:"profile_id,1" json:"profile_id"`
}

func NewProfileServiceGetProfileArgs() *ProfileServiceGetProfileArgs {
	return &ProfileServiceGetProfileArgs{}
}

func (p *ProfileServiceGetProfileArgs) GetProfileID() string {
	return p.ProfileID
}
func (p *ProfileServiceGetProfileArgs) Read(iprot thrift.TProtocol) error {
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

func (p *ProfileServiceGetProfileArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ProfileID = v
	}
	return nil
}

func (p *ProfileServiceGetProfileArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetProfile_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
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

func (p *ProfileServiceGetProfileArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("profile_id", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:profile_id: ", p), err)
	}
	if err := oprot.WriteString(string(p.ProfileID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.profile_id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:profile_id: ", p), err)
	}
	return err
}

func (p *ProfileServiceGetProfileArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ProfileServiceGetProfileArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ProfileServiceGetProfileResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewProfileServiceGetProfileResult() *ProfileServiceGetProfileResult {
	return &ProfileServiceGetProfileResult{}
}

var ProfileServiceGetProfileResult_Success_DEFAULT string

func (p *ProfileServiceGetProfileResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return ProfileServiceGetProfileResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *ProfileServiceGetProfileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProfileServiceGetProfileResult) Read(iprot thrift.TProtocol) error {
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

func (p *ProfileServiceGetProfileResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *ProfileServiceGetProfileResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetProfile_result"); err != nil {
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

func (p *ProfileServiceGetProfileResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *ProfileServiceGetProfileResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ProfileServiceGetProfileResult(%+v)", *p)
}

// Attributes:
//  - ProfileSearchCondition
type ProfileServiceSearchProfilesArgs struct {
	ProfileSearchCondition *ProfileSearchCondition `thrift:"profile_search_condition,1" json:"profile_search_condition"`
}

func NewProfileServiceSearchProfilesArgs() *ProfileServiceSearchProfilesArgs {
	return &ProfileServiceSearchProfilesArgs{}
}

var ProfileServiceSearchProfilesArgs_ProfileSearchCondition_DEFAULT *ProfileSearchCondition

func (p *ProfileServiceSearchProfilesArgs) GetProfileSearchCondition() *ProfileSearchCondition {
	if !p.IsSetProfileSearchCondition() {
		return ProfileServiceSearchProfilesArgs_ProfileSearchCondition_DEFAULT
	}
	return p.ProfileSearchCondition
}
func (p *ProfileServiceSearchProfilesArgs) IsSetProfileSearchCondition() bool {
	return p.ProfileSearchCondition != nil
}

func (p *ProfileServiceSearchProfilesArgs) Read(iprot thrift.TProtocol) error {
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

func (p *ProfileServiceSearchProfilesArgs) readField1(iprot thrift.TProtocol) error {
	p.ProfileSearchCondition = &ProfileSearchCondition{}
	if err := p.ProfileSearchCondition.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.ProfileSearchCondition), err)
	}
	return nil
}

func (p *ProfileServiceSearchProfilesArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SearchProfiles_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
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

func (p *ProfileServiceSearchProfilesArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("profile_search_condition", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:profile_search_condition: ", p), err)
	}
	if err := p.ProfileSearchCondition.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.ProfileSearchCondition), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:profile_search_condition: ", p), err)
	}
	return err
}

func (p *ProfileServiceSearchProfilesArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ProfileServiceSearchProfilesArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ProfileServiceSearchProfilesResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewProfileServiceSearchProfilesResult() *ProfileServiceSearchProfilesResult {
	return &ProfileServiceSearchProfilesResult{}
}

var ProfileServiceSearchProfilesResult_Success_DEFAULT string

func (p *ProfileServiceSearchProfilesResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return ProfileServiceSearchProfilesResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *ProfileServiceSearchProfilesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProfileServiceSearchProfilesResult) Read(iprot thrift.TProtocol) error {
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

func (p *ProfileServiceSearchProfilesResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *ProfileServiceSearchProfilesResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SearchProfiles_result"); err != nil {
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

func (p *ProfileServiceSearchProfilesResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *ProfileServiceSearchProfilesResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ProfileServiceSearchProfilesResult(%+v)", *p)
}
