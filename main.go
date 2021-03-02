package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	//sobackup "github.com/huacnlee/gobackup"
	log "github.com/sirupsen/logrus"
)

var (
	DEBUG_GOBACKUP_EVENTS = false
)

func run_gobackup_config_test() {
	gobackup_config_object := get_gobackup_config_object()
	fmt.Printf("get_gobackup_config_object=%v", gobackup_config_object)

	gobackup_config_yaml_bytes := get_gobackup_cfg_bytes()
	fmt.Printf("gobackup_config_yaml_bytes=%s", gobackup_config_yaml_bytes)
}

func get_gobackup_cfg_bytes() []byte {
	gobackup_cfg_str, err := yaml.Marshal(get_gobackup_config_object())
	if err != nil {
		log.Fatal(err)
	}
	return gobackup_cfg_str
}

func get_gobackup_config_object() *GobackupModels {
	gobackup_yaml_cfg := GobackupModels{
		Models: GobackupConfig{
			Etc: GobackupModelConfig{
				Compression: GobackupCompression{
					Type: "tgz",
				},
				Encryption: GobackupEncryption{},
				Storage: GobackupStorage{
					Type: "local",
					Keep: 2,
					Path: "/backups/etc",
				},
				Archive: GobackupArchive{
					Includes: []string{
						"/etc/supervisord.conf",
						"/etc/supervisord.d",
					},
					Excludes: []string{},
				},
			},
			VarLibPersistent: GobackupModelConfig{
				Compression: GobackupCompression{
					Type: "tar",
				},
				Encryption: GobackupEncryption{},
				Storage: GobackupStorage{
					Type: "local",
					Keep: 2,
					Path: "/backups/var_lib_persistent",
				},
				Archive: GobackupArchive{
					Includes: []string{
						"/var/lib/persistent",
					},
					Excludes: []string{
						"/var/lib/persistent/collectd",
						"/var/lib/persistent/mysql",
					},
				},
			},
			VarLog: GobackupModelConfig{
				Compression: GobackupCompression{
					Type: "tgz",
				},
				Encryption: GobackupEncryption{},
				Storage: GobackupStorage{
					Type: "local",
					Keep: 2,
					Path: "/backups/var_log",
				},
				Archive: GobackupArchive{
					Includes: []string{
						"/var/log",
					},
					Excludes: []string{
						"/var/log/xxxxxxxxxx",
					},
				},
			},
			MySQL_Data: GobackupModelConfig{
				Compression: GobackupCompression{
					Type: "tgz",
				},
				Encryption: GobackupEncryption{},
				Storage: GobackupStorage{
					Type: "local",
					Keep: 1000,
					Path: "/backups/mysql_data",
				},
				Archive: GobackupArchive{
					Includes: []string{},
					Excludes: []string{},
				},
			},
			MySQL_Schema: GobackupModelConfig{
				Compression: GobackupCompression{
					Type: "tgz",
				},
				Encryption: GobackupEncryption{},
				Storage: GobackupStorage{
					Type: "local",
					Keep: 1000,
					Path: "/backups/mysql_schema",
				},
				Archive: GobackupArchive{
					Includes: []string{},
					Excludes: []string{},
				},
			},
		},
	}
	return &gobackup_yaml_cfg
}
