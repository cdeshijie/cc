package settings

//settings用来读取配置文件的信息，使用viper包
import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig) //这里声明一个全局指针来进行后面的初始化和使用

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	MachineID    int64  `mapstructure:"machine_id"`
	StartTime    string `mapstructure:"start_time"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct{} //这里作为实例，没写参数,主要是作为上面结构体的参数
type MySQLConfig struct{}
type RedisConfig struct{}

func Init() (err error) {
	//两种文件读取方式
	//viper.SetConfigFile("./config.yaml"),这行可以直接替代下面两行，作用相同,这里可以使用相对路径，也可以使用绝对路径比如：F:\goprojects\gin_demo2\config.yaml

	viper.SetConfigName("config") //配置指定文件名称不带后缀

	//配置指定文件的后缀，如果是json这里就是json，这里第二种方法可以不指定文件类型，也可以指定为yaml或者json，不指定就会在指定路径进行搜寻名字相同的，后缀不同名字相同，先找到哪个算哪个，不会读第二个
	//SetConfigType一般不用，一般是配合远程读取的时候才使用，远程传输是字节流，需要告诉viper说明格式来进行字节流解析
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".") //配置文件路径（这里使用相对路径）,一个.代表当前路径,假如有config包，就是./config；而且path可以不唯一，不同路径相同名字，按路径顺序找，先找到为止

	err = viper.ReadInConfig() //从文件读取配置信息

	//viper.Unmarshal(conf),假设有个结构体可以用来存配置文件的信息，可以用unmarshal来进行反序列化为结构体，结构体的成员标签都是mapstructure，配置文件类型在上面那里改，而不是改标签
	if err != nil {
		fmt.Println("read config err:", err)
		return
		//panic(fmt.Errorf("read config err:%s \n", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改！！！")
	})
	return
	/*r := gin.Default()
	err = r.Run(fmt.Sprintf(":%d", viper.Get("port")))
	if err != nil {
		panic(err)
	}*/
}
