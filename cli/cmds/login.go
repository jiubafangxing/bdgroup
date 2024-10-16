package cmds

import (
	"errors"
	"fmt"
	"github.com/bdgroup/service"
	"os"
	"strconv"
	"time"

	"github.com/bdgroup/config"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

// Login login data
type Login struct {
	bduss   string
	stoken  string
	cookies string
}

// LoginConfig config
var LoginConfig Login

// NewLoginCommand login command
func NewLoginCommand() []cli.Command {
	return []cli.Command{
		{
			Name:      "login",
			Usage:     "登录百度网盘",
			UsageText: appName + " login [OPTIONS]",
			Action:    loginAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "bduss",
					Usage:       "bdn Cookie",
					Destination: &LoginConfig.bduss,
				},
				cli.StringFlag{
					Name:        "stoken",
					Usage:       "bdn Cookie",
					Destination: &LoginConfig.stoken,
				},
				cli.StringFlag{
					Name:        "cookies",
					Usage:       "bdn Cookie",
					Destination: &LoginConfig.cookies,
				},
			},
			After: configSaveFunc,
		},
	}
}

func loginAction(c *cli.Context) error {
	//通过 cookie 登录
	var (
		bduss   = LoginConfig.bduss
		stoken  = LoginConfig.stoken
		cookies = LoginConfig.cookies
		bdnInfo *config.BdnInfo
		err     error
	)

	if bduss != "" && stoken != "" {
		bdnInfo, err = config.Instance.SetBdnInfoByBdussAndStoken(bduss, stoken)
		if err != nil {
			return err
		}
		fmt.Printf("百度网盘登录验证登录成功, 昵称:%s,bdstoken:%s", bdnInfo.Username, bdnInfo.Bdstoken)
		return nil
	}
	if cookies != "" {
		bduser, err := config.Instance.SetBdnInfoByCookies(cookies)
		existUser, err := service.GetUserByUK(bduser.UK)
		curtime := time.Now()
		if nil == existUser {
			service.InsertUser(bduser.Username, bduser.UK, curtime, curtime, cookies)
		} else {
			service.UpdateUser(bduser.UK, cookies, curtime)
		}
		if err != nil {
			return err
		}
		fmt.Printf("百度网盘登录验证登录成功, 昵称:%s,", bduser.Username)
		return nil
	}

	return errors.New("请输入登录凭证信息")
}

func usersAction(c *cli.Context) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Uk", "昵称"})
	bdnInfo := config.Instance.BdnInfo
	table.Append([]string{strconv.Itoa(0), strconv.Itoa(bdnInfo.Uk), bdnInfo.Username})
	table.Render()
	return nil
}
