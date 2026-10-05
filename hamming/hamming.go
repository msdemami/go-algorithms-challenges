package hamming
import  "errors"


func Distance(a, b string) (int, error) {
    dis :=0 
    if len(a) != len(b){
        return 0 , errors.New("DNA lengths do not match")
    }else {
        for i := 0 ; i < len(a) ; i++ {
            if a[i] != b[i]{
                dis++
            }
         
        }
    }
    return dis , nil
}
