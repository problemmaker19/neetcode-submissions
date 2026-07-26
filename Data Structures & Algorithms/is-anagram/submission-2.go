import "slices"

func isAnagram(s string, t string) bool {
    
    sRuneSlice := []rune(s)
    tRuneSlice := []rune(t)

    slices.Sort(sRuneSlice)
    slices.Sort(tRuneSlice)

    return slices.Equal(sRuneSlice, tRuneSlice)
}
