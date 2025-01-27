// Code generated by unionize.

package main

import (
	bytes "bytes"
	template1 "html/template"
	template0 "text/template"
)

type UnionSafeType uint8

const (
	UnionSafeType_Uint64      UnionSafeType = 0
	UnionSafeType_Int64       UnionSafeType = 1
	UnionSafeType_String      UnionSafeType = 2
	UnionSafeType_importTest1 UnionSafeType = 3
	UnionSafeType_importTest2 UnionSafeType = 4
	UnionSafeType_importTest3 UnionSafeType = 5
	UnionSafeType_importTest4 UnionSafeType = 6
	UnionSafeType_importTest5 UnionSafeType = 7
	UnionSafeType_importTest6 UnionSafeType = 8
	UnionSafeType_importTest7 UnionSafeType = 9
)

func (x UnionSafeType) String() string {
	switch x {
	case UnionSafeType_Uint64:
		return "Uint64"
	case UnionSafeType_Int64:
		return "Int64"
	case UnionSafeType_String:
		return "String"
	case UnionSafeType_importTest1:
		return "importTest1"
	case UnionSafeType_importTest2:
		return "importTest2"
	case UnionSafeType_importTest3:
		return "importTest3"
	case UnionSafeType_importTest4:
		return "importTest4"
	case UnionSafeType_importTest5:
		return "importTest5"
	case UnionSafeType_importTest6:
		return "importTest6"
	case UnionSafeType_importTest7:
		return "importTest7"
	default:
		panic("unreachable")
	}
}

func (u *UnionSafe) Type() UnionSafeType {
	return u.typ
}

type UnionSafe struct {
	typ          UnionSafeType
	_Uint64      uint64
	_Int64       int64
	_String      string
	_importTest1 []bytes.Buffer
	_importTest2 [1]bytes.Buffer
	_importTest3 struct{ bytes.Buffer }
	_importTest4 interface{ Buffer() bytes.Buffer }
	_importTest5 Dummy
	_importTest6 template0.Template
	_importTest7 template1.Template
}

func (u *UnionSafe) Uint64() uint64 {
	return u._Uint64
}
func (u *UnionSafe) SetUint64(v uint64) {
	u.typ = 0
	u._Uint64 = v
}
func (u *UnionSafe) Uint64Ptr() *uint64 {
	return &u._Uint64
}

func (u *UnionSafe) Int64() int64 {
	return u._Int64
}
func (u *UnionSafe) SetInt64(v int64) {
	u.typ = 1
	u._Int64 = v
}
func (u *UnionSafe) Int64Ptr() *int64 {
	return &u._Int64
}

func (u *UnionSafe) String() string {
	return u._String
}
func (u *UnionSafe) SetString(v string) {
	u.typ = 2
	u._String = v
}
func (u *UnionSafe) StringPtr() *string {
	return &u._String
}

func (u *UnionSafe) importTest1() []bytes.Buffer {
	return u._importTest1
}
func (u *UnionSafe) setImportTest1(v []bytes.Buffer) {
	u.typ = 3
	u._importTest1 = v
}
func (u *UnionSafe) importTest1Ptr() *[]bytes.Buffer {
	return &u._importTest1
}

func (u *UnionSafe) importTest2() [1]bytes.Buffer {
	return u._importTest2
}
func (u *UnionSafe) setImportTest2(v [1]bytes.Buffer) {
	u.typ = 4
	u._importTest2 = v
}
func (u *UnionSafe) importTest2Ptr() *[1]bytes.Buffer {
	return &u._importTest2
}

func (u *UnionSafe) importTest3() struct{ bytes.Buffer } {
	return u._importTest3
}
func (u *UnionSafe) setImportTest3(v struct{ bytes.Buffer }) {
	u.typ = 5
	u._importTest3 = v
}
func (u *UnionSafe) importTest3Ptr() *struct{ bytes.Buffer } {
	return &u._importTest3
}

func (u *UnionSafe) importTest4() interface{ Buffer() bytes.Buffer } {
	return u._importTest4
}
func (u *UnionSafe) setImportTest4(v interface{ Buffer() bytes.Buffer }) {
	u.typ = 6
	u._importTest4 = v
}
func (u *UnionSafe) importTest4Ptr() *interface{ Buffer() bytes.Buffer } {
	return &u._importTest4
}

func (u *UnionSafe) importTest5() Dummy {
	return u._importTest5
}
func (u *UnionSafe) setImportTest5(v Dummy) {
	u.typ = 7
	u._importTest5 = v
}
func (u *UnionSafe) importTest5Ptr() *Dummy {
	return &u._importTest5
}

func (u *UnionSafe) importTest6() template0.Template {
	return u._importTest6
}
func (u *UnionSafe) setImportTest6(v template0.Template) {
	u.typ = 8
	u._importTest6 = v
}
func (u *UnionSafe) importTest6Ptr() *template0.Template {
	return &u._importTest6
}

func (u *UnionSafe) importTest7() template1.Template {
	return u._importTest7
}
func (u *UnionSafe) setImportTest7(v template1.Template) {
	u.typ = 9
	u._importTest7 = v
}
func (u *UnionSafe) importTest7Ptr() *template1.Template {
	return &u._importTest7
}
