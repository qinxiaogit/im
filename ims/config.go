package main

import (
	"log"
	"strconv"
)

//超级群离线消息数量限制 ，超过的部分会丢弃

const GROUP_OFFLINE_LIMIT = 100

//离线消息返回的数量限制
const OFFLINE_DEFAULT_LIMIT = 3000

const GROUP_OFFLINE_DEFAULT_LIMIT = 0

//unlimit
const OFFLINE_DEFAULT_HARD_LIMIT = 0

type StorageConfig struct {
	rpc_listen          string
	storage_root        string
	http_listen_address string

	sync_listen    string
	master_address string
	is_push_system string
	group_limit int //普通群离线消息数量限制
	limit int //单次离线消息数量限制
	hard_limit int //离线消息总的数量限制

	log_filename string
	log_level string
	log_backup int
	log_age int
	log_caller bool
}

func get_int(appConf map[string]string,key string)int64{
	concurrency , present := appConf[key]
	if !present{
		log.Fatalf("key:%s non exist", key)
	}
	n, err := strconv.ParseInt(concurrency, 10, 64)
	if err != nil {
		log.Fatalf("key:%s is't integer", key)
	}
	return n
}



func get_opt_int(app_cfg map[string]string, key string, default_value int64) int64 {
	concurrency, present := app_cfg[key]
	if !present {
		return default_value
	}
	n, err := strconv.ParseInt(concurrency, 10, 64)
	if err != nil {
		log.Fatalf("key:%s is't integer", key)
	}
	return n
}


func get_string(app_cfg map[string]string, key string) string {
	concurrency, present := app_cfg[key]
	if !present {
		log.Fatalf("key:%s non exist", key)
	}
	return concurrency
}

func get_opt_string(app_cfg map[string]string, key string) string {
	concurrency, present := app_cfg[key]
	if !present {
		return ""
	}
	return concurrency
}

func read_storage_cfg(cfg_path string) *StorageConfig {
	config := new(StorageConfig)
	app_cfg := make(map[string]string)
	err := cfg.Load(cfg_path, app_cfg)
	if err != nil {
		log.Fatal(err)
	}

	config.rpc_listen = get_string(app_cfg, "rpc_listen")
	config.http_listen_address = get_opt_string(app_cfg, "http_listen_address")
	config.storage_root = get_string(app_cfg, "storage_root")
	config.sync_listen = get_string(app_cfg, "sync_listen")
	config.master_address = get_opt_string(app_cfg, "master_address")
	config.is_push_system = get_opt_int(app_cfg, "is_push_system", 0) == 1
	config.limit = int(get_opt_int(app_cfg, "limit", OFFLINE_DEFAULT_LIMIT))
	config.group_limit = int(get_opt_int(app_cfg, "group_limit", GROUP_OFFLINE_DEFAULT_LIMIT))
	config.hard_limit = int(get_opt_int(app_cfg, "hard_limit", OFFLINE_DEFAULT_HARD_LIMIT))

	config.log_filename = get_opt_string(app_cfg, "log_filename")
	config.log_level = get_opt_string(app_cfg, "log_level")
	config.log_backup = int(get_opt_int(app_cfg, "log_backup", 0))
	config.log_age = int(get_opt_int(app_cfg, "log_age", 0))
	config.log_caller = get_opt_int(app_cfg, "log_caller", 0) != 0

	return config
}