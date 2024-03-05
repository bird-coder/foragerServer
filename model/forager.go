package model

// Goods [...]
type Goods struct {
	ID     int     `gorm:"primaryKey;column:id;type:int;not null" json:"id"`
	ShopID int     `gorm:"index:shop_id;column:shop_id;type:int;not null;default:0" json:"shopId"` // 店铺id
	Name   string  `gorm:"column:name;type:varchar(50);not null;default:''" json:"name"`           // 商品名
	Thumb  string  `gorm:"column:thumb;type:varchar(200);not null;default:''" json:"thumb"`        // 商品缩略图
	Intro  string  `gorm:"column:intro;type:varchar(300);not null;default:''" json:"intro"`        // 商品简介
	Price  float64 `gorm:"column:price;type:decimal(6,2);not null;default:0.00" json:"price"`      // 商品价格
	Active bool    `gorm:"column:active;type:tinyint(1);not null;default:0" json:"active"`         // 商品状态 0不可用 1可用
	Ctime  int     `gorm:"column:ctime;type:int;not null;default:0" json:"ctime"`                  // 创建时间
}

// Group [...]
type Group struct {
	ID      int    `gorm:"primaryKey;column:id;type:int;not null" json:"id"`
	Title   string `gorm:"column:title;type:varchar(50);not null;default:''" json:"title"`  // 小组名称
	Thumb   string `gorm:"column:thumb;type:varchar(200);not null;default:''" json:"thumb"` // 小组缩略图
	Intro   string `gorm:"column:intro;type:varchar(300);not null;default:''" json:"intro"` // 小组简介
	Ctime   int    `gorm:"column:ctime;type:int;not null;default:0" json:"ctime"`           // 小组创建时间
	Members int    `gorm:"column:members;type:int;not null;default:0" json:"members"`       // 小组成员数
	Active  bool   `gorm:"column:active;type:tinyint(1);not null;default:0" json:"active"`  // 小组状态 0可用1不可用
}

// GroupMember [...]
type GroupMember struct {
	ID      int  `gorm:"primaryKey;column:id;type:int;not null" json:"id"`
	GroupID int  `gorm:"uniqueIndex:group_user;column:group_id;type:int;not null;default:0" json:"groupId"` // 小组id
	UID     int  `gorm:"uniqueIndex:group_user;column:uid;type:int;not null;default:0" json:"uid"`          // 用户id
	Status  bool `gorm:"column:status;type:tinyint(1);not null;default:0" json:"status"`                    // 用户关联状态 0退出 1加入
	Ctime   int  `gorm:"column:ctime;type:int;not null;default:0" json:"ctime"`                             // 用户加入时间
}

// LogUserLogin [...]
type LogUserLogin struct {
	ID    int    `gorm:"primaryKey;column:id;type:int;not null" json:"id"`         // 用户登录流水id
	UID   int    `gorm:"column:uid;type:int;not null;default:0" json:"uid"`        // 用户ID
	Ctime int    `gorm:"column:ctime;type:int;not null;default:0" json:"ctime"`    // 记录时间
	IP    string `gorm:"column:ip;type:varchar(15);not null;default:''" json:"ip"` // ip
}

// Questions [...]
type Questions struct {
	ID      int    `gorm:"primaryKey;column:id;type:int;not null" json:"id"`               // 问题ID
	Qtype   int16  `gorm:"column:qtype;type:smallint;not null;default:0" json:"qtype"`     // 问题类型
	Title   string `gorm:"column:title;type:varchar(30);not null;default:''" json:"title"` // 问题标题
	Content string `gorm:"column:content;type:text;not null" json:"content"`               // 问题内容
	Ctime   int    `gorm:"column:ctime;type:int;not null;default:0" json:"ctime"`          // 添加时间
	Index   int    `gorm:"column:index;type:int;not null;default:0" json:"index"`          // 排序
}

// ReportRetention [...]
type ReportRetention struct {
	Day         int `gorm:"primaryKey;column:day;type:int;not null;default:0" json:"day"`      // Ymd
	Register    int `gorm:"column:register;type:int;not null;default:0" json:"register"`       // 注册用户数
	Retention1  int `gorm:"column:retention1;type:int;not null;default:0" json:"retention1"`   // 1日留存人数
	Retention2  int `gorm:"column:retention2;type:int;not null;default:0" json:"retention2"`   // 2日留存人数
	Retention3  int `gorm:"column:retention3;type:int;not null;default:0" json:"retention3"`   // 3日留存人数
	Retention4  int `gorm:"column:retention4;type:int;not null;default:0" json:"retention4"`   // 4日留存人数
	Retention5  int `gorm:"column:retention5;type:int;not null;default:0" json:"retention5"`   // 5日留存人数
	Retention6  int `gorm:"column:retention6;type:int;not null;default:0" json:"retention6"`   // 6日留存人数
	Retention7  int `gorm:"column:retention7;type:int;not null;default:0" json:"retention7"`   // 7日留存人数
	Retention14 int `gorm:"column:retention14;type:int;not null;default:0" json:"retention14"` // 14日留存人数
	Retention30 int `gorm:"column:retention30;type:int;not null;default:0" json:"retention30"` // 30日留存人数
}

// ReportSummary [...]
type ReportSummary struct {
	Day        int    `gorm:"primaryKey;column:day;type:int;not null;default:0" json:"day"`               // Ymd
	Utype      string `gorm:"primaryKey;column:utype;type:varchar(20);not null;default:tel" json:"utype"` // 用户来源(渠道)
	TotalUsers int    `gorm:"column:total_users;type:int;not null;default:0" json:"totalUsers"`           // 总用户数
	Registers  int    `gorm:"column:registers;type:int;not null;default:0" json:"registers"`              // 注册用户数
	Logins     int    `gorm:"column:logins;type:int;not null;default:0" json:"logins"`                    // 登录用户数
}

// SendCode [...]
type SendCode struct {
	Phone string `gorm:"primaryKey;column:phone;type:varchar(11);not null;default:''" json:"phone"` // 手机号
	Code  int    `gorm:"column:code;type:int;not null;default:0" json:"code"`                       // 验证码
	Ltime int    `gorm:"column:ltime;type:int;not null;default:0" json:"ltime"`                     // 时间
}

// Shop [...]
type Shop struct {
	ID     int     `gorm:"primaryKey;column:id;type:int;not null" json:"id"`
	Title  string  `gorm:"column:title;type:varchar(50);not null;default:''" json:"title"`                   // 店铺名
	Thumb  string  `gorm:"column:thumb;type:varchar(200);not null;default:''" json:"thumb"`                  // 店铺缩略图
	Score  float64 `gorm:"column:score;type:decimal(3,1);not null;default:0.0" json:"score"`                 // 店铺评分
	Lat    float64 `gorm:"index:lat_lng;column:lat;type:decimal(10,6);not null;default:0.000000" json:"lat"` // 维度
	Lng    float64 `gorm:"index:lat_lng;column:lng;type:decimal(10,6);not null;default:0.000000" json:"lng"` // 经度
	Active bool    `gorm:"column:active;type:tinyint(1);not null;default:0" json:"active"`                   // 是否可用 0否1是
	Ctime  int     `gorm:"column:ctime;type:int;not null;default:0" json:"ctime"`                            // 创建时间
}

// Trend [...]
type Trend struct {
	ID       int    `gorm:"primaryKey;column:id;type:int;not null" json:"id"`
	UID      int    `gorm:"index:index_uid;column:uid;type:int;not null;default:0" json:"uid"`       // 用户ID
	Content  string `gorm:"column:content;type:varchar(1000);not null;default:''" json:"content"`    // 动态内容
	Pics     string `gorm:"column:pics;type:varchar(1000);not null;default:''" json:"pics"`          // 动态图片
	Ctime    int    `gorm:"index:index_ctime;column:ctime;type:int;not null;default:0" json:"ctime"` // 发布时间
	Thumb    int    `gorm:"column:thumb;type:int;not null;default:0" json:"thumb"`                   // 点赞数
	Browse   int    `gorm:"column:browse;type:int;not null;default:0" json:"browse"`                 // 浏览数
	Comments int    `gorm:"column:comments;type:int;not null;default:0" json:"comments"`             // 评论数
	State    bool   `gorm:"column:state;type:tinyint(1);not null;default:1" json:"state"`            // 是否显示 0否 1是
}

// TrendComment [...]
type TrendComment struct {
	ID       int    `gorm:"primaryKey;column:id;type:int;not null" json:"id"`
	Tid      int    `gorm:"index:index_tid_uid;column:tid;type:int;not null;default:0" json:"tid"`   // 动态ID
	UID      int    `gorm:"index:index_tid_uid;column:uid;type:int;not null;default:0" json:"uid"`   // 用户ID
	Content  string `gorm:"column:content;type:varchar(300);not null;default:''" json:"content"`     // 评论内容
	Ctime    int    `gorm:"index:index_ctime;column:ctime;type:int;not null;default:0" json:"ctime"` // 评论时间
	Thumb    int    `gorm:"column:thumb;type:int;not null;default:0" json:"thumb"`                   // 点赞数
	Comments int    `gorm:"column:comments;type:int;not null;default:0" json:"comments"`             // 回复数
}

// TrendCommentReply [...]
type TrendCommentReply struct {
	ID      int    `gorm:"primaryKey;column:id;type:int;not null" json:"id"`
	Tcid    int    `gorm:"index:index_tcid_uid;column:tcid;type:int;not null;default:0" json:"tcid"` // 动态评论ID
	UID     int    `gorm:"index:index_tcid_uid;column:uid;type:int;not null;default:0" json:"uid"`   // 用户ID
	Target  int    `gorm:"column:target;type:int;not null;default:0" json:"target"`                  // 回复目标用户
	Content string `gorm:"column:content;type:varchar(300);not null;default:''" json:"content"`      // 评论内容
	Ctime   int    `gorm:"index:index_ctime;column:ctime;type:int;not null;default:0" json:"ctime"`  // 评论时间
	Thumb   int    `gorm:"column:thumb;type:int;not null;default:0" json:"thumb"`                    // 点赞数
}

// TrendReport [...]
type TrendReport struct {
	Tid   int   `gorm:"primaryKey;column:tid;type:int;not null;default:0" json:"tid"`        // 动态ID
	Type  int16 `gorm:"primaryKey;column:type;type:smallint;not null;default:0" json:"type"` // 举报类别
	Ctime int   `gorm:"column:ctime;type:int;not null;default:0" json:"ctime"`               // 举报时间
	Times int   `gorm:"column:times;type:int;not null;default:0" json:"times"`               // 举报次数
}

// User [...]
type User struct {
	UID         int    `gorm:"primaryKey;column:uid;type:int;not null" json:"uid"`                               // 用户ID
	Account     string `gorm:"unique;column:account;type:varchar(64);not null;default:''" json:"account"`        // 账号
	Username    string `gorm:"column:username;type:varchar(20);not null;default:''" json:"username"`             // 用户昵称
	Salt        string `gorm:"column:salt;type:varchar(6);not null;default:''" json:"salt"`                      // salt
	Password    string `gorm:"column:password;type:varchar(32);not null;default:''" json:"password"`             // 用户密码
	Phone       string `gorm:"index:index_phone;column:phone;type:varchar(20);not null;default:''" json:"phone"` // 手机号
	Realname    string `gorm:"column:realname;type:varchar(20);not null;default:''" json:"realname"`             // 真实姓名
	Cardno      string `gorm:"column:cardno;type:varchar(18);not null;default:''" json:"cardno"`                 // 身份证号
	RegIP       string `gorm:"column:reg_ip;type:varchar(15);not null;default:''" json:"regIp"`                  // 注册ip
	RegDate     int    `gorm:"column:reg_date;type:int;not null;default:0" json:"regDate"`                       // 注册时间
	Sex         bool   `gorm:"column:sex;type:tinyint(1);not null;default:1" json:"sex"`                         // 性别 1男 2女
	Birthday    int    `gorm:"column:birthday;type:int;not null;default:0" json:"birthday"`                      // 生日
	Utype       string `gorm:"column:utype;type:varchar(20);not null;default:tel" json:"utype"`                  // 用户类型(渠道)
	Intro       string `gorm:"column:intro;type:varchar(100);not null;default:''" json:"intro"`                  // 简介
	City        string `gorm:"column:city;type:varchar(50);not null;default:''" json:"city"`                     // 城市
	Avatar      string `gorm:"column:avatar;type:varchar(255);not null;default:''" json:"avatar"`                // 头像
	Vip         bool   `gorm:"column:vip;type:tinyint(1);not null;default:0" json:"vip"`                         // 是否vip 0否 1是
	State       bool   `gorm:"column:state;type:tinyint(1);not null;default:1" json:"state"`                     // 账号状态
	Online      bool   `gorm:"column:online;type:tinyint(1);not null;default:0" json:"online"`                   // 是否在线 0否 1是
	Agree       bool   `gorm:"column:agree;type:tinyint(1);not null;default:1" json:"agree"`                     // 是否同意用户协议 0否 1是
	LoginTimes  int    `gorm:"column:login_times;type:int;not null;default:0" json:"loginTimes"`                 // 登录次数
	LoginIP     string `gorm:"column:login_ip;type:varchar(15);not null;default:''" json:"loginIp"`              // 登陆ip
	LoginDate   int    `gorm:"column:login_date;type:int;not null;default:0" json:"loginDate"`                   // 最后一次登陆时间
	LoginDevice string `gorm:"column:login_device;type:varchar(50);not null;default:''" json:"loginDevice"`      // 登录设备uniqueId
	RegCity     string `gorm:"column:reg_city;type:varchar(50);not null;default:''" json:"regCity"`              // 注册城市
	LoginCity   string `gorm:"column:login_city;type:varchar(50);not null;default:''" json:"loginCity"`          // 登陆城市
}

// UserAddress [...]
type UserAddress struct {
	ID          int     `gorm:"primaryKey;column:id;type:int;not null" json:"id"`
	UID         int     `gorm:"index:index_uid;column:uid;type:int;not null;default:0" json:"uid"`                // 用户id
	Address     string  `gorm:"column:address;type:varchar(30);not null;default:''" json:"address"`               // 地址
	FullAddress string  `gorm:"column:full_address;type:varchar(50);not null;default:''" json:"fullAddress"`      // 详细地址
	City        string  `gorm:"column:city;type:varchar(30);not null;default:''" json:"city"`                     // 城市
	District    string  `gorm:"column:district;type:varchar(30);not null;default:''" json:"district"`             // 地区
	Lat         float64 `gorm:"column:lat;type:decimal(10,6);not null;default:0.000000" json:"lat"`               // 纬度
	Lng         float64 `gorm:"column:lng;type:decimal(10,6);not null;default:0.000000" json:"lng"`               // 经度
	Houseno     string  `gorm:"column:houseno;type:varchar(30);not null;default:''" json:"houseno"`               // 门牌号
	Consignee   string  `gorm:"column:consignee;type:varchar(30);not null;default:''" json:"consignee"`           // 收货人
	Phone       string  `gorm:"column:phone;type:char(11);not null;default:''" json:"phone"`                      // 收货人手机号
	Tag         string  `gorm:"column:tag;type:enum('home','company','school');not null;default:home" json:"tag"` // 标签
	IsDefault   int8    `gorm:"column:is_default;type:tinyint;not null;default:0" json:"isDefault"`               // 是否默认 0否 1是
	Ctime       int     `gorm:"column:ctime;type:int;not null;default:0" json:"ctime"`                            // 创建时间
}

// UserApple [...]
type UserApple struct {
	UID     int    `gorm:"primaryKey;column:uid;type:int;not null;default:0" json:"uid"`                           // 用户id
	AppleID string `gorm:"index:index_appleId;column:appleId;type:varchar(50);not null;default:''" json:"appleId"` // 苹果用户唯一id
	Ltime   int    `gorm:"column:ltime;type:int;not null;default:0" json:"ltime"`                                  // 绑定时间
}

// UserFeedback [...]
type UserFeedback struct {
	ID         int    `gorm:"primaryKey;column:id;type:int;not null" json:"id"`
	UID        int    `gorm:"column:uid;type:int;not null;default:0" json:"uid"`                         // 用户ID
	IP         string `gorm:"column:ip;type:varchar(15);not null;default:''" json:"ip"`                  // ip
	Address    string `gorm:"column:address;type:varchar(50);not null;default:''" json:"address"`        // 地址
	Type       bool   `gorm:"column:type;type:tinyint(1);not null;default:0" json:"type"`                // 类型 0系统自动反馈 1用户建议 2游戏bug
	Text       string `gorm:"column:text;type:text;not null" json:"text"`                                // 反馈内容
	ContactWay string `gorm:"column:contact_way;type:varchar(50);not null;default:''" json:"contactWay"` // 联系方式
	Timeline   int    `gorm:"column:timeline;type:int;not null;default:0" json:"timeline"`               // 时间
	Pic        string `gorm:"column:pic;type:varchar(200);not null;default:''" json:"pic"`               // 图片
	State      bool   `gorm:"column:state;type:tinyint(1);not null;default:1" json:"state"`              // 问题状态 1未解决 2已解决
	Ver        string `gorm:"column:ver;type:varchar(20);not null;default:''" json:"ver"`
	Md5        string `gorm:"index:md5;column:md5;type:char(32);not null;default:''" json:"md5"`
	Times      int    `gorm:"column:times;type:int;not null;default:0" json:"times"` // 反馈次数
	LastTime   int    `gorm:"column:last_time;type:int;not null;default:0" json:"lastTime"`
}

// UserField [...]
type UserField struct {
	UID       int  `gorm:"primaryKey;column:uid;type:int;not null;default:0" json:"uid"`          // 用户ID
	Follow    int  `gorm:"column:follow;type:int;not null;default:0" json:"follow"`               // 关注人数
	Fans      int  `gorm:"column:fans;type:int;not null;default:0" json:"fans"`                   // 粉丝人数
	Trends    int  `gorm:"column:trends;type:int;not null;default:0" json:"trends"`               // 发布动态数
	BindQq    bool `gorm:"column:bind_qq;type:tinyint(1);not null;default:0" json:"bindQq"`       // 是否绑定qq
	BindWx    bool `gorm:"column:bind_wx;type:tinyint(1);not null;default:0" json:"bindWx"`       // 是否绑定weixin
	BindApple bool `gorm:"column:bind_apple;type:tinyint(1);not null;default:0" json:"bindApple"` // 是否绑定apple
}

// UserFollow [...]
type UserFollow struct {
	UID   int `gorm:"primaryKey;column:uid;type:int;not null;default:0" json:"uid"`                    // 用户ID
	Fuid  int `gorm:"primaryKey;index:index_fuid;column:fuid;type:int;not null;default:0" json:"fuid"` // 关注用户ID
	Ctime int `gorm:"column:ctime;type:int;not null;default:0" json:"ctime"`                           // 关注时间
}

// UserPlatform [...]
type UserPlatform struct {
	UID      int    `gorm:"primaryKey;column:uid;type:int;not null;default:0" json:"uid"`                           // 用户id
	Platform int16  `gorm:"primaryKey;column:platform;type:smallint;not null;default:0" json:"platform"`            // 第三方平台 0wx 1qq
	UnionID  string `gorm:"index:index_unionId;column:unionId;type:varchar(50);not null;default:''" json:"unionId"` // 微信unionid
	OpenID   string `gorm:"index:index_openId;column:openId;type:varchar(50);not null;default:''" json:"openId"`    // 微信openid
	Invitor  int    `gorm:"index:index_invitor;column:invitor;type:int;not null;default:0" json:"invitor"`          // 邀请人id
	Ltime    int    `gorm:"column:ltime;type:int;not null;default:0" json:"ltime"`                                  // 绑定时间
}

// UserShield [...]
type UserShield struct {
	UID    int `gorm:"primaryKey;column:uid;type:int;not null;default:0" json:"uid"`                          // 用户ID
	Target int `gorm:"primaryKey;index:index_target;column:target;type:int;not null;default:0" json:"target"` // 被屏蔽用户
	Ctime  int `gorm:"column:ctime;type:int;not null;default:0" json:"ctime"`                                 // 屏蔽时间
}

// UserThumb [...]
type UserThumb struct {
	UID   int `gorm:"primaryKey;column:uid;type:int;not null;default:0" json:"uid"`                 // 用户ID
	Tid   int `gorm:"primaryKey;index:index_tid;column:tid;type:int;not null;default:0" json:"tid"` // 动态ID
	Tcid  int `gorm:"primaryKey;column:tcid;type:int;not null;default:0" json:"tcid"`               // 动态评论ID
	Tcrid int `gorm:"primaryKey;column:tcrid;type:int;not null;default:0" json:"tcrid"`             // 动态评论回复ID
	Ctime int `gorm:"column:ctime;type:int;not null;default:0" json:"ctime"`                        // 点赞时间
}

// Version [...]
type Version struct {
	BundleID string `gorm:"primaryKey;column:bundleId;type:varchar(50);not null;default:''" json:"bundleId"` // app包名
	Version  string `gorm:"primaryKey;column:version;type:varchar(20);not null;default:''" json:"version"`   // 版本号
	Remark   string `gorm:"column:remark;type:varchar(1000);not null;default:''" json:"remark"`              // 更新内容
	State    bool   `gorm:"column:state;type:tinyint(1);not null;default:0" json:"state"`                    // 是否可更新 0否 1是
}
