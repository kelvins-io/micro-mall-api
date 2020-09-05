#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import os
import sys
# py3 commands模块并入了subprocess
try:
    import commands
except ImportError:
    import subprocess as commands

project_name = os.path.basename(os.path.abspath("."))


def genpb(proto_arr):
    '''复制proto文件到工程的proto文件夹下，且执行.go, .pb, swagger.json的生成'''
    proto_pkg_path = get_proto_pkg_path()
    if not os.path.exists("./proto"):
        os.mkdir("./proto")
    for proto_dir in proto_arr:
        gen_proto(proto_pkg_path, proto_dir)
    # 在./proto目录下生成swagger.json.go
    if os.system(r'''go-bindata -ignore=\\.go -ignore=\\.proto -ignore=\\.md -o=./proto/swagger_json.go -pkg=proto ./proto/...''') != 0: # 防止转义，原始字符串
        print("GO_BINDATA ./proto/swagger.json.go ERROR!")
        sys.exit(1)


def gen_proto(proto_pkg_path, proto_dir):
    '''复制proto文件夹到工程的proto文件夹下,并生成相应的go文件'''
    # proto文件夹在工程下对应的文件夹名
    sub_dir = os.path.basename(os.path.abspath(proto_dir)).replace("-", "_")
    # proto文件夹在当前项目目录下的相对目录
    rel_pro_dir = "./proto/" + sub_dir
    # 如果文件夹不存在就创建
    if not os.path.exists(rel_pro_dir):
        os.system("mkdir %s" % rel_pro_dir)
    # 删除文件夹下所有文件
    os.system("rm -rf %s/*" % rel_pro_dir)
    # 将proto_dir下的所有文件复制到相应的地方
    os.system("cp -r %s/* %s" % (proto_dir, rel_pro_dir))
    # 获取rel_pro_dir下所有的子目录
    status, result = commands.getstatusoutput("find %s -type d" % rel_pro_dir)
    if status != 0:
        print(result)
        sys.exit(1)
    dirs = result.split("\n")
    for d in dirs:
        status, result = commands.getstatusoutput(
            "ls %s/*.proto 2>/dev/null | wc -l" % d)
        if status != 0:
            print(result)
            sys.exit(1)
        if result.strip() == "0":
            continue
        if os.system("protoc -I. -I%s --go_out=plugins=grpc:. %s/*.proto" % (proto_pkg_path, d)) != 0:
            print("PROTOC PB ERROR!!!")
            sys.exit(1)
        if os.system("protoc -I. -I%s --grpc-gateway_out=logtostderr=true:. %s/*.proto" % (proto_pkg_path, d)) != 0:
            print("PTOTOC PB.GW ERROR!!!")
            sys.exit(1)
        # 如果是项目对应的proto目录，则需要生成swagger文件
        if sub_dir == project_name.replace("-", "_") + "_proto":
            if os.system("protoc -I. -I%s --swagger_out=logtostderr=true:. %s/*.proto" % (proto_pkg_path, d)) != 0:
                print("PTOTOC SWAGGER ERROR!!!")
                sys.exit(1)


def get_proto_pkg_path():
    '''获取包含proto相关所有包的目录，优先使用vendor目录，然后使用gopath'''
    vendor = os.path.abspath("./vendor")
    if check_proto_pkg(vendor):
        return vendor
    gopath = os.environ.get("GOPATH")
    if gopath is None:
        print("未定义GOPATH环境变量")
        sys.exit(1)
    # gopath 可能有多个，unix以":"分隔，windows以";"分隔
    for path in gopath.split(":"):
        proto_pkg = path + "/src"
        if check_proto_pkg(proto_pkg):
            return proto_pkg
    print("在vendor下面和GOPATH下面都找不到proto相关目录")
    sys.exit(1)


def check_proto_pkg(base_dir):
    '''检查proto相关的目录是否都存在'''
    rel_paths = [
        "/gitee.com/kelvins-io/common/proto",
        "/gitee.com/kelvins-io/common/proto/common",
        "/gitee.com/kelvins-io/common/proto/google/api",
    ]
    for rel_path in rel_paths:
        if not os.path.exists(base_dir + rel_path):
            return False
    return True


if __name__ == "__main__":
    proto_arr = sys.argv[1:]
    if len(proto_arr) > 0:
        genpb(proto_arr)
