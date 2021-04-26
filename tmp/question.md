> TODO: iap delete tag 为什么需要 app <11-01-21,> >

```goland
	//put 更新资源，执行多次一样的结果
		router.PostWithStat("/gifts/tag/", "updateGiftsByTag", h.updateGiftsByTag)


	//连续两次解析请求，是否可以
		ids := &struct {
		IDs []int64 `json:"ids"`
	}{}
	if err := jhttp.DecodeRequest(r, ids); err != nil {
		jhttp.WriteErrorResponse(w, http.StatusBadRequest, httpError.RequestParamInvalid)
		return
	}
	appName := httputils.GetStringParamOrDefault(r, "app", "")
	if appName == "" {
		jhttp.WriteErrorResponse(w, http.StatusBadRequest, httpError.RequestParamInvalid)
		return
	}
	p := control.NewUpdatePayload()
	if err := jhttp.DecodeRequest(r, p); err != nil {
		jhttp.WriteErrorDetailResponse(w, err, "")
	}

	//校验意义，必须校验吗
	if !h.checkGiftScopeValid(p.Scope, p.RoomIDs) {
		jhttp.WriteErrorResponse(w, http.StatusBadRequest, httpError.RequestParamInvalid)
		return
	}

	//单独返回 ok 可以吗
		data := utils.StringMap{}
	data["is_update"] = ok
	jhttp.WriteResponse(w, data)

```

> 待重构

```
func GetWaitOnlineGoods(session *xorm.Session, app string, typ, country string) (model.GoodsKindRecordList, error) {
	var t *Test
	nowStr := time.Now().UTC().Format("2006-01-02 15:04:05")
	goods := make([]*model.GoodsKindRecord, 0)
	whereStr := "goods_kind.app=? and gtype=? and status=? and can_buy=? and online_time > ? and goods_country.country=?"
	t.nowStr =nowStr
	err := plan(session, whereStr, t).Find(&goods)
/*	err := session.Join("INNER", "goods_country", "goods_kind.id = goods_country.goods_id").
		Where("goods_kind.app=? and gtype=? and status=? and can_buy=? and offline_time > ? and online_time < ? and goods_country.country=?",
			app, typ, consts.GoodsKindStatusOnline, true, nowStr, nowStr, country).
		Asc("goods_country.`corder`").Find(&goods)
*/	if err != nil {
		sentry.Error(err)
		return nil, err
	}
	return goods, nil
}
type Test struct {
	app string
	typ string
	country string
	name string
	nowStr string
}
func plan(s *xorm.Session,str string,t *Test) *xorm.Session{
	session :=s.Join("INNER", "goods_country", "goods_kind.id = goods_country.goods_id").Where(
		str,t.app,t.typ,t.country,consts.GoodsKindStatusOnline,true,t.nowStr).Asc("goods_country.`corder`")
	return session
}

func where(){}

func sqlPlan(s *xorm.Session,fun func(session2 *xorm.Session,args...interface{}) *xorm.Session) *xorm.Session{
	return s.Join("INNER", "goods_country", "goods_kind.id = goods_country.goods_id").
		fun.Asc("goods_country.`corder`")
}


```
