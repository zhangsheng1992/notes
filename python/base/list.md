> list时一种可以存储任意类型的有序序列

### 定义list

使用 `[]` 来定义一个list, 如:

```python
a = [1,"a",-1,True,["a","b"]]
```

### list元素访问

使用 **下标** 来访问list,下标从0开始,注:访问一个不存在的下标会造成越界

```python
a = [1, "a", -1, True, ["a", "b"]]
print(a[0])  # 1
print(a[4])  # ['a', 'b']
print(a[5])  # index out of range
```

### list遍历

python中使用 ```for...in...```语句对list进行遍历

```python
a = [1, "a", -1, True, ["a", "b"]]

# 遍历元素
for val in a:
    print(val)

# 遍历key
for i in range(len(a)):
    print(a[i])
```

使用`len()`获取list的长度  `range(i)`表示从0开始遍历到i但并不包括i

### list截取
可以使用`list[start:end:step]`方式截取list,**start**省略表示从0开始,
**end**省略表示到最后的位置 python中下标负值如表示从尾部开始,注意是从**-1**开始的
-1表示最后一个元素  -2表示倒数第二个 以此类推
step表示步幅,即隔几个元素截取一下, demo:

```python
a = ["a", "b", "c", "d", "e", "f", "g", "h", "i"]
# 从list头(0)位置开始,截取到下标为4的元素位置(不包括这个元素)
print(a[:4])  # ['a', 'b', 'c', 'd']
# 从list下标为2的位置开始,截取到下标为4的元素位置(不包括这个元素)
print(a[2:4])  # ['c', 'd']
# 从list头开始,每隔一个元素截取一次,截取到list结尾
print(a[::2])  # ['a', 'c', 'e', 'g', 'i']
# 从list倒数第三个元素位置截取到list结尾处
print(a[-3:])  # ['g', 'h', 'i']
```
### list运算
list支持两种运算 **"+"** **"*"**
**"+"** 表示连接list,将两个list组合成一个list

```python
a = ["a", "b", "c", "d"]
b = ["e", "f", "g", "h", "i"]
print(a + b)  # 拼接 ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i']
```

**"*"** 表示重复list,即将一个list重复多次变成一个新的list
使用时只能指定int类型的非负值(负值会清空list,类似与clear())

```python
a = ["a", "b", "c", "d"]
print(a * 2)  # repeat多次 ['a', 'b', 'c', 'd', 'a', 'b', 'c', 'd']
print(a * -1)  # 清空list 等同于 clear()[]
# print(a * 1.5)  # 报错 an't multiply sequence by non-int of type 'float'
```


### list方法

**append()**

向list后端追加元素

```python
a = []
a.append("hello")
a.append("World")
print(a)  # ['Hello', 'World']
```

**pop()**

从list中取一个元素,无参数情况下默认取最后一个元素,可以指定一个参数,表示弹出元素的下标

```python
a = [1, 2, 3, 4, 5, 6, 7, 8, 9]
print(a.pop())  # 9
print(a.pop(5))  # 6
```

**count()**

统计list中某个字符的出现的次数.

```python
a = [1, 2, 3, "a", "b", "c", 1]
print(a.count("a"))  # 1
print(a.count(1))  # 2
```

**copy()**

拷贝list中内容, list为引用类型, 修改元素会修改原始引用

```python
a = [1, 2, 3, 4, 5]
print(a)  # [1, 2, 3, 4, 5]
b = a
b[0] = "aaa"
print(a)  # ['aaa', 2, 3, 4, 5]
```
如果不想修改b中元素的同时影响a的,那么需要使用**copy()**方法,
**copy()**会将原list拷贝一份,分配新的地址,对copy对象的修改不会影响原始数据

```python
a = [1, 2, 3, 4, 5]
print(a)  # [1, 2, 3, 4, 5]
b = a.copy()
b[0] = "aaa"
print(a)  # [1, 2, 3, 4, 5]
```

**clear()**

清空list中的内容

```python
a = [1, 2, 3, 4, 5]
print(a)  # [1, 2, 3, 4, 5]
a.clear()  # []
print(a)
```

**extend()**

从可迭代的list中继承元素,简单理解就是从一个list中拷贝元素并追加

```python
a = [1, 2, 3, 4, 5]
b = [6, 7, 8]
a.extend(b)
print(a)  # [1, 2, 3, 4, 5, 6, 7, 8]
```

与 **+** 号作用相同

```python
a = [1, 2, 3, 4, 5]
b = [6, 7, 8]
c = a + b
print(c)  # [1, 2, 3, 4, 5, 6, 7, 8]
```python

**index()**

查重list中某个元素的位置,如果存在返回对应下标,如果不存在,将抛出异常.

```python
a = [1, 2, 3, 4, 5]
print(a.index(2))  # 1
```

配合`try...expect...`使用

```python
a = [1, 2, 3, 4, 5]
try:
    print(a.index(8))
except Exception as err:
    print(err)  # 8 is not in list
```

此外可以通过地2,3个参数来定义查找范围

```python
a = [1, 2, 3, 4, 5, 1, 2, 3, 4, 5]
# 从list下标为4的位置开始到下标为10的位置, 查询值为2的元素的下标.
print(a.index(2, 4, 10))  # 6
```

**insert()**

在list指定位置插入元素,接收两个参数,第一个参数表示插入的位置,第二个参数为插入的元素
如果指定的位置在list中不存在,会在list后面追加元素,类似与append
如果指定的位置存在,将会插入指定位置,后续元素的位置将会依次向后移

```python
a = [1, 2, 3, 4, 5]
a.insert(7, 6)  # 下标7不存在,将类似于append()
print(a)  # [1, 2, 3, 4, 5, 6]
a.insert(2, "aaa")  # 在下标2的位置插入"aaa",后续元素依次后移
print(a)  # [1, 2, 'aaa', 3, 4, 5, 6]
```

**remove()**

移除list中指定的元素,如果元素不存在将会抛出异常

```python
a = ["a", "b", "c", "d"]
a.remove("a")
print(a)
a.remove("fff")  # 报错, x not in list 不存在, 配合try except 使用最佳
print(a)
```

**reverse()**

逆向排序整个list

```python
a = [1, 2, 3, 4, 5]
a.reverse()
print(a)  # [5, 4, 3, 2, 1]
```

**sort()**

正向排序整个list

```python
a = [5, 4, 3, 2, 1]
a.sort()
print(a)  # [1, 2, 3, 4, 5]
```

**enumerate()**

用于将一个可遍历的数据对象(如列表、元组或字符串)组合为一个索引序列,同时列出数据和数据下标,一般用在`for`循环当中


```
a = [{"a": 1}, 2, ["a", "b"], True, ("mp4", "mp3"), {"apple", "orange", "banana"}]
for key, value in enumerate(a):
    print(key, value)

# 0 {'a': 1}
# 1 2
# 2 ['a', 'b']
# 3 True
# 4 ('mp4', 'mp3')
# 5 {'banana', 'orange', 'apple'}
```
