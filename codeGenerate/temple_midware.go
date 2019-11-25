package main

var midwareTemp=`
package midware

import "context"

//定义func
type MidWareFunc func(ctx context.Context,req interface{})(interface{},error)

//定义midware链
type MidWare func(wareFunc MidWareFunc)MidWareFunc

//自定义中间件
var userMidWareLink []MidWare

//定义链传递函数
/*
  next为下一层调用函数
*/
func Chain(outer MidWare,other ...MidWare)MidWare{
	return func(next MidWareFunc) MidWareFunc {
		//next为传递的最后参数，代表处理函数本身
		//传递链参数
		for i:=len(other)-1;i>=0 ;i--  {
			//设置调用参数
			next=other[i](next)
		}
		//将最前的参数传递给outer
		return outer(next)
	}
}

//添加中间件
func AddUserMidWare(midware ...MidWare){
	userMidWareLink= append(userMidWareLink, midware...)
}

//连接中间件
func BuildUserMidWareChain(handler MidWareFunc)MidWareFunc{
    var midwareLink []MidWare
    //添加监控中间件
    midwareLink=append(midwareLink,PromeScanMidWare)
	if len(userMidWareLink)!=0 {
		midwareLink=append(midwareLink,userMidWareLink...)
	}
	headFunc:=Chain(midwareLink[0],midwareLink[1:]...)
	outerFunc:=headFunc(handler)
	return outerFunc
}
`

var midwareProme=`
package midware

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type serverContextKey struct {}

type ServerScanMeta struct {
	ServerName string
	Method string
}

//获得数据元素
func GetServerScanMeta(ctx context.Context)*ServerScanMeta{
	v,ok:=ctx.Value(serverContextKey{}).(*ServerScanMeta)
	if ok==false {
		return &ServerScanMeta{}
	}
	return v
}

//初始化元素
func InitServerScanMeta(ctx context.Context,serverName string,method string)context.Context{
	serverScan:=&ServerScanMeta{
		ServerName: serverName,
		Method:     method,
	}
	ctxNew:=context.WithValue(ctx,serverContextKey{},serverScan)
	return ctxNew
}

type ServerScanTool struct {
	requestNum *prometheus.CounterVec     //统计访问量
	errorNum *prometheus.CounterVec       //统计各个错误的次数
	methodTime *prometheus.SummaryVec     //统计访问时间的发布
}

func NewServerScanTool()*ServerScanTool{
	Tool:=&ServerScanTool{
		requestNum: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name:"wfuproject_request_num",
				Help:"the request time of wfuproject",
			},
			[]string{"service","method"}),
		errorNum:   prometheus.NewCounterVec(prometheus.CounterOpts{
			Name:"wfuproject_error_count",
			Help:"the num of every error",
		},
		[]string{"service","method","error_code"},
		),
		methodTime: prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Name:        "wfuproject_method_usetime",
			Help:        "the time len of wfuproject method",
            Objectives: map[float64]float64{0.5:0.05,0.9:0.01,0.99:0.001},
		},
		[]string{"service","method"}),
	}
	prometheus.MustRegister(Tool.requestNum)
	prometheus.MustRegister(Tool.methodTime)
	prometheus.MustRegister(Tool.errorNum)
	return Tool
}

//增加统计次数
func(s *ServerScanTool)IncrRequestTime(server string,method string){
	s.requestNum.With(prometheus.Labels{
		"service":server,
		"method":method,
	}).Inc()
}

//增加错误次数
func(s *ServerScanTool)IncrErrorTime(server string,method string,err error){
    var errMess string
    if err!=nil{
        errMess=err.Error()
    }
	s.errorNum.With(prometheus.Labels{
		"service":server,
		"method":method,
		"error_code":errMess,
	}).Inc()
}

//统计耗时
func(s *ServerScanTool)CalTimeUse(server string,method string,us int64){
	s.methodTime.With(prometheus.Labels{
		"service":server,
		"method":method,
	}).Observe(float64(us))
}


//中间件
var promeScanTool =NewServerScanTool()

func PromeScanMidWare(next MidWareFunc)MidWareFunc{
	return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		//统计次数
		serverMeta:=GetServerScanMeta(ctx)
		promeScanTool.IncrRequestTime(serverMeta.ServerName,serverMeta.Method)
		//计算耗时
		startTime:=time.Now().UnixNano()
		resp,err=next(ctx,req)
		//统计错诶次数
		promeScanTool.IncrErrorTime(serverMeta.ServerName,serverMeta.Method,err)
		//计算耗时
		promeScanTool.CalTimeUse(serverMeta.ServerName,serverMeta.Method,(time.Now().UnixNano()-startTime)/1000)
		return 
	}
}
`