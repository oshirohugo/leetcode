function countConsistentStrings(allowed: string, words: string[]): number {
    const lookup = new Set<string>()
    for (let c of allowed) {
        lookup.add(c);
    }
    let consistent = words.length;

    for (let word of words) {
        for (let c of word) {
            if (!lookup.has(c)) {
                consistent--;
                break;
            }
        }
    }

    return consistent;

};
