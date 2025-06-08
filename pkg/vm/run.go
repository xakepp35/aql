package vm

import (
	"github.com/xakepp35/aql/pkg/aqc"
	"github.com/xakepp35/aql/pkg/util"
)

func Run(src []byte, input any) (any, error) {
	e := NewProgrammer()
	err := aqc.Compile(src, e)
	if err != nil {
		return nil, err
	}
	this := NewVM()
	if input != nil {
		this.Push(input)
	}
	this.Run(e.)
	if err := this.Err(); err != nil {
		return nil, util.EWrap(ErrRuntime, err, this.PC(), e.Program()[this.PC()])
	}
	return this.Pop(), nil
}
