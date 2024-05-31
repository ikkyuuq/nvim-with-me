package variables

func main() {
	// Declaring var with type
	// var <var_name> <type : optional> = <value>
	var a int = 10
	// int can be int8, int16, int32, int64
	// uint can be uint8, uint16, uint32, uint64
	// byte is an alias for uint8
	// rune is an alias for int32
	var b string = "Hello, World!"
	var c float32 = 3.14
	// float can be float32, float64
	// complex can be complex64, complex128
	var d bool = true

	// Declaring with walrus operator
	// <var_name> := <value>
	e := 10
	f := "Hello, World!"
	g := 3.14
	h := true

	// Declaring multiple variables
	// var <var_name1>, <var_name2> <type>
	var i, j = 10, 20
	// Declaring multiple variables with walrus operator
	// <var_name1>, <var_name2> := <value1>, <value2>
	k, l := 10, 20.5

	// Declaring multiple variables with different types with walrus operator
	// <var_name1>, <var_name2> := <value1>, <value2>
	o, p := "Hello", 10
}
