package test

import (
	"bufio"
	"context"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Yux77Yux/platform_backend/generated/auth"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	"github.com/Yux77Yux/platform_backend/generated/review"
	api "github.com/Yux77Yux/platform_backend/scripts/api"
	"google.golang.org/protobuf/encoding/protojson"
)

const Token = "olHmHcQLxQxlZqAcxII1ygKP17SDXl9K3guOlhRKpkxL48p7FbyEGOG5HadnW9sJPBIydnBRb1xONMUsj+l7+GewJ6nNKJxT3pjXawUYjP5J+OiJ6aWbamwM/Knfc9lf7cztNCguTs2Mm2tCMTZzU1ruXVW/TWYyxvihkPcuni+qniyiu8pumqB4X1+cV9dAbx6OvlKIniB9RHnum8S6TFDdFoCVdc5tuU2vVWk9ddvWPjhPyFq7FblwWMVp5rVYKjPixdWozkTPruDcOClXIw=="

func TestGetAndUpdateNewReviews(t *testing.T) {
	reviewerId := int64(1901557963315744768)
	refreshToken := &auth.RefreshToken{
		Value: Token,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	response, err := api.Refresh(ctx, refreshToken)
	cancel()
	if err != nil {
		log.Fatalf("token fetch error %s", err.Error())
	}
	msg := response.GetMsg()
	status := msg.GetStatus()
	if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
		log.Fatalf("token fetch error %s", msg.GetDetails())
	}
	accessToken := response.GetAccessToken()

	totalRequests := int32(1)
	okCh := make(chan *review.Review, totalRequests)
	concurrencyLimit := int32(1)
	done := make(chan any, concurrencyLimit)

	var (
		okWg sync.WaitGroup
		wg   sync.WaitGroup
	)

	go func() {
		path := "E:/xuexi/platform/platform_backend/scripts/result/review/review_ok.jsonl"
		file, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("file err %s", err.Error())
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		defer writer.Flush() // 最后统一刷新缓冲区

		for e := range okCh {
			b, err := protojson.Marshal(e)
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

	startTime := time.Now() // 记录整个测试开始时间
	for {
		_ctx, _cancel := context.WithTimeout(context.Background(), 4*time.Second)
		atomic.AddInt32(&totalRequests, 1)
		r_response, err := api.GetNewReviews(_ctx, reviewerId, review.TargetType_CREATION)
		_cancel()
		if err != nil {
			log.Printf("GetReviews %s", err.Error())
			return
		}

		newReviews := r_response.GetReviews()
		msg := r_response.GetMsg()
		status := msg.GetStatus()
		if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
			return
		}

		if len(newReviews) <= 0 {
			break
		}

		for _, _review := range newReviews {
			_review.Status = review.ReviewStatus_APPROVED
			wg.Add(1)
			done <- struct{}{}

			go func(_review *review.Review) {
				defer func() {
					wg.Done()
					<-done
				}()
				ctx1, cancel1 := context.WithTimeout(context.Background(), 5*time.Second)
				atomic.AddInt32(&totalRequests, 1)
				response, err := api.UpdateReview(ctx1, accessToken, _review)
				cancel1()
				if err != nil {
					log.Printf("UpdateReview %s", err.Error())
					return
				}
				msg := response.GetMsg()
				status := msg.GetStatus()
				if status != common.ApiResponse_SUCCESS && status != common.ApiResponse_PENDING {
					log.Printf("UpdateReview %s", msg.GetDetails())
					return
				}

				okWg.Add(1)
				okCh <- _review
			}(_review)
		}
	}
	endTime := time.Now() // 记录整个测试结束时间
	totalDuration := endTime.Sub(startTime).Seconds()
	// 计算吞吐量
	throughput := float64(totalRequests) / totalDuration

	log.Printf("ConcurrencyLimit: %d\n", concurrencyLimit)
	log.Printf("Total Requests: %d\n", totalRequests)
	log.Printf("Total Duration: %.2f seconds\n", totalDuration)
	log.Printf("Throughput: %.2f requests/second\n", throughput)
	wg.Wait()
	okWg.Wait()

	close(okCh)
}

/* TestGetAndUpdateNewReviews
2025/03/19 00:53:48 ConcurrencyLimit: 1
2025/03/19 00:53:48 Total Requests: 518
2025/03/19 00:53:48 Total Duration: 4.75 seconds
2025/03/19 00:53:48 Throughput: 188.73 requests/second
*/
