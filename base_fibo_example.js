function fibonacci(n) {
    if (n === 0 || n === 1)
      return n;
    else
      return fibonacci(n - 1) + fibonacci(n - 2);
 /*   3
        f(3 - 1) + f(3 - 2)
           (f(2) + f(0)) + 1
                (f(1) + f(0)) + 1
                    1 + 1   */
        
}

// for(let i = 0; i < 10; i++) {
//     console.log(fibonacci(i))
// }