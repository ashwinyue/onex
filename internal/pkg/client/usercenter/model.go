// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

package usercenter

import v1 "github.com/rosas/onex/pkg/api/usercenter/v1"

type GetSecretRequest struct {
	UserID string
	Name   string
}

type GetSecretResponse = v1.SecretReply
