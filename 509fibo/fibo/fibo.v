module fibo

fn fibonaci(n int) int {
    mut a := 1
    mut b := 0
    for i := 0; i < n; i++ {
        a = a + b
        b = a - b
    }
    return b
}

fn main() {
    println(fibonaci(35))
}
