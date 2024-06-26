package config

import "strconv"

type Mysql struct {
	Host     string `yaml:"host"`      //服务地址
	Port     int    `yaml:"port"`      //端口
	DB       string `yaml:"db"`        //数据库名
	Username string `yaml:"username"`  //数据库用户名
	Password string `yaml:"password"`  //数据库密码
	Config   string `yaml:"config"`      //高级配置,例如charset
	LogLevel string `yaml:"log_level"` // 日志等级，debug就是输出全部sql，dev开发环境, release线上环境
	MaxIdleConns int `json:"max-idle-conns" yaml:"max-idle-conns"` //空闲的最大连接数
	MaxOpenConns int `json:"max-open-conns" yaml:"max-open-conns"` //打开到数据库的最大连接数
	LogMode  string `yaml:"log-mode"`  //是否开启Gorm全局日志
}

//mysql连接配置，用于给gorm用
func (m *Mysql) Dsn() string {
	//Port是一个数字，拼成字符串时要转换一下格式
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}
