const lookup = new Set(['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U']);
function halvesAreAlike(s: string): boolean {
    let l = 0, r = s.length - 1;
    let leftCount = 0, rightCount = 0;

    while (l < r) {
        if (lookup.has(s[l])) {
            leftCount++;
        }
        if (lookup.has(s[r])) {
            rightCount++;
        }

        l++;
        r--;
    }

    return leftCount === rightCount;
};
