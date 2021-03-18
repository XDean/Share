# Git 个人建议

1. 信任git
    - git不会弄丢你的文件
    - 仔细阅读git给你的信息，里面包括了重要的当前信息以及可能的下一步信息
1. 尽快提交
    - 细粒度的提交可以帮助你或协作者review代码和追溯定位历史
    - 提交之前文件不受git保护，修改一旦丢失无法找回
1. 不要rebase
    - 历史记录是版本控制的核心，不要为了所谓的"clean commit tree"而掩盖历史
    - rebase只在必要时使用
1. 必须写commit message，必须在第一行就描述清楚commit的改动
1. 保持最新状态
    - 随时随地fetch
    - 定时和（且只和）source branch合并
1. 不要提交（包括但不限于）
    - 第三方依赖（如node_modules）
    - 可生成的文件
    - 二进制文件（如exe）
    - 临时测试文件
1. 不要和他人共享工作分支
1. 遵循workflow