# go-lisp

Toy LIPS (s-expression) lexer and parser in golang

Features:
- handwritten greedy LISP lexer
- handwritten greedy top-down parser combinator for LISP

TODO:
- code generation

Example
```
(defun fact (x) (if (<= x 0) 1 (* x (fact (- x 1)))))
(defun fib (x) (if (<= x 0) 0 (if (<= x 1) 1 (+ (fib (- x 1)) (fib (- x 2))))))
(defun fn1 (l) (l 3))
(defun fn2 (l x) (l x))
(fn1 fact)
(fn2 fact 3)
```