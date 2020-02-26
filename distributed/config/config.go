package config

const(
	//parser names
	ParseCity = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile = "ProfileParser"
	NilParser = "NilParser"

	ElasticDataBase = "data_profile"
	ElasticHost = "http://192.168.99.101:9200"
	ItemSaverRpcMethod = "ItemSaverService.Save"
	CrawlServiceMethod = "CrawlService.Process"

	//qps 请求间隔
	Qps = 20
)
