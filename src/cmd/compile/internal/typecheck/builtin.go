// Code generated by mkbuiltin.go. DO NOT EDIT.

package typecheck

import (
	"cmd/compile/internal/types"
	"cmd/internal/src"
)

// Not inlining this function removes a significant chunk of init code.
//
//go:noinline
func newSig(params, results []*types.Field) *types.Type {
	return types.NewSignature(nil, params, results)
}

func params(tlist ...*types.Type) []*types.Field {
	flist := make([]*types.Field, len(tlist))
	for i, typ := range tlist {
		flist[i] = types.NewField(src.NoXPos, nil, typ)
	}
	return flist
}

var runtimeDecls = [...]struct {
	name string
	tag  int
	typ  int
}{
	{"newobject", funcTag, 4},
	{"mallocgc", funcTag, 8},
	{"panicdivide", funcTag, 9},
	{"panicshift", funcTag, 9},
	{"panicmakeslicelen", funcTag, 9},
	{"panicmakeslicecap", funcTag, 9},
	{"throwinit", funcTag, 9},
	{"panicwrap", funcTag, 9},
	{"gopanic", funcTag, 11},
	{"gorecover", funcTag, 14},
	{"goschedguarded", funcTag, 9},
	{"goPanicIndex", funcTag, 16},
	{"goPanicIndexU", funcTag, 18},
	{"goPanicSliceAlen", funcTag, 16},
	{"goPanicSliceAlenU", funcTag, 18},
	{"goPanicSliceAcap", funcTag, 16},
	{"goPanicSliceAcapU", funcTag, 18},
	{"goPanicSliceB", funcTag, 16},
	{"goPanicSliceBU", funcTag, 18},
	{"goPanicSlice3Alen", funcTag, 16},
	{"goPanicSlice3AlenU", funcTag, 18},
	{"goPanicSlice3Acap", funcTag, 16},
	{"goPanicSlice3AcapU", funcTag, 18},
	{"goPanicSlice3B", funcTag, 16},
	{"goPanicSlice3BU", funcTag, 18},
	{"goPanicSlice3C", funcTag, 16},
	{"goPanicSlice3CU", funcTag, 18},
	{"goPanicSliceConvert", funcTag, 16},
	{"printbool", funcTag, 19},
	{"printfloat", funcTag, 21},
	{"printint", funcTag, 23},
	{"printhex", funcTag, 25},
	{"printuint", funcTag, 25},
	{"printcomplex", funcTag, 27},
	{"printstring", funcTag, 29},
	{"printpointer", funcTag, 30},
	{"printuintptr", funcTag, 31},
	{"printiface", funcTag, 30},
	{"printeface", funcTag, 30},
	{"printslice", funcTag, 30},
	{"printnl", funcTag, 9},
	{"printsp", funcTag, 9},
	{"printlock", funcTag, 9},
	{"printunlock", funcTag, 9},
	{"concatstring2", funcTag, 34},
	{"concatstring3", funcTag, 35},
	{"concatstring4", funcTag, 36},
	{"concatstring5", funcTag, 37},
	{"concatstrings", funcTag, 39},
	{"cmpstring", funcTag, 40},
	{"intstring", funcTag, 43},
	{"slicebytetostring", funcTag, 44},
	{"slicebytetostringtmp", funcTag, 45},
	{"slicerunetostring", funcTag, 48},
	{"stringtoslicebyte", funcTag, 50},
	{"stringtoslicerune", funcTag, 53},
	{"slicecopy", funcTag, 54},
	{"decoderune", funcTag, 55},
	{"countrunes", funcTag, 56},
	{"convI2I", funcTag, 58},
	{"convT", funcTag, 59},
	{"convTnoptr", funcTag, 59},
	{"convT16", funcTag, 61},
	{"convT32", funcTag, 63},
	{"convT64", funcTag, 64},
	{"convTstring", funcTag, 65},
	{"convTslice", funcTag, 68},
	{"assertE2I", funcTag, 69},
	{"assertE2I2", funcTag, 70},
	{"assertI2I", funcTag, 69},
	{"assertI2I2", funcTag, 70},
	{"panicdottypeE", funcTag, 71},
	{"panicdottypeI", funcTag, 71},
	{"panicnildottype", funcTag, 72},
	{"interfaceSwitch", funcTag, 73},
	{"ifaceeq", funcTag, 74},
	{"efaceeq", funcTag, 74},
	{"deferrangefunc", funcTag, 75},
	{"fastrand", funcTag, 76},
	{"makemap64", funcTag, 78},
	{"makemap", funcTag, 79},
	{"makemap_small", funcTag, 80},
	{"mapaccess1", funcTag, 81},
	{"mapaccess1_fast32", funcTag, 82},
	{"mapaccess1_fast64", funcTag, 83},
	{"mapaccess1_faststr", funcTag, 84},
	{"mapaccess1_fat", funcTag, 85},
	{"mapaccess2", funcTag, 86},
	{"mapaccess2_fast32", funcTag, 87},
	{"mapaccess2_fast64", funcTag, 88},
	{"mapaccess2_faststr", funcTag, 89},
	{"mapaccess2_fat", funcTag, 90},
	{"mapassign", funcTag, 81},
	{"mapassign_fast32", funcTag, 82},
	{"mapassign_fast32ptr", funcTag, 91},
	{"mapassign_fast64", funcTag, 83},
	{"mapassign_fast64ptr", funcTag, 91},
	{"mapassign_faststr", funcTag, 84},
	{"mapiterinit", funcTag, 92},
	{"mapdelete", funcTag, 92},
	{"mapdelete_fast32", funcTag, 93},
	{"mapdelete_fast64", funcTag, 94},
	{"mapdelete_faststr", funcTag, 95},
	{"mapiternext", funcTag, 96},
	{"mapclear", funcTag, 97},
	{"makechan64", funcTag, 99},
	{"makechan", funcTag, 100},
	{"chanrecv1", funcTag, 102},
	{"chanrecv2", funcTag, 103},
	{"chansend1", funcTag, 105},
	{"closechan", funcTag, 30},
	{"writeBarrier", varTag, 107},
	{"typedmemmove", funcTag, 108},
	{"typedmemclr", funcTag, 109},
	{"typedslicecopy", funcTag, 110},
	{"selectnbsend", funcTag, 111},
	{"selectnbrecv", funcTag, 112},
	{"selectsetpc", funcTag, 113},
	{"selectgo", funcTag, 114},
	{"block", funcTag, 9},
	{"makeslice", funcTag, 115},
	{"makeslice64", funcTag, 116},
	{"makeslicecopy", funcTag, 117},
	{"growslice", funcTag, 119},
	{"unsafeslicecheckptr", funcTag, 120},
	{"panicunsafeslicelen", funcTag, 9},
	{"panicunsafeslicenilptr", funcTag, 9},
	{"unsafestringcheckptr", funcTag, 121},
	{"panicunsafestringlen", funcTag, 9},
	{"panicunsafestringnilptr", funcTag, 9},
	{"memmove", funcTag, 122},
	{"memclrNoHeapPointers", funcTag, 123},
	{"memclrHasPointers", funcTag, 123},
	{"memequal", funcTag, 124},
	{"memequal0", funcTag, 125},
	{"memequal8", funcTag, 125},
	{"memequal16", funcTag, 125},
	{"memequal32", funcTag, 125},
	{"memequal64", funcTag, 125},
	{"memequal128", funcTag, 125},
	{"f32equal", funcTag, 126},
	{"f64equal", funcTag, 126},
	{"c64equal", funcTag, 126},
	{"c128equal", funcTag, 126},
	{"strequal", funcTag, 126},
	{"interequal", funcTag, 126},
	{"nilinterequal", funcTag, 126},
	{"memhash", funcTag, 127},
	{"memhash0", funcTag, 128},
	{"memhash8", funcTag, 128},
	{"memhash16", funcTag, 128},
	{"memhash32", funcTag, 128},
	{"memhash64", funcTag, 128},
	{"memhash128", funcTag, 128},
	{"f32hash", funcTag, 129},
	{"f64hash", funcTag, 129},
	{"c64hash", funcTag, 129},
	{"c128hash", funcTag, 129},
	{"strhash", funcTag, 129},
	{"interhash", funcTag, 129},
	{"nilinterhash", funcTag, 129},
	{"int64div", funcTag, 130},
	{"uint64div", funcTag, 131},
	{"int64mod", funcTag, 130},
	{"uint64mod", funcTag, 131},
	{"float64toint64", funcTag, 132},
	{"float64touint64", funcTag, 133},
	{"float64touint32", funcTag, 134},
	{"int64tofloat64", funcTag, 135},
	{"int64tofloat32", funcTag, 137},
	{"uint64tofloat64", funcTag, 138},
	{"uint64tofloat32", funcTag, 139},
	{"uint32tofloat64", funcTag, 140},
	{"complex128div", funcTag, 141},
	{"getcallerpc", funcTag, 142},
	{"getcallersp", funcTag, 142},
	{"racefuncenter", funcTag, 31},
	{"racefuncexit", funcTag, 9},
	{"raceread", funcTag, 31},
	{"racewrite", funcTag, 31},
	{"racereadrange", funcTag, 143},
	{"racewriterange", funcTag, 143},
	{"msanread", funcTag, 143},
	{"msanwrite", funcTag, 143},
	{"msanmove", funcTag, 144},
	{"asanread", funcTag, 143},
	{"asanwrite", funcTag, 143},
	{"checkptrAlignment", funcTag, 145},
	{"checkptrArithmetic", funcTag, 147},
	{"libfuzzerTraceCmp1", funcTag, 148},
	{"libfuzzerTraceCmp2", funcTag, 149},
	{"libfuzzerTraceCmp4", funcTag, 150},
	{"libfuzzerTraceCmp8", funcTag, 151},
	{"libfuzzerTraceConstCmp1", funcTag, 148},
	{"libfuzzerTraceConstCmp2", funcTag, 149},
	{"libfuzzerTraceConstCmp4", funcTag, 150},
	{"libfuzzerTraceConstCmp8", funcTag, 151},
	{"libfuzzerHookStrCmp", funcTag, 152},
	{"libfuzzerHookEqualFold", funcTag, 152},
	{"addCovMeta", funcTag, 154},
	{"x86HasPOPCNT", varTag, 6},
	{"x86HasSSE41", varTag, 6},
	{"x86HasFMA", varTag, 6},
	{"armHasVFPv4", varTag, 6},
	{"arm64HasATOMICS", varTag, 6},
	{"asanregisterglobals", funcTag, 123},
}

func runtimeTypes() []*types.Type {
	var typs [155]*types.Type
	typs[0] = types.ByteType
	typs[1] = types.NewPtr(typs[0])
	typs[2] = types.Types[types.TANY]
	typs[3] = types.NewPtr(typs[2])
	typs[4] = newSig(params(typs[1]), params(typs[3]))
	typs[5] = types.Types[types.TUINTPTR]
	typs[6] = types.Types[types.TBOOL]
	typs[7] = types.Types[types.TUNSAFEPTR]
	typs[8] = newSig(params(typs[5], typs[1], typs[6]), params(typs[7]))
	typs[9] = newSig(nil, nil)
	typs[10] = types.Types[types.TINTER]
	typs[11] = newSig(params(typs[10]), nil)
	typs[12] = types.Types[types.TINT32]
	typs[13] = types.NewPtr(typs[12])
	typs[14] = newSig(params(typs[13]), params(typs[10]))
	typs[15] = types.Types[types.TINT]
	typs[16] = newSig(params(typs[15], typs[15]), nil)
	typs[17] = types.Types[types.TUINT]
	typs[18] = newSig(params(typs[17], typs[15]), nil)
	typs[19] = newSig(params(typs[6]), nil)
	typs[20] = types.Types[types.TFLOAT64]
	typs[21] = newSig(params(typs[20]), nil)
	typs[22] = types.Types[types.TINT64]
	typs[23] = newSig(params(typs[22]), nil)
	typs[24] = types.Types[types.TUINT64]
	typs[25] = newSig(params(typs[24]), nil)
	typs[26] = types.Types[types.TCOMPLEX128]
	typs[27] = newSig(params(typs[26]), nil)
	typs[28] = types.Types[types.TSTRING]
	typs[29] = newSig(params(typs[28]), nil)
	typs[30] = newSig(params(typs[2]), nil)
	typs[31] = newSig(params(typs[5]), nil)
	typs[32] = types.NewArray(typs[0], 32)
	typs[33] = types.NewPtr(typs[32])
	typs[34] = newSig(params(typs[33], typs[28], typs[28]), params(typs[28]))
	typs[35] = newSig(params(typs[33], typs[28], typs[28], typs[28]), params(typs[28]))
	typs[36] = newSig(params(typs[33], typs[28], typs[28], typs[28], typs[28]), params(typs[28]))
	typs[37] = newSig(params(typs[33], typs[28], typs[28], typs[28], typs[28], typs[28]), params(typs[28]))
	typs[38] = types.NewSlice(typs[28])
	typs[39] = newSig(params(typs[33], typs[38]), params(typs[28]))
	typs[40] = newSig(params(typs[28], typs[28]), params(typs[15]))
	typs[41] = types.NewArray(typs[0], 4)
	typs[42] = types.NewPtr(typs[41])
	typs[43] = newSig(params(typs[42], typs[22]), params(typs[28]))
	typs[44] = newSig(params(typs[33], typs[1], typs[15]), params(typs[28]))
	typs[45] = newSig(params(typs[1], typs[15]), params(typs[28]))
	typs[46] = types.RuneType
	typs[47] = types.NewSlice(typs[46])
	typs[48] = newSig(params(typs[33], typs[47]), params(typs[28]))
	typs[49] = types.NewSlice(typs[0])
	typs[50] = newSig(params(typs[33], typs[28]), params(typs[49]))
	typs[51] = types.NewArray(typs[46], 32)
	typs[52] = types.NewPtr(typs[51])
	typs[53] = newSig(params(typs[52], typs[28]), params(typs[47]))
	typs[54] = newSig(params(typs[3], typs[15], typs[3], typs[15], typs[5]), params(typs[15]))
	typs[55] = newSig(params(typs[28], typs[15]), params(typs[46], typs[15]))
	typs[56] = newSig(params(typs[28]), params(typs[15]))
	typs[57] = types.NewPtr(typs[5])
	typs[58] = newSig(params(typs[1], typs[57]), params(typs[57]))
	typs[59] = newSig(params(typs[1], typs[3]), params(typs[7]))
	typs[60] = types.Types[types.TUINT16]
	typs[61] = newSig(params(typs[60]), params(typs[7]))
	typs[62] = types.Types[types.TUINT32]
	typs[63] = newSig(params(typs[62]), params(typs[7]))
	typs[64] = newSig(params(typs[24]), params(typs[7]))
	typs[65] = newSig(params(typs[28]), params(typs[7]))
	typs[66] = types.Types[types.TUINT8]
	typs[67] = types.NewSlice(typs[66])
	typs[68] = newSig(params(typs[67]), params(typs[7]))
	typs[69] = newSig(params(typs[1], typs[1]), params(typs[1]))
	typs[70] = newSig(params(typs[1], typs[2]), params(typs[2]))
	typs[71] = newSig(params(typs[1], typs[1], typs[1]), nil)
	typs[72] = newSig(params(typs[1]), nil)
	typs[73] = newSig(params(typs[1], typs[1]), params(typs[15], typs[1]))
	typs[74] = newSig(params(typs[57], typs[7], typs[7]), params(typs[6]))
	typs[75] = newSig(nil, params(typs[10]))
	typs[76] = newSig(nil, params(typs[62]))
	typs[77] = types.NewMap(typs[2], typs[2])
	typs[78] = newSig(params(typs[1], typs[22], typs[3]), params(typs[77]))
	typs[79] = newSig(params(typs[1], typs[15], typs[3]), params(typs[77]))
	typs[80] = newSig(nil, params(typs[77]))
	typs[81] = newSig(params(typs[1], typs[77], typs[3]), params(typs[3]))
	typs[82] = newSig(params(typs[1], typs[77], typs[62]), params(typs[3]))
	typs[83] = newSig(params(typs[1], typs[77], typs[24]), params(typs[3]))
	typs[84] = newSig(params(typs[1], typs[77], typs[28]), params(typs[3]))
	typs[85] = newSig(params(typs[1], typs[77], typs[3], typs[1]), params(typs[3]))
	typs[86] = newSig(params(typs[1], typs[77], typs[3]), params(typs[3], typs[6]))
	typs[87] = newSig(params(typs[1], typs[77], typs[62]), params(typs[3], typs[6]))
	typs[88] = newSig(params(typs[1], typs[77], typs[24]), params(typs[3], typs[6]))
	typs[89] = newSig(params(typs[1], typs[77], typs[28]), params(typs[3], typs[6]))
	typs[90] = newSig(params(typs[1], typs[77], typs[3], typs[1]), params(typs[3], typs[6]))
	typs[91] = newSig(params(typs[1], typs[77], typs[7]), params(typs[3]))
	typs[92] = newSig(params(typs[1], typs[77], typs[3]), nil)
	typs[93] = newSig(params(typs[1], typs[77], typs[62]), nil)
	typs[94] = newSig(params(typs[1], typs[77], typs[24]), nil)
	typs[95] = newSig(params(typs[1], typs[77], typs[28]), nil)
	typs[96] = newSig(params(typs[3]), nil)
	typs[97] = newSig(params(typs[1], typs[77]), nil)
	typs[98] = types.NewChan(typs[2], types.Cboth)
	typs[99] = newSig(params(typs[1], typs[22]), params(typs[98]))
	typs[100] = newSig(params(typs[1], typs[15]), params(typs[98]))
	typs[101] = types.NewChan(typs[2], types.Crecv)
	typs[102] = newSig(params(typs[101], typs[3]), nil)
	typs[103] = newSig(params(typs[101], typs[3]), params(typs[6]))
	typs[104] = types.NewChan(typs[2], types.Csend)
	typs[105] = newSig(params(typs[104], typs[3]), nil)
	typs[106] = types.NewArray(typs[0], 3)
	typs[107] = types.NewStruct([]*types.Field{types.NewField(src.NoXPos, Lookup("enabled"), typs[6]), types.NewField(src.NoXPos, Lookup("pad"), typs[106]), types.NewField(src.NoXPos, Lookup("needed"), typs[6]), types.NewField(src.NoXPos, Lookup("cgo"), typs[6]), types.NewField(src.NoXPos, Lookup("alignme"), typs[24])})
	typs[108] = newSig(params(typs[1], typs[3], typs[3]), nil)
	typs[109] = newSig(params(typs[1], typs[3]), nil)
	typs[110] = newSig(params(typs[1], typs[3], typs[15], typs[3], typs[15]), params(typs[15]))
	typs[111] = newSig(params(typs[104], typs[3]), params(typs[6]))
	typs[112] = newSig(params(typs[3], typs[101]), params(typs[6], typs[6]))
	typs[113] = newSig(params(typs[57]), nil)
	typs[114] = newSig(params(typs[1], typs[1], typs[57], typs[15], typs[15], typs[6]), params(typs[15], typs[6]))
	typs[115] = newSig(params(typs[1], typs[15], typs[15]), params(typs[7]))
	typs[116] = newSig(params(typs[1], typs[22], typs[22]), params(typs[7]))
	typs[117] = newSig(params(typs[1], typs[15], typs[15], typs[7]), params(typs[7]))
	typs[118] = types.NewSlice(typs[2])
	typs[119] = newSig(params(typs[3], typs[15], typs[15], typs[15], typs[1]), params(typs[118]))
	typs[120] = newSig(params(typs[1], typs[7], typs[22]), nil)
	typs[121] = newSig(params(typs[7], typs[22]), nil)
	typs[122] = newSig(params(typs[3], typs[3], typs[5]), nil)
	typs[123] = newSig(params(typs[7], typs[5]), nil)
	typs[124] = newSig(params(typs[3], typs[3], typs[5]), params(typs[6]))
	typs[125] = newSig(params(typs[3], typs[3]), params(typs[6]))
	typs[126] = newSig(params(typs[7], typs[7]), params(typs[6]))
	typs[127] = newSig(params(typs[3], typs[5], typs[5]), params(typs[5]))
	typs[128] = newSig(params(typs[7], typs[5]), params(typs[5]))
	typs[129] = newSig(params(typs[3], typs[5]), params(typs[5]))
	typs[130] = newSig(params(typs[22], typs[22]), params(typs[22]))
	typs[131] = newSig(params(typs[24], typs[24]), params(typs[24]))
	typs[132] = newSig(params(typs[20]), params(typs[22]))
	typs[133] = newSig(params(typs[20]), params(typs[24]))
	typs[134] = newSig(params(typs[20]), params(typs[62]))
	typs[135] = newSig(params(typs[22]), params(typs[20]))
	typs[136] = types.Types[types.TFLOAT32]
	typs[137] = newSig(params(typs[22]), params(typs[136]))
	typs[138] = newSig(params(typs[24]), params(typs[20]))
	typs[139] = newSig(params(typs[24]), params(typs[136]))
	typs[140] = newSig(params(typs[62]), params(typs[20]))
	typs[141] = newSig(params(typs[26], typs[26]), params(typs[26]))
	typs[142] = newSig(nil, params(typs[5]))
	typs[143] = newSig(params(typs[5], typs[5]), nil)
	typs[144] = newSig(params(typs[5], typs[5], typs[5]), nil)
	typs[145] = newSig(params(typs[7], typs[1], typs[5]), nil)
	typs[146] = types.NewSlice(typs[7])
	typs[147] = newSig(params(typs[7], typs[146]), nil)
	typs[148] = newSig(params(typs[66], typs[66], typs[17]), nil)
	typs[149] = newSig(params(typs[60], typs[60], typs[17]), nil)
	typs[150] = newSig(params(typs[62], typs[62], typs[17]), nil)
	typs[151] = newSig(params(typs[24], typs[24], typs[17]), nil)
	typs[152] = newSig(params(typs[28], typs[28], typs[17]), nil)
	typs[153] = types.NewArray(typs[0], 16)
	typs[154] = newSig(params(typs[7], typs[62], typs[153], typs[28], typs[15], typs[66], typs[66]), params(typs[62]))
	return typs[:]
}

var coverageDecls = [...]struct {
	name string
	tag  int
	typ  int
}{
	{"initHook", funcTag, 1},
}

func coverageTypes() []*types.Type {
	var typs [2]*types.Type
	typs[0] = types.Types[types.TBOOL]
	typs[1] = newSig(params(typs[0]), nil)
	return typs[:]
}
