class OrderedStream {
    private stream: string[];
    private index: number;
    constructor(n: number) {
        this.stream = new Array(n);
        this.index = 0;
    }

    insert(idKey: number, value: string): string[] {
        const streamKey = idKey - 1;
        this.stream[streamKey] = value;

        if (this.index === streamKey) {
            while(this.stream[this.index]) {
                this.index++;
            }
        }

        return this.stream.slice(streamKey, this.index);
    }
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * var obj = new OrderedStream(n)
 * var param_1 = obj.insert(idKey,value)
 */