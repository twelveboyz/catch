package utils

import "github.com/bytedance/gopkg/util/gopool"

var GoPool gopool.Pool = gopool.NewPool("gopool", 1000, gopool.NewConfig())
