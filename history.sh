###
 # @Author: EnderbyEndera
 # @Date: 2021-01-05 15:36:26
 # @LastEditTime: 2021-01-05 17:23:34
 # @LastEditors: Please set LastEditors
 # @Description: In User Settings Edit
 # @FilePath: /commdetection/history.sh
### 

export HISTTIMEFORMAT="%F %T  `cat /sys/class/net/eth0/address`  `whoami`  "
export HISTSIZE=10000
export HISTFILE=~/.bash_history
history -r
history
history -c