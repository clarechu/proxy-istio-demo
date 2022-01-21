# 安装nfs

## 前言

NFS 是 Network File System 的缩写,即网络文件系统。功能是让客户端通过网络访问不同主机上磁盘里的数据,主要用在类Unix系统上实现文件共享的一种方法。 本例演示 CentOS 7 下安装和配置 NFS 的基本步骤。

## 环境说明

CentOS 7（Minimal Install）

```bash
$ cat /etc/redhat-release
CentOS Linux release 7.5.1804 (Core)
```

## 本例演示环境如下

|Name | IP Addr      |        Descprition |
| ------------ |--------------| ------------ | 
|Server| 10.10.13.125 |    服务端 IP|
|Client| 10.10.13.118 |    客户端 IP|

根据官网说明 Chapter 8. Network File System (NFS) - Red Hat Customer Portal,
CentOS 7.4 以后,支持 NFS v4.2 不需要 rpcbind 了,
但是如果客户端只支持 NFC v3 则需要 rpcbind 这个服务。

## 挂载磁盘

```bash

# 查看多余的磁盘

$ fdisk -l 

磁盘 /dev/sdb：214.7 GB, 214748364800 字节,419430400 个扇区
Units = 扇区 of 1 * 512 = 512 bytes
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节


磁盘 /dev/sda：214.7 GB, 214748364800 字节,419430400 个扇区
Units = 扇区 of 1 * 512 = 512 bytes
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节
磁盘标签类型：dos
磁盘标识符：0x000052fd

   设备 Boot      Start         End      Blocks   Id  System
/dev/sda1   *        2048     2099199     1048576   83  Linux
/dev/sda2         2099200   209715199   103808000   8e  Linux LVM

磁盘 /dev/mapper/centos-root：53.7 GB, 53687091200 字节,104857600 个扇区
Units = 扇区 of 1 * 512 = 512 bytes
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节


磁盘 /dev/mapper/centos-swap：8455 MB, 8455716864 字节,16515072 个扇区
Units = 扇区 of 1 * 512 = 512 bytes
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节


磁盘 /dev/mapper/centos-home：44.1 GB, 44149243904 字节,86228992 个扇区
Units = 扇区 of 1 * 512 = 512 bytes
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节

# 格式化磁盘  

$ mkfs.ext4 /dev/sdb


# 挂载磁盘

$ mount /dev/sdb /nfs/data

# 查看磁盘的uid

$ blkid

/dev/mapper/centos-root: UUID="a3d00a85-4d1b-4587-8470-b1ae75370bdb" TYPE="xfs"
/dev/sda2: UUID="EQlih9-djWE-hK12-MALH-ffix-9hHG-HsSh44" TYPE="LVM2_member"
/dev/sdb: UUID="614164ac-db0e-47c7-8c41-2ffaca617e3f" TYPE="ext4"
/dev/sda1: UUID="c44cf072-e06b-4919-b19b-ec72eb6d4a37" TYPE="xfs"
/dev/mapper/centos-swap: UUID="cda5e626-53e3-4900-b3bd-a6042fcc41c0" TYPE="swap"
/dev/mapper/centos-home: UUID="2c9a7d50-4f02-4c74-8088-ebb9f6164d24" TYPE="xfs"

# 永久挂载磁盘添加以下磁盘目录

$  vim /etc/fstab

UUID="614164ac-db0e-47c7-8c41-2ffaca617e3f" /nfs/data ext4 defaults 0 0


```


## 服务端

服务端安装
使用 yum 安装 NFS 安装包。

```bash
$ sudo yum install nfs-utils
```

`注意`

| 只安装 nfs-utils 即可,rpcbind 属于它的依赖,也会安装上。

### 服务端配置

设置 NFS 服务开机启动
```bash
$ sudo systemctl enable rpcbind
$ sudo systemctl enable nfs
```
启动 NFS 服务

```bash
$ sudo systemctl start rpcbind
$ sudo systemctl start nfs
```
防火墙需要打开 rpc-bind 和 nfs 的服务

```bash
$ sudo firewall-cmd --zone=public --permanent --add-service={rpc-bind,mountd,nfs}
success
$ sudo firewall-cmd --reload
success
```

### 配置共享目录

服务启动之后,我们在服务端配置一个共享目录

```bash
$ sudo mkdir /data
$ sudo chmod 755 /data
```
根据这个目录,相应配置导出目录

```bash
$ sudo vi /etc/exports
# 添加如下配置
/data/     10.10.13.125/24(rw,sync,no_root_squash,no_all_squash)
```

* /data: 共享目录位置。

* 192.168.0.0/24: 客户端 IP 范围,* 代表所有,即没有限制。
* rw: 权限设置,可读可写。
* sync: 同步共享目录。
* no_root_squash: 可以使用 root 授权。
* no_all_squash: 可以使用普通用户授权。



保存设置,重启 NFS 服务。

```bash
$ sudo systemctl restart nfs
```

可以检查一下本地的共享目录

```bash
$ showmount -e localhost

Export list for localhost:
/data 192.168.0.0/24
```
这样,服务端就配置好了,接下来配置客户端,连接服务端,使用共享目录。


## 测试helm
```bash
helm install --dry-run --debug nfs --generate-name --set nfs.server=10.10.13.125    
```
## 参考文档

[CentOS 7 下 yum 安装和配置 NFS](https://qizhanming.com/blog/2018/08/08/how-to-install-nfs-on-centos-7)