>tuple与list类似,区别在于tuple中的元素无法修改

### 定义tuple
使用 `()` 来定义tuple

```python
a = (1, 2, 3, 4, 5)
```
可以把tuple看作是一个只读list,一经定义,无法修改或增加元素
```python
a = (1, 2, 3, 4, 5)
a[6] = 1  # 'tuple' object does not support item assignment
```

### tuple读取
和list类似,使用下标访问tuple

```python
a = (1, 2, 3, 4, 5)
print(a[0])
```

### tuple遍历
遍历方式也与list相同 使用 `for...in...`遍历

```python
a = (1, 2, 3, 4, 5)
for value in a:
    print(value)
# 1
# 2
# 3
# 4
# 5
```

### tuple运算
tuple运算与list基本一致
```python
a = (1, 2, 3, 4, 5)
b = (4, 5, 6)
print(a * 2)  # (1, 2, 3, 4, 5, 1, 2, 3, 4, 5)
print(a + b)  # (1, 2, 3, 4, 5, 4, 5, 6)
```

### tuple截取
tuple截取方式与list完全一致,使用`tuple[start:end:step]`来生成新的tuple

```python
a = ("a", "b", "c", "d", "e", "f", "g")
# 从下标2的位置截取到下标4的位置,不包含4这个元素
print(a[2:4])  # ('c', 'd')
# 从tuple开始截取到下标4的位置,不包含4这个元素
print(a[:4])  # ('a', 'b', 'c', 'd')
# 从tuple开始 间隔一个元素 截取到tuple结尾
print(a[::2])  # ('a', 'c', 'e', 'g')
# # 从tuple倒数第三个元素开始 截取到tuple结尾处
print(a[-3:])  # ('e', 'f', 'g')
```
**注: tuple没有运算符,字符串可以看作一种特殊的tuple**


### tuple方法

**count()**
与list的**count()**完全一致, 统计tuple中元素的数量,如果不存在会抛出异常
配合`try...except...`使用以免造成不必要的麻烦.

```python
a = (1, 2, 3, 4, 5)
print(a.count(2))  # 1
```

**index()**
与list的**index()**完全一致,返回指定元素的下标,如果不存在会抛出异常
配合`try...except...`使用以免造成不必要的麻烦.

```python
a = ("a", "b", "c", "d", "e")
print(a.index("d"))  # 3
```

**len()**
计算tuple长度
```python
a = (1, 2, 3, 4, 5)
print(len(a))  # 5
```
