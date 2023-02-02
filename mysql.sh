ARG PWD=E:/dockerMySql/root/mysql
docker run -p 3306:3306 --name my-mysql --network slinknet -v $PWD/conf:/etc/mysql -v $PWD/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root -d mysql:latest
docker exec -it my-mysql bashgo
mysql -u root -p
"root"
GRANT ALL ON *.* TO 'root'@'%';
flush privileges;
CREATE USER 'link'@'%' IDENTIFIED BY 'link';
CREATE DATABASE slink;
GRANT ALL ON *.* TO 'link'@'%';
flush privileges;