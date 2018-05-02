# Service-now 
A wrapper API for some functions in the Service-now API

ref: https://developer.servicenow.com/app.do#!/rest_api_doc?v=jakarta&id=c_TableAPI

# Install
```Bash
go get github.com/frikky/service-now
```

```Go
import "github.com/frikky/service-now"
```

# Example 
Set logindata, used for any interactive APIcall 
```Go
url := "service-now.com"
username := "foo"
pw := "bar"
verify := false
client := servicenow.CreateLogin(url, username, pw, verify)

ret, err := client.GetTable("incident", 10, "&active=true^ORDERBYDESCsys_created_on")
```

This will return active tickets in the incident table ordered by newest first.
It returns a struct that can be used to get info. The struct is still in the works as can be seen.

# Todo
* Fix the struct as it currently is just a universal one
* Yes, service-now 
