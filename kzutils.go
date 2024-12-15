package kzutils

import (
	"errors"
	"math"
	"slices"
	"strings"
	"time"

	"golang.org/x/exp/constraints"
)

func TakePtr[T any](val T) *T     { return &val }
func TakePointer[T any](val T) *T { return &val }
func TakeAddress[T any](val T) *T { return &val }

func TakeValue[T any](ptr *T) T {
	if ptr == nil {
		var z T
		return z
	}
	return *ptr
}

func IsIn[T comparable](value T, okValues []T) bool {
	for _, okValue := range okValues {
		if value == okValue {
			return true
		}
	}
	return false
}
func IsOneOf[T comparable](value T, okValues ...T) bool {
	return IsIn(value, okValues)
}

func TrimStringWithSpaces(str string, cutset string) string {
	str1 := ""
	str2 := str
	for str1 != str2 {
		str1 = str2
		str2 = strings.TrimSpace(str2)
		str2 = strings.Trim(str2, cutset)
	}
	return str2
}
func TrimStringSpacesAndQuotes(str string) string {
	return TrimStringWithSpaces(str, `'"`)
}
func TrimStringWithSpacesAndQuotes(str string, cutset string) string {
	return TrimStringWithSpaces(str, `'"`+cutset)
}

func InitSlice[V any](sl []V, cnts ...int) []V {
	if sl == nil {
		var makeCnt int
		for _, cnt := range cnts {
			makeCnt += cnt
		}
		sl = make([]V, 0, makeCnt)
	}
	return sl
}
func InitMap[K comparable, V any](mp map[K]V, cnts ...int) map[K]V {
	if mp == nil {
		var makeCnt int
		for _, cnt := range cnts {
			makeCnt += cnt
		}
		mp = make(map[K]V, makeCnt)
	}
	return mp
}
func InitMapKey[K comparable, V any](mp map[K]V, key K, cnts ...int) map[K]V {
	mp = InitMap[K, V](mp, cnts...)
	if _, ok := mp[key]; !ok {
		var zv V
		mp[key] = zv
	}
	return mp
}

func Abs[T constraints.Integer | constraints.Float](vl T) T {
	return T(math.Abs(float64(vl)))
}

func RoundDuration(d time.Duration, base time.Duration, decimals int) (ret time.Duration) {
	if d == 0 {
		return d
	}
	for Abs(base) > Abs(d) {
		base = base / 10
	}

	decimalsPow10 := math.Pow10(-1 * decimals)

	baseFloat := float64(base)
	baseFloat = baseFloat * decimalsPow10
	if baseFloat == 0 {
		return 0
	}

	df := float64(d)
	df = df / baseFloat
	df = math.Round(df)
	df = df * baseFloat

	ret = time.Duration(int64(df))
	if ret == 0 {
		return d
	}

	return ret
}

func GrowSliceOn[T any](sl []T, additionalCapacity int) (ret []T) {
	return GrowSliceTo(sl, len(sl)+additionalCapacity)
}
func GrowSliceTo[T any](sl []T, targetCapacity int) (ret []T) {
	if targetCapacity <= cap(sl) {
		return sl
	}
	return slices.Grow(ret, targetCapacity)
}

func SlicesIntersect[T comparable](sl1 []T, sl2 []T) (ret []T) {
	ret = make([]T, 0, min(len(sl1), len(sl2)))
	if len(sl1) <= 0 {
		return ret
	}
	if len(sl2) <= 0 {
		return ret
	}
	type emp struct{}
	mp2 := make(map[T]emp, len(sl2))
	for _, sl2v := range sl2 {
		mp2[sl2v] = emp{}
	}
	for _, sl1v := range sl1 {
		if _, ok := mp2[sl1v]; ok {
			ret = append(ret, sl1v)
		}
	}
	return ret
}
func IsSliceContainsOneOf[T comparable](sl []T, okValues ...T) bool {
	isect := SlicesIntersect(sl, okValues)
	return len(isect) > 0
}

var (
	ErrGetOnlyOneSliceItemSliceIsEmpty    = errors.New("cannot get only one slice item - slice is empty")
	ErrGetOnlyOneSliceItemMoreThanOneItem = errors.New("cannot get only one slice item - slice contains more than one item")
)

func GetOnlyOneSliceItem[T any](items []T, errs ...error) (ret T, err error) {
	var zt T
	if len(items) < 1 {
		if len(errs) > 0 {
			return zt, errs[0]
		} else {
			return zt, ErrGetOnlyOneSliceItemSliceIsEmpty
		}
	}
	if len(items) > 1 {
		if len(errs) > 1 {
			return zt, errs[1]
		} else if len(errs) > 0 {
			return zt, errs[0]
		} else {
			return zt, ErrGetOnlyOneSliceItemMoreThanOneItem
		}
	}
	return items[0], nil
}
