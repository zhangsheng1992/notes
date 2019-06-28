> python中的类在定义时就包含了一系列的专有方法和内置属性,这些内置的属性和方法在操作类时非常便利

## 内置属性

### `__name__`
显示类名,注意和显示**module**名称的`__name__`的区别

```python
class Animal: pass


print(__name__)  # 显示模块名 __main__
print(Animal.__name__)  # 显示类名 Animal
```

### `__slots__`
python是一门动态语言,允许我们在运行时修改代码,如给对象增加属性
```python
class A:
    def __init__(self):
        self.x = 1
        self.y = 2

if __name__ == "__main__":
    a = A()
    a.z = 3
    print(a.x, a.y,a.z)
```
动态添加对程序阅读性与健壮性均有影响,内置变量`__slots__`就是为了约束这种行为

```python
class A:
    __slots__ = ("x", "y")

    def __init__(self):
        self.x = 1
        self.y = 2


if __name__ == "__main__":
    a = A()
    a.z = 3
```
如demo所示,当试图给给实例a增加一个属性`z`的时候,会抛出`AttributeError: 'A' object has no attribute 'z'`异常

### `__bases__`

显示父类信息
```python
class Animal: pass


class Life: pass


# Dog继承Animal与Life
class Dog(Animal, Life): pass


print(Dog.__bases__)  # (<class '__main__.Animal'>, <class '__main__.Life'>)
```

### `__module__`
类所在的模块,如果是直接运行脚本,值为`__main__`,如果类作为模块被导入,值为`包名.模块名`,新建一个package `human`,然后在`teacher模块中`定义一个类`Teacher`
```python
class Teacher(): pass
```
脚本中导入`teacher`模块
```python
import human.teacher

t = human.teacher.Teacher()
print(t.__module__)  # 作为模块被导入,__module__值为human.teacher
```
直接在脚本中定义一个类`Animal`
```python
class Animal: pass


print(Animal.__module__) # 作为入口使用,__module__值为__main__
```

### `__dict__`
返回一个包含类中所有属性的字典,包含类属性,实例属性,私有属性和内置属性
```python
class Animal:
    _sex = ""
    age = ""

    def __init__(self):
        self.name = ""


print(Animal.__dict__)  # name属性由于实例化时创建,所以此处是得不到的
dog = Animal()
print(dog.__dict__)  # 这样就可以取到name属性了
```

### `__mro__`
python中的类是多继承的,如果一个子类同时继承多个父类,那么访问父类中的方法,将按照MRO顺序访问,可以输出`__mro__`来看看查找顺序
```python
class A: pass

class B(A): pass

class C(A): pass

class D(A): pass

class E(B, C, D): pass

print(E.__mro__) # E,B,C,D,A
```
## 类的专有方法

### `__new__()`与`__init__()`
`__new__()` 是构造方法,用于产生实例化对象,必须返回一个对象
`__init__()` 初始化方法,负责对实例化对象进行属性值初始化,此方法必须返回None

在实例化对象时,先调用`__new__()`创建一个空对象,然后再调用`__init__()`初始化对象,`__init__()`之前介绍过,可以给实例增加实例属性

```python
class Animal:
    def __new__(cls, *args, **kwargs):
        print("生产实例化对象")
        return object.__new__(cls)

    def __init__(self):
        print("初始化对象")
        # 增加实例属性
        self.sex = ""
        self.age = ""


dog = Animal()
```
python类与**objective-c**语言很像,所有的类均继承**object**类

### `__str__()`与`__repr__()`
调用`print()`与`repr()`输出对象时会调用对应的`__str__()`与`__repr__()`,必须返回一个string对象,demo中以`__str__()`为例,`__repr__()`与其完全一致,可以鸡贼的定义`__repr__ = __str__`保持一直

```python
class Animal:

    def __str__(self):
        return "\033[1;31;32m[Info]  " + super(object, Animal).__str__()

    __repr__xxz = __str__


dog = Animal()
print(dog)
```
打印绿色字体的类信息

### `__call__()`
`__call__()`可以让类对象像函数一样执行,本质上,函数就是一个拥有__call__方法的对象

```python
class Animal:

    def __call__(self, *args, **kwargs):
        print("wang wang wang!")


dog = Animal()
dog()  # wang wang wang!
```

### `__del__`
调用`__del__`会将对象的引用计数置0, 并被垃圾回收器回收, 当程序结束调用时也会调用此方法

```python
class Animal:
    def __init__(self, name):
        self.name = name

    def __del__(self):
        print("删除对象",self.name)


dog = Animal("dog")
cat = Animal("cat")
del dog  # 调用时删除dog对象
# 程序运行结束 自动删除cat对象
```
注:python中的垃圾回收策略是**引用计数**为主, **分代收集**为辅
### `__iter__()`与`__next__()`
python中的**list**,**tuple**,**set**和**dictionary**都可以通过`iter()`转换为迭代器,然后通过`next()`一次取出其中一个元素, 迭代器其实就是通过调用内置的`__iter__()`与`__next__()`实现,而类作为一种类型,自然也可以实现迭代器,只需要实现`__iter__()`与`__next__()`即可

```python
class Animal:
    name = ["dog", "cat"]

    def __iter__(self):
        print("获取类迭代器对象")
        return self

    def __next__(self):
        if len(self.name) > 0:
            return self.name.pop()
        else:
            return StopIteration("没有更多元素")


b = iter(Animal())  # 获取类迭代器对象
print(next(b))  # cat
print(next(b))  # dog
print(next(b))  # 没有更多元素
```
**注意:**实现`__next__()`一定要判断边界


### `__getitem__()`、`__setitem__()`和`__delitem__()`
除了将类作为迭代器,也可以将类作为一个**list**或者**dictionary**来访问,只需要实现相应的方法即可
```
class Animal:
    voice = {"dog": "wang", "cat": "miao"}

    def __setitem__(self, key, value):
        print("新增动物{} 叫声:{}".format(key, value))
        self.voice[key] = value

    def __getitem__(self, item):
        if self.voice.get(item) is None:
            return KeyError("{key}不存在".format(key=key))
        else:
            print("{}的叫声是{}".format(item, self.voice[item]))
            return self.voice[item]

    def __delitem__(self, key):
        if self.voice.get(key) is None:
            return KeyError("{key}不存在".format(key=key))
        else:
            print("删除{key}成功".format(key=key))
            del self.voice[key]


a = Animal()
a["dog"]  # 访问 dog的叫声是wang
a["bridge"] = "zhi"  # 设置 新增bridge 值:zhi
del (a["cat"])  # 删除 删除cat成功
```
**注意:**实现时一定要判断key的存在,否则会造成**except**


### `__getattr__()`、`__setattr__()`与`__delattr__()`
访问、设置和删除对象属性时会触发相应方法

```python
class Animal:
    def __init__(self):
        self.age = 25
        self.sex = "公"

    def __getattr__(self, item):
        print("要访问的属性是" + item)
        return item

    def __setattr__(self, key, value):
        print("要设置的属性是{key},值为{value}".format(key=key, value=value))

    def __delattr__(self, item):
        print("要删除的属性是" + item)


singleDog = Animal()
print(singleDog.age)
del singleDog.age
```

### `__getatrribute__()`
`__getatrribute__()`是一个属性访问截断器,当访问对象的属性时,会优先触发这个方法,访问一个属性的顺序如下

1. **实例的`__getattribute__()`方法**
2. **实例对象字典**
3. **实例所在类字典**
4. **实例所在类的父类(MRO顺序）字典**
5. **实例所在类的getattr**
6. **报错**

demo示例如下
```python
class A:
    age = 10


class B(A):
    name = "狗"

    def __init__(self):
        self.sex = "公"

    def __getattribute__(self, item):
        print("1.进入__getattribute__(), 查找的属性", item)
        return super(object, B).__getattribute__(item)

    def __getattr__(self, item):
        print("5.进入__getattr__(), 查找的属性", item)
        # 6.没有找到 应该报错


b = B()
print(b.__dict__)  # 2.在实例的字典中查询
print(B.__dict__)  # 3.在实例所在类字典中给查询
print(A.__dict__)  # 4.在实例父类类字典查询
print(b.name)
```

### `__enter__()`与`__exit__()`
重写这两个方法可以允许对一个类对象使用with方法.

```python
class A:
    def __enter__(self):
        print("open文件")

    def __exit__(self, exc_type, exc_val, exc_tb):
        print("退出")


with A() as a:
    # do something
    pass
```

##  总结
从上面的方法可以看出一些共同点,在此整理一下
1. python中的类都是继承自**object**
2. 内置方法不需要主动调用,会在满足条件的时候自动调用(如果重写过的化)
3. 重写内置方法需要特别注意边界等信息,否则很容易造成**except**,影响程序运行
