package scrabblescore
import  "strings"

func Score(word string) int {
    point := 0
    a:=strings.ToUpper(word)
    for _,char := range a {
    	switch char {
            case 'A','E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
            	point += 1
            case  'D', 'G':
            	point += 2
            case 'B', 'C', 'M', 'P' :
            	point +=3
            case 'F', 'H', 'V', 'W', 'Y':
            	point +=4
            case 'K':
            	point +=5
            case 'J', 'X':
            	point +=8
        	case 'Q', 'Z':
            	point +=10
        }
	}
    return point
}
