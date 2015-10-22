#! /usr/bin/env bash

# if file not found, display an error and die
[ $# -eq 0 ] && { echo "Usage: $0 moviefile"; exit 1; }

echo "working on "$1 
NAME=`echo "$1" | cut -d'.' -f1`
ffmpeg -i ${1} -vf scale=600:-1 -pix_fmt rgb24 -r 10 -f gif - | gifsicle --optimize=3 --delay=3 > $NAME.gif 

# PIDS=()

# MACHINE_LIST=("M6600M519" "M6600Master" "M6600M521" "M6600M522")

# for MACHINE in ${MACHINE_LIST[@]}; do
# 	echo $MACHINE
# 	scp jessie.mlj root@$MACHINE.local:/home/user/script/defaults/. 
# 	ssh root@$MACHINE.local -t "cd /home/user/datafile/; rm jessie.img; wget -O jessie.img $1" &
# 	PIDS+=($!)
# done

# for PID in ${PIDS[@]}; do
#     wait $PID
# done
# echo $PIDS
# echo "All sleeps done!"
# curl -X POST --data-urlencode 'payload={"text": "@daniel @ivan @justin @kodie just finished copying the image over to the flashing station", "channel": "#notifications"}' https://hooks.slack.com/services/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