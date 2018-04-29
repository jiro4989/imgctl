tkimgutil
================================================================================

RPGツクール用の表情タイル画像を自動生成するためのCLIツールです。

目的
--------------------------------------------------------------------------------

用意したイラストをゲーム内で使用したいと考えます。
用意したイラストは表情の差分以外は全く同じポーズをしているとします。
この時、RPGツクール内で使用できるように画像を加工するためにどのような作業が必要
になるでしょうか。

大まかな手順は下記のようになるでしょう。

1. Photoshop、GIMPといった加工ツールを開く
2. ツクールの画像規格の新規ファイルを作成する
3. 画像をトリミングしてピクセル位置を調節する。
4. 用意した画像すべてにそれらを適用する。
5. ツクールを起動してメッセージ上で画像を表示して切り替えてみて
   表情がずれたりしていないか確認する。

これらのルーチンワークはイラストをゲーム内で使用することの本質からは大きくそれて
いることで、非常に不毛な作業だと考えています。

イラストを用意して、ゲーム内で使用したいだけなのに、位置の調整をしていても楽しい
はずがありません。

本ツールはこの上記の手順を大幅に短縮し、用意したイラストを即座にゲーム内で使用で
きるように整形する機能を提供します。

使い方
--------------------------------------------------------------------------------

### Linux

`script/run.sh`を編集します。

下記のコマンドを実行します。

```shell
bash ./script/run.sh
```

### Windows

整備中
