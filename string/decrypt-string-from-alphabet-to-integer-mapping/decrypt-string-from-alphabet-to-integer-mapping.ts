function freqAlphabets(s: string): string {
    let res = '';
    for (let i = 0; i < s.length;) {
        if (i + 2 < s.length && s[i + 2] === "#") {
            res += String.fromCharCode('j'.charCodeAt(0) +
                (s[i].charCodeAt(0) - '1'.charCodeAt(0)) * 10 +
                s[i + 1].charCodeAt(0) - '0'.charCodeAt(0)
            );
            i += 3;
        } else {
            res += String.fromCharCode('a'.charCodeAt(0)
                + s[i].charCodeAt(0) - '1'.charCodeAt(0))
            i++;
        }

    }

    return res;
};
