package bike

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	battlePrefix = "参加者募集！参戦ID："
	battleReg    = regexp.MustCompile(`参加者募集！参戦ID：\w{8}\nLv\d{2}`)
)

type BattleInfo struct {
	Id       int64  `json:"id"`
	Level    string `json:"level"`
	RoomId   string `json:"roomId"`
	MobName  string `json:"mobName"`
	Creator  string `json:"creator"`
	CreateAt string `json:"createAt"`
}

func IsGBFBattle(msg string) bool {
	return battleReg.MatchString(msg)
}

func ConvertGBFBattleInfo(msg string) (*BattleInfo, error) {
	if msg == "" {
		return nil, errors.New("Msg is empty")
	}
	if !IsGBFBattle(msg) {
		return nil, fmt.Errorf("%s is not GBF Battle", msg)
	}
	index := strings.Index(msg, battlePrefix)
	msg = msg[index:]
	msg = strings.Replace(msg, battlePrefix, "", -1)
	msg = strings.Replace(msg, "\n", " ", -1)
	msg = strings.Replace(msg, "Lv", " ", -1)
	strs := strings.Split(msg, " ")
	strs = delete_empty(strs)
	if len(strs) > 4 {
		return nil, fmt.Errorf("%s result not match. %#v", msg, strs)
	}
	ret := &BattleInfo{
		RoomId:  strs[0],
		Level:   strs[1],
		MobName: strs[2],
	}
	return ret, nil
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
