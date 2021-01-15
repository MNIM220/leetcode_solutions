func twoSum(nums []int, target int) []int {
    for i, val := range nums {
        for j, valin := range nums {
            if i == j {
                continue
            }
            if val + valin ==target { 
                var result []int 
                return append(result, i, j)
            }
        }     
    }
    return nil
}
