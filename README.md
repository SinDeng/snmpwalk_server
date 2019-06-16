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

测试命令：

```
curl -X POST http://127.0.0.1:8085/api/snmpwalk --data '{"ip":"172.16.101.1","community":"public","targetOid":"1.3.6.1.2.1.1"}'
```
