package maximumProductofTwoElementsinanArray

func maxProduct(nums []int) int {
    var firstBiggest int
    var secBiggest int
    if nums[0] > nums[1] {
        firstBiggest = 0
        secBiggest = 1
    } else {
        firstBiggest = 1
        secBiggest = 0
    }
    
    for i := 2; i < len(nums); i++ {
        if nums[i] > nums[firstBiggest] {
            secBiggest = firstBiggest
            firstBiggest = i
        } else if nums[i] > nums[secBiggest] {
            secBiggest = i
        }
    }
    return (nums[firstBiggest] - 1) * (nums[secBiggest] - 1)
}
