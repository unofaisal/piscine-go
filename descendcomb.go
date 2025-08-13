package piscine
import "github.com/01-edu/z01"

func DescendComb() {
  for i := 99; i >= 1; i-- {
    for x := i - 1; x >= 0; x-- {
      print2dig(i)
      z01.PrintRune(' ')
      print2dig(x)
      if !(i == 1 && x == 0) {
        z01.PrintRune(',')
      z01.PrintRune(' ')
      // fmt.Print(i, " ", x, ", ")
      }
      
    }
  }
}

func print2dig(in int) {
  z01.PrintRune(rune('0' + in/10))
  z01.PrintRune(rune('0' + in%10))
}
