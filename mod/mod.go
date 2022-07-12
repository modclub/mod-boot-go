package mod

type D map[string]interface{}

type Plugin interface {
	Name() string
}

func RegisterPlugin(p Plugin) {

}

func UnregisterPlugin(name string) {

}
func RegisterService() {

}

func UnregisterService() {

}

func Call() {

}

func Run() {

}
