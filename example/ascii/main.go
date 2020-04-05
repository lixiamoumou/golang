package main

import (
	"bytes"
	"fmt"
	"os"
)

type Ascci struct {
	Dec         int
	Symble      string
	Description string
}

var arr []Ascci = []Ascci{
	{0, "NUL", "null"},
	{1, "SOH", "start of heading"},
	{2, "STX", "start of text"},
	{3, "ETX", "end of text"},
	{4, "EOT", "end of transmission"},
	{5, "ENQ", "enquiry"},
	{6, "ACK", "acknowledge"},
	{7, "BEL", "bell"},
	{8, "BS", "backspace"},
	{9, "TAB", "horizontal tab"},
	{10, "LF", "NL line feed, new line"},
	{11, "VT", "vertical tab"},
	{12, "FF", "NP form feed, new page"},
	{13, "CR", "carriage return"},
	{14, "SO", "shift out"},
	{15, "SI", "shift in"},
	{16, "DLE", "data link escape"},
	{17, "DC1", "device control 1"},
	{18, "DC2", "device control 2"},
	{19, "DC3", "device control 3"},
	{20, "DC4", "device control 4"},
	{21, "NAK", "negative acknowledge"},
	{22, "SYN", "synchronous idle"},
	{23, "ETB", "end of trans. block"},
	{24, "CAN", "cancel"},
	{25, "EM", "end of medium"},
	{26, "SUB", "substitute"},
	{27, "ESC", "escape"},
	{28, "FS", "file separator"},
	{29, "GS", "group separator"},
	{30, "RS", "record separator"},
	{31, "US", "unit separator"},
}

//func Getmax(s []Ascci) int{
//	max := 0
//	05for _, v := range s{
//		length := len(k) + len(v.Symble) + len(v.Description) +
//			len(strconv.FormatInt(int64(v.Dec), 16)) + 3 * 2
//		if length >= max {
//			max = length
//		}
//	}
//	return max
//}

func ShowAscci() {
	var col, row int
	//Dec  Hex  Symble-16 __Des-28
	title1 := "Dec  Hex  Symble  Description          "
	title2 := "  Dec  Hex  Char "

	//print title
	for col = 0; col < 4; col++ {
		if col == 0 {
			fmt.Printf(title1)
		} else {
			fmt.Printf(title2)
		}
	}
	fmt.Printf("\n")

	//print ascci
	for row = 0; row < 32; row++ {
		buf := bytes.Buffer{}
		//offset := 0
		for col = 0; col < 4; col++ {
			if col == 0 {
				s := fmt.Sprintf(" %-3d  %02[1]x  %-6s  %s", arr[row].Dec,
					arr[row].Symble, arr[row].Description)
				buf.Write([]byte(s))
				if len(s) < len(title1) {
					space := make([]byte, len(title1)-len(s)+1)
					buf.Write([]byte(space))
				}

			} else {
				num := col*32 + row
				s := fmt.Sprintf("| %-3d  %02[1]x  %2[1]c", num)
				buf.Write([]byte(s))
				if len(s) < len(title2) {
					space := make([]byte, len(title2)-len(s))
					buf.Write([]byte(space))
				}
			}
		}
		buf.WriteTo(os.Stdout)
		buf.Reset()
		fmt.Println()
	}
}

func main() {
	ShowAscci()
}
