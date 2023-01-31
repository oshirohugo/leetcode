function reverseWords(s: string): string {
    const reversed = Array<string>(s.length)

    let j = 0;
    for (let i = 0; i < s.length; i++) {
        if (s[i] === " ") {
            reverseWord(s.slice(j, i), reversed, j)
            j = i + 1;
        }
    }

    return reversed.join(" ")
};

function reverseWord(s: string, arr: string[], start: number) {
    let j = start + s.length - 1;
    for(let i = 0; i < s.length; i++) {
        arr[j--] = s[i];
    }
}
