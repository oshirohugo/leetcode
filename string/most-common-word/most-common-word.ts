function mostCommonWord(paragraph: string, banned: string[]): string {
    let ans = '';
    let max = -1;

    const lookup = new Set<string>(banned)
    const count: { [key: string]: number } = {};

    paragraph = paragraph.replace(/[!?',;.]/g, ' ').toLowerCase();

    for (let word of paragraph.split(/\s/)) {
        if (lookup.has(word)) {
            continue;
        }
        count[word] = count[word] ? count[word] + 1 : 1;
        if (count[word] > max) {
            max = count[word];
            ans = word;
        }
    }

    return ans;
};
