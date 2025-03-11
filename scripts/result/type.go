package result

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ResponseTime struct {
	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
}

type FileName string

const (
	Register_Type      FileName = "register_response.jsonl"
	Login_Type         FileName = "login_response.jsonl"
	Update_Avatar_Type FileName = "update_avatar_response.jsonl"
	Update_Person_Type FileName = "update_person_response.jsonl"

	Upload_Type             FileName = "upload_response.jsonl"
	Get_User_Creations_Type FileName = "get_user_creations_type"
	Publish_Type            FileName = "publish_response.jsonl"
)

func SaveError(err error) error {
	filename := "error.log"
	file, fileErr := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if fileErr != nil {
		return fileErr
	}
	defer file.Close()

	// 写入错误信息并换行
	_, writeErr := fmt.Fprintln(file, err.Error())
	if writeErr != nil {
		return writeErr
	}
	return nil
}

func SaveResponseTime(filename FileName, r *ResponseTime) error {
	file, err := os.OpenFile(string(filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}
	if _, err := writer.Write(data); err != nil {
		return err
	}
	writer.WriteString("\n") // 确保每行一个 JSON 对象
	return writer.Flush()    // 刷新缓冲区
}
