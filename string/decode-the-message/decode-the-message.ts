const A_CODE = 'a'.charCodeAt(0);

function decodeMessage(key: string, message: string): string {
    const dic = {' ': ' '};
    let code = A_CODE;
    for (let c of key) {
        if (c === ' ' || dic[c]) continue;
        const newKey = String.fromCharCode(code++);
        dic[c] = newKey;
    }

    const ans = Array<string>(message.length)
    for (let i = 0; i < message.length; i++) {
        ans[i] = dic[message[i]];
    }

    return ans.join('');
};
