package cmd

import (
	"fmt"
	"io/ioutil"
	"jim/global"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// protocCmd represents the protoc command
var protocCmd = &cobra.Command{
	Use:   "protoc",
	Short: "build command for protoc buffers",
	Long:  `build command for protoc buffers`,
	Run: func(cmd *cobra.Command, args []string) {
		isWindows := false
		if runtime.GOOS == "windows" {
			isWindows = true
		}
		protoPath := fmt.Sprintf("%s/internal/rpc/proto", global.Root)
		if isWindows {
			protoPath = windowsReplace(protoPath)
		}
		files, err := ioutil.ReadDir(protoPath)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Print(files)
		for _, f := range files {
			println(f.Name())
			content := fmt.Sprintf("protoc --proto_path=%s/internal/rpc/proto/ --go_out=%s/internal/rpc --go-grpc_out=%s/internal/rpc %s", global.Root, global.Root, global.Root, f.Name())
			if isWindows {
				content = windowsReplace(content)
			}
			println(content)
			_ = executeCommand(content)
		}
		fmt.Println("protoc called")
	},
}

func windowsReplace(path string) string {
	return strings.Replace(path, "/", "\\", -1)
}
func executeCommand(content string) error {
	var name string
	var arg string
	switch runtime.GOOS {
	case "windows":
		name = "cmd"
		arg = "/C"
		break
	case "linux":
		name = "/bin/sh"
		arg = "-c"
	default:
		name = ""
		arg = ""
		break
	}
	c := exec.Command(name, arg, content)
	return c.Start()
}

func init() {
	rootCmd.AddCommand(protocCmd)
}
