package test

import (
  lru "github.com/hashicorp/golang-lru"
  ctx "github.com/gorilla/context"
  "golang.org/x/text/unicode/norm"
)

func main() { 
  lru.New(1)
  ctx.Purge(0)
  _ = norm.Version
}
