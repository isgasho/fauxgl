package soft

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type plyDataType int

const (
	plyNone plyDataType = iota
	plyChar
	plyUchar
	plyShort
	plyUshort
	plyInt
	plyUint
	plyFloat
	plyDouble
	plyInt8
	plyUint8
	plyInt16
	plyUint16
	plyInt32
	plyUint32
	plyFloat32
	plyFloat64
)

var plyDataTypeMapping = map[string]plyDataType{
	"char":   plyInt8,
	"uchar":  plyUint8,
	"short":  plyInt16,
	"ushort": plyUint16,
	"int":    plyInt32,
	"uint":   plyUint32,
}

type plyProperty struct {
	name      string
	countType plyDataType
	dataType  plyDataType
}

type plyElement struct {
	name       string
	count      int
	properties []plyProperty
}

func LoadPLY(path string) (*Mesh, error) {
	// open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		fmt.Println(fields)
	}

	// [ply]
	// [format ascii 1.0]
	// [comment generated by ply_writer]
	// [element vertex 437645]
	// [property float x]
	// [property float y]
	// [property float z]
	// [element face 871414]
	// [property list uchar int vertex_indices]
	// [end_header]

	// check for errors
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return nil, err
}