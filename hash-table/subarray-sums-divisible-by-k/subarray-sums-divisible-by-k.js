/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
function subarraysDivByK (nums, k) {
  const freq = new Array(k).fill(0);
  freq[0] = 1;
 
  let sum = 0;
  return nums.reduce((acc, num) => {
    sum += num;

    let remainder = sum % k;

    // only consider positive remainders
    if (remainder < 0) {
      remainder += k;
    }
    
    acc += freq[remainder];

    freq[remainder]++;
        
    return acc;
  }, 0);
}
