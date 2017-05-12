package bike

type BattleInfoReceiver interface {
	NewBattleInfo(battleInfo *BattleInfo) error
}
