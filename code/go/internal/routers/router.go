package router

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/global"
	"github.com/mwqnice/oh-admin/internal/handler"
	"github.com/mwqnice/oh-admin/internal/middleware"
	"github.com/mwqnice/oh-admin/internal/widget"
	"github.com/mwqnice/oh-admin/pkg/app"
	"github.com/mwqnice/oh-admin/pkg/limiter"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Cors())
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	// 创建基于cookie的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("MsW32dQN2342434I5C43E6"))
	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("ohadmin", store))

	//加载模板
	r.HTMLRender = loadTemplates("views")
	// 设置静态资源路由
	r.Static("/static", "./static")

	r.NoRoute(HandleNotFound)
	r.NoMethod(HandleNotFound)
	deviceRecord := r.Group("deviceRecord")
	{
		deviceRecord.GET("/index", handler.DeviceRecordHandler.Index)
		deviceRecord.POST("/list", handler.DeviceRecordHandler.List)
		deviceRecord.Any("/add", handler.DeviceRecordHandler.Add)
		deviceRecord.Any("/edit", handler.DeviceRecordHandler.Edit)
		deviceRecord.POST("/delete/:id", handler.DeviceRecordHandler.Delete)
		deviceRecord.POST("/scan", handler.DeviceRecordHandler.Add)
		deviceRecord.POST("/app/list", handler.DeviceRecordHandler.AppList)

	}
	sampInfo := r.Group("samp_info")
	{
		sampInfo.Any("/index", handler.SampeHandler.Index)
		sampInfo.POST("/list", handler.SampeHandler.List)
		sampInfo.POST("/app/list", handler.SampeHandler.AppList)
		sampInfo.Any("/add", handler.SampeHandler.Add)
		sampInfo.POST("/scan", handler.SampeHandler.Scan)
		sampInfo.Any("/edit", handler.SampeHandler.Edit)
		sampInfo.POST("/delete/:id", handler.SampeHandler.Delete) //删除处方
	}
	buss := r.Group("buss")
	{
		buss.GET("/bigscreen/bigScreenData/getLatestScreenData", handler.BigScreenHandler.BigScreenData)
	}
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	//r.Use(middleware.Translations())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/check", func(c *gin.Context) {
		c.JSON(200, "OK")
		return
	})

	/* 文件上传 */
	upload := r.Group("upload")
	{
		// 上传图片
		upload.POST("/uploadImage", handler.PublicHandler.UploadImage)
	}
	trace_chinese_medicine := r.Group("trace_chinese_medicine")
	{

		trace_chinese_medicine.Any("/index", handler.Trace_chinese_medicineHandler.Index) //业务流程扫描接口
		trace_chinese_medicine.GET("/flow", handler.Trace_chinese_medicineHandler.Flow)   //业务流程扫描接口
	}
	//r.Use(middleware.Tracer())
	r.Use(middleware.AdminAuth()) //验证登录
	index := r.Group("/")
	{
		index.GET("/", handler.IndexHandler.Index)                 //首页
		index.GET("/index", handler.IndexHandler.Index)            //首页
		index.GET("/welcome", handler.IndexHandler.Welcome)        //欢迎页
		index.Any("/login", handler.PublicHandler.Login)           //登录
		index.GET("/captcha", handler.PublicHandler.Captcha)       //获取验证码
		index.GET("/logout", handler.PublicHandler.LoginOut)       //退出
		index.POST("/update_pwd", handler.PublicHandler.UpdatePwd) //修改密码
		index.POST("/check_pwd", handler.PublicHandler.CheckPwd)   //校验密码
		index.Any("/user_info", handler.AdminUserHandler.UserInfo) //获取用户信息

	}
	/* 管理员管理 */
	user := r.Group("user")
	{
		user.GET("/index", handler.AdminUserHandler.Index)          //用户详情
		user.GET("/info", handler.AdminUserHandler.UserInfo)        //用户详情
		user.POST("/list", handler.AdminUserHandler.List)           //菜单列表
		user.Any("/add", handler.AdminUserHandler.Add)              //添加用户
		user.POST("/setStatus", handler.AdminUserHandler.SetStatus) //设置状态
		user.Any("/edit", handler.AdminUserHandler.Edit)            //修改用户
		user.POST("/delete/:id", handler.AdminUserHandler.Delete)   //删除
	}
	/* 菜单管理 */
	menu := r.Group("menu")
	{
		menu.GET("/index", handler.MenuHandler.Index)        //菜单首页
		menu.POST("/list", handler.MenuHandler.List)         //菜单列表
		menu.Any("/add", handler.MenuHandler.Add)            //添加菜单
		menu.Any("/edit", handler.MenuHandler.Edit)          //修改菜单
		menu.POST("/delete/:id", handler.MenuHandler.Delete) //删除
	}
	/* 角色管理 */
	role := r.Group("role")
	{
		role.GET("/index", handler.RoleHandler.Index)                      //角色首页
		role.POST("/list", handler.RoleHandler.List)                       //角色列表
		role.Any("/add", handler.RoleHandler.Add)                          //添加角色
		role.Any("/edit", handler.RoleHandler.Edit)                        //修改角色
		role.POST("/delete/:ids", handler.RoleHandler.Delete)              //删除
		role.POST("/setStatus", handler.RoleHandler.SetStatus)             //设置状态
		role.GET("/menu_list/:role_id", handler.RoleHandler.MenuList)      //角色菜单列表
		role.POST("/menu_list/save", handler.RoleHandler.SaveRoleMenuList) //角色菜单列表保存
	}
	/* 接方管理 */
	prescription := r.Group("prescription")
	{
		prescription.GET("/index", handler.PrescriptionHandler.Index)        //处方首页
		prescription.POST("/list", handler.PrescriptionHandler.List)         //接方列表
		prescription.Any("/add", handler.PrescriptionHandler.Add)            //添加处方
		prescription.Any("/edit", handler.PrescriptionHandler.Edit)          //修改处方
		prescription.Any("/detail", handler.PrescriptionHandler.Detail)      //处方明细
		prescription.POST("/delete/:id", handler.PrescriptionHandler.Delete) //删除处方

	}
	//审核
	prescriptionAudit := r.Group("prescription_audit")
	{
		prescriptionAudit.GET("/index", handler.PrescriptionAuditHandler.Index)      //处方审核首页
		prescriptionAudit.POST("/list", handler.PrescriptionAuditHandler.List)       //审核列表
		prescriptionAudit.POST("/audit/:id", handler.PrescriptionAuditHandler.Audit) //审核处方

	}
	//调剂
	adjustment := r.Group("adjustment")
	{
		adjustment.GET("/index", handler.AdjustmentHandler.Index) //调剂首页
		adjustment.POST("/list", handler.AdjustmentHandler.List)  //调剂列表

	}
	//复核
	herb_decoction_audit := r.Group("herb_decoction_audit")
	{
		herb_decoction_audit.GET("/index", handler.HerbDecoctionAuditHandler.Index) //复核首页
		herb_decoction_audit.POST("/list", handler.HerbDecoctionAuditHandler.List)  //复核列表

	}
	//泡药
	herb_soaking := r.Group("herb_soaking")
	{
		herb_soaking.GET("/index", handler.SoakHandler.Index) //泡药首页
		herb_soaking.POST("/list", handler.SoakHandler.List)  //泡药列表
	}
	//煎药
	herbal_decoction_info := r.Group("herbal_decoction_info")
	{
		herbal_decoction_info.GET("/index", handler.DecHandler.Index) //煎药首页
		herbal_decoction_info.POST("/list", handler.DecHandler.List)  //煎药列表

	}
	//包装
	medicine_packing := r.Group("medicine_packing")
	{
		medicine_packing.GET("/index", handler.PackHandler.Index) //包装首页
		medicine_packing.POST("/list", handler.PackHandler.List)  //包装列表

	}
	//发货
	herb_delivery := r.Group("herb_delivery")
	{
		herb_delivery.GET("/index", handler.DeliveryHandler.Index) //发货首页
		herb_delivery.POST("/list", handler.DeliveryHandler.List)  //发货列表
	}
	prescriptionQuery := r.Group("prescription_query")
	{
		prescriptionQuery.Any("/index", handler.PrescriptionHandler.PrescriptionQueryIndex) //业务流程扫描接口
		prescriptionQuery.POST("/list", handler.WxHandler.PcList)                           //煎药流转信息列表
	}

	work_statistics := r.Group("work_statistics")
	{
		work_statistics.Any("/index", handler.PrescriptionHandler.WorkStatisticsIndex) //业务流程扫描接口
		work_statistics.POST("/list", handler.PrescriptionHandler.WorkStatisticsList)
	}
	hospital := r.Group("hospitaldata")
	{
		hospital.Any("/index", handler.HospitalDataHandler.Index) //业务流程扫描接口
		hospital.POST("/list", handler.HospitalDataHandler.List)
	}
	// 设备路由组
	device := r.Group("/device")
	{
		device.GET("/index", handler.DevicelDataHandler.Index)
		device.POST("/list", handler.DevicelDataHandler.List)
		device.Any("/add", handler.DevicelDataHandler.Add)
		device.Any("/edit", handler.DevicelDataHandler.Edit)
		device.POST("/delete/:id", handler.DevicelDataHandler.Delete)
		device.Any("/detail", handler.DevicelDataHandler.Detail) //处方明细
	}
	dose := r.Group("/drug_dosage")
	{
		dose.GET("/index", handler.DrugDosageHandler.Index)
		dose.POST("/list", handler.DrugDosageHandler.List)
		dose.Any("/add", handler.DrugDosageHandler.Add)
		dose.Any("/edit", handler.DrugDosageHandler.Edit)
		dose.POST("/delete/:id", handler.DrugDosageHandler.Delete)

	}

	stock_in := r.Group("stock_in")
	{
		stock_in.GET("/index", handler.StockInDataHandler.Index) //业务流程扫描接口
		stock_in.POST("/list", handler.StockInDataHandler.List)
		stock_in.Any("/add", handler.StockInDataHandler.Add) //添加
		stock_in.Any("/detail", handler.StockInDataHandler.Detail)
	}
	inventory := r.Group("inventory")
	{
		inventory.GET("/index", handler.InventoryDataHandler.Index) //业务流程扫描接口
		inventory.POST("/list", handler.InventoryDataHandler.List)
		inventory.Any("/edit", handler.InventoryDataHandler.Edit) //添加
	}
	stock_out := r.Group("stock_out")
	{
		stock_out.GET("/index", handler.StockOutDataHandler.Index) //业务流程扫描接口
		stock_out.POST("/list", handler.StockOutDataHandler.List)
		stock_out.Any("/detail", handler.StockOutDataHandler.Detail)
	}

	stock_transfer := r.Group("stock_transfer")
	{
		stock_transfer.GET("/index", handler.StockTransferDataHandler.Index) //业务流程扫描接口
		stock_transfer.POST("/list", handler.StockTransferDataHandler.List)
		stock_transfer.GET("/detail", handler.StockTransferDataHandler.Detail)
		stock_transfer.Any("/add", handler.StockTransferDataHandler.Add) //添加
	}
	stock_check := r.Group("stock_check")
	{
		stock_check.GET("/index", handler.StockCheckDataHandler.Index) //业务流程扫描接口
		stock_check.POST("/list", handler.StockCheckDataHandler.List)
		stock_check.Any("/detail", handler.StockCheckDataHandler.Detail)
		stock_check.Any("/add", handler.StockCheckDataHandler.Add) //添加
	}
	//pda
	pc := r.Group("prescription_scan")
	{
		pc.Any("/index", handler.WxHandler.Index)
		pc.POST("/list", handler.WxHandler.List)                    //煎药流转信息列表
		pc.Any("/add", handler.WxHandler.ScanBarcode)               //业务流程扫描接口
		pc.POST("/upload_chain/:id", handler.WxHandler.UploadChain) //上传数据到区块链
	}
	//app api
	app := r.Group("wx_program")
	{
		app.POST("/getMedicineProcess", handler.PrescriptionHandler.SearchView)
		app.POST("/login", handler.WxHandler.Login)
		app.POST("/info/tisane/list", handler.WxHandler.PrescriptionList)
		app.POST("/prescription/list", handler.WxHandler.PrescriptionList)
		app.POST("/prescription/drug/detail/:id", handler.WxHandler.PrescriptionDetail)
		app.POST("/prescription/scan", handler.WxHandler.AppScan)
		app.POST("/prescription/delete", handler.WxHandler.PrescriptionDelete)
	}

	return r
}

func HandleNotFound(c *gin.Context) {
	c.JSON(app.CODE_SUCCESS, &app.ResponseCommonStruct{Code: app.CODE_ROUTE_ERROR, Msg: fmt.Sprintf("路由%s不存在或不支持%s请求", c.Request.URL.String(), c.Request.Method)})
	return
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// 非模板嵌套
	adminHtmls, err := filepath.Glob(templatesDir + "/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, html := range adminHtmls {
		r.AddFromGlob(filepath.Base(html), html)
	}

	// 布局模板
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	// 嵌套的内容模板
	includes, err := filepath.Glob(templatesDir + "/includes/**/*.html")
	if err != nil {
		panic(err.Error())
	}

	// template自定义函数
	funcMap := template.FuncMap{
		"StringToLower": func(str string) string {
			return strings.ToLower(str)
		},
		"date2": func() string {
			return time.Now().Format("2006-01-02 15:04:05.00000")
		},
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
		"query":        widget.Query,
		"add":          widget.Add,
		"edit":         widget.Edit,
		"delete":       widget.Delete,
		"audit":        widget.Audit,
		"detail":       widget.Detail,
		"expand":       widget.Expand,
		"collapse":     widget.Collapse,
		"addz":         widget.Addz,
		"in":           widget.In,
		"outstock":     widget.Outstock,
		"chain_upload": widget.ChainUpload,
	}

	// 将主模板，include页面，layout子模板组合成一个完整的html页面
	for _, include := range includes {
		// 文件名称
		baseName := filepath.Base(include)
		files := []string{}
		if strings.Contains(baseName, "edit") || strings.Contains(baseName, "add") || strings.Contains(baseName, "detail") {
			files = append(files, templatesDir+"/layouts/form.html", include)
		} else {
			files = append(files, templatesDir+"/layouts/layout.html", include)
		}
		files = append(files, layouts...)
		r.AddFromFilesFuncs(baseName, funcMap, files...)
	}
	return r
}
