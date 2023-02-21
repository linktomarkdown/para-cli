#!/usr/bin/env bash

check_os(){
    if [ "$(uname)" == "Darwin" ]; then
      echo "Mac OS X"
      mac_install
    elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
      echo "Linux"
      linux_install
    elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW64_NT" ]; then
      echo "Windows"
      window_install
    else
      echo "Your platform ($(uname -a)) is not supported."
      exit 1
    fi
  }

window_install(){
  echo "window install"
  # 下载安装包
  curl -o- https://github.com/linktomarkdown/para-cli/releases/download/v0.1.5/para-cli_0.1.5_windows_amd64.tar.gz | tar -xz
  # 移动到系统目录,如果已经存在，删除
  if [ -f /c/Windows/System32/para-cli.exe ]; then
    rm /c/Windows/System32/para-cli.exe
  fi
  mv para-cli.exe /c/Windows/System32
  # 创建系统变量
  setx para /c/Windows/System32/para-cli.exe
  # 查看版本
  para -v
}

mac_install(){
  echo "mac install"
  # 创建/Users/用户名/.para目录
  mkdir ~/.para
  # 下载安装包到/Users/用户名/.para目录
  cd ~/.para
  curl -o- https://github.com/linktomarkdown/para-cli/releases/download/v0.1.5/para-cli_0.1.5_darwin_arm64.tar.gz | tar -xz
  # 判断是否已经安装，如果已经安装，删除，重新创建软链接
  if [ -f /usr/local/bin/para ]; then
    rm /usr/local/bin/para
  fi
  # 创建软链接
  ln -s ~/.para/para /usr/local/bin/para
  # 给予执行权限
  chmod +x para
  # 查看版本
  para -v
}

linux_install(){
  echo "linux install"
  # 创建/usr/local/para目录
  mkdir /usr/local/para
  # 下载安装包到/usr/local/para目录
  cd /usr/local/para
  curl -o- https://github.com/linktomarkdown/para-cli/releases/download/v0.1.5/para-cli_0.1.5_linux_arm64.tar.gz | tar -xz
  # 判断是否已经安装，如果已经安装，删除，重新创建软链接
  if [ -f /usr/local/bin/para ]; then
    rm /usr/local/bin/para
  fi
  # 创建软链接
  ln -s /usr/local/para/para /usr/local/bin/para
  # 给予执行权限
  chmod +x para
  # 查看版本
  para -v
}

# 安装入口
check_os