func hasDuplicate(nums []int) bool {
    vistos := make(map[int]bool)
    for _, n := range nums{
        if vistos[n] {
            return true
        }
        vistos[n] = true
    }
    return false
}
