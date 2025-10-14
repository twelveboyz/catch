package bump

import (
	run "Catch/bump/comparison/execute"
	"Catch/bump/internal/bootstrap"
	"Catch/bump/internal/utils"
	"Catch/bump/pull/compute/ecs"
	"Catch/bump/pull/compute/ironic"
	"github.com/sirupsen/logrus"
)

func Run() {
	f := utils.LogFunc("model")
	defer f.Close()

	bootstrap.Token = utils.Cache()

	err := utils.DetermineTokenValid()
	if err != nil {
		logrus.Errorln(err)
		return
	} else {
		logrus.Info("Valid Token!")
	}

	root := utils.InputRoot()

	logrus.SetLevel(utils.GetLogLevel())

	//var root = `D:\桌面\上云\D010_中间号IT凤凰节点资源`

	run.ComputeRun(root)

	run.StorageRun(root)

	run.NetworkRun(root)
}

func RunFind() {
	f := utils.LogFunc("model")
	defer f.Close()

	bootstrap.Token = utils.Cache()

	err := utils.DetermineTokenValid()
	if err != nil {
		logrus.Errorln(err)
		return
	} else {
		logrus.Info("Valid Token!")
	}

	choose := utils.InputChoose()
	switch choose {
	case "1":
		ecs.ExecuteMatchByName()
		ironic.ExecuteMatchByName()
	case "2":
		ecs.ExecuteMatchByID()
		ironic.ExecuteMatchByID()
	}
}
