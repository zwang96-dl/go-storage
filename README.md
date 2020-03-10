# go-storage
mysql -u root -h 157.230.169.141 -p
123456

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

SET GLOBAL server_id=2;

注意！！！主从server id必须不一样！！

CREATE TABLE tbl_test (`user` varchar(64) NOT NULL, `age` int(11) NOT NULL) DEFAULT charset utf8;

INSERT INTO tbl_test (user, age) VALUES ('xiaowang', 18);


```
CREATE DATABASE fileserver DEFAULT character SET utf8;

CREATE TABLE `tbl_file` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `file_sha1` char(40) NOT NULL DEFAULT '' COMMENT 'file hash',
    `file_name` varchar(256) NOT NULL DEFAULT '' COMMENT 'file name',
    `file_size` bigint(20) DEFAULT '0' COMMENT 'file size',
    `file_addr` varchar(1024) NOT NULL DEFAULT '' COMMENT 'file storeage location',
    `create_at` datetime DEFAULT NOW() COMMENT 'create time',
    `updated_at` datetime DEFAULT NOW() ON UPDATE current_timestamp() COMMENT 'update date',
    `status` int(11) NOT NULL DEFAULT '0' COMMENT 'status available disable deleted',
    `ext1` int(11) DEFAULT '0' COMMENT 'extension 1',
    `ext2` text COMMENT 'extension 2',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_file_hash` (`file_sha1`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

SHOW CREATE TABLE tbl_file;

/data/mysql/conf/master.conf
[client]
default-character-set=utf8
[mysql]
default-character-set=utf8
[mysqld]
log_bin = log  #开启二进制日志，用于从节点的历史复制回放
collation-server = utf8_unicode_ci
init-connect='SET NAMES utf8'
character-set-server = utf8
server_id = 1  #需保证主库和从库的server_id不同， 假设主库设为1
replicate-do-db=fileserver  #需要复制的数据库名，需复制多个数据库的话则重复设置这个选项 (如果想同步所有的数据库，则直接注释这一行配置)

/data/mysql/conf/slave.conf
[client]
default-character-set=utf8
[mysql]
default-character-set=utf8
[mysqld]
log_bin = log  #开启二进制日志，用于从节点的历史复制回放
collation-server = utf8_unicode_ci
init-connect='SET NAMES utf8'
character-set-server = utf8
server_id = 2  #需保证主库和从库的server_id不同， 假设从库设为2
replicate-do-db=fileserver  #需要复制的数据库名，需复制多个数据库的话则重复设置这个选项 (如果想同步所有的数据库，则直接注释这一行配置)


mkdir -p /data/mysql/datam
docker run -d \
    --name mysql-master \
    -p 13306:3306 \
    -v /data/mysql/conf/master.conf:/etc/mysql/mysql.conf.d/mysqld.cnf \
    -v /data/mysql/datam:/var/lib/mysql  \
    -e MYSQL_ROOT_PASSWORD
    mysql:5.7


CREATE TABLE `tbl_user` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT 'Use name',
    `user_pwd` varchar(256) NOT NULL DEFAULT '' COMMENT 'User encoded password',
    `email` varchar(64) DEFAULT '' COMMENT 'Email address',
    `phone` varchar(128) DEFAULT '' COMMENT 'Phone number',
    `email_validated` tinyint(1) DEFAULT 0 COMMENT 'email validated',
    `phone_validated` tinyint(1) DEFAULT 0 COMMENT 'phone number is validated',
    `signup_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'Register date',
    `last_active` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last active timestamp',
    `profile` text COMMENT 'user attribute',
    `status` int(11) NOT NULL DEFAULT '0' COMMENT 'Account status',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_phone` (`phone`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;


CREATE TABLE `tbl_user_token` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT 'username',
    `user_token` char(40) NOT NULL DEFAULT '' COMMENT 'user login token',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `tbl_user_file` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_name` varchar(64) NOT NULL,
  `file_sha1` varchar(64) NOT NULL DEFAULT '' COMMENT '文件hash',
  `file_size` bigint(20) DEFAULT '0' COMMENT '文件大小',
  `file_name` varchar(256) NOT NULL DEFAULT '' COMMENT '文件名',
  `upload_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
  `last_update` datetime DEFAULT CURRENT_TIMESTAMP 
          ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '文件状态(0正常1已删除2禁用)',
  UNIQUE KEY `idx_user_file` (`user_name`, `file_sha1`),
  KEY `idx_status` (`status`),
  KEY `idx_user_id` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;