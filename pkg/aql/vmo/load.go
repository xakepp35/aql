package vmo

//go:inline
func LdI64(this *VM) {
	this.Stack = append(this.Stack, this.I64())
}

//go:inline
func LdF64(this *VM) {
	this.Stack = append(this.Stack, this.F64())
}

//go:inline
func LdString(this *VM) {
	this.Stack = append(this.Stack, this.String())
}
