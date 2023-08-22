function reverseVowels(s: string): string {
    const ans = s.split("");
    const vowels = new Set(["a", "e", "i", "o", "u", "A", "E", "I", "O", "U"]);

    let i = 0;
    let j = s.length - 1;
    while (i < j) {
        while (i < j && !vowels.has(ans[i])) {
            i++;
        }
        while (i < j && !vowels.has(ans[j])) {
            j--;
        }
        let aux = s[i];
        ans[i++] = ans[j];
        ans[j--] = aux;
    }

    return ans.join('');
};
