package piscine

func ToUpper(s string) string {
    s1 := []rune(s)
    str := ""
    for i := 0; i <= len(s1)-1; i++ {
        if s1[i] >= 'a' && s1[i] <= 'z' {
            s1[i] = s1[i] - 32
        }
        str = str + string(s1[i])
    }
    return str
}