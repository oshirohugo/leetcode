/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
function subarraySum(nums, k) {
  let sum = 0;
  const prefix = nums.map(el => {
    sum += el;
    return sum;
  })

  const sumsMap = {};
  return nums.reduce((acc, curr, index) => {
    if(prefix[index] === k) {
      acc++;
    }
    
    if (sumsMap[prefix[index] - k]) {
      acc += sumsMap[prefix[index] - k];
    }
    
    sumsMap[prefix[index]] =(sumsMap[prefix[index]] || 0) + 1;
    return acc;

  }, 0);
}
