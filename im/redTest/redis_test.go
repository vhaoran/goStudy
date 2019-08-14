package redTest

import (
	"fmt"
	"testing"

	"chat/Library/Inner/Nodes/yredis"
	"chat/Library/Inner/sysinit"
)

func init() {
	sysinit.InitAllCnt()
	sysinit.InitBundelConf()
}

func Test_red_up(t *testing.T) {
	script := `
         local key = KEYS[1]
         local field = KEYS[2]
         

         local amt = ARGV[1]

         redis.call("hincrbyfloat",key,field,"0.0");
         redis.call("expire",key,8640000000);

         local ye = redis.call("hget",key,field);
         ye = ye + 0

         local r = redis.call("hincrbyfloat",key,field,tostring(amt));
         return r
    `

	red := yredis.Redix
	//  prepare for feature change

	keys := []string{"ye_h", "user1"}
	s, err := red.Eval(script, keys, "50000", "msg2_b").Float64()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(s)
}

func Test_red_down(t *testing.T) {
	script := `
         local key = KEYS[1]
         local field = KEYS[2]
         local amt = ARGV[1]

         redis.call("hincrbyfloat",key,field,"0.0");

         local ye = redis.call("hget",key,field);
         ye = ye + 0

         -- tran sign 
         amt = (amt + 0)*(-1)

         if ( ye + amt < 0) 
			then
            return -1
         end

         local r = redis.call("hincrbyfloat",key,field,tostring(amt));
         return r
    `

	red := yredis.Redix
	//  prepare for feature change

	keys := []string{"test_h", "1"}
	s, err := red.Eval(script, keys, "55.55", "msg2_b").Float64()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(s)
}

func Test_red_ye(t *testing.T) {
	script := `
         local key = KEYS[1]
         local field = KEYS[2]

         redis.call("hincrbyfloat",key,field,"0.0");

         local r = redis.call("hget",key,field);
         return r
    `

	red := yredis.Redix
	//  prepare for feature change

	keys := []string{"ye_h", "user1"}
	s, err := red.Eval(script, keys).Float64()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(s)
}
func Test_red_bet(t *testing.T) {
	// remainder - amt(判断大小)
	// 加入订单队列数据
	// 加入订单金额数据

	script := `
         -- 余额hash
         local keyRemainder = KEYS[1]
         local fieldRemainder = KEYS[2]
         -- 订单队列 list
         local orderL = KEYS[3]
         -- 订单item汇总队列 list
         local orderAmtL = KEYS[4]

         

         -- 订单串
         local orderStr = ARGV[1]
         -- 单项汇总值
         local amt = ARGV[2]

         local ye = redis.call("hincrbyfloat",keyRemainder,fieldRemainder,"0.0");

         -- 方向转换,并转为数值型 
         amt = (amt + 0) * (-1)
         ye = ye + 0
         -- 判断余额是否够用
         if (ye + amt < 0)
         then
           return -1   
         end

         --  改余额
         ye = redis.call("HINCRBYFLOAT",keyRemainder,fieldRemainder,tostring(amt))
         redis.call("expire",keyRemainder, "86400000000")
         
        -- 余额转正数
          amt = (amt + 0)*(-1)

         -- 加入orderList
         redis.call("lpush",orderL,orderStr)
         redis.call("expire",orderL, "86400000000")

         -- 加入余额list
         redis.call("lpush",orderAmtL,tostring(amt))
         redis.call("expire",orderAmtL, "86400000000")

         -- 返回投注后的余额
         return ye
    `

	red := yredis.Redix
	//  prepare for feature change

	keys := []string{"ye_h", "user1", "orderL", "orderAmtL"}
	s, err := red.Eval(script, keys, "order-data", 50).Float64()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(s)
}

func Test_red_back(t *testing.T) {
	// remainder - amt(判断大小)
	// 加入订单队列数据
	// 加入订单金额数据

	script := `
         -- 余额hash
         local keyRemainder = KEYS[1]
         local fieldRemainder = KEYS[2]
         -- 订单队列 list
         local orderL = KEYS[3]
         -- 订单item汇总队列 list
         local orderAmtL = KEYS[4]


         -- 是否有投注记录
         local exists = redis.call("exists",orderL)
         exists = exists + 0
         if exists < 1
         then
            -- 无单可撤返回-100
           return -100
         end


         -- 单项值
         local amt = redis.call("lpop",orderAmtL)
         if not amt  
         then
            -- 无单可撤返回-100
           return -100
         end

         amt = amt + 0

         local ye = redis.call("hincrbyfloat",keyRemainder,fieldRemainder,tostring(amt));

         -- 返回新增后的余额
         return ye
    `

	red := yredis.Redix
	//  prepare for feature change

	keys := []string{"ye_h", "user1", "orderL", "orderAmtL"}
	s, err := red.Eval(script, keys).Float64()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(s)
}
func Test_red_backAll(t *testing.T) {
	// remainder - amt(判断大小)
	// 加入订单队列数据
	// 加入订单金额数据

	script := `
         -- 余额hash
         local keyRemainder = KEYS[1]
         local fieldRemainder = KEYS[2]
         -- 订单队列 list
         local orderL = KEYS[3]
         -- 订单item汇总队列 list
         local orderAmtL = KEYS[4]


         -- 是否有投注记录
         local exists = redis.call("exists",orderL)
         exists = exists + 0
         if exists < 1
         then
            -- 无单可撤返回-100
           return -100
         end

         local len = redis.call("llen",orderAmtL)
         if  not len
         then
            -- 无单可撤返回-100
           return -100
         end


         -- 获取所有列表
         local lst = redis.call("lrange",orderAmtL,0,-1)

         if  lst == nil  
         then
            -- 无单可撤返回-100
           return -100
         end

         -- total
        local amt = 0
        for i,v in  pairs(lst) do
         amt = amt + (v + 0)
        end

        local ye = redis.call("hincrbyfloat",keyRemainder,fieldRemainder,tostring(amt));

         redis.call("del",orderL)
         redis.call("del",orderAmtL)

         -- 返回新增后的余额
         return ye
    `

	red := yredis.Redix
	//  prepare for feature change

	keys := []string{"ye_h", "user1", "orderL", "orderAmtL"}
	s, err := red.Eval(script, keys).Float64()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(s)
}
