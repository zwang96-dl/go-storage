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