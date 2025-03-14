package logger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"

	"github.com/gofrs/flock"
)

var (
	lm   *LoggerManager
	once sync.Once
)

func GetLoggerManager() *LoggerManager {
	if lm != nil {
		return lm
	}
	once.Do(func() {
		lm = &LoggerManager{
			files:      make(map[string]*bufio.Writer),
			logChan:    make(chan *LogFile, 600),
			sharedChan: make(chan *LogMessage, 1000),
			flushFreq:  time.Second * 4,
		}

		go lm.listen()
	})
	return lm
}

func init() {
	lm = GetLoggerManager()
}

type Level string

const (
	INFO    Level = "[INFO]"
	DEBUG   Level = "[DEBUG]"
	WARNING Level = "[WARNING]"
	ERROR   Level = "[ERROR]"
	SUPER   Level = "[SUPER]"
)

// 日志消息结构
type LogMessage struct {
	Level     Level                  `json:"level"`
	TraceId   string                 `json:"traceId"`
	Timestamp time.Time              `json:"timestamp"`
	Position  string                 `json:"position"`
	Message   string                 `json:"message"`
	Extra     map[string]interface{} `json:"extra,omitempty"`
}

type LogFile struct {
	Path string
	*LogMessage
}

type LoggerManager struct {
	files map[string]*bufio.Writer
	mu    sync.Mutex

	sharedChan chan *LogMessage
	logChan    chan *LogFile
	flushFreq  time.Duration
}

func (lm *LoggerManager) listen() {
	flushTicker := time.NewTicker(lm.flushFreq)
	defer flushTicker.Stop()

	go lm.flushShared()

	for {
		select {
		case logMsg := <-lm.logChan:
			lm.writeLog(logMsg)
		case <-flushTicker.C:
			lm.flushAll()
		}
	}
}

func (lm *LoggerManager) GetMatches(str, pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindStringSubmatch(str)
}

func (lm *LoggerManager) SplitFullName(fullName string) (string, string) {
	matches := lm.GetMatches(fullName, `^/([^/.]+)\.([^/]+)Service/([^/]+)$`)
	if len(matches) != 4 {
		return "server", fullName
	}
	return matches[1], matches[3]
}

func (lm *LoggerManager) writeLog(msg *LogFile) {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	key := msg.Path
	writer, exists := lm.files[key]
	if !exists {
		// 获取文件所在目录
		dir := filepath.Dir(key)
		// 确认目录是否存在，如果不存在就创建
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Println("创建目录失败:", err)
				return
			}
		}

		file, err := os.OpenFile(key, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Failed to open log file: %s\n", err)
			return
		}
		writer = bufio.NewWriter(file)
		lm.files[key] = writer
	}

	body, err := json.Marshal(msg.LogMessage)
	if err != nil {
		log.Printf("error: LogMessage Marshal Failed %s", err.Error())
		return
	}
	writer.Write(append(body, '\n'))
}

func (lm *LoggerManager) flushAll() {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	for _, writer := range lm.files {
		_ = writer.Flush()
	}
}

func (lm *LoggerManager) flushShared() {
	filePath := "../../log/services.log"
	// 获取文件所在目录
	dir := filepath.Dir(filePath)
	// 确认目录是否存在，如果不存在就创建
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("创建目录失败:", err)
			return
		}
	}

	for {
		lock := flock.New(filePath + ".lock")
		if err := lock.Lock(); err != nil {
			err = fmt.Errorf("failed to acquire lock: %v", err)
			log.Printf("%s", err.Error())
			time.Sleep(2 * time.Second) // 没拿到锁就等一会儿再试
			continue
		}

		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			lock.Unlock()
			log.Printf("%s", err.Error())
			time.Sleep(time.Second) // 确保释放锁后再等一会儿
			continue
		}

		writer := bufio.NewWriter(file)
		ticker := time.NewTicker(4 * time.Second)
		defer ticker.Stop()

		var exitSelect bool
		for {
			select {
			case <-ticker.C:
				if err := writer.Flush(); err != nil {
					log.Printf("failed to flush log writer: %v", err)
				}
				file.Close()
				lock.Unlock()
				time.Sleep(time.Second * 2) // 给其他服务释放锁的机会
				exitSelect = true
			case msg, ok := <-lm.sharedChan:
				if !ok {
					writer.Flush()
					file.Close()
					lock.Unlock()
					return
				}
				body, err := json.Marshal(msg)
				if err != nil {
					log.Printf("failed to marshal log message: %v", err)
					return
				}
				if _, err := writer.Write(append(body, '\n')); err != nil {
					log.Printf("failed to write log message: %v", err)
				}
			}
			if exitSelect {
				break
			}
		}
	}
}

func (lm *LoggerManager) Log(msg *LogFile) {
	lm.logChan <- msg
}

func (lm *LoggerManager) SharedLog(msg *LogMessage) {
	lm.sharedChan <- msg
}

func (lm *LoggerManager) Close() {
	lm.flushAll()
	close(lm.logChan)
}
