// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package category

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// Attributes:
//  - ID
//  - Name
//  - Subcategories
type Category struct {
	ID            string         `thrift:"id,1" json:"id"`
	Name          string         `thrift:"name,2" json:"name"`
	Subcategories []*SubCategory `thrift:"subcategories,3" json:"subcategories"`
}

func NewCategory() *Category {
	return &Category{}
}

func (p *Category) GetID() string {
	return p.ID
}

func (p *Category) GetName() string {
	return p.Name
}

func (p *Category) GetSubcategories() []*SubCategory {
	return p.Subcategories
}
func (p *Category) Read(iprot thrift.TProtocol) error {
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
		case 3:
			if err := p.readField3(iprot); err != nil {
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

func (p *Category) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *Category) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *Category) readField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*SubCategory, 0, size)
	p.Subcategories = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &SubCategory{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.Subcategories = append(p.Subcategories, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *Category) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Category"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
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

func (p *Category) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err)
	}
	if err := oprot.WriteString(string(p.ID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err)
	}
	return err
}

func (p *Category) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:name: ", p), err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.name (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:name: ", p), err)
	}
	return err
}

func (p *Category) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("subcategories", thrift.LIST, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:subcategories: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Subcategories)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Subcategories {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:subcategories: ", p), err)
	}
	return err
}

func (p *Category) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Category(%+v)", *p)
}

// Attributes:
//  - ID
//  - Name
type SubCategory struct {
	ID   string `thrift:"id,1" json:"id"`
	Name string `thrift:"name,2" json:"name"`
}

func NewSubCategory() *SubCategory {
	return &SubCategory{}
}

func (p *SubCategory) GetID() string {
	return p.ID
}

func (p *SubCategory) GetName() string {
	return p.Name
}
func (p *SubCategory) Read(iprot thrift.TProtocol) error {
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

func (p *SubCategory) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *SubCategory) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *SubCategory) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SubCategory"); err != nil {
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

func (p *SubCategory) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err)
	}
	if err := oprot.WriteString(string(p.ID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err)
	}
	return err
}

func (p *SubCategory) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:name: ", p), err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.name (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:name: ", p), err)
	}
	return err
}

func (p *SubCategory) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SubCategory(%+v)", *p)
}
