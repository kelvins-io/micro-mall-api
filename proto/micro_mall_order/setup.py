#!/usr/bin/env python3
# -*- coding:utf-8 -*-
import os
import sys
import json

project_name = os.path.basename(os.path.abspath("."))


def get_git_branch():
    '''GO_ENV 环境变量配置了默认拉取的git分支'''
    go_env = os.environ.get("GO_ENV")
    if go_env is None:
        print("未配置GO_ENV环境变量")
        sys.exit(1)
    return "master" if go_env == "prod" else go_env


def get_proto_arr():
    '''获取服务依赖和pb命令'''
    try:
        with open("/usr/local/etc/global-conf/config.json") as f:
            conf = json.load(f)
            before_build = conf.get(project_name).get("before_build")
            pb_cmd = before_build[:2]
            proto_arr = before_build[2:]
            return pb_cmd, proto_arr
    except IOError:
        print("open /usr/local/etc/global-conf/config.json failed, please check!")
        sys.exit(1)


def git_pull(proto_arr):
    '''拉取git仓库'''
    branch = get_git_branch()
    project_abs_path = os.path.abspath(".")
    proto_abspath_arr = [os.path.abspath(
        proto_rel_path) for proto_rel_path in proto_arr]
    for proto_abs_path in proto_abspath_arr:
        cd_pull(proto_abs_path, branch)
    cd_pull(project_abs_path, branch)


def cd_pull(abspath, branch):
    if not os.path.exists(abspath):
        print("directory or file %s not exists!" % abspath)
        sys.exit(1)
    # cd abspath
    os.chdir(abspath)
    # 切换分支不存在就创建分支并切换
    if os.system("git checkout %s" % branch) != 0:
        if os.system("git checkout -b %s" % branch) != 0:
            print("git checkout -b %s failed!" % branch)
            sys.exit(1)
    if os.system("git pull origin %s" % branch) != 0:
        print("git pull origin %s failed!" % branch)
        sys.exit(1)


def deploy():
    '''部署服务'''
    if os.system("go build -o %s" % project_name) != 0:
        print("go build -o %s failed!" % project_name)
        sys.exit(1)
    # 服务进程可以已经存在，需先kill调服务进程
    os.system("ps -ef | grep %s | grep -v setup.sh | grep -v grep | cut -c '9-15' | xargs kill 2>/dev/null" % project_name)
    if os.system("nohup ./%s server &" % project_name) != 0:
        print("start process failed")
        sys.exit(1)


def genpb(pb_cmd, proto_arr):
    cmd = " ".join(pb_cmd) + " " + " ".join(proto_arr)
    if os.system(cmd) != 0:
        print("exec %s failed!" % cmd)
        sys.exit(1)


if __name__ == "__main__":
    # 获取依赖的pb仓库目录
    pb_cmd, proto_arr = get_proto_arr()
    # 拉取项目和其依赖的pb仓库的分支代码
    git_pull(proto_arr)
    # 生成pb文件
    genpb(pb_cmd, proto_arr)
    # 部署
    deploy()
