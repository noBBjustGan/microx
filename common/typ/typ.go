package typ

type Header struct {
	Token      string // 登录令牌
	UserId     int64  // 用户id
	AppId      int    // appid
	AppVersion string // 客户端版本
	OsType     string // 操作系统类型
	OsVersion  string // 操作系统版本
	Resolution string // 分辨率
	Model      string // 机型信息
	Channel    string // 渠道信息
	Net        string // 网络类型
	DeviceId   string // 设备Id
}
