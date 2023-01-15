function cellsInRange(s: string): string[] {
  const colLow = s[0].charCodeAt(0);
  const colHigh = s[3].charCodeAt(0);
  const rowLow = s[1].charCodeAt(0);
  const rowHigh = s[4].charCodeAt(0);

  const ans: string[] = Array<string>((colHigh - colLow + 1) * (rowHigh - rowLow + 1));

  let i = 0;
  for (let col = colLow; col <= colHigh; col++) {
    for (let row = rowLow; row <= rowHigh; row++) {
      ans[i++] = String.fromCharCode(col) + String.fromCharCode(row);
    }
  }


  return ans;
};
