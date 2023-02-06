### Database selection options 

 

In general, MySQL database is suitable for all deployment situations and is also our preferred database. Other data processing capabilities and concurrent compatibility are better than SQLite. 

| Deployment status   | Testing status         | Stable deployment |
| ---------- | ---------------- | -------- |
| Applicable database | SQLite/MySQL | MySQL    |

 

### SQLite details 

The SQLite database is used by the HFish system by default, and the specific path of the self-contained initialized db is /usr/share/db/hfish.db 

 

### Replace SQLite with MySQL 

 

HFish currently provides the "database configuration" function, which can quickly replace the database 

![image-20211116210129137](https://hfish.net/images/image-20211116210129137.png)