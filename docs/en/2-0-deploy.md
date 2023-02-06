#### Deployment hosts supported by HFish 

HFish adopts B/S architecture. The system consists of a console and a node terminal. The console is used to generate and manage node terminal, and receive, analyze and display the data returned from the node terminal. The node terminal accepts the control of the console and is responsible for building the honeypot service. 

|                        | Windows                   | Linux X86                 |
| ---------------------- | ------------------------- | ------------------------- |
| Console (Server)       | 64-bit support            | 64-bit support            |
| Node terminal (Client) | 64-bit and 32-bit support | 64-bit and 32-bit support |

 

#### Configuration required for HFish intranet 

\> Generally speaking, honeypots deployed on the internal network have lower performance requirements, and honeypots connected to public network will have greater performance requirements. 

For the past test situation, we give two configurations. Note that if your honeypot is deployed on the Internet, it will suffer from large attack traffic. It is recommended to upgrade the configuration of the host. 

|                           | Console          | Node terminal   |
| ------------------------- | ---------------- | --------------- |
| Recommended configuration | 2 core 4 g 200 G | 1 core 2 g 50 G |
| Minimum configuration     | 1 core 2 g 100 G | 1 core 1 g 50 G |

 

Note: The occupation of log disks is greatly affected by the number of attacks. It is recommended to configure the console with more than 200 G of hard disk space. 

#### Configuration required for HFish extranet (mysql database must be replaced) 

\> Generally speaking, honeypots connected to public network will have greater performance requirements. 

For the past test situation, we give two configurations. Note that if your honeypot is deployed on the Internet, it will suffer from large attack traffic. It is recommended to upgrade the configuration of the host. 

|                           | Console (mysql database must be replaced) | Node terminal  |
| ------------------------- | ----------------------------------------- | -------------- |
| Recommended configuration | Within 5 nodes, 4-core 8 g200 G.          | 1 core 2 g50 G |
| Minimum configuration     | 2 core 4 g100 G                           | 1 core 1 g50 G |

 

Note: The occupation of log disks is greatly affected by the number of attacks. It is recommended to configure the console with more than 200 G of hard disk space. 

#### Deployment permission requirements 

\> Requirement for root privileges on the Console 

\1. If you use the install.sh script recommended by the official website to install,  you need root permission, and the installation directory will be located in the opt directory; 

\2. If you download the installation package and manual installation , in the case of using the SQLite database by default, the deployment and use of the Console does not require root privileges, but if you want to replace SQLite with MySQL data, MySQL installation and configuration requires root privileges; 

\> Requirement for root privileges on the node terminal 

The installation and operation of the node terminal do not require root permission. However, due to the limitation of the operating system, the node running with non root permission cannot monitor ports lower than tcp/1024; 