> hive使用中一些常见的错误整理

### 内存不足导致的query错误
执行一条查询大量数据集的sql `select * from xxx`时返回 `Execution Error, return code 2`,解决步骤如下

1.查看hive日志
定位包到报错部分如下:

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
