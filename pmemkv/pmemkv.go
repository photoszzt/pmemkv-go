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

type KVFunction func(key string)
type KVEachFunction func(key string, value string)
type KVStartFailureFunction func(engine string, config string, msg string)

type kvCallback struct {
	Func KVFunction
}

type eachCallback struct {
	Func KVEachFunction
}

func cbStartWrapper(context unsafe.Pointer, engine string, config string, msg string) {
	cb := pointer.Restore(context).(*startFailureCallback)
	cb.Func(engine, config, msg)
}

func cbWrapper(context unsafe.Pointer, keybytes int32, key string) {
	cb := pointer.Restore(context).(*kvCallback)
	cb.Func(key)
}

func cbEachWrapper(context unsafe.Pointer, keybytes int32, key string, valuebytes int32, value string) {
	cb := pointer.Restore(context).(*eachCallback)
	cb.Func(key, value)
}

/// Start function
func StartKVEngine(engine string, config string, callback KVStartFailureFunction) (KVEngine, error) {
	engine = engine + "\x00"
	config = config + "\x00"
	ke := pmemkv_sys.Start(pointer.Save(&startFailureCallback{
		Func: callback,
	}), engine, config, cbStartWrapper)
	if ke == nil {
		return KVEngine{engine: nil}, errors.New("Fail to start KVEngine")
	} else {
		return KVEngine{engine: ke}, nil
	}
}

func (k *KVEngine) Stop() {
	pmemkv_sys.Stop(k.engine)
}

func (k *KVEngine) All(callback KVFunction) {
	pmemkv_sys.All(k.engine, pointer.Save(&kvCallback{
		Func: callback,
	}), cbWrapper)
}

func (k *KVEngine) AllAbove(key string, callback KVFunction) {
	key = key + "\x00"
	pmemkv_sys.AllAbove(k.engine, pointer.Save(&kvCallback{
		Func: callback,
	}), int32(len(key)), key, cbWrapper)
}

func (k *KVEngine) AllBelow(key string, callback KVFunction) {
	key = key + "\x00"
	pmemkv_sys.AllBelow(k.engine, pointer.Save(&kvCallback{
		Func: callback,
	}), int32(len(key)), key, cbWrapper)
}

func (k *KVEngine) AllBetween(key1 string, key2 string, callback KVFunction) {
	key1 = key1 + "\x00"
	key2 = key2 + "\x00"
	pmemkv_sys.AllBetween(k.engine, pointer.Save(&kvCallback{
		Func: callback,
	}), int32(len(key1)), key1, int32(len(key2)), key2, cbWrapper)
}

func (k *KVEngine) Count() int {
	return pmemkv_sys.Count(k.engine)
}

func (k *KVEngine) CountAbove(key string) int {
	key = key + "\x00"
	return pmemkv_sys.CountAbove(k.engine, int32(len(key)), key)
}

func (k *KVEngine) CountBelow(key string) int {
	key = key + "\x00"
	return pmemkv_sys.CountBelow(k.engine, int32(len(key)), key)
}

func (k *KVEngine) CountBetween(key1 string, key2 string) int {
	key1 = key1 + "\x00"
	key2 = key2 + "\x00"
	return pmemkv_sys.CountBetween(k.engine, int32(len(key1)), key1, int32(len(key2)), key2)
}

func (k *KVEngine) Each(callback KVEachFunction) {
	pmemkv_sys.Each(k.engine, pointer.Save(&eachCallback{
		Func: callback,
	}), cbEachWrapper)
}

func (k *KVEngine) EachAbove(key string, callback KVEachFunction) {
	key = key + "\x00"
	pmemkv_sys.EachAbove(k.engine, pointer.Save(&eachCallback{
		Func: callback,
	}), int32(len(key)), key, cbEachWrapper)
}

func (k *KVEngine) EachBelow(key string, callback KVEachFunction) {
	key = key + "\x00"
	pmemkv_sys.EachBelow(k.engine, pointer.Save(&eachCallback{
		Func: callback,
	}), int32(len(key)), key, cbEachWrapper)
}

func (k *KVEngine) EachBetween(key1 string, key2 string, callback KVEachFunction) {
	key1 = key1 + "\x00"
	key2 = key2 + "\x00"
	pmemkv_sys.EachBetween(k.engine, pointer.Save(&eachCallback{
		Func: callback,
	}), int32(len(key1)), key1, int32(len(key2)), key2, cbEachWrapper)
}

func (k *KVEngine) Get(key string, callback KVFunction) {
	key = key + "\x00"
	pmemkv_sys.Get(k.engine, pointer.Save(&kvCallback{
		Func: callback,
	}), int32(len(key)), key, cbWrapper)
}

func (k *KVEngine) Exists(key string) error {
	key = key + "\x00"
	res := pmemkv_sys.Exists(k.engine, int32(len(key)), key)
	return pmemkvError(res)
}

func (k *KVEngine) GetCopy(key string, value []byte) error {
	key = key + "\x00"
	res := pmemkv_sys.GetCopy(k.engine, int32(len(key)), key, int32(cap(value)), (*byte)(unsafe.Pointer(&value)))
	return pmemkvError(res)
}

func (k *KVEngine) Put(key string, v string) error {
	key = key + "\x00"
	v = v + "\x00"
	res := pmemkv_sys.Put(k.engine, int32(len(key)), key, int32(len(v)), v)
	return pmemkvError(res)
}

func (k *KVEngine) Remove(key string) error {
	key = key + "\x00"
	res := pmemkv_sys.Remove(k.engine, int32(len(key)), key)
	return pmemkvError(res)
}
