#!/usr/bin/env bash

symbols="ETH#14#<7000 EOS#29#<50"
cookie="xxx"

while true
do
    message=""
    for item in $symbols;
    do
        symbol=`echo "$item" | awk -F '#' '{print $1}'`
        symbolNum=`echo "$item" | awk -F '#' '{print $2}'`
        price=0

        price=$(curl -sS "https://www.coinw.com/real/market2.html?symbol=${symbolNum}" \
        -H 'Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7' \
        -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36' \
        -H "Cookie: ${cookie}" | jq '.buy1[0]')

        condition=`echo $item | awk -F '#' '{print $3}'`
        cmpSymbol=${condition:0:1}
        cmpNum=${condition:1}


        if [ "$cmpSymbol"x = ">"x ]; then
            c=$(echo "$price > $cmpNum" | bc)
            if [ "$c"x == "1"x ]; then
                message="$message$symbol  "\￥"$price \t(高于: $cmpNum)\n"
           fi
        elif [ "$cmpSymbol"x = "<"x ]; then
            c=$(echo "$price < $cmpNum" | bc)
            if [ "$c"x == "1"x ]; then
                message="$message$symbol  "\￥"$price \t(低于: $cmpNum)\n"
            fi
        fi
        # message="$message$symbol    "￥"$price\n"
        echo "$symbol"    \￥ "$price"
        # echo "$price $cmpSymbol $cmpNum"
    done

    if [ "$message" != "" ]; then
        osascript -e "display notification \"$message\" with title \"Coinw Monitor\""
    fi

    sleep 3
done