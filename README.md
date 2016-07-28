# IPgo
IPgo - class for handling of IPv4.It can create subnets ip.

IPgo是go语言写的生成子网IPv4地址的类。
# Usage


```
package main

package main

import (
	"fmt"
	"github.com/the-404/IPgo"
)

func main() {
	ips := IPgo.Iplist("127.0.0.1/28")
	fmt.Println(ips)
}

```
Result:

	```
	[127.0.0.0 127.0.0.1 127.0.0.2 127.0.0.3 127.0.0.4 127.0.0.5 127.0.0.6 127.0.0.7 127.0.0.8 127.0.0.9 127.0.0.10 127.0.0.11 127.0.0.12 127.0.0.13 127.0.0.14 127.0.0.15]
```


