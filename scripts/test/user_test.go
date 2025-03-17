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
	"sync/atomic"
	"testing"
	"time"

	"github.com/Yux77Yux/platform_backend/generated/common"
	api "github.com/Yux77Yux/platform_backend/scripts/api"
)

// 改个人资料前初始化
func UpdateSpaceInit() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		filename := "E:/xuexi/platform/platform_backend/scripts/result/login_ok.jsonl"

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
func TestUpdateSpace(t *testing.T) {
	UpdateSpaceInit()
	totalRequests := len(LoginOkMap)
	errCh := make(chan *User_ER, totalRequests)
	okCh := make(chan *Id, totalRequests)
	concurrencyLimit := int32(3)
	var okWg, errWg sync.WaitGroup

	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/update_space_err.jsonl"
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
			errWg.Done()
		}
	}()

	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/update_space_ok.jsonl"
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
			okWg.Done()
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
				errWg.Add(1)
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
				errWg.Add(1)
				errCh <- &User_ER{
					User:  User,
					Error: err.Error(),
				}
				return
			}

			accessToken := response.GetAccessToken()
			// 更新头像
			if User.Name == "" {
				err := fmt.Errorf("id %s name %s is null", User.Id, User.Name)
				errWg.Add(1)
				errCh <- &User_ER{
					User:  User,
					Error: err.Error(),
				}
				return
			}
			start := time.Now()
			_ctx, _cancel := context.WithTimeout(context.Background(), 4*time.Second)
			_response, err := api.UpdateUserSpace(_ctx, User.Name, user.Bio, accessToken)
			_cancel()
			end := time.Now()
			if err != nil {
				errWg.Add(1)
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
				errWg.Add(1)
				errCh <- &User_ER{
					User:  User,
					Error: err.Error(),
				}
				return
			}

			okWg.Add(1)
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

	okWg.Wait()
	errWg.Wait()
	close(errCh)
	close(okCh)
}

/* TestUpdateSpace
2025/03/17 17:29:16 ConcurrencyLimit: 3
2025/03/17 17:29:16 Total Requests: 357
2025/03/17 17:29:16 Total Duration: 1.64 seconds
2025/03/17 17:29:16 Throughput: 217.68 requests/second
*/

// 改头像前初始化
func UpdateAvatarInit() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		filename := "E:/xuexi/platform/platform_backend/scripts/result/login_ok.jsonl"

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
func TestUpdateAvatar(t *testing.T) {
	UpdateAvatarInit()
	totalRequests := len(LoginOkMap)
	errCh := make(chan *User_ER, totalRequests)
	okCh := make(chan *Id, totalRequests)
	var okWg, errWg sync.WaitGroup
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
			errWg.Done()
		}
	}()

	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/update_avatar_ok.jsonl"
		file, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
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
			okWg.Done()
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
				errWg.Add(1)
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
				errWg.Add(1)
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
				errWg.Add(1)
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
				errWg.Add(1)
				errCh <- &User_ER{
					User:  User,
					Error: err.Error(),
				}
				return
			}

			okWg.Add(1)
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

	okWg.Wait()
	errWg.Wait()

	close(errCh)
	close(okCh)
}

/* TestUpdateAvatar
2025/03/17 17:00:43 ConcurrencyLimit: 3
2025/03/17 17:00:43 Total Requests: 357
2025/03/17 17:00:43 Total Duration: 1.28 seconds
2025/03/17 17:00:43 Throughput: 279.73 requests/second
*/

// 计算平均值
// LoginByDbInit 38.14 ms
// LoginByCacheInit 29.93 ms
func LoginByInit() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		filename := "E:/xuexi/platform/platform_backend/scripts/result/login_ok.jsonl"

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
	LoginByInit()

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
			var p Register_OK
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
	concurrencyLimit := int32(3)
	var okWg, errWg sync.WaitGroup

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
			errWg.Done()
		}
	}()

	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/login_ok.jsonl"
		file, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
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
			okWg.Done()
		}
	}()

	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrencyLimit) // 信号量控制并发数
	startTime := time.Now()                      // 记录整个测试开始时间
	for _, user := range RegisterOkMap {
		wg.Add(1)
		sem <- struct{}{} // 信号量申请，超出则阻塞
		go func(user *Register_OK) {
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
				errWg.Add(1)
				errCh <- &Login_ER{
					User:  user.User,
					Error: err.Error(),
				}
				return
			}

			msg := response.GetMsg()
			status := msg.GetStatus()
			if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
				err := fmt.Errorf("code %s error %s", msg.GetCode(), msg.GetDetails())
				errWg.Add(1)
				errCh <- &Login_ER{
					User:  user.User,
					Error: err.Error(),
				}
				return
			}

			lgUser := response.GetUserLogin()
			tokens := response.GetTokens()
			okWg.Add(1)
			okCh <- &Login_OK{
				User:         user.User,
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

	okWg.Wait()
	errWg.Wait()
	close(errCh)
	close(okCh)
}

/* TestLogin
2025/03/17 20:21:52 ConcurrencyLimit: 3
2025/03/17 20:21:52 Total Requests: 357
2025/03/17 20:21:52 Total Duration: 11.33 seconds
2025/03/17 20:21:52 Throughput: 31.51 requests/second

2025/03/17 16:36:04 ConcurrencyLimit: 3
2025/03/17 16:36:04 Total Requests: 354
2025/03/17 16:36:04 Total Duration: 10.11 seconds
2025/03/17 16:36:04 Throughput: 35.01 requests/second
*/

// 集体注册
func RegisterInit() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		filename := "E:/xuexi/platform/platform_backend/scripts/users.jsonl"

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
		filename := "E:/xuexi/platform/platform_backend/scripts/result/register_ok.jsonl"

		f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
		if err != nil {
			log.Fatalf("无法创建临时文件 %s: %v", filename, err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var p Register_OK
			if err := json.Unmarshal(scanner.Bytes(), &p); err != nil {
				log.Fatalf("error: json %s", err.Error())
			}
			RegisterOkMap[p.Id] = &p
		}
	}()

	wg.Wait()
	log.Printf("init OKOK")
}
func TestRegisterDuration(t *testing.T) {
	RegisterInit()
	var durations []float64
	for _, user := range RegisterOkMap {
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
func TestRegister(t *testing.T) {
	RegisterInit()
	totalRequests := int32(0)
	concurrencyLimit := 3
	errCh := make(chan *Register_ER, 5)
	okCh := make(chan *Register_OK, 5)
	var okWg, errWg sync.WaitGroup
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
			errWg.Done()
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
			okWg.Done()
		}
	}()
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrencyLimit) // 信号量控制并发数
	startTime := time.Now()                      // 记录整个测试开始时间

	for _, user := range Users {
		if _, exist := RegisterOkMap[user.Id]; exist {
			continue
		}
		atomic.AddInt32(&totalRequests, 1)
		wg.Add(1)
		sem <- struct{}{} // 信号量申请，超出则阻塞

		go func(user *User) {
			defer func() {
				wg.Done()
				<-sem // 释放信号量
			}()

			ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
			start := time.Now()
			response, err := api.Register(ctx, user.Id)
			cancel()
			end := time.Now()
			if err != nil {
				errWg.Add(1)
				errCh <- &Register_ER{
					User:  user,
					Error: err.Error(),
				}
				return
			}

			msg := response.GetMsg()
			status := msg.GetStatus()
			if status != common.ApiResponse_PENDING && status != common.ApiResponse_SUCCESS {
				err := fmt.Errorf("code %s error %s", msg.GetCode(), msg.GetDetails())
				errWg.Add(1)
				errCh <- &Register_ER{
					User:  user,
					Error: err.Error(),
				}
				return
			}

			okWg.Add(1)
			okCh <- &Register_OK{
				User:     user,
				Duration: math.Round(float64(end.Sub(start).Milliseconds())*100) / 100,
			}
		}(user)
	}
	endTime := time.Now() // 记录整个测试结束时间
	okWg.Wait()
	errWg.Wait()
	close(errCh)
	close(okCh)

	totalDuration := endTime.Sub(startTime).Seconds()
	// 计算吞吐量
	throughput := float64(totalRequests) / totalDuration

	log.Printf("ConcurrencyLimit: %d\n", concurrencyLimit)
	log.Printf("Total Requests: %d\n", totalRequests)
	log.Printf("Total Duration: %.2f seconds\n", totalDuration)
	log.Printf("Throughput: %.2f requests/second\n", throughput)
}

/* TestRegister
2025/03/17 16:46:14 ConcurrencyLimit: 3
2025/03/17 16:46:14 Total Requests: 357
2025/03/17 16:46:14 Total Duration: 3.06 seconds
2025/03/17 16:46:14 Throughput: 116.65 requests/second
*/
