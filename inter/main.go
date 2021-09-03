package main

import (
	"os"
)

func main() {

	if len(os.Args) == 3 {
		str1 := os.Args[1]
		str2 := os.Args[2]
		res := ""
		for _, a := range str1 {
			for _, b := range str2 {
				if a == b {
					isUnique := true
					for _, c := range res {
						if a == c {
							isUnique = false
						}
					}
					if isUnique {
						res += string(a)
					}
				}
			}
		}
		os.Stdout.WriteString(res + "\n")
	}
}

// $ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
// padinto
// $ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
// df6ewg4
// $
