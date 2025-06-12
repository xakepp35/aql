package vmo

func LdI64(this *VM) {
	this.Stack = append(this.Stack, this.I64())
}

func LdF64(this *VM) {
	this.Stack = append(this.Stack, this.F64())
}

func LdString(this *VM) {
	this.Stack = append(this.Stack, this.String())
}
