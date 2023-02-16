package enum

type PlayerStatus map[string]int

const (
	LifeStatusDead = iota
	LifeStatusSurvival
	LifeStatusNearDeath
)

const (
	PlayerStatusLive          = "Live"          // 存活状态
	PlayerStatusLifeValue     = "LifeValue"     // 体力值
	PlayerStatusLifeLimit     = "LifeLimit"     // 体力上限
	PlayerStatusKillCount     = "KillCount"     // 出杀次数
	PlayerStatusCardsLimit    = "CardsLimit"    // 手牌上限
	PlayerStatusDrawnCardsNum = "DrawnCardsNum" // 摸牌数量
)

func DefaultPlayerStatus(lifeValue, lifeLimit int) PlayerStatus {
	return map[string]int{
		PlayerStatusLive:          LifeStatusSurvival,
		PlayerStatusLifeValue:     lifeValue,
		PlayerStatusLifeLimit:     lifeLimit,
		PlayerStatusKillCount:     1,
		PlayerStatusCardsLimit:    lifeValue,
		PlayerStatusDrawnCardsNum: DefaultDrawnCardsNum,
	}
}
