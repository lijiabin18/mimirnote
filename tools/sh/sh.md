> 日志

- 清理日志

```shell
# 查看某个目录使用情况
#du -h --max-depth=1 /var/log

# 删除脚本，使用/dev/null 来回收
cat /dev/null > /var/log/alternatives.log
cat /dev/null > /var/log/auth.log
cat /dev/null > /var/log/boot.log
cat /dev/null > /var/log/btmp
cat /dev/null > /var/log/daemon.log
cat /dev/null > /var/log/debug
cat /dev/null > /var/log/dpkg.log
cat /dev/null > /var/log/kern.log
cat /dev/null > /var/log/logfile
cat /dev/null > /var/log/amessages
cat /dev/null > /var/log/php7.0-fpm.log
cat /dev/null > /var/log/shairport.err
cat /dev/null > /var/log/shairport.log
cat /dev/null > /var/log/syslog
cat /dev/null > /var/log/user.log
cat /dev/null > /var/log/wtmp
cat /dev/null > /var/log/Xorg.0.log
cat /dev/null > /var/log/Xorg.1.log
cat /dev/null > /var/log/Xorg.0.log.old
cat /dev/null > /var/log/Xorg.1.log.old
cat /dev/null > /var/log/auth.log.1
cat /dev/null > /var/log/auth.log.2.gz
cat /dev/null > /var/log/auth.log.3.gz

cat /dev/null > /var/log/dmesg.1.gz
cat /dev/null > /var/log/dmesg.2.gz
cat /dev/null > /var/log/dmesg.3.gz
cat /dev/null > /var/log/dmesg.4.gz

cat /dev/null > /var/log/kern.log.1
cat /dev/null > /var/log/kern.log.2.gz
cat /dev/null > /var/log/kern.log.3.gz


cat /dev/null > /var/log/syslog.1
cat /dev/null > /var/log/syslog.2.gz
cat /dev/null > /var/log/syslog.3.gz
cat /dev/null > /var/log/syslog.4.gz
cat /dev/null > /var/log/syslog.5.gz
cat /dev/null > /var/log/syslog.6.gz
cat /dev/null > /var/log/syslog.7.gz

# 备选 find /var/log -mtime +30 -name "\*.gz" -exec rm -rf {} \;
```
