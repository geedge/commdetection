<!--
 * @Author: your name
 * @Date: 2020-12-23 14:42:44
 * @LastEditTime: 2021-01-05 18:18:06
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/README.md
-->
# commdetection

#### 介绍
commdetection是基于命令行的指令检测系统
commdetection拥有多种分析方式和指令评估模式，用户也可以自定义指令评估分数和指令评估策略

#### 软件架构
commdetection包含以下几个模块
1. cmd模块，使用了cobra作为命令行生成工具，提供了几类子命令和flags以方便用户自定义模式
2. comm模块，包含Command Get和Command Flush功能，以获得和清除无效指令
3. logger模块，基于logrus自定义logger和middleware，middleware方便提供给gin后端框架以自定义日志分析
4. model模块，存放各类model
5. router模块，基于gin-gonic框架，自定义martini-like API，提供多种API接口给用户，方便用户进行命令分析和监控
6. rules模块，指令检测系统核心，基于用户自定义规则，可选择两种模式进行指令评估
7. yaml模块，解析系统yaml配置文件


#### 安装教程

1.  使用`git clone https://gitee.com/EnderByEndera/commdetection`下载本项目
2.  定义环境变量`export set COMMDEPATH=本机克隆地址`（最好能够放置于.bashrc或/etc/profile中并source）
3.  使用`go build`命令

#### 使用说明

1.  不同指令的评估分数可以由用户自定义，自定义指令分数需要创建一个json文件保存（未来会提供API和图形化界面）
2.  可选择两种评估方式：command和path，也可以同时选择

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
