function fizzBuzz(n: number): string[] {
    const answer = Array<string>(n);

    for (let i = 0; i < n; i++) {
        const num = i + 1;
        let ans = "";
        if (num % 3 === 0) {
            ans += "Fizz";
        }
        if (num % 5 === 0) {
            ans += "Buzz";
        }
        if (ans === "") {
            ans = `${num}`;
        }
        answer[i] = ans;
    }

    return answer;
};
