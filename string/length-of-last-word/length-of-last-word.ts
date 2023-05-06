
function lengthOfLastWord(s: string): number {
    let tail = s.length - 1;
    while (tail >= 0 && s[tail] === ' ') tail--;
    let ans = 0;
    while (tail >= 0 && s[tail] !== ' ') {
        ans++;
        tail--;
    };

    return ans;
}

function easyLengthOfLastWord(s: string): number {
    const w = s.trim().split(/\s/)
    return w[w.length - 1].length;
}
