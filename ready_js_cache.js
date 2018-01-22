const hash = require('object-hash');

const pow = (arg, power) => arg * power;
const myFunc = (arg) => pow(arg, 5);

const showEmplInfo = (empl) => {
    return `${empl.job}, ${empl.age}`;
}

function memoize(foo) {
    let memo = {};
    return function() {
        const slice = Array.prototype.slice;
        const args = hash(slice.apply(arguments));
        if (args in memo) {
            console.log('cached');
            return memo[args]
        }
        else {
            console.log('not cached');
            return memo[args] = foo(...arguments)
        }
    }
};

const testFunc = memoize(myFunc);
console.log(testFunc(5));
console.log(testFunc(5));
console.log(testFunc(5));

const testEmpInfo = memoize(showEmplInfo);
const vasia = { job: 'ba', age: 20 };
const masha = { job: 'qa', age: 21 };
console.log(testEmpInfo(vasia));
console.log(testEmpInfo(vasia));
console.log(testEmpInfo(masha));
console.log(testEmpInfo(vasia));
console.log(testEmpInfo({ job: 'ba', age: 20 }));

// const arr0 = [3,2,1];
// const arr1 = [1,3,2];
// console.log(hash(arr0) == hash(arr1));