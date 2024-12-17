package keeper

import (
	"time"

	"github.com/wwqdrh/launcher/internal/app"
)

type Keeper struct {
	appManager *app.Manager
	interval   time.Duration
}

func NewKeeper(appManager *app.Manager) *Keeper {
	return &Keeper{
		appManager: appManager,
		interval:   time.Minute, // 每分钟检查一次
	}
}

func (k *Keeper) Start() {
	ticker := time.NewTicker(k.interval)
	for range ticker.C {
		k.checkAndRestart()
	}
}

func (k *Keeper) checkAndRestart() {
	// 实现检查进程是否存在的逻辑
	// 如果进程不存在，则重新启动
}
