function sortString(s: string): string {
    const count = new Array<number>(26).fill(0);
    for (let c of s) {
        count[c.charCodeAt(0) - 'a'.charCodeAt(0)]++;
    }

    let ans: string[] = [];

    while (ans.length !== s.length) {
        for (let i = 0; i < 26; i++) {
            if (count[i] > 0) {
                ans.push(String.fromCharCode(i + 'a'.charCodeAt(0)))
                count[i]--;
            }
        }
        for (let i = 25; i >= 0; i--) {
            if (count[i] > 0) {
                ans.push(String.fromCharCode(i + 'a'.charCodeAt(0)))
                count[i]--;
            }
        }
    }

    return ans.join('');
};
