package system

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	sysModel "github.com/weilaim/wm-admin/server/model/system"
	"github.com/weilaim/wm-admin/server/service/system"
	"github.com/weilaim/wm-admin/server/utils"
	"gorm.io/gorm"
)

const initOrderUser = initOrderAuthority + 1

type initUser struct{}

func init() {
	fmt.Println("init-----user -------")
	system.RegisterInit(initOrderUser, &initUser{})
}

func (i *initUser) MigrateTable(c context.Context) (context.Context, error) {
	db, ok := c.Value("db").(*gorm.DB)
	if !ok {
		return c, system.ErrMissingDBContext
	}
	return c, db.AutoMigrate(&sysModel.SysUser{})
}

func (i initUser) InitializerName() string {
	return sysModel.SysUser{}.TableName()
}

func (i *initUser) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	adminPassword := utils.BcryptHash("123456")
	password := utils.BcryptHash("123456")
	entitiles := []sysModel.SysUser{
		{
			UUID:        uuid.NewV4(),
			Username:    "admin",
			Password:    adminPassword,
			NickName:    "Mr.莫",
			HeaderImg:   "https://qmplusimg.henrongyi.top/gva_header.jpg",
			AuthorityId: 888,
			Phone:       "18776803246",
			Email:       "879181054@qq.com",
		},
		{
			UUID:        uuid.NewV4(),
			Username:    "weilaim",
			Password:    password,
			NickName:    "Mr.Weilaim",
			HeaderImg:   "https://qmplusimg.henrongyi.top/gva_header.jpg",
			AuthorityId: 888,
			Phone:       "18776803246",
			Email:       "879181054@qq.com",
		},
	}

	if err = db.Create(&entitiles).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysUser{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entitiles)
	// 创建用户权限关联
	authorityEntitles, ok := ctx.Value(initAuthority{}.InitializerName()).([]sysModel.SysAuthority)
	if !ok {
		return next, errors.Wrapf(system.ErrMissingDependentContext, "创建[用户-权限]关联失败,未找到权限表初始化数据")
	}
	if err = db.Model(&entitiles[0]).Association("Authorities").Replace(authorityEntitles); err != nil {
		return next, err
	}

	if err = db.Model(&entitiles[1]).Association("Authorities").Replace(authorityEntitles[:1]); err != nil {
		return next, err
	}
	return next, nil
}

func (i *initUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}

	return db.Migrator().HasTable(&sysModel.SysUser{})
}

func (i *initUser) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}

	var record sysModel.SysUser
	if errors.Is(db.Where("username= ?", "weilaim").Preload("Authorities").First(&record).Error, gorm.ErrRecordNotFound) {
		// 判断是否存在数据
		return false
	}
	return len(record.Authorities) > 0 && record.Authorities[0].AuthorityId == 888

}
