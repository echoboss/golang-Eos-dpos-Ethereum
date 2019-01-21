
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//版权所有2016 Go Ethereum作者
//此文件是Go以太坊库的一部分。
//
//Go-Ethereum库是免费软件：您可以重新分发它和/或修改
//根据GNU发布的较低通用公共许可证的条款
//自由软件基金会，或者许可证的第3版，或者
//（由您选择）任何更高版本。
//
//Go以太坊图书馆的发行目的是希望它会有用，
//但没有任何保证；甚至没有
//适销性或特定用途的适用性。见
//GNU较低的通用公共许可证，了解更多详细信息。
//
//你应该收到一份GNU较低级别的公共许可证副本
//以及Go以太坊图书馆。如果没有，请参见<http://www.gnu.org/licenses/>。

package abi

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
)

//typewithoutStringer是类型类型类型的别名，该类型只是不实现
//Stringer接口允许在下面的测试中打印类型详细信息。
type typeWithoutStringer Type

//测试类型分析器识别所有允许的类型。
func TestTypeRegexp(t *testing.T) {
	tests := []struct {
		blob string
		kind Type
	}{
		{"bool", Type{Kind: reflect.Bool, T: BoolTy, Type: reflect.TypeOf(bool(false)), stringKind: "bool"}},
		{"bool[]", Type{Kind: reflect.Slice, T: SliceTy, Type: reflect.TypeOf([]bool(nil)), Elem: &Type{Kind: reflect.Bool, T: BoolTy, Type: reflect.TypeOf(bool(false)), stringKind: "bool"}, stringKind: "bool[]"}},
		{"bool[2]", Type{Size: 2, Kind: reflect.Array, T: ArrayTy, Type: reflect.TypeOf([2]bool{}), Elem: &Type{Kind: reflect.Bool, T: BoolTy, Type: reflect.TypeOf(bool(false)), stringKind: "bool"}, stringKind: "bool[2]"}},
		{"bool[2][]", Type{Kind: reflect.Slice, T: SliceTy, Type: reflect.TypeOf([][2]bool{}), Elem: &Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2]bool{}), Elem: &Type{Kind: reflect.Bool, T: BoolTy, Type: reflect.TypeOf(bool(false)), stringKind: "bool"}, stringKind: "bool[2]"}, stringKind: "bool[2][]"}},
		{"bool[][]", Type{Kind: reflect.Slice, T: SliceTy, Type: reflect.TypeOf([][]bool{}), Elem: &Type{Kind: reflect.Slice, T: SliceTy, Type: reflect.TypeOf([]bool{}), Elem: &Type{Kind: reflect.Bool, T: BoolTy, Type: reflect.TypeOf(bool(false)), stringKind: "bool"}, stringKind: "bool[]"}, stringKind: "bool[][]"}},
		{"bool[][2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2][]bool{}), Elem: &Type{Kind: reflect.Slice, T: SliceTy, Type: reflect.TypeOf([]bool{}), Elem: &Type{Kind: reflect.Bool, T: BoolTy, Type: reflect.TypeOf(bool(false)), stringKind: "bool"}, stringKind: "bool[]"}, stringKind: "bool[][2]"}},
		{"bool[2][2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2][2]bool{}), Elem: &Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2]bool{}), Elem: &Type{Kind: reflect.Bool, T: BoolTy, Type: reflect.TypeOf(bool(false)), stringKind: "bool"}, stringKind: "bool[2]"}, stringKind: "bool[2][2]"}},
		{"bool[2][][2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2][][2]bool{}), Elem: &Type{Kind: reflect.Slice, T: SliceTy, Type: reflect.TypeOf([][2]bool{}), Elem: &Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2]bool{}), Elem: &Type{Kind: reflect.Bool, T: BoolTy, Type: reflect.TypeOf(bool(false)), stringKind: "bool"}, stringKind: "bool[2]"}, stringKind: "bool[2][]"}, stringKind: "bool[2][][2]"}},
		{"bool[2][2][2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2][2][2]bool{}), Elem: &Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2][2]bool{}), Elem: &Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2]bool{}), Elem: &Type{Kind: reflect.Bool, T: BoolTy, Type: reflect.TypeOf(bool(false)), stringKind: "bool"}, stringKind: "bool[2]"}, stringKind: "bool[2][2]"}, stringKind: "bool[2][2][2]"}},
		{"bool[][][]", Type{T: SliceTy, Kind: reflect.Slice, Type: reflect.TypeOf([][][]bool{}), Elem: &Type{T: SliceTy, Kind: reflect.Slice, Type: reflect.TypeOf([][]bool{}), Elem: &Type{T: SliceTy, Kind: reflect.Slice, Type: reflect.TypeOf([]bool{}), Elem: &Type{Kind: reflect.Bool, T: BoolTy, Type: reflect.TypeOf(bool(false)), stringKind: "bool"}, stringKind: "bool[]"}, stringKind: "bool[][]"}, stringKind: "bool[][][]"}},
		{"bool[][2][]", Type{T: SliceTy, Kind: reflect.Slice, Type: reflect.TypeOf([][2][]bool{}), Elem: &Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2][]bool{}), Elem: &Type{T: SliceTy, Kind: reflect.Slice, Type: reflect.TypeOf([]bool{}), Elem: &Type{Kind: reflect.Bool, T: BoolTy, Type: reflect.TypeOf(bool(false)), stringKind: "bool"}, stringKind: "bool[]"}, stringKind: "bool[][2]"}, stringKind: "bool[][2][]"}},
		{"int8", Type{Kind: reflect.Int8, Type: int8T, Size: 8, T: IntTy, stringKind: "int8"}},
		{"int16", Type{Kind: reflect.Int16, Type: int16T, Size: 16, T: IntTy, stringKind: "int16"}},
		{"int32", Type{Kind: reflect.Int32, Type: int32T, Size: 32, T: IntTy, stringKind: "int32"}},
		{"int64", Type{Kind: reflect.Int64, Type: int64T, Size: 64, T: IntTy, stringKind: "int64"}},
		{"int256", Type{Kind: reflect.Ptr, Type: bigT, Size: 256, T: IntTy, stringKind: "int256"}},
		{"int8[]", Type{Kind: reflect.Slice, T: SliceTy, Type: reflect.TypeOf([]int8{}), Elem: &Type{Kind: reflect.Int8, Type: int8T, Size: 8, T: IntTy, stringKind: "int8"}, stringKind: "int8[]"}},
		{"int8[2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2]int8{}), Elem: &Type{Kind: reflect.Int8, Type: int8T, Size: 8, T: IntTy, stringKind: "int8"}, stringKind: "int8[2]"}},
		{"int16[]", Type{Kind: reflect.Slice, T: SliceTy, Type: reflect.TypeOf([]int16{}), Elem: &Type{Kind: reflect.Int16, Type: int16T, Size: 16, T: IntTy, stringKind: "int16"}, stringKind: "int16[]"}},
		{"int16[2]", Type{Size: 2, Kind: reflect.Array, T: ArrayTy, Type: reflect.TypeOf([2]int16{}), Elem: &Type{Kind: reflect.Int16, Type: int16T, Size: 16, T: IntTy, stringKind: "int16"}, stringKind: "int16[2]"}},
		{"int32[]", Type{Kind: reflect.Slice, T: SliceTy, Type: reflect.TypeOf([]int32{}), Elem: &Type{Kind: reflect.Int32, Type: int32T, Size: 32, T: IntTy, stringKind: "int32"}, stringKind: "int32[]"}},
		{"int32[2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2]int32{}), Elem: &Type{Kind: reflect.Int32, Type: int32T, Size: 32, T: IntTy, stringKind: "int32"}, stringKind: "int32[2]"}},
		{"int64[]", Type{Kind: reflect.Slice, T: SliceTy, Type: reflect.TypeOf([]int64{}), Elem: &Type{Kind: reflect.Int64, Type: int64T, Size: 64, T: IntTy, stringKind: "int64"}, stringKind: "int64[]"}},
		{"int64[2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2]int64{}), Elem: &Type{Kind: reflect.Int64, Type: int64T, Size: 64, T: IntTy, stringKind: "int64"}, stringKind: "int64[2]"}},
		{"int256[]", Type{Kind: reflect.Slice, T: SliceTy, Type: reflect.TypeOf([]*big.Int{}), Elem: &Type{Kind: reflect.Ptr, Type: bigT, Size: 256, T: IntTy, stringKind: "int256"}, stringKind: "int256[]"}},
		{"int256[2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2]*big.Int{}), Elem: &Type{Kind: reflect.Ptr, Type: bigT, Size: 256, T: IntTy, stringKind: "int256"}, stringKind: "int256[2]"}},
		{"uint8", Type{Kind: reflect.Uint8, Type: uint8T, Size: 8, T: UintTy, stringKind: "uint8"}},
		{"uint16", Type{Kind: reflect.Uint16, Type: uint16T, Size: 16, T: UintTy, stringKind: "uint16"}},
		{"uint32", Type{Kind: reflect.Uint32, Type: uint32T, Size: 32, T: UintTy, stringKind: "uint32"}},
		{"uint64", Type{Kind: reflect.Uint64, Type: uint64T, Size: 64, T: UintTy, stringKind: "uint64"}},
		{"uint256", Type{Kind: reflect.Ptr, Type: bigT, Size: 256, T: UintTy, stringKind: "uint256"}},
		{"uint8[]", Type{Kind: reflect.Slice, T: SliceTy, Type: reflect.TypeOf([]uint8{}), Elem: &Type{Kind: reflect.Uint8, Type: uint8T, Size: 8, T: UintTy, stringKind: "uint8"}, stringKind: "uint8[]"}},
		{"uint8[2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2]uint8{}), Elem: &Type{Kind: reflect.Uint8, Type: uint8T, Size: 8, T: UintTy, stringKind: "uint8"}, stringKind: "uint8[2]"}},
		{"uint16[]", Type{T: SliceTy, Kind: reflect.Slice, Type: reflect.TypeOf([]uint16{}), Elem: &Type{Kind: reflect.Uint16, Type: uint16T, Size: 16, T: UintTy, stringKind: "uint16"}, stringKind: "uint16[]"}},
		{"uint16[2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2]uint16{}), Elem: &Type{Kind: reflect.Uint16, Type: uint16T, Size: 16, T: UintTy, stringKind: "uint16"}, stringKind: "uint16[2]"}},
		{"uint32[]", Type{T: SliceTy, Kind: reflect.Slice, Type: reflect.TypeOf([]uint32{}), Elem: &Type{Kind: reflect.Uint32, Type: uint32T, Size: 32, T: UintTy, stringKind: "uint32"}, stringKind: "uint32[]"}},
		{"uint32[2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2]uint32{}), Elem: &Type{Kind: reflect.Uint32, Type: uint32T, Size: 32, T: UintTy, stringKind: "uint32"}, stringKind: "uint32[2]"}},
		{"uint64[]", Type{T: SliceTy, Kind: reflect.Slice, Type: reflect.TypeOf([]uint64{}), Elem: &Type{Kind: reflect.Uint64, Type: uint64T, Size: 64, T: UintTy, stringKind: "uint64"}, stringKind: "uint64[]"}},
		{"uint64[2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2]uint64{}), Elem: &Type{Kind: reflect.Uint64, Type: uint64T, Size: 64, T: UintTy, stringKind: "uint64"}, stringKind: "uint64[2]"}},
		{"uint256[]", Type{T: SliceTy, Kind: reflect.Slice, Type: reflect.TypeOf([]*big.Int{}), Elem: &Type{Kind: reflect.Ptr, Type: bigT, Size: 256, T: UintTy, stringKind: "uint256"}, stringKind: "uint256[]"}},
		{"uint256[2]", Type{Kind: reflect.Array, T: ArrayTy, Type: reflect.TypeOf([2]*big.Int{}), Size: 2, Elem: &Type{Kind: reflect.Ptr, Type: bigT, Size: 256, T: UintTy, stringKind: "uint256"}, stringKind: "uint256[2]"}},
		{"bytes32", Type{Kind: reflect.Array, T: FixedBytesTy, Size: 32, Type: reflect.TypeOf([32]byte{}), stringKind: "bytes32"}},
		{"bytes[]", Type{T: SliceTy, Kind: reflect.Slice, Type: reflect.TypeOf([][]byte{}), Elem: &Type{Kind: reflect.Slice, Type: reflect.TypeOf([]byte{}), T: BytesTy, stringKind: "bytes"}, stringKind: "bytes[]"}},
		{"bytes[2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2][]byte{}), Elem: &Type{T: BytesTy, Type: reflect.TypeOf([]byte{}), Kind: reflect.Slice, stringKind: "bytes"}, stringKind: "bytes[2]"}},
		{"bytes32[]", Type{T: SliceTy, Kind: reflect.Slice, Type: reflect.TypeOf([][32]byte{}), Elem: &Type{Kind: reflect.Array, Type: reflect.TypeOf([32]byte{}), T: FixedBytesTy, Size: 32, stringKind: "bytes32"}, stringKind: "bytes32[]"}},
		{"bytes32[2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2][32]byte{}), Elem: &Type{Kind: reflect.Array, T: FixedBytesTy, Size: 32, Type: reflect.TypeOf([32]byte{}), stringKind: "bytes32"}, stringKind: "bytes32[2]"}},
		{"string", Type{Kind: reflect.String, T: StringTy, Type: reflect.TypeOf(""), stringKind: "string"}},
		{"string[]", Type{T: SliceTy, Kind: reflect.Slice, Type: reflect.TypeOf([]string{}), Elem: &Type{Kind: reflect.String, Type: reflect.TypeOf(""), T: StringTy, stringKind: "string"}, stringKind: "string[]"}},
		{"string[2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2]string{}), Elem: &Type{Kind: reflect.String, T: StringTy, Type: reflect.TypeOf(""), stringKind: "string"}, stringKind: "string[2]"}},
		{"address", Type{Kind: reflect.Array, Type: addressT, Size: 20, T: AddressTy, stringKind: "address"}},
		{"address[]", Type{T: SliceTy, Kind: reflect.Slice, Type: reflect.TypeOf([]common.Address{}), Elem: &Type{Kind: reflect.Array, Type: addressT, Size: 20, T: AddressTy, stringKind: "address"}, stringKind: "address[]"}},
		{"address[2]", Type{Kind: reflect.Array, T: ArrayTy, Size: 2, Type: reflect.TypeOf([2]common.Address{}), Elem: &Type{Kind: reflect.Array, Type: addressT, Size: 20, T: AddressTy, stringKind: "address"}, stringKind: "address[2]"}},
//正确实现固定类型时的TODO
//“固定”，类型，
//“Fixed128x128”，类型，
//“Fixed[]”，类型，
//“Fixed[2]”，类型，
//“Fixed128x128[]”，类型，
//“Fixed128x128[2]”，类型，
	}

	for _, tt := range tests {
		typ, err := NewType(tt.blob)
		if err != nil {
			t.Errorf("type %q: failed to parse type string: %v", tt.blob, err)
		}
		if !reflect.DeepEqual(typ, tt.kind) {
			t.Errorf("type %q: parsed type mismatch:\nGOT %s\nWANT %s ", tt.blob, spew.Sdump(typeWithoutStringer(typ)), spew.Sdump(typeWithoutStringer(tt.kind)))
		}
	}
}

func TestTypeCheck(t *testing.T) {
	for i, test := range []struct {
		typ   string
		input interface{}
		err   string
	}{
		{"uint", big.NewInt(1), "unsupported arg type: uint"},
		{"int", big.NewInt(1), "unsupported arg type: int"},
		{"uint256", big.NewInt(1), ""},
		{"uint256[][3][]", [][3][]*big.Int{{{}}}, ""},
		{"uint256[][][3]", [3][][]*big.Int{{{}}}, ""},
		{"uint256[3][][]", [][][3]*big.Int{{{}}}, ""},
		{"uint256[3][3][3]", [3][3][3]*big.Int{{{}}}, ""},
		{"uint8[][]", [][]uint8{}, ""},
		{"int256", big.NewInt(1), ""},
		{"uint8", uint8(1), ""},
		{"uint16", uint16(1), ""},
		{"uint32", uint32(1), ""},
		{"uint64", uint64(1), ""},
		{"int8", int8(1), ""},
		{"int16", int16(1), ""},
		{"int32", int32(1), ""},
		{"int64", int64(1), ""},
		{"uint24", big.NewInt(1), ""},
		{"uint40", big.NewInt(1), ""},
		{"uint48", big.NewInt(1), ""},
		{"uint56", big.NewInt(1), ""},
		{"uint72", big.NewInt(1), ""},
		{"uint80", big.NewInt(1), ""},
		{"uint88", big.NewInt(1), ""},
		{"uint96", big.NewInt(1), ""},
		{"uint104", big.NewInt(1), ""},
		{"uint112", big.NewInt(1), ""},
		{"uint120", big.NewInt(1), ""},
		{"uint128", big.NewInt(1), ""},
		{"uint136", big.NewInt(1), ""},
		{"uint144", big.NewInt(1), ""},
		{"uint152", big.NewInt(1), ""},
		{"uint160", big.NewInt(1), ""},
		{"uint168", big.NewInt(1), ""},
		{"uint176", big.NewInt(1), ""},
		{"uint184", big.NewInt(1), ""},
		{"uint192", big.NewInt(1), ""},
		{"uint200", big.NewInt(1), ""},
		{"uint208", big.NewInt(1), ""},
		{"uint216", big.NewInt(1), ""},
		{"uint224", big.NewInt(1), ""},
		{"uint232", big.NewInt(1), ""},
		{"uint240", big.NewInt(1), ""},
		{"uint248", big.NewInt(1), ""},
		{"int24", big.NewInt(1), ""},
		{"int40", big.NewInt(1), ""},
		{"int48", big.NewInt(1), ""},
		{"int56", big.NewInt(1), ""},
		{"int72", big.NewInt(1), ""},
		{"int80", big.NewInt(1), ""},
		{"int88", big.NewInt(1), ""},
		{"int96", big.NewInt(1), ""},
		{"int104", big.NewInt(1), ""},
		{"int112", big.NewInt(1), ""},
		{"int120", big.NewInt(1), ""},
		{"int128", big.NewInt(1), ""},
		{"int136", big.NewInt(1), ""},
		{"int144", big.NewInt(1), ""},
		{"int152", big.NewInt(1), ""},
		{"int160", big.NewInt(1), ""},
		{"int168", big.NewInt(1), ""},
		{"int176", big.NewInt(1), ""},
		{"int184", big.NewInt(1), ""},
		{"int192", big.NewInt(1), ""},
		{"int200", big.NewInt(1), ""},
		{"int208", big.NewInt(1), ""},
		{"int216", big.NewInt(1), ""},
		{"int224", big.NewInt(1), ""},
		{"int232", big.NewInt(1), ""},
		{"int240", big.NewInt(1), ""},
		{"int248", big.NewInt(1), ""},
		{"uint30", uint8(1), "abi: cannot use uint8 as type ptr as argument"},
		{"uint8", uint16(1), "abi: cannot use uint16 as type uint8 as argument"},
		{"uint8", uint32(1), "abi: cannot use uint32 as type uint8 as argument"},
		{"uint8", uint64(1), "abi: cannot use uint64 as type uint8 as argument"},
		{"uint8", int8(1), "abi: cannot use int8 as type uint8 as argument"},
		{"uint8", int16(1), "abi: cannot use int16 as type uint8 as argument"},
		{"uint8", int32(1), "abi: cannot use int32 as type uint8 as argument"},
		{"uint8", int64(1), "abi: cannot use int64 as type uint8 as argument"},
		{"uint16", uint16(1), ""},
		{"uint16", uint8(1), "abi: cannot use uint8 as type uint16 as argument"},
		{"uint16[]", []uint16{1, 2, 3}, ""},
		{"uint16[]", [3]uint16{1, 2, 3}, ""},
		{"uint16[]", []uint32{1, 2, 3}, "abi: cannot use []uint32 as type [0]uint16 as argument"},
		{"uint16[3]", [3]uint32{1, 2, 3}, "abi: cannot use [3]uint32 as type [3]uint16 as argument"},
		{"uint16[3]", [4]uint16{1, 2, 3}, "abi: cannot use [4]uint16 as type [3]uint16 as argument"},
		{"uint16[3]", []uint16{1, 2, 3}, ""},
		{"uint16[3]", []uint16{1, 2, 3, 4}, "abi: cannot use [4]uint16 as type [3]uint16 as argument"},
		{"address[]", []common.Address{{1}}, ""},
		{"address[1]", []common.Address{{1}}, ""},
		{"address[1]", [1]common.Address{{1}}, ""},
		{"address[2]", [1]common.Address{{1}}, "abi: cannot use [1]array as type [2]array as argument"},
		{"bytes32", [32]byte{}, ""},
		{"bytes31", [31]byte{}, ""},
		{"bytes30", [30]byte{}, ""},
		{"bytes29", [29]byte{}, ""},
		{"bytes28", [28]byte{}, ""},
		{"bytes27", [27]byte{}, ""},
		{"bytes26", [26]byte{}, ""},
		{"bytes25", [25]byte{}, ""},
		{"bytes24", [24]byte{}, ""},
		{"bytes23", [23]byte{}, ""},
		{"bytes22", [22]byte{}, ""},
		{"bytes21", [21]byte{}, ""},
		{"bytes20", [20]byte{}, ""},
		{"bytes19", [19]byte{}, ""},
		{"bytes18", [18]byte{}, ""},
		{"bytes17", [17]byte{}, ""},
		{"bytes16", [16]byte{}, ""},
		{"bytes15", [15]byte{}, ""},
		{"bytes14", [14]byte{}, ""},
		{"bytes13", [13]byte{}, ""},
		{"bytes12", [12]byte{}, ""},
		{"bytes11", [11]byte{}, ""},
		{"bytes10", [10]byte{}, ""},
		{"bytes9", [9]byte{}, ""},
		{"bytes8", [8]byte{}, ""},
		{"bytes7", [7]byte{}, ""},
		{"bytes6", [6]byte{}, ""},
		{"bytes5", [5]byte{}, ""},
		{"bytes4", [4]byte{}, ""},
		{"bytes3", [3]byte{}, ""},
		{"bytes2", [2]byte{}, ""},
		{"bytes1", [1]byte{}, ""},
		{"bytes32", [33]byte{}, "abi: cannot use [33]uint8 as type [32]uint8 as argument"},
		{"bytes32", common.Hash{1}, ""},
		{"bytes31", common.Hash{1}, "abi: cannot use common.Hash as type [31]uint8 as argument"},
		{"bytes31", [32]byte{}, "abi: cannot use [32]uint8 as type [31]uint8 as argument"},
		{"bytes", []byte{0, 1}, ""},
		{"bytes", [2]byte{0, 1}, "abi: cannot use array as type slice as argument"},
		{"bytes", common.Hash{1}, "abi: cannot use array as type slice as argument"},
		{"string", "hello world", ""},
		{"string", string(""), ""},
		{"string", []byte{}, "abi: cannot use slice as type string as argument"},
		{"bytes32[]", [][32]byte{{}}, ""},
		{"function", [24]byte{}, ""},
		{"bytes20", common.Address{}, ""},
		{"address", [20]byte{}, ""},
		{"address", common.Address{}, ""},
		{"bytes32[]]", "", "invalid arg type in abi"},
		{"invalidType", "", "unsupported arg type: invalidType"},
		{"invalidSlice[]", "", "unsupported arg type: invalidSlice"},
	} {
		typ, err := NewType(test.typ)
		if err != nil && len(test.err) == 0 {
			t.Fatal("unexpected parse error:", err)
		} else if err != nil && len(test.err) != 0 {
			if err.Error() != test.err {
				t.Errorf("%d failed. Expected err: '%v' got err: '%v'", i, test.err, err)
			}
			continue
		}

		err = typeCheck(typ, reflect.ValueOf(test.input))
		if err != nil && len(test.err) == 0 {
			t.Errorf("%d failed. Expected no err but got: %v", i, err)
			continue
		}
		if err == nil && len(test.err) != 0 {
			t.Errorf("%d failed. Expected err: %v but got none", i, test.err)
			continue
		}

		if err != nil && len(test.err) != 0 && err.Error() != test.err {
			t.Errorf("%d failed. Expected err: '%v' got err: '%v'", i, test.err, err)
		}
	}
}