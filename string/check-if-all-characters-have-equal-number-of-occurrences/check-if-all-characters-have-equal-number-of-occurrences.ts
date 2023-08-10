const A = "a".charCodeAt(0);

function areOccurrencesEqual(s: string): boolean {
    const lookup = new Array(27).fill(0);

    for (let c of s) {
        lookup[c.charCodeAt(0) - A]++;
    }
    const count = lookup[s[0].charCodeAt(0) - A];

    for (let i = 0; i < lookup.length; i++) {
        if (lookup[i] !== 0 && lookup[i] !== count) {
            return false;
        }
    }

    return true;
};
