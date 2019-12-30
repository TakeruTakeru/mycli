# mycli

## How to use
`go build [this repo]`

## Commands

### base64
Converter of base64 treats plain text.

Encode
```
./mycli base64 hogehoge
 -> aG9nZWhvZ2U=
```

Decode
```
./mycli base64 -d aG9nZWhvZ2U=
 -> hogehoge
```

