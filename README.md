# time-calculator

## Install

```shell
go install github.com/golfz/time-calculator@latest
```

## How to use

### Sum command

#### Without flag, sum from command arguments

Run:

```shell
time-calculator sum 01:03 14:45
```

result:

```shell
00:15:48
```

#### With -f flag, sum from a file

Example file: ([time.txt](./time.txt))

```text
13:57
18:00
9:20
17:26
13:57
8:53
8:03
1:26
14:04
4:00
9:46
2:28
7:20
10:42
8:17
6:52
9:49
1:11
21:25
18:15
6:35
10:24
8:00
6:40
8:46
9:39
4:37
12:20
10:36
```

Run:

```shell
time-calculator sum -f time.txt
```

result:

```shell
04:42:48
```