package main

import (
  "fmt"
  "os"
  "strings"
)

type Environment map[string]string

func getEnvironment() Environment {
  env := make(Environment)

  for _, e := range os.Environ() {
    pair := strings.SplitN(e, "=", 2)

    env[pair[0]] = pair[1]
  }

  return env
}

type OSInfo struct {
  Hostname string
	Env      Environment
}

func (osInfo *OSInfo) fromOS() error {
  if hostname, err := os.Hostname(); err != nil {
    return fmt.Errorf("os.Hostname: %v", err)
  } else {
    osInfo.Hostname = hostname
  }

  osInfo.Env = getEnvironment()

  return nil
}
