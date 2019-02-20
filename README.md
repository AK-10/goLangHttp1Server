# goLangHttp1Server
Http1.0 server with golang


Go言語でhttp/1.1に対応したwebサーバを実装する
実装を経験してwebサーバがどう動いているのか, httpとはどういうものかを学ぶ

## 制限
- net/httpは使わない, netだけ

## 完了
- getメソッドの対応

未実装
- 多分複数のリクエストを捌けない(goroutine 使えば良い？)
- post, put, delete
- cookie
