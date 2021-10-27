package cobra

import (
	"github.com/zhixian0949/five/framework"
)

// SetContainer 设置服务容器
func (c *Command) SetContainer(container framework.Container) {
	c.container = container
}

// GetContainer 获取容器
func (c *Command) GetContainer() framework.Container {
	return c.Root().container
}

type CronSpec struct {
	Type        string
	Cmd         *Command
	Spec        string
	ServiceName string
}

// func (c *Command) AddCronCommand(spec string, cmd *Command) {
// 	root := c.Root()
// 	if root.Cron == nil {
// 		root.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
// 		root.CronSpecs = []CronSpec{}
// 	}
// 	root.CronSpecs = append(root.CronSpecs, CronSpec{
// 		Type: "normal-cron",
// 		Cmd:  cmd,
// 		Spec: spec,
// 	})

// 	var cronCmd Command
// 	ctx := root.Context()
// 	cronCmd = *cmd
// 	cronCmd.args = []string{}
// 	cronCmd.setPa
// }
