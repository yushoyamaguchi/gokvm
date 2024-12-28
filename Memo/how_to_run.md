# make initrd するために
```
mkdir -p $GOPATH/src/github.com/u-root 
cd $GOPATH/src/github.com/u-root
git clone https://github.com/u-root/u-root.git
```

# 動かすコマンド
```
sudo ./gokvm boot -k ./bzImage -i ./initrd -t mytap0
```

# 止めるコマンド
```
# 他のターミナルで
sudo pkill -f gokvm
```



