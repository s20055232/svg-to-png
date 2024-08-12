# SVG to JPEG

這是一個簡單的 CLI 工具，輸入 SVG 的路徑以及要輸出的 JPEG 路徑就可以進行轉換，但目前不支援 utf-8 相關字元，只能轉換簡單的 pixel。

```shell
$ go run . -svg=input.svg -jpeg=output.jpeg
Draw SVG to PNG
Convert success
```
