function mostWordsFound(sentences: string[]): number {
    return sentences.reduce((max, sentence) => Math.max(sentence.split(" ").length, max), 0);
};
