package piscine


func Capitalize(s string) string {
 r := []rune(s)
  for i, c := range r {
     if i < len(r) - 1 && !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9') {
       if r[i+1] >= 'a' && r[i+1] <= 'z' {
         r[i+1] = 'A' + r[i + 1] - 'a'
       }
       
     }
  }
 return string(r)
}