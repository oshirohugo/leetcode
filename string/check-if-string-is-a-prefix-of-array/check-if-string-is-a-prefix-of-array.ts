function isPrefixString(s: string, words: string[]): boolean {
    let check = "";

    for (let word of words) {
        check += word;
        if (check === s) {
            return true;
        }
        if(check.length > s.length) {
            return false;
        }
    }

    return false;
};
