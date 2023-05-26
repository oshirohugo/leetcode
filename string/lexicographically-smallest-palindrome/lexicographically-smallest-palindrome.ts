function makeSmallestPalindrome(s: string): string {
    let i = 0, j = s.length - 1;
    const output = Array.from(s)

    while (i < j) {
        if (s[i] < s[j]) {
            output[j] = s[i];
        } else {
            output[i] = s[j];
        }
        i++;
        j--;
    }

    return output.join('')
};
