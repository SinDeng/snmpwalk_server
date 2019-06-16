# SnmpWalk Server

## API使用

POST格式，json包括 ip community targetOid 字段。  

返回json，是一个数字对应String的KV表，数字是walk结果。比如查询结果：  

```
1.1 a
1.2 b
```

targetOid传1，json返回：
```
{"1":"a","2":"b"}
```
