# Snowflake

Snowflake 结构是一个 64bit 的 int64 类型的数据。

![pic](https://img.luozhiyun.com/20210502181302.png)

## 实现步骤

1. 获取当前的毫秒时间戳；
2. 用当前的毫秒时间戳和上次保存的时间戳进行比较；
    1. 如果和上次保存的时间戳相等，那么对序列号 sequence 加一；
    2. 如果不相等，那么直接设置 sequence 为 0 即可；
3. 然后通过或运算拼接雪花算法需要返回的 int64 返回值。
