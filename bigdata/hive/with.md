>with...as...也叫做子查询部分,语句允许hive定义一个sql片段,供整个sql使用


## 简介
`with...as...`需要定义一个sql片段,会将这个片段产生的结果集保存在内存中,
后续的sql均可以访问这个结果集和,作用与视图或临时表类似.


## 语法限制
1. `with...as...`必须和其他sql一起使用(可以定义一个`with`但在后续语句中不使用他)
2. `with...as...`是一次性的

`with...as...`的完整格式是这样的

```sql
-- with table_name as(子查询语句) 其他sql 
with temp as (
    select * from xxx
)
select * from temp;
```

只定义不实用
```sql
with temp as (
    select * from xxx
)
select * from othertable;
```

同级的多个temp之间用`,`分割`with`只需要一次,`as`后的子句必须用`()`,

```sql
with temp1 as (
    select * from xxx
),temp2 as (
    select * from xxx
)
select * from temp1,temp2;
```

`with...as...`当然是可以嵌套的,此处举一个简单例子

```sql
with temp2 as (
    with temp1 as (
        select * from xxx
    )
    select * from temp1
)
select * from temp2;
```

`with...as...`只能在一条sql中使用

```sql
with temp1 as (
    select * from xxx
)
select * from temp1;
select xxx from temp1; -- error! no table named temp1;
```

## 语句的优点
1. 提高代码可读性(结构清晰)
2. 简化sql,优化执行速度(`with`子句只需要执行一次)

### 栗子
现有 **city** 表,结构如下:

| city_number   |  city_name  | province |
|:----------:|:--------:|:--------:|
| 010 | 北京 | 北京 |
| 021 | 上海 | 上海 |
| 025 | 南京 | 江苏 |
| 0512 | 昆山 | 江苏 |
| 0531 | 济南 | 山东 |
| 0533 | 淄博 | 山东 |

然后分别有商品表**good**

| city_number   |  good_name  |
|:----------:|:--------:|
| 010 | A |
| 021 | B |

现在需要分别统计这上海商品,一般sql如下:
```sql
select * from `good`  where city_number in (select city_number from city where city_name = "上海");
```
除了子查询,上述的的例子还可以用`join`来实现,

如果用**with...as...**语句实现,如下
```sql
with tmp_shanghai as(
    select city_number from city where city_name = "上海"
)
select * from `good` where tmp_shanghai in (select * from tmp_shanghai) 
```
看起来使用 **with...as...** 语句反而更复杂一点,但如果**tmp_shanghai**要被多次使用的使用,就很有必要

来看一个实际的例子,有一张操作表**event**主要字段如下:

| date   |  event_key  |
|:----------:|:--------:|
| 20190530 | Delete |
| 20190530 | Put |
| 20190530 | Get |
| 20190530 | Get |
| 20190601 | Set |
......

现在要求一条sql统计出`Get`与`Set` 操作的数量,先使用子查询实现
```sql
select (
    select count(*) from event where event_key = "Get"
) as get_num,(
    select count(*) from event where event_key = "Set"
) as set_num
```
如果再增加其他项的统计呢,是否每一个均需要增加一个对**event**表进行扫描的自查询

使用 **with...as...**
```sql
with temp as(
    select * from event where event_key = "Get" or event_key = "Set"
)
select 
    sum(case when event_key = "Get" then 1 else 0 end) as get_num,
    sum(case when event_key = "Set" then 1 else 0 end) as Set_num
from temp
```
阅读性是否比之前有所提高?此外,这条语句只对event表进行了一次扫描,将符合条件的数据存入temp中供后续计算,
在event表数据集非常大的情况下,性能将比子查询的方式优秀很多