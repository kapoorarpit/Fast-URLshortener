# HLD Document 

# Functional Requirements
* returns shortened URL 
* re-direct to input URL 
* removes shortened URL's data from DB after some time(for example - 6 months) 
<br>
<br>

# Estimation
1. Traffic estimates     
* Let there are 10 million (10,000,000) new urls per month.
* Let read:write == 100:1 i.e 1 shortened url will redirect 100 times approximately.
* 10,000,000 request per month = 10,000,000 / (30 x 24 x 3600)m = 4 i.e 4 read operations per second and 4x 100 = 400 write operations per second.
* Hence, Servers should handle 400+4 i.e approximately 450 urls /per second (upper limit).

2. Storage estimates 
* Let every URL will expire after 6 months.
* 10 millions urls will join storage to persist for 6 months.
* hence 10 millions x 6 = 60 million URLs.
* Let 100 bytes will be consumed by 1 URL.
* 60,000,000 x 100 = 60,000,000,000 bytes.
* 60 GB storage will be required.

3. Chacheing estimates 
* Let size of cache be 1 kilobyte i.e 1000 URL at max.
* LRU cache will be best option for this structure.

# DataBase Architecture
​​​​​​No-SQL will be better choice as there will be single relational table.
Table - 
URL 	Shortened URL(primary key)	Creation Date

1. Entity table with Attributes - 
* Shortened URL as primary as - shortened URL will be unique always(according to algorithm used).
* URL to redirect it.
* Creation date to deallocate assigned resources after expiry date of the data. 

# Algorithm 
* Generate hash of a URL by giving a unique counter to new URL. 
* We can use base 100 number to get better combinations in shorter URLs.
* Let shortened URL lentgh be 6 characters.
* 256 ^ 6 shortened URL can exist at same time.
## Problem - System may curropt if there are multiple servers.
* Services like zookeper can solve this problem.
## Problem - Multiple shortened URL can exist for same URL.
* A check using <set> can solve this. 

# Major Classes 
* There are a total of four major classes: Administrators, SystemSettings, TrafficHistory, Users. 
* To GET input URL
* To shorten the URL
* To save the relational data.
* To redirect to the original URL with the help of shortened URL.
* To delete relational data after fixed period.