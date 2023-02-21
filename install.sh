#!/usr/bin/env bash

check_os(){
    if [ "$(uname)" == "Darwin" ]; then
      echo "Mac OS X"
      mac_install
    elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
      echo "Linux"
      linux_install
    elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW32_NT" ]; then
      echo "Windows"
      window_install
    else
      echo "Your platform ($(uname -a)) is not supported."
      exit 1
    fi
  }

window_install(){
  echo "window install"
  # 下载安装包 https://github.com/linktomarkdown/para-cli/releases/download/v0.1.0/para_windows_amd64.exe 并重命名为para.exe，放到C:\Windows\System32目录下，然后在cmd中输入para -v，如果出现版本号，说明安装成功
  # 下载安装包
  wget -O para.exe https://github.com/linktomarkdown/para-cli/releases/download/v0.1.0/para_windows_amd64.exe
  # 移动到系统目录
  mv para.exe /c/Windows/System32
  # 创建系统变量
  setx para /c/Windows/System32/para.exe
  # 查看版本
  para -v
}

mac_install(){
  echo "mac install"
  # 创建/Users/用户名/.para目录
  mkdir ~/.para
  # 下载安装包到/Users/用户名/.para目录
  cd ~/.para
  wget -O para https://github.com/linktomarkdown/para-cli/releases/download/v0.1.0/para_darwin_amd64
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
  wget -O para https://github.com/linktomarkdown/para-cli/releases/download/v0.1.0/para_linux_amd64
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