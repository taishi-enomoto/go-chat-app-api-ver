## チャットアプリケーション

### Go言語の標準・準標準ライブラリを中心に実装を目指したウェブアプリケーションです。
ウェブ関連技術について基礎から学ぶため、またGo言語はフレームワークを用いずに実装することも多いと耳にしたため、できる限りGo言語の標準・準標準ライブラリのみを使用しての実装を目指して作成しました。
題材としては、CRUDなど基礎的な機能から、非同期通信やAPIなど様々な技術を試せる会員制のチャットアプリケーションを選択しました。
### - 技術的特徴
  - **会員登録・ログイン機能**<br>
    チャットアプリケーションに不可欠な機能のため、会員登録・ログイン機能を実装しています。パスワードは平文ではなくハッシュ化して保存しています。
  - **Cookieによるセッション管理**<br>
    Cookieによるセッション管理を行い、セッション変数にはユーザー名を保存しています。セッションIDはメモリ上で保管し、失効期限を設定しています。
  - **双方向通信(WebSocket)によるリアルタイム更新(`/mypage/chatroom`内)**<br>
    チャットアプリケーションにとってリアルタイムのやり取りは不可欠な機能のため、各チャットルーム毎に独立してWebSocket接続が行われるよう実装を行っています。
  - **APIによるチャット内容の取得(`/mypage/chatroom`内)**<br>
    チャット内容の取得は通信負荷が大きくなりがちなため、チャットの取得はチャットアプリのwebサーバーではなくAPIサーバーによって行っています。(フレームワークにgoaを使用)
  - **コンテナ内での作動(Docker)**<br>
    より実際の運用環境に近づけるため、webサーバー、DBサーバー、APIサーバーを別々のDockerコンテナ上で起動し、Docker-composeにより一つのシステムとして構築しています。(クラウドサーバー上でこのアプリ公開し続けることは費用・セキュリティ面などの問題で避けたく、ローカル環境で本アプリケーションの動作を確かめていただけるようにしたかったことも動機の一つです。)
  - **セキュリティ**<br>
    冒頭のパスワードのハッシュ化以外にも、メタ文字のエスケープによるXSS対策、プリペアドステートメントの使用によるSQLインジェクション対策、APIキーによる認証など基礎的なセキュリティ対策を行っています。また、あるチャットルームURLへ参加メンバーではないユーザーからアクセスがあっても、アクセスが拒否される設計になっています。

### - 今後の課題
  - ~~管理ページの実装~~ **(2021/7/25 管理ページの実装)**
  - ~~API認証の実装~~ **(2021/7/23 APIキーによる認証実装)**
  - 新着メッセージの通知
  - ~~ルーム一覧とチャットスペースを同一ページで表示~~ **(2021/7/24 レイアウト変更)**
  - あるユーザー名を削除した後、同名のユーザー名で登録し直すと、前のユーザーが使用していたチャットが表示されるのを防ぐ
  - ユーザー管理用とチャット管理用のデータベースの分離
  - フレームワークを用いた同様のアプリケーションの作成

### 【使用方法】
①コンテナを立ち上げ後、ログイン画面(`http://172.25.0.2/login`)へ移動します。「新規登録はこちら」から`/resistration`ページへ移動し、フォームにユーザー名・パスワード(10文字以内、パスワードは英数字)を入力し登録ボタンをクリックしてください。チャット機能の試行には最低2名の登録が必要なため、2人分のユーザーを登録してください。
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126492301-8c039aa4-55aa-42f4-93af-92d0e50b9e9e.jpg" width="700px">
</div>

<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126493332-e7309442-2dea-40a8-97df-9088a99401fd.jpg" width="700px">
</div>

②登録が完了すると、登録完了ページに遷移するので、「ログインページに戻る」をクリックしてください。
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126493427-622e52ba-fabf-47da-a2f4-e5ba13da3bbe.jpg" width="700px">
</div>

③登録した、いずれかのユーザーでログインしてください。

④マイページが表示されます。新規ルーム名(20文字以内)と、もう一人の作成済ユーザーを相手ユーザー名に入力し、「作成する」ボタンを押してください。「ルーム一覧」に作成したルームが表示されます。
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126493536-013a0c7a-52ce-44ce-8ee1-55392acf6046.jpg" width="700px">
</div>

⑤作成したルームをクリックすると、チャットページへ遷移します。  「新規投稿」フォームにチャット内容(255文字以内)を入力し、「メッセージを投稿」ボタンを押すと投稿が行われます。投稿内容はページの更新を経ずに反映されます。
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126494044-eb99c71d-69aa-4831-baa6-1cadbb6f3309.jpg" width="700px">
</div>

<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126505084-2b3625ff-41ea-48ae-9905-b7da1ca7e207.gif" width="700px">
</div>

⑥「このルームを削除する」ボタンを押すと、そのチャットルームが削除されます。(クリック一回で削除されます。)
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126494276-a693c618-8d81-415b-8520-4d7e435e5122.jpg" width="700px">
</div>

⑦マイページにて「退会はこちら」リンクをクリックすると、退会ページに遷移します。  ログイン中のユーザーのユーザーID・パスワードを入力して「送信」ボタンを押すと、ユーザーが削除されます。
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126592571-8f59b69b-f902-49c7-907e-ba182489b32a.jpg" width="700px">
</div>

<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126498692-b551b827-a576-4952-9084-a095e246969b.jpg" width="700px">
</div>
