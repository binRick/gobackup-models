package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	//sobackup "github.com/huacnlee/gobackup"
	log "github.com/sirupsen/logrus"
)

type GobackupModelConfig struct {
	Storage     GobackupStorage     `yaml:"store_with";`
	Archive     GobackupArchive     `yaml:"archive";`
	Compression GobackupCompression `yaml:"compress_with";`
	Databases   []GobackupDatabase  `yaml:"databases";`
	Encryption  GobackupEncryption  `yaml:"encrypt_with";`
}

type GobackupConfig struct {
	MySQL_Schema     GobackupModelConfig `yaml:"mysql_schema";`
	MySQL_Data       GobackupModelConfig `yaml:"mysql_data";`
	Etc              GobackupModelConfig `yaml:"etc";`
	VarLog           GobackupModelConfig `yaml:"var_log";`
	VarLibPersistent GobackupModelConfig `yaml:"var_lib_persistent";`
}
type GobackupModels struct {
	Models GobackupConfig `yaml:"models";`
}

type GobackupStorage struct {
	Keep int    `yaml:keep;`
	Type string `yaml:type;`
	Path string `yaml:path;`
}
type GobackupEncryption struct {
}
type GobackupCompression struct {
	Type string `yaml:type;`
}
type GobackupDatabase struct {
}
type GobackupArchive struct {
	Excludes []string `yaml:excludes;`
	Includes []string `yaml:includes;`
}
