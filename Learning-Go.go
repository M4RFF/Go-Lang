package main

// Type - Logic
var a bool 
b := true // so, 'b' is is highlighted now, i can't understand why 
c := false
// d := True (error)
// e := FALSE (error)
println(true, false)

// Type - Numbers
// #integers
a := int8(42)
var b int64
c := 7 // int
// b = a (error)
// b = int64(a)

// #floats
d := float32(42)
var e = float64(0.5)
// e * d  (error)
f := 4.2 // float64

// strings
// #EXAMPLE 1#
var s string
a := "hel \n lo" 
b := `hel \n lo`
c := "hel \\n lo"
println(a)
println(b)
println(C)

// #EXAMPLE 2#
//a := "hel
//lo"   (error)
b := `hel
lo`
println(b)

// #EXAMPLE 3#
a := "Hello!"
println(len(a), a)

// #EXAMPLE 4#  (I dont have to do it like this)
func привет(имя string){
	println("Привет, ", имя)
}
func main(){
	имя := "Леша"
	привет(имя)
}

// #EXAMPLE 5# 
s := "Hello"
a := " World"
fmt.Println(s + "\n"+ a) // ' .Println() 'accepts strings as arguments and concatenates them without any spacing


// Type - Symbol 

a := 'F'
a := 'Щ'
fmt.Println("%T %T", 'F', 'Щ')
// %T - "output TYPE value" (int32 int32 string string)

// #EXAMPLE 1# 
a := int32(40)
b := rune(2) // rune == int32
println(a+b)