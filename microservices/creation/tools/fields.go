package tools

type Category struct {
	Name        string
	Parent      int
	Description string
}

var (
	Categories = map[int32]Category{
		1:  {Name: "动画", Parent: 0, Description: "动画相关的内容，包括经典动画和现代动画"},
		2:  {Name: "MAD·AMV", Parent: 1, Description: "音乐动画制作与混剪"},
		3:  {Name: "MMD·3D", Parent: 1, Description: "MikuMikuDance和3D动画"},
		4:  {Name: "同人·手书", Parent: 1, Description: "同人作品及手书内容"},
		5:  {Name: "配音", Parent: 1, Description: "动画角色或创意内容的配音"},
		6:  {Name: "动漫杂谈", Parent: 1, Description: "关于动漫的讨论和评论"},
		7:  {Name: "游戏", Parent: 0, Description: "关于电子游戏的实况、评测和讨论"},
		8:  {Name: "单机游戏", Parent: 7, Description: "单机游戏内容和玩法"},
		9:  {Name: "电子竞技", Parent: 7, Description: "电竞比赛和选手相关内容"},
		10: {Name: "手机游戏", Parent: 7, Description: "移动端游戏相关内容"},
		11: {Name: "网络游戏", Parent: 7, Description: "多人在线网络游戏"},
		12: {Name: "桌游棋牌", Parent: 7, Description: "桌面游戏和棋牌内容"},
		13: {Name: "音游", Parent: 7, Description: "音乐游戏相关内容"},
		14: {Name: "音乐", Parent: 0, Description: "与音乐相关的内容，包括表演和教程"},
		15: {Name: "原创音乐", Parent: 14, Description: "原创音乐创作"},
		16: {Name: "翻唱", Parent: 14, Description: "翻唱歌曲分享"},
		17: {Name: "演奏", Parent: 14, Description: "乐器演奏相关内容"},
		18: {Name: "影视", Parent: 0, Description: "与影视相关的内容，包括影评和电影解说"},
		19: {Name: "影视杂谈", Parent: 18, Description: "对影视剧的评论和分析"},
		20: {Name: "影视剪辑", Parent: 18, Description: "影视片段的剪辑和创意编辑"},
		21: {Name: "知识", Parent: 0, Description: "关于各种主题的教育和知识内容"},
		22: {Name: "科学科普", Parent: 21, Description: "科学知识普及"},
		23: {Name: "社科·法律·心理", Parent: 21, Description: "社会科学、法律、心理学内容"},
		24: {Name: "人文历史", Parent: 21, Description: "历史、人文学科相关内容"},
		25: {Name: "财经商业", Parent: 21, Description: "财经和商业相关知识"},
		26: {Name: "科技", Parent: 0, Description: "与科技相关的话题，包括科技产品和创新"},
		27: {Name: "数码", Parent: 26, Description: "数码产品评测与资讯"},
		28: {Name: "计算机技术", Parent: 26, Description: "计算机技术与知识分享"},
		29: {Name: "极客DIY", Parent: 26, Description: "极客精神下的创意DIY"},
		30: {Name: "美食", Parent: 0, Description: "美食的制作、食谱和烹饪技巧"},
		31: {Name: "美食制作", Parent: 30, Description: "美食的制作方法和教程"},
		32: {Name: "美食侦探", Parent: 30, Description: "美食探索与发现"},
		33: {Name: "美食测评", Parent: 30, Description: "美食的测评和体验"},
		34: {Name: "美食记录", Parent: 30, Description: "记录日常美食和生活"},
		35: {Name: "动物圈", Parent: 0, Description: "关于动物的相关内容"},
		36: {Name: "猫", Parent: 35, Description: "关于猫的内容"},
		37: {Name: "狗", Parent: 35, Description: "关于狗的内容"},
		38: {Name: "异宠", Parent: 35, Description: "关于奇特宠物的内容"},
		39: {Name: "野生动物", Parent: 35, Description: "关于野生动物的内容"},
	}
)
