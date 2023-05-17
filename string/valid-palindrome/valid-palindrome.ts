function isPalindrome(s: string): boolean {
    let i = 0;
    let j = s.length - 1;

    while (i < j) {
        if (!isAlphanumeric(s[i])) {
            i++;
            continue;
        }
        if (!isAlphanumeric(s[j])) {
            j--;
            continue;
        }
        if (s[i].toLowerCase() !== s[j].toLocaleLowerCase()) {
            return false;
        }
        i++;
        j--;
    }

    return true;
};

const A_CODE = 'A'.charCodeAt(0);
const a_CODE = 'a'.charCodeAt(0);
const Z_CODE = 'Z'.charCodeAt(0);
const z_CODE = 'z'.charCodeAt(0);
const ZERO_CODE = '0'.charCodeAt(0);
const NINE_CODE = '9'.charCodeAt(0);

function isAlphanumeric(s: string) {
    const code = s.charCodeAt(0);
    return code >= A_CODE && code <= Z_CODE ||
        code >= a_CODE && code <= z_CODE ||
        code >= ZERO_CODE && code <= NINE_CODE;
}
