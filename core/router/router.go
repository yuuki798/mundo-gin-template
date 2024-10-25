package router

import (
	"github.com/gin-gonic/gin"
	"github.com/trancecho/mundo-be-template/core/middleware/response"
	"github.com/trancecho/mundo-be-template/core/router/base"
	"github.com/trancecho/mundo-be-template/core/router/protected"
)

func GenerateRouters(r *gin.Engine) *gin.Engine {

	r.Use(response.ResponseMiddleware())

	//// 注册接口
	//r.POST("/api/register", auth2.EmailRegister)
	//// 邮箱验证接口
	//r.GET("/api/verify", auth2.VerifyEmail)
	//// 登录接口
	//r.POST("/api/login", auth2.Login)
	base.Router(r)
	//r.GET("/api/preset/job/list", job.GetPresetJobList)
	//
	//// 购买套餐接口
	//r.POST("/api/buy-plan", payment.BuyPlanHandler)

	// 受保护的路由
	protected.Router(r)
	//{
	//	adminGroup := protected.Group("")
	//	adminGroup.Use(middleware3.AdminMiddleware())
	//	{
	//		adminGroup.GET("/admin/ping", utils.Ping)
	//		// 开发者使用
	//		// 预设岗位添加
	//		adminGroup.POST("/preset/job/create", job.CreateJob)
	//		adminGroup.POST("/redeem-code/create", redeem2.CreateCode)
	//		adminGroup.POST("/redeem-code/create/batch", redeem2.CreateBatchCode)
	//		adminGroup.GET("/redeem-code/list", redeem2.GetAllCodes)
	//		adminGroup.GET("/data/interview/duration", data.InterviewDurationService)
	//		adminGroup.GET("/data/user/registered", data.UserRegisteredService)
	//
	//	}
	//	// 需要激活码验证的受保护路由组
	//	activated := protected.Group("")
	//	activated.Use(redeem2.RedeemMiddleware())
	//	{
	//
	//		activated.GET("/activated/ping", utils.Ping)
	//
	//		activated.POST("/interview/register", common.UpsertPresetAndCreateInterview)
	//		activated.POST("/simulation/answer", common.CreateOrUpdateAnswer)
	//		// 简历评价接口
	//		activated.GET("/preset/resume/evaluate", common.ResumeSuggestion)
	//		// 流代理
	//		activated.Any("/:mod/:task/proxy", proxy.ProxyLLM("http://localhost:5000", db.DB))
	//		// 关闭面试
	//		activated.POST("/interview/close", interview.CloseInterview)
	//
	//		// 心跳接口
	//		activated.POST("/interview/heartbeat", heartbeat.Heartbeat)
	//	}
	//	// 注册 gRPC 服务的调用路由
	//	protected.POST("/make-payment", payment.MakePaymentHandler)
	//	// 激活码验证通过后的受保护路由
	//	// 获取面试预设接口
	//	protected.GET("/preset", common.GetPreset)
	//	// 更新预设
	//	protected.POST("/preset/upsert", common.UpsertPreset)
	//	// 上传简历
	//	protected.POST("/preset/resume/upload/pdf", common.UploadResumePDF)
	//	// 获取单个岗位信息
	//	protected.GET("/preset/job", job.GetJobByTitle)
	//	// 获取简历列表
	//	protected.GET("/resume", resume.GetResumeList)
	//	//获取问答记录
	//	protected.POST("/interview/record", common.QueryInterviewResult(db.DB))
	//	//获取单个面试信息
	//	protected.GET("/simulation/:id", common.GetSimulatedInterview)
	//	// 一个暂时废弃的接口
	//	//activated.GET("/interview/question", common2.GetQuestionIdByInterviewId)
	//	// 获取面试（历史）列表
	//	protected.GET("/interview/list", common.GetInterviewListByUid)
	//	// 删除简历
	//	protected.POST("/resume/delete", resume.DeleteResumeByID)
	//	// 设置生日
	//	protected.POST("/user/birth/set", model.SetBirth)
	//	protected.POST("/redeem-code/verify", redeem2.VerifyCode)
	//
	//	protected.GET("/profile", auth2.GetProfile)
	//}
	return r
}
