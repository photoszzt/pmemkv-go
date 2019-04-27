package pmemkv

import (
	"errors"
	"github.com/mattn/go-pointer"
	"github.com/photoszzt/pmemkv-go/pmemkv_sys"
	"unsafe"
)

type KVEngine struct {
	engine *pmemkv_sys.KVEngineSys
}

type StartFailureCallback struct {
	Func KVStartFailureFunction
}

type AllCallback struct {
	Func pmemkv_sys.KVAllFunction
}

type KVStartFailureFunction func(engine string, config string, msg string)

func cbStartWrapper(context unsafe.Pointer, engine string, config string, msg string) {
	cb := pointer.Restore(context).(*StartFailureCallback)
	cb.Func(engine, config, msg)
}

func cbAllWrapper(context unsafe.Pointer, keybytes int32, key string) {
	cb := pointer.Restore(context).(*AllCallback)
	cb.Func(keybytes, key)
}

/// Start function
func (k *KVEngine) Start(engine string, config string, callback KVStartFailureFunction) (KVEngine, error) {
	ke := pmemkv_sys.Kvengine_start(pointer.Save(&StartFailureCallback{
		Func: callback,
	}), engine, config, cbStartWrapper)
	if ke == nil {
		return KVEngine{engine: nil}, errors.New("Fail to start KVEngine")
	} else {
		return KVEngine{engine: ke}, nil
	}
}

func (k *KVEngine) Stop() {
	pmemkv_sys.Kvengine_stop(k.engine)
}

func (k *KVEngine) All(callback pmemkv_sys.KVAllFunction) {
	pmemkv_sys.Kvengine_all(k.engine, pointer.Save(&AllCallback{
		Func: callback,
	}), cbAllWrapper)
}

func (k *KVEngine) AllAbove(kb int32, key string, callback pmemkv_sys.KVAllFunction) {
	pmemkv_sys.Kvengine_all_above(k.engine, pointer.Save(&AllCallback{
		Func: callback,
	}), kb, key, cbAllWrapper)
}

func (k *KVEngine) AllBelow(kb int32, key string, callback pmemkv_sys.KVAllFunction) {
	pmemkv_sys.Kvengine_all_below(k.engine, pointer.Save(&AllCallback{
		Func: callback,
	}), kb, key, cbAllWrapper)
}

func (k *KVEngine) AllBetween(kb1 int32, key1 string, kb2 int32, key2 string, callback pmemkv_sys.KVAllFunction) {
	pmemkv_sys.Kvengine_all_between(k.engine, pointer.Save(&AllCallback{
		Func: callback,
	}), kb1, key1, kb2, key2, cbAllWrapper)
}

func (k *KVEngine) Count() int {
	return pmemkv_sys.Kvengine_count(k.engine)
}
