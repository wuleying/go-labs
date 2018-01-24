#!/usr/bin/env bash

URL="http://bj.122.gov.cn/m/mvehxh/getTfhdList"
PARAMS="page=0&glbm=110000000400&hpzl=52&type=0"

while true
do
    result=$(curl -s -XPOST -d "${PARAMS}" "${URL}" | jq ".data.list.content[0].subhd")
    echo ${result}
    echo ${result:3:1}
    echo ${result:4:1}
    echo ${result:5:1}

    if [ "${result:3:1}" == "${result:4:1}" -a "${result:4:1}" == "${result:5:1}" ]; then
        osascript -e "display notification \"豹子号段出现: ${result}\" with title \"122 Monitor\""
    fi
    sleep 60
done