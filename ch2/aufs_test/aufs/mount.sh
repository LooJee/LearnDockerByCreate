#!/usr/bin/env bash

# mount aufs 命令，默认第一个挂载的目录是可读写权限，后面的目录都是只读权限
sudo mount -t aufs -o dirs=./container-layer:./image-layer4:./image-layer3:./image-layer2:./image-layer1 none ./mnt

# 挂载好后，系统的aufs目录会生成对应的si_**目录，使用一下目录可以查看对应的信息，将si_**替换成本地的目录
cat /sys/fs/aufs/si_dd7e172a86ad78a5/*

# 往 image-layer4.txt 追加数据
echo -e "\nwrite to mnt's image-layer1.txt" >> ./mnt/image-layer4.txt

# 查看源路径下的 image-layer4.txt，发现数据没有发生改变
cat image-layer4/image-layer4.txt

# 但是挂载的 image-layer4.txt的数据发生了改变
cat mnt/image-layer4.txt

# 查看 container-layer 目录，发现多了一个 image-layer.txt 文件
ls container-layer/

# 查看 container-layer/image-layer4.txt 里的数据，发现是追加了之后的数据，这是后发生了 AUFS 的写时拷贝
cat container-layer/image-layer4.txt
