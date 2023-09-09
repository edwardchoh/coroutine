// Code generated by coroc. DO NOT EDIT

//go:build durable

package testdata

import "github.com/stealthrocket/coroutine"

func Identity(n int) {
	_c := coroutine.LoadContext[int, any]()
	_f := _c.Push()
	if _f.IP > 0 {
		n = _f.Get(0).(int)
	}
	defer func() {
		if _c.Unwinding() {
			_f.Set(0, n)
		} else {
			_c.Pop()
		}
	}()
	coroutine.Yield[int, any](n)
}

func SquareGenerator(n int) {
	_c := coroutine.LoadContext[int, any]()
	_f := _c.Push()
	var _v0 int
	if _f.IP > 0 {
		n = _f.Get(0).(int)
		_v0 = _f.Get(1).(int)
	}
	defer func() {
		if _c.Unwinding() {
			_f.Set(0, n)
			_f.Set(1, _v0)
		} else {
			_c.Pop()
		}
	}()
	switch {
	case _f.IP < 2:
		_v0 = 1
		_f.IP = 2
		fallthrough
	case _f.IP < 3:
		for ; _v0 <= n; _v0++ {
			coroutine.Yield[int, any](_v0 * _v0)
			_f.IP = 2
		}
	}
}

func EvenSquareGenerator(n int) {
	_c := coroutine.LoadContext[int, any]()
	_f := _c.Push()
	var (
		_v0 int
		_v1 int
	)
	if _f.IP > 0 {
		n = _f.Get(0).(int)
		_v0 = _f.Get(1).(int)
		_v1 = _f.Get(2).(int)
	}
	defer func() {
		if _c.Unwinding() {
			_f.Set(0, n)
			_f.Set(1, _v0)
			_f.Set(2, _v1)
		} else {
			_c.Pop()
		}
	}()
	switch {
	case _f.IP < 2:
		_v0 = 1
		_f.IP = 2
		fallthrough
	case _f.IP < 4:
		for ; _v0 <= n; _v0++ {
			switch {
			case _f.IP < 3:
				_v1 = _v0 % 2
				_f.IP = 3
				fallthrough
			case _f.IP < 4:
				if _v1 == 0 {
					coroutine.Yield[int, any](_v0 * _v0)
				}
			}
			_f.IP = 2
		}
	}
}

func NestedLoops(n int) {
	_c := coroutine.LoadContext[int, any]()
	_f := _c.Push()
	var (
		_v0 int
		_v1 int
		_v2 int
	)
	if _f.IP > 0 {
		n = _f.Get(0).(int)
		_v0 = _f.Get(1).(int)
		_v1 = _f.Get(2).(int)
		_v2 = _f.Get(3).(int)
	}
	defer func() {
		if _c.Unwinding() {
			_f.Set(0, n)
			_f.Set(1, _v0)
			_f.Set(2, _v1)
			_f.Set(3, _v2)
		} else {
			_c.Pop()
		}
	}()
	switch {
	case _f.IP < 2:
		_v0 = 1
		_f.IP = 2
		fallthrough
	case _f.IP < 5:
		for ; _v0 <= n; _v0++ {
			switch {
			case _f.IP < 3:
				_v1 = 1
				_f.IP = 3
				fallthrough
			case _f.IP < 5:
				for ; _v1 <= n; _v1++ {
					switch {
					case _f.IP < 4:
						_v2 = 1
						_f.IP = 4
						fallthrough
					case _f.IP < 5:
						for ; _v2 <= n; _v2++ {
							coroutine.Yield[int, any](_v0 * _v1 * _v2)
							_f.IP = 4
						}
					}
					_f.IP = 3
				}
			}
			_f.IP = 2
		}
	}
}

func FizzBuzzIfGenerator(n int) {
	_c := coroutine.LoadContext[int, any]()
	_f := _c.Push()
	var (
		_v0 int
		_v1 int
	)
	if _f.IP > 0 {
		n = _f.Get(0).(int)
		_v0 = _f.Get(1).(int)
		_v1 = _f.Get(2).(int)
	}
	defer func() {
		if _c.Unwinding() {
			_f.Set(0, n)
			_f.Set(1, _v0)
			_f.Set(2, _v1)
		} else {
			_c.Pop()
		}
	}()
	switch {
	case _f.IP < 2:
		_v0 = 1
		_f.IP = 2
		fallthrough
	case _f.IP < 7:
		for ; _v0 <= n; _v0++ {
			if _v0%3 == 0 && _v0%5 == 0 {
				coroutine.Yield[int, any](FizzBuzz)
			} else if _v0%3 == 0 {
				coroutine.Yield[int, any](Fizz)
			} else {
				_v1 = _v0 % 5
				if _v1 == 0 {
					coroutine.Yield[int, any](Buzz)
				} else {
					coroutine.Yield[int, any](_v0)
				}
			}
			_f.IP = 2
		}
	}
}

func FizzBuzzSwitchGenerator(n int) {
	_c := coroutine.LoadContext[int, any]()
	_f := _c.Push()
	var _v0 int
	if _f.IP > 0 {
		n = _f.Get(0).(int)
		_v0 = _f.Get(1).(int)
	}
	defer func() {
		if _c.Unwinding() {
			_f.Set(0, n)
			_f.Set(1, _v0)
		} else {
			_c.Pop()
		}
	}()
	switch {
	case _f.IP < 2:
		_v0 = 1
		_f.IP = 2
		fallthrough
	case _f.IP < 6:
		for ; _v0 <= n; _v0++ {
			switch {
			case _v0%3 == 0 && _v0%5 == 0:
				coroutine.Yield[int, any](FizzBuzz)
			case _v0%3 == 0:
				coroutine.Yield[int, any](Fizz)
			case _v0%5 == 0:
				coroutine.Yield[int, any](Buzz)
			default:
				coroutine.Yield[int, any](_v0)
			}
			_f.IP = 2
		}
	}
}

func Shadowing(_ int) {
	_c := coroutine.LoadContext[int, any]()
	_f := _c.Push()
	var (
		_v0 int
		_v1 int
		_v2 int
		_v3 int
	)
	if _f.IP > 0 {
		_v0 = _f.Get(0).(int)
		_v1 = _f.Get(1).(int)
		_v2 = _f.Get(2).(int)
		_v3 = _f.Get(3).(int)
	}
	defer func() {
		if _c.Unwinding() {
			_f.Set(0, _v0)
			_f.Set(1, _v1)
			_f.Set(2, _v2)
			_f.Set(3, _v3)
		} else {
			_c.Pop()
		}
	}()
	switch {
	case _f.IP < 2:
		_v0 = 0
		_f.IP = 2
		fallthrough
	case _f.IP < 3:
		coroutine.Yield[int, any](_v0)
		_f.IP = 3
		fallthrough
	case _f.IP < 5:
		switch {
		case _f.IP < 4:
			_v1 = 1
			_f.IP = 4
			fallthrough
		case _f.IP < 5:
			if true {
				coroutine.Yield[int, any](_v1)
			}
		}
		_f.IP = 5
		fallthrough
	case _f.IP < 6:
		coroutine.Yield[int, any](_v0)
		_f.IP = 6
		fallthrough
	case _f.IP < 8:
		switch {
		case _f.IP < 7:
			_v2 = 1
			_f.IP = 7
			fallthrough
		case _f.IP < 8:
			for ; _v2 < 3; _v2++ {
				coroutine.Yield[int, any](_v2)
				_f.IP = 7
			}
		}
		_f.IP = 8
		fallthrough
	case _f.IP < 9:
		coroutine.Yield[int, any](_v0)
		_f.IP = 9
		fallthrough
	case _f.IP < 11:
		switch {
		case _f.IP < 10:
			_v3 = 1
			_f.IP = 10
			fallthrough
		case _f.IP < 11:
			switch _v3 {
			case 1:
				coroutine.Yield[int, any](_v3)
			}
		}
		_f.IP = 11
		fallthrough
	case _f.IP < 12:
		coroutine.Yield[int, any](_v0)
	}
}

func RangeSliceIndexGenerator(_ int) {
	_c := coroutine.LoadContext[int, any]()
	_f := _c.Push()
	var (
		_v0 []int
		_v1 int
	)
	if _f.IP > 0 {
		_v0 = _f.Get(0).([]int)
		_v1 = _f.Get(1).(int)
	}
	defer func() {
		if _c.Unwinding() {
			_f.Set(0, _v0)
			_f.Set(1, _v1)
		} else {
			_c.Pop()
		}
	}()
	switch {
	case _f.IP < 2:
		_v0 = []int{10, 20, 30}
		_f.IP = 2
		fallthrough
	case _f.IP < 4:
		switch {
		case _f.IP < 3:
			_v1 = 0
			_f.IP = 3
			fallthrough
		case _f.IP < 4:
			for ; _v1 < len(_v0); _v1++ {
				coroutine.Yield[int, any](_v1)
				_f.IP = 3
			}
		}
	}
}

func RangeArrayIndexValueGenerator(_ int) {
	_c := coroutine.LoadContext[int, any]()
	_f := _c.Push()
	var (
		_v0 [3]int
		_v1 int
		_v2 int
	)
	if _f.IP > 0 {
		_v0 = _f.Get(0).([3]int)
		_v1 = _f.Get(1).(int)
		_v2 = _f.Get(2).(int)
	}
	defer func() {
		if _c.Unwinding() {
			_f.Set(0, _v0)
			_f.Set(1, _v1)
			_f.Set(2, _v2)
		} else {
			_c.Pop()
		}
	}()
	switch {
	case _f.IP < 2:
		_v0 = [...]int{10, 20, 30}
		_f.IP = 2
		fallthrough
	case _f.IP < 6:
		switch {
		case _f.IP < 3:
			_v1 = 0
			_f.IP = 3
			fallthrough
		case _f.IP < 6:
			for ; _v1 < len(_v0); _v1++ {
				switch {
				case _f.IP < 4:
					_v2 = _v0[_v1]
					_f.IP = 4
					fallthrough
				case _f.IP < 5:
					coroutine.Yield[int, any](_v1)
					_f.IP = 5
					fallthrough
				case _f.IP < 6:
					coroutine.Yield[int, any](_v2)
				}
				_f.IP = 3
			}
		}
	}
}
