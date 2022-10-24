# Examples

### Hello World!
```
(<< "Hello world!")
```
### Comment
```
(<< 111) @ This code prints 111
```
### Variables
```
(= $i 1) @ Creates variable called i with value 1
(= $i 123) @ Changes value called i to 123
(@= $j "a") @ Creates constant variable called j with value "a"
(= $i "abd") @ Error! Can't modify constant variable
(<< $i) @ Prints 123
```
### If,else,loop
```
@ If else else-if
(? (== 1 2) (<< (- 1 2)) (:? (!= 11) (<< (+ 1 1))) (: (<< 1)))
@ Loop
(= $i 1)
(? ($i != 10) (
  (= $i (+ $i 1))
  (~*)
))
```
### List
```
(# "build/arrays") # Imports arrays lib 
(= $list [1,2,3])
(<< (arrays.get 2 $list)) # Returns 3
```
### Modules/Functions
lib.hr
```
(=> $i $@a $@b (+ $a $b)) @ Creates function i
(## $i) @ Exports it
```
main.hr
```
(# "coolproject/lib")
(<< lib.i 1 2) @ Prints out 3
```
