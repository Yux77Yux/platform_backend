package test

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/Yux77Yux/platform_backend/generated/common"
	api "github.com/Yux77Yux/platform_backend/scripts/api"
)

// 改头像前初始化
func UpdateAvatarInit() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		filename := "E:/xuexi/platform/platform_backend/scripts/result/login_ok_by_db.jsonl"

		f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
		if err != nil {
			log.Fatalf("无法创建临时文件 %s: %v", filename, err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var p Login_OK
			if err := json.Unmarshal(scanner.Bytes(), &p); err != nil {
				log.Fatalf("error: json %s", err.Error())
			}
			LoginOkMap[p.Id] = &p
		}
	}()

	go func() {
		defer wg.Done()
		filename := "E:/xuexi/platform/platform_backend/scripts/result/update_avatar_ok.jsonl"

		f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
		if err != nil {
			log.Fatalf("无法创建临时文件 %s: %v", filename, err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var p Id
			if err := json.Unmarshal(scanner.Bytes(), &p); err != nil {
				log.Fatalf("error: json %s", err.Error())
			}
			UpdateAvatarOkMap[p.Id] = &p
		}
	}()

	wg.Wait()
	log.Printf("init OKOK")
}
func TestUpdateAvatar(t *testing.T) {
	UpdateAvatarInit()
	totalRequests := len(RegisterOkMap)
	errCh := make(chan *User_ER, totalRequests)
	okCh := make(chan *Id, totalRequests)
	concurrencyLimit := int32(3)

	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/update_avatar_err.jsonl"
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("file err %s", err.Error())
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		defer writer.Flush() // 最后统一刷新缓冲区

		for e := range errCh {
			b, err := json.Marshal(e)
			if err != nil {
				log.Fatalf("error Marshal %s", err.Error())
			}

			// 确保每条 JSON 都是独立一行
			if _, err := writer.Write(append(b, '\n')); err != nil {
				log.Fatalf("error writing to file: %s", err.Error())
			}
		}
	}()

	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/update_avatar_ok.jsonl"
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("file err %s", err.Error())
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		defer writer.Flush() // 最后统一刷新缓冲区

		for e := range okCh {
			b, err := json.Marshal(e)
			if err != nil {
				log.Fatalf("error Marshal %s", err.Error())
			}

			// 确保每条 JSON 都是独立一行
			if _, err := writer.Write(append(b, '\n')); err != nil {
				log.Fatalf("error writing to file: %s", err.Error())
			}
		}
	}()

	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrencyLimit) // 信号量控制并发数
	startTime := time.Now()                      // 记录整个测试开始时间
	for _, user := range LoginOkMap {
		wg.Add(1)
		sem <- struct{}{} // 信号量申请，超出则阻塞
		go func(user *Login_OK) {
			User := user.User
			defer func() {
				wg.Done()
				<-sem // 释放信号量
			}()
			// 拿accessToken
			ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
			response, err := api.Refresh(ctx, user.RefreshToken)
			cancel()
			if err != nil {
				errCh <- &User_ER{
					User:  User,
					Error: err.Error(),
				}
				return
			}
			msg := response.GetMsg()
			status := msg.GetStatus()
			if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
				err := fmt.Errorf("code %s error %s", msg.GetCode(), msg.GetDetails())
				errCh <- &User_ER{
					User:  User,
					Error: err.Error(),
				}
				return
			}

			accessToken := response.GetAccessToken()
			// 更新头像
			start := time.Now()
			_ctx, _cancel := context.WithTimeout(context.Background(), 4*time.Second)
			_response, err := api.UpdateUserAvatar(_ctx, User.Avatar, accessToken)
			_cancel()
			end := time.Now()
			if err != nil {
				errCh <- &User_ER{
					User:  User,
					Error: err.Error(),
				}
				return
			}
			msg = _response.GetMsg()
			status = msg.GetStatus()
			if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
				err := fmt.Errorf("code %s error %s", msg.GetCode(), msg.GetDetails())
				errCh <- &User_ER{
					User:  User,
					Error: err.Error(),
				}
				return
			}

			okCh <- &Id{
				Id:       user.Id,
				Duration: math.Round(float64(end.Sub(start).Milliseconds())*100) / 100,
			}
		}(user)
	}
	wg.Wait()
	endTime := time.Now() // 记录整个测试结束时间
	totalDuration := endTime.Sub(startTime).Seconds()
	// 计算吞吐量
	throughput := float64(totalRequests) / totalDuration

	log.Printf("ConcurrencyLimit: %d\n", concurrencyLimit)
	log.Printf("Total Requests: %d\n", totalRequests)
	log.Printf("Total Duration: %.2f seconds\n", totalDuration)
	log.Printf("Throughput: %.2f requests/second\n", throughput)

	close(errCh)
	close(okCh)
}

// 计算平均值
// LoginByDbInit 38.14 ms
// LoginByCacheInit 29.93 ms
func LoginByDbInit() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		filename := "E:/xuexi/platform/platform_backend/scripts/result/login_ok_by_db.jsonl"

		f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
		if err != nil {
			log.Fatalf("无法创建临时文件 %s: %v", filename, err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var p Login_OK
			if err := json.Unmarshal(scanner.Bytes(), &p); err != nil {
				log.Fatalf("error: json %s", err.Error())
			}
			LoginOkMap[p.Id] = &p
		}
	}()

	wg.Wait()
	log.Printf("init OKOK")
}
func LoginByCacheInit() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		filename := "E:/xuexi/platform/platform_backend/scripts/result/login_ok_by_cache.jsonl"

		f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
		if err != nil {
			log.Fatalf("无法创建临时文件 %s: %v", filename, err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var p Login_OK
			if err := json.Unmarshal(scanner.Bytes(), &p); err != nil {
				log.Fatalf("error: json %s", err.Error())
			}
			LoginOkMap[p.Id] = &p
		}
	}()

	wg.Wait()
	log.Printf("init OKOK")
}
func TestLoginPreMs(t *testing.T) {
	LoginByCacheInit()

	var durations []float64
	for _, user := range LoginOkMap {
		durations = append(durations, user.Duration)
	}

	if len(durations) < 3 {
		log.Println("样本数量不足，无法去掉最高最低值")
		return
	}

	sort.Float64s(durations)
	durations = durations[1 : len(durations)-1] // 去掉最高和最低

	sum := float64(0)
	for _, duration := range durations {
		sum += duration
	}

	average := sum / float64(len(durations))
	log.Printf("去掉最高和最低后的平均值为 %.2f ms", average)
}

// 集体登录
func LoginInit() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		filename := "E:/xuexi/platform/platform_backend/scripts/result/register_ok.jsonl"

		f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
		if err != nil {
			log.Fatalf("无法创建临时文件 %s: %v", filename, err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var p User
			if err := json.Unmarshal(scanner.Bytes(), &p); err != nil {
				log.Fatalf("error: json %s", err.Error())
			}
			RegisterOkMap[p.Id] = &p
		}
	}()

	wg.Wait()
	log.Printf("init OKOK")
}
func TestLogin(t *testing.T) {
	LoginInit()
	totalRequests := len(RegisterOkMap)
	errCh := make(chan *Login_ER, totalRequests)
	okCh := make(chan *Login_OK, totalRequests)
	concurrencyLimit := int32(2)

	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/login_err.jsonl"
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("file err %s", err.Error())
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		defer writer.Flush() // 最后统一刷新缓冲区

		for e := range errCh {
			b, err := json.Marshal(e)
			if err != nil {
				log.Fatalf("error Marshal %s", err.Error())
			}

			// 确保每条 JSON 都是独立一行
			if _, err := writer.Write(append(b, '\n')); err != nil {
				log.Fatalf("error writing to file: %s", err.Error())
			}
		}
	}()

	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/login_ok.jsonl"
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("file err %s", err.Error())
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		defer writer.Flush() // 最后统一刷新缓冲区

		for e := range okCh {
			b, err := json.Marshal(e)
			if err != nil {
				log.Fatalf("error Marshal %s", err.Error())
			}

			// 确保每条 JSON 都是独立一行
			if _, err := writer.Write(append(b, '\n')); err != nil {
				log.Fatalf("error writing to file: %s", err.Error())
			}
		}
	}()

	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrencyLimit) // 信号量控制并发数
	startTime := time.Now()                      // 记录整个测试开始时间
	for _, user := range RegisterOkMap {
		wg.Add(1)
		sem <- struct{}{} // 信号量申请，超出则阻塞
		go func(user *User) {
			defer func() {
				wg.Done()
				<-sem // 释放信号量
			}()
			ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
			start := time.Now()
			response, err := api.Login(ctx, user.Id)
			cancel()
			end := time.Now()
			if err != nil {
				errCh <- &Login_ER{
					User:  user,
					Error: err.Error(),
				}
				return
			}

			msg := response.GetMsg()
			status := msg.GetStatus()
			if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
				err := fmt.Errorf("code %s error %s", msg.GetCode(), msg.GetDetails())
				errCh <- &Login_ER{
					User:  user,
					Error: err.Error(),
				}
				return
			}

			lgUser := response.GetUserLogin()
			tokens := response.GetTokens()
			okCh <- &Login_OK{
				User:         user,
				IdInDb:       lgUser.UserDefault.GetUserId(),
				RefreshToken: tokens.GetRefreshToken(),
				Duration:     math.Round(float64(end.Sub(start).Milliseconds())*100) / 100,
			}
		}(user)
	}
	wg.Wait()
	endTime := time.Now() // 记录整个测试结束时间
	totalDuration := endTime.Sub(startTime).Seconds()
	// 计算吞吐量
	throughput := float64(totalRequests) / totalDuration

	log.Printf("ConcurrencyLimit: %d\n", concurrencyLimit)
	log.Printf("Total Requests: %d\n", totalRequests)
	log.Printf("Total Duration: %.2f seconds\n", totalDuration)
	log.Printf("Throughput: %.2f requests/second\n", throughput)

	close(errCh)
	close(okCh)
}

// 集体注册
func RegisterInit() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		filename := "E:/xuexi/platform/platform_backend/scripts/Users.jsonl"
		// filename := "E:/xuexi/platform/platform_backend/scripts/Users.jsonl"

		f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
		if err != nil {
			log.Fatalf("无法读取文件 %s: %v", filename, err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var user User
			if err := json.Unmarshal(scanner.Bytes(), &user); err != nil {
				log.Fatalf("error: json %s", err.Error())
			}
			Users = append(Users, &user)
		}
	}()

	go func() {
		defer wg.Done()
		filename := "E:/xuexi/platform/platform_backend/scripts/Videos.jsonl"

		f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
		if err != nil {
			log.Fatalf("无法创建临时文件 %s: %v", filename, err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var video Creation
			if err := json.Unmarshal(scanner.Bytes(), &video); err != nil {
				log.Fatalf("error: json %s", err.Error())
			}
			Videos = append(Videos, &video)
		}
	}()

	go func() {
		defer wg.Done()
		filename := "E:/xuexi/platform/platform_backend/scripts/result/register_ok.jsonl"

		f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
		if err != nil {
			log.Fatalf("无法创建临时文件 %s: %v", filename, err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var p User
			if err := json.Unmarshal(scanner.Bytes(), &p); err != nil {
				log.Fatalf("error: json %s", err.Error())
			}
			RegisterOkMap[p.Id] = &p
		}
	}()

	wg.Wait()
	log.Printf("init OKOK")
}
func TestRegister(t *testing.T) {
	RegisterInit()
	errCh := make(chan *Register_ER, 5)
	okCh := make(chan *Register_OK, 5)

	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/register_err.jsonl"
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("file err %s", err.Error())
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		defer writer.Flush() // 最后统一刷新缓冲区

		for e := range errCh {
			b, err := json.Marshal(e)
			if err != nil {
				log.Fatalf("error Marshal %s", err.Error())
			}

			// 确保每条 JSON 都是独立一行
			if _, err := writer.Write(append(b, '\n')); err != nil {
				log.Fatalf("error writing to file: %s", err.Error())
			}
		}
	}()

	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/register_ok.jsonl"
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("file err %s", err.Error())
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		defer writer.Flush() // 最后统一刷新缓冲区

		for e := range okCh {
			b, err := json.Marshal(e)
			if err != nil {
				log.Fatalf("error Marshal %s", err.Error())
			}

			// 确保每条 JSON 都是独立一行
			if _, err := writer.Write(append(b, '\n')); err != nil {
				log.Fatalf("error writing to file: %s", err.Error())
			}
		}
	}()

	for _, user := range Users {
		if _, exist := RegisterOkMap[user.Id]; exist {
			continue
		}
		ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
		start := time.Now()
		response, err := api.Register(ctx, user.Id)
		cancel()
		end := time.Now()
		if err != nil {
			errCh <- &Register_ER{
				User:  user,
				Error: err.Error(),
			}
			continue
		}

		msg := response.GetMsg()
		status := msg.GetStatus()
		if status != common.ApiResponse_PENDING && status != common.ApiResponse_SUCCESS {
			err := fmt.Errorf("code %s error %s", msg.GetCode(), msg.GetDetails())
			errCh <- &Register_ER{
				User:  user,
				Error: err.Error(),
			}
			continue
		}

		okCh <- &Register_OK{
			User:     user,
			Duration: math.Round(float64(end.Sub(start).Milliseconds())*100) / 100,
		}
	}
	close(errCh)
	close(okCh)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
