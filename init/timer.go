package init

import (
	"HertzBoot/config"
	"HertzBoot/pkg"
	"HertzBoot/pkg/global"
	"fmt"
)

func Timer() {
	if global.CONFIG.Timer.Start {
		for _, detail := range global.CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				_, _ = global.Timer.AddTaskByFunc("ClearDB", global.CONFIG.Timer.Spec, func() {
					err := pkg.ClearTable(global.DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(detail)
		}
	}
}
