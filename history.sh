###
 # @Author: EnderbyEndera
 # @Date: 2021-01-05 15:36:26
 # @LastEditTime: 2021-01-05 18:15:26
 # @LastEditors: Please set LastEditors
 # @Description: history.sh is used in the commget.go to get history commands in the host with the defined format
 # @FilePath: /commdetection/history.sh
### 

export HISTTIMEFORMAT="%F %T  `cat /sys/class/net/eth0/address`  `whoami`  "
export HISTSIZE=10000
export HISTFILE=~/.bash_history
history -r
history
history -c