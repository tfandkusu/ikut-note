# iKut Note

# 要件

## シーン記録

- OBS の仮想カメラを取り込んで画像認識を行い、シーンが発生した時刻を保存する。

### 対応シーン

- 試合の開始
- たおした
- やられた

試合の終了はスプラトゥーン2の時代に作ったが精度がいまいちで、試合の終了と誤認識すると、以後のたおした、やられたが無いことになってしまうので、今回は作らない。

### 画像認識モデルを作り直す

現在スプラトゥーン2のスクリーンショットで学習したものしか無いので、スプラトゥーン3のスクリーンショットで学習したモデルを新たに作成する

## シーンブラウザ

- 保存されたシーンを時系列に並べる
- ユーザは特定のディレクトリを指定する
- ユーザはシーンを選択できる
- 特定ディレクトリに選択されたシーンに該当する動画があれば、動画プレイヤーで開く

### 技術的詳細

シーンと動画の紐付けは時刻によってのみ行われる。よって、他人の動画をディレクトリに置いたら、シーンと動画がずれてしまう。
動画の録画開始時刻を取得する方法は要検証。

## 動画プレイヤー

- 再生開始
- 一時停止
- N秒進む
- N秒戻る
- Nをユーザが1〜10秒の範囲で調整できる
- スロー再生
- ユーザが1/2, 1/3, 1/4 で調整できる
- 次のシーンに進む
- 前のシーンに進む

## 動画メモ

- 動画の特定のフレームに対して140字程度のメモを残せる

ここら辺はミニマムで作成し、学習効果を見て検討する。

# 開発環境

- Flutter Web による Web アプリとして作成する
- ディレクトリの指定と一覧の取得は　[showDirectoryPicker](https://developer.mozilla.org/ja/docs/Web/API/Window/showDirectoryPicker) で行う。
- シーンと時刻は [IndexedDb](https://developer.mozilla.org/ja/docs/Web/API/IndexedDB_API) に保存する。
- 私は Android エンジニアだがサーバサイドが Go 言語の会社に転職したので、Go 言語を使いたい件は諦める。
