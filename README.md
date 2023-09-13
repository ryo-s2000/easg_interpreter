# easg_interpreter

## run
```
go run main.go
>>write code here!
```

## let
```
>>let a = 1
>>a
1

>>let b = "b"
>>b
b
```

## Array
```
>>[1,2,"3"][1]
2

>>[1,2,"3"][-2]
2
```

## Hash
```
>>let h = {"a": "a_text", "b": "b_text"}
>>h["a"]
a_text

>>h["b"]
b_text
```

## Func
```
>>let add = fn(a, b) { return a + b }
>>add(1, 3)
4

>>let sub = fn(a, b) { return a - b }
>>sub(5, 11)
-6
```

## Advanced Func
```
>>let map = fn(arr, f) { let iter = fn(arr, accumulated) { if (len(arr) == 0) { accumulated } else { iter(rest(arr), push(accumulated, f(first(arr)))); } }; iter(arr, []) };
>>let a = [1, 2, 3];
>>let double = fn(x) { x * 2 };
>>map(a, double)
[2, 4, 6]

>>let reduce = fn(arr, initial, f) { let iter = fn(arr, result) { if (len(arr) == 0) { result } else { iter(rest(arr), f(result, first(arr))); } }; iter(arr, initial) }
>>let sum = fn(arr) { reduce(arr, 0, fn(initial, el) { initial + el }); }
>>let a = [1, 2, 3];
>>sum(a)
6

>>let mean = fn(arr) { sum(arr) / len(arr) }
>>mean(a)
2
```
