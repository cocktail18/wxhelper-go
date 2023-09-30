# wxhelper-go
[wxhelper](https://github.com/ttttupup/wxhelper) 项目的go sdk、例子

#### 免责声明:
本仓库发布的内容，仅用于学习研究，请勿用于非法用途和商业用途！如因此产生任何法律纠纷，均与作者无关！
且不保证其可用性、稳定性。


#### 使用说明:
1. 下载注入工具，如果是 3.9.5.81 可以跳过
- 3.9.2.23 是x86版本，使用第三方注入工具，请自行分辨该工具。下载后请放到 example 文件夹
- 附带地址：injector.exe 下载地址 https://github.com/nefarius/Injector
- 3.9.5.81 是x64版本，不需要使用第三方注入工具
2. 下载dll文件，自行到[wxhelper](https://github.com/ttttupup/wxhelper) 下载对应的版本，然后放到 example 文件夹
3. 修改example/main.go
```go
const (
  apiVersion = api.ApiVersionV1 // 3.9.5.81 使用v2
  dllPath    = "wxhelper.dll" // 对应的 dll 文件路径
  port       = 10086 // 监听的端口
  )
```

4. `go run main.go`
   







