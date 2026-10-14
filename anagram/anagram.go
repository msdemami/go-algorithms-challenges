package anagram
import  ("strings"
         "slices"
)


func Detect(subject string, candidates []string) []string {
	var result []string
    
	lower_sub := strings.ToLower(subject)
    sub_rune := []rune(lower_sub)
    slices.Sort(sub_rune)
    sort_sub := string(sub_rune)

    for _,candida := range candidates {
        lower_candida := strings.ToLower(candida)

		if lower_candida == lower_sub{
            continue
        }
        
        candida_rune := []rune(lower_candida)
        slices.Sort(candida_rune)
        sort_candida := string(candida_rune)

        if sort_sub == sort_candida {
        	result = append(result , candida)
        }
        
        
    }
    return result
}
