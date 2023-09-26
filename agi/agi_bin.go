package agi

import (
	"context"
	"os"

	"go.uber.org/zap"
)

type ConfEnv struct {
	ConfigDir  string
	ConfigFile string
	ModuleDir  string
	SpoolDir   string
	MonitorDir string
	VarDir     string
	DataDir    string
	LogDir     string
	AGIDir     string
	KeyDir     string
	RunDir     string
}

func LoadConfEnv() ConfEnv {
	env := func(s string) string {
		v, _ := os.LookupEnv(s)
		return v
	}
	return ConfEnv{
		ConfigDir:  env("AST_CONFIG_DIR"),
		ConfigFile: env("AST_CONFIG_FILE"),
		ModuleDir:  env("AST_MODULE_DIR"),
		SpoolDir:   env("AST_SPOOL_DIR"),
		MonitorDir: env("AST_MONITOR_DIR"),
		VarDir:     env("AST_VAR_DIR"),
		DataDir:    env("AST_DATA_DIR"),
		LogDir:     env("AST_LOG_DIR"),
		AGIDir:     env("AST_AGI_DIR"),
		KeyDir:     env("AST_KEY_DIR"),
		RunDir:     env("AST_RUN_DIR"),
	}
}

func Run(handler HandlerFunc) {
	ctx := context.Background()
	session, err := NewSession(ctx, os.Stdin, os.Stdout, nil)
	if err != nil {
		zap.S().With("err", err).Error("init session failed")
		os.Exit(1)
	}
	handler(session)
}
