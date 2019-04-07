## Linux
### 1.文件管理
1. pwd 查看当前路径
2. shell: 命令解析器.默认运行在终端当中的程序 (unix--born)
3. bash(born again shell): linux版的shell
4. date: 日期
5. 文件: linux系统中所见皆文件
6. 清屏: Ctrl + l  或者 clear
7. 用户: linux是一款多用户多任务的分时复用操作系统
8. 家目录 宿主目录 cd回到家目录
9. history: 执行过的命令记录
10. ls -l: 获取文件详情的描述
  ```shell
    ls -a 查看所有文件,包括隐藏文件,以"."开头的文件
    ls -d 查看目录属性
    ls test[0-9].? -l > out 将test的文件写入到out文件中
    ls test[0-9].? -l >> out 将test的文件追加到out文件中
    ls -lh 以人类的方式显示
    man 手册
  ```
11. mkfifo name: 创建管道文件
```shell 
  -rw-r--r--. 1 root root 47 2月  21 19:31 vm.sh
  文件属性,硬链接计数,文件所属用户,文件所属用户组,文件所占存储空间大小(字节),文件最后一次修改时间, 文件名
  文件属性: 10个字符,分为两组,
        第一组: 第一个字符,代表文件类型,7种类型,不以文件后缀名作为区分文件类型的依据
              普通文件: -
              目录文件: d
              软连接: l (相当于windows快捷键)
              字符设备文件: c
              块设备文件: b
              套接字文件: s
              管道文件: p
              unknown文件
        第一组: 9个字符,分为3组,文件所有者,文件所属组,其他,r(读),w(写),x(执行)
  硬链接计数: 
        概念: 表示硬链接数有几个
        创建硬链接: ln 旧文件 新文件
        特征: 文件和硬链接文件之间,除文件名不一样以外,其他信息完全一致,并能实时同步
  文件所属用户:
        谁创建默认属于谁
  文件所属用户组:
        谁创建默认文件所属用户所在的用户组
  文件所占存储空间大小: 文件: 时间大小.目录:4k的整数倍
  文件最后一次修改时间: 初创文件: 时间创建时间,修改过的时间: 修改时间
  文件名
  ```
12. 重定向
```shell
  >: 重定向到指定文件,文件不存在就自动创建
  >>: 文件不覆盖,直接追加 
```
13. 分屏显示  
  more less cat tac 
  head -5 xx.txt 显示前5条
  tail -5 xx.txt 显示最后5条
14. 创建目录
  mkdir
  -p: 递归一次性创建多个目录 mkdir -p a/b/c/d
15. 删除文件
  rm -r  递归删除目录子内容
16. ln命令
  硬链接: ln 源文件 连接文件 不能给目录创建
  软连接: ln -s 源文件 连接文件 相当于快捷键
17. 目录拷贝
  cp 源文件 目标位置 (目录位置/文件名,可以改名)
  cp -r 递归拷贝目录内容
  cp -a 递归拷贝目录内容,并保留属性
18. mv
  mv 源文件 目标位置
19. 压缩
  ```shell
  gzip, bzip2,tar
  tar
    压缩: tar -zcvf xxx.tar.gz 
    解压: tar -zxvf xxx.tar.gz
    z: gzip格式
    c: 创建压缩文件
    x: 解压压缩文件
    v: 输出压缩详情
    f: 指定压缩后的文件名
  zip: 
    压缩 zip -r 压缩包名(没有.zip后缀)
    解压: unzip -d 解压目录 
  ```
20. 切换用户目录
  su 用户名 切换用户,不改变工作目录  
  su - 用户名 切换用户,改变工作目录  
  sudo: 零时获取root权限  
21. 添加新用户
  sudo adduser 用户名  
  sudo deluser 用户名  
  chown 新用户名 待修改用户的文件名  
  sudo addgroup 新组名  
  sudo group 组名  
  chgrp 新组名 带修改组名的文件名  
22. chmod
  修改文件访问权限
  r,w,x: 4/2/1
  示例: r-x-w-r--: 4+1=5 2 4
  1. 数字修改权限 chmod 777 ./123.txt 用户读写执行权限
  2. 加号减号修改权限 chmod +x ./123.txt 添加执行权限 r读权限, w写权限
23. w,r,x对于目录,文件的含义是否相同?
  r: 查看文件内容  查看目录的内容(目录项)
  w: 修改,删除文件内容 目录的内容可以被删除,修改,增加
  x: 该文件可以执行   如果没有执行权限,目录不能进入

### 2.系统管理
1. ps进程命令
  ps aux |grep xxx
2. top进程查看
3. kill 种子进程  
  kill -l 查看所有的选项,9表示强行终止    
  kill -9 id    
4. cat & 暂停的含义,前台到后台
5. fg 将暂停的启动
  fg 恢复优先级最高  
  fg 1 选择恢复的  
6. jobs 查看暂停的程序
7. 关机重启
  reboot 重启  
  shutdown -h now 关机 // 现在关机  
  shutdown -h 20:25 // 20:25关机  
  shutdown -h +10 // 加10分钟关机  
  init 0 // 关机  
  init 9 // 重启  
8. ifconfig
  ifconfig ens33 192.168.1.251 // 临时设置ip,重启失效  
  ifconfig ens33 down // 关闭网络  
  ifconfig ens33 up // 启动网络  
  eth0 网口  
  link encap 链路封装  
9. 虚拟机网络设备  
  桥接模式: 路由器给虚拟Linux单独分配一个ip地址,与windows所在同一网段  
  NAT模式: linux虚拟机借助windows网口,访问外网,linux和Windows公用访问外网的IP

10. find 查找文件
  格式: find 带搜索目录 参数 "关键字"  
  * -name  
    find ./ -name "*.tar" 在目录中查找.tar结尾的文件  
  * -type: f(普通文件),d,l,c,b,p,s  
    find ./ -type "f"  
  * -size 小写的k,大写的M  
    find ./ -size 1000 搜索文件大小1000kb的  
    find ./ -size +1k -size -5M  
  * -maxdepth 1 访问目录的层数,防止在其他参数的之前  
    find ./ -maxdepth 1 -type "f"  
  * 命令可以一起使用  
    find ./ -maxdepth 1 -type "f" -name "*.png"  
  * -exec 对结果加执行  
    find ./ -maxdepth 1 -type "f" -name "*.png" -exec ls -l {} \;  
  * -| xargs ls -l  
    find ./ -maxdepth 1 -type "f" -name "*.png" | xargs ls -l  
11. grep 查找内容
  grep -r/-R 'love' 文件目录, r递归执行  
  结合find,xargs,grep  
  find -maxdepth 1 -type "f" -print() | xargs -0 grep "love" -n (n表示行号) 
### 目录结构:
  * /bin 可执行文件,绿色的是可执行二进制文件的目录,ls,bash,ping
  * boot 放在系统启动是用到一些文件,linux内核文件:/boot/vmlinuz,系统引导管理器:/boot/grub
  * dev 鼠标,应用设备的文件,入挂载的光驱 mount /dev/cdrom .mnt
  * etc 系统配置文件存放的目录,
  * home 系统默认的用户目录,新增用户账号时,用户的家目录都存放在此目录下
  * lib 系统使用的函数库的目录,
  * lost+fount: 系统异常产生错误是,会将一些遗失的片段放在此目录下
  * /mnt: 光盘默认挂载点,通常挂载在 /mnt/cdrom下,也可以挂载在其他地方
  * opt: 给主机额外安装软件所摆放的目录
  * proc: 此目录的数据都在每次中,如系统核心,外部设备,网络状态,不占磁盘空间,/proc/cupinfo./proc/interrupts,/proc/dma,/proc/ioports,/proc/net
  * root: 系统管理员root的家目录
  * sbin: 放置系统管理员使用的可执行命令


### 远程登录
1. ssh 
  * ssh -l user 127.0.0.1
  * exit
2. telnet
3. 文件传输scp
  scp -r ./test.txt user@127.0.0.1:/home/
  scp -r 源目标 用户名@IP地址:目标位置