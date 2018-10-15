package main

import (
	"../bolt"
	"log"
	"fmt"
)

func main()  {
	//1.打开数据库
	db,err := bolt.Open("test.db",0600,nil) //打开数据库test.db 如果不存在就新建一个 0600是文件打开修改删除的权限模式
	if err != nil {
		log.Panic("数据库打开失败")
	}

	//2.操作数据库
	db.Update(func(tx *bolt.Tx) error {
		//找抽屉bucket 如果没有 就创建
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			bucket,err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				panic("bucket b1创建失败")
			}
		}

		//3.写数据
		bucket.Put([]byte("11111"),[]byte("hello"))
		bucket.Put([]byte("22222"),[]byte("world"))

		return nil
	})

	//4.读数据
	db.View(func(tx *bolt.Tx) error {
		//找抽屉bucket 如果没有 就报错退出
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			log.Panic("bucket b1不应该为空 请检查")
		}

		//直接读取数据
		v1 := bucket.Get([]byte("11111"))
		v2 := bucket.Get([]byte("22222"))

		fmt.Printf("v1:%s\n",v1)
		fmt.Printf("v2:%s\n",v2)

		return nil
	})
}
