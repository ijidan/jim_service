package dispatch


const (
	BusinessHeaderFlag = "jim1"
)

const (
	BusinessCmdPing        = "A001" //ping
	BusinessCmdPong        = "A002" //pong
	BusinessCmdServerClose = "A999" //close：服务端断开
	BusinessCmdClientClose = "A998" //close：客户端断开

	BusinessCmdAuthLogin   = "B001" //请求登录
	BusinessCmdAuthSuccess = "B002" //认证成功
	BusinessCmdAuthFail    = "B003" //认证失败
	BusinessCmdAuthLogout    = "B004" //退出

	BusinessCmdC2C = "B101" //C2C单聊消息
	BusinessCmdC2G = "B102" //C2G群聊消息
	BusinessCmdS2C = "B103" //S2C系统推送消息
)