function mergeAlternately(word1: string, word2: string): string {
    const l1 = word1.length;
    const l2 = word2.length;

    const ans = new Array(l1 + l2);

    let i1 = 0; 
    let i2 = 0; 
    let i = 0;
    while (i1 < l1 && i2 < l2) {
        if (i % 2 === 0) {
            ans[i++] = word1[i1++];
        } else {
            ans[i++] = word2[i2++];
        }
    }

    if (i1 < l1) {
        while(i1 < l1) {
            ans[i++] = word1[i1++];
        }
    }

    if (i2 < l2) {
        while(i2 < l2) {
            ans[i++] = word2[i2++];
        }
    }

    return ans.join("");
};
