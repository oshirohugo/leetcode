function reverseOnlyLetters(s: string): string {
    const ans = Array.from(s);
    let i = 0;
    let j = s.length - 1;
    while (i < j) {
        if (!isLetter(s[i])) {
            i++;
            continue;
        }
        if (!isLetter(s[j])) {
            j--;
            continue;
        }
        ans[i] = s[j];
        ans[j--] = s[i++];
    }

    return ans.join("")
};

const a_CODE = 'a'.charCodeAt(0);
const z_CODE = 'z'.charCodeAt(0);
const A_CODE = 'A'.charCodeAt(0);
const Z_CODE = 'Z'.charCodeAt(0);
function isLetter(c: string): boolean {
    const code = c.charCodeAt(0);
    return (code >= a_CODE && code <= z_CODE) || (code >= A_CODE && code <= Z_CODE)
}
