package Kg

import "github.com/gin-gonic/gin_bak/json"

//创建任务 create_task api url
const CREATE_TASK_URL = "https://aip.baidubce.com/rest/2.0/kg/v1/pie/task_create"
//更新任务 update_task api url
const UPDATE_TASK_URL  = "https://aip.baidubce.com/rest/2.0/kg/v1/pie/task_update"
//获取任务详情 task_info api url
const TASK_INFO_URL  = "https://aip.baidubce.com/rest/2.0/kg/v1/pie/task_info"
//以分页的方式查询当前用户所有的任务信息 task_query api url
const TASK_QUERY_URL  = "https://aip.baidubce.com/rest/2.0/kg/v1/pie/task_query"
//启动任务 task_start api url
const TASK_START_URL  = "https://aip.baidubce.com/rest/2.0/kg/v1/pie/task_start"
//查询任务状态 task_status api url
const TASK_STATUS_URL  = "https://aip.baidubce.com/rest/2.0/kg/v1/pie/task_status"

type Task struct {
	name string
	template_content string
	input_mapping_file string
	output_file string
	url_pattern string
	limit_count int64
}


type Result struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	LogID            int64 `json:"log_id"`
	EntityAnnotation []Entity `json:"entity_annotation"`
}


type Entity struct {
	Status     string `json:"status"`
	Confidence string `json:"confidence"`
	Concept    struct {
		Level1 string `json:"level1"`
		Level2 string `json:"level2"`
	} `json:"concept"`
	BdbkKgID string `json:"_bdbkKgId"`
	Mention  string `json:"mention"`
	BdbkURL  string `json:"_bdbkUrl"`
	Offset   string `json:"offset"`
	Desc     string `json:"desc"`
}

type ErrResult struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	LogID     int64  `json:"log_id"`
}

//var errList = make(map[int64]string)

var errList=map[int64]string{100:"Invalid parameter",200:"Invalid parameter"}

func (task Task) CreateTask(name , templateContent , inputMappingFile, outputFile, urlPattern string, limitCount int64)(result Result){
	//var reqeustTask = new(Task)
	reqeustTask :=Task{
		name:name,
		template_content:templateContent,
		input_mapping_file:inputMappingFile,
		output_file:outputFile,
		url_pattern:urlPattern,
		limit_count:limitCount,
		}

	return
}

func validate(url string,request interface{}) bool{
	return true
}

func proccessResult(data string){

}

func request(url string, data interface{}){
	if ret :=validate(url, data); !ret {
		return
	}
	requestStr,err := json.Marshal(data)
	if err!=nil {
		return
	}
}