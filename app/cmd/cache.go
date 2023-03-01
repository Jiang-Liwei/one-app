package cmd

import (
	"fmt"
	"forum/pkg/cache"
	"forum/pkg/console"

	"github.com/spf13/cobra"
)

var Cache = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
}

var CacheClear = &cobra.Command{
	Use:   "clear",
	Short: "删除所有缓存",
	Run:   runCacheClear,
}

var CacheForget = &cobra.Command{
	Use:   "forget",
	Short: "清除单个key, 示例: cache forget cache-key",
	Run:   runCacheForget,
}

// forget 命令的选项
var cacheKey string

func init() {
	// 注册 cache 命令的子命令
	Cache.AddCommand(CacheClear, CacheForget)

	// 设置 cache forget 命令的选项
	CacheForget.Flags().StringVarP(&cacheKey, "key", "k", "", "KEY of the cache")
	CacheForget.MarkFlagRequired("key")
}

func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared.")
}

func runCacheForget(cmd *cobra.Command, args []string) {
	cache.Forget(cacheKey)
	console.Success(fmt.Sprintf("Cache key [%s] deleted.", cacheKey))
}
