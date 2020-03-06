# go-storage

mysql -u root --password=May@20140515 -h 157.230.169.141

docker run --name mysqlmaster -v /Users/zwang/Documents/go-storage/mysql/mysql_db_master:/var/lib/mysql -v /Users/zwang/Documents/go-storage/mysql/masterconf.d:/etc/mysql/conf.d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql
docker run --name mysqlslave -v /Users/zwang/Documents/go-storage/mysql/mysql_db_slave:/var/lib/mysql -v /Users/zwang/Documents/go-storage/mysql/slaveconf.d:/etc/mysql/conf.d -p 3307:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql

CHANGE MASTER TO MASTER_HOST='172.17.0.2',MASTER_USER='root',MASTER_PASSWORD='123456',MASTER_LOG_FILE='binlog.000002',MASTER_LOG_POS=0;

START SLAVE;
STOP SLAVE;

SHOW SLAVE STATUS\G;
SHOW MASTER STATUS;
CREATE DATABASE test1 DEFAULT character set utf8;

https://www.jianshu.com/p/ab20e835a73f

docker inspect --format='{{.NetworkSettings.IPAddress}}' 346f2c17bd9a # mysqlslave 172.17.0.3 mysqlmaster 172.17.0.2

mysql -uroot -h127.0.0.1 -p # connect master

mysql -uroot -h127.0.0.1 -P3307 -p # connect slave

注意！！！主从server id必须不一样！！

