package main

import (
	"errors"
	"github.com/afex/hystrix-go/hystrix"
	"log"
	"net/http"
	"time"
)

func main()  {

	// 限流接口
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(":8074",hystrixStreamHandler)

	hystrix.ConfigureCommand("aaa",hystrix.CommandConfig{
		Timeout :1000, // 单次请求 超时时间
		MaxConcurrentRequests :1, // 最大并发量
		RequestVolumeThreshold :5000, // 熔断后多久去尝试服务是否可用
		SleepWindow :1,// 验证熔断的请求数量，10秒内采样
		ErrorPercentThreshold:1, // 验证熔断的错误百分比
	})

	for i:= 0; i<1000; i++ {
		// 异步调用使用hystrix.go
		err := hystrix.Do("aaa", func() error {
			if i == 0 {
				return errors.New("service error")
			}
			log.Println("do services")
			return nil
		},nil)

		if err != nil {
			log.Println("hystrix err:" +err.Error())
			time.Sleep(1*time.Second)
			log.Println("sleep 1 second")
		}

	}





}
