package main

import 
(	
	"fmt"
    "reflect"
)


func main() {
	// Integer types
	var i int = -10
	var i8 int8 = -128
	var i16 int16 = -32768
	var i32 int32 = -2147483648
	var i64 int64 = -9223372036854775808

	// Unsigned integer types
	var ui uint = 10
	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615

	// Floating-point types
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793

	// Complex types
	var c64 complex64 = 3 + 4i
	var c128 complex128 = 3 + 4i
	
	// Boolean type
	var b bool = true

	// String type
	var s string = "Hello, World!"
 
	// Displaying the types
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(i8))
	fmt.Println(reflect.TypeOf(i16))
	fmt.Println(reflect.TypeOf(i32))
	fmt.Println(reflect.TypeOf(i64))
	fmt.Println(reflect.TypeOf(ui))
	fmt.Println(reflect.TypeOf(ui8))
	fmt.Println(reflect.TypeOf(ui16))
	fmt.Println(reflect.TypeOf(ui32))
	fmt.Println(reflect.TypeOf(ui64))
	fmt.Println(reflect.TypeOf(f32))
	fmt.Println(reflect.TypeOf(f64))
	fmt.Println(reflect.TypeOf(c64))
	fmt.Println(reflect.TypeOf(c128))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(s))
}