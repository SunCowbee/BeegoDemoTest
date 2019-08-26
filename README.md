# BeegoDemoTest

ENVIRONMENT REQUIRED:

#ubuntu

#go

#beego

$ go get -u -v github.com/astaxie/beego

$ go get -u -v github.com/beego/bee

##配置beego执行文件环境变量：$GOPATH/bin

	$ vim .bashrc
	
	//在最后一行插入
	
	export PATH="$GOPATH/bin:$PATH"
	
	//然后保存退出
	
	$ source .bashrc
	
#mysql

#mysql-driver

go get -u -v github.com/go-sql-driver/mysql


#database

mysql -uroot -proot

drop database test;

create database test charset=utf8;
