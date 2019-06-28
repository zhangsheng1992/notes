>使用hive进行实际分析时,数据集往往非常大,因此经常需要从原始数据集提取数据到临时表(中间表)

## 情景
假设现在有一张学生表**student**,包含四个字段**name,sex,grade,other**,结构如下:

| name   |  sex  |  grade | other |
|:----------:|:-------------:|:------:|:------:|
| zhangsan | 男 | 一年级 |...|
| lisi | 女 | 二年级  |...|
| wangwu | 女 | 三年级 |...|


假设这张表的记录数非常大(这里不讨论为什么不分区分桶),
我们需要将学生按照grade和性别生成临时表,比如table1(一年级的男生),table2(一年级的女生)...
然后再分别对每个table中的other信息进行分析

## 生成临时表
一般常见的方式如下:

```sql
insert overwrite table table1 select from student where sex="男" and grade="一年级"
insert overwrite table table2 select from student where sex="女" and grade="一年级"
......
```

这种方式语法是正确的,但是效率过于低下,因每一次**insert**时都需要去扫描一遍student

## 优化
hive提供了`from table`语句,在达到同样目的的时候,只扫描一遍表

```sql
from student
insert overwrite table table1 select * where sex="男" and grade="一年级"
insert overwrite table table2 select * where sex="女" and grade="一年级"
```