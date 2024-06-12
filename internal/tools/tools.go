package tools

import (
	"crowform/internal/vector"
	"math"
	"math/rand"
	"reflect"
	"sort"
)

func RandSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// https://github.com/juliangruber/go-intersect/blob/master/intersect.go
//
// Hash has complexity: O(n * x) where x is a factor of hash function efficiency (between 1 and 2)
// func HashGeneric[T comparable](a []T, b []T) []T {
func GetIntersects[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}

func IndexOf[T comparable](slice []T, item T) int {
	for idx := range slice {
		sliceItem := slice[idx]
		if item == sliceItem {
			return idx
		}
	}

	return -1
}

func Splice[T any](slice []T, index int, count int) []T {
	if index < 0 {
		return slice
	}
	if count < 1 {
		count = 1
	}

	return append(slice[:index], slice[index+count:]...)
}

func Remove[T comparable](slice []T, item T) []T {
	idx := IndexOf[T](slice, item)

	return Splice[T](slice, idx, 1)
}
func RemoveAll[T comparable](slice []T, item T) []T {
	var result []T = slice

	idx := IndexOf(result, item)
	for idx != -1 {
		result = Splice(result, idx, 1)
		idx = IndexOf(result, item)
	}

	return result
}

func FindIndex[T comparable](slice []T, predicate func(item T) bool) int {
	for idx := range slice {
		item := slice[idx]
		if predicate(item) {
			return idx
		}
	}

	return -1
}

func InsertSorted[T comparable](s []T, e T, compare func(T) bool) []T {
	insertIdx := sort.Search(len(s), func(i int) bool { return compare(s[i]) })
	if len(s) == 0 || insertIdx == -1 || insertIdx == len(s) {
		return append(s, e)
	}

	s2 := make([]T, len(s)+1)
	// if insertIdx > 0 {
	// 	copy(s2[:insertIdx], s[:insertIdx-1])
	// }
	// copy(s2[insertIdx+1:], s[insertIdx:])
	// s2[insertIdx] = e

	for i := 0; i < len(s2); i++ {
		if i < insertIdx {
			s2[i] = s[i]
		} else if i > insertIdx {
			s2[i] = s[i-1]
		} else {
			s2[insertIdx] = e
		}
	}

	return s2
}

func FilterSlice[T any](slice []T, condition func(T) bool) []T {
	s2 := make([]T, 0)
	for i := range slice {
		item := slice[i]
		if condition(item) {
			s2 = append(s2, item)
		}
	}

	return s2
}

func MapSlice[T any, U any](slice []T, mapper func(T) U) []U {
	s2 := make([]U, 0)
	for i := range slice {
		item := slice[i]
		s2 = append(s2, mapper(item))
	}

	return s2
}

/** Pointer Safe ForEach **/
func ForEach[T any](slice []T, condition func(T)) {
	for i := range slice {
		item := slice[i]
		condition(item)
	}
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func RandomIntFromFloats(min, max float64) int {
	if min == 0 && max == 0 {
		return 0
	}
	if max == min {
		return int(max)
	}
	return int(min) + rand.Intn(int(max)-int(min))
}

func RandomInt(min, max int) int {
	if min == 0 && max == 0 {
		return 0
	}
	if max == min {
		return max
	}
	return min + rand.Intn(max-min)
}

func RandomFloat(min, max float64) float64 {
	if min == 0 && max == 0 {
		return 0
	}
	if max == min {
		return max
	}
	return min + rand.Float64()*(max-min)
}

func GetRandomItem[T any](slice []T) T {
	min := 0
	max := len(slice)

	return slice[RandomInt(min, max)]
}

func Trunc(n float64, prec float64) float64 {
	p := math.Pow(10, prec)
	return float64(int(n*p)) / p
}

func TruncFloat32(n float32) float64 {
	return math.Trunc(float64(n))
}

// a + ((b - a) * t)
func Lerp(v1 vector.Vector, v2 vector.Vector, t float64) vector.Vector {
	// return v1.Add(v2.Subtract(v1).MultiplyScalar(t))

	// Clamping makes it smoother
	clampd := Trunc(t, 2)

	if clampd == 0 {
		return v1
	}
	if clampd == 1 {
		return v2
	}

	if clampd == .5 {
		return vector.Midpoint(v1, v2)
	}

	return v1.Add(v2.Subtract(v1).MultiplyScalar(clampd))
}

func InterfaceToString(iface interface{}) string {
	return reflect.ValueOf(iface).String()
}

func DeepCopyMap[K, V comparable](m1 map[K]V, parse func(interface{}) V) map[K]V {
	copiedMap := make(map[K]V)
	for key, value := range m1 {
		copiedMap[key] = parse(deepCopy(value))
	}
	return copiedMap
}

func deepCopy(item interface{}) interface{} {
	if item == nil {
		return nil
	}

	typ := reflect.TypeOf(item)
	val := reflect.ValueOf(item)

	if typ.Kind() == reflect.Ptr {
		newVal := reflect.New(typ.Elem())
		newVal.Elem().Set(reflect.ValueOf(deepCopy(val.Elem().Interface())))
		return newVal.Interface()
	} else if typ.Kind() == reflect.Map {
		newMap := reflect.MakeMap(typ)
		for _, k := range val.MapKeys() {
			newMap.SetMapIndex(k, reflect.ValueOf(deepCopy(val.MapIndex(k).Interface())))
		}
		return newMap.Interface()
	} else if typ.Kind() == reflect.Slice {
		newSlice := reflect.MakeSlice(typ, val.Len(), val.Cap())
		for i := 0; i < val.Len(); i++ {
			newSlice.Index(i).Set(reflect.ValueOf(deepCopy(val.Index(i).Interface())))
		}
		return newSlice.Interface()
	}

	return item
}

// Got tired/scared of nil checking
type Maybe[T any] struct {
	Value    T
	hasValue bool
}

func MaybeWithValue[T any](value T) *Maybe[T] {
	m := &Maybe[T]{
		Value:    value,
		hasValue: true,
	}
	return m
}

func (maybe *Maybe[T]) SetValue(value T) *Maybe[T] {
	maybe.Value = value
	maybe.hasValue = true
	return maybe
}
func (maybe *Maybe[T]) HasValue() bool {
	return maybe.hasValue
}

func MaxIntOf(nInts ...int32) int32 {
	var m int32 = 0

	for _, n := range nInts {
		if m < n {
			m = n
		}
	}

	return m
}
func MaxFloat32Of(nFloats ...float32) float32 {
	var m float32 = 0

	for _, n := range nFloats {
		if m < n {
			m = n
		}
	}

	return m
}

func MinIntOf(nInts ...int32) int32 {
	var m int32 = math.MaxInt32

	for _, n := range nInts {
		if m > n {
			m = n
		}
	}

	return m
}

func ToEvenFloat32(f32 float32) float32 {
	return float32(math.RoundToEven(float64(f32)))
}
