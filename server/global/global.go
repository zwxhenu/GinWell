package global

import (
	"GinWell-Server/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	GW_DB     *gorm.DB
	GW_DBList map[string]*gorm.DB
	GW_REDIS  *redis.Client
	GW_CONFIG config.Server
	GW_VP     *viper.Viper
	// GVA_LOG    *oplogging.Logger
	GW_LOG *zap.Logger
	//GW_Timer               timer.Timer = timer.NewTimerTask()
	//GW_Concurrency_Control             = &singleflight.Group{}

	//BlackCache local_cache.Cache
	lock sync.RWMutex
)
