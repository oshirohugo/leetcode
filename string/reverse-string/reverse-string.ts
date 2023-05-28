
/**
 Do not return anything, modify s in-place instead.
 */
function reverseString(s: string[]): void {
    let i = 0, j = s.length - 1;

    while (i < j) {
        const aux = s[i];
        s[i++] = s[j];
        s[j--] = aux;
    }
};
