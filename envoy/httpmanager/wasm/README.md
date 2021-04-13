# envoy 支持wasm 编译


## 生成wasm 文件

```bash
$ tinygo build -o filter.wasm -target=wasi -wasm-abi=generic .
```

## 启动 envoy

```bash
$ envoy -c wasm.yaml
```