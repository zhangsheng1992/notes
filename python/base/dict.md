> 字典是另一种可变容器模型，且可存储任意类型对象

### 定义dictionary
使用`{}`来定义一个字典,字典为键值对对应形式,如:

```python
a = {"a": True, "b": 4}
print(a)  # {'a': True, 'b': 4}
```

需要注意的是,字典的键必须是唯一切不可变的,即:
**number** **string** **boolean** **float** **tuple**等
如果同一个键被赋值两次,后一个值会被记住.对于值,则可以使用任何类型
```
a = {1: "a", False: "a", 1.1: "a", -1: "a", ("b", "c"): "a", "b": "a"}
print(a)  # {1: 'a', False: 'a', 1.1: 'a', -1: 'a', ('b', 'c'): 'a', 'b': 'a'}

b = {[1, 2]: "a"}
print(b)  # TypeError: unhashable type: 'list'
```

### 访问字典

通过键来访问字典值,如:
```python
a = {"a": 1, "b": 2, "c": 3, "d": 4}
print(a["a"])  # 1
print(a["b"])  # 2
print(a["c"])  # 3
print(a["d"])  # 4
```

访问一个不存在的key会导致程序抛出`KeyError`异常, 因此在不确定的key是否存在的情况下建议使用```try...except...```访问.

```python
a = {"a": 1, "b": 2, "c": 3, "d": 4}
try:
    print(a["e"])
except KeyError as e:
    print("e not found in dict.")
```

通过key来修改指定键对应的值,如:

```python
a = {"a": 1, "b": 2, "c": 3, "d": 4}
print(a["a"])  # 1
a["a"] = "hello world!"
print(a["a"])  # hello world!
```

### 内置函数

使用**del()**来删除指定键,同样的,在字典中不存在key时会抛出`KeyError`异常

```python
a = {"a": 1, "b": 2, "c": 3, "d": 4}
print(a)  # {'a': 1, 'b': 2, 'c': 3, 'd': 4}
del (a["a"])
print(a)  # {'b': 2, 'c': 3, 'd': 4}

try:
    del(a["e"])
except KeyError:
    print("e not in dictionary.")
```

使用**len()**计算字典长度

```python
a = {"a": 1, "b": 2, "c": 3}
print(len(a))  # 3
```

使用**str()**转换字典为字符串形式

```python
a = {"a": 1, "b": 2, "c": 3}
print(str(a))  # {'a': 1, 'b': 2, 'c': 3}
```

### 字典方法
**copy()**

拷贝一个字典,由于字典是引用类型,单纯赋值后对字典进行操作,会影响到原始字典,此时就需要用到拷贝

```python
a = {"a": 1, "b": 2, "c": 3, "d": 4}
print(a)  # {'a': 1, 'b': 2, 'c': 3, 'd': 4}
b = a
b["a"] = "aaaaaa"
print(a)  # {'a': 'aaaaaa', 'b': 2, 'c': 3, 'd': 4}

a = {"a": 1, "b": 2, "c": 3, "d": 4}
print(a)  # {'a': 1, 'b': 2, 'c': 3, 'd': 4}
b = a.copy()
b["a"] = "aaaaaa"
print(a)  # {'a': 1, 'b': 2, 'c': 3, 'd': 4}
```

**clear()**

清空字典

```python
a = {"a": 1, "b": 2, "c": 3, "d": 4}
a.clear()
print(a)  # {}
```

**pop()**

从字典取对应值,并从字典中删除这个键值对,接收两个参数:

第一个参数为要取值的key,如果key存在,则返回key对应的值,如果key不存在,抛出KeyError异常

第二个参数为如果key不存在时返回的默认值,如果设置了默认值并且key不存在的清空下,会返回默认值,
并且不会抛出KeyError异常

```python
a = {"a": 1, "b": 2, "c": 3, "d": 4}
print(a.pop("a"))  # 1
print(a.pop("e", "not found"))  # not found
print(a)  # {'b': 2, 'c': 3, 'd': 4}
```

**fromkeys()**

接收可迭代对象,创建一个新字典,接收两个参数:

第一个参数必须是**可迭代类型**(string list tuple set dictionary),也可以是**生成器** **迭代器**
有可迭代返回值的**函数**, 一般用此函数提取字典的key或批量重置字典中的值.

第二个参数为新元素的初始值,可以不设置

```python
a = {}
# 字符串
b = a.fromkeys("abc", "default value")
print(b)  # {'a': 'default value', 'b': 'default value', 'c': 'default value'}

# list
c = a.fromkeys([1, 2, 3], 0)
print(c)  # {1: '0', 2: '0', 3: '0'}

# tuple
d = a.fromkeys(("apple", "banana", "orange"), "")
print(d)  # {'apple': '', 'banana': '', 'orange': ''}

# set 不指定默认值
f = a.fromkeys({"cat", "dog", "fish", "bird"})
print(f)  # {'bird': None, 'fish': None, 'cat': None, 'dog': None}

# 生成器
x = (i for i in range(30) if i % 2 == 0 and i % 3 == 0)
g = a.fromkeys(x)
print(g)  # {0: None, 6: None, 12: None, 18: None, 24: None}

# 迭代器
newList = [1, 2, 3, 4]
listIter = iter(newList)
h = a.fromkeys(listIter)
print(h)  # {1: None, 2: None, 3: None, 4: None}


# 有返回值的函数
def test():
    return "key"


i = a.fromkeys(test())
print(i)  # {'k': None, 'e': None, 'y': None}

j = a.fromkeys({"a": "1", "b": "2"})
print(j)  # {'a': None, 'b': None}
```

**get()**

从字典中取指定key,接收两个参数

第一个参数为指定key

第二个参数为指定key不存在的时候返回的默认值


```python
a = {"a": 1, "b": 2, "c": 3}
print(a.get("a"))  # 1
print(a.get("d"))  # None
print(a.get("e", "not found!"))  # not found!
```

**items()**

将字典转换为一个迭代器,包含可以遍历的元组

```python
a = {"a": 1, "b": 2, "c": 3}
b = a.items()
for i in b:
    print(i)

# ('a', 1)
# ('b', 2)
# ('c', 3)
```

**keys()**

返回一个包含字典所有键的迭代器,可以转换成list()

```python
a = {"a": 1, "b": 2, "c": 3}
b = a.keys()
# 迭代
for value in b:
    print(value)

# a
# b
# c

# 使用list()将迭代器转换成list
print(list(b))
```

**popitem()**

从字典中取一个键值对,一般为最后一个,取出后字典中将不包含此键值

```python
a = {"a": 1, "b": 2, "c": 3}
b = a.popitem()
print(b)  # ('c', 3)
print(a)  # {'a': 1, 'b': 2}
```

如果字典为空,将会抛出**KeyError**异常
```python
c = {}
print(c.popitem())  # KeyError: 'popitem(): dictionary is empty'
```

**setdefault()**

向字典中插入一个key,如果key已经在字典中存在,将返回这个键对应的值,否则返回设置的默认值,接收两个参数:

第一个参数要插入的key

第二个参数设置插入key默认值, 如果key不冲突的清空下(字典中没有这个key)

```python
a = {"a": 1, "b": 2, "c": 3}
print(a.setdefault("d", 4))  # 4
print(a.setdefault("e"))  # 不知道默认值 None
print(a.setdefault("a", 123))  # a存在与字典中  返回对应值
print(a)  # {'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': None}
```

**update()**

将一个字典组合到指定字典中,接收一个字典类型参数,如果参数不是字典类型,将抛出**TypeError**异常

```python
a = {"a": 1, "b": 2, "c": 3}
b = {"d": 4, "e": 5}
a.update(b)
print(a)  # {'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5}
```

**values()**
返回一个包含字典所有值的迭代器,可以转换成list()

```python
a = {"a": 1, "b": 2, "c": 3}
print(a.values())  # dict_values([1, 2, 3])
# 遍历
for i in a.values():
    print(i)
# 1
# 2
# 3
```

