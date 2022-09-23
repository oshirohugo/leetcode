/**
 * @param {number[]} nums
 * @param {number} p
 * @return {number}
 */
function minSubarray(nums, p) {
  const target = (nums.reduce((acc, curr) => acc + curr, 0)) % p;

  if (!target) {
    return 0;
  }

  const remains = {};
  remains[0] = -1;

  let res = nums.length;
  let sum = 0;
  for (let j = 0; j < nums.length; j++) {
    sum += nums[j];
    const currentRemain = sum % p;
    let key = (currentRemain - target)%p;
    key = key < 0 ? key + p : key;
    const i = remains[key];

    if (i !== undefined) {
      res = Math.min(j - i, res);
    }

    remains[currentRemain] = j;
  }

  return res === nums.length ? -1 : res;
}
