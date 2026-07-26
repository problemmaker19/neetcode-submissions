func hasDuplicate(nums []int) bool {
    memory := make(map[int]bool)
    
    for _, v := range nums {
        if memory[v] {
            return true
        }
        memory[v] = true
    }

    return false
}
