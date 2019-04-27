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

type startFailureCallback struct {
	Func KVStartFailureFunction
}

type allCallback struct {
	Func pmemkv_sys.KVAllFunction
}

type eachCallback struct {
	Func pmemkv_sys.KVEachFunction
}

type getCallback struct {
	Func pmemkv_sys.KVGetFunction
}

type KVStartFailureFunction func(engine string, config string, msg string)

func cbStartWrapper(context unsafe.Pointer, engine string, config string, msg string) {
	cb := pointer.Restore(context).(*startFailureCallback)
	cb.Func(engine, config, msg)
}

func cbAllWrapper(context unsafe.Pointer, keybytes int32, key string) {
	cb := pointer.Restore(context).(*allCallback)
	cb.Func(keybytes, key)
}

func cbGetWrapper(context unsafe.Pointer, keybytes int32, key string) {
	cb := pointer.Restore(context).(*getCallback)
	cb.Func(keybytes, key)
}

func cbEachWrapper(context unsafe.Pointer, keybytes int32, key string, valuebytes int32, value string) {
	cb := pointer.Restore(context).(*eachCallback)
	cb.Func(keybytes, key, valuebytes, value)
}

/// Start function
func (k *KVEngine) Start(engine string, config string, callback KVStartFailureFunction) (KVEngine, error) {
	ke := pmemkv_sys.Kvengine_start(pointer.Save(&startFailureCallback{
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
	pmemkv_sys.Kvengine_all(k.engine, pointer.Save(&allCallback{
		Func: callback,
	}), cbAllWrapper)
}

func (k *KVEngine) AllAbove(kb int32, key string, callback pmemkv_sys.KVAllFunction) {
	pmemkv_sys.Kvengine_all_above(k.engine, pointer.Save(&allCallback{
		Func: callback,
	}), kb, key, cbAllWrapper)
}

func (k *KVEngine) AllBelow(kb int32, key string, callback pmemkv_sys.KVAllFunction) {
	pmemkv_sys.Kvengine_all_below(k.engine, pointer.Save(&allCallback{
		Func: callback,
	}), kb, key, cbAllWrapper)
}

func (k *KVEngine) AllBetween(kb1 int32, key1 string, kb2 int32, key2 string, callback pmemkv_sys.KVAllFunction) {
	pmemkv_sys.Kvengine_all_between(k.engine, pointer.Save(&allCallback{
		Func: callback,
	}), kb1, key1, kb2, key2, cbAllWrapper)
}

func (k *KVEngine) Count() int {
	return pmemkv_sys.Kvengine_count(k.engine)
}

func (k *KVEngine) CountAbove(kb int32, key string) int {
	return pmemkv_sys.Kvengine_count_above(k.engine, kb, key)
}

func (k *KVEngine) CountBelow(kb int32, key string) int {
	return pmemkv_sys.Kvengine_count_below(k.engine, kb, key)
}

func (k *KVEngine) CountBetween(kb1 int32, key1 string, kb2 int32, key2 string) int {
	return pmemkv_sys.Kvengine_count_between(k.engine, kb1, key1, kb2, key2)
}

func (k *KVEngine) Each(callback pmemkv_sys.KVEachFunction) {
	pmemkv_sys.Kvengine_each(k.engine, pointer.Save(&eachCallback{
		Func: callback,
	}), cbEachWrapper)
}

func (k *KVEngine) EachAbove(kb int32, key string, callback pmemkv_sys.KVEachFunction) {
	pmemkv_sys.Kvengine_each_above(k.engine, pointer.Save(&eachCallback{
		Func: callback,
	}), kb, key, cbEachWrapper)
}

func (k *KVEngine) EachBelow(kb int32, key string, callback pmemkv_sys.KVEachFunction) {
	pmemkv_sys.Kvengine_each_below(k.engine, pointer.Save(&eachCallback{
		Func: callback,
	}), kb, key, cbEachWrapper)
}

func (k *KVEngine) EachBetween(kb1 int32, key1 string, kb2 int32, key2 string, callback pmemkv_sys.KVEachFunction) {
	pmemkv_sys.Kvengine_each_between(k.engine, pointer.Save(&eachCallback{
		Func: callback,
	}), kb1, key1, kb2, key2, cbEachWrapper)
}

func (k *KVEngine) Get(kb int32, key string, callback pmemkv_sys.KVGetFunction) {
	pmemkv_sys.Kvengine_get(k.engine, pointer.Save(&getCallback{
		Func: callback,
	}), kb, key, cbGetWrapper)
}

func (k *KVEngine) Exists(kb int32, key string) error {
	res := pmemkv_sys.Kvengine_exists(k.engine, kb, key)
	if res == 1 {
		return nil
	} else if res == -1 {
		return errors.New("Fail to execute")
	} else {
		return errors.New("Key not found")
	}
}

func (k *KVEngine) Get_copy(kb int32, key string, maxvaluebytes int32, value []byte) error {
	res := pmemkv_sys.Kvengine_get_copy(k.engine, kb, key, maxvaluebytes, value)
	if res == 1 {
		return nil
	} else if res == -1 {
		return errors.New("Fail to execute")
	} else {
		return errors.New("Key not found")
	}
}

func (k *KVEngine) Put(kb int32, key string, vb int32, v string) error {
	res := pmemkv_sys.Kvengine_put(k.engine, kb, key, vb, v)
	if res == 1 {
		return nil
	} else if res == -1 {
		return errors.New("Fail to execute")
	} else {
		return errors.New("Key not found")
	}
}

func (k *KVEngine) Remove(kb int32, key string) error {
	res := pmemkv_sys.Kvengine_remove(k.engine, kb, key)
	if res == 1 {
		return nil
	} else if res == -1 {
		return errors.New("Fail to execute")
	} else {
		return errors.New("Key not found")
	}
}
