
function twoLargest(nums: number[]): number[] {
    let largest = -1;
    let second = -1;
    for (let i = 0; i < nums.length; i++) {
        let current = nums[i];
        if (current > second) {
            if (current > largest) {
                largest = current
            } else {
                second = current;
            }
        }
    }

    return [second, largest];
}