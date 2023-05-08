function repeatedCharacter(s: string): string {
    const lookup: { [key: string]: boolean } = {};

    for (let c of s) {
        if (lookup[c] !== undefined) {
            return c;
        } else {
            lookup[c] = true;
        }
    }

    return s;
};
