>itertools提供了高效的循环和迭代函数集合,
>此模块中的所有函数返回的迭代器都可以与for循环语句以及其他包含迭代器的函数联合使用

## 无限迭代器

### count()
`itertools.count()`会创建一个无限迭代器,格式 `count(start=0,step=1)`,一般常用来计数等操作.
* start 表示开始位置
* step 表示递增的幅度

```python
import itertools

it = itertools.count(1, 5)
j = 0
for i in it:
    print(i, end=",")
    if j > 10:
        break
    j += 1
```
结果: 1,6,11,16,21,26,31,36,41,46,51,56

### cycle()
`itertools.cycle(iter)`创建一个迭代器,当迭代完成全部对象后,将重新开始.
```python
import itertools

it = itertools.cycle([1, 2, 3])
j = 0
for i in it:
    print(i, end=",")
    if j > 10:
        break
    j += 1
```
结果: 1,2,3,1,2,3,1,2,3,1,2,3,

### islice()
`itertools.cycle(iter,stop)`创建一个迭代器,当满足条件时停止,stop代表迭代次数

```python
import itertools

it = itertools.islice([1, 2, 3, 4, 5], 4)
for i in it:
    print(i, end=",")

```
结果: 1,2,3,4,

## 有限迭代器

### chain
**chain**类允许将多个**可迭代对象**组合,每次调用**__next__**方法时输出一个元素

```python
import itertools

t = itertools.chain("ABC", [1, 2, 3], (4, 5, 6), {"a": 1, "b": 2}, [x for x in range(0, 10, 2)])
for i in t:
    print(i, end="")
```
输出结果:ABC123456ab02468

### chain.from_iterable(iterables)
接收一个迭代对象**iterables**,生成一个迭代序列,等同于一下函数

```python
def from_iterable(iterables):
    for i in iterables:
        for j in i:
            yield j
```
一般用来展开一个**dict**或者二维的**list()**

```python
for i in itertools.chain.from_iterable([["a", "b"], ["c", "d"], ["e", "f"]]):
    print(i, end="")
```
输出结果:abcdef

### accumulate
返回累计求和的结果,比如需要计算数字1-10的和,demo如下:

```python
import itertools

it = itertools.accumulate(range(1, 11))
for i in it:
    print(i, end=",")
```
结果: 1,3,6,10,15,21,28,36,45,55