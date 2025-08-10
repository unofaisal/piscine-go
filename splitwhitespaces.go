package piscine

func SplitWhiteSpaces(s string) []string {
    word := ""
    sl := []string{}
    for _, l := range s {
        if l != ' ' && l != '\t' && l != '\n' {
            word += string(l)
        } else if word != "" {
            sl = append(sl, word)
            word = ""
        }
    }
    if word != "" {
        sl = append(sl, word)
    }
    return sl
}
