#!/bin/sh

HFISH_DIR=/opt/HFish

if [ ! -d $HFISH_DIR ];then
  mv /tmp/HFish $HFISH_DIR
  sed -i "s/status = 0/status = 1/g" $HFISH_DIR/config.ini
  sed -i "s/127.0.0.1/0.0.0.0/g" $HFISH_DIR/config.ini
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

cd $HFISH_DIR && ./HFish run

