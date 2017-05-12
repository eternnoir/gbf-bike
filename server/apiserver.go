package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/eternnoir/gbf-bike/bike"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var log = logrus.New().WithFields(logrus.Fields{
	"system": "GBF-Bike",
	"module": "APIServer",
})

var (
	ErrConvertTimeout = errors.New("Can not convert timeout value")
)

var DefaultTimeout = "5"

func NewApi(port string) *ApiServer {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return &ApiServer{hs: e, revChanList: make(map[chan (*bike.BattleInfo)]bool), port: port}
}

type ApiServer struct {
	hs          *echo.Echo
	port        string
	revChanList map[chan (*bike.BattleInfo)]bool
}

func (api *ApiServer) Start() error {
	e := api.hs
	e.GET("/query", api.query)
	e.Logger.Fatal(e.Start(":" + api.port))
	return nil
}

func (api *ApiServer) NewBattleInfo(battleInfo *bike.BattleInfo) error {
	log.Debugf("%#v", battleInfo)
	for infoCh, alive := range api.revChanList {
		if !alive {
			delete(api.revChanList, infoCh)
			continue
		}
		go func() {
			select {
			case infoCh <- battleInfo:
				log.Debugf("Push battle info to channel %#v", battleInfo)
			}
		}()
	}
	return nil
}

func (api *ApiServer) query(c echo.Context) error {
	// Get team and member from the query string
	level := c.QueryParam("level")
	mobs := c.QueryParam("mobs")
	timeoutStr := c.QueryParam("timeout")
	if timeoutStr == "" {
		timeoutStr = DefaultTimeout
	}
	timeout, err := convertTimeout(timeoutStr)
	if err != nil {
		return ErrConvertTimeout
	}
	recCh := make(chan (*bike.BattleInfo), 50)
	api.revChanList[recCh] = true
	stopCh := make(chan bool)
	go func() {
		time.Sleep(time.Second * time.Duration(timeout))
		stopCh <- true
	}()
	result := api.pushBattleToList(level, mobs, recCh, stopCh)
	close(stopCh)
	go func() {
		time.Sleep(3 * time.Second)
		close(recCh)
	}()
	return c.JSON(http.StatusOK, result)
}

func (api *ApiServer) pushBattleToList(level, mobs string, ch chan (*bike.BattleInfo), stopCh chan bool) []*bike.BattleInfo {
	list := make([]*bike.BattleInfo, 0)
	mobMap := convertStringMap(mobs)
	levelMap := convertStringMap(level)
	for {
		select {
		case <-stopCh:
			api.revChanList[ch] = false
			return list
		default:
		}
		bi := <-ch
		if isWantedMission(levelMap, mobMap, bi) {
			log.Debug("Push BattleInfo to list")
			list = append(list, bi)
		}
	}
	return list
}

func isWantedMission(levels, mobs map[string]struct{}, info *bike.BattleInfo) bool {
	return isWantedLevel(levels, info) && isWantedMob(mobs, info)
}

func isWantedMob(mobs map[string]struct{}, info *bike.BattleInfo) bool {
	if len(mobs) == 0 {
		return true
	}
	_, ok := mobs[info.MobName]
	return ok
}

func isWantedLevel(levels map[string]struct{}, info *bike.BattleInfo) bool {
	if len(levels) == 0 {
		return true
	}
	_, ok := levels[info.Level]
	return ok
}

func convertStringMap(mobstr string) map[string]struct{} {
	if mobstr == "" {
		return make(map[string]struct{})
	}
	moblist := strings.Split(mobstr, ";")
	ret := make(map[string]struct{})
	for _, mobname := range moblist {
		ret[mobname] = struct{}{}
	}
	return ret
}

func convertTimeout(timeout string) (int64, error) {
	i, err := strconv.Atoi(timeout)
	if err != nil {
		return int64(0), err
	}
	return int64(i), nil
}
