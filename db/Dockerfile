FROM mysql:8.0

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

EXPOSE 3306

#MySQL設定ファイルをイメージ内にコピー
ADD ./my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]