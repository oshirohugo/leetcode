function checkPossibility(nums: number[]): boolean {
  let count = 0;
  for (let i = 1; i < nums.length; i++) {
    if (nums[i] < nums[i - 1]) {
        if (count++ || (
            i > 1 && i < nums.length - 1 &&
            nums[i-2] > nums[i] && nums[i+1] < nums[i-1]) ) {
            return false;
        }
    }
  }

  return true;

};

[1, 2, 3, 4, 15, 6, 7, 8, 9]