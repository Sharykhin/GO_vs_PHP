Author: Siarhei Sharykhin

there are two branch: php and golang
checkout to neccessary branch and run tests

Results:
-------

php: 5.5
Go: 1.4
MySQL: 5.5

Machine:
i3 3.30 GHzx4
15.4 RAM

2000 users in 10 second
2 cycle



10 todos in database (in common way, with pagination of 10 items):

```sh
php: 
	average: 9800/msec
	throughput: 91,1/sec
go:
	avegar: 9/msec
	throughput: 350/sec
```

1000 todos in one page:

```sh
php:
	average: 60622
	throughput: 19.0/sec

go:
	average: 10049
	throughput: 90.3/sec
```




