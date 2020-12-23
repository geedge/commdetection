<!--
 * @Author: EnderByEndera
 * @Date: 2020-12-23 14:42:44
 * @LastEditTime: 2020-12-23 15:04:02
 * @LastEditors: Please set LastEditors
 * @Description: Describe the project by English
 * @FilePath: /commdetection/README.en.md
-->
# commdetection

#### Description
commdetection is a command-line tool used for Linux command detection 

#### Software Architecture
Commdetection contains the following modules
1. The cmd module, which uses cobra as a command line generation tool, provides several types of subcommands and flags to facilitate user-defined modes
2. The comm module includes Command Get and Command Flush functions to obtain all commands and clear invalid commands
3. logger module, based on logrus custom logger and middleware, middleware is conveniently provided to the gin back-end framework to customize log analysis
4. The router module, based on the gin-gonic framework, customizes the martini-like API and provides a variety of API interfaces for users to facilitate command analysis and monitoring
5. The rules module is the core of the instruction detection system. Based on user-defined rules, two modes can be selected for instruction evaluation
6. yaml module, parsing system yaml configuration file

#### Installation

1. Use `git clone https://gitee.com/EnderByEndera/commdetection` to download this project
2. Define the environment variable `export set COMMDEPATH=your local clone address`
3. Use the `go build` command

#### Instructions

1. The evaluation scores of different instructions can be customized by users, and the scores of custom instructions need to be saved in a json file (API and graphical interface will be provided in the future)
2. Two evaluation methods can be selected: command and path, or both

#### Contribution

1.  Fork the repository
2.  Create Feat_xxx branch
3.  Commit your code
4.  Create Pull Request


#### Gitee Feature

1.  You can use Readme\_XXX.md to support different languages, such as Readme\_en.md, Readme\_zh.md
2.  Gitee blog [blog.gitee.com](https://blog.gitee.com)
3.  Explore open source project [https://gitee.com/explore](https://gitee.com/explore)
4.  The most valuable open source project [GVP](https://gitee.com/gvp)
5.  The manual of Gitee [https://gitee.com/help](https://gitee.com/help)
6.  The most popular members  [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
