### tmux 基本操作

> vi 模式

1. setw -g mode-keys vi

2. 进入 vi mode，`prefix + [` ，退出，`q`

### 会话

- prefix
  |key|map|
  |----|----|
  |`prefix + d` |分离当前会话|
  |`prefix + s` |列出所有会话|
  |`prefix + $` |重命名当前会话|

### 窗口

| key              | map                   |
| ---------------- | --------------------- |
| `prefix + c`     | 创建一个新窗口        |
| `prefix + p`     | 切换到上一个窗口      |
| `prefix + n`     | 切换到下一个窗口      |
| `prefix + <num>` | 切换到执行 num 的窗口 |
| `prefix + w`     | 从列表中选择窗口      |
| `prefix + ,`     | 窗口重命名            |

### 窗格

| key          | map                          |
| ------------ | ---------------------------- |
| `prefix + %` | 划分左右窗格                 |
| `prefix + "` | 划分上下窗格                 |
| `prefix + ;` | 光标切换到上一个窗格         |
| `prefix + o` | 光标切换到下一个窗格         |
| `prefix + x` | 关闭当前窗格                 |
| `prefix + !` | 将当前窗格拆分独立           |
| `prefix + z` | 当前窗格全屏，再按一次回原样 |
| `prefix + q` | 显示窗格编号                 |
