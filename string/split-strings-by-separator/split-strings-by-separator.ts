function splitWordsBySeparator(words: string[], separator: string): string[] {
    const ans: string[] = []
    for (let word of words) {
        ans.push(...word.split(separator).filter(word => word !== ""))
    }
    return ans;
};
