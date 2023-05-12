function canBeTypedWords(text: string, brokenLetters: string): number {
    const broken = new Array<boolean>(26).fill(false);

    for (let c of brokenLetters) {
        broken[c.charCodeAt(0) - 'a'.charCodeAt(0)] = true;
    }

    let res = 0;
    let shouldCount = true;
    for (let c of text) {
        if (c === ' ') {
            res += Number(shouldCount);
            shouldCount = true;
        } else if (broken[c.charCodeAt(0) - 'a'.charCodeAt(0)]) {
            shouldCount = false;
        }
    }

    return res + Number(shouldCount);
};
