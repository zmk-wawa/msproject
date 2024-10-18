package model

import common "project-common"

// 返回的JSON Code 意义

const (
	LegalMobile   common.BusinessCode = 2000 // 手机号合法
	NoLegalMobile common.BusinessCode = 2001 // 手机号不合法
)
