>面向对象是一种编程思想,将具有相同的属性和方法的对象的抽象为一个类的,具有封装,继承,多态三个基本特征

## 类的定义
一个类由属性和方法组成,使用关键字`class`定义一个类,定义一个**Person**类:

```python
class Person: pass
```

## 实例化一个类
使用`类型()`的方式来实例一个类

```python
# 定义Person类
class Person: pass

# 实例化Person类
p = Person()
print(p)  # <__main__.Person object at 0x10fde2828>
```

## 属性
类中的属性分为**实例属性**,**类属性**,**私有属性**三种,每种属性的定义和访问略有不同

### 实例属性
在实例化类的时候定义的属性,外部访问需要使用 `实例名.实例属性名来访问`, 内部使用 `self.实例属性名`访问,一般在类的专有方法`__init__`中初始化.

```python
class Person:

    def __init__(self):
        self.age = "18"  # 定义实例属性age
        self.name = "Peter"  # 定义实例属性name

    def read(self):
        # 在类内使用self.属性名访问实例属性
        print("类内:", self.age, self.name)  # 类内: 18 Peter


# Person类的实例
p = Person()
# 在类外使用实例名.实例属性的方式访问
print("类外:", p.age, p.name)  # 类外: 18 Peter
# 调用实例方法
p.read()
```

### 类属性
没有在专有方法`__init__`中定义的属性称作类属性

```python
class Men:
    sex = "男人"

    def who(self):
        print("类内使用self访问:", self.sex)
        print("类内使用类名Men访问", Men.sex)


print("类外使用类名Men访问", Men.sex)
Peter = Men()
print("类外使用实例来访问", Peter.sex)
Peter.who()
```
类属性有4钟方式访问,无论是在类内还是类外,使用`类名.类属性名`均可以访问类属性,前提是该属性不是私有的,如果是在类外,还可以使用`实例.类属性`的方式来访问,在类内,还可以使用`self.类属性的方式访问`



### 私有属性
使用`__属性名`来定义私有属性,私有属性可以是类属性,也可以是实例属性, 但都只能在类内访问

```
class Women:
    # 私有的类属性
    __Money = "$"

    def __init__(self):
        # 私有的实例属性
        self.__age = 18

    def info(self):
        print("在类内访问私有属性:", self.__age, Women.__Money)


# print(Women.Money)  # 类外无法访问 AttributeError: type object 'Women' has no attribute 'Money'
lily = Women()
# print("lily.age)  # 类外无法访问 AttributeError 'Women' object has no attribute 'age'
lily.info()  # 在类内访问私有属性: 18 $
```

## 方法
同函数一样,在类中使用使用`def`来定义方法,区别在于方法的第一个参数必须是`self`或者`cls`,分别代表实例对象和类对象,类中的方法分为**类方法**,**实例方法**,**静态方法**,**私有方法**

### 类方法
类方法使用`@classmethod`装饰,第一个参数必须是`cls`,代表类本身,访问方式与类属性相同

```
class Animal:
    @classmethod
    def eat(cls):
        print("吃饭")

    @classmethod
    def sleep(cls):
        print("类内访问eat方法,两种方法均可", cls.eat(), Animal.eat())


# 使用类名直接访问类方法
Animal.eat()
Animal.sleep()

# 可以使用实例直接访问类方法
dog = Animal()
dog.eat()
dog.sleep()
```

### 实例方法
不使用直接使用`def`定义的方法为实例方法,第一个参数必须为`self`,代表实例本身
```python
class Animal:

    def eat(self):
        print("吃饭")

    def sleep(self):
        print("类内使用self访问eat方法", self.eat())


# 不可以使用类名直接访问类方法
# Animal.eat()
# Animal.sleep()

# 使用实例直接访问
dog = Animal()
dog.eat()
dog.sleep()
```

### 静态方法
从上述的例子可以看出,无论是类方法还是实例方法,第一个参数分别为类本身和实例本身,便于访问类或实例的其他属性方法,如果一个类方法,不需要访问实例或者类的其他属性方法,那么他也就不需要`self`或者`cls`参数这类方法被称为**静态方法**,使用`@staticmethod`装饰

```python
class Animal:
    @staticmethod
    def eat(food):
        print("吃", food)

    # 静态方法在类内既可以 使用类名.方法名调用
    # 也可以使用self或者cls调用

    def t1(self):
        self.eat("牛肉")

    @classmethod
    def t2(cls):
        cls.eat("鸡腿")


# 静态方法可以使用类名直接调用
Animal.eat("馒头")
# 也可以使用实例调用
dog = Animal()
dog.eat("骨头")
dog.t1()
Animal.t2()

```
静态方法既可以通过类名来访问,也可以通过实例来访问,可以把他理解称一个嵌在类内的函数.

## 私有方法
无论是类方法,实例方法,或者是静态方法,均可以通过在函数名前增加`__`来使方法私有,类外不能访问

```python
class Animal:
    def __t1(self):
        self.eat("牛肉")

# 类外无法访问私有方法
dog = Animal()
# dog.t1() # AttributeError: 'Animal' object has no attribute 't1'
```
