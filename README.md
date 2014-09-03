## bqq-go

![Build Status](http://img.shields.io/travis/heroicyang/bqq-go.svg?style=flat-square)

> A Tencent Business QQ API Wrapper In Go.

## Usage

```go
package main

import (
  "fmt"
  bqq "github.com/heroicyang/bqq-go"
)

func main() {
  app := bqq.Init("APP_ID", "APP_SECRET")
  app.BaseEndPoint = "https://openapi.b.qq.com"
  app.RedirectUri = "http://yourdomain.com/oauth/callback"

  // Requset company id and token
  res, _ := app.GetCompanyToken("code", "state")
  companyId := res.Data["company_id"]
  companyToken := res.Data["company_token"]

  // Create a session based on the app
  session := app.CreateSession(companyId, companyToken)

  // Request API with the session
  res, _ := session.GetCompanyInfo()

  fmt.Println("company name: ", res.Data["company_name"])
  fmt.Println("company fullname: ", res.Data["company_fullname"])
}
```
