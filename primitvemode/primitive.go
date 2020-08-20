package primitivemode

import (
	"os/exec"
	"strings"
	"fmt"
)

// PrimitiveMode Type
type PrimitiveMode int

const (
	combo PrimitiveMode = iota
	triangle
	rect
	ellipse
	circle
	rotatedrect
	beziers
	rotatedellipse
	polygon
)

// Primitive method
func Primitive(inputFile, outputFile string, numShapes int, mode PrimitiveMode) (string, error) {
	argStr := fmt.Sprintf("-i %s -o %s -n %d -m %d", inputFile, outputFile, numShapes, mode)
	cmd := exec.Command("primitive", strings.Fields(argStr)...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return string(b), err
}
