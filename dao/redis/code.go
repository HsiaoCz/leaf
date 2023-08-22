package redis

// 这里放phone验证码相关的redis方法
// 生成的验证码存到redis里面，用户传过来的验证码和redis里面的做比较
// 相同则验证成功
// 当然这里要设置过期时间，一般为60s
