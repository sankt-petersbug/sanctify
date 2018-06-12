package main

import (
	"bytes"
	"strings"
    "fmt"
    "os"
    "io/ioutil"
)

const sig = `
  SSSSS   TTTTTTTT      PPPPPPP     EEEEEE  TTTTTTTT   EEEEEE  RRRRRRR     SSSSS   BB        UU    UU   GGGGG      
 SS   SS  T  TT  T      PP     PP  EE       T  TT  T  EE       RR     RR  SS   SS  BB        UU    UU  GG    GG
 SS          TT         PP     PP  EE          TT     EE       RR     RR  SS       BB        UU    UU  G  
   SSS       TT         PPPPPPP     EEEEEE     TT      EEEEEE  RRRRRRR      SSS    BBBBBB    UU    UU  G    GGG  
      SS     TT         PP         EE          TT     EE       RR  R           SS  BB    BB  UU    UU  G      G
 SS   SS     TT     .   PP         EE          TT     EE       RR   RR    SS   SS  BB    BB  UU    UU  GG    GG
  SSSSS      TT    ...  PP          EEEEEE     TT      EEEEEE  RR     RR   SSSSS   BBBBBB      UUUU     GGGGG
`

func Tokenize(s string) []rune {
	var chars []rune

	for _, r := range s {
		if r != '\n' && r != '\r' {
			if r == ' ' {
				chars = append(chars, '0')
			} else {
				chars = append(chars, r)
			}
		}
	}

	return chars
}

func String(pos [][]int, chars []rune) string {
	if pos == nil || chars == nil {
		return ""
	}

	var buf bytes.Buffer
	var r = 0

	// TODO: use generator.. NextChar
	// separate decision making func, (int) -> nextChar or space

	for _, row := range pos {
		for _, col := range row {
			switch col {
			case 0:
				buf.WriteRune(' ')
			case 1:
				if r >= len(chars) {
					r = 0
				}

				buf.WriteRune(chars[r])
				r++
			}
		}
		buf.WriteRune('\n')
	}

	return buf.String()
}

func Position(s string) [][]int {
	var pos [][]int

	for _, line := range strings.Split(s, "\n") {
		row := make([]int, len(line))
		for i, c := range line {
			switch c {
			case ' ':
				row[i] = 0
			default:
				row[i] = 1
			}
		}

		pos = append(pos, row)
	}

	return pos
}

func main() {
    fpath := os.Args[1]
    data, err := ioutil.ReadFile(fpath)

    if err != nil {
        panic(err)
    }

	pos := Position(sig)
	chars := Tokenize(string(data))

	fmt.Println(String(pos, chars))

}
