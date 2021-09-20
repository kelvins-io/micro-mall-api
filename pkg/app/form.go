package app

import (
	"errors"
	"fmt"
	"gitee.com/cristiane/micro-mall-api/internal/logging"
	"gitee.com/cristiane/micro-mall-api/vars"
	"gitee.com/kelvins-io/common/json"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"strings"
)

func BindAndValid(c *gin.Context, form interface{}) error {
	if err := c.Bind(form); err != nil {
		return err
	}
	valid := validation.Validation{}
	ok, err := valid.Valid(form)
	if err != nil {
		return err
	}
	if !ok {
		markErrors(c, valid.Errors)
		return buildFormErr(valid.Errors)
	}
	return nil
}

func buildFormErr(errs []*validation.Error) error {
	var msg strings.Builder
	for _, v := range errs {
		if v.Field != "" {
			msg.WriteString(v.Field)
		} else if v.Key != "" {
			msg.WriteString(v.Key)
		} else {
			msg.WriteString(v.Name)
		}
		msg.WriteString(" : ")
		msg.WriteString(json.MarshalToStringNoError(v.Value))
		msg.WriteString(" => ")
		msg.WriteString(v.Error())
		msg.WriteString(" should=> ")
		msg.WriteString(json.MarshalToStringNoError(v.LimitValue))
	}
	return errors.New(msg.String())
}

func markErrors(ctx *gin.Context, errors []*validation.Error) {
	buf := strings.Builder{}
	buf.WriteString(fmt.Sprintf("%v %v %v 400 ", ctx.ClientIP(), ctx.Request.Method, ctx.Request.RequestURI))
	buf.WriteString("{")
	for _, err := range errors {
		buf.WriteString(fmt.Sprintf("%v：%v ", err.Key, err.Message))
	}
	buf.WriteString(" }")
	buf.WriteString(fmt.Sprintf(" %v", ctx.Request.Header))
	if vars.AccessLogger != nil {
		vars.AccessLogger.Error(ctx, buf.String())
	} else {
		logging.Info(buf.String())
	}

	return
}
