### Overview of Docker Versions 

Docker is one of our recommended deployment methods. The current version has the following features: 

\- Automatic upgrade: request the latest image every hour to upgrade, and the upgrade will not lose data. 

\- Data persistence: create a data directory under the /usr/share/hfish directory of the host to store attack data, and create a logs directory to store logs. 

Note: The current Docker version uses host mode to start. If you do not want the Docker console, open ports other than 4433 and 4434, you can suspend the local node. 

### Default installation instructions on Docker 

#### Install the console in Docker: 

\> Step 1: Confirm that Docker is installed and started 

```
docker version 
```

\> Step 2: Run the version HFish (copy all in the box, paste and execute) 

```
docker run -itd --name hfish \
-v /usr/share/hfish:/usr/share/hfish \
--network host \
--privileged=true \
threatbook/hfish-server:latest
```

<img src="https://hfish.net/images/4351638188574_.pic_hd.jpg" alt="4351638188574_.pic_hd" style="zoom:50%;" /> 

\> Step 3: Configure subsequent automatic upgrades (copy all in the box, paste and execute) 

```
Docker run -d   \
 --name watchtower \
 --restart unless-stopped \
 -v /var/run/docker.sock:/var/run/docker.sock  \
 --label=com.centurylinklabs.watchtower.enable=false \
--privileged=true \
 containrrr/watchtower  \
 --cleanup  \
 hfish \
 --interval 3600
```

![4381638189986_.pic_hd](https://hfish.net/images/4381638189986_.pic_hd.jpg)

 

\> Step 4: Login to HFish 

Login address: https://ip:4433/web/ 

Initial username: admin 

Initial password: HFish2021 

### Upgrade Failure of Docker 

If you configure the Docker image proxy, it may result in the failure of watchower. Here you can manually execute:  

```
docker pull threatbook/hfish-server:3.1.4  
docker tag threatbook/hfish-server:3.1.4  threatbook/hfish-server:latest  
docker rm -f hfish  
docker run -itd --name hfish \
-v /usr/share/hfish:/usr/share/hfish \
--network host \
--privileged=true \
threatbook/hfish-server:latest
```



### Without automatic upgrade configured, there are single manual upgrade instructions of Docker. 

\> Step 1: configure watchover (copy, paste and execute all in the box) 

```
docker run -d   \
 --name watchtower \
 --restart unless-stopped \
 -v /var/run/docker.sock:/var/run/docker.sock  \
 --label=com.centurylinklabs.watchtower.enable=false \
--privileged=true \
 containrrr/watchtower  \
 --cleanup  \
 hfish \
 --interval 10
```

 

\> Step 2: after the upgrade is successful, log in to the page and confirm that the upgrade is complete 

\> Step 3: Cancel watchover auto-upgrade 

```
docker stop watchtower
```

 

After configuration, if you upgrade regularly, you can continue to manually upgrade the hfish image through docker start watchtower and docker stop watchtower 



### Docker restart instructions after modifying persistent configuration 

\> Step 1: modify the configuration under usr/share/hfish/config. toml 

\> Step 2: restart the docker container 

```
docker restart hfish
```

 

 

 

 

 

 