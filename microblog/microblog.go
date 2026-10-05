package microblog

func Truncate(phrase string) string {
	myruns := []rune(phrase) 
    if len(myruns)<5{
        return phrase
    }else{
        return string(myruns[:5])
    }
}
