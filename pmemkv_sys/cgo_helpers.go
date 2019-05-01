// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 01 May 2019 10:58:26 CDT.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package pmemkv_sys

/*
#cgo LDFLAGS: -lpmemkv -ltbb
#include "libpmemkv.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = unsafe.Pointer(p)
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - uintptr(h.Data))
	}
	return
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(h.Data), C.int(h.Len))
}

func (x KVAllCallback) PassRef() (ref *C.KVAllCallback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if kVAllCallback5910E1EFunc == nil {
		kVAllCallback5910E1EFunc = x
	}
	return (*C.KVAllCallback)(C.KVAllCallback_5910e1e), nil
}

func NewKVAllCallbackRef(ref unsafe.Pointer) *KVAllCallback {
	return (*KVAllCallback)(ref)
}

//export kVAllCallback5910E1E
func kVAllCallback5910E1E(ccontext unsafe.Pointer, ckeybytes C.int, ckey *C.char) {
	if kVAllCallback5910E1EFunc != nil {
		context5910e1e := (unsafe.Pointer)(unsafe.Pointer(ccontext))
		keybytes5910e1e := (int32)(ckeybytes)
		key5910e1e := packPCharString(ckey)
		kVAllCallback5910E1EFunc(context5910e1e, keybytes5910e1e, key5910e1e)
		return
	}
	panic("callback func has not been set (race?)")
}

var kVAllCallback5910E1EFunc KVAllCallback

func (x KVEachCallback) PassRef() (ref *C.KVEachCallback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if kVEachCallback770981B9Func == nil {
		kVEachCallback770981B9Func = x
	}
	return (*C.KVEachCallback)(C.KVEachCallback_770981b9), nil
}

func NewKVEachCallbackRef(ref unsafe.Pointer) *KVEachCallback {
	return (*KVEachCallback)(ref)
}

//export kVEachCallback770981B9
func kVEachCallback770981B9(ccontext unsafe.Pointer, ckeybytes C.int, ckey *C.char, cvaluebytes C.int, cvalue *C.char) {
	if kVEachCallback770981B9Func != nil {
		context770981b9 := (unsafe.Pointer)(unsafe.Pointer(ccontext))
		keybytes770981b9 := (int32)(ckeybytes)
		key770981b9 := packPCharString(ckey)
		valuebytes770981b9 := (int32)(cvaluebytes)
		value770981b9 := packPCharString(cvalue)
		kVEachCallback770981B9Func(context770981b9, keybytes770981b9, key770981b9, valuebytes770981b9, value770981b9)
		return
	}
	panic("callback func has not been set (race?)")
}

var kVEachCallback770981B9Func KVEachCallback

func (x KVGetCallback) PassRef() (ref *C.KVGetCallback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if kVGetCallback6290FFEFunc == nil {
		kVGetCallback6290FFEFunc = x
	}
	return (*C.KVGetCallback)(C.KVGetCallback_6290ffe), nil
}

func NewKVGetCallbackRef(ref unsafe.Pointer) *KVGetCallback {
	return (*KVGetCallback)(ref)
}

//export kVGetCallback6290FFE
func kVGetCallback6290FFE(ccontext unsafe.Pointer, cvaluebytes C.int, cvalue *C.char) {
	if kVGetCallback6290FFEFunc != nil {
		context6290ffe := (unsafe.Pointer)(unsafe.Pointer(ccontext))
		valuebytes6290ffe := (int32)(cvaluebytes)
		value6290ffe := packPCharString(cvalue)
		kVGetCallback6290FFEFunc(context6290ffe, valuebytes6290ffe, value6290ffe)
		return
	}
	panic("callback func has not been set (race?)")
}

var kVGetCallback6290FFEFunc KVGetCallback

func (x KVStartFailureCallback) PassRef() (ref *C.KVStartFailureCallback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if kVStartFailureCallback206EE003Func == nil {
		kVStartFailureCallback206EE003Func = x
	}
	return (*C.KVStartFailureCallback)(C.KVStartFailureCallback_206ee003), nil
}

func NewKVStartFailureCallbackRef(ref unsafe.Pointer) *KVStartFailureCallback {
	return (*KVStartFailureCallback)(ref)
}

//export kVStartFailureCallback206EE003
func kVStartFailureCallback206EE003(ccontext unsafe.Pointer, cengine *C.char, cconfig *C.char, cmsg *C.char) {
	if kVStartFailureCallback206EE003Func != nil {
		context206ee003 := (unsafe.Pointer)(unsafe.Pointer(ccontext))
		engine206ee003 := packPCharString(cengine)
		config206ee003 := packPCharString(cconfig)
		msg206ee003 := packPCharString(cmsg)
		kVStartFailureCallback206EE003Func(context206ee003, engine206ee003, config206ee003, msg206ee003)
		return
	}
	panic("callback func has not been set (race?)")
}

var kVStartFailureCallback206EE003Func KVStartFailureCallback

// Ref returns a reference to C object as it is.
func (x *KVEngineSys) Ref() *C.KVEngine {
	if x == nil {
		return nil
	}
	return (*C.KVEngine)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *KVEngineSys) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewKVEngineSysRef converts the C object reference into a raw struct reference without wrapping.
func NewKVEngineSysRef(ref unsafe.Pointer) *KVEngineSys {
	return (*KVEngineSys)(ref)
}

// NewKVEngineSys allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewKVEngineSys() *KVEngineSys {
	return (*KVEngineSys)(allocKVEngineSysMemory(1))
}

// allocKVEngineSysMemory allocates memory for type C.KVEngine in C.
// The caller is responsible for freeing the this memory via C.free.
func allocKVEngineSysMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfKVEngineSysValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfKVEngineSysValue = unsafe.Sizeof([1]C.KVEngine{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *KVEngineSys) PassRef() *C.KVEngine {
	if x == nil {
		x = (*KVEngineSys)(allocKVEngineSysMemory(1))
	}
	return (*C.KVEngine)(unsafe.Pointer(x))
}

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(h.Data), cgoAllocsUnknown
}
