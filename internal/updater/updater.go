package updater

import (
	"io"
	"net/http"
	"os"

	"github.com/robfig/cron/v3"
	"github.com/wwqdrh/launcher/internal/app"
)

type Updater struct {
	appManager *app.Manager
	cron       *cron.Cron
}

func NewUpdater(appManager *app.Manager) *Updater {
	return &Updater{
		appManager: appManager,
		cron:       cron.New(),
	}
}

func (u *Updater) Start() {
	// 每天午夜检查更新
	u.cron.AddFunc("@midnight", u.checkUpdates)
	u.cron.Start()

	// 启动时检查一次更新
	u.checkUpdates()
}

func (u *Updater) checkUpdates() {
	// 实现检查更新逻辑
	// 1. 获取远程版本
	// 2. 比较本地版本
	// 3. 如果需要更新，下载新版本
	// 4. 替换可执行文件
}

func (u *Updater) downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
