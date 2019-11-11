## 安装

Docker Machine 可以在多种操作系统平台上安装，包括 Linux、macOS，以及 Windows。

### macOS、Windows

`Docker Desktop for Mac/Windows` 自带 `docker-machine` 二进制包，安装之后即可使用。

查看版本信息。

```bash
$ docker-machine -v
docker-machine version 0.16.1, build cce350d7
```

### Linux

在 Linux 上的也安装十分简单，从 [官方 GitHub Release](https://github.com/docker/machine/releases) 处直接下载编译好的二进制文件即可。

例如，在 Linux 64 位系统上直接下载对应的二进制包。

```bash
$ sudo curl -L https://github.com/docker/machine/releases/download/v0.16.1/docker-machine-`uname -s`-`uname -m` > /usr/local/bin/docker-machine

$ sudo chmod +x /usr/local/bin/docker-machine
```

完成后，查看版本信息。

```bash
$ docker-machine -v
docker-machine version 0.16.1, build cce350d7
```
