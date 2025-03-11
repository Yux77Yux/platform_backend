package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type UserProfile struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Bio    string `json:"bio"`
	Avatar string `json:"avatar"`

	Creations []*Creation `json:"creations"`
}
type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Bio    string `json:"bio"`
	Avatar string `json:"avatar"`
}
type Creation struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Bio        string `json:"bio"`
	Uid        string `json:"uid"`
	Src        string `json:"src"`
	Thumbnail  string `json:"thumbnail"`
	Duration   int32  `json:"duration"`
	CategoryId int32  `json:"categoryId"`
}

var (
	videos = make([]*Creation, 0, 500)
	users  = make([]*User, 0, 500)
)

func main() {
	filePath := "E:/xuexi/platform/platform_backend/scripts/data/ok_profiles.jsonl"
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("error: file can't open %s", err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) > 0 { // 确保最后一行没有丢
					fmt.Println(line)
				}
				break // 正常结束读取
			}
			log.Fatalf("%s", err.Error())
		}
		var profile UserProfile                      // 不需要指针，直接初始化
		err = json.Unmarshal([]byte(line), &profile) // 传入指针
		if err != nil {
			log.Fatalf("%s", err.Error())
		}
		users = append(users, &User{
			Id:     profile.Id,
			Avatar: profile.Avatar,
			Name:   profile.Name,
			Bio:    profile.Bio,
		})
		creations := profile.Creations
		videos = append(videos, creations...)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		filename := "users.jsonl"
		tempFilename := filename + ".tmp"

		f, err := os.OpenFile(tempFilename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatalf("无法创建临时文件 %s: %v", tempFilename, err)
		}

		enc := json.NewEncoder(f)

		for _, user := range users {
			if err := enc.Encode(user); err != nil {
				f.Close()
				os.Remove(tempFilename)
				log.Fatalf("写入临时文件失败: %v", err)
			}
		}

		// 先关闭文件，确保完全释放文件句柄
		if err := f.Close(); err != nil {
			os.Remove(tempFilename)
			log.Fatalf("关闭临时文件失败: %v", err)
		}

		// 确保写入完成后替换原文件
		if err := os.Rename(tempFilename, filename); err != nil {
			os.Remove(tempFilename)
			log.Fatalf("替换原文件失败: %v", err)
		}

		log.Printf("✅ 用户资料已保存至: %s\n", filename)
	}()

	go func() {
		defer wg.Done()
		filename := "videos.jsonl"
		tempFilename := filename + ".tmp"

		f, err := os.OpenFile(tempFilename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatalf("无法创建临时文件 %s: %v", tempFilename, err)
		}

		enc := json.NewEncoder(f)

		for _, video := range videos {
			if err := enc.Encode(video); err != nil {
				f.Close()
				os.Remove(tempFilename)
				log.Fatalf("写入临时文件失败: %v", err)
			}
		}

		// 先关闭文件，确保完全释放文件句柄
		if err := f.Close(); err != nil {
			os.Remove(tempFilename)
			log.Fatalf("关闭临时文件失败: %v", err)
		}

		// 确保写入完成后替换原文件
		if err := os.Rename(tempFilename, filename); err != nil {
			os.Remove(tempFilename)
			log.Fatalf("替换原文件失败: %v", err)
		}

		log.Printf("✅ 用户资料已保存至: %s\n", filename)
	}()

	wg.Wait()
	log.Printf("OKOK")

}
