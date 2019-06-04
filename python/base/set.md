>与list和tuple类似,集合(set)是一个序列,区别在于,集合中的元素无序且不重复

## 定义集合
集合使用`{}`或`set()`来定义,注意一点,在定义一个空集合时,需要使用`set()`,否则会被识别为`dict()`

```python
a = {1, 2, 3, 6, 2, 2, 4, 5, 6, 4, 5}
print(a)  # {1, 2, 3, 4, 5, 6}
```

b = set()
print(b)
# 增加元素
b.add(1)
print(b)
# 清空set
b.clear()
print(b)

# 判断元素是否存在
c = {1, 2, 3}
if 1 in c:
    print("have")
else:
    print("no")

# 遍历
for value in c:
    print(value)

# set运算
# a - b 集合a中包含 集合b中不包含的元素
a = {"a", "b", "c", 1, 2, 3}
b = {"a", 2}
print(a - b)  # {'c', 1, 3, 'b'}

# a|b 求并集
print(a | b)  # {1, 2, 3, 'a', 'c', 'b'}

# a&b 求交集
print(a & b)  # {2, 'a'}

# a^b 求差集
print(a ^ b)  # {'b', 1, 3, 'c'}