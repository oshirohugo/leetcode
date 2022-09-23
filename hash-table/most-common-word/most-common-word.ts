function mostCommonWord(paragraph: string, banned: string[]): string {

  const counter: { [key: string]: number } = banned.reduce((acc, el) => {
    return {
      ...acc,
      el: 0,
    };
  }, {});

  const heightFrequency = "";
  paragraph.replace(/[!?',;.]/, '')
    .split(" ")
    .forEach(word => {
      if (counter[word]) {
        counter[word]++;

      }
    })


};