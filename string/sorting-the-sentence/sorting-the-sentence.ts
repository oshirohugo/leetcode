function sortSentence(s: string): string {
    const words = s.split(" ");

    const sentence = Array<string>(words.length);

    for (let word of words) {
        const i = word[word.length - 1];
        sentence[parseInt(i, 10) - 1] = word.slice(0, -1);
    }

    return sentence.join(" ")
};
