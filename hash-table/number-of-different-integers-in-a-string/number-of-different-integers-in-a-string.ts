function numDifferentIntegers(word: string): number {
    let lookup = {};
    let uniques = 0;

    let buffer: string[] = [];
    for (let current of word) {
        if (!isLetter(current)) {
            buffer.push(current);
        } else {
            uniques += getUniqueFromBuffer(buffer, lookup)
            buffer = [];
        }
    }

    uniques += getUniqueFromBuffer(buffer, lookup)

    return uniques;
};

const A_CODE = 'a'.charCodeAt(0);
const Z_CODE = 'z'.charCodeAt(0);

function isLetter(letter: string): boolean {
    let charCode = letter.charCodeAt(0)
    return charCode >= A_CODE && charCode <= Z_CODE;
}

function getUniqueFromBuffer(buffer: string[], lookup: any): number {
    let unique = 0;
    if (buffer.length) {
        const str = buffer.join('')
        const num =  Number.parseInt(str);
        const key = num !== Infinity ? num : str;
        if (!lookup[key]) {
            lookup[key] = 1;
            unique = 1;
        } else {
            lookup[key]++;
        }
    }

    return unique;
}