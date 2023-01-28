const a_code = "a".charCodeAt(0);
const A_CODE = "A".charCodeAt(0);
const Z_CODE = "Z".charCodeAt(0);
const LOWER_UPPER_DIFF = a_code - A_CODE;

function toLowerCase(s: string): string {
    const lower = Array<string>(s.length);

    for (let i = 0; i < s.length; i++) {
        lower[i] = s[i];

        const code = s[i].charCodeAt(0);
        if (code >= A_CODE && code <= Z_CODE) {
            const newChar = String.fromCharCode(code + LOWER_UPPER_DIFF);
            lower[i] = newChar;
        }
    }

    return lower.join("");
};
