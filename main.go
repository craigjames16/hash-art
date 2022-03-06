package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"

	"github.com/craigjames16/hash-art/colors"
	"github.com/fogleman/gg"
)

func parseSHA256(inputString []byte) string {
	sum := sha256.Sum256(inputString)
	return fmt.Sprintf("%x", sum)
}

func getHashSlice(inputString string) []int64 {
	var hashSlice []int64
	for i := 0; i < len(inputString); i += 2 {
		hashSlice = append(hashSlice, decodeHex(inputString[i:i+2]))
	}

	return hashSlice
}

func getPropertyGroup(inputSlice []int64) [][]int64 {
	var groups [][]int64
	for i := 0; i < len(inputSlice); i += 4 {
		groups = append(groups, inputSlice[i:i+4])
	}

	return groups
}

func decodeHex(hexValue string) int64 {
	num, err := strconv.ParseInt(hexValue, 16, 64)
	if err != nil {
		panic(err)
	}

	return num
}

func main() {
	arg := os.Args[1]
	hash := parseSHA256([]byte(arg))
	hashSlice := getHashSlice(hash)
	propertyGroups := getPropertyGroup(hashSlice)

	drawImage(propertyGroups)
}

func drawImage(propertyGroups [][]int64) {
	var (
		dim    = 1000
		radius float64
	)

	dc := gg.NewContext(dim, dim)

	for i := 0; i < len(propertyGroups); i++ {
		properties := propertyGroups[i]
		radius = float64(float64(dim/2)-(float64(i)*float64(dim)/16-(float64(dim)/16*(float64(properties[0])/255)))) - 50
		dc.Push()
		dc.DrawCircle(500, 500, radius)

		circleColor := colors.GetColor(properties[1])
		cicleOpacity := float64(properties[2]) / 255
		dc.SetRGBA(circleColor[0], circleColor[1], circleColor[2], cicleOpacity)
		dc.FillPreserve()

		strokeColor := colors.GetColor(properties[3])
		dc.SetRGBA(strokeColor[0], strokeColor[1], strokeColor[2], cicleOpacity)
		dc.SetLineWidth(3)
		dc.Stroke()

		dc.Pop()
	}

	dc.SavePNG("out.png")
}
