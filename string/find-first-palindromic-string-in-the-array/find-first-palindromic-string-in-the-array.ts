function firstPalindrome(words: string[]): string {
    for (let word of words) {
        if (isPalindromic(word)) {
            return word;
        }
    }
    return '';
};

function isPalindromic(word: string): boolean {
    let i = 0;
    let j = word.length - 1;
    while (i < j) {
        if (word[i++] !== word[j--]) return false;
    }

    return true;
}
