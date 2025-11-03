package code

import (
	"code.byted.org/h_cloud/easymock"
	"code.byted.org/storage/leetcode/code/strings"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CodeTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller
}

func (suite *CodeTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
}

func (suite *CodeTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func TestCodeTestSuite(t *testing.T) {
	suite.Run(t, new(CodeTestSuite))
}

func (suite *CodeTestSuite) TestLengthOfLongestSubstring() {
	easymock.EasyMockConvey("TestLengthOfLongestSubstring test", suite.T(), func() {
		length := strings.LengthOfLongestSubstring("abcabcaaaaa")
		suite.Equal(3, length)
	})
}

func (suite *CodeTestSuite) TestCheckInclusion() {
	easymock.EasyMockConvey("TestCheckInclusion test", suite.T(), func() {
		b := strings.CheckInclusion("ab", "eidbaooo")
		suite.True(b)
	})
}

func (suite *CodeTestSuite) TestMySqrt() {
	easymock.EasyMockConvey("TestMySqrt test", suite.T(), func() {
		b := strings.LongestCommonPrefix2([]string{"ab", "a"})
		suite.Equal(b, 0)
	})
}
