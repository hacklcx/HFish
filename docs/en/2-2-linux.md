### Do it by one-click installation if you are in the online environment 

> Special note: Centos is our native development and a main test system, and we recommend Centos system as the optimal option for installation. 

At present, there will be two processes after HFish is started. The "hfish" process is the management process, which is responsible for monitoring, launching and upgrading the honeypot master. The "Console" process is the honeypot master process, which executes the honeypot software program**The **Linux version HFish console database and configuration files are stored in the /usr/share/hfish directory, and the configuration and data in the directory will be automatically read during reinstallation.

If your deployment environment is Linux and has internet access. We have prepared a one-click deployment script for you to install and configure. Before using the one-click scripting, please configure the firewall. 

 

> Please open 4433 and 4434 on the firewall, and confirm to return to success (if the Honeypot Service shall occupy other ports later, you can use the same command to open it.) 

```shell
firewall-cmd --add-port=4433/tcp --permanent # (for web interface startup) 

firewall-cmd --add-port=4434/tcp --permanent # (for node and console communication ) 

firewall- cmd --reload 
```



> Using root user, run the following script.

```shell
bash <(curl -sS -L https://hfish.net/webinstall.sh)
```

> Complete the installation  

```
Login link: https://[ip]:4433/web/ 
Account: admin 
Password: HFish2021 
```

If the ip of the console is 192.168.1.1, the login link is: https://192.168.1.1:4433/web/ 

After the installation is completed, HFish will automatically create a node on the console . After logging in, you can view it in the "Node Management" list. 

 

<img src="https://hfish.net/images/image-20210914113134975.png" alt="image-20210914113134975" style="zoom: 25%;" />


### If you can't connect to the internet, it can be done by manual installation. 

 

\> **<u>Special note: Centos is our native development and main test system, we give priority to recommend you to use it </u>** 

\> **<u>Centos system is to be installed. </u>** 

If your environment does not have internet access, you can try the manual installation . 



### Step 1: Download the[installation package: [HFish-Linux-3.3.0](https://hfish.cn-bj.ufileos.com/hfish-3.1.4-linux-amd64.tgz) (Linux x86 architecture 64-bit system) 

Follow the steps below to install (take a Linux 64-bit system as an example): 

### Step 2: Create a path decompression installer 

```
mkdir hfish
```

 

### Step 3: Unzip the Installing file packages to the hfish directory 

 

```
tar zxvf hfish-3.1.4-linux-amd64.tgz -C hfish

```

#### **Step 4: Please open 4433, 4434 and 7879 on the firewall, and confirm that success is returned (if there are other services that need to open the ports, use the same command to open them).\***

 

```shell
Firewall-cmd --add-port=4433/tcp --permanent (used for web interface startup) 
firewall-cmd --add-port=4434/tcp --permanent (used for node and console communication) 
firewall-cmd - -reload 
```

 

### Step 5: Enter the installation directory and run install.sh directly 

```
cd hfish
sudo ./install.sh
```

 

### Step 6: Log in to the web interface 

```
Login link: https://[ip]:4433/web/ 
Account: admin 
Password: HFish2021
```



If the ip of the console is 192.168.1.1, the login link is: https://192.168.1.1:4433/web/ 

 

After the installation is complete, HFish will automatically create a node on the console. It can be viewed in node management. 

 

<img src="https://hfish.net/images/image-20210914113134975.png" alt="image-20210914113134975" style="zoom: 25%;" />

 

 

  

Special attention:

```
If the deployment node is on the extranet, the tcp connections may exceed the maximum connection limits (1024), causing other connections to be rejected. In this case, you can manually release the maximum number of machine tcp connections. 

 

Refer to the solution link https://www. cnblogs. com/lemon-flm/p/7975812.html
```

