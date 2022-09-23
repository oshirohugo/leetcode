function getFolderNames(names: string[]): string[] {
  
  const namesHash: {[hash:string]: any} = {};
  const res: string[] = [];
  
  names.forEach(name => {
    const [prefix, extractedIndex] = getPrefix(name);
    
    //console.log(namesHash)
    
    const counterMap = namesHash[prefix];
    if (counterMap) {
    
      const index = getNextAvailableIndex(counterMap, extractedIndex);
      const newName =  `${prefix}(${index})`;
      
      namesHash[prefix] = counterMap;
      res.push(newName);
    } else {
      namesHash[prefix] = {};
      res.push(name);
    }
  });
  

  return res;
};
                
function getNextAvailableIndex(indexMap: any, proposed: number) {
    if (proposed && !indexMap[proposed]) {
      indexMap[proposed] = true;
      return proposed;
    }
    let index = 1;
    for (_ in indexMap) {
      if (!indexMap[index]) {
        indexMap[index] = true;
        return index;
      }
      index++;
    }
    indexMap[index] = true;
    return index;
}
                
function getPrefix(name: string): [string, number] {
  const match = name.match(/([a-zA-Z0-9_]+)(\((\d+)\))?/)
  
  const count = match[3] ? parseInt(match[3]) : undefined;
  const prefix = match[1]
  
  return  [prefix, count];
}
                
                