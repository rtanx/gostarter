package security_test

import (
	"testing"

	"github.com/rtanx/gostarter/internal/infrastructure/security"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	rawPwd := "myV3ryS3cReTP4assW0Rd!@#$%_+"
	hashedPwd, err := security.HashPassword(rawPwd)
	assert.NoError(t, err)

	valid := security.CheckPassword(hashedPwd, rawPwd)
	assert.True(t, valid)
}
