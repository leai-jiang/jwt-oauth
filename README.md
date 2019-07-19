"# spata" 

# https://github.com
头像 -> setting -> Developer settings -> OAuth Apps


# Client ID
ecf4d78a2de563fbf68a

# Client Secret
01f41a42bfdd5564f4b6d7191c3d70d268f445cf

# Authorization callback URL
应用服务接口 + "?code=${code}"


##1
https://github.com/login/oauth/authorize?client_id=${ClientId}

##2
登录后重定向到 github上配置地 Authorization callback URL

##3
拿到 url 上的 code
请求 https://github.com/login/oauth/access_token?client_id=${CientId}&client_secret=${ClientSecret}&code=${code}

获取 access_token

##4
拿access_token 获取用户信息
https://api.github.com/user

``
