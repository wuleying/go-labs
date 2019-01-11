package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(ctx)

	for index := 0; index < 3; index++ {
		indexTmp := index
		group.Go(func() error {
			fmt.Printf("indexTmp = %d\n", indexTmp)
			if indexTmp == 0 {
				fmt.Println("indexTmp == 0 end ")
			} else if indexTmp == 1 {
				fmt.Println("indexTmp == 1 start ")
				// 这里一般都是某个协程发生异常之后，调用cancel()
				// 样别的协程就可以通过errCtx获取到err信息，以便决定是否需要取消后续操作
				cancel()
				fmt.Println("indexTmp == 1 had err ")
			} else if indexTmp == 2 {
				fmt.Println("indexTmp == 2 begin ")
				time.Sleep(1 * time.Second)
				// 检查 其他协程已经发生错误，如果已经发生异常，则不再执行下面的代码
				err := CheckGoroutineErr(errCtx)
				if err != nil {
					return err
				}
				fmt.Println("indexTmp == 2 end ")
			}
			return nil
		})
	}

	err := group.Wait()
	if err == nil {
		fmt.Println("都完成了")
	} else {
		fmt.Printf("get error:%v\n", err)
	}
}

// CheckGoroutineErr 校验是否有协程已发生错误
func CheckGoroutineErr(errContext context.Context) error {
	select {
	case <-errContext.Done():
		return errContext.Err()
	default:
		return nil
	}
}
