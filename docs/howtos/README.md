# My Golang

Just have started playing with https://pages.github.com/
https://mirkar.github.io/mygolang/

```
/home/dm0434/go
go install github.com/mirkar/mygolang/hello
```

#### Set proxy when executing “go get” command

In Windows, if you have configured proxy for connecting Internet, then executing “go get” command may get error, like this:
changed
```text
C:\Users\xiaona>go get -v gopkg.in/fatih/pool.v2
Fetching https://gopkg.in/fatih/pool.v2?go-get=1
https fetch failed.
Fetching http://gopkg.in/fatih/pool.v2?go-get=1
import "gopkg.in/fatih/pool.v2": http/https fetch: Get http://gopkg.in/fatih/poo
l.v2?go-get=1: dial tcp 107.178.216.236:80: ConnectEx tcp: A connection attempt
failed because the connected party did not properly respond after a period of ti
me, or established connection failed because connected host has failed to respon
d.
package gopkg.in/fatih/pool.v2: unrecognized import path "gopkg.in/fatih/pool.v2
"
```

The solution is set “http_proxy” and/or “https_proxy” environment variables according to your reality need:

```
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
