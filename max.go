package piscine

func Max(a []int) int {
  max := 0
  if len(a) == 0 {
    return max
  }
  for _, in := range a {
    if in > max {
      max = in
    }
  }
  return max
}