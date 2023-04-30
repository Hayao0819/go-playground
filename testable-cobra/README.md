## 個人的最強のCobraのテンプレート

標準のテンプレートと比較してこんな感じの利点があります。

- グローバル変数でコマンドを定義するのではなく関数を用いて動的に定義している
- 標準のテンプレートと同じように`func init()`を用いてrootコマンドにサブコマンドを追加している
    - そのサブコマンドに関する処理が全てまとめられるので非常にきれいだと思う
- コマンド内から別のコマンドをかんたんに呼び出しできる
    - ネットの記事とかだと他のコマンドの`Run()`を直に叩くみたいな処理してる場合があるけど、`RunE`を用いている場合も多い
    - 直接`Run()`を叩くと`PreRun()`が実行されないので整合性が取れない場合がある
    - なので適切なI/Oを設定してあげた上で`Execute()`を実行してあげるのが正解だと思う（わからないのでいいアイディアあったら教えてください）
- I/Oを横取りしてかんたんにテストができる

### 現状の課題(?)

- [root.go](./cmd/root.go)の`Execute()`はどうやってテストするんだろう...
- 

### 参考サイト

ほとんど、先人様のサイトの内容を書き換えただけになっていたりします。ありがとうございます。

- [【Golang】cobraでコマンドラインツール\(CLI\)を作る \- Simple minds think alike](https://simple-minds-think-alike.moritamorie.com/entry/golang-cobra)
- [Golangのコマンドライブラリcobraを使って少しうまく実装する \- Qiita](https://qiita.com/tkit/items/3cdeafcde2bd98612428)
- [Cobra でテストしやすい CLI を構成する](https://zenn.dev/spiegel/articles/20201018-cli-with-cobra-and-golang)
- [spf13/cobra を testable に使う \- chroju\.dev/blog](https://chroju.dev/blog/go_cobra_testable_refinement)
- [Cobra の使い方とテスト \| text\.Baldanders\.info](https://text.baldanders.info/golang/using-and-testing-cobra/)
- [go \- Cobra Commander: How to call a command from another command? \- Stack Overflow](https://stackoverflow.com/questions/43747075/)
- [is there a way to add flags dynamically? · Issue \#1758 · spf13/cobra](https://github.com/spf13/cobra/issues/1758)
