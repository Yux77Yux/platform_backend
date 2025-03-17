package test

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	api "github.com/Yux77Yux/platform_backend/scripts/api"
)

// 测试用户发布自己的作品
// 初始化
func GetPublishVideoInit() {
	var wg sync.WaitGroup

	// 登录用户，用于获取accessToken
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
			LoginOKMapIdInDb[p.IdInDb] = &p
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		filename := "E:/xuexi/platform/platform_backend/scripts/result/get_video_ok.jsonl"

		f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
		if err != nil {
			log.Fatalf("无法创建临时文件 %s: %v", filename, err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var p CreationInfo_OK
			if err := json.Unmarshal(scanner.Bytes(), &p); err != nil {
				log.Fatalf("error: json %s", err.Error())
			}
			GetVideosOkMapIdInDb[p.CreationId] = &p
		}
	}()

	wg.Wait()
}
func TestPublishVideo(t *testing.T) {
	GetPublishVideoInit()
	totalRequests := int32(0)
	errCh := make(chan *CreationInfo_ER, totalRequests)
	okCh := make(chan *CreationInfo_OK, totalRequests)
	concurrencyLimit := int32(19)
	var okWg, errWg sync.WaitGroup

	// 初始化错误通道
	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/pending_video_err.jsonl"
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

	// 初始化成功通道
	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/pending_video_ok.jsonl"
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
	for _, video := range GetVideosOkMapIdInDb {
		creationId := video.CreationId

		wg.Add(1)
		sem <- struct{}{} // 信号量申请，超出则阻塞
		go func(video *CreationInfo_OK) {
			defer func() {
				wg.Done()
				<-sem // 释放信号量
			}()
			// 拿accessToken
			ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
			authorId := video.AuthorId
			atomic.AddInt32(&totalRequests, 1)
			response, err := api.Refresh(ctx, LoginOKMapIdInDb[authorId].RefreshToken)
			cancel()
			if err != nil {
				errWg.Add(1)
				errCh <- &CreationInfo_ER{
					CreationId: creationId,
					AuthorId:   authorId,
					Error:      err.Error(),
				}
				return
			}
			msg := response.GetMsg()
			status := msg.GetStatus()
			if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
				err := fmt.Errorf("code %s error %s", msg.GetCode(), msg.GetDetails())
				errWg.Add(1)
				errCh <- &CreationInfo_ER{
					CreationId: creationId,
					AuthorId:   authorId,
					Error:      err.Error(),
				}
				return
			}

			accessToken := response.GetAccessToken()
			start := time.Now()
			_ctx, _cancel := context.WithTimeout(context.Background(), 4*time.Second)

			atomic.AddInt32(&totalRequests, 1)
			_response, err := api.PublishDraftCreation(_ctx, accessToken, creationId)
			_cancel()
			end := time.Now()
			if err != nil {
				errWg.Add(1)
				errCh <- &CreationInfo_ER{
					CreationId: creationId,
					AuthorId:   authorId,
					Error:      err.Error(),
				}
				return
			}
			msg = _response.GetMsg()
			status = msg.GetStatus()
			if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
				err := fmt.Errorf("code %s error %s", msg.GetCode(), msg.GetDetails())
				errWg.Add(1)
				errCh <- &CreationInfo_ER{
					CreationId: creationId,
					AuthorId:   authorId,
					Error:      err.Error(),
				}
				return
			}

			okWg.Add(1)
			okCh <- &CreationInfo_OK{
				CreationId: creationId,
				AuthorId:   authorId,
				Title:      video.Title,
				Duration:   math.Round(float64(end.Sub(start).Milliseconds())*100) / 100,
			}
		}(video)
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

/* TestPublishVideo
2025/03/17 22:37:10 ConcurrencyLimit: 19
2025/03/17 22:37:10 Total Requests: 916
2025/03/17 22:37:10 Total Duration: 0.88 seconds
2025/03/17 22:37:10 Throughput: 1044.43 requests/second

2025/03/17 19:08:10 ConcurrencyLimit: 19
2025/03/17 19:08:10 Total Requests: 916
2025/03/17 19:08:10 Total Duration: 1.47 seconds
2025/03/17 19:08:10 Throughput: 622.75 requests/second

2025/03/17 19:11:42 ConcurrencyLimit: 19
2025/03/17 19:11:42 Total Requests: 916
2025/03/17 19:11:42 Total Duration: 1.40 seconds
2025/03/17 19:11:42 Throughput: 652.25 requests/second

2025/03/17 20:05:19 ConcurrencyLimit: 19
2025/03/17 20:05:19 Total Requests: 916
2025/03/17 20:05:19 Total Duration: 1.40 seconds
2025/03/17 20:05:19 Throughput: 654.75 requests/second

2025/03/17 19:09:07 ConcurrencyLimit: 19
2025/03/17 19:09:07 Total Requests: 916
2025/03/17 19:09:07 Total Duration: 1.43 seconds
2025/03/17 19:09:07 Throughput: 640.65 requests/second

2025/03/17 23:16:24 ConcurrencyLimit: 19
2025/03/17 23:16:24 Total Requests: 916
2025/03/17 23:16:24 Total Duration: 0.99 seconds
2025/03/17 23:16:24 Throughput: 921.80 requests/second
*/

// 测试用户发布自己的作品
// 初始化
func GetDbVideoInit() {
	var wg sync.WaitGroup

	// 登录用户，用于获取accessToken
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
}
func TestGetDbVideo(t *testing.T) {
	GetDbVideoInit()
	totalRequests := int32(0)
	errCh := make(chan *CreationInfo_ER, totalRequests)
	okCh := make(chan *CreationInfo_OK, totalRequests)
	concurrencyLimit := int32(4)
	var okWg, errWg sync.WaitGroup

	// 初始化错误通道
	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/get_video_err.jsonl"
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

	// 初始化成功通道
	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/get_video_ok.jsonl"
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
			defer func() {
				wg.Done()
				<-sem // 释放信号量
			}()
			// 拿accessToken
			ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
			atomic.AddInt32(&totalRequests, 1)
			response, err := api.Refresh(ctx, user.RefreshToken)
			cancel()
			if err != nil {
				errWg.Add(1)
				errCh <- &CreationInfo_ER{
					AuthorId: user.IdInDb,
					Error:    err.Error(),
				}
				return
			}
			msg := response.GetMsg()
			status := msg.GetStatus()
			if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
				err := fmt.Errorf("code %s error %s", msg.GetCode(), msg.GetDetails())
				errWg.Add(1)
				errCh <- &CreationInfo_ER{
					AuthorId: user.IdInDb,
					Error:    err.Error(),
				}
				return
			}

			accessToken := response.GetAccessToken()
			start := time.Now()

			_ctx, _cancel := context.WithTimeout(context.Background(), 4*time.Second)
			atomic.AddInt32(&totalRequests, 1)
			_response, err := api.GetUserCreations(_ctx, accessToken, generated.CreationStatus_DRAFT, 1)
			_cancel()
			end := time.Now()
			if err != nil {
				errWg.Add(1)
				errCh <- &CreationInfo_ER{
					AuthorId: user.IdInDb,
					Error:    err.Error(),
				}
				return
			}
			msg = _response.GetMsg()
			status = msg.GetStatus()
			if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
				err := fmt.Errorf("code %s error %s", msg.GetCode(), msg.GetDetails())
				errWg.Add(1)
				errCh <- &CreationInfo_ER{
					AuthorId: user.IdInDb,
					Error:    err.Error(),
				}
				return
			}

			CreationInfos := _response.GetCreationInfoGroup()
			for _, info := range CreationInfos {
				okWg.Add(1)
				okCh <- &CreationInfo_OK{
					CreationId: info.Creation.CreationId,
					AuthorId:   user.IdInDb,
					Title:      info.Creation.BaseInfo.Title,
					Duration:   math.Round(float64(end.Sub(start).Milliseconds())*100) / 100,
				}
			}

			count := _response.GetCount()

			for i := int32(2); i <= count; i++ {
				_ctx, _cancel := context.WithTimeout(context.Background(), 4*time.Second)
				start := time.Now()
				atomic.AddInt32(&totalRequests, 1)
				i_response, err := api.GetUserCreations(_ctx, accessToken, generated.CreationStatus_DRAFT, i)
				_cancel()
				end := time.Now()
				if err != nil {
					errWg.Add(1)
					errCh <- &CreationInfo_ER{
						AuthorId: user.IdInDb,
						Error:    err.Error(),
					}
					return
				}
				msg = i_response.GetMsg()
				status = msg.GetStatus()
				if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
					err := fmt.Errorf("code %s error %s", msg.GetCode(), msg.GetDetails())
					errWg.Add(1)
					errCh <- &CreationInfo_ER{
						AuthorId: user.IdInDb,
						Error:    err.Error(),
					}
					continue
				}

				CreationInfos := i_response.GetCreationInfoGroup()
				for _, info := range CreationInfos {
					okWg.Add(1)
					okCh <- &CreationInfo_OK{
						CreationId: info.Creation.CreationId,
						AuthorId:   user.IdInDb,
						Title:      info.Creation.BaseInfo.Title,
						Duration:   math.Round(float64(end.Sub(start).Milliseconds())*100) / 100,
					}
				}
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
	time.Sleep(2 * time.Second)
	close(errCh)
	close(okCh)
}

/* TestGetDbVideo
2025/03/17 18:35:15 ConcurrencyLimit: 4
2025/03/17 18:35:15 Total Requests: 717
2025/03/17 18:35:15 Total Duration: 3.67 seconds
2025/03/17 18:35:15 Throughput: 195.58 requests/second
*/

// 上传视频前初始化
func UplodaVideoInit() {
	var wg sync.WaitGroup
	wg.Add(3)

	// 登录用户，用于获取accessToken
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

	// 上传成功部分，不再重新上传
	go func() {
		defer wg.Done()
		filename := "E:/xuexi/platform/platform_backend/scripts/result/upload_ok.jsonl"

		f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
		if err != nil {
			log.Fatalf("无法创建临时文件 %s: %v", filename, err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var p Creation_OK
			if err := json.Unmarshal(scanner.Bytes(), &p); err != nil {
				log.Fatalf("error: json %s", err.Error())
			}
			UploadOkMap[p.Id] = &p
		}
	}()

	// 获取要上传的视频集
	go func() {
		defer wg.Done()
		filename := "E:/xuexi/platform/platform_backend/scripts/videos.jsonl"

		f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
		if err != nil {
			log.Fatalf("无法创建临时文件 %s: %v", filename, err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var p Creation
			if err := json.Unmarshal(scanner.Bytes(), &p); err != nil {
				log.Fatalf("error: json %s", err.Error())
			}
			Videos = append(Videos, &p)
		}
	}()

	wg.Wait()
}
func TestUplodaVideo(t *testing.T) {
	UplodaVideoInit()
	totalRequests := len(Videos)
	errCh := make(chan *Creation_ER, totalRequests)
	okCh := make(chan *Creation_OK, totalRequests)
	concurrencyLimit := int32(3)
	var okWg, errWg sync.WaitGroup

	// 初始化错误通道
	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/upload_err.jsonl"
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

	// 初始化成功通道
	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/upload_ok.jsonl"
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
	for _, video := range Videos {
		if _, exist := UploadOkMap[video.Id]; exist {
			continue
		}

		wg.Add(1)
		sem <- struct{}{} // 信号量申请，超出则阻塞
		go func(video *Creation) {
			defer func() {
				wg.Done()
				<-sem // 释放信号量
			}()
			// 拿accessToken
			ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
			response, err := api.Refresh(ctx, LoginOkMap[video.Uid].RefreshToken)
			cancel()
			if err != nil {
				errWg.Add(1)
				errCh <- &Creation_ER{
					Creation: video,
					Error:    err.Error(),
				}
				return
			}
			msg := response.GetMsg()
			status := msg.GetStatus()
			if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
				err := fmt.Errorf("code %s error %s", msg.GetCode(), msg.GetDetails())
				errWg.Add(1)
				errCh <- &Creation_ER{
					Creation: video,
					Error:    err.Error(),
				}
				return
			}

			accessToken := response.GetAccessToken()
			start := time.Now()
			_ctx, _cancel := context.WithTimeout(context.Background(), 4*time.Second)
			_response, err := api.UploadCreation(_ctx, accessToken, video)
			_cancel()
			end := time.Now()
			if err != nil {
				errWg.Add(1)
				errCh <- &Creation_ER{
					Creation: video,
					Error:    err.Error(),
				}
				return
			}
			msg = _response.GetMsg()
			status = msg.GetStatus()
			if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
				err := fmt.Errorf("code %s error %s", msg.GetCode(), msg.GetDetails())
				errWg.Add(1)
				errCh <- &Creation_ER{
					Creation: video,
					Error:    err.Error(),
				}
				return
			}

			okWg.Add(1)
			okCh <- &Creation_OK{
				Creation:       video,
				UploadDuration: math.Round(float64(end.Sub(start).Milliseconds())*100) / 100,
			}
		}(video)
	}
	endTime := time.Now() // 记录整个测试结束时间
	wg.Wait()
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

/* TestUplodaVideo
2025/03/17 17:40:23 ConcurrencyLimit: 3
2025/03/17 17:40:23 Total Requests: 357
2025/03/17 17:40:23 Total Duration: 4.06 seconds
2025/03/17 17:40:23 Throughput: 87.96 requests/second
*/
