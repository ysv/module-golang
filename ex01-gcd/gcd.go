package gcd

// Example for acm.ru.
//func main(){
//    a, b := read("INPUT.TXT")
//    write("OUTPUT.TXT", Gcd(a, b))
//}
//
//func read(fName string) (int, int) {
//    in, _ := os.Open(fName)
//    var a, b int
//    fmt.Fscanf(in, "%d %d", &a, &b)
//
//    return a, b
//}
//
//func write(fName string, num int){
//    out, _ := os.Create(fName)
//
//    fmt.Fprintf(out, "%d", num)
//}

func Gcd(a, b int) (int, error) {
    if a < b{
        a, b = b, a
    }

    for a != 0 && b != 0 {
        a, b = b, a % b
    }
    return a, nil
}
