/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package healthcheck

import (
	"github.com/polarismesh/polaris-server/plugin"
	"time"
)

// Config 健康检查配置
type Config struct {
	Open             bool                   `yaml:"open"`
	Service          string                 `yaml:"service"`
	SlotNum          int                    `yaml:"slotNum"`
	LocalHost        string                 `yaml:"localHost"`
	MinCheckInterval time.Duration          `yaml:"minCheckInterval"`
	MaxCheckInterval time.Duration          `yaml:"maxCheckInterval"`
	Checkers         []plugin.ConfigEntry   `yaml:"checkers"`
	Batch            map[string]interface{} `yaml:"batch"`
}

const (
	minCheckInterval = 1 * time.Second
	maxCheckInterval = 30 * time.Second
)

// SetDefault 设置默认值
func (c *Config) SetDefault() {
	if len(c.Service) == 0 {
		c.Service = "polaris.checker"
	}
	if c.SlotNum == 0 {
		c.SlotNum = 30
	}
	if c.MinCheckInterval == 0 {
		c.MinCheckInterval = minCheckInterval
	}
	if c.MaxCheckInterval == 0 {
		c.MaxCheckInterval = maxCheckInterval
	}
	if c.MinCheckInterval > c.MaxCheckInterval {
		c.MinCheckInterval = minCheckInterval
		c.MaxCheckInterval = maxCheckInterval
	}
}
