#!/bin/bash

# 获取kmesh-system命名空间下的Pod名称
PODS=$(kubectl get pods -n kmesh-system --no-headers -o custom-columns=:metadata.name)

# 遍历每个Pod名称
for POD in $PODS; do
    # 为每个Pod生成日志文件名
    LOGFILE="logs/${POD}.log"

    # 创建logs目录（如果不存在）
    mkdir -p logs

    # 获取Pod的日志并保存到文件
    kubectl logs -f -n kmesh-system "$POD" > "$LOGFILE"
done

echo "日志已保存至logs目录下。"