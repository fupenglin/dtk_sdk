package dtk_sdk

type DtkOpenPlatform struct {
	appSecret string //
	appKey    string
}

func NewOpenPlatform(appKey, appSecret string) *DtkOpenPlatform {
	return &DtkOpenPlatform{
		appSecret: appSecret,
		appKey:    appKey,
	}
}

// 查询大淘客所有的一级分类和二级分类
func (this *DtkOpenPlatform) GetSuperCategory() (string, error) {
	url := "https://openapi.dataoke.com/api/category/get-super-category"
	args := NewParam()
	args.Set("appKey", this.appKey)
	args.Set("version", "v1.1.0")
	return DtkClient{}.Get(url, args, this.appSecret)
}

// 根据二级分类获取商品列表
func (this *DtkOpenPlatform) GetGoodsList(subcid, pageid string) (string, error) {
	url := "https://openapi.dataoke.com/api/goods/get-goods-list"
	args := NewParam()
	args.Set("appKey", this.appKey)
	args.Set("version", "v1.2.0")
	args.Set("subcid", subcid)
	args.Set("pageId", pageid)
	args.Set("pageSize", "10")
	return DtkClient{}.Get(url, args, this.appSecret)
}

// 通过淘宝id查询单个商品的详细信息
func (this *DtkOpenPlatform) GetGoodsDetails(goodsId string) (string, error) {
	url := "https://openapi.dataoke.com/api/goods/get-goods-details"
	args := NewParam()
	args.Set("version", "v1.2.0")
	args.Set("appKey", this.appKey)
	args.Set("goodsId", goodsId)
	return DtkClient{}.Get(url, args, this.appSecret)
}

// 转链接口
func (this *DtkOpenPlatform) GetPrivilegeLink(goodsId string) (string, error) {
	url := "https://openapi.dataoke.com/api/tb-service/get-privilege-link"
	args := NewParam()
	args.Set("version", "v1.1.1")
	args.Set("appKey", this.appKey)
	args.Set("goodsId", goodsId)
	return DtkClient{}.Get(url, args, this.appSecret)
}

// 超级搜索
func (this *DtkOpenPlatform) List_super_goods(keyWords, pageId string) (string, error) {
	url := "https://openapi.dataoke.com/api/goods/list-super-goods"
	args := NewParam()
	args.Set("version", "v1.2.1")
	args.Set("appKey", this.appKey)
	args.Set("keyWords", keyWords)
	args.Set("pageId", pageId)
	args.Set("pageSize", "20")
	args.Set("type", "0")
	return DtkClient{}.Get(url, args, this.appSecret)
}
