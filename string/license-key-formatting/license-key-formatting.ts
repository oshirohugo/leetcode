function licenseKeyFormatting(s: string, k: number): string {
    let res = '';
    for (let i = s.length - 1; i >= 0; i--) {
        if (s[i] !== '-') {
            if (res.length % (k + 1) === k) {
                res += '-';
            }
            res += s[i].toUpperCase();
        }
    }
    return res.split('').reverse().join('');
};
