​
#LLD Document 

#System architecture

![2](img/3.png)

#Methods Involved

1. "ShortenURL"   
    * GenerateURL(URL)
    * SetURL(shortenedURL)

2. "ErrorRedirect"
    * returnErr()
      
3. "Redirect"
    * GetURL(shortenedURL)
    * Redirect(URL)

4. "RemoveExpiredData"
    * Collect_Data()
    * RemoveEntries()
​