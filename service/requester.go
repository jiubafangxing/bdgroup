package service

import (
	"fmt"
	"io"
)

type ShareInfoParam struct {
	FromUk int64  `json:"from_uk"` // 将 JSON 字段 from_uk 解析为 FromUk
	ToUk   int64  `json:"to_uk"`   // 将 JSON 字段 to_uk 解析为 ToUk
	MsgId  string `json:"msg_id"`  // 将 JSON 字段 msg_id 解析为 MsgId
	Num    int    `json:"num"`     // 将 JSON 字段 num 解析为 Num
	Page   int    `json:"page"`    // 将 JSON 字段 page 解析为 Page
	FsId   string `json:"fs_id"`   // 将 JSON 字段 fs_id 解析为 FsId
	Gid    string `json:"gid"`     // 将 JSON 字段 gid 解析为 Gid
}

// 获取用户信息
func (s *Service) requestHistorySessions() (io.ReadCloser, error) {
	res, err := s.client.Req("POST", "https://pan.baidu.com/mbox/msg/historysession?clienttype=0&app_id=250528&web=1", nil, map[string]string{})
	return handleHTTPResponse(res, err)
}

// 我的群组
func (s *Service) requestShareGroups() (io.ReadCloser, error) {
	res, err := s.client.Req("POST", "https://pan.baidu.com/mbox/group/list?clienttype=0&app_id=250528&web=1&dp-logid=85740600462994540141&start=0&limit=100", nil, map[string]string{})

	return handleHTTPResponse(res, err)
}

// 获取分享群文件列表
func (s *Service) requestShareGroupFileList(gid string) (io.ReadCloser, error) {
	url := fmt.Sprintf(`https://pan.baidu.com/mbox/group/listshare`)
	queries := fmt.Sprintf(`clienttype=0&app_id=250528&web=1&type=2&gid=%s&limit=50&desc=1&dp-logid=44249000966212050142`, gid)
	url = fmt.Sprintf("%s?%s", url, queries)
	res, err := s.client.Req("POST", url, nil, map[string]string{})
	return handleHTTPResponse(res, err)
}

// 获取分享群文件列表
func (s *Service) requestShareInfo(param ShareInfoParam) (io.ReadCloser, error) {
	url := fmt.Sprintf("https://pan.baidu.com/mbox/msg/shareinfo?type=2&from_uk=%d&msg_id=%s&to_uk=%d&num=%d&page=%d&fs_id=%s&gid=%s&clienttype=0&app_id=250528&web=1",
		param.FromUk, param.MsgId, param.ToUk, param.Num, param.Page, param.FsId, param.Gid)
	res, err := s.client.Req("POST", url, nil, map[string]string{})
	return handleHTTPResponse(res, err)
}

// 获取用户登录信息
func (s *Service) requestLoginStatus() (io.ReadCloser, error) {
	res, err := s.client.Req("GET", "https://pan.baidu.com/api/loginStatus?clienttype=0", nil, map[string]string{})
	return handleHTTPResponse(res, err)
}

func (s *Service) groups() {

}
