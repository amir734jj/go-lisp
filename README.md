# go-lisp

Toy LIPS (s-expression) lexer and parser in golang

Features:
- handwritten greedy LISP lexer
- handwritten greedy top-down parser combinator
- simple semantic analyzer
- JavaScript code generation

TODO:
- ~~code generation~~

Example

```lisp
(defun fact (x) (if (<= x 0) 1 (* x (fact (- x 1)))))
(defun fib (x) (if (<= x 0) 0 (if (<= x 1) 1 (+ (fib (- x 1)) (fib (- x 2))))))
(defun fn1 (l) (l 3))
(defun fn2 (l x) (l x))
(fn1 fact)
(fn2 fact 3)
```

JavaScript
```javascript
function fact (x) {
    return (x <= 0.000000 ) ? 1.000000 : (x * fact((x - 1.000000 )) );
}
function fib (x) {
    return (x <= 0.000000 ) ? 0.000000 : (x <= 1.000000 ) ? 1.000000 : (fib((x - 1.000000 )) + fib((x - 2.000000 )) );
}
function fn1 (l) {
    return l(3.000000);
}
function fn2 (l, x) {
    return l(x);
}
fn1(fact)
fn2(fact, 3.000000)

```

Notes:

Thanks to the unambiguous s-expression syntax, top-down parser will successfully pasrer any s-expression, unless we have a syntax error.
