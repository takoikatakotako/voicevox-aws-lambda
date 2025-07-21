# VOICEVOX AWS Lambda

VOICEVOX AWS Lambda は VOICEVOX を Lambda で動かすためのサンプルコードです。


## セットアップ

### VOICEVOXと関連ファイルのダウンロード

AWS Lambda で VOICEVOX を動作させるためには、Lambda の実行環境に対応したバイナリ（Linux ARM 64bit）が必要です。しかし、ダウンローダースクリプト download は ARM アーキテクチャの Linux 環境でしか実行できません。

以下の例では、ARM アーキテクチャの EC2インスタンスを一時的に用意し、そこで VOICEVOX Core をダウンロードしたうえで、ファイルを S3 にアップロードし、最終的にローカル環境へ移動させました。

以下のコマンドを EC2インスタンス上で実行します。

```bash
curl -sSfL https://github.com/VOICEVOX/voicevox_core/releases/download/0.16.0/download-linux-arm64 -o download
chmod +x download
./download
```

実行中に利用規約への同意が求められますので、内容を確認のうえ「同意する」選択肢を選んでください。
ダウンロードが完了すると、以下のようなファイル構成になります。


```
.
├── download
└── voicevox_core
    ├── c_api
    │   ├── include
    │   │   └── voicevox_core.h
    │   └── lib
    │       └── libvoicevox_core.so
    ├── dict
    │   └── open_jtalk_dic_utf_8-1.11
    │       └── ...
    ├── models
    │   └── vvms
    │       ├── 0.vvm
    │       ├── 1.vvm
    │       └── ...
    └── onnxruntime
        └── lib
            └── libvoicevox_onnxruntime.so.1.17.3
```

`voicevox_core` ディレクトリを S3 にアップロードします。
まず、アップロード用の S3 バケットを作成し、EC2 インスタンスにそのバケットへアップロードできる権限を付与します。
以下のコマンドを EC2 上で実行してください。

```bash
aws s3 sync voicevox_core s3://${S3_BUCKET_NAME}/voicevox_core
```

S3 にアップロードされたファイルをローカル環境にダウンロードします。
以下のコマンドをローカル環境で実行してください。

```bash
aws s3 sync s3://${S3_BUCKET_NAME}/voicevox_core voicevox_core
```

[Releases/0.16.0](https://github.com/VOICEVOX/voicevox_core/releases/tag/0.16.0)

[VOICEVOX コア ユーザーガイド](https://github.com/VOICEVOX/voicevox_core/blob/main/docs/guide/user/usage.md)


### VOICEVOXと関連ファイルの配置

`lib` ディレクトリに `libvoicevox_core.so`, `libvoicevox_onnxruntime.so.1.17.3`, `open_jtalk_dic_utf_8-1.11`, `voicevox_core.h`, `vvms` を配置します。
以下のようなファイル構成にしてください。

```
.
├── app
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── voicevox_core.go
│   └── voicevox_wrapper.go
├── Dockerfile
├── lib
│   ├── libvoicevox_core.so
│   ├── libvoicevox_onnxruntime.so.1.17.3
│   ├── open_jtalk_dic_utf_8-1.11
│   │   ├── char.bin
│   │   ├── COPYING
│   │   ├── left-id.def
│   │   ├── matrix.bin
│   │   ├── pos-id.def
│   │   ├── rewrite.def
│   │   ├── right-id.def
│   │   ├── sys.dic
│   │   └── unk.dic
│   ├── voicevox_core.h
│   └── vvms
│       ├── 0.vvm
│       ├── 1.vvm
│       └── ...
└── README.md
```


## ビルド

Dockerfile のビルドを行います。

```bash
docker build . -t voicevox-aws-lambda:latest
```


## 参考にさせていただいたリポジトリ

以下のリポジトリを参考にさせていただきました。

- [voicevoxcore.go](https://github.com/sh1ma/voicevoxcore.go)
