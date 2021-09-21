package piscine

func FindNextPrime(nb int) int {
    if nb < 1 {
        return 2
    }
    for i := nb; i < 9223372036854775807; i++ {
        for a := 2; a < i; a++ {
            if i%a != 0 {
                return i
            }
        }
    }
    return 0
}