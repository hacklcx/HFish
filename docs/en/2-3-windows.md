#### Windows 3.3.0 environment installation process 

 

> Step 1: Download the[ **installation package of [HFish-Windows-amd** ](https://hfish.cn-bj.ufileos.com/hfish-3.3.0-windows-amd64.tgz)64](https://hfish.cn-bj.ufileos.com/hfish-3.1.4-windows-amd64.tgz)  (Windows x86 architecture 64-bit system), and unzip it 

 

> Step 2: Open TCP / 4433 &4434 ports on the firewall in both directions (if you need to use other services, you also need to open ports) 

 

> Step 3: Enter the folder of HFish-Windows-amd64 and run install.bat in the file directory (the script will install HFish in the current directory) 

 

> Step 4: Log in web interface

 

``` 
Login link: https://[ip]:4433/web/ 
Account: admin 
Password: HFish2021
```

 

```
If the address of the Console is 192.168.1.1, the login link shall be: https://192.168.1.1:4433/web/
```

 

Note: 



```
1. The HFish is divided into two processes. The "hfish" process is the management process, which is responsible for monitoring, launching and upgrading Honeypot master. The "console" process is the honeypot master process, which executes the honeypot software program. Therefore, when installing, ensure to execute the hfish process as required. If the console program is executed directly, the program may be unstable and the upgrade failure.
```

 

```
2. The database files of windows version of HFish are currently stored in the C:\Users\Public\hfish directory. After reinstalling HFish, HFish will automatically read the configuration and data in this directory by default".
```

 

```
3. If page access fails, check if the windows firewall releases TCP / 4433 &4434
```

