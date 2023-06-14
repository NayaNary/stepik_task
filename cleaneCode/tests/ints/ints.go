package tests


func IntMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func Sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}