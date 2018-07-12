package algorithms

import (
	"testing"
	"math/rand"
)

//
//func TestSplitter_2digit(t *testing.T) {
//	input := float64(28)
//	expected1 := float64(2)
//	expected2 := float64(8)
//
//	result1, result2 := Splitter(input)
//
//	if result1 != expected1 {
//		t.Errorf("%v != %v", expected1, result1)
//	}
//	if result2 != expected2 {
//		t.Errorf("%v != %v", expected2, result2)
//	}
//}
//
//func TestSplitter_3digit(t *testing.T) {
//	input := float64(287)
//	expected1 := float64(2)
//	expected2 := float64(87)
//
//	result1, result2 := Splitter(input)
//
//	if result1 != expected1 {
//		t.Errorf("%v != %v", expected1, result1)
//	}
//	if result2 != expected2 {
//		t.Errorf("%v != %v", expected2, result2)
//	}
//}
//
//func TestSplitter_4digit(t *testing.T) {
//	input := float64(2874)
//	expected1 := float64(28)
//	expected2 := float64(74)
//
//	result1, result2 := Splitter(input)
//
//	if result1 != expected1 {
//		t.Errorf("%v != %v", expected1, result1)
//	}
//	if result2 != expected2 {
//		t.Errorf("%v != %v", expected2, result2)
//	}
//}
//
//func TestSplitter_5digit(t *testing.T) {
//	input := float64(28754)
//	expected1 := float64(28)
//	expected2 := float64(754)
//
//	result1, result2 := Splitter(input)
//
//	if result1 != expected1 {
//		t.Errorf("%v != %v", expected1, result1)
//	}
//	if result2 != expected2 {
//		t.Errorf("%v != %v", expected2, result2)
//	}
//}
//
//func TestSplitter_6digit(t *testing.T) {
//	input := float64(123456)
//	expected1 := float64(123)
//	expected2 := float64(456)
//
//	result1, result2 := Splitter(input)
//
//	if result1 != expected1 {
//		t.Errorf("%v != %v", expected1, result1)
//	}
//	if result2 != expected2 {
//		t.Errorf("%v != %v", expected2, result2)
//	}
//}
//
//func TestSplitter_7digit(t *testing.T) {
//	input := float64(1234567)
//	expected1 := float64(123)
//	expected2 := float64(4567)
//
//	result1, result2 := Splitter(input)
//
//	if result1 != expected1 {
//		t.Errorf("%v != %v", expected1, result1)
//	}
//	if result2 != expected2 {
//		t.Errorf("%v != %v", expected2, result2)
//	}
//}
//
//func TestSplitter_8digit(t *testing.T) {
//	input := float64(12345678)
//	expected1 := float64(1234)
//	expected2 := float64(5678)
//
//	result1, result2 := Splitter(input)
//
//	if result1 != expected1 {
//		t.Errorf("%v != %v", expected1, result1)
//	}
//	if result2 != expected2 {
//		t.Errorf("%v != %v", expected2, result2)
//	}
//}
//
//func TestSplitter_9digit(t *testing.T) {
//	input := float64(123456789)
//	expected1 := float64(1234)
//	expected2 := float64(56789)
//
//	result1, result2 := Splitter(input)
//
//	if result1 != expected1 {
//		t.Errorf("%v != %v", expected1, result1)
//	}
//	if result2 != expected2 {
//		t.Errorf("%v != %v", expected2, result2)
//	}
//}
//
//func TestSplitter_10digit(t *testing.T) {
//	input := float64(1234567890)
//	expected1 := float64(12345)
//	expected2 := float64(67890)
//
//	result1, result2 := Splitter(input)
//
//	if result1 != expected1 {
//		t.Errorf("%v != %v", expected1, result1)
//	}
//	if result2 != expected2 {
//		t.Errorf("%v != %v", expected2, result2)
//	}
//}
//
//func TestSplitter_1digit(t *testing.T) {
//	input := float64(5)
//	expected1 := float64(0)
//	expected2 := float64(5)
//
//	result1, result2 := Splitter(input)
//
//	if result1 != expected1 {
//		t.Errorf("%v != %v", expected1, result1)
//	}
//	if result2 != expected2 {
//		t.Errorf("%v != %v", expected2, result2)
//	}
//}
//
//func TestIntMult_two2digit(t *testing.T) {
//	input1 := float64(25)
//	input2 := float64(30)
//	expected := float64(750)
//
//	result := IntMult(input1, input2)
//
//	if result != expected {
//		t.Errorf("%v != %v", expected, result)
//	}
//}
//
//func TestSplitter2_1_1(t *testing.T) {
//	input1 := float64(1)
//	input2 := float64(7)
//	e1 := float64(0)
//	e2 := float64(1)
//	e3 := float64(0)
//	e4 := float64(7)
//
//	r1, r2, r3, r4 := Splitter2(input1, input2)
//
//	if r1 != e1 {
//		t.Errorf("%v != %v", e1, r1)
//	}
//	if r2 != e2 {
//		t.Errorf("%v != %v", e2, r2)
//	}
//	if r3 != e3 {
//		t.Errorf("%v != %v", e3, r3)
//	}
//	if r4 != e4 {
//		t.Errorf("%v != %v", e4, r4)
//	}
//}
//
//func TestSplitter2_2_1(t *testing.T) {
//	input1 := float64(13)
//	input2 := float64(7)
//	e1 := float64(1)
//	e2 := float64(3)
//	e3 := float64(0)
//	e4 := float64(7)
//
//	r1, r2, r3, r4 := Splitter2(input1, input2)
//
//	if r1 != e1 {
//		t.Errorf("%v != %v", e1, r1)
//	}
//	if r2 != e2 {
//		t.Errorf("%v != %v", e2, r2)
//	}
//	if r3 != e3 {
//		t.Errorf("%v != %v", e3, r3)
//	}
//	if r4 != e4 {
//		t.Errorf("%v != %v", e4, r4)
//	}
//}
//
//func TestSplitter2_2_2(t *testing.T) {
//	input1 := float64(13)
//	input2 := float64(78)
//	e1 := float64(1)
//	e2 := float64(3)
//	e3 := float64(7)
//	e4 := float64(8)
//
//	r1, r2, r3, r4 := Splitter2(input1, input2)
//
//	if r1 != e1 {
//		t.Errorf("%v != %v", e1, r1)
//	}
//	if r2 != e2 {
//		t.Errorf("%v != %v", e2, r2)
//	}
//	if r3 != e3 {
//		t.Errorf("%v != %v", e3, r3)
//	}
//	if r4 != e4 {
//		t.Errorf("%v != %v", e4, r4)
//	}
//}
//
//func TestSplitter2_3_2(t *testing.T) {
//	input1 := float64(138)
//	input2 := float64(78)
//	e1 := float64(1)
//	e2 := float64(38)
//	e3 := float64(0)
//	e4 := float64(78)
//
//	r1, r2, r3, r4 := Splitter2(input1, input2)
//
//	if r1 != e1 {
//		t.Errorf("%v != %v", e1, r1)
//	}
//	if r2 != e2 {
//		t.Errorf("%v != %v", e2, r2)
//	}
//	if r3 != e3 {
//		t.Errorf("%v != %v", e3, r3)
//	}
//	if r4 != e4 {
//		t.Errorf("%v != %v", e4, r4)
//	}
//}

//func TestIntMult(t *testing.T) {
//	for i := 0; i < 1000000; i++ {
//		input1 := rand.Intn(4294967296 - 1)
//		input2 := rand.Intn(4294967296 - 1)
//		expected := input2 * input1
//		result := IterativeDoubleHalf(input1, input2)
//
//		if expected != result {
//			t.Fatalf("%v != %v", expected, result)
//		}
//	}
//}

func BenchmarkIntMult(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		//b.StopTimer()
		input1 := rand.Intn(100)
		input2 := rand.Intn(100)
		//b.StartTimer()
		IntMult(input1, input2)
	}
}

func BenchmarkIntNative(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		//b.StopTimer()
		input1 := rand.Intn(100)
		input2 := rand.Intn(100)
		//b.StartTimer()
		Nativemult(input1, input2)
	}
}

func BenchmarkIntBinaryMul(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		//b.StopTimer()
		input1 := rand.Intn(100)
		input2 := rand.Intn(100)
		//b.StartTimer()
		IterativeDoubleHalf(input1, input2)
	}
}

