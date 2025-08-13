package piscine

func Rot14(s string) string {
  nS := ""
  for _, c := range s {
    if c >= ' ' && c <= '~' && !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z') {
      nS += string(c)
    }
    if c >= 'a' && c <= 'z' {
      if c - 'a' + 14 > 25 {
        diff := c - 'a' + 14 - 25
        nS += string('a' - 1 + diff)
      }else {
        nS += string(c + 14)
      }
    }
    if c >= 'A' && c <= 'Z' {
        if c - 'A' + 14 > 25 {
        diff := c - 'A' + 14 - 25
        nS += string('A' - 1 + diff)
      }else {
        nS += string(c + 14)
      }
    }
  }
  return nS
}
