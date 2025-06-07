package vm

import (
	"github.com/xakepp35/aql/pkg/aqc"
)

func Run(src []byte, input any) (any, error) {
	e := NewCompiler()
	err := aqc.Compile(src, e)
	if err != nil {
		return nil, err
	}
	this := NewState()
	e.Init(this)
	if input != nil {
		this.Push(input)
	}
	this.Run(e.Program())
	if err := this.Err(); err != nil {
		return nil, err
	}
	return this.Pop(), nil
}
