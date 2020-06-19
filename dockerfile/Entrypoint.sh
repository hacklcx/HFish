#!/bin/sh

HFISH_DIR=/opt/HFish

if [ ! -d $HFISH_DIR ];then
  mv /tmp/HFish $HFISH_DIR
  sed -i "2s/status = 0/status = 1/" $HFISH_DIR/config.ini
fi

if [ ! -z "$CLUSTER_IP" ];then
  sed -i "2s/status = 1/status = 2/" $HFISH_DIR/config.ini
  sed -i "s/addr = 0.0.0.0:7879/addr = $CLUSTER_IP/" $HFISH_DIR/config.ini
fi

if [ ! -z "$NODE_NAME" ];then
  sed -i "s/name = Server/name = $NODE_NAME/" $HFISH_DIR/config.ini
fi

if [ ! -z "$USERNAME" ];then
  sed -i "s/account = admin/account = $USERNAME/" $HFISH_DIR/config.ini
fi

if [ ! -z "$PASSWORD" ];then
  sed -i "s/password = admin/password = $PASSWORD/" $HFISH_DIR/config.ini
fi

if [ ! -z "$MYSQL_USER" ] && [ ! -z "$MYSQL_PASSWORD" ] && [ ! -z "$MYSQL_IP" ] && [ ! -z "$MYSQL_PORT" ] && [ ! -z "$MYSQL_DATABASE" ];then
  sed -i "s/db_type = sqlite/db_type = mysql/" $HFISH_DIR/config.ini
  sed -i "s#^db_str = .*rwc#db_str = $MYSQL_USER:$MYSQL_PASSWORD@tcp\($MYSQL_IP:$MYSQL_PORT\)\/$MYSQL_DATABASE\?charset=utf8\&parseTime=true\&loc=Local#" $HFISH_DIR/config.ini
  if [ ! -f $HFISH_DIR/db/sql/import_sql.log ];then
    mysql -h $MYSQL_IP -P $MYSQL_PORT -u$MYSQL_USER -p$MYSQL_PASSWORD -D $MYSQL_DATABASE < $HFISH_DIR/db/sql/hfish_db.sql &&
    echo "SQL import time: `date "+%Y-%m-%d %H:%M:%S"`" > $HFISH_DIR/db/sql/import_sql.log
  fi
fi

cd $HFISH_DIR && ./HFish run
