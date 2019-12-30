const randomRange = (min, max) => {
  return Math.floor(Math.random() * (max - min)) + min;
}

const getRoundNum = (length) => {
  let list = [];
  let numList = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];

  for (let i = 1; i <= length; i++) {
    let randomNum = Math.floor(Math.random() * numList.length);
    list.push(numList[randomNum]);
    numList.splice(randomNum, 1);
  }
  return list;
}

const getRoundLetter = (length) => {
  let list = [];
  let letterList = [
    "A",
    "B",
    "C",
    "D",
    "E",
    "F",
    "G",
    "H",
    "I",
    "J",
    "K",
    "L",
    "M",
    "N",
    "O",
    "P",
    "Q",
    "R",
    "S",
    "T",
    "U",
    "V",
    "W",
    "X",
    "Y",
    "Z"
  ];

  for (let i = 1; i <= length; i++) {
    let randomNum = Math.floor(Math.random() * letterList.length);
    list.push(letterList[randomNum]);
    letterList.splice(randomNum, 1);
  }
  return list;
}

const guid2 = () => {
  function S4() {
    return (((1 + Math.random()) * 0x10000) | 0).toString(16).substring(1);
  }
  return (S4() + S4() + "-" + S4() + "-" + S4() + "-" + S4() + "-" + S4() + S4() + S4());
}

const getNumberTask = (type,num) => {
 
  let typeTextList = ["顺序", "倒叙", "从小到大", "按字母表顺序", "加和"]
  if(!num){
    num = randomRange(2, 5)
  }
  let list = getRoundNum(num)
  return {
    type: type,
    typeName: typeTextList[type],
    list: list,
    right:getRight(list,type)
  };
}
const getLetterTask = (type,num) => {
  let typeTextList = ["顺序", "倒叙", "从小到大", "按字母表顺序", "加和"]
  if(!num){
    num = randomRange(2, 5)
  }
  let list = getRoundLetter(num)
  return {
    type: type,
    typeName: typeTextList[type],
    list: list,
    right:getRight(list,type)
  };
}

const getRight = (list,type) => {
  let right = JSON.parse(JSON.stringify(list));
  switch (type) {
    case 0:
      right = right.join('');
      break;
    case 1:
      right = right.reverse().join('');
      break;
    case 2:
      right = right
        .sort()
        .join('');
      break;
    case 3:
      right = right
        .sort()
        .join('');
      break;
    case 4:
      right = eval(right.join("+"));
      right = right.toString();
      break;
  }
  return right
}
const Utils = {
  randomRange,
  getRoundNum,
  getRoundLetter,
  guid2,
  getNumberTask,
  getLetterTask,
  getRight,
}
const install = (Vue) => {
  Vue.prototype.Utils = Utils
}

export default install
export { Utils }