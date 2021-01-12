1. head 、head.next 其中一個是 null，回傳現在節點 head
2. head 、head.next 相同，找到下一個不同的數字，繼續遞迴下去
3. head 、head.next 不同，把目前的 head 記錄下來，後面的節點要繼續遞迴下去找，最後組起來 return 。