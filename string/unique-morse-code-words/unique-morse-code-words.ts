function uniqueMorseRepresentations(words: string[]): number {
    const uniqueCodes = new Set<string>();
    for (let word of words) {
        const code = getCode(word);
        if (!uniqueCodes.has(code)) {
            uniqueCodes.add(code);
        }
    }

    return uniqueCodes.size;
};


const DIC =[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."];
const A_CODE = "a".charCodeAt(0);
function getCode(word: string) : string {
    const ans : string[] = [];

    for (let c of word) {
        const code = c.charCodeAt(0) - A_CODE;
        ans.push(DIC[code]);
    }

    return ans.join("");
}
