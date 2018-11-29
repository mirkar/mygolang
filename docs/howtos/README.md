# How To

#### Set proxy when executing “go get” command

In Windows, if you have configured proxy for connecting Internet, then executing “go get” command may get error


The solution is set “http_proxy” and/or “https_proxy” environment variables according to your reality need:

```text
C:\Users\xiaona>set https_proxy=https://web-proxy.corp.xx.com:8080/

C:\Users\xiaona>set http_proxy=https://web-proxy.corp.xx.com:8080/

C:\Users\xiaona>go get -v gopkg.in/fatih/pool.v2
Fetching https://gopkg.in/fatih/pool.v2?go-get=1
Parsing meta tags from https://gopkg.in/fatih/pool.v2?go-get=1 (status code 200)

get "gopkg.in/fatih/pool.v2": found meta tag main.metaImport{Prefix:"gopkg.in/fa
tih/pool.v2", VCS:"git", RepoRoot:"https://gopkg.in/fatih/pool.v2"} at https://g
opkg.in/fatih/pool.v2?go-get=1
gopkg.in/fatih/pool.v2 (download)
gopkg.in/fatih/pool.v2 
```

All is OK now!
