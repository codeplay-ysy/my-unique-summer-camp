#Linux操作
##Linux基本命令
###ls:(list)
* ls -a 显示所有文件，包括隐藏文件
* ls -l 显示文件的详细信息
* ls -d 显示目录本身的详细信息
* ls -h 以人类可读的方式显示文件大小
* ls -i 显示文件的inode(inode是文件系统中的文件索引，每个文件都有唯一的inode号)
* ls -R 递归显示目录下的所有文件
* ls -t 按照文件修改时间排序
* ls -S 按照文件大小排序
* ls -r 逆序显示
* ls -1 每行显示一个文件
###cd:(change directory)
* cd - 返回上一次所在的目录
* cd .. 返回上一级目录
* cd ~ 返回当前用户的home目录
* cd / 返回根目录
###pwd:(print working directory)
* pwd 显示当前所在的目录
* pwd -P 显示当前所在的目录的真实路径
* pwd -L 显示当前所在的目录的逻辑路径(逻辑路径是指通过链接进入的路径)
###mkdir:(make directory)
* mkdir -p 创建多级目录
* mkdir -m 创建目录并设置权限(例如：mkdir -m 777 test)
* mkdir -v 显示创建的目录信息
###rm:(remove)
* rm -f 强制删除文件
* rm -r 递归删除目录
* rm -i 删除前提示(例如：rm -i test,前提示是：rm: remove regular empty file 'test'?)
* rm -v 显示删除的文件信息
* rm -rf 强制递归删除目录
###cp:(copy)
* cp -a 复制文件，保留源文件的属性(例如：cp -a test test1,复制test文件为test1,保留了文件的所有者、组、权限、时间等)
* cp -d 复制链接文件
* cp -f 强制复制(无视目标文件是否存在以及权限等因素)
* cp -i 复制前提示
* cp -l 复制文件为硬链接文件(硬链接文件是指多个文件指向同一个inode,软连接是指多个文件指向同一个文件名)
###mv:(move)
* mv -f 强制移动
* mv -i 移动前提示
* mv -u 只有源文件比目标文件新或者目标文件不存在时才移动
* mv -v 显示移动的文件信息
* mv -b 覆盖前备份
* mv -n 不覆盖已存在的文件
###touch:
* touch -a 修改文件的访问时间(访问时间是指最后一次读取文件的时间)
* touch -c 如果文件不存在，不创建文件
###find:
* find -name 按照文件名查找文件
* find -type 按照文件类型查找文件
* find -size 按照文件大小查找文件(例如：find -size +1M,查找大于1M的文件,find -size -1M,查找小于1M的文件)
* find -user 按照文件属主查找文件
* find -group 按照文件属组查找文件
* find -perm 按照文件权限查找文件(例如：find -perm 777,查找权限为777的文件)
* find -mtime 按照文件修改时间查找文件(例如：find -mtime +1,查找修改时间大于1天的文件,find -mtime -1,查找修改时间小于1天的文件)
* find -atime 按照文件访问时间查找文件
* find -ctime 按照文件创建时间查找文件
##文件权限
###chmod:(change mode)
* chmod -R 递归修改文件权限
* chmod -c 修改前提示
* chmod -f 强制修改
* chmod -v 显示修改的文件信息
* chmod -x 删除执行权限
* chmod -w 删除写权限
* chmod -r 删除读权限
###chown:(change owner)(同chmod)
###粘滞位：
* 是指在目录的权限中的最后一位，当粘滞位为1时，只有文件的所有者和root用户才能删除文件，其他用户只能查看文件内容，不能删除文件。比如/tmp目录的权限为drwxrwxrwt,其中最后一位就是粘滞位。前面依次是文件类型(d)、所有者权限(rwx)、组权限(rwx)、其他用户权限(rwx)。rwx分别代表读、写、执行权限，当某一位为1时，表示有该权限，为0时表示没有该权限。还可以用1、2、4来表示读、写、执行权限，例如：rwx=1+2+4=7,表示有读、写、执行权限。r-x=4+0+1=5,表示有读、执行权限。粘滞位可以避免别人删除自己创建的文件，尽管对方可能拥有修改权限。
###隐藏权限：
常见的隐藏权限：
* i权限：只能查看目录或文件内容，不可以删除和修改目录及文件内容，包括root用户也不可以。
* a权限：可以查看目录和文件内容，也可以追加文件内容，也可以在目录中创建新子目录和文件，但不可以删除目录和文件的内容，当然也不可以修改只能追加。
* 查看隐藏权限
命令语法：lsattr  文件或目录
##什么是进程
* 进程(process)是指正在运行的程序，每个进程都有一个唯一的进程号(PID)，进程号是一个非负整数，进程号为0的进程是调度进程，进程号为1的进程是init进程，init进程是所有进程的父进程。
* 进程的几个特征:
  * 1.动态性：进程是一个动态的概念，是程序的一次执行过程，是系统进行资源分配和调度的一个独立单位。
  * 2.并发性：多个进程可以同时运行，每个进程都可以独立运行，互不影响。
  * 3.独立性:进程是系统进行资源分配和调度的一个独立单位，进程之间不会相互影响。
  * 4.异步性:进程是异步执行的，进程按各自独立的、不可预知的速度向前推进。
  * 5.结构性:进程是由程序、数据和进程控制块三部分组成。
* Linux进程结构:
  可由三部分组成：代码段、数据段、堆栈段。也就是程序、数据、进程控制块PCB（Process Control Block）组成。进程控制块是进程存在的惟一标识，系统通过PCB的存在而感知进程的存在。
* 进程的状态:进程的状态有5种，分别是运行、就绪、阻塞、创建、终止。
  * 运行状态:进程正在运行。
  * 就绪状态:进程已经准备好运行，正在等待CPU调度。
  * 阻塞状态:进程正在等待某个事件的发生，例如等待用户输入、等待磁盘IO等。
  * 创建状态:进程正在被创建。
  * 终止状态:进程已经终止。
##进程操作
###ps:(process status)
* ps -a(-e) 显示所有进程
* ps -u 显示指定用户的进程
* ps -x 显示没有控制终端的进程
* ps -f 显示进程的详细信息
* ps -l 显示进程的详细信息(与-f的区别是，-l显示的信息更详细)
###top:(table of processes)
* top 动态显示进程信息
* top -d 指定刷新时间
* top -p 指定进程号
* top -u 指定用户
* top -c 显示完整的命令
* top -b 以批处理模式运行
* top -n 指定刷新次数
> top显示的cpu状态中:
> us(user)表示运行用户进程的CPU时间
> sy(system)表示运行内核进程的CPU时间
> ni(niced)表示运行用户进程的优先级
> wa(wait)表示等待IO的CPU时间
> hi(hardware interrupt)表示硬件中断的CPU时间
> si(system interrupt)表示软件中断的CPU时间
> st(steal time)表示虚拟机偷取的CPU时间
###kill:
* kill -l 显示所有信号
* kill -s 发送信号(还可以写成kill -信号名)
* kill -1(-HUP) 重启进程
* kill -9 强制杀死进程
* kill -15 正常杀死进程
* kill匹配PID
###killall:
* killall -u 指定用户
* killall -e name 精确匹配进程名
* killall -i name 交互式
* killall -I name 忽略进程名大小写
* kill匹配进程名
###pkill:
* pkill -u 指定用户
* pkill -n(--newest) 按照进程的启动时间排序，杀死最新的进程
* pkill -o(--oldest) 按照进程的启动时间排序，杀死最旧的进程
* pkill -v(--inverse) 杀死不匹配的进程
* pkill -t(--terminal) 按照终端号杀死进程
* pkill用于杀死某类或某个用户的进程
##端口
端口是指计算机与外部通信的接口，端口号是一个16位的无符号整数，范围是0-65535。端口号分为三类，分别是公认端口、注册端口和动态端口。公认端口的范围是0-1023，注册端口的范围是1024-49151，动态端口的范围是49152-65535。公认端口是指某些服务的默认端口，例如80端口是http服务的默认端口，22端口是ssh服务的默认端口。注册端口是指某些服务的默认端口，但是并不是所有的系统都会使用这些端口，例如mysql的默认端口是3306，但是并不是所有的系统都会使用3306端口。动态端口是指系统自动分配的端口，范围是49152-65535。
* TCP端口和UDP端口:由于TCP和UDP 两个协议是独立的，因此各自的端口号也相互独立，比如TCP有235端口，UDP也 可以有235端口，两者并不冲突
##端口查看:
* netstat -a 显示所有端口
* netstat -t 显示tcp端口
* netstat -u 显示udp端口
* netstat -l 显示监听状态的端口
* netstat -p 显示进程信息
* 还可以用lsof -i:端口号查看指定端口的进程信息
* fuser -v 端口号/协议 显示指定端口的进程信息
##Systemd守护进程
这一部分是根据阮一峰的教程学习的，不在此记录
##terminal和shell
terminal过去是指一台计算机，由于主机设备昂贵，所以通过多个终端串口连接到主机上，用于输入和输出数据，降低成本。而shell是一个命令解释器，用于解释用户输入的命令，然后执行命令。现在的terminal是指一个模拟终端，用于输入和输出数据，而shell是指一个命令解释器，用于解释用户输入的命令，然后执行命令。zsh，bash，fish等都是shell的一种。
##防火墙
 Arch的防火墙由 netfilter 和 iptables 两个组件组成
  ### iptables
* iptables 通过创建不同的过滤链来处理数据包。常见的过滤链有：INPUT、FORWARD 和 OUTPUT。INPUT 链用于处理目标地址是本机的数据包，FORWARD 链用于处理目标地址不是本机的数据包，OUTPUT 链用于处理源地址是本机的数据包。当数据包到达时，iptables 会根据数据包的目标地址和目标端口来查找相应的过滤链，然后按照过滤链中的规则来处理数据包。如果数据包没有匹配任何规则，iptables 就会按照默认策略来处理数据包。默认策略一般是丢弃（DROP）或放行（ACCEPT）。
#### filter表:
* filter 表用于过滤数据包，它包含三个过滤链：INPUT、FORWARD 和 OUTPUT。filter 表是默认的表，如果没有指定表，iptables 就会使用 filter 表。
#### nat表:
* nat 表用于修改数据包的源地址和目标地址，它包含三个过滤链：PREROUTING、OUTPUT 和 POSTROUTING。PREROUTING 链用于修改数据包的目标地址，OUTPUT 链用于修改数据包的源地址，POSTROUTING 链用于修改数据包的源地址和目标地址。
#### mangle表:
* mangle 表用于修改数据包的 IP 头部，它包含五个过滤链：PREROUTING、INPUT、FORWARD、OUTPUT 和 POSTROUTING。PREROUTING 链用于修改数据包的 IP 头部，INPUT 链用于修改数据包的 IP 头部，FORWARD 链用于修改数据包的 IP 头部，OUTPUT 链用于修改数据包的 IP 头部，POSTROUTING 链用于修改数据包的 IP 头部。
#### raw表:
* raw 表用于配置数据包的连接追踪机制，它包含两个过滤链：PREROUTING 和 OUTPUT。PREROUTING 链用于配置数据包的连接追踪机制，OUTPUT 链用于配置数据包的连接追踪机制。