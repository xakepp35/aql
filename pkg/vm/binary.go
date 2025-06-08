package vm

import (
	"encoding/binary"
	"fmt"

	"github.com/xakepp35/aql/pkg/vm/op"
	"github.com/xakepp35/aql/pkg/vmi"
)

// [32:N]   Payload:
//     [ops]   uint8 * N               // длина = nOps
//     [args]  int64 * nArgs           // 8 байт на аргумент
//     [types] uint8 * nArgs           // типы
//     [pad16] 0..15                   // если нужно, для выравнивания
//     [data]  []byte                  // всё остальное

func (e *Programmer) MarshalBinary() ([]byte, error) {
	nOps := len(e.prog)
	nArgs := len(e.args)

	opsSize := nOps
	argsSize := nArgs * 8
	typesSize := nArgs
	headerSize := 32

	offset := headerSize + opsSize + argsSize + typesSize

	totalSize := uint64(offset + len(e.data))
	buf := make([]byte, 0, totalSize)

	// --- Header ---
	buf = binary.LittleEndian.AppendUint64(buf, totalSize)
	buf = binary.LittleEndian.AppendUint32(buf, uint32(nOps))
	buf = binary.LittleEndian.AppendUint32(buf, uint32(nArgs))
	buf = append(buf, make([]byte, 16)...)

	// --- Payload ---
	for _, op := range e.prog {
		buf = append(buf, byte(op))
	}
	for _, arg := range e.args {
		buf = binary.LittleEndian.AppendUint64(buf, uint64(arg))
	}
	for _, t := range e.types {
		buf = append(buf, byte(t))
	}
	buf = append(buf, e.data...)

	return buf, nil
}

func (e *Programmer) UnmarshalBinary(data []byte) error {
	if len(data) < 32 {
		return fmt.Errorf("invalid program format")
	}
	total := binary.LittleEndian.Uint64(data[0:8])
	nOps := binary.LittleEndian.Uint32(data[8:12])
	nArgs := binary.LittleEndian.Uint32(data[12:16])

	offset := 32
	if int(total) > len(data) {
		return fmt.Errorf("unexpected total size")
	}

	e.prog = make([]op.Code, nOps)
	for i := range e.prog {
		e.prog[i] = op.Code(data[offset])
		offset++
	}

	e.args = make([]int64, nArgs)
	for i := range e.args {
		e.args[i] = int64(binary.LittleEndian.Uint64(data[offset : offset+8]))
		offset += 8
	}

	e.types = make([]vmi.Type, nArgs)
	for i := range e.types {
		e.types[i] = vmi.Type(data[offset])
		offset++
	}

	// padding := (16 - (offset % 16)) % 16
	// offset += padding

	e.data = make([]byte, len(data)-offset)
	copy(e.data, data[offset:])

	return nil
}
