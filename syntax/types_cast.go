package syntax

func ToIdent(in Value) (Ident_, bool) {
	val, ok := in.(Ident_)
	return val, ok
}

func ToArray(in Value) (Array, bool) {
	val, ok := in.(Array)
	return val, ok
}

func ToString(in Value) (StringLiteral, bool) {
	val, ok := in.(StringLiteral)
	return val, ok
}

func ToInt(in Value) (IntLiteral, bool) {
	val, ok := in.(IntLiteral)
	return val, ok
}

func ToFloat(in Value) (FloatLiteral, bool) {
	val, ok := in.(FloatLiteral)
	return val, ok
}

func ToNull(in Value) (Null_, bool) {
	val, ok := in.(Null_)
	return val, ok
}

func ToBool(in Value) (BoolLiteral, bool) {
	val, ok := in.(BoolLiteral)
	return val, ok
}

func ToMap(in Value) (Map, bool) {
	val, ok := in.(Map)
	return val, ok
}
