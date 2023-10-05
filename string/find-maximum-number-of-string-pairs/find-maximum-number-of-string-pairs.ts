function maximumNumberOfStringPairs(words: string[]): number {
    const lookup = new Set<string>();
    let pairs = 0;

    for (let word of words) {
        if (lookup.has(word[1] + word[0])) {
            pairs++;
        }
        lookup.add(word)
    }

    return pairs;
};
