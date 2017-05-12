package bike

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBattleMsg(t *testing.T) {
	ast := assert.New(t)
	msg := `参加者募集！参戦ID：E5088E86
Lv60 青竜
https://t.co/RbqZBFIUBz`
	ast.True(IsGBFBattle(msg))
}


func TestBattleMsgFail(t *testing.T) {
	ast := assert.New(t)
	msg := `参加者募集！参戦ID：37F7B348
ALv60 リヴァイアサン・マグナ
https://t.co/RbqZBFIUBz`
	ast.False(IsGBFBattle(msg))
}


func TestConvertBattleMsg(t *testing.T) {
	ast := assert.New(t)
	msg := `参加者募集！参戦ID：37F7B348
ALv60 リヴァイアサン・マグナ
https://t.co/RbqZBFIUBz`
	result,err:= ConvertGBFBattleInfo(msg)
	if err!=nil{
		ast.Fail(err.Error())
		return
	}
	ast.Equal("37F7B348",result.RoomId )
	ast.Equal("60",result.Level)
	ast.Equal("リヴァイアサン・マグナ",result.MobName )
}
