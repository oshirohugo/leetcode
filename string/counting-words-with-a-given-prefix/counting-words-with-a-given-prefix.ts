function prefixCount(words: string[], pref: string): number {
    return words.reduce((sum, word) => word.indexOf(pref) === 0 ? sum + 1 : sum, 0)
};
