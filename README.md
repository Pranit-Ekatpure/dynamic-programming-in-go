# Dynamic Programming in Golang

## Introduction

Dynamic programming is a way of breaking a complex problem into the sub-problems, solving each of the these sub-problems once, and saving the solutions for later use.

## Dynamic Programming Techniques
* **_Memoization_**  
    _Memoization_ is a technique for improving the performance of a recursive function/algorithm. It offers an optimization to speed up programs by storing the solution of expensive fuction calls and returning the cached solution when the same inputs are fed to the program again.

    * Example  
    [See memoization go code example here](memoization/memoization.go)  
    Running the code will result in:  
    ```
    2023/07/06 01:32:38 --> Loop solution | result: 12586269025 | took: 1.201µs  
    2023/07/06 01:32:38 --> Memoization solution | result: 12586269025 | took: 491ns  
    2023/07/06 01:32:38 --> Recursion solution | result: 12586269025 | took: 3m57.0162348s
    ```  
    Running the memoization solution with an input of `50` result in `12586269025` with execution time of `491µs`, much faster and more efficient that the recursion solution. In addition to the performance gain, we also have the flexibility of making the solution modular and readable.  

    [Reference](https://betterprogramming.pub/dynamic-programming-in-go-a95d32ee9953)


