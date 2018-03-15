# tkimgutil
RPGツクール用の画像処理ツール群です。
***開発中。まだ何も作ってない***

## 概要


## 使い方
下記のコマンドを実行します。

### 通常の使い方

```shell
find ./img -name *.png | 
  ./bin/tkimgutil scale -s 50 |
  ./bin/tkimgutil trim -x 50 -y 28 -w 144 -h 144 |
  ./bin/tkimgutil paste -r 2 -c 4
```

### 左右反転版

```shell
find ./img -name *.png | 
  ./bin/tkimgutil scale -s 50 |
  ./bin/tkimgutil trim -x 50 -y 28 -w 144 -h 144 |
  ./bin/tkimgutil flip |
  ./bin/tkimgutil paste -r 2 -c 4
```

