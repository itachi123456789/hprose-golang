/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * io/reader_test.go                                      *
 *                                                        *
 * hprose Reader Test for Go.                             *
 *                                                        *
 * LastModified: Sep 9, 2016                              *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

import (
	"math"
	"reflect"
	"testing"
	"time"
)

func TestReadBool(t *testing.T) {
	trueValue := "true"
	data := map[interface{}]bool{
		true:            true,
		false:           false,
		nil:             false,
		"":              false,
		0:               false,
		1:               true,
		9:               true,
		100:             true,
		100000000000000: true,
		0.0:             false,
		"t":             true,
		"f":             false,
		&trueValue:      true,
		&trueValue:      true,
		"false":         false,
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&trueValue)
	reader := NewReader(w.Bytes(), false)
	for _, k := range keys {
		b := reader.ReadBool()
		if b != data[k] {
			t.Error(k, data[k], b)
		}
	}
	b := reader.ReadBool()
	if b != true {
		t.Error(trueValue, true, b)
	}
	w.Close()
}

func TestUnserializeBool(t *testing.T) {
	trueValue := "true"
	data := map[interface{}]bool{
		true:            true,
		false:           false,
		nil:             false,
		"":              false,
		0:               false,
		1:               true,
		9:               true,
		100:             true,
		100000000000000: true,
		0.0:             false,
		"t":             true,
		"f":             false,
		&trueValue:      true,
		"false":         false,
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&trueValue)
	reader := NewReader(w.Bytes(), false)
	var p bool
	for _, k := range keys {
		reader.Unserialize(&p)
		if p != data[k] {
			t.Error(k, data[k], p)
		}
	}
	reader.Unserialize(&p)
	if p != true {
		t.Error(trueValue, true, p)
	}
	w.Close()
}

func BenchmarkReadBool(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(true)
	bytes := w.Bytes()
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.ReadBool()
	}
	w.Close()
}

func BenchmarkUnserializeBool(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(true)
	bytes := w.Bytes()
	var p bool
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
	w.Close()
}

func TestReadInt(t *testing.T) {
	intValue := "1234567"
	u := uint(math.MaxUint64)
	data := map[interface{}]int64{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		-100:                             -100,
		math.MinInt32:                    int64(math.MinInt32),
		math.MaxInt64:                    int64(math.MaxInt64),
		math.MinInt64:                    int64(math.MinInt64),
		u:                                int64(u),
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&intValue:                        1234567,
		time.Unix(123, 456):              123000000456,
		time.Unix(1234567890, 123456789): 1234567890123456789,
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&intValue)
	reader := NewReader(w.Bytes(), false)
	for _, k := range keys {
		i := reader.ReadInt()
		if i != data[k] {
			t.Error(k, data[k], i)
		}
	}
	i := reader.ReadInt()
	if i != 1234567 {
		t.Error(intValue, 1234567, i)
	}
	w.Close()
}

func TestUnserializeInt(t *testing.T) {
	intValue := "1234567"
	u := uint(math.MaxUint64)
	data := map[interface{}]int{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		-100:                             -100,
		math.MinInt32:                    int(math.MinInt32),
		math.MaxInt64:                    int(math.MaxInt64),
		math.MinInt64:                    int(math.MinInt64),
		u:                                int(u),
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&intValue:                        1234567,
		time.Unix(123, 456):              123000000456,
		time.Unix(1234567890, 123456789): 1234567890123456789,
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&intValue)
	reader := NewReader(w.Bytes(), false)
	var p int
	for _, k := range keys {
		reader.Unserialize(&p)
		if p != data[k] {
			t.Error(k, data[k], p)
		}
	}
	reader.Unserialize(&p)
	if p != 1234567 {
		t.Error(intValue, 1234567, p)
	}
	w.Close()
}

func BenchmarkReadInt(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(12345)
	bytes := w.Bytes()
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.ReadInt()
	}
	w.Close()
}

func BenchmarkUnserializeInt(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(12345)
	bytes := w.Bytes()
	var p int
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
	w.Close()
}

func TestReadUint(t *testing.T) {
	intValue := "1234567"
	u := uint(math.MaxUint64)
	data := map[interface{}]uint64{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		math.MaxInt64:                    uint64(math.MaxInt64),
		u:                                uint64(u),
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&intValue:                        1234567,
		time.Unix(123, 456):              123000000456,
		time.Unix(1234567890, 123456789): 1234567890123456789,
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&intValue)
	reader := NewReader(w.Bytes(), false)
	for _, k := range keys {
		i := reader.ReadUint()
		if i != data[k] {
			t.Error(k, data[k], i)
		}
	}
	i := reader.ReadUint()
	if i != 1234567 {
		t.Error(intValue, 1234567, i)
	}
	w.Close()
}

func TestUnserializeUint(t *testing.T) {
	intValue := "1234567"
	u := uint(math.MaxUint64)
	data := map[interface{}]uint{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		math.MaxInt64:                    uint(math.MaxInt64),
		u:                                uint(u),
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&intValue:                        1234567,
		time.Unix(123, 456):              123000000456,
		time.Unix(1234567890, 123456789): 1234567890123456789,
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&intValue)
	reader := NewReader(w.Bytes(), false)
	var p uint
	for _, k := range keys {
		reader.Unserialize(&p)
		if p != data[k] {
			t.Error(k, data[k], p)
		}
	}
	reader.Unserialize(&p)
	if p != 1234567 {
		t.Error(intValue, 1234567, p)
	}
	w.Close()
}

func BenchmarkReadUint(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(12345)
	bytes := w.Bytes()
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.ReadUint()
	}
	w.Close()
}

func BenchmarkUnserializeUint(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(12345)
	bytes := w.Bytes()
	var p uint
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
	w.Close()
}

func TestReadFloat32(t *testing.T) {
	floatValue := "3.14159"
	data := map[interface{}]float32{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		math.MaxInt64:                    float32(math.MaxInt64),
		math.MaxFloat32:                  math.MaxFloat32,
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&floatValue:                      3.14159,
		time.Unix(123, 456):              float32(123.000000456),
		time.Unix(1234567890, 123456789): float32(1234567890.123456789),
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&floatValue)
	reader := NewReader(w.Bytes(), false)
	for _, k := range keys {
		x := reader.ReadFloat32()
		if x != data[k] {
			t.Error(k, data[k], x)
		}
	}
	x := reader.ReadFloat32()
	if x != float32(3.14159) {
		t.Error(floatValue, 3.14159, x)
	}
	w.Close()
}

func TestUnserializeFloat32(t *testing.T) {
	floatValue := "3.14159"
	data := map[interface{}]float32{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		math.MaxInt64:                    float32(math.MaxInt64),
		math.MaxFloat32:                  math.MaxFloat32,
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&floatValue:                      3.14159,
		time.Unix(123, 456):              float32(123.000000456),
		time.Unix(1234567890, 123456789): float32(1234567890.123456789),
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&floatValue)
	reader := NewReader(w.Bytes(), false)
	var p float32
	for _, k := range keys {
		reader.Unserialize(&p)
		if p != data[k] {
			t.Error(k, data[k], p)
		}
	}
	reader.Unserialize(&p)
	if p != float32(3.14159) {
		t.Error(floatValue, 3.14159, p)
	}
	w.Close()
}

func BenchmarkReadFloat32(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(3.14159)
	bytes := w.Bytes()
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.ReadFloat32()
	}
	w.Close()
}

func BenchmarkUnserializeFloat32(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(3.14159)
	bytes := w.Bytes()
	var p float32
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
	w.Close()
}

func TestReadFloat64(t *testing.T) {
	floatValue := "3.14159"
	data := map[interface{}]float64{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		math.MaxFloat32:                  math.MaxFloat32,
		math.MaxFloat64:                  math.MaxFloat64,
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&floatValue:                      3.14159,
		time.Unix(123, 456):              float64(123.000000456),
		time.Unix(1234567890, 123456789): float64(1234567890.123456789),
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&floatValue)
	reader := NewReader(w.Bytes(), false)
	for _, k := range keys {
		x := reader.ReadFloat64()
		if x != data[k] {
			t.Error(k, data[k], x)
		}
	}
	x := reader.ReadFloat64()
	if x != float64(3.14159) {
		t.Error(floatValue, 3.14159, x)
	}
	w.Close()
}

func TestUnserializeFloat64(t *testing.T) {
	floatValue := "3.14159"
	data := map[interface{}]float64{
		true:                             1,
		false:                            0,
		nil:                              0,
		"":                               0,
		0:                                0,
		1:                                1,
		9:                                9,
		100:                              100,
		math.MaxFloat32:                  math.MaxFloat32,
		math.MaxFloat64:                  math.MaxFloat64,
		0.0:                              0,
		"1":                              1,
		"9":                              9,
		&floatValue:                      3.14159,
		time.Unix(123, 456):              float64(123.000000456),
		time.Unix(1234567890, 123456789): float64(1234567890.123456789),
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&floatValue)
	reader := NewReader(w.Bytes(), false)
	var p float64
	for _, k := range keys {
		reader.Unserialize(&p)
		if p != data[k] {
			t.Error(k, data[k], p)
		}
	}
	reader.Unserialize(&p)
	if p != float64(3.14159) {
		t.Error(floatValue, 3.14159, p)
	}
	w.Close()
}

func BenchmarkReadFloat64(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(3.14159)
	bytes := w.Bytes()
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.ReadFloat64()
	}
	w.Close()
}

func BenchmarkUnserializeFloat64(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(3.14159)
	bytes := w.Bytes()
	var p float64
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
	w.Close()
}

func TestUnserializeArray(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]byte{'h', 'e', 'l', 'l', 'o'}
	w := NewWriter(false)
	w.Serialize(&a)
	w.Serialize(&b)
	w.Serialize(&a)
	w.Serialize(&b)
	w.Serialize(&a)
	w.Serialize(&b)
	reader := NewReader(w.Bytes(), false)
	var a1 [5]int
	reader.Unserialize(&a1)
	if !reflect.DeepEqual(a1, a) {
		t.Error(a1, a)
	}
	var b1 [5]byte
	reader.Unserialize(&b1)
	if !reflect.DeepEqual(b1, b) {
		t.Error(b1, b)
	}
	var a2 [4]int
	reader.Unserialize(&a2)
	if !reflect.DeepEqual(a2[:4], a[:4]) {
		t.Error(a2[:4], a[:4])
	}
	var b2 [4]byte
	reader.Unserialize(&b2)
	if !reflect.DeepEqual(b2[:4], b[:4]) {
		t.Error(b2[:4], b[:4])
	}
	var a3 [6]int
	reader.Unserialize(&a3)
	if !reflect.DeepEqual(a3, [6]int{1, 2, 3, 4, 5, 0}) {
		t.Error(a3)
	}
	var b3 [6]byte
	reader.Unserialize(&b3)
	if !reflect.DeepEqual(b3, [6]byte{'h', 'e', 'l', 'l', 'o', 0}) {
		t.Error(b3)
	}
	w.Close()
}

func BenchmarkUnserializeByteArray(b *testing.B) {
	w := NewWriter(true)
	w.Serialize([5]byte{'h', 'e', 'l', 'l', 'o'})
	bytes := w.Bytes()
	var p [5]byte
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
	w.Close()
}

func BenchmarkUnserializeIntArray(b *testing.B) {
	w := NewWriter(true)
	w.Serialize([5]int{1, 2, 3, 4, 5})
	bytes := w.Bytes()
	var p [5]int
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
	w.Close()
}

func TestUnserializeSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []byte{'h', 'e', 'l', 'l', 'o'}
	w := NewWriter(false)
	w.Serialize(a)
	w.Serialize(b)
	w.Serialize(a)
	w.Serialize(b)
	w.Serialize(a)
	w.Serialize(b)
	reader := NewReader(w.Bytes(), false)
	var a1 []int
	reader.Unserialize(&a1)
	if !reflect.DeepEqual(a1, a) {
		t.Error(a1, a)
	}
	var b1 []byte
	reader.Unserialize(&b1)
	if !reflect.DeepEqual(b1, b) {
		t.Error(b1, b)
	}
	a2 := []int{}
	reader.Unserialize(&a2)
	if !reflect.DeepEqual(a2, a) {
		t.Error(a2, a)
	}
	b2 := []byte{}
	reader.Unserialize(&b2)
	if !reflect.DeepEqual(b2, b) {
		t.Error(b2, b)
	}
	a2 = make([]int, 10)
	reader.Unserialize(&a2)
	if !reflect.DeepEqual(a2, a) {
		t.Error(a2, a)
	}
	b2 = make([]byte, 10)
	reader.Unserialize(&b2)
	if !reflect.DeepEqual(b2, b) {
		t.Error(b2, b)
	}
	w.Close()
}

func TestUnserializeSliceRef(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []byte{'h', 'e', 'l', 'l', 'o'}
	w := NewWriter(false)
	w.Serialize(&a)
	w.Serialize(&b)
	w.Serialize(&a)
	w.Serialize(&b)
	w.Serialize(&a)
	w.Serialize(&b)
	reader := NewReader(w.Bytes(), false)
	var a1 []int
	reader.Unserialize(&a1)
	if !reflect.DeepEqual(a1, a) {
		t.Error(a1, a)
	}
	var b1 []byte
	reader.Unserialize(&b1)
	if !reflect.DeepEqual(b1, b) {
		t.Error(b1, b)
	}
	a2 := []int{}
	reader.Unserialize(&a2)
	if !reflect.DeepEqual(a2, a) {
		t.Error(a2, a)
	}
	b2 := []byte{}
	reader.Unserialize(&b2)
	if !reflect.DeepEqual(b2, b) {
		t.Error(b2, b)
	}
	a2 = make([]int, 10)
	reader.Unserialize(&a2)
	if !reflect.DeepEqual(a2, a) {
		t.Error(a2, a)
	}
	b2 = make([]byte, 10)
	reader.Unserialize(&b2)
	if !reflect.DeepEqual(b2, b) {
		t.Error(b2, b)
	}
	w.Close()
}

func BenchmarkUnserializeByteSlice(b *testing.B) {
	w := NewWriter(true)
	w.Serialize([]byte{'h', 'e', 'l', 'l', 'o'})
	bytes := w.Bytes()
	var p []byte
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
	w.Close()
}

func BenchmarkUnserializeIntSlice(b *testing.B) {
	w := NewWriter(true)
	w.Serialize([]int{1, 2, 3, 4, 5})
	bytes := w.Bytes()
	var p []int
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
	w.Close()
}

func TestUnserializeComplex64(t *testing.T) {
	complex64Value := "3.14159"
	complex64Slice := []float32{math.MaxFloat32, math.MaxFloat32}
	data := map[interface{}]complex64{
		true:                                      1,
		false:                                     0,
		nil:                                       0,
		"":                                        0,
		0:                                         0,
		1:                                         1,
		9:                                         9,
		100:                                       100,
		math.MaxInt64:                             complex(float32(math.MaxInt64), 0),
		math.MaxFloat32:                           complex(math.MaxFloat32, 0),
		0.0:                                       0,
		"1":                                       1,
		"9":                                       9,
		&complex64Value:                           complex(float32(3.14159), 0),
		complex(math.MaxFloat32, math.MaxFloat32): complex(math.MaxFloat32, math.MaxFloat32),
		&complex64Slice:                           complex(math.MaxFloat32, math.MaxFloat32),
	}
	w := NewWriter(false)
	keys := []interface{}{}
	for k := range data {
		w.Serialize(k)
		keys = append(keys, k)
	}
	w.Serialize(&complex64Value)
	w.Serialize(&complex64Slice)
	reader := NewReader(w.Bytes(), false)
	var p complex64
	for _, k := range keys {
		reader.Unserialize(&p)
		if p != data[k] {
			t.Error(k, data[k], p)
		}
	}
	reader.Unserialize(&p)
	if p != complex64(3.14159) {
		t.Error(complex64Value, 3.14159, p)
	}
	reader.Unserialize(&p)
	if p != complex(math.MaxFloat32, math.MaxFloat32) {
		t.Error(complex64Value, complex(math.MaxFloat32, math.MaxFloat32), p)
	}
	w.Close()
}

func BenchmarkReadComplex64(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(complex(math.MaxFloat32, math.MaxFloat32))
	bytes := w.Bytes()
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.ReadComplex64()
	}
	w.Close()
}

func BenchmarkUnserializeComplex64(b *testing.B) {
	w := NewWriter(true)
	w.Serialize(complex(math.MaxFloat32, math.MaxFloat32))
	bytes := w.Bytes()
	var p complex64
	for i := 0; i < b.N; i++ {
		reader := NewReader(bytes, true)
		reader.Unserialize(&p)
	}
	w.Close()
}