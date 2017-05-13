package bike

import (
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var log = logrus.WithFields(logrus.Fields{
	"system": "GBF-Bike",
	"module": "GbfBike",
})

type GbfBike struct {
	twitterClient     *twitter.Client
	battleReceivers   []BattleInfoReceiver
	receiverListMutex *sync.Mutex
}

func NewGbfBike(consumerKey, consumerSecret, accessToken, accessSecret string) (*GbfBike, error) {
	log.Infof("%s, %s, %s, %s", consumerKey, consumerSecret, accessToken, accessSecret)
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	tc := twitter.NewClient(httpClient)
	gb := &GbfBike{
		twitterClient:     tc,
		battleReceivers:   make([]BattleInfoReceiver, 0),
		receiverListMutex: &sync.Mutex{},
	}
	return gb, nil
}

func (gb *GbfBike) AddBattleReceiver(r BattleInfoReceiver) {
	gb.receiverListMutex.Lock()
	defer gb.receiverListMutex.Unlock()
	gb.battleReceivers = append(gb.battleReceivers, r)
}

func (gb *GbfBike) Start() error {
	log.Info("Start GBFBike")
	params := &twitter.StreamFilterParams{
		Track:         []string{"ID", "Lv"},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := gb.twitterClient.Streams.Filter(params)
	if err != nil {
		return err
	}
	for message := range stream.Messages {
		twit, ok := message.(*twitter.Tweet)
		if !ok {
			log.Warnf("Get not Twitte. %#v", message)
			continue
		}
		btInfo, err := ConvertGBFBattleInfo(twit.Text)
		if err != nil {
			log.Debugf("Cannot convert %s to battle Info. Error %s", twit.Text, err.Error())
			continue
		}
		btInfo.Id = twit.ID
		btInfo.CreateAt = twit.CreatedAt
		if twit.User != nil {
			btInfo.Creator = twit.User.ScreenName
		}
		go gb.triggerReceivers(btInfo)
	}
	return nil
}

func (gb *GbfBike) triggerReceivers(btinfo *BattleInfo) {
	for _, rec := range gb.battleReceivers {
		go triggerReceiver(rec, btinfo)
	}
}

func triggerReceiver(rev BattleInfoReceiver, binfo *BattleInfo) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Recovery in %#v. %#v. %#v", r, rev, binfo)
		}
	}()
	err := rev.NewBattleInfo(binfo)
	if err != nil {
		log.Errorf("Rev %#v Get Error %s. With Info %#v", rev, err.Error(), binfo)
	}

}
