# ゲスト側のルーティングテーブルを設定
```
ip route add default via 192.168.20.20
```

# tapデバイスに対する設定
```
sudo ip addr add 192.168.20.20/24 dev mytap0
sudo ip link set mytap0 up
```

# iptablesでVMからのパケットをnatする
```
sudo iptables -t nat -A POSTROUTING -s 192.168.20.0/24 -j MASQUERADE
```
## 消す時
```
sudo iptables -t nat -L POSTROUTING -v -n --line-numbers
sudo iptables -t nat -D POSTROUTING <行番号>
```