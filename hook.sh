#!/bin/bash
#"Got character ZDOID from Sbeve The Dim : -753911200:72"
# line=$1;# l=${line//*ZDOID from /};# l=${l// :*/};# num=${line//*:};# if (($num>30));then
#     if [ "$l" = "Sbeve The Dim" ];then
#         curl catbox:8081/rob
#     elif [ "$l" = "DAN" ];then
#         curl catbox:8081/dan
#     elif [ "$l" = "Brandon" ];then
#         curl catbox:8081/brandon
#     else
#         curl catbox:8081/default
#     fi
# fi
#read line;l=${line//*ZDOID from /};l=${l// :*/};num=${line//*:};comp=30;if [[ $num -gt $comp ]];then if [ "$l" = "Sbeve The Dim" ];then curl catbox:8081/rob;elif [ "$l" = "DAN" ];then curl catbox:8081/dan;elif [ "$l" = "Brandon" ];then curl catbox:8081/brandon;else curl catbox:8081/default;fi fi
##{ read l;line=$l; l=${l//*ZDOID from /}; l=${l// :*/};num=${line//*:};echo "\nnum:$num\nplayer:$l\nline:$line\n";comp=30;if [[ $num -gt $comp ]];then if [ "$l" = "Sbeve The Dim" ];then curl catbox:8081/rob;elif [ "$l" = "DAN" ];then curl catbox:8081/dan;elif [ "$l" = "Brandon" ];then curl catbox:8081/brandon;else curl catbox:8081/default;fi fi
curl -d "Got character ZDOID from Sbeve The Dim : -753911200:72" catbox:8081/
#  ffmpeg -i odin.mp3 -filter:a "volume=0.5" odins.mp3
#  ffmpeg -f concat -safe 0 -i files.txt -c copy odin.mp3