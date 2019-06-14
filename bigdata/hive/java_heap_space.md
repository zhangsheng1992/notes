[08S01][2] Error while processing statement: FAILED: Execution Error, return code 2 from org.apache.hadoop.hive.ql.exec.mr.MapRedTask

2019-06-12T17:49:15,906  WARN [Thread-3212] mapred.LocalJobRunner: job_local114462102_0010
java.lang.Exception: java.lang.OutOfMemoryError: Java heap space
	at org.apache.hadoop.mapred.LocalJobRunner$Job.runTasks(LocalJobRunner.java:492) ~[hadoop-mapreduce-client-common-3.1.2.jar:?]
	at org.apache.hadoop.mapred.LocalJobRunner$Job.run(LocalJobRunner.java:552) ~[hadoop-mapreduce-client-common-3.1.2.jar:?]
Caused by: java.lang.OutOfMemoryError: Java heap space

在hive命令行中!env 看一下

vim $HIVE_HOME/config/hive-env.sh

# Larger heap size may be required when running queries over large number of files or partitions.
# By default hive shell scripts use a heap size of 256 (MB).  Larger heap size would also be
# appropriate for hive server.

默认是256MB 修改大一些
export HADOOP_HEAPSIZE=2048
