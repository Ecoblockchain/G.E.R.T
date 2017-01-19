// Do not edit. Bootstrap copy of /Users/fruit/Documents/biscuit/golang1.7/src/cmd/compile/internal/ssa/rewritedec.go

//line /Users/fruit/Documents/biscuit/golang1.7/src/cmd/compile/internal/ssa/rewritedec.go:1
// autogenerated from gen/dec.rules: do not edit!
// generated with: cd gen; go run *.go

package ssa

import "math"

var _ = math.MinInt8 // in case not otherwise used
func rewriteValuedec(v *Value, config *Config) bool {
	switch v.Op {
	case OpComplexImag:
		return rewriteValuedec_OpComplexImag(v, config)
	case OpComplexReal:
		return rewriteValuedec_OpComplexReal(v, config)
	case OpIData:
		return rewriteValuedec_OpIData(v, config)
	case OpITab:
		return rewriteValuedec_OpITab(v, config)
	case OpLoad:
		return rewriteValuedec_OpLoad(v, config)
	case OpSliceCap:
		return rewriteValuedec_OpSliceCap(v, config)
	case OpSliceLen:
		return rewriteValuedec_OpSliceLen(v, config)
	case OpSlicePtr:
		return rewriteValuedec_OpSlicePtr(v, config)
	case OpStore:
		return rewriteValuedec_OpStore(v, config)
	case OpStringLen:
		return rewriteValuedec_OpStringLen(v, config)
	case OpStringPtr:
		return rewriteValuedec_OpStringPtr(v, config)
	}
	return false
}
func rewriteValuedec_OpComplexImag(v *Value, config *Config) bool {
	b := v.Block
	_ = b
	// match: (ComplexImag (ComplexMake _ imag ))
	// cond:
	// result: imag
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpComplexMake {
			break
		}
		imag := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = imag.Type
		v.AddArg(imag)
		return true
	}
	return false
}
func rewriteValuedec_OpComplexReal(v *Value, config *Config) bool {
	b := v.Block
	_ = b
	// match: (ComplexReal (ComplexMake real _  ))
	// cond:
	// result: real
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpComplexMake {
			break
		}
		real := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = real.Type
		v.AddArg(real)
		return true
	}
	return false
}
func rewriteValuedec_OpIData(v *Value, config *Config) bool {
	b := v.Block
	_ = b
	// match: (IData (IMake _ data))
	// cond:
	// result: data
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpIMake {
			break
		}
		data := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = data.Type
		v.AddArg(data)
		return true
	}
	return false
}
func rewriteValuedec_OpITab(v *Value, config *Config) bool {
	b := v.Block
	_ = b
	// match: (ITab (IMake itab _))
	// cond:
	// result: itab
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpIMake {
			break
		}
		itab := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = itab.Type
		v.AddArg(itab)
		return true
	}
	return false
}
func rewriteValuedec_OpLoad(v *Value, config *Config) bool {
	b := v.Block
	_ = b
	// match: (Load <t> ptr mem)
	// cond: t.IsComplex() && t.Size() == 8
	// result: (ComplexMake     (Load <config.fe.TypeFloat32()> ptr mem)     (Load <config.fe.TypeFloat32()>       (OffPtr <config.fe.TypeFloat32().PtrTo()> [4] ptr)       mem)     )
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsComplex() && t.Size() == 8) {
			break
		}
		v.reset(OpComplexMake)
		v0 := b.NewValue0(v.Line, OpLoad, config.fe.TypeFloat32())
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Line, OpLoad, config.fe.TypeFloat32())
		v2 := b.NewValue0(v.Line, OpOffPtr, config.fe.TypeFloat32().PtrTo())
		v2.AuxInt = 4
		v2.AddArg(ptr)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.IsComplex() && t.Size() == 16
	// result: (ComplexMake     (Load <config.fe.TypeFloat64()> ptr mem)     (Load <config.fe.TypeFloat64()>       (OffPtr <config.fe.TypeFloat64().PtrTo()> [8] ptr)       mem)     )
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsComplex() && t.Size() == 16) {
			break
		}
		v.reset(OpComplexMake)
		v0 := b.NewValue0(v.Line, OpLoad, config.fe.TypeFloat64())
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Line, OpLoad, config.fe.TypeFloat64())
		v2 := b.NewValue0(v.Line, OpOffPtr, config.fe.TypeFloat64().PtrTo())
		v2.AuxInt = 8
		v2.AddArg(ptr)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.IsString()
	// result: (StringMake     (Load <config.fe.TypeBytePtr()> ptr mem)     (Load <config.fe.TypeInt()>       (OffPtr <config.fe.TypeInt().PtrTo()> [config.PtrSize] ptr)       mem))
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsString()) {
			break
		}
		v.reset(OpStringMake)
		v0 := b.NewValue0(v.Line, OpLoad, config.fe.TypeBytePtr())
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Line, OpLoad, config.fe.TypeInt())
		v2 := b.NewValue0(v.Line, OpOffPtr, config.fe.TypeInt().PtrTo())
		v2.AuxInt = config.PtrSize
		v2.AddArg(ptr)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.IsSlice()
	// result: (SliceMake     (Load <t.ElemType().PtrTo()> ptr mem)     (Load <config.fe.TypeInt()>       (OffPtr <config.fe.TypeInt().PtrTo()> [config.PtrSize] ptr)       mem)     (Load <config.fe.TypeInt()>       (OffPtr <config.fe.TypeInt().PtrTo()> [2*config.PtrSize] ptr)       mem))
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsSlice()) {
			break
		}
		v.reset(OpSliceMake)
		v0 := b.NewValue0(v.Line, OpLoad, t.ElemType().PtrTo())
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Line, OpLoad, config.fe.TypeInt())
		v2 := b.NewValue0(v.Line, OpOffPtr, config.fe.TypeInt().PtrTo())
		v2.AuxInt = config.PtrSize
		v2.AddArg(ptr)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		v3 := b.NewValue0(v.Line, OpLoad, config.fe.TypeInt())
		v4 := b.NewValue0(v.Line, OpOffPtr, config.fe.TypeInt().PtrTo())
		v4.AuxInt = 2 * config.PtrSize
		v4.AddArg(ptr)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v.AddArg(v3)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.IsInterface()
	// result: (IMake     (Load <config.fe.TypeBytePtr()> ptr mem)     (Load <config.fe.TypeBytePtr()>       (OffPtr <config.fe.TypeBytePtr().PtrTo()> [config.PtrSize] ptr)       mem))
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsInterface()) {
			break
		}
		v.reset(OpIMake)
		v0 := b.NewValue0(v.Line, OpLoad, config.fe.TypeBytePtr())
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Line, OpLoad, config.fe.TypeBytePtr())
		v2 := b.NewValue0(v.Line, OpOffPtr, config.fe.TypeBytePtr().PtrTo())
		v2.AuxInt = config.PtrSize
		v2.AddArg(ptr)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuedec_OpSliceCap(v *Value, config *Config) bool {
	b := v.Block
	_ = b
	// match: (SliceCap (SliceMake _ _ cap))
	// cond:
	// result: cap
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSliceMake {
			break
		}
		cap := v_0.Args[2]
		v.reset(OpCopy)
		v.Type = cap.Type
		v.AddArg(cap)
		return true
	}
	return false
}
func rewriteValuedec_OpSliceLen(v *Value, config *Config) bool {
	b := v.Block
	_ = b
	// match: (SliceLen (SliceMake _ len _))
	// cond:
	// result: len
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSliceMake {
			break
		}
		len := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = len.Type
		v.AddArg(len)
		return true
	}
	return false
}
func rewriteValuedec_OpSlicePtr(v *Value, config *Config) bool {
	b := v.Block
	_ = b
	// match: (SlicePtr (SliceMake ptr _ _ ))
	// cond:
	// result: ptr
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSliceMake {
			break
		}
		ptr := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = ptr.Type
		v.AddArg(ptr)
		return true
	}
	return false
}
func rewriteValuedec_OpStore(v *Value, config *Config) bool {
	b := v.Block
	_ = b
	// match: (Store [8] dst (ComplexMake real imag) mem)
	// cond:
	// result: (Store [4]     (OffPtr <config.fe.TypeFloat32().PtrTo()> [4] dst)     imag     (Store [4] dst real mem))
	for {
		if v.AuxInt != 8 {
			break
		}
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpComplexMake {
			break
		}
		real := v_1.Args[0]
		imag := v_1.Args[1]
		mem := v.Args[2]
		v.reset(OpStore)
		v.AuxInt = 4
		v0 := b.NewValue0(v.Line, OpOffPtr, config.fe.TypeFloat32().PtrTo())
		v0.AuxInt = 4
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(imag)
		v1 := b.NewValue0(v.Line, OpStore, TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v1.AddArg(real)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Store [16] dst (ComplexMake real imag) mem)
	// cond:
	// result: (Store [8]     (OffPtr <config.fe.TypeFloat64().PtrTo()> [8] dst)     imag     (Store [8] dst real mem))
	for {
		if v.AuxInt != 16 {
			break
		}
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpComplexMake {
			break
		}
		real := v_1.Args[0]
		imag := v_1.Args[1]
		mem := v.Args[2]
		v.reset(OpStore)
		v.AuxInt = 8
		v0 := b.NewValue0(v.Line, OpOffPtr, config.fe.TypeFloat64().PtrTo())
		v0.AuxInt = 8
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(imag)
		v1 := b.NewValue0(v.Line, OpStore, TypeMem)
		v1.AuxInt = 8
		v1.AddArg(dst)
		v1.AddArg(real)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Store [2*config.PtrSize] dst (StringMake ptr len) mem)
	// cond:
	// result: (Store [config.PtrSize]     (OffPtr <config.fe.TypeInt().PtrTo()> [config.PtrSize] dst)     len     (Store [config.PtrSize] dst ptr mem))
	for {
		if v.AuxInt != 2*config.PtrSize {
			break
		}
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStringMake {
			break
		}
		ptr := v_1.Args[0]
		len := v_1.Args[1]
		mem := v.Args[2]
		v.reset(OpStore)
		v.AuxInt = config.PtrSize
		v0 := b.NewValue0(v.Line, OpOffPtr, config.fe.TypeInt().PtrTo())
		v0.AuxInt = config.PtrSize
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(len)
		v1 := b.NewValue0(v.Line, OpStore, TypeMem)
		v1.AuxInt = config.PtrSize
		v1.AddArg(dst)
		v1.AddArg(ptr)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Store [3*config.PtrSize] dst (SliceMake ptr len cap) mem)
	// cond:
	// result: (Store [config.PtrSize]     (OffPtr <config.fe.TypeInt().PtrTo()> [2*config.PtrSize] dst)     cap     (Store [config.PtrSize]       (OffPtr <config.fe.TypeInt().PtrTo()> [config.PtrSize] dst)       len       (Store [config.PtrSize] dst ptr mem)))
	for {
		if v.AuxInt != 3*config.PtrSize {
			break
		}
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSliceMake {
			break
		}
		ptr := v_1.Args[0]
		len := v_1.Args[1]
		cap := v_1.Args[2]
		mem := v.Args[2]
		v.reset(OpStore)
		v.AuxInt = config.PtrSize
		v0 := b.NewValue0(v.Line, OpOffPtr, config.fe.TypeInt().PtrTo())
		v0.AuxInt = 2 * config.PtrSize
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(cap)
		v1 := b.NewValue0(v.Line, OpStore, TypeMem)
		v1.AuxInt = config.PtrSize
		v2 := b.NewValue0(v.Line, OpOffPtr, config.fe.TypeInt().PtrTo())
		v2.AuxInt = config.PtrSize
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(len)
		v3 := b.NewValue0(v.Line, OpStore, TypeMem)
		v3.AuxInt = config.PtrSize
		v3.AddArg(dst)
		v3.AddArg(ptr)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Store [2*config.PtrSize] dst (IMake itab data) mem)
	// cond:
	// result: (Store [config.PtrSize]     (OffPtr <config.fe.TypeBytePtr().PtrTo()> [config.PtrSize] dst)     data     (Store [config.PtrSize] dst itab mem))
	for {
		if v.AuxInt != 2*config.PtrSize {
			break
		}
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpIMake {
			break
		}
		itab := v_1.Args[0]
		data := v_1.Args[1]
		mem := v.Args[2]
		v.reset(OpStore)
		v.AuxInt = config.PtrSize
		v0 := b.NewValue0(v.Line, OpOffPtr, config.fe.TypeBytePtr().PtrTo())
		v0.AuxInt = config.PtrSize
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(data)
		v1 := b.NewValue0(v.Line, OpStore, TypeMem)
		v1.AuxInt = config.PtrSize
		v1.AddArg(dst)
		v1.AddArg(itab)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuedec_OpStringLen(v *Value, config *Config) bool {
	b := v.Block
	_ = b
	// match: (StringLen (StringMake _ len))
	// cond:
	// result: len
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpStringMake {
			break
		}
		len := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = len.Type
		v.AddArg(len)
		return true
	}
	return false
}
func rewriteValuedec_OpStringPtr(v *Value, config *Config) bool {
	b := v.Block
	_ = b
	// match: (StringPtr (StringMake ptr _))
	// cond:
	// result: ptr
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpStringMake {
			break
		}
		ptr := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = ptr.Type
		v.AddArg(ptr)
		return true
	}
	return false
}
func rewriteBlockdec(b *Block, config *Config) bool {
	switch b.Kind {
	}
	return false
}
