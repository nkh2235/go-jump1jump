package request

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	. "github.com/youfu9527/go-jump1jump/crypto"
	"hdgit.com/golang/common/errors"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type ActionData struct {
	GameData string `json:"game_data"`
	Score    int64  `json:"score"`
	Times    int64  `json:"times"`
}

type GameData struct {
	Seed      int64           `json:"seed"`
	Action    [][]interface{} `json:"action"`
	TouchList [][]interface{} `json:"touchLisst"`
	MusicList []bool          `json:"musicList"`
	Version   string          `json:"version"`
	Steps     [][]interface{} `json:"steps"`
}

type Base struct {
	Session string `json:"session_id"`
	Fast    int    `json:"fast"`
}

type Req struct {
	Base   Base        `json:"base_req"`
	Action interface{} `json:"action_data"`
}

type Resp struct {
	Base struct {
		Code int    `json:"errcode"`
		Ts   string `json:"ts"`
	} `json:"base_resp"`
}

func Gogogo(score int64, sessionid string) error {
	seed, _ := strconv.Atoi(fmt.Sprintf("%v", time.Now().UnixNano())[0:16])
	gameData := GameData{
		Seed:      int64(seed),
		Action:    make([][]interface{}, 0),
		TouchList: make([][]interface{}, 0),
		MusicList: make([]bool, 0),
		Version:   "9",
		Steps:     make([][]interface{}, 0),
	}
	var m Math
	for i := score; i > 0; i-- {
		gameData.Action = append(gameData.Action,
			[]interface{}{m.Random().ToFixed(3), m.Random().Multiply(2).ToFixed(2), i/3 == 0})
		gameData.MusicList = append(gameData.MusicList, false)
		gameData.TouchList = append(gameData.TouchList,
			[]interface{}{m.Random().Multiply(10).Minus(300).ToFixed(4), m.Random().Multiply(20).Minus(700).ToFixed(4)})
	}

	bs, err := json.Marshal(gameData)
	if err != nil {
		log.Println(`[Debug]: `, err.Error())
		return err
	}

	action := ActionData{
		GameData: string(bs),
		Score:    score,
		Times:    int64(rand.Intn(109)),
	}

	bs, err = json.Marshal(action)
	if err != nil {
		log.Println(`[Debug]: `, err.Error())
		return err
	}

	key := []byte(sessionid[0:16])
	bs, err = Encrypt(bs, key, key)
	if err != nil {
		log.Println(`[Debug]: `, err.Error())
		return err
	}

	bs, err = json.Marshal(&Req{
		Base: Base{
			Session: sessionid,
			Fast:    1,
		},
		Action: base64.StdEncoding.EncodeToString(bs),
	})
	if err != nil {
		log.Println(`[Debug]:失败了，  `, err.Error())
		return err
	}
	return DoIt(string(bs))
}

func DoIt(data string) (err error) {
	s := gorequest.New()
	_, r, _ := s.Post(`https://mp.weixin.qq.com/wxagame/wxagame_settlement`).
		Set("charset", "utf-8").
		Set("Accept-Encoding", "gzip").
		Set("referer", "https://servicewechat.com/wx7c8d593b2c3a7703/6/page-frame.html").
		Set("content-type", "application/json").
		Set("User-Agent", "MicroMessenger/6.6.1.1220(0x26060133) NetType/WIFI Language/zh_CN").
		Send(data).End()
	var resp Resp
	json.Unmarshal([]byte(r), &resp)
	if resp.Base.Code == 0 {
		log.Println(`[Debug] 恭喜你～修改成功了哦`)
		return nil
	}
	log.Println(`[Debug] 哎呀～失败了哦`)
	return errors.New("哎呀～失败了哦")
}
