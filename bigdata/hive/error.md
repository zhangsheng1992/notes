> hive使用中一些常见的错误整理

### 内存不足导致的query错误
执行一条查询大量数据集的sql `select * from xxx`时返回 `Execution Error, return code 2`,解决步骤如下

1.查看hive日志,定位到报错部分如下:

```text
[08S01][2] Error while processing statement: FAILED: Execution Error, return code 2 from org.apache.hadoop.hive.ql.exec.mr.MapRedTask

2019-06-12T17:49:15,906  WARN [Thread-3212] mapred.LocalJobRunner: job_local114462102_0010
java.lang.Exception: java.lang.OutOfMemoryError: Java heap space
	at org.apache.hadoop.mapred.LocalJobRunner$Job.runTasks(LocalJobRunner.java:492) ~[hadoop-mapreduce-client-common-3.1.2.jar:?]
	at org.apache.hadoop.mapred.LocalJobRunner$Job.run(LocalJobRunner.java:552) ~[hadoop-mapreduce-client-common-3.1.2.jar:?]
Caused by: java.lang.OutOfMemoryError: Java heap space
```
日志中提示`Java heap space`, 堆内存不足

2.登陆hive
在命令行中`!env`查看一下配置,发现配置项

```text
HADOOP_HEAPSIZE=256m
```

3.修改配置

```text
vim $HIVE_HOME/config/hive-env.sh
```

找到`export HADOOP_HEAPSIZE`这项,这里的注释也提示如果需要执行大数据集的query,需要增大该项配置

```text
# Larger heap size may be required when running queries over large number of files or partitions.
# By default hive shell scripts use a heap size of 256 (MB).  Larger heap size would also be
# appropriate for hive server.
```
默认是256MB,修改大一些,重启,问题解决
```
export HADOOP_HEAPSIZE=2048
```


### mysql连接池不足导致连接失败

执行sql返回`Connection is not available, request timed out after 30001ms`,排查步骤如下

1.查看hive日志,定位到报错部分如下:
```text
ERROR [pool-6-thread-151] metastore.RetryingHMSHandler: Retrying HMSHandler after 2000 ms (attempt 8 of 10) with error: javax.jdo.JDODataStoreException: HikariPool-1 - Connection is not available, request timed out after 30001ms.
	at org.datanucleus.api.jdo.NucleusJDOHelper.getJDOExceptionForNucleusException(NucleusJDOHelper.java:543)
	at org.datanucleus.api.jdo.JDOQuery.executeInternal(JDOQuery.java:391)
	at org.datanucleus.api.jdo.JDOQuery.execute(JDOQuery.java:216)
	at org.apache.hadoop.hive.metastore.ObjectStore.getCurrentNotificationEventId(ObjectStore.java:9608)
	at sun.reflect.GeneratedMethodAccessor24.invoke(Unknown Source)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:498)
	at org.apache.hadoop.hive.metastore.RawStoreProxy.invoke(RawStoreProxy.java:97)
	at com.sun.proxy.$Proxy25.getCurrentNotificationEventId(Unknown Source)
	at org.apache.hadoop.hive.metastore.HiveMetaStore$HMSHandler.get_current_notificationEventId(HiveMetaStore.java:7485)
	at sun.reflect.GeneratedMethodAccessor23.invoke(Unknown Source)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:498)
	at org.apache.hadoop.hive.metastore.RetryingHMSHandler.invokeInternal(RetryingHMSHandler.java:147)
	at org.apache.hadoop.hive.metastore.RetryingHMSHandler.invoke(RetryingHMSHandler.java:108)
	at com.sun.proxy.$Proxy26.get_current_notificationEventId(Unknown Source)
	at org.apache.hadoop.hive.metastore.api.ThriftHiveMetastore$Processor$get_current_notificationEventId.getResult(ThriftHiveMetastore.java:18364)
	at org.apache.hadoop.hive.metastore.api.ThriftHiveMetastore$Processor$get_current_notificationEventId.getResult(ThriftHiveMetastore.java:18349)
	at org.apache.thrift.ProcessFunction.process(ProcessFunction.java:39)
	at org.apache.hadoop.hive.metastore.TUGIBasedProcessor$1.run(TUGIBasedProcessor.java:111)
	at org.apache.hadoop.hive.metastore.TUGIBasedProcessor$1.run(TUGIBasedProcessor.java:107)
	at java.security.AccessController.doPrivileged(Native Method)
	at javax.security.auth.Subject.doAs(Subject.java:422)
	at org.apache.hadoop.security.UserGroupInformation.doAs(UserGroupInformation.java:1729)
	at org.apache.hadoop.hive.metastore.TUGIBasedProcessor.process(TUGIBasedProcessor.java:119)
	at org.apache.thrift.server.TThreadPoolServer$WorkerProcess.run(TThreadPoolServer.java:286)
	at java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1149)
	at java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:624)
	at java.lang.Thread.run(Thread.java:748)
NestedThrowablesStackTrace:
......
```
对所有hive元数据和分区的访问都要通过**Hive Metastore**,元数据库没有使用自带的**derby**,使用了**mysql**,因此考虑可能是**mysql**出了问题


2.登陆mysql
使用`show full processlist`查看mysql连接信息

```text
+-------+------------+----------------------+--------+---------+-------+-------+-----------------------+
| Id    | User       | Host                 | db     | Command | Time  | State | Info                  |
+-------+------------+----------------------+--------+---------+-------+-------+-----------------------+
| 41954 | zhangsheng | 220.249.22.178:55664 | hive   | Sleep   |    11 |       | NULL                  |
| 41955 | zhangsheng | 220.249.22.178:55689 | hive   | Sleep   |    11 |       | NULL                  |
| 41956 | zhangsheng | 220.249.22.178:55704 | hive   | Sleep   |  1747 |       | NULL                  |
| 41960 | zhangsheng | 220.249.22.178:55751 | hive   | Sleep   |  1732 |       | NULL                  |                
| .............
+-------+------------+----------------------+--------+---------+-------+-------+-----------------------+
```
发现sleep中的连接数过多

再次查看一下连接异常和失败信息,发现客户端未正确关闭和中断连接的次数都很多

```text
show status like "%Aborted_clients%";
+-----------------+-------+
| Variable_name   | Value |
+-----------------+-------+
| Aborted_clients | 26005 |
+-----------------+-------+
show status like "%Aborted_connects%";
+------------------+-------+
| Variable_name    | Value |
+------------------+-------+
| Aborted_connects | 5091  |
+------------------+-------+
```

查看一下最大连接数与响应的连接数,最大使用的连接数已经达到最大连接限制的百分之80

```text
show variables like 'max_connections';
+-----------------+-------+
| Variable_name   | Value |
+-----------------+-------+
| max_connections | 151   |
+-----------------+-------+
show status like 'max_used_connections';
+----------------------+-------+
| Variable_name        | Value |
+----------------------+-------+
| max_used_connections | 122   |
+----------------------+-------+
```
查看一下超时设置

```text
show variables like '%timeout%';
+-----------------------------+----------+
| Variable_name               | Value    |
+-----------------------------+----------+
| connect_timeout             | 10       |
| ............
| interactive_timeout         | 28800    |
| ............
| wait_timeout                | 28800    |
+-----------------------------+----------+
```

3.解决

首先减少超时时间,让空闲的连接尽快释放,使用**root**账号执行

```text
set wait_timeout=3600;
set interactive_timeout 3600;
```
**interactive_timeout**与**wait_timeout**两个参数只有一个起作用,
和用户连接时指定的连接参数相关,缺省情况下是使用**wait_timeout**,建议都设置

然后增大mysql的最大连接数和
```text
set global max_connections=200;
```

4.注意
MySQL服务器允许的最大连接数16384,MySQL无论如何都会保留一个用于管理员(SUPER)登陆的连接,用于管理员连接数据库进行维护操作,即使当前连接数已经达到了max_connections.
因此MySQL的实际最大可连接数为**max_connections+1**
此外,通过set设置的配置在数据库重启后会失败,如果要永久,修改**my.ini**中并重启mysql