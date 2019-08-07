package try

import "reflect"

func Try(f func()) *tryStruct {
	return &tryStruct{
		catches: make(map[reflect.Type]ExeceptionHandler),
		hold:    f,
	}
}

type ExeceptionHandler func(interface{})

type tryStruct struct {
	catches map[reflect.Type]ExeceptionHandler
	hold    func()
}

func (t *tryStruct) Catch(f func()) {
	defer func() {
		if e := recover(); nil != e {
			if h, ok := t.catches[reflect.TypeOf(e)]; ok {
				h(e)
			} else {
				f()
			}
		}
	}()

	t.hold()
}
