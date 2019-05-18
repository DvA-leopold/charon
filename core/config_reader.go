package core

var Config = make(map[string]string)

func ReadConfig(filename string) {
	Config["HTTPListen"] = "0.0.0.0:8081"
	Config["TCPListen"] = "127.0.0.1:123456"

	//data, err := ioutil.ReadFile(filename)
	//log.Println("reading config file[", filename, "]")
	//if err != nil {
	//	log.Println("cannot read file[", filename, "], err[", err, "]")
	//	return
	//}
	//str := string(data)
	//
	//Config["test"] = str
	//
	//log.Println(Config["test"])
}
