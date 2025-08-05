package piscine

func LastRune(s string) rune {
  r := []rune(s)
  var x rune
  for i, c := range r {
    if i == len(s) - 1 {
      x = c
    }
  }
  return x
}