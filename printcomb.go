// Empty file

package piscine

import "github.com/01-edu/z01"

func PrintComb() {
   for a := 0; a <= 7; a++{
     for b := (a + 1); b < 8; b++{
       for c := (b + 1); c < 9; c++{
         z01.PrintRune('0' + rune(a))
         z01.PrintRune('0' + rune(b))
         z01.PrintRune('0' + rune(c))
         if !(a==7 && b==8 && c==9){
           z01.PrintRune(',')
           z01.PrintRune(' ')
         }
       }
     }
   }
}