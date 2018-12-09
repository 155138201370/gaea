package tmpls

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/micro-plat/gaea/cmds/project/tmpls/dev"
)

var templateFiles map[string][]string
var names map[string]string
var templates map[string]string

func init() {
	templateFiles = make(map[string][]string)
	templateFiles["install.dev.go"] = []string{"api.port", "api.jwt", "db", "cache", "api.cros", "queue", "api.appconf", "api.metric"}
	templateFiles["install.prod.go"] = []string{"api.port", "api.jwt", "db", "cache", "api.cros", "queue", "api.appconf", "api.metric"}
	templateFiles["init.go"] = []string{"db.init", "queue.init", "cache.init", "appconf.func", "appconf.struct"}
	templateFiles["handling.go"] = []string{"handling.jwt"}

	names = make(map[string]string)
	names["api.port"] = apiPort
	names["api.cros"] = apiCros
	names["api.jwt"] = apiJWT
	names["api.metric"] = apiMetric
	names["db"] = db
	names["cache"] = cache
	names["queue"] = queue
	names["api.appconf"] = apiAppConf
	names["db.init"] = dbInit
	names["queue.init"] = queueInit
	names["cache.init"] = cacheInit
	names["appconf.struct"] = appconfStruct
	names["appconf.func"] = appconfFunc
	names["handling.jwt"] = handlingJWT

	templates = make(map[string]string)
	templates["api.port"] = dev.APIMainPort
	templates["api.jwt"] = dev.APISubAuth
	templates["db"] = dev.PlatVarDB
	templates["cache"] = dev.PlatVarCache
	templates["queue"] = dev.PlatVarQueue
	templates["api.cros"] = dev.APISubCros
	templates["api.metric"] = dev.APISubMetric
	templates["api.appconf"] = dev.APIApp
	templates["db.init"] = DBInit
	templates["queue.init"] = QueueInit
	templates["cache.init"] = CacheInit
	templates["appconf.struct"] = APPConfStruct
	templates["appconf.func"] = APPConfFunc
	templates["handling.jwt"] = dev.HandingJWT

}

const (
	apiPort       = `api.port`
	apiCros       = `api.cros`
	apiJWT        = `api.jwt`
	apiAppConf    = `api.appconf`
	apiMetric     = `api.metric`
	db            = `db`
	cache         = `cache`
	queue         = `queue`
	cronApp       = `cron.app`
	cronTask      = `cron.task`
	cronMetric    = `cron.Metric`
	mqcServer     = `mqc.server`
	mqcQueue      = `mqc.queue`
	mqcMetric     = `mqc.Metric`
	webPort       = `web.port`
	webStatic     = `web.static`
	webMetric     = `web.Metric`
	wsApp         = `ws.app`
	wsAuth        = `ws.auth`
	wsMetric      = `ws.Metric`
	rpcPort       = `rpc.port`
	rpcMetric     = `rpc.Metric`
	dbInit        = `db.init`
	queueInit     = "queue.init"
	cacheInit     = "cache.init"
	appconfStruct = "appconf.struct"
	appconfFunc   = "appconf.func"
	handlingJWT   = "handling.jwt"
)

//GetTmpls 获取模板
func GetTmpls(projectName string, input map[string]interface{}) (out map[string]string, err error) {
	fmt.Println("input:", input)
	//input := makeParams(projectName, serverType, port, db, jwt, domain)
	out = make(map[string]string)
	if out["main.go"], err = translate(mainTmpl, input); err != nil {
		return nil, err
	}
	if out["init.go"], err = translate(initTmpl, input); err != nil {
		return nil, err
	}
	if out["install.dev.go"], err = translate(strings.Replace(strings.Replace(installDevTmpl, "\"", "`", -1), "'", "\"", -1), input); err != nil {
		return nil, err
	}
	if out["install.prod.go"], err = translate(strings.Replace(strings.Replace(installProdTmpl, "\"", "`", -1), "'", "\"", -1), input); err != nil {
		return nil, err
	}
	if out["handling.go"], err = translate(strings.Replace(strings.Replace(handingTmpl, "\"", "`", -1), "'", "\"", -1), input); err != nil {
		return nil, err
	}
	if out[".gitignore"], err = translate(gitignoreTmpl, input); err != nil {
		return nil, err
	}
	out["modules/const/sql/sql.go"] = "package sql"
	out["services/server.go"] = "package server"

	return out, nil
}

//GetConfTmpls 获取配置模板
func GetConfTmpls(blocks []string, input map[string]interface{}) (out map[string]map[string]string, err error) {

	out = make(map[string]map[string]string)
	for fname, f := range templateFiles {
		for _, n := range f {
			for _, name := range blocks {

				if !strings.Contains(n, name) {
					continue
				}
				fmt.Printf("name:%s n: %s \n", name, n)
				if _, ok := out[fname]; !ok {
					out[fname] = make(map[string]string)
				}

				out[fname][names[n]], err = translate(strings.Replace(strings.Replace(templates[n], "\"", "`", -1), "'", "\"", -1), input)
				if err != nil {
					return nil, err
				}
			}
		}

	}

	return out, err
}

func getServices(serverType string) []string {
	s := make([]string, 0, 2)
	if strings.Contains(serverType, "api") || strings.Contains(serverType, "rpc") {
		s = append(s, "Micro")
	}
	if strings.Contains(serverType, "mqc") || strings.Contains(serverType, "cron") {
		s = append(s, "Autoflow")
	}
	if strings.Contains(serverType, "web") {
		s = append(s, "Page")
	}
	return s
}

func translate(c string, input interface{}) (string, error) {
	var tmpl = template.New("api").Funcs(makeFunc())
	np, err := tmpl.Parse(c)
	if err != nil {
		return "", err
	}
	buff := bytes.NewBufferString("")
	if err := np.Execute(buff, input); err != nil {
		return "", err
	}
	return buff.String(), nil
}

//获取生成项目的数据
func makeParams(projectName, serverType, port, db string, jwt, domain bool) map[string]interface{} {
	if !strings.Contains(port, ":") {
		port = ":" + port
	}
	return map[string]interface{}{
		"projectName": projectName,
		"serverType":  serverType,
		//	"pkgs":    gGetModulePackageName(modules),
		"rss":    getServices(serverType),
		"port":   port,
		"dbname": strings.Split(db, ":")[0],
		"db":     db,
		"jwt":    jwt,
		"domain": domain,
	}
}

func getRModules(modules []string) []string {
	nmodule := make([]string, 0, len(modules)+1)
	for _, m := range modules {
		ms := strings.Split(m, " ")
		for _, s := range ms {
			nmodule = append(nmodule, filepath.Join("/", s))
		}
	}
	return nmodule
}

func makeFunc() map[string]interface{} {
	return map[string]interface{}{
		"humpName": fGetHumpName,           //多个单词首字符大写
		"spkgName": fGetServicePackageName, //包路径
		"mpkgName": fGetModulePackageName,  //包路径
		"lName":    fGetLastName,           //取最后一个单词
		"fName":    fGetFirstName,          //取第一个单词
		"fServer":  fServer,                //判断是否有这个服务
	}
}

func fServer(s, substr string) bool {
	return strings.Contains(s, substr)
}

func fGetFirstName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	return names[0]
}

func fGetHumpName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	buff := bytes.NewBufferString("")
	for _, v := range names {
		buff.WriteString(fGetLoopHumpName(v, "."))
	}
	return strings.Replace(buff.String(), ".", "", -1)
}

func fGetLoopHumpName(n string, s string) string {
	names := strings.Split(strings.Trim(n, s), s)
	buff := bytes.NewBufferString("")
	for _, v := range names {
		buff.WriteString(strings.ToUpper(v[0:1]))
		buff.WriteString(v[1:])
	}
	return strings.Replace(buff.String(), ".", "", -1)
}

func fGetServicePackageName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	if len(names) == 1 {
		return "services"
	}
	return strings.ToLower(strings.Join(names[0:len(names)-1], "/"))
}

func fGetPackageName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	if len(names) == 1 {
		return names[0]
	}
	return strings.Join(names[0:len(names)-1], "/")
}

func fGetModulePackageName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	if len(names) == 1 {
		return "modules"
	}
	return strings.ToLower(strings.Join(names[0:len(names)-1], "/"))
}

func fGetServicePackagePath(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	if len(names) == 1 {
		return "services"
	}
	return strings.ToLower(filepath.Join("services", strings.Join(names[0:len(names)-1], "/")))
}

func fGetLastName(n string) string {
	names := strings.Split(strings.Trim(n, "/"), "/")
	return names[len(names)-1]
}

func gGetModulePackageName(module []string) []string {
	npkgs := make([]string, 0, len(module)/2)
	n := make(map[string]string)
	for _, m := range module {
		nm := fGetServicePackagePath(m)
		if _, ok := n[nm]; !ok {
			npkgs = append(npkgs, nm)
			n[nm] = nm
		}
	}
	return npkgs
}
