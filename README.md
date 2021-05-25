# zkl
Simple Zookeeper distribute lock

## start zookeeper container
`docker-compose -f stack.yml up`

Run:
```
zookeeper-test$ go run main.go 
check if root dir exist: /lockTest
2021/05/24 23:06:28 connected to 127.0.0.1:2181
2021/05/24 23:06:28 authenticated: id=72069276631171101, timeout=10000
2021/05/24 23:06:28 re-submitting `0` credentials after reconnect
2021/05/24 23:06:28 connected to 127.0.0.1:2181
2021/05/24 23:06:28 connected to 127.0.0.1:2181
C02FL1TFMD6N:zookeeper-test zixuan.xu$ go run main.go 
check if root dir exist: /lockTest
2021/05/24 23:07:43 connected to 127.0.0.1:2181
2021/05/24 23:07:43 authenticated: id=72069276631171105, timeout=10000
2021/05/24 23:07:43 re-submitting `0` credentials after reconnect
C02FL1TFMD6N:zookeeper-test zixuan.xu$ go run main.go 
check if root dir exist: /lockTest
2021/05/24 23:08:24 connected to 127.0.0.1:2181
2021/05/24 23:08:24 authenticated: id=72069276631171109, timeout=10000
C02FL1TFMD6N:zookeeper-test zixuan.xu$ go run main.go 
check if root dir exist: /lockTest
2021/05/24 23:10:10 connected to 127.0.0.1:2181
2021/05/24 23:10:10 authenticated: id=72069276631171113, timeout=10000
2021/05/24 23:10:10 re-submitting `0` credentials after reconnect
2021/05/24 23:10:10 connected to 127.0.0.1:2181
2021/05/24 23:10:10 connected to 127.0.0.1:2181
2021/05/24 23:10:10 connected to 127.0.0.1:2181
2021/05/24 23:10:10 authenticated: id=72069276631171114, timeout=10000
2021/05/24 23:10:10 re-submitting `0` credentials after reconnect
2021/05/24 23:10:10 authenticated: id=72069276631171115, timeout=10000
2021/05/24 23:10:10 re-submitting `0` credentials after reconnect
2021/05/24 23:10:10 authenticated: id=72069276631171116, timeout=10000
2021/05/24 23:10:10 re-submitting `0` credentials after reconnect
CreateLock: /lockTest/0000000013
CreateLock: /lockTest/0000000014
CreateLock: /lockTest/0000000015
Child of rootPath: 0000000013 0000000014 0000000015
/lockTest/0000000013 acquire lock success
Release lock: /lockTest/0000000013
Child of rootPath: 0000000013 0000000014 0000000015
start monitoring path: /lockTest/0000000013
Child of rootPath: 0000000013 0000000014 0000000015
start monitoring path: /lockTest/0000000014
/lockTest/0000000014 Monitor node goes offline, get lock
Release lock: /lockTest/0000000014
/lockTest/0000000015 Monitor node goes offline, get lock
Release lock: /lockTest/0000000015
3
Stop...
```
