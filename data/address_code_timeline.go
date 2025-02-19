package data

// 行政区划代码（地址码）更新时间线
// 中华人民共和国民政部权威数据
// 注1：台湾省、香港特别行政区和澳门特别行政区暂缺地市和区县信息
// 注2：每月发布的区划变更表是根据区划变更地的统计人员在统计信息系统更新后的情况所绘制，与区划变更文件发布的时间有一定的延迟性，但在每年的最后一次发布变更情况后与区划全年变更文件保持一致。
// Data Source: http://www.mca.gov.cn/article/sj/xzqh/
var AddressCodeTimeline = map[int][]map[string]string{
	110000: {
		{
			"address":    "北京市",
			"start_year": "",
			"end_year":   "",
		},
	},
	110101: {
		{
			"address":    "东城区",
			"start_year": "",
			"end_year":   "",
		},
	},
	110102: {
		{
			"address":    "西城区",
			"start_year": "",
			"end_year":   "",
		},
	},
	110103: {
		{
			"address":    "崇文区",
			"start_year": "",
			"end_year":   "2009",
		},
	},
	110104: {
		{
			"address":    "宣武区",
			"start_year": "",
			"end_year":   "2009",
		},
	},
	110105: {
		{
			"address":    "朝阳区",
			"start_year": "",
			"end_year":   "",
		},
	},
	110106: {
		{
			"address":    "丰台区",
			"start_year": "",
			"end_year":   "",
		},
	},
	110107: {
		{
			"address":    "石景山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	110108: {
		{
			"address":    "海淀区",
			"start_year": "",
			"end_year":   "",
		},
	},
	110109: {
		{
			"address":    "门头沟区",
			"start_year": "",
			"end_year":   "",
		},
	},
	110110: {
		{
			"address":    "燕山区",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	110111: {
		{
			"address":    "房山区",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	110112: {
		{
			"address":    "通州区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	110113: {
		{
			"address":    "顺义区",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	110114: {
		{
			"address":    "昌平区",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	110115: {
		{
			"address":    "大兴区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	110116: {
		{
			"address":    "怀柔区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	110117: {
		{
			"address":    "平谷区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	110118: {
		{
			"address":    "密云区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	110119: {
		{
			"address":    "延庆区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	110201: {
		{
			"address":    "昌平县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	110202: {
		{
			"address":    "顺义县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	110203: {
		{
			"address":    "通县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	110204: {
		{
			"address":    "大兴县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	110205: {
		{
			"address":    "房山县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	110206: {
		{
			"address":    "平谷县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	110207: {
		{
			"address":    "怀柔县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	110208: {
		{
			"address":    "密云县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	110209: {
		{
			"address":    "延庆县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	110221: {
		{
			"address":    "昌平县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	110222: {
		{
			"address":    "顺义县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	110223: {
		{
			"address":    "通县",
			"start_year": "1982",
			"end_year":   "1996",
		},
	},
	110224: {
		{
			"address":    "大兴县",
			"start_year": "1982",
			"end_year":   "2000",
		},
	},
	110225: {
		{
			"address":    "房山县",
			"start_year": "1982",
			"end_year":   "1985",
		},
	},
	110226: {
		{
			"address":    "平谷县",
			"start_year": "1982",
			"end_year":   "2000",
		},
	},
	110227: {
		{
			"address":    "怀柔县",
			"start_year": "1982",
			"end_year":   "2000",
		},
	},
	110228: {
		{
			"address":    "密云县",
			"start_year": "1982",
			"end_year":   "2014",
		},
	},
	110229: {
		{
			"address":    "延庆县",
			"start_year": "1982",
			"end_year":   "2014",
		},
	},
	120000: {
		{
			"address":    "天津市",
			"start_year": "",
			"end_year":   "",
		},
	},
	120101: {
		{
			"address":    "和平区",
			"start_year": "",
			"end_year":   "",
		},
	},
	120102: {
		{
			"address":    "河东区",
			"start_year": "",
			"end_year":   "",
		},
	},
	120103: {
		{
			"address":    "河西区",
			"start_year": "",
			"end_year":   "",
		},
	},
	120104: {
		{
			"address":    "南开区",
			"start_year": "",
			"end_year":   "",
		},
	},
	120105: {
		{
			"address":    "河北区",
			"start_year": "",
			"end_year":   "",
		},
	},
	120106: {
		{
			"address":    "红桥区",
			"start_year": "",
			"end_year":   "",
		},
	},
	120107: {
		{
			"address":    "塘沽区",
			"start_year": "",
			"end_year":   "2008",
		},
	},
	120108: {
		{
			"address":    "汉沽区",
			"start_year": "",
			"end_year":   "2008",
		},
	},
	120109: {
		{
			"address":    "大港区",
			"start_year": "",
			"end_year":   "2008",
		},
	},
	120110: {
		{
			"address":    "东郊区",
			"start_year": "",
			"end_year":   "1991",
		},
		{
			"address":    "东丽区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	120111: {
		{
			"address":    "西郊区",
			"start_year": "",
			"end_year":   "1991",
		},
		{
			"address":    "西青区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	120112: {
		{
			"address":    "南郊区",
			"start_year": "",
			"end_year":   "1991",
		},
		{
			"address":    "津南区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	120113: {
		{
			"address":    "北郊区",
			"start_year": "",
			"end_year":   "1991",
		},
		{
			"address":    "北辰区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	120114: {
		{
			"address":    "武清区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	120115: {
		{
			"address":    "宝坻区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	120116: {
		{
			"address":    "滨海新区",
			"start_year": "2009",
			"end_year":   "",
		},
	},
	120117: {
		{
			"address":    "宁河区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	120118: {
		{
			"address":    "静海区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	120119: {
		{
			"address":    "蓟州区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	120201: {
		{
			"address":    "宁河县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	120202: {
		{
			"address":    "武清县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	120203: {
		{
			"address":    "静海县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	120204: {
		{
			"address":    "宝坻县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	120205: {
		{
			"address":    "蓟县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	120221: {
		{
			"address":    "宁河县",
			"start_year": "1982",
			"end_year":   "2014",
		},
	},
	120222: {
		{
			"address":    "武清县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	120223: {
		{
			"address":    "静海县",
			"start_year": "1982",
			"end_year":   "2014",
		},
	},
	120224: {
		{
			"address":    "宝坻县",
			"start_year": "1982",
			"end_year":   "2000",
		},
	},
	120225: {
		{
			"address":    "蓟县",
			"start_year": "1982",
			"end_year":   "2015",
		},
	},
	130000: {
		{
			"address":    "河北省",
			"start_year": "",
			"end_year":   "",
		},
	},
	130100: {
		{
			"address":    "石家庄市",
			"start_year": "",
			"end_year":   "",
		},
	},
	130102: {
		{
			"address":    "长安区",
			"start_year": "",
			"end_year":   "",
		},
	},
	130103: {
		{
			"address":    "桥东区",
			"start_year": "",
			"end_year":   "2014",
		},
	},
	130104: {
		{
			"address":    "桥西区",
			"start_year": "",
			"end_year":   "",
		},
	},
	130105: {
		{
			"address":    "新华区",
			"start_year": "",
			"end_year":   "",
		},
	},
	130106: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	130107: {
		{
			"address":    "井陉矿区",
			"start_year": "",
			"end_year":   "",
		},
	},
	130108: {
		{
			"address":    "裕华区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	130109: {
		{
			"address":    "藁城区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	130110: {
		{
			"address":    "鹿泉区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	130111: {
		{
			"address":    "栾城区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	130121: {
		{
			"address":    "井陉县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130122: {
		{
			"address":    "获鹿县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	130123: {
		{
			"address":    "栾城县",
			"start_year": "1986",
			"end_year":   "1992",
		},
		{
			"address":    "正定县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130124: {
		{
			"address":    "正定县",
			"start_year": "1986",
			"end_year":   "1992",
		},
		{
			"address":    "栾城县",
			"start_year": "1993",
			"end_year":   "2013",
		},
	},
	130125: {
		{
			"address":    "行唐县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130126: {
		{
			"address":    "灵寿县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130127: {
		{
			"address":    "高邑县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130128: {
		{
			"address":    "深泽县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130129: {
		{
			"address":    "赞皇县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130130: {
		{
			"address":    "无极县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130131: {
		{
			"address":    "平山县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130132: {
		{
			"address":    "元氏县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130133: {
		{
			"address":    "赵县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130181: {
		{
			"address":    "辛集市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	130182: {
		{
			"address":    "藁城市",
			"start_year": "1995",
			"end_year":   "2013",
		},
	},
	130183: {
		{
			"address":    "晋州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	130184: {
		{
			"address":    "新乐市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	130185: {
		{
			"address":    "鹿泉市",
			"start_year": "1995",
			"end_year":   "2013",
		},
	},
	130200: {
		{
			"address":    "唐山市",
			"start_year": "",
			"end_year":   "",
		},
	},
	130202: {
		{
			"address":    "路南区",
			"start_year": "",
			"end_year":   "",
		},
	},
	130203: {
		{
			"address":    "路北区",
			"start_year": "",
			"end_year":   "",
		},
	},
	130204: {
		{
			"address":    "东矿区",
			"start_year": "",
			"end_year":   "1994",
		},
		{
			"address":    "古冶区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	130205: {
		{
			"address":    "开平区",
			"start_year": "",
			"end_year":   "",
		},
	},
	130206: {
		{
			"address":    "新区",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	130207: {
		{
			"address":    "丰南区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	130208: {
		{
			"address":    "丰润区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	130209: {
		{
			"address":    "曹妃甸区",
			"start_year": "2012",
			"end_year":   "",
		},
	},
	130221: {
		{
			"address":    "丰润县",
			"start_year": "1983",
			"end_year":   "2001",
		},
	},
	130222: {
		{
			"address":    "丰南县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	130223: {
		{
			"address":    "滦县",
			"start_year": "1983",
			"end_year":   "2017",
		},
	},
	130224: {
		{
			"address":    "滦南县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130225: {
		{
			"address":    "乐亭县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130226: {
		{
			"address":    "玉田县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "迁安县",
			"start_year": "1984",
			"end_year":   "1995",
		},
	},
	130227: {
		{
			"address":    "唐海县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "迁西县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	130228: {
		{
			"address":    "迁安县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "遵化县",
			"start_year": "1984",
			"end_year":   "1991",
		},
	},
	130229: {
		{
			"address":    "迁西县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "玉田县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	130230: {
		{
			"address":    "遵化县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "唐海县",
			"start_year": "1984",
			"end_year":   "2011",
		},
	},
	130281: {
		{
			"address":    "遵化市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	130282: {
		{
			"address":    "丰南市",
			"start_year": "1995",
			"end_year":   "2001",
		},
	},
	130283: {
		{
			"address":    "迁安市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	130284: {
		{
			"address":    "滦州市",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	130300: {
		{
			"address":    "秦皇岛市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130302: {
		{
			"address":    "海港区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130303: {
		{
			"address":    "山海关区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130304: {
		{
			"address":    "北戴河区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130305: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	130306: {
		{
			"address":    "抚宁区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	130321: {
		{
			"address":    "青龙县",
			"start_year": "1983",
			"end_year":   "1985",
		},
		{
			"address":    "青龙满族自治县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	130322: {
		{
			"address":    "昌黎县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130323: {
		{
			"address":    "抚宁县",
			"start_year": "1983",
			"end_year":   "2014",
		},
	},
	130324: {
		{
			"address":    "卢龙县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130400: {
		{
			"address":    "邯郸市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130402: {
		{
			"address":    "邯山区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130403: {
		{
			"address":    "丛台区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130404: {
		{
			"address":    "复兴区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130405: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	130406: {
		{
			"address":    "峰峰矿区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130407: {
		{
			"address":    "肥乡区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	130408: {
		{
			"address":    "永年区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	130421: {
		{
			"address":    "邯郸县",
			"start_year": "1983",
			"end_year":   "2015",
		},
	},
	130422: {
		{
			"address":    "武安县",
			"start_year": "1986",
			"end_year":   "1987",
		},
	},
	130423: {
		{
			"address":    "临漳县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130424: {
		{
			"address":    "成安县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130425: {
		{
			"address":    "大名县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130426: {
		{
			"address":    "涉县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130427: {
		{
			"address":    "磁县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130428: {
		{
			"address":    "肥乡县",
			"start_year": "1993",
			"end_year":   "2015",
		},
	},
	130429: {
		{
			"address":    "永年县",
			"start_year": "1993",
			"end_year":   "2015",
		},
	},
	130430: {
		{
			"address":    "丘县",
			"start_year": "1993",
			"end_year":   "1995",
		},
		{
			"address":    "邱县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	130431: {
		{
			"address":    "鸡泽县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130432: {
		{
			"address":    "广平县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130433: {
		{
			"address":    "馆陶县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130434: {
		{
			"address":    "魏县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130435: {
		{
			"address":    "曲周县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130481: {
		{
			"address":    "武安市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	130500: {
		{
			"address":    "邢台市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130502: {
		{
			"address":    "桥东区",
			"start_year": "1983",
			"end_year":   "2019",
		},
		{
			"address":    "襄都区",
			"start_year": "2020",
			"end_year":   "",
		},
	},
	130503: {
		{
			"address":    "桥西区",
			"start_year": "1983",
			"end_year":   "2019",
		},
		{
			"address":    "信都区",
			"start_year": "2020",
			"end_year":   "",
		},
	},
	130504: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	130505: {
		{
			"address":    "任泽区",
			"start_year": "2020",
			"end_year":   "",
		},
	},
	130506: {
		{
			"address":    "南和区",
			"start_year": "2020",
			"end_year":   "",
		},
	},
	130521: {
		{
			"address":    "邢台县",
			"start_year": "1986",
			"end_year":   "2019",
		},
	},
	130522: {
		{
			"address":    "临城县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130523: {
		{
			"address":    "内丘县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130524: {
		{
			"address":    "柏乡县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130525: {
		{
			"address":    "隆尧县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130526: {
		{
			"address":    "任县",
			"start_year": "1993",
			"end_year":   "2019",
		},
	},
	130527: {
		{
			"address":    "南和县",
			"start_year": "1993",
			"end_year":   "2019",
		},
	},
	130528: {
		{
			"address":    "宁晋县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130529: {
		{
			"address":    "巨鹿县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130530: {
		{
			"address":    "新河县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130531: {
		{
			"address":    "广宗县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130532: {
		{
			"address":    "平乡县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130533: {
		{
			"address":    "威县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130534: {
		{
			"address":    "清河县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130535: {
		{
			"address":    "临西县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130581: {
		{
			"address":    "南宫市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	130582: {
		{
			"address":    "沙河市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	130600: {
		{
			"address":    "保定市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130602: {
		{
			"address":    "新市区",
			"start_year": "1983",
			"end_year":   "2014",
		},
		{
			"address":    "竞秀区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	130603: {
		{
			"address":    "北市区",
			"start_year": "1983",
			"end_year":   "2014",
		},
	},
	130604: {
		{
			"address":    "南市区",
			"start_year": "1983",
			"end_year":   "2014",
		},
	},
	130605: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "1986",
		},
	},
	130606: {
		{
			"address":    "莲池区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	130607: {
		{
			"address":    "满城区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	130608: {
		{
			"address":    "清苑区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	130609: {
		{
			"address":    "徐水区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	130621: {
		{
			"address":    "满城县",
			"start_year": "1983",
			"end_year":   "2014",
		},
	},
	130622: {
		{
			"address":    "清苑县",
			"start_year": "1986",
			"end_year":   "2014",
		},
	},
	130623: {
		{
			"address":    "涞水县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	130624: {
		{
			"address":    "阜平县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	130625: {
		{
			"address":    "徐水县",
			"start_year": "1994",
			"end_year":   "2014",
		},
	},
	130626: {
		{
			"address":    "定兴县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	130627: {
		{
			"address":    "唐县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	130628: {
		{
			"address":    "高阳县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	130629: {
		{
			"address":    "容城县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	130630: {
		{
			"address":    "涞源县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	130631: {
		{
			"address":    "望都县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	130632: {
		{
			"address":    "安新县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	130633: {
		{
			"address":    "易县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	130634: {
		{
			"address":    "曲阳县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	130635: {
		{
			"address":    "蠡县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	130636: {
		{
			"address":    "顺平县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	130637: {
		{
			"address":    "博野县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	130638: {
		{
			"address":    "雄县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	130681: {
		{
			"address":    "涿州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	130682: {
		{
			"address":    "定州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	130683: {
		{
			"address":    "安国市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	130684: {
		{
			"address":    "高碑店市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	130700: {
		{
			"address":    "张家口市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130702: {
		{
			"address":    "桥东区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130703: {
		{
			"address":    "桥西区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130704: {
		{
			"address":    "茶坊区",
			"start_year": "1983",
			"end_year":   "1988",
		},
	},
	130705: {
		{
			"address":    "宣化区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130706: {
		{
			"address":    "下花园区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130707: {
		{
			"address":    "庞家堡区",
			"start_year": "1983",
			"end_year":   "1988",
		},
	},
	130708: {
		{
			"address":    "万全区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	130709: {
		{
			"address":    "崇礼区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	130721: {
		{
			"address":    "宣化县",
			"start_year": "1983",
			"end_year":   "2015",
		},
	},
	130722: {
		{
			"address":    "张北县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130723: {
		{
			"address":    "康保县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130724: {
		{
			"address":    "沽源县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130725: {
		{
			"address":    "尚义县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130726: {
		{
			"address":    "蔚县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130727: {
		{
			"address":    "阳原县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130728: {
		{
			"address":    "怀安县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130729: {
		{
			"address":    "万全县",
			"start_year": "1993",
			"end_year":   "2015",
		},
	},
	130730: {
		{
			"address":    "怀来县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130731: {
		{
			"address":    "涿鹿县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130732: {
		{
			"address":    "赤城县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130733: {
		{
			"address":    "崇礼县",
			"start_year": "1993",
			"end_year":   "2015",
		},
	},
	130800: {
		{
			"address":    "承德市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130802: {
		{
			"address":    "双桥区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130803: {
		{
			"address":    "双滦区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130804: {
		{
			"address":    "鹰手营子矿区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130821: {
		{
			"address":    "承德县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130822: {
		{
			"address":    "兴隆县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130823: {
		{
			"address":    "平泉县",
			"start_year": "1993",
			"end_year":   "2016",
		},
	},
	130824: {
		{
			"address":    "滦平县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130825: {
		{
			"address":    "隆化县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130826: {
		{
			"address":    "丰宁满族自治县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130827: {
		{
			"address":    "宽城满族自治县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130828: {
		{
			"address":    "围场满族蒙古族自治县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130881: {
		{
			"address":    "平泉市",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	130900: {
		{
			"address":    "沧州市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130902: {
		{
			"address":    "新华区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130903: {
		{
			"address":    "运河区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130904: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "1996",
		},
	},
	130921: {
		{
			"address":    "沧县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	130922: {
		{
			"address":    "青县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	130923: {
		{
			"address":    "东光县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130924: {
		{
			"address":    "海兴县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130925: {
		{
			"address":    "盐山县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130926: {
		{
			"address":    "肃宁县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130927: {
		{
			"address":    "南皮县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130928: {
		{
			"address":    "吴桥县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130929: {
		{
			"address":    "献县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130930: {
		{
			"address":    "孟村回族自治县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	130981: {
		{
			"address":    "泊头市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	130982: {
		{
			"address":    "任丘市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	130983: {
		{
			"address":    "黄骅市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	130984: {
		{
			"address":    "河间市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	131000: {
		{
			"address":    "廊坊市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	131002: {
		{
			"address":    "安次区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	131003: {
		{
			"address":    "广阳区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	131021: {
		{
			"address":    "三河县",
			"start_year": "1988",
			"end_year":   "1992",
		},
	},
	131022: {
		{
			"address":    "固安县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	131023: {
		{
			"address":    "永清县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	131024: {
		{
			"address":    "香河县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	131025: {
		{
			"address":    "大城县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	131026: {
		{
			"address":    "文安县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	131027: {
		{
			"address":    "霸县",
			"start_year": "1988",
			"end_year":   "1989",
		},
	},
	131028: {
		{
			"address":    "大厂回族自治县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	131081: {
		{
			"address":    "霸州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	131082: {
		{
			"address":    "三河市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	131100: {
		{
			"address":    "衡水市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	131102: {
		{
			"address":    "桃城区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	131103: {
		{
			"address":    "冀州区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	131121: {
		{
			"address":    "枣强县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	131122: {
		{
			"address":    "武邑县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	131123: {
		{
			"address":    "武强县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	131124: {
		{
			"address":    "饶阳县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	131125: {
		{
			"address":    "安平县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	131126: {
		{
			"address":    "故城县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	131127: {
		{
			"address":    "景县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	131128: {
		{
			"address":    "阜城县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	131181: {
		{
			"address":    "冀州市",
			"start_year": "1996",
			"end_year":   "2015",
		},
	},
	131182: {
		{
			"address":    "深州市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	132100: {
		{
			"address":    "邯郸地区",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132101: {
		{
			"address":    "邯郸市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132102: {
		{
			"address":    "邯山区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132103: {
		{
			"address":    "丛台区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132104: {
		{
			"address":    "复兴区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132105: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132106: {
		{
			"address":    "峰峰矿区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132121: {
		{
			"address":    "大名县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132122: {
		{
			"address":    "魏县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132123: {
		{
			"address":    "曲周县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132124: {
		{
			"address":    "丘县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132125: {
		{
			"address":    "鸡泽县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132126: {
		{
			"address":    "肥乡县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132127: {
		{
			"address":    "广平县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132128: {
		{
			"address":    "成安县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132129: {
		{
			"address":    "临漳县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132130: {
		{
			"address":    "磁县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132131: {
		{
			"address":    "武安县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	132132: {
		{
			"address":    "涉县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132133: {
		{
			"address":    "永年县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132134: {
		{
			"address":    "邯郸县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132135: {
		{
			"address":    "馆陶县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132200: {
		{
			"address":    "邢台地区",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132201: {
		{
			"address":    "邢台市",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "南宫市",
			"start_year": "1986",
			"end_year":   "1992",
		},
	},
	132202: {
		{
			"address":    "桥东区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "沙河市",
			"start_year": "1987",
			"end_year":   "1992",
		},
	},
	132203: {
		{
			"address":    "桥西区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132204: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132221: {
		{
			"address":    "邢台县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	132222: {
		{
			"address":    "沙河县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	132223: {
		{
			"address":    "临城县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132224: {
		{
			"address":    "内丘县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132225: {
		{
			"address":    "柏乡县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132226: {
		{
			"address":    "隆尧县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132227: {
		{
			"address":    "任县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132228: {
		{
			"address":    "南和县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132229: {
		{
			"address":    "宁晋县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132230: {
		{
			"address":    "南宫县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	132231: {
		{
			"address":    "巨鹿县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132232: {
		{
			"address":    "新河县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132233: {
		{
			"address":    "广宗县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132234: {
		{
			"address":    "平乡县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132235: {
		{
			"address":    "威县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132236: {
		{
			"address":    "清河县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132237: {
		{
			"address":    "临西县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132300: {
		{
			"address":    "石家庄地区",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132301: {
		{
			"address":    "辛集市",
			"start_year": "1986",
			"end_year":   "1992",
		},
	},
	132302: {
		{
			"address":    "藁城市",
			"start_year": "1989",
			"end_year":   "1992",
		},
	},
	132303: {
		{
			"address":    "晋州市",
			"start_year": "1991",
			"end_year":   "1992",
		},
	},
	132304: {
		{
			"address":    "新乐市",
			"start_year": "1992",
			"end_year":   "1992",
		},
	},
	132321: {
		{
			"address":    "束鹿县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	132322: {
		{
			"address":    "晋县",
			"start_year": "",
			"end_year":   "1990",
		},
	},
	132323: {
		{
			"address":    "深泽县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132324: {
		{
			"address":    "无极县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132325: {
		{
			"address":    "藁城县",
			"start_year": "",
			"end_year":   "1988",
		},
	},
	132326: {
		{
			"address":    "赵县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132327: {
		{
			"address":    "栾城县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	132328: {
		{
			"address":    "正定县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	132329: {
		{
			"address":    "新乐县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	132330: {
		{
			"address":    "高邑县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132331: {
		{
			"address":    "元氏县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132332: {
		{
			"address":    "赞皇县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132333: {
		{
			"address":    "井陉县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132334: {
		{
			"address":    "获鹿县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132335: {
		{
			"address":    "平山县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132336: {
		{
			"address":    "灵寿县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132337: {
		{
			"address":    "行唐县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132400: {
		{
			"address":    "保定地区",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132401: {
		{
			"address":    "保定市",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "定州市",
			"start_year": "1986",
			"end_year":   "1993",
		},
	},
	132402: {
		{
			"address":    "新市区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "涿州市",
			"start_year": "1986",
			"end_year":   "1993",
		},
	},
	132403: {
		{
			"address":    "北市区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "安国市",
			"start_year": "1991",
			"end_year":   "1993",
		},
	},
	132404: {
		{
			"address":    "南市区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "高碑店市",
			"start_year": "1993",
			"end_year":   "1993",
		},
	},
	132405: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132421: {
		{
			"address":    "易县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132422: {
		{
			"address":    "满城县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132423: {
		{
			"address":    "徐水县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132424: {
		{
			"address":    "涞源县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132425: {
		{
			"address":    "定兴县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132426: {
		{
			"address":    "完县",
			"start_year": "",
			"end_year":   "1992",
		},
		{
			"address":    "顺平县",
			"start_year": "1993",
			"end_year":   "1993",
		},
	},
	132427: {
		{
			"address":    "唐县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132428: {
		{
			"address":    "望都县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132429: {
		{
			"address":    "涞水县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132430: {
		{
			"address":    "涿县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	132431: {
		{
			"address":    "清苑县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	132432: {
		{
			"address":    "高阳县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132433: {
		{
			"address":    "安新县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132434: {
		{
			"address":    "雄县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132435: {
		{
			"address":    "容城县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132436: {
		{
			"address":    "新城县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132437: {
		{
			"address":    "曲阳县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132438: {
		{
			"address":    "阜平县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132439: {
		{
			"address":    "定县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	132440: {
		{
			"address":    "安国县",
			"start_year": "",
			"end_year":   "1990",
		},
	},
	132441: {
		{
			"address":    "博野县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132442: {
		{
			"address":    "蠡县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	132500: {
		{
			"address":    "张家口地区",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132501: {
		{
			"address":    "张家口市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132502: {
		{
			"address":    "桥东区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132503: {
		{
			"address":    "桥西区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132504: {
		{
			"address":    "茶坊区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132505: {
		{
			"address":    "宣化区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132506: {
		{
			"address":    "下花园区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132507: {
		{
			"address":    "庞家堡区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132521: {
		{
			"address":    "张北县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132522: {
		{
			"address":    "康保县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132523: {
		{
			"address":    "沽源县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132524: {
		{
			"address":    "尚义县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132525: {
		{
			"address":    "蔚县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132526: {
		{
			"address":    "阳原县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132527: {
		{
			"address":    "怀安县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132528: {
		{
			"address":    "万全县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132529: {
		{
			"address":    "怀来县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132530: {
		{
			"address":    "涿鹿县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132531: {
		{
			"address":    "宣化县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132532: {
		{
			"address":    "赤城县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132533: {
		{
			"address":    "崇礼县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132600: {
		{
			"address":    "承德地区",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132601: {
		{
			"address":    "承德市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132602: {
		{
			"address":    "双桥区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132603: {
		{
			"address":    "双滦区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132604: {
		{
			"address":    "鹰手营子矿区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132621: {
		{
			"address":    "青龙县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132622: {
		{
			"address":    "宽城县",
			"start_year": "",
			"end_year":   "1988",
		},
		{
			"address":    "宽城满族自治县",
			"start_year": "1989",
			"end_year":   "1992",
		},
	},
	132623: {
		{
			"address":    "兴隆县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132624: {
		{
			"address":    "平泉县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132625: {
		{
			"address":    "承德县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132626: {
		{
			"address":    "滦平县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132627: {
		{
			"address":    "丰宁县",
			"start_year": "",
			"end_year":   "1985",
		},
		{
			"address":    "丰宁满族自治县",
			"start_year": "1986",
			"end_year":   "1992",
		},
	},
	132628: {
		{
			"address":    "隆化县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132629: {
		{
			"address":    "围场县",
			"start_year": "",
			"end_year":   "1988",
		},
		{
			"address":    "围场满族蒙古族自治县",
			"start_year": "1989",
			"end_year":   "1992",
		},
	},
	132700: {
		{
			"address":    "唐山地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132701: {
		{
			"address":    "秦皇岛市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132702: {
		{
			"address":    "海港区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132703: {
		{
			"address":    "山海关区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132704: {
		{
			"address":    "北戴河区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132705: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132721: {
		{
			"address":    "丰润县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132722: {
		{
			"address":    "丰南县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132723: {
		{
			"address":    "滦县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132724: {
		{
			"address":    "滦南县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132725: {
		{
			"address":    "乐亭县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132726: {
		{
			"address":    "昌黎县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132727: {
		{
			"address":    "抚宁县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132728: {
		{
			"address":    "卢龙县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132729: {
		{
			"address":    "迁安县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132730: {
		{
			"address":    "迁西县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132731: {
		{
			"address":    "遵化县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132732: {
		{
			"address":    "玉田县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132733: {
		{
			"address":    "唐海县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	132800: {
		{
			"address":    "廊坊地区",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	132801: {
		{
			"address":    "廊坊市",
			"start_year": "1981",
			"end_year":   "1987",
		},
	},
	132821: {
		{
			"address":    "三河县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	132822: {
		{
			"address":    "大厂回族自治县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	132823: {
		{
			"address":    "香河县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	132824: {
		{
			"address":    "安次县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132825: {
		{
			"address":    "永清县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	132826: {
		{
			"address":    "固安县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	132827: {
		{
			"address":    "霸县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	132828: {
		{
			"address":    "文安县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	132829: {
		{
			"address":    "大城县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	132900: {
		{
			"address":    "沧州地区",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132901: {
		{
			"address":    "沧州市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132902: {
		{
			"address":    "新华区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "泊头市",
			"start_year": "1984",
			"end_year":   "1992",
		},
	},
	132903: {
		{
			"address":    "运河区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "任丘市",
			"start_year": "1986",
			"end_year":   "1992",
		},
	},
	132904: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "黄骅市",
			"start_year": "1989",
			"end_year":   "1992",
		},
	},
	132905: {
		{
			"address":    "泊头市",
			"start_year": "1982",
			"end_year":   "1983",
		},
		{
			"address":    "河间市",
			"start_year": "1990",
			"end_year":   "1992",
		},
	},
	132921: {
		{
			"address":    "沧县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132922: {
		{
			"address":    "河间县",
			"start_year": "",
			"end_year":   "1989",
		},
	},
	132923: {
		{
			"address":    "肃宁县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132924: {
		{
			"address":    "献县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132925: {
		{
			"address":    "交河县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	132926: {
		{
			"address":    "吴桥县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132927: {
		{
			"address":    "东光县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132928: {
		{
			"address":    "南皮县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132929: {
		{
			"address":    "盐山县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132930: {
		{
			"address":    "黄骅县",
			"start_year": "",
			"end_year":   "1988",
		},
	},
	132931: {
		{
			"address":    "孟村回族自治县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	132932: {
		{
			"address":    "青县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	132933: {
		{
			"address":    "任丘县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	132934: {
		{
			"address":    "海兴县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	133000: {
		{
			"address":    "衡水地区",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	133001: {
		{
			"address":    "衡水市",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	133002: {
		{
			"address":    "冀州市",
			"start_year": "1993",
			"end_year":   "1995",
		},
	},
	133003: {
		{
			"address":    "深州市",
			"start_year": "1994",
			"end_year":   "1995",
		},
	},
	133021: {
		{
			"address":    "衡水县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	133022: {
		{
			"address":    "冀县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	133023: {
		{
			"address":    "枣强县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	133024: {
		{
			"address":    "武邑县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	133025: {
		{
			"address":    "深县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	133026: {
		{
			"address":    "武强县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	133027: {
		{
			"address":    "饶阳县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	133028: {
		{
			"address":    "安平县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	133029: {
		{
			"address":    "故城县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	133030: {
		{
			"address":    "景县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	133031: {
		{
			"address":    "阜城县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	139001: {
		{
			"address":    "武安市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	139002: {
		{
			"address":    "霸州市",
			"start_year": "1990",
			"end_year":   "1994",
		},
	},
	139003: {
		{
			"address":    "遵化市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	139004: {
		{
			"address":    "辛集市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	139005: {
		{
			"address":    "藁城市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	139006: {
		{
			"address":    "晋州市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	139007: {
		{
			"address":    "新乐市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	139008: {
		{
			"address":    "泊头市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	139009: {
		{
			"address":    "任丘市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	139010: {
		{
			"address":    "黄骅市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	139011: {
		{
			"address":    "河间市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	139012: {
		{
			"address":    "三河县",
			"start_year": "1993",
			"end_year":   "1993",
		},
		{
			"address":    "三河市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	139013: {
		{
			"address":    "南宫市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	139014: {
		{
			"address":    "沙河市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	139015: {
		{
			"address":    "定州市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	139016: {
		{
			"address":    "涿州市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	139017: {
		{
			"address":    "安国市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	139018: {
		{
			"address":    "高碑店市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	139019: {
		{
			"address":    "鹿泉市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	139020: {
		{
			"address":    "丰南市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	140000: {
		{
			"address":    "山西省",
			"start_year": "",
			"end_year":   "",
		},
	},
	140100: {
		{
			"address":    "太原市",
			"start_year": "",
			"end_year":   "",
		},
	},
	140102: {
		{
			"address":    "南城区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	140103: {
		{
			"address":    "北城区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	140104: {
		{
			"address":    "河西区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	140105: {
		{
			"address":    "小店区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	140106: {
		{
			"address":    "迎泽区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	140107: {
		{
			"address":    "杏花岭区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	140108: {
		{
			"address":    "尖草坪区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	140109: {
		{
			"address":    "万柏林区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	140110: {
		{
			"address":    "晋源区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	140111: {
		{
			"address":    "古交工矿区",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	140112: {
		{
			"address":    "南郊区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	140113: {
		{
			"address":    "北郊区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	140121: {
		{
			"address":    "清徐县",
			"start_year": "",
			"end_year":   "",
		},
	},
	140122: {
		{
			"address":    "阳曲县",
			"start_year": "",
			"end_year":   "",
		},
	},
	140123: {
		{
			"address":    "娄烦县",
			"start_year": "",
			"end_year":   "",
		},
	},
	140181: {
		{
			"address":    "古交市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	140200: {
		{
			"address":    "大同市",
			"start_year": "",
			"end_year":   "",
		},
	},
	140202: {
		{
			"address":    "城区",
			"start_year": "",
			"end_year":   "2017",
		},
	},
	140203: {
		{
			"address":    "矿区",
			"start_year": "",
			"end_year":   "2017",
		},
	},
	140211: {
		{
			"address":    "南郊区",
			"start_year": "",
			"end_year":   "2017",
		},
	},
	140212: {
		{
			"address":    "新荣区",
			"start_year": "",
			"end_year":   "",
		},
	},
	140213: {
		{
			"address":    "平城区",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	140214: {
		{
			"address":    "云冈区",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	140215: {
		{
			"address":    "云州区",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	140221: {
		{
			"address":    "阳高县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	140222: {
		{
			"address":    "天镇县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	140223: {
		{
			"address":    "广灵县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	140224: {
		{
			"address":    "灵丘县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	140225: {
		{
			"address":    "浑源县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	140226: {
		{
			"address":    "左云县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	140227: {
		{
			"address":    "大同县",
			"start_year": "1993",
			"end_year":   "2017",
		},
	},
	140300: {
		{
			"address":    "阳泉市",
			"start_year": "",
			"end_year":   "",
		},
	},
	140302: {
		{
			"address":    "城区",
			"start_year": "",
			"end_year":   "",
		},
	},
	140303: {
		{
			"address":    "矿区",
			"start_year": "",
			"end_year":   "",
		},
	},
	140311: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "",
		},
	},
	140321: {
		{
			"address":    "平定县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	140322: {
		{
			"address":    "盂县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	140400: {
		{
			"address":    "长治市",
			"start_year": "",
			"end_year":   "",
		},
	},
	140402: {
		{
			"address":    "城区",
			"start_year": "",
			"end_year":   "2017",
		},
	},
	140403: {
		{
			"address":    "潞州区",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	140404: {
		{
			"address":    "上党区",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	140405: {
		{
			"address":    "屯留区",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	140406: {
		{
			"address":    "潞城区",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	140411: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "2017",
		},
	},
	140421: {
		{
			"address":    "长治县",
			"start_year": "1983",
			"end_year":   "2017",
		},
	},
	140422: {
		{
			"address":    "潞城县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	140423: {
		{
			"address":    "襄垣县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	140424: {
		{
			"address":    "屯留县",
			"start_year": "1985",
			"end_year":   "2017",
		},
	},
	140425: {
		{
			"address":    "平顺县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	140426: {
		{
			"address":    "黎城县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	140427: {
		{
			"address":    "壶关县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	140428: {
		{
			"address":    "长子县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	140429: {
		{
			"address":    "武乡县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	140430: {
		{
			"address":    "沁县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	140431: {
		{
			"address":    "沁源县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	140481: {
		{
			"address":    "潞城市",
			"start_year": "1995",
			"end_year":   "2017",
		},
	},
	140500: {
		{
			"address":    "晋城市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	140502: {
		{
			"address":    "城区",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	140503: {
		{
			"address":    "郊区",
			"start_year": "1985",
			"end_year":   "1995",
		},
	},
	140521: {
		{
			"address":    "沁水县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	140522: {
		{
			"address":    "阳城县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	140523: {
		{
			"address":    "高平县",
			"start_year": "1985",
			"end_year":   "1992",
		},
	},
	140524: {
		{
			"address":    "陵川县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	140525: {
		{
			"address":    "泽州县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	140581: {
		{
			"address":    "高平市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	140600: {
		{
			"address":    "朔州市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	140602: {
		{
			"address":    "朔城区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	140603: {
		{
			"address":    "平鲁区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	140621: {
		{
			"address":    "山阴县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	140622: {
		{
			"address":    "应县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	140623: {
		{
			"address":    "右玉县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	140624: {
		{
			"address":    "怀仁县",
			"start_year": "1993",
			"end_year":   "2017",
		},
	},
	140681: {
		{
			"address":    "怀仁市",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	140700: {
		{
			"address":    "晋中市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	140702: {
		{
			"address":    "榆次区",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	140703: {
		{
			"address":    "太谷区",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	140721: {
		{
			"address":    "榆社县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	140722: {
		{
			"address":    "左权县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	140723: {
		{
			"address":    "和顺县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	140724: {
		{
			"address":    "昔阳县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	140725: {
		{
			"address":    "寿阳县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	140726: {
		{
			"address":    "太谷县",
			"start_year": "1999",
			"end_year":   "2018",
		},
	},
	140727: {
		{
			"address":    "祁县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	140728: {
		{
			"address":    "平遥县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	140729: {
		{
			"address":    "灵石县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	140781: {
		{
			"address":    "介休市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	140800: {
		{
			"address":    "运城市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140802: {
		{
			"address":    "盐湖区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140821: {
		{
			"address":    "临猗县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140822: {
		{
			"address":    "万荣县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140823: {
		{
			"address":    "闻喜县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140824: {
		{
			"address":    "稷山县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140825: {
		{
			"address":    "新绛县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140826: {
		{
			"address":    "绛县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140827: {
		{
			"address":    "垣曲县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140828: {
		{
			"address":    "夏县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140829: {
		{
			"address":    "平陆县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140830: {
		{
			"address":    "芮城县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140881: {
		{
			"address":    "永济市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140882: {
		{
			"address":    "河津市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140900: {
		{
			"address":    "忻州市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140902: {
		{
			"address":    "忻府区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140921: {
		{
			"address":    "定襄县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140922: {
		{
			"address":    "五台县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140923: {
		{
			"address":    "代县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140924: {
		{
			"address":    "繁峙县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140925: {
		{
			"address":    "宁武县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140926: {
		{
			"address":    "静乐县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140927: {
		{
			"address":    "神池县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140928: {
		{
			"address":    "五寨县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140929: {
		{
			"address":    "岢岚县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140930: {
		{
			"address":    "河曲县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140931: {
		{
			"address":    "保德县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140932: {
		{
			"address":    "偏关县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	140981: {
		{
			"address":    "原平市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141000: {
		{
			"address":    "临汾市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141002: {
		{
			"address":    "尧都区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141021: {
		{
			"address":    "曲沃县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141022: {
		{
			"address":    "翼城县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141023: {
		{
			"address":    "襄汾县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141024: {
		{
			"address":    "洪洞县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141025: {
		{
			"address":    "古县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141026: {
		{
			"address":    "安泽县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141027: {
		{
			"address":    "浮山县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141028: {
		{
			"address":    "吉县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141029: {
		{
			"address":    "乡宁县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141030: {
		{
			"address":    "大宁县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141031: {
		{
			"address":    "隰县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141032: {
		{
			"address":    "永和县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141033: {
		{
			"address":    "蒲县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141034: {
		{
			"address":    "汾西县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141081: {
		{
			"address":    "侯马市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141082: {
		{
			"address":    "霍州市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	141100: {
		{
			"address":    "吕梁市",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	141102: {
		{
			"address":    "离石区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	141121: {
		{
			"address":    "文水县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	141122: {
		{
			"address":    "交城县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	141123: {
		{
			"address":    "兴县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	141124: {
		{
			"address":    "临县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	141125: {
		{
			"address":    "柳林县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	141126: {
		{
			"address":    "石楼县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	141127: {
		{
			"address":    "岚县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	141128: {
		{
			"address":    "方山县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	141129: {
		{
			"address":    "中阳县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	141130: {
		{
			"address":    "交口县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	141181: {
		{
			"address":    "孝义市",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	141182: {
		{
			"address":    "汾阳市",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	142100: {
		{
			"address":    "雁北地区",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	142121: {
		{
			"address":    "大同县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "阳高县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	142122: {
		{
			"address":    "阳高县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "天镇县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	142123: {
		{
			"address":    "天镇县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "广灵县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	142124: {
		{
			"address":    "广灵县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "灵丘县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	142125: {
		{
			"address":    "灵丘县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "浑源县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	142126: {
		{
			"address":    "浑源县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "应县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	142127: {
		{
			"address":    "怀仁县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "山阴县",
			"start_year": "1982",
			"end_year":   "1987",
		},
	},
	142128: {
		{
			"address":    "应县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "朔县",
			"start_year": "1982",
			"end_year":   "1987",
		},
	},
	142129: {
		{
			"address":    "山阴县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "平鲁县",
			"start_year": "1982",
			"end_year":   "1987",
		},
	},
	142130: {
		{
			"address":    "朔县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "左云县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	142131: {
		{
			"address":    "平鲁县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "右玉县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	142132: {
		{
			"address":    "左云县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "大同县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	142133: {
		{
			"address":    "右玉县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "怀仁县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	142200: {
		{
			"address":    "忻县地区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "忻州地区",
			"start_year": "1983",
			"end_year":   "1999",
		},
	},
	142201: {
		{
			"address":    "忻州市",
			"start_year": "1983",
			"end_year":   "1999",
		},
	},
	142202: {
		{
			"address":    "原平市",
			"start_year": "1993",
			"end_year":   "1999",
		},
	},
	142221: {
		{
			"address":    "忻县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	142222: {
		{
			"address":    "原平县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "定襄县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142223: {
		{
			"address":    "代县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "五台县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142224: {
		{
			"address":    "繁峙县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "原平县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	142225: {
		{
			"address":    "五台县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "代县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142226: {
		{
			"address":    "定襄县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "繁峙县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142227: {
		{
			"address":    "静乐县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宁武县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142228: {
		{
			"address":    "岢岚县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "静乐县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142229: {
		{
			"address":    "保德县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "神池县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142230: {
		{
			"address":    "五寨县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	142231: {
		{
			"address":    "河曲县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "岢岚县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142232: {
		{
			"address":    "偏关县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "河曲县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142233: {
		{
			"address":    "神池县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "保德县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142234: {
		{
			"address":    "宁武县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "偏关县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142300: {
		{
			"address":    "晋中地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "吕梁地区",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	142301: {
		{
			"address":    "榆次市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "孝义市",
			"start_year": "1992",
			"end_year":   "2002",
		},
	},
	142302: {
		{
			"address":    "离石市",
			"start_year": "1996",
			"end_year":   "2002",
		},
	},
	142303: {
		{
			"address":    "汾阳市",
			"start_year": "1996",
			"end_year":   "2002",
		},
	},
	142321: {
		{
			"address":    "榆次县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "汾阳县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	142322: {
		{
			"address":    "寿阳县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "文水县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	142323: {
		{
			"address":    "盂县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "交城县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	142324: {
		{
			"address":    "平定县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "孝义县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	142325: {
		{
			"address":    "昔阳县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "兴县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	142326: {
		{
			"address":    "和顺县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "临县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	142327: {
		{
			"address":    "左权县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "柳林县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	142328: {
		{
			"address":    "榆社县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "石楼县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	142329: {
		{
			"address":    "太谷县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "岚县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	142330: {
		{
			"address":    "祁县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "方山县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	142331: {
		{
			"address":    "平遥县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "离石县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	142332: {
		{
			"address":    "介休县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "中阳县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	142333: {
		{
			"address":    "灵石县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "交口县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	142400: {
		{
			"address":    "吕梁地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "晋中地区",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	142401: {
		{
			"address":    "榆次市",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	142402: {
		{
			"address":    "介休市",
			"start_year": "1992",
			"end_year":   "1998",
		},
	},
	142421: {
		{
			"address":    "离石县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "榆社县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	142422: {
		{
			"address":    "孝义县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "左权县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	142423: {
		{
			"address":    "兴县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "和顺县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	142424: {
		{
			"address":    "交口县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "昔阳县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	142425: {
		{
			"address":    "方山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "平定县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	142426: {
		{
			"address":    "石楼县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "盂县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	142427: {
		{
			"address":    "岚县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "寿阳县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	142428: {
		{
			"address":    "中阳县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "榆次县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	142429: {
		{
			"address":    "交城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "太谷县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	142430: {
		{
			"address":    "临县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "祁县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	142431: {
		{
			"address":    "文水县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "平遥县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	142432: {
		{
			"address":    "柳林县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "介休县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	142433: {
		{
			"address":    "汾阳县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "灵石县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	142500: {
		{
			"address":    "晋东南地区",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	142501: {
		{
			"address":    "晋城市",
			"start_year": "1983",
			"end_year":   "1984",
		},
	},
	142521: {
		{
			"address":    "长治县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	142522: {
		{
			"address":    "潞城县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	142523: {
		{
			"address":    "襄垣县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "屯留县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	142524: {
		{
			"address":    "武乡县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "长子县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	142525: {
		{
			"address":    "黎城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "沁水县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	142526: {
		{
			"address":    "平顺县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "阳城县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	142527: {
		{
			"address":    "壶关县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "晋城县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	142528: {
		{
			"address":    "陵川县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "高平县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	142529: {
		{
			"address":    "高平县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "陵川县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	142530: {
		{
			"address":    "晋城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "壶关县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	142531: {
		{
			"address":    "阳城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "平顺县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	142532: {
		{
			"address":    "沁水县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "黎城县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	142533: {
		{
			"address":    "长子县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "武乡县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	142534: {
		{
			"address":    "屯留县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "襄垣县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	142535: {
		{
			"address":    "沁源县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "沁县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	142536: {
		{
			"address":    "沁县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "沁源县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	142600: {
		{
			"address":    "临汾地区",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	142601: {
		{
			"address":    "临汾市",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	142602: {
		{
			"address":    "侯马市",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	142603: {
		{
			"address":    "霍州市",
			"start_year": "1989",
			"end_year":   "1999",
		},
	},
	142621: {
		{
			"address":    "临汾县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "曲沃县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142622: {
		{
			"address":    "隰县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "翼城县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142623: {
		{
			"address":    "汾西县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "襄汾县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142624: {
		{
			"address":    "永和县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "临汾县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	142625: {
		{
			"address":    "安泽县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "洪洞县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142626: {
		{
			"address":    "洪洞县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "霍县",
			"start_year": "1982",
			"end_year":   "1988",
		},
	},
	142627: {
		{
			"address":    "古县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	142628: {
		{
			"address":    "霍县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "安泽县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142629: {
		{
			"address":    "翼城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "浮山县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142630: {
		{
			"address":    "浮山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "吉县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142631: {
		{
			"address":    "曲沃县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "乡宁县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142632: {
		{
			"address":    "襄汾县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "蒲县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142633: {
		{
			"address":    "吉县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "大宁县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142634: {
		{
			"address":    "乡宁县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "永和县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142635: {
		{
			"address":    "大宁县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "隰县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142636: {
		{
			"address":    "蒲县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "汾西县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142700: {
		{
			"address":    "运城地区",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	142701: {
		{
			"address":    "运城市",
			"start_year": "1983",
			"end_year":   "1999",
		},
	},
	142702: {
		{
			"address":    "永济市",
			"start_year": "1994",
			"end_year":   "1999",
		},
	},
	142703: {
		{
			"address":    "河津市",
			"start_year": "1994",
			"end_year":   "1999",
		},
	},
	142721: {
		{
			"address":    "运城县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	142722: {
		{
			"address":    "夏县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "永济县",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	142723: {
		{
			"address":    "闻喜县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "芮城县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142724: {
		{
			"address":    "绛县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "临猗县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142725: {
		{
			"address":    "垣曲县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "万荣县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142726: {
		{
			"address":    "平陆县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "新绛县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142727: {
		{
			"address":    "芮城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "稷山县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142728: {
		{
			"address":    "永济县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "河津县",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	142729: {
		{
			"address":    "临猗县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "闻喜县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142730: {
		{
			"address":    "万荣县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "夏县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142731: {
		{
			"address":    "新绛县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "绛县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142732: {
		{
			"address":    "稷山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "平陆县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	142733: {
		{
			"address":    "河津县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "垣曲县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	149001: {
		{
			"address":    "古交市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	149002: {
		{
			"address":    "高平市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	149003: {
		{
			"address":    "潞城市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	150000: {
		{
			"address":    "内蒙古自治区",
			"start_year": "",
			"end_year":   "",
		},
	},
	150100: {
		{
			"address":    "呼和浩特市",
			"start_year": "",
			"end_year":   "",
		},
	},
	150102: {
		{
			"address":    "新城区",
			"start_year": "",
			"end_year":   "",
		},
	},
	150103: {
		{
			"address":    "回民区",
			"start_year": "",
			"end_year":   "",
		},
	},
	150104: {
		{
			"address":    "玉泉区",
			"start_year": "",
			"end_year":   "",
		},
	},
	150105: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1999",
		},
		{
			"address":    "赛罕区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	150121: {
		{
			"address":    "土默特左旗",
			"start_year": "",
			"end_year":   "",
		},
	},
	150122: {
		{
			"address":    "托克托县",
			"start_year": "",
			"end_year":   "",
		},
	},
	150123: {
		{
			"address":    "和林格尔县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	150124: {
		{
			"address":    "清水河县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	150125: {
		{
			"address":    "武川县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	150200: {
		{
			"address":    "包头市",
			"start_year": "",
			"end_year":   "",
		},
	},
	150202: {
		{
			"address":    "东河区",
			"start_year": "",
			"end_year":   "",
		},
	},
	150203: {
		{
			"address":    "昆都仑区",
			"start_year": "",
			"end_year":   "",
		},
	},
	150204: {
		{
			"address":    "青山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	150205: {
		{
			"address":    "石拐矿区",
			"start_year": "",
			"end_year":   "1998",
		},
		{
			"address":    "石拐区",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	150206: {
		{
			"address":    "白云矿区",
			"start_year": "",
			"end_year":   "2006",
		},
		{
			"address":    "白云鄂博矿区",
			"start_year": "2007",
			"end_year":   "",
		},
	},
	150207: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1998",
		},
		{
			"address":    "九原区",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	150221: {
		{
			"address":    "土默特右旗",
			"start_year": "",
			"end_year":   "",
		},
	},
	150222: {
		{
			"address":    "固阳县",
			"start_year": "",
			"end_year":   "",
		},
	},
	150223: {
		{
			"address":    "达尔罕茂明安联合旗",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	150300: {
		{
			"address":    "乌海市",
			"start_year": "",
			"end_year":   "",
		},
	},
	150302: {
		{
			"address":    "海勃湾区",
			"start_year": "",
			"end_year":   "",
		},
	},
	150303: {
		{
			"address":    "海南区",
			"start_year": "",
			"end_year":   "",
		},
	},
	150304: {
		{
			"address":    "乌达区",
			"start_year": "",
			"end_year":   "",
		},
	},
	150400: {
		{
			"address":    "赤峰市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	150402: {
		{
			"address":    "红山区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	150403: {
		{
			"address":    "元宝山区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	150404: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "1992",
		},
		{
			"address":    "松山区",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	150421: {
		{
			"address":    "阿鲁科尔沁旗",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	150422: {
		{
			"address":    "巴林左旗",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	150423: {
		{
			"address":    "巴林右旗",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	150424: {
		{
			"address":    "林西县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	150425: {
		{
			"address":    "克什克腾旗",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	150426: {
		{
			"address":    "翁牛特旗",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	150428: {
		{
			"address":    "喀喇沁旗",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	150429: {
		{
			"address":    "宁城县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	150430: {
		{
			"address":    "敖汉旗",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	150500: {
		{
			"address":    "通辽市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	150502: {
		{
			"address":    "科尔沁区",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	150521: {
		{
			"address":    "科尔沁左翼中旗",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	150522: {
		{
			"address":    "科尔沁左翼后旗",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	150523: {
		{
			"address":    "开鲁县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	150524: {
		{
			"address":    "库伦旗",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	150525: {
		{
			"address":    "奈曼旗",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	150526: {
		{
			"address":    "扎鲁特旗",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	150581: {
		{
			"address":    "霍林郭勒市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	150600: {
		{
			"address":    "鄂尔多斯市",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150602: {
		{
			"address":    "东胜区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150603: {
		{
			"address":    "康巴什区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	150621: {
		{
			"address":    "达拉特旗",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150622: {
		{
			"address":    "准格尔旗",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150623: {
		{
			"address":    "鄂托克前旗",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150624: {
		{
			"address":    "鄂托克旗",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150625: {
		{
			"address":    "杭锦旗",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150626: {
		{
			"address":    "乌审旗",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150627: {
		{
			"address":    "伊金霍洛旗",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150700: {
		{
			"address":    "呼伦贝尔市",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150702: {
		{
			"address":    "海拉尔区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150703: {
		{
			"address":    "扎赉诺尔区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	150721: {
		{
			"address":    "阿荣旗",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150722: {
		{
			"address":    "莫力达瓦达斡尔族自治旗",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150723: {
		{
			"address":    "鄂伦春自治旗",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150724: {
		{
			"address":    "鄂温克族自治旗",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150725: {
		{
			"address":    "陈巴尔虎旗",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150726: {
		{
			"address":    "新巴尔虎左旗",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150727: {
		{
			"address":    "新巴尔虎右旗",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150781: {
		{
			"address":    "满洲里市",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150782: {
		{
			"address":    "牙克石市",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150783: {
		{
			"address":    "扎兰屯市",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150784: {
		{
			"address":    "额尔古纳市",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150785: {
		{
			"address":    "根河市",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	150800: {
		{
			"address":    "巴彦淖尔市",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150802: {
		{
			"address":    "临河区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150821: {
		{
			"address":    "五原县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150822: {
		{
			"address":    "磴口县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150823: {
		{
			"address":    "乌拉特前旗",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150824: {
		{
			"address":    "乌拉特中旗",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150825: {
		{
			"address":    "乌拉特后旗",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150826: {
		{
			"address":    "杭锦后旗",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150900: {
		{
			"address":    "乌兰察布市",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150902: {
		{
			"address":    "集宁区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150921: {
		{
			"address":    "卓资县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150922: {
		{
			"address":    "化德县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150923: {
		{
			"address":    "商都县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150924: {
		{
			"address":    "兴和县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150925: {
		{
			"address":    "凉城县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150926: {
		{
			"address":    "察哈尔右翼前旗",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150927: {
		{
			"address":    "察哈尔右翼中旗",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150928: {
		{
			"address":    "察哈尔右翼后旗",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150929: {
		{
			"address":    "四子王旗",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	150981: {
		{
			"address":    "丰镇市",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	152100: {
		{
			"address":    "呼伦贝尔盟",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	152101: {
		{
			"address":    "海拉尔市",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	152102: {
		{
			"address":    "满洲里市",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	152103: {
		{
			"address":    "扎兰屯市",
			"start_year": "1983",
			"end_year":   "2000",
		},
	},
	152104: {
		{
			"address":    "牙克石市",
			"start_year": "1983",
			"end_year":   "2000",
		},
	},
	152105: {
		{
			"address":    "根河市",
			"start_year": "1994",
			"end_year":   "2000",
		},
	},
	152106: {
		{
			"address":    "额尔古纳市",
			"start_year": "1994",
			"end_year":   "2000",
		},
	},
	152121: {
		{
			"address":    "布特哈旗",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	152122: {
		{
			"address":    "阿荣旗",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	152123: {
		{
			"address":    "莫力达瓦达斡尔族自治旗",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	152124: {
		{
			"address":    "喜桂图旗",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	152125: {
		{
			"address":    "额尔古纳右旗",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	152126: {
		{
			"address":    "额尔古纳左旗",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	152127: {
		{
			"address":    "鄂伦春自治旗",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	152128: {
		{
			"address":    "鄂温克族自治旗",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	152129: {
		{
			"address":    "新巴尔虎右旗",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	152130: {
		{
			"address":    "新巴尔虎左旗",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	152131: {
		{
			"address":    "陈巴尔虎旗",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	152200: {
		{
			"address":    "兴安盟",
			"start_year": "",
			"end_year":   "",
		},
	},
	152201: {
		{
			"address":    "乌兰浩特市",
			"start_year": "",
			"end_year":   "",
		},
	},
	152202: {
		{
			"address":    "阿尔山市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	152221: {
		{
			"address":    "科尔沁右翼前旗",
			"start_year": "",
			"end_year":   "",
		},
	},
	152222: {
		{
			"address":    "科尔沁右翼中旗",
			"start_year": "",
			"end_year":   "",
		},
	},
	152223: {
		{
			"address":    "扎赉特旗",
			"start_year": "",
			"end_year":   "",
		},
	},
	152224: {
		{
			"address":    "突泉县",
			"start_year": "",
			"end_year":   "",
		},
	},
	152300: {
		{
			"address":    "哲里木盟",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	152301: {
		{
			"address":    "通辽市",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	152302: {
		{
			"address":    "霍林郭勒市",
			"start_year": "1985",
			"end_year":   "1998",
		},
	},
	152321: {
		{
			"address":    "通辽县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	152322: {
		{
			"address":    "科尔沁左翼中旗",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	152323: {
		{
			"address":    "科尔沁左翼后旗",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	152324: {
		{
			"address":    "开鲁县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	152325: {
		{
			"address":    "库伦旗",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	152326: {
		{
			"address":    "奈曼旗",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	152327: {
		{
			"address":    "扎鲁特旗",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	152400: {
		{
			"address":    "昭乌达盟",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	152401: {
		{
			"address":    "赤峰市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	152421: {
		{
			"address":    "赤峰县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "阿鲁科尔沁旗",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	152422: {
		{
			"address":    "巴林左旗",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	152423: {
		{
			"address":    "巴林右旗",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	152424: {
		{
			"address":    "林西县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	152425: {
		{
			"address":    "克什克腾旗",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	152426: {
		{
			"address":    "翁牛特旗",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	152427: {
		{
			"address":    "赤峰县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	152428: {
		{
			"address":    "喀喇沁旗",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	152429: {
		{
			"address":    "宁城县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	152430: {
		{
			"address":    "敖汉旗",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	152431: {
		{
			"address":    "阿鲁科尔沁旗",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	152500: {
		{
			"address":    "伊克昭盟",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "锡林郭勒盟",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	152501: {
		{
			"address":    "二连浩特市",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	152502: {
		{
			"address":    "锡林浩特市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	152521: {
		{
			"address":    "东胜县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "阿巴哈纳尔旗",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	152522: {
		{
			"address":    "达拉特旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "阿巴嘎旗",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	152523: {
		{
			"address":    "准格尔旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "苏尼特左旗",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	152524: {
		{
			"address":    "鄂托克前旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "苏尼特右旗",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	152525: {
		{
			"address":    "鄂托克旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "东乌珠穆沁旗",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	152526: {
		{
			"address":    "杭锦旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "西乌珠穆沁旗",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	152527: {
		{
			"address":    "乌审旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "太仆寺旗",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	152528: {
		{
			"address":    "伊金霍洛旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "镶黄旗",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	152529: {
		{
			"address":    "正镶白旗",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	152530: {
		{
			"address":    "正蓝旗",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	152531: {
		{
			"address":    "多伦县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	152600: {
		{
			"address":    "锡林郭勒盟",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "乌兰察布盟",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152601: {
		{
			"address":    "二连浩特市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "集宁市",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152602: {
		{
			"address":    "丰镇市",
			"start_year": "1990",
			"end_year":   "2002",
		},
	},
	152621: {
		{
			"address":    "阿巴哈纳尔旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "武川县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	152622: {
		{
			"address":    "阿巴嘎旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "和林格尔县",
			"start_year": "1982",
			"end_year":   "1994",
		},
	},
	152623: {
		{
			"address":    "苏尼特左旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "清水河县",
			"start_year": "1982",
			"end_year":   "1994",
		},
	},
	152624: {
		{
			"address":    "苏尼特右旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "卓资县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152625: {
		{
			"address":    "东乌珠穆沁旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "化德县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152626: {
		{
			"address":    "西乌珠穆沁旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "商都县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152627: {
		{
			"address":    "太仆寺旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "兴和县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152628: {
		{
			"address":    "镶黄旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "丰镇县",
			"start_year": "1982",
			"end_year":   "1989",
		},
	},
	152629: {
		{
			"address":    "正镶白旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "凉城县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152630: {
		{
			"address":    "正蓝旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "察哈尔右翼前旗",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152631: {
		{
			"address":    "多伦县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "察哈尔右翼中旗",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152632: {
		{
			"address":    "察哈尔右翼后旗",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152633: {
		{
			"address":    "达尔罕茂明安联合旗",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	152634: {
		{
			"address":    "四子王旗",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152700: {
		{
			"address":    "巴彦淖尔盟",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "伊克昭盟",
			"start_year": "1982",
			"end_year":   "2000",
		},
	},
	152701: {
		{
			"address":    "东胜市",
			"start_year": "1983",
			"end_year":   "2000",
		},
	},
	152721: {
		{
			"address":    "临河县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "东胜县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	152722: {
		{
			"address":    "五原县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "达拉特旗",
			"start_year": "1982",
			"end_year":   "2000",
		},
	},
	152723: {
		{
			"address":    "磴口县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "准格尔旗",
			"start_year": "1982",
			"end_year":   "2000",
		},
	},
	152724: {
		{
			"address":    "乌拉特前旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "鄂托克前旗",
			"start_year": "1982",
			"end_year":   "2000",
		},
	},
	152725: {
		{
			"address":    "乌拉特中后联合旗",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "乌拉特中旗",
			"start_year": "1981",
			"end_year":   "1981",
		},
		{
			"address":    "鄂托克旗",
			"start_year": "1982",
			"end_year":   "2000",
		},
	},
	152726: {
		{
			"address":    "杭锦后旗",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "杭锦旗",
			"start_year": "1982",
			"end_year":   "2000",
		},
	},
	152727: {
		{
			"address":    "潮格旗",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "乌拉特后旗",
			"start_year": "1981",
			"end_year":   "1981",
		},
		{
			"address":    "乌审旗",
			"start_year": "1982",
			"end_year":   "2000",
		},
	},
	152728: {
		{
			"address":    "伊金霍洛旗",
			"start_year": "1982",
			"end_year":   "2000",
		},
	},
	152800: {
		{
			"address":    "乌兰察布盟",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "巴彦淖尔盟",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152801: {
		{
			"address":    "集宁市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "临河市",
			"start_year": "1984",
			"end_year":   "2002",
		},
	},
	152821: {
		{
			"address":    "武川县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "临河县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	152822: {
		{
			"address":    "和林格尔县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "五原县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152823: {
		{
			"address":    "清水河县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "磴口县",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152824: {
		{
			"address":    "卓资县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "乌拉特前旗",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152825: {
		{
			"address":    "化德县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "乌拉特中旗",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152826: {
		{
			"address":    "商都县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "乌拉特后旗",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152827: {
		{
			"address":    "兴和县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "杭锦后旗",
			"start_year": "1982",
			"end_year":   "2002",
		},
	},
	152828: {
		{
			"address":    "丰镇县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	152829: {
		{
			"address":    "凉城县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	152830: {
		{
			"address":    "察哈尔右翼前旗",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	152831: {
		{
			"address":    "察哈尔右翼中旗",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	152832: {
		{
			"address":    "察哈尔右翼后旗",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	152833: {
		{
			"address":    "达尔罕茂明安联合旗",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	152834: {
		{
			"address":    "四子王旗",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	152900: {
		{
			"address":    "阿拉善盟",
			"start_year": "",
			"end_year":   "",
		},
	},
	152921: {
		{
			"address":    "阿拉善左旗",
			"start_year": "",
			"end_year":   "",
		},
	},
	152922: {
		{
			"address":    "阿拉善右旗",
			"start_year": "",
			"end_year":   "",
		},
	},
	152923: {
		{
			"address":    "额济纳旗",
			"start_year": "",
			"end_year":   "",
		},
	},
	210000: {
		{
			"address":    "辽宁省",
			"start_year": "",
			"end_year":   "",
		},
	},
	210100: {
		{
			"address":    "沈阳市",
			"start_year": "",
			"end_year":   "",
		},
	},
	210102: {
		{
			"address":    "和平区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210103: {
		{
			"address":    "沈河区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210104: {
		{
			"address":    "大东区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210105: {
		{
			"address":    "皇姑区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210106: {
		{
			"address":    "铁西区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210111: {
		{
			"address":    "苏家屯区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210112: {
		{
			"address":    "东陵区",
			"start_year": "",
			"end_year":   "2015",
		},
		{
			"address":    "浑南区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	210113: {
		{
			"address":    "新城子区",
			"start_year": "",
			"end_year":   "2005",
		},
		{
			"address":    "沈北新区",
			"start_year": "2006",
			"end_year":   "",
		},
	},
	210114: {
		{
			"address":    "于洪区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210115: {
		{
			"address":    "辽中区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	210121: {
		{
			"address":    "新民县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	210122: {
		{
			"address":    "辽中县",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	210123: {
		{
			"address":    "康平县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	210124: {
		{
			"address":    "法库县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	210181: {
		{
			"address":    "新民市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	210200: {
		{
			"address":    "旅大市",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "大连市",
			"start_year": "1981",
			"end_year":   "",
		},
	},
	210202: {
		{
			"address":    "中山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210203: {
		{
			"address":    "西岗区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210204: {
		{
			"address":    "沙河口区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210211: {
		{
			"address":    "甘井子区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210212: {
		{
			"address":    "旅顺口区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210213: {
		{
			"address":    "金州区",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	210214: {
		{
			"address":    "普兰店区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	210219: {
		{
			"address":    "瓦房店市",
			"start_year": "1985",
			"end_year":   "1985",
		},
	},
	210221: {
		{
			"address":    "金县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	210222: {
		{
			"address":    "新金县",
			"start_year": "",
			"end_year":   "1990",
		},
	},
	210223: {
		{
			"address":    "复县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	210224: {
		{
			"address":    "长海县",
			"start_year": "",
			"end_year":   "",
		},
	},
	210225: {
		{
			"address":    "庄河县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	210281: {
		{
			"address":    "瓦房店市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	210282: {
		{
			"address":    "普兰店市",
			"start_year": "1995",
			"end_year":   "2014",
		},
	},
	210283: {
		{
			"address":    "庄河市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	210300: {
		{
			"address":    "鞍山市",
			"start_year": "",
			"end_year":   "",
		},
	},
	210302: {
		{
			"address":    "铁东区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210303: {
		{
			"address":    "铁西区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210304: {
		{
			"address":    "立山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210311: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1983",
		},
		{
			"address":    "旧堡区",
			"start_year": "1984",
			"end_year":   "1995",
		},
		{
			"address":    "千山区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	210319: {
		{
			"address":    "海城市",
			"start_year": "1985",
			"end_year":   "1985",
		},
	},
	210321: {
		{
			"address":    "台安县",
			"start_year": "",
			"end_year":   "",
		},
	},
	210322: {
		{
			"address":    "海城县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	210323: {
		{
			"address":    "岫岩满族自治县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	210381: {
		{
			"address":    "海城市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	210400: {
		{
			"address":    "抚顺市",
			"start_year": "",
			"end_year":   "",
		},
	},
	210402: {
		{
			"address":    "新抚区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210403: {
		{
			"address":    "露天区",
			"start_year": "",
			"end_year":   "1998",
		},
		{
			"address":    "东洲区",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	210404: {
		{
			"address":    "望花区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210411: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1987",
		},
		{
			"address":    "顺城区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	210421: {
		{
			"address":    "抚顺县",
			"start_year": "",
			"end_year":   "",
		},
	},
	210422: {
		{
			"address":    "新宾县",
			"start_year": "",
			"end_year":   "1984",
		},
		{
			"address":    "新宾满族自治县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	210423: {
		{
			"address":    "清原县",
			"start_year": "",
			"end_year":   "1988",
		},
		{
			"address":    "清原满族自治县",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	210500: {
		{
			"address":    "本溪市",
			"start_year": "",
			"end_year":   "",
		},
	},
	210502: {
		{
			"address":    "平山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210503: {
		{
			"address":    "溪湖区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210504: {
		{
			"address":    "明山区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	210505: {
		{
			"address":    "南芬区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	210511: {
		{
			"address":    "立新区",
			"start_year": "",
			"end_year":   "1983",
		},
		{
			"address":    "南芬区",
			"start_year": "1985",
			"end_year":   "1985",
		},
	},
	210521: {
		{
			"address":    "本溪县",
			"start_year": "",
			"end_year":   "1988",
		},
		{
			"address":    "本溪满族自治县",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	210522: {
		{
			"address":    "桓仁县",
			"start_year": "",
			"end_year":   "1988",
		},
		{
			"address":    "桓仁满族自治县",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	210600: {
		{
			"address":    "丹东市",
			"start_year": "",
			"end_year":   "",
		},
	},
	210602: {
		{
			"address":    "元宝区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210603: {
		{
			"address":    "振兴区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210604: {
		{
			"address":    "振安区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210621: {
		{
			"address":    "凤城县",
			"start_year": "",
			"end_year":   "1984",
		},
		{
			"address":    "凤城满族自治县",
			"start_year": "1985",
			"end_year":   "1993",
		},
	},
	210622: {
		{
			"address":    "岫岩县",
			"start_year": "",
			"end_year":   "1984",
		},
		{
			"address":    "岫岩满族自治县",
			"start_year": "1985",
			"end_year":   "1991",
		},
	},
	210623: {
		{
			"address":    "东沟县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	210624: {
		{
			"address":    "宽甸县",
			"start_year": "",
			"end_year":   "1988",
		},
		{
			"address":    "宽甸满族自治县",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	210681: {
		{
			"address":    "东港市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	210682: {
		{
			"address":    "凤城市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	210700: {
		{
			"address":    "锦州市",
			"start_year": "",
			"end_year":   "",
		},
	},
	210702: {
		{
			"address":    "古塔区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210703: {
		{
			"address":    "凌河区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210704: {
		{
			"address":    "南票区",
			"start_year": "1982",
			"end_year":   "1988",
		},
	},
	210705: {
		{
			"address":    "葫芦岛区",
			"start_year": "1982",
			"end_year":   "1988",
		},
	},
	210706: {
		{
			"address":    "太和区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	210711: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "太和区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	210719: {
		{
			"address":    "锦西市",
			"start_year": "1985",
			"end_year":   "1985",
		},
	},
	210721: {
		{
			"address":    "锦西县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	210722: {
		{
			"address":    "兴城县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	210723: {
		{
			"address":    "绥中县",
			"start_year": "",
			"end_year":   "1988",
		},
	},
	210724: {
		{
			"address":    "锦县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	210725: {
		{
			"address":    "北镇县",
			"start_year": "",
			"end_year":   "1988",
		},
		{
			"address":    "北镇满族自治县",
			"start_year": "1989",
			"end_year":   "1994",
		},
	},
	210726: {
		{
			"address":    "黑山县",
			"start_year": "",
			"end_year":   "",
		},
	},
	210727: {
		{
			"address":    "义县",
			"start_year": "",
			"end_year":   "",
		},
	},
	210781: {
		{
			"address":    "凌海市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	210782: {
		{
			"address":    "北宁市",
			"start_year": "1995",
			"end_year":   "2005",
		},
		{
			"address":    "北镇市",
			"start_year": "2006",
			"end_year":   "",
		},
	},
	210800: {
		{
			"address":    "营口市",
			"start_year": "",
			"end_year":   "",
		},
	},
	210802: {
		{
			"address":    "站前区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210803: {
		{
			"address":    "西市区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210804: {
		{
			"address":    "鲅鱼圈区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	210811: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1983",
		},
		{
			"address":    "老边区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	210821: {
		{
			"address":    "营口县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	210822: {
		{
			"address":    "盘山县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	210823: {
		{
			"address":    "大洼县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	210824: {
		{
			"address":    "盖县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	210881: {
		{
			"address":    "盖州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	210882: {
		{
			"address":    "大石桥市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	210900: {
		{
			"address":    "阜新市",
			"start_year": "",
			"end_year":   "",
		},
	},
	210902: {
		{
			"address":    "海州区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210903: {
		{
			"address":    "新邱区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210904: {
		{
			"address":    "太平区",
			"start_year": "",
			"end_year":   "",
		},
	},
	210905: {
		{
			"address":    "清河门区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	210911: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1983",
		},
		{
			"address":    "细河区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	210921: {
		{
			"address":    "阜新蒙古族自治县",
			"start_year": "",
			"end_year":   "",
		},
	},
	210922: {
		{
			"address":    "彰武县",
			"start_year": "",
			"end_year":   "",
		},
	},
	211000: {
		{
			"address":    "辽阳市",
			"start_year": "",
			"end_year":   "",
		},
	},
	211002: {
		{
			"address":    "白塔区",
			"start_year": "",
			"end_year":   "",
		},
	},
	211003: {
		{
			"address":    "文圣区",
			"start_year": "",
			"end_year":   "",
		},
	},
	211004: {
		{
			"address":    "宏伟区",
			"start_year": "",
			"end_year":   "",
		},
	},
	211005: {
		{
			"address":    "弓长岭区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211011: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1983",
		},
		{
			"address":    "太子河区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211021: {
		{
			"address":    "辽阳县",
			"start_year": "",
			"end_year":   "",
		},
	},
	211022: {
		{
			"address":    "灯塔县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	211081: {
		{
			"address":    "灯塔市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	211100: {
		{
			"address":    "盘锦市",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211102: {
		{
			"address":    "盘山区",
			"start_year": "1984",
			"end_year":   "1985",
		},
		{
			"address":    "双台子区",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	211103: {
		{
			"address":    "兴隆台区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211104: {
		{
			"address":    "大洼区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	211111: {
		{
			"address":    "郊区",
			"start_year": "1984",
			"end_year":   "1985",
		},
	},
	211121: {
		{
			"address":    "大洼县",
			"start_year": "1984",
			"end_year":   "2015",
		},
	},
	211122: {
		{
			"address":    "盘山县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	211200: {
		{
			"address":    "铁岭市",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211202: {
		{
			"address":    "银州区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211203: {
		{
			"address":    "铁法区",
			"start_year": "1984",
			"end_year":   "1985",
		},
	},
	211204: {
		{
			"address":    "清河区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211221: {
		{
			"address":    "铁岭县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211222: {
		{
			"address":    "开原县",
			"start_year": "1984",
			"end_year":   "1987",
		},
	},
	211223: {
		{
			"address":    "西丰县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211224: {
		{
			"address":    "昌图县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211225: {
		{
			"address":    "康平县",
			"start_year": "1984",
			"end_year":   "1991",
		},
	},
	211226: {
		{
			"address":    "法库县",
			"start_year": "1984",
			"end_year":   "1991",
		},
	},
	211281: {
		{
			"address":    "铁法市",
			"start_year": "1995",
			"end_year":   "2001",
		},
		{
			"address":    "调兵山市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	211282: {
		{
			"address":    "开原市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	211300: {
		{
			"address":    "朝阳市",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211302: {
		{
			"address":    "双塔区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211303: {
		{
			"address":    "龙城区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211319: {
		{
			"address":    "北票市",
			"start_year": "1985",
			"end_year":   "1985",
		},
	},
	211321: {
		{
			"address":    "朝阳县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211322: {
		{
			"address":    "建平县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211323: {
		{
			"address":    "凌源县",
			"start_year": "1984",
			"end_year":   "1990",
		},
	},
	211324: {
		{
			"address":    "喀喇沁左翼蒙古族自治县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	211325: {
		{
			"address":    "建昌县",
			"start_year": "1984",
			"end_year":   "1988",
		},
	},
	211326: {
		{
			"address":    "北票县",
			"start_year": "1984",
			"end_year":   "1984",
		},
	},
	211381: {
		{
			"address":    "北票市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	211382: {
		{
			"address":    "凌源市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	211400: {
		{
			"address":    "锦西市",
			"start_year": "1989",
			"end_year":   "1993",
		},
		{
			"address":    "葫芦岛市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	211402: {
		{
			"address":    "连山区",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	211403: {
		{
			"address":    "龙港区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	211404: {
		{
			"address":    "南票区",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	211405: {
		{
			"address":    "葫芦岛区",
			"start_year": "1989",
			"end_year":   "1993",
		},
	},
	211421: {
		{
			"address":    "绥中县",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	211422: {
		{
			"address":    "建昌县",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	211481: {
		{
			"address":    "兴城市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	212100: {
		{
			"address":    "铁岭地区",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	212101: {
		{
			"address":    "铁岭市",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	212102: {
		{
			"address":    "铁法市",
			"start_year": "1981",
			"end_year":   "1983",
		},
	},
	212121: {
		{
			"address":    "铁岭县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	212122: {
		{
			"address":    "开原县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	212123: {
		{
			"address":    "西丰县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	212124: {
		{
			"address":    "昌图县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	212125: {
		{
			"address":    "康平县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	212126: {
		{
			"address":    "法库县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	212200: {
		{
			"address":    "朝阳地区",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	212201: {
		{
			"address":    "朝阳市",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	212221: {
		{
			"address":    "朝阳县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	212222: {
		{
			"address":    "建平县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	212223: {
		{
			"address":    "凌源县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	212224: {
		{
			"address":    "喀喇沁左翼蒙古族自治县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	212225: {
		{
			"address":    "建昌县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	212226: {
		{
			"address":    "北票县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	219001: {
		{
			"address":    "瓦房店市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	219002: {
		{
			"address":    "海城市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	219003: {
		{
			"address":    "锦西市",
			"start_year": "1986",
			"end_year":   "1988",
		},
	},
	219004: {
		{
			"address":    "兴城市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	219005: {
		{
			"address":    "铁法市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	219006: {
		{
			"address":    "北票市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	219007: {
		{
			"address":    "开原市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	219008: {
		{
			"address":    "普兰店市",
			"start_year": "1991",
			"end_year":   "1994",
		},
	},
	219009: {
		{
			"address":    "凌源市",
			"start_year": "1991",
			"end_year":   "1994",
		},
	},
	219010: {
		{
			"address":    "庄河市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	219011: {
		{
			"address":    "大石桥市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	219012: {
		{
			"address":    "盖州市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	219013: {
		{
			"address":    "新民市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	219014: {
		{
			"address":    "东港市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	219015: {
		{
			"address":    "凌海市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	219016: {
		{
			"address":    "凤城市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	220000: {
		{
			"address":    "吉林省",
			"start_year": "",
			"end_year":   "",
		},
	},
	220100: {
		{
			"address":    "长春市",
			"start_year": "",
			"end_year":   "",
		},
	},
	220102: {
		{
			"address":    "南关区",
			"start_year": "",
			"end_year":   "",
		},
	},
	220103: {
		{
			"address":    "宽城区",
			"start_year": "",
			"end_year":   "",
		},
	},
	220104: {
		{
			"address":    "朝阳区",
			"start_year": "",
			"end_year":   "",
		},
	},
	220105: {
		{
			"address":    "二道河子区",
			"start_year": "",
			"end_year":   "1994",
		},
		{
			"address":    "二道区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	220106: {
		{
			"address":    "绿园区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	220111: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	220112: {
		{
			"address":    "双阳区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	220113: {
		{
			"address":    "九台区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	220121: {
		{
			"address":    "榆树县",
			"start_year": "",
			"end_year":   "1989",
		},
	},
	220122: {
		{
			"address":    "农安县",
			"start_year": "",
			"end_year":   "",
		},
	},
	220123: {
		{
			"address":    "九台县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	220124: {
		{
			"address":    "德惠县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	220125: {
		{
			"address":    "双阳县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	220181: {
		{
			"address":    "九台市",
			"start_year": "1995",
			"end_year":   "2013",
		},
	},
	220182: {
		{
			"address":    "榆树市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	220183: {
		{
			"address":    "德惠市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	220184: {
		{
			"address":    "公主岭市",
			"start_year": "2020",
			"end_year":   "",
		},
	},
	220200: {
		{
			"address":    "吉林市",
			"start_year": "",
			"end_year":   "",
		},
	},
	220202: {
		{
			"address":    "昌邑区",
			"start_year": "",
			"end_year":   "",
		},
	},
	220203: {
		{
			"address":    "龙潭区",
			"start_year": "",
			"end_year":   "",
		},
	},
	220204: {
		{
			"address":    "船营区",
			"start_year": "",
			"end_year":   "",
		},
	},
	220211: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1991",
		},
		{
			"address":    "丰满区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	220221: {
		{
			"address":    "永吉县",
			"start_year": "",
			"end_year":   "",
		},
	},
	220222: {
		{
			"address":    "舒兰县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	220223: {
		{
			"address":    "磐石县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	220224: {
		{
			"address":    "蛟河县",
			"start_year": "",
			"end_year":   "1988",
		},
	},
	220225: {
		{
			"address":    "桦甸县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	220281: {
		{
			"address":    "蛟河市",
			"start_year": "1995",
			"end_year":   "",
		},
		{
			"address":    "桦甸市",
			"start_year": "2003",
			"end_year":   "2003",
		},
	},
	220282: {
		{
			"address":    "桦甸市",
			"start_year": "1995",
			"end_year":   "",
		},
		{
			"address":    "蛟河市",
			"start_year": "2003",
			"end_year":   "2003",
		},
	},
	220283: {
		{
			"address":    "舒兰市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	220284: {
		{
			"address":    "磐石市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	220300: {
		{
			"address":    "四平市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	220301: {
		{
			"address":    "铁西区",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	220302: {
		{
			"address":    "铁东区",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "铁西区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	220303: {
		{
			"address":    "铁东区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	220319: {
		{
			"address":    "公主岭市",
			"start_year": "1985",
			"end_year":   "1987",
		},
	},
	220321: {
		{
			"address":    "怀德县",
			"start_year": "1983",
			"end_year":   "1984",
		},
	},
	220322: {
		{
			"address":    "梨树县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	220323: {
		{
			"address":    "伊通县",
			"start_year": "1983",
			"end_year":   "1987",
		},
		{
			"address":    "伊通满族自治县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	220324: {
		{
			"address":    "双辽县",
			"start_year": "1983",
			"end_year":   "1995",
		},
	},
	220381: {
		{
			"address":    "公主岭市",
			"start_year": "1995",
			"end_year":   "2019",
		},
	},
	220382: {
		{
			"address":    "双辽市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	220400: {
		{
			"address":    "辽源市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	220401: {
		{
			"address":    "龙山区",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	220402: {
		{
			"address":    "西安区",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "龙山区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	220403: {
		{
			"address":    "西安区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	220421: {
		{
			"address":    "东丰县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	220422: {
		{
			"address":    "东辽县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	220500: {
		{
			"address":    "通化市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	220502: {
		{
			"address":    "东昌区",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	220503: {
		{
			"address":    "二道江区",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	220519: {
		{
			"address":    "梅河口市",
			"start_year": "1985",
			"end_year":   "1987",
		},
	},
	220521: {
		{
			"address":    "通化县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	220522: {
		{
			"address":    "集安县",
			"start_year": "1985",
			"end_year":   "1987",
		},
	},
	220523: {
		{
			"address":    "辉南县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	220524: {
		{
			"address":    "柳河县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	220581: {
		{
			"address":    "梅河口市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	220582: {
		{
			"address":    "集安市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	220600: {
		{
			"address":    "浑江市",
			"start_year": "1985",
			"end_year":   "1993",
		},
		{
			"address":    "白山市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	220602: {
		{
			"address":    "八道江区",
			"start_year": "1986",
			"end_year":   "2009",
		},
		{
			"address":    "浑江区",
			"start_year": "2010",
			"end_year":   "",
		},
	},
	220603: {
		{
			"address":    "三岔子区",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	220604: {
		{
			"address":    "临江区",
			"start_year": "1986",
			"end_year":   "1991",
		},
	},
	220605: {
		{
			"address":    "江源区",
			"start_year": "2006",
			"end_year":   "",
		},
	},
	220621: {
		{
			"address":    "抚松县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	220622: {
		{
			"address":    "靖宇县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	220623: {
		{
			"address":    "长白朝鲜族自治县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	220624: {
		{
			"address":    "临江县",
			"start_year": "1992",
			"end_year":   "1992",
		},
	},
	220625: {
		{
			"address":    "江源县",
			"start_year": "1995",
			"end_year":   "2005",
		},
	},
	220681: {
		{
			"address":    "临江市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	220700: {
		{
			"address":    "松原市",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	220702: {
		{
			"address":    "扶余区",
			"start_year": "1992",
			"end_year":   "1994",
		},
		{
			"address":    "宁江区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	220721: {
		{
			"address":    "前郭尔罗斯蒙古族自治县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	220722: {
		{
			"address":    "长岭县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	220723: {
		{
			"address":    "乾安县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	220724: {
		{
			"address":    "扶余县",
			"start_year": "1995",
			"end_year":   "2012",
		},
	},
	220781: {
		{
			"address":    "扶余市",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	220800: {
		{
			"address":    "白城市",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	220802: {
		{
			"address":    "洮北区",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	220821: {
		{
			"address":    "镇赉县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	220822: {
		{
			"address":    "通榆县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	220881: {
		{
			"address":    "洮南市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	220882: {
		{
			"address":    "大安市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	222100: {
		{
			"address":    "四平地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	222101: {
		{
			"address":    "四平市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	222102: {
		{
			"address":    "辽源市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	222121: {
		{
			"address":    "怀德县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	222122: {
		{
			"address":    "梨树县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	222123: {
		{
			"address":    "伊通县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	222124: {
		{
			"address":    "东丰县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	222125: {
		{
			"address":    "双辽县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	222200: {
		{
			"address":    "通化地区",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	222201: {
		{
			"address":    "通化市",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	222202: {
		{
			"address":    "浑江市",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	222221: {
		{
			"address":    "海龙县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	222222: {
		{
			"address":    "通化县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	222223: {
		{
			"address":    "柳河县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	222224: {
		{
			"address":    "辉南县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	222225: {
		{
			"address":    "集安县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	222226: {
		{
			"address":    "抚松县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	222227: {
		{
			"address":    "靖宇县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	222228: {
		{
			"address":    "长白朝鲜族自治县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	222300: {
		{
			"address":    "白城地区",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	222301: {
		{
			"address":    "白城市",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	222302: {
		{
			"address":    "洮南市",
			"start_year": "1987",
			"end_year":   "1992",
		},
	},
	222303: {
		{
			"address":    "扶余市",
			"start_year": "1987",
			"end_year":   "1991",
		},
	},
	222304: {
		{
			"address":    "大安市",
			"start_year": "1988",
			"end_year":   "1992",
		},
	},
	222321: {
		{
			"address":    "扶余县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	222322: {
		{
			"address":    "洮安县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	222323: {
		{
			"address":    "长岭县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	222324: {
		{
			"address":    "前郭尔罗斯蒙古族自治县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	222325: {
		{
			"address":    "大安县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	222326: {
		{
			"address":    "镇赉县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	222327: {
		{
			"address":    "通榆县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	222328: {
		{
			"address":    "乾安县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	222400: {
		{
			"address":    "延边朝鲜族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	222401: {
		{
			"address":    "延吉市",
			"start_year": "",
			"end_year":   "",
		},
	},
	222402: {
		{
			"address":    "图们市",
			"start_year": "",
			"end_year":   "",
		},
	},
	222403: {
		{
			"address":    "敦化市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	222404: {
		{
			"address":    "珲春市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	222405: {
		{
			"address":    "龙井市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	222406: {
		{
			"address":    "和龙市",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	222421: {
		{
			"address":    "延吉县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "龙井县",
			"start_year": "1984",
			"end_year":   "1987",
		},
	},
	222422: {
		{
			"address":    "敦化县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	222423: {
		{
			"address":    "和龙县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	222424: {
		{
			"address":    "汪清县",
			"start_year": "",
			"end_year":   "",
		},
	},
	222425: {
		{
			"address":    "珲春县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	222426: {
		{
			"address":    "安图县",
			"start_year": "",
			"end_year":   "",
		},
	},
	222427: {
		{
			"address":    "龙井县",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	222500: {
		{
			"address":    "德惠地区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	222521: {
		{
			"address":    "榆树县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	222522: {
		{
			"address":    "农安县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	222523: {
		{
			"address":    "九台县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	222524: {
		{
			"address":    "德惠县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	222525: {
		{
			"address":    "双阳县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	222600: {
		{
			"address":    "永吉地区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	222621: {
		{
			"address":    "永吉县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	222622: {
		{
			"address":    "舒兰县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	222623: {
		{
			"address":    "磐石县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	222624: {
		{
			"address":    "蛟河县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	222625: {
		{
			"address":    "桦甸县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	229001: {
		{
			"address":    "公主岭市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	229002: {
		{
			"address":    "梅河口市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	229003: {
		{
			"address":    "集安市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	229004: {
		{
			"address":    "桦甸市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	229005: {
		{
			"address":    "九台市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	229006: {
		{
			"address":    "蛟河市",
			"start_year": "1989",
			"end_year":   "1994",
		},
	},
	229007: {
		{
			"address":    "榆树市",
			"start_year": "1990",
			"end_year":   "1994",
		},
	},
	229008: {
		{
			"address":    "舒兰市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	229009: {
		{
			"address":    "大安市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	229010: {
		{
			"address":    "洮南市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	229011: {
		{
			"address":    "临江市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	229012: {
		{
			"address":    "德惠市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	230000: {
		{
			"address":    "黑龙江省",
			"start_year": "",
			"end_year":   "",
		},
	},
	230100: {
		{
			"address":    "哈尔滨市",
			"start_year": "",
			"end_year":   "",
		},
	},
	230102: {
		{
			"address":    "道里区",
			"start_year": "",
			"end_year":   "",
		},
	},
	230103: {
		{
			"address":    "南岗区",
			"start_year": "",
			"end_year":   "",
		},
	},
	230104: {
		{
			"address":    "道外区",
			"start_year": "",
			"end_year":   "",
		},
	},
	230105: {
		{
			"address":    "太平区",
			"start_year": "",
			"end_year":   "2003",
		},
	},
	230106: {
		{
			"address":    "香坊区",
			"start_year": "",
			"end_year":   "2003",
		},
	},
	230107: {
		{
			"address":    "动力区",
			"start_year": "",
			"end_year":   "2005",
		},
	},
	230108: {
		{
			"address":    "平房区",
			"start_year": "",
			"end_year":   "",
		},
	},
	230109: {
		{
			"address":    "松北区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	230110: {
		{
			"address":    "香坊区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	230111: {
		{
			"address":    "呼兰区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	230112: {
		{
			"address":    "阿城区",
			"start_year": "2006",
			"end_year":   "",
		},
	},
	230113: {
		{
			"address":    "双城区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	230121: {
		{
			"address":    "阿城县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "呼兰县",
			"start_year": "1984",
			"end_year":   "2003",
		},
	},
	230122: {
		{
			"address":    "呼兰县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "阿城县",
			"start_year": "1984",
			"end_year":   "1986",
		},
	},
	230123: {
		{
			"address":    "依兰县",
			"start_year": "1991",
			"end_year":   "",
		},
	},
	230124: {
		{
			"address":    "方正县",
			"start_year": "1991",
			"end_year":   "",
		},
	},
	230125: {
		{
			"address":    "宾县",
			"start_year": "1991",
			"end_year":   "",
		},
	},
	230126: {
		{
			"address":    "巴彦县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	230127: {
		{
			"address":    "木兰县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	230128: {
		{
			"address":    "通河县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	230129: {
		{
			"address":    "延寿县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	230181: {
		{
			"address":    "阿城市",
			"start_year": "1995",
			"end_year":   "2005",
		},
	},
	230182: {
		{
			"address":    "双城市",
			"start_year": "1996",
			"end_year":   "2013",
		},
	},
	230183: {
		{
			"address":    "尚志市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	230184: {
		{
			"address":    "五常市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	230200: {
		{
			"address":    "齐齐哈尔市",
			"start_year": "",
			"end_year":   "",
		},
	},
	230202: {
		{
			"address":    "龙沙区",
			"start_year": "",
			"end_year":   "",
		},
	},
	230203: {
		{
			"address":    "建华区",
			"start_year": "",
			"end_year":   "",
		},
	},
	230204: {
		{
			"address":    "铁锋区",
			"start_year": "",
			"end_year":   "",
		},
	},
	230205: {
		{
			"address":    "昂昂溪区",
			"start_year": "",
			"end_year":   "",
		},
	},
	230206: {
		{
			"address":    "富拉尔基区",
			"start_year": "",
			"end_year":   "",
		},
	},
	230207: {
		{
			"address":    "华安区",
			"start_year": "",
			"end_year":   "1983",
		},
		{
			"address":    "碾子山区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	230208: {
		{
			"address":    "梅里斯区",
			"start_year": "",
			"end_year":   "1987",
		},
		{
			"address":    "梅里斯达斡尔族区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	230221: {
		{
			"address":    "龙江县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	230222: {
		{
			"address":    "讷河县",
			"start_year": "1984",
			"end_year":   "1991",
		},
	},
	230223: {
		{
			"address":    "依安县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	230224: {
		{
			"address":    "泰来县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	230225: {
		{
			"address":    "甘南县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	230226: {
		{
			"address":    "杜尔伯特蒙古族自治县",
			"start_year": "1984",
			"end_year":   "1991",
		},
	},
	230227: {
		{
			"address":    "富裕县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	230228: {
		{
			"address":    "林甸县",
			"start_year": "1984",
			"end_year":   "1991",
		},
	},
	230229: {
		{
			"address":    "克山县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	230230: {
		{
			"address":    "克东县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	230231: {
		{
			"address":    "拜泉县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	230281: {
		{
			"address":    "讷河市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	230300: {
		{
			"address":    "鸡西市",
			"start_year": "",
			"end_year":   "",
		},
	},
	230302: {
		{
			"address":    "鸡冠区",
			"start_year": "",
			"end_year":   "",
		},
	},
	230303: {
		{
			"address":    "恒山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	230304: {
		{
			"address":    "滴道区",
			"start_year": "",
			"end_year":   "",
		},
	},
	230305: {
		{
			"address":    "梨树区",
			"start_year": "",
			"end_year":   "",
		},
	},
	230306: {
		{
			"address":    "城子河区",
			"start_year": "",
			"end_year":   "",
		},
	},
	230307: {
		{
			"address":    "麻山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	230321: {
		{
			"address":    "鸡东县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	230322: {
		{
			"address":    "虎林县",
			"start_year": "1993",
			"end_year":   "1995",
		},
	},
	230381: {
		{
			"address":    "虎林市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	230382: {
		{
			"address":    "密山市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	230400: {
		{
			"address":    "大庆市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "鹤岗市",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230402: {
		{
			"address":    "萨尔图区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "向阳区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230403: {
		{
			"address":    "龙凤区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "工农区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230404: {
		{
			"address":    "让胡路区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "南山区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230405: {
		{
			"address":    "红岗区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "兴安区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230406: {
		{
			"address":    "大同区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "东山区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230407: {
		{
			"address":    "兴山区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230421: {
		{
			"address":    "萝北县",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	230422: {
		{
			"address":    "绥滨县",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	230500: {
		{
			"address":    "鹤岗市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "双鸭山市",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230502: {
		{
			"address":    "向阳区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "尖山区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230503: {
		{
			"address":    "工农区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "岭东区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230504: {
		{
			"address":    "南山区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "岭西区",
			"start_year": "1982",
			"end_year":   "1986",
		},
	},
	230505: {
		{
			"address":    "兴安区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "四方台区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230506: {
		{
			"address":    "东山区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宝山区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230507: {
		{
			"address":    "兴山区",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	230521: {
		{
			"address":    "集贤县",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	230522: {
		{
			"address":    "友谊县",
			"start_year": "1991",
			"end_year":   "",
		},
	},
	230523: {
		{
			"address":    "宝清县",
			"start_year": "1991",
			"end_year":   "",
		},
	},
	230524: {
		{
			"address":    "饶河县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	230600: {
		{
			"address":    "双鸭山市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "大庆市",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230602: {
		{
			"address":    "尖山区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "萨尔图区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230603: {
		{
			"address":    "岭东区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "龙凤区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230604: {
		{
			"address":    "岭西区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "让胡路区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230605: {
		{
			"address":    "四方台区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "红岗区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230606: {
		{
			"address":    "宝山区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "大同区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	230621: {
		{
			"address":    "肇州县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	230622: {
		{
			"address":    "肇源县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	230623: {
		{
			"address":    "林甸县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	230624: {
		{
			"address":    "杜尔伯特蒙古族自治县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	230700: {
		{
			"address":    "伊春市",
			"start_year": "",
			"end_year":   "",
		},
	},
	230702: {
		{
			"address":    "伊春区",
			"start_year": "",
			"end_year":   "2018",
		},
	},
	230703: {
		{
			"address":    "南岔区",
			"start_year": "",
			"end_year":   "2018",
		},
	},
	230704: {
		{
			"address":    "友好区",
			"start_year": "",
			"end_year":   "2018",
		},
	},
	230705: {
		{
			"address":    "西林区",
			"start_year": "",
			"end_year":   "2018",
		},
	},
	230706: {
		{
			"address":    "翠峦区",
			"start_year": "",
			"end_year":   "2018",
		},
	},
	230707: {
		{
			"address":    "新青区",
			"start_year": "",
			"end_year":   "2018",
		},
	},
	230708: {
		{
			"address":    "美溪区",
			"start_year": "",
			"end_year":   "2018",
		},
	},
	230709: {
		{
			"address":    "大丰区",
			"start_year": "",
			"end_year":   "1983",
		},
		{
			"address":    "金山屯区",
			"start_year": "1984",
			"end_year":   "2018",
		},
	},
	230710: {
		{
			"address":    "五营区",
			"start_year": "",
			"end_year":   "2018",
		},
	},
	230711: {
		{
			"address":    "乌敏河区",
			"start_year": "",
			"end_year":   "1983",
		},
		{
			"address":    "乌马河区",
			"start_year": "1984",
			"end_year":   "2018",
		},
	},
	230712: {
		{
			"address":    "东风区",
			"start_year": "",
			"end_year":   "1983",
		},
		{
			"address":    "汤旺河区",
			"start_year": "1984",
			"end_year":   "2018",
		},
	},
	230713: {
		{
			"address":    "带岭区",
			"start_year": "",
			"end_year":   "2018",
		},
	},
	230714: {
		{
			"address":    "乌伊岭区",
			"start_year": "",
			"end_year":   "2018",
		},
	},
	230715: {
		{
			"address":    "红星区",
			"start_year": "",
			"end_year":   "2018",
		},
	},
	230716: {
		{
			"address":    "上甘岭区",
			"start_year": "",
			"end_year":   "2018",
		},
	},
	230717: {
		{
			"address":    "伊美区",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	230718: {
		{
			"address":    "乌翠区",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	230719: {
		{
			"address":    "友好区",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	230721: {
		{
			"address":    "铁力县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	230722: {
		{
			"address":    "嘉荫县",
			"start_year": "",
			"end_year":   "",
		},
	},
	230723: {
		{
			"address":    "汤旺县",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	230724: {
		{
			"address":    "丰林县",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	230725: {
		{
			"address":    "大箐山县",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	230726: {
		{
			"address":    "南岔县",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	230751: {
		{
			"address":    "金林区",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	230781: {
		{
			"address":    "铁力市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	230800: {
		{
			"address":    "佳木斯市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	230802: {
		{
			"address":    "永红区",
			"start_year": "1983",
			"end_year":   "2005",
		},
	},
	230803: {
		{
			"address":    "向阳区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	230804: {
		{
			"address":    "前进区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	230805: {
		{
			"address":    "东风区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	230811: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	230821: {
		{
			"address":    "富锦县",
			"start_year": "1984",
			"end_year":   "1987",
		},
	},
	230822: {
		{
			"address":    "桦南县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	230823: {
		{
			"address":    "依兰县",
			"start_year": "1984",
			"end_year":   "1990",
		},
	},
	230824: {
		{
			"address":    "友谊县",
			"start_year": "1984",
			"end_year":   "1990",
		},
	},
	230825: {
		{
			"address":    "集贤县",
			"start_year": "1984",
			"end_year":   "1986",
		},
	},
	230826: {
		{
			"address":    "桦川县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	230827: {
		{
			"address":    "宝清县",
			"start_year": "1984",
			"end_year":   "1990",
		},
	},
	230828: {
		{
			"address":    "汤原县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	230829: {
		{
			"address":    "绥滨县",
			"start_year": "1984",
			"end_year":   "1986",
		},
	},
	230830: {
		{
			"address":    "萝北县",
			"start_year": "1984",
			"end_year":   "1986",
		},
	},
	230831: {
		{
			"address":    "同江县",
			"start_year": "1984",
			"end_year":   "1986",
		},
	},
	230832: {
		{
			"address":    "饶河县",
			"start_year": "1984",
			"end_year":   "1992",
		},
	},
	230833: {
		{
			"address":    "抚远县",
			"start_year": "1984",
			"end_year":   "2015",
		},
	},
	230881: {
		{
			"address":    "同江市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	230882: {
		{
			"address":    "富锦市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	230883: {
		{
			"address":    "抚远市",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	230900: {
		{
			"address":    "七台河市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	230902: {
		{
			"address":    "新兴区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	230903: {
		{
			"address":    "桃山区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	230904: {
		{
			"address":    "茄子河区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	230921: {
		{
			"address":    "勃利县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	231000: {
		{
			"address":    "牡丹江市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	231002: {
		{
			"address":    "东安区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	231003: {
		{
			"address":    "阳明区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	231004: {
		{
			"address":    "爱民区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	231005: {
		{
			"address":    "西安区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	231006: {
		{
			"address":    "爱民区",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	231007: {
		{
			"address":    "阳明区",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	231011: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "1996",
		},
	},
	231020: {
		{
			"address":    "绥芬河市",
			"start_year": "1984",
			"end_year":   "1985",
		},
	},
	231021: {
		{
			"address":    "宁安县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	231022: {
		{
			"address":    "海林县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	231023: {
		{
			"address":    "穆棱县",
			"start_year": "1983",
			"end_year":   "1994",
		},
	},
	231024: {
		{
			"address":    "东宁县",
			"start_year": "1983",
			"end_year":   "2014",
		},
	},
	231025: {
		{
			"address":    "林口县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	231026: {
		{
			"address":    "密山县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	231027: {
		{
			"address":    "虎林县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	231081: {
		{
			"address":    "绥芬河市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	231083: {
		{
			"address":    "海林市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	231084: {
		{
			"address":    "宁安市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	231085: {
		{
			"address":    "穆棱市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	231086: {
		{
			"address":    "东宁市",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	231100: {
		{
			"address":    "黑河市",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	231101: {
		{
			"address":    "绥芬河市",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	231102: {
		{
			"address":    "爱辉区",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	231121: {
		{
			"address":    "嫩江县",
			"start_year": "1993",
			"end_year":   "2018",
		},
	},
	231122: {
		{
			"address":    "德都县",
			"start_year": "1993",
			"end_year":   "1995",
		},
	},
	231123: {
		{
			"address":    "逊克县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	231124: {
		{
			"address":    "孙吴县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	231181: {
		{
			"address":    "北安市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	231182: {
		{
			"address":    "五大连池市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	231183: {
		{
			"address":    "嫩江市",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	231200: {
		{
			"address":    "绥化市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	231202: {
		{
			"address":    "北林区",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	231221: {
		{
			"address":    "望奎县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	231222: {
		{
			"address":    "兰西县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	231223: {
		{
			"address":    "青冈县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	231224: {
		{
			"address":    "庆安县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	231225: {
		{
			"address":    "明水县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	231226: {
		{
			"address":    "绥棱县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	231281: {
		{
			"address":    "安达市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	231282: {
		{
			"address":    "肇东市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	231283: {
		{
			"address":    "海伦市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	232100: {
		{
			"address":    "绥化地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "松花江地区",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	232101: {
		{
			"address":    "双城市",
			"start_year": "1988",
			"end_year":   "1995",
		},
	},
	232102: {
		{
			"address":    "尚志市",
			"start_year": "1988",
			"end_year":   "1995",
		},
	},
	232103: {
		{
			"address":    "五常市",
			"start_year": "1993",
			"end_year":   "1995",
		},
	},
	232121: {
		{
			"address":    "海伦县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "阿城县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232122: {
		{
			"address":    "肇东县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宾县",
			"start_year": "1982",
			"end_year":   "1990",
		},
	},
	232123: {
		{
			"address":    "绥化县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "呼兰县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232124: {
		{
			"address":    "安达县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "双城县",
			"start_year": "1982",
			"end_year":   "1987",
		},
	},
	232125: {
		{
			"address":    "望奎县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "五常县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	232126: {
		{
			"address":    "兰西县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "巴彦县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	232127: {
		{
			"address":    "青冈县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "木兰县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	232128: {
		{
			"address":    "肇源县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "通河县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	232129: {
		{
			"address":    "肇州县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "尚志县",
			"start_year": "1982",
			"end_year":   "1987",
		},
	},
	232130: {
		{
			"address":    "庆安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "方正县",
			"start_year": "1982",
			"end_year":   "1990",
		},
	},
	232131: {
		{
			"address":    "明水县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "延寿县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	232132: {
		{
			"address":    "绥棱县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	232200: {
		{
			"address":    "松花江地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "嫩江地区",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232221: {
		{
			"address":    "五常县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "龙江县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232222: {
		{
			"address":    "双城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "讷河县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232223: {
		{
			"address":    "巴彦县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "依安县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232224: {
		{
			"address":    "呼兰县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "泰来县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232225: {
		{
			"address":    "宾县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "甘南县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232226: {
		{
			"address":    "阿城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "杜尔伯特蒙古族自治县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232227: {
		{
			"address":    "尚志县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "富裕县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232228: {
		{
			"address":    "木兰县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "林甸县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232229: {
		{
			"address":    "延寿县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "克山县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232230: {
		{
			"address":    "通河县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "克东县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232231: {
		{
			"address":    "方正县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "拜泉县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232300: {
		{
			"address":    "嫩江地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "绥化地区",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	232301: {
		{
			"address":    "绥化市",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	232302: {
		{
			"address":    "安达市",
			"start_year": "1984",
			"end_year":   "1998",
		},
	},
	232303: {
		{
			"address":    "肇东市",
			"start_year": "1986",
			"end_year":   "1998",
		},
	},
	232304: {
		{
			"address":    "海伦市",
			"start_year": "1989",
			"end_year":   "1998",
		},
	},
	232321: {
		{
			"address":    "讷河县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "海伦县",
			"start_year": "1982",
			"end_year":   "1988",
		},
	},
	232322: {
		{
			"address":    "拜泉县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "肇东县",
			"start_year": "1982",
			"end_year":   "1985",
		},
	},
	232323: {
		{
			"address":    "龙江县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	232324: {
		{
			"address":    "依安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "望奎县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	232325: {
		{
			"address":    "克山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "兰西县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	232326: {
		{
			"address":    "甘南县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "青冈县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	232327: {
		{
			"address":    "泰来县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "安达县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232328: {
		{
			"address":    "克东县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "肇源县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	232329: {
		{
			"address":    "富裕县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "肇州县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	232330: {
		{
			"address":    "林甸县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "庆安县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	232331: {
		{
			"address":    "杜尔伯特蒙古族自治县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "明水县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	232332: {
		{
			"address":    "绥棱县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	232400: {
		{
			"address":    "合江地区",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	232401: {
		{
			"address":    "佳木斯市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	232402: {
		{
			"address":    "七台河市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	232403: {
		{
			"address":    "永红区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	232404: {
		{
			"address":    "向阳区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	232405: {
		{
			"address":    "前进区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	232406: {
		{
			"address":    "东风区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	232411: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	232421: {
		{
			"address":    "桦南县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "富锦县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232422: {
		{
			"address":    "集贤县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "桦南县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232423: {
		{
			"address":    "宝清县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "依兰县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232424: {
		{
			"address":    "富锦县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "勃利县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232425: {
		{
			"address":    "依兰县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "集贤县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232426: {
		{
			"address":    "勃利县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "桦川县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232427: {
		{
			"address":    "汤原县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宝清县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232428: {
		{
			"address":    "桦川县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "汤原县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232429: {
		{
			"address":    "萝北县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "绥滨县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232430: {
		{
			"address":    "绥滨县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "萝北县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232431: {
		{
			"address":    "饶河县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "同江县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232432: {
		{
			"address":    "同江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "饶河县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	232433: {
		{
			"address":    "抚远县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	232500: {
		{
			"address":    "牡丹江地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	232501: {
		{
			"address":    "牡丹江市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	232502: {
		{
			"address":    "东凤区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "绥芬河市",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232503: {
		{
			"address":    "先锋区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "东凤区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232504: {
		{
			"address":    "爱民区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "先锋区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232505: {
		{
			"address":    "阳明区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "爱民区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232506: {
		{
			"address":    "阳明区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232511: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	232521: {
		{
			"address":    "海林县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宁安县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232522: {
		{
			"address":    "宁安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "海林县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232523: {
		{
			"address":    "林口县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "穆棱县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232524: {
		{
			"address":    "密山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "东宁县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232525: {
		{
			"address":    "穆棱县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "林口县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232526: {
		{
			"address":    "虎林县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "鸡东县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232527: {
		{
			"address":    "鸡东县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "密山县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232528: {
		{
			"address":    "东宁县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "虎林县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232581: {
		{
			"address":    "绥芬河市",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	232600: {
		{
			"address":    "黑河地区",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	232601: {
		{
			"address":    "黑河市",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	232602: {
		{
			"address":    "北安市",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	232603: {
		{
			"address":    "五大连池市",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	232621: {
		{
			"address":    "嫩江县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	232622: {
		{
			"address":    "北安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "嫩江县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	232623: {
		{
			"address":    "德都县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	232624: {
		{
			"address":    "爱辉县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	232625: {
		{
			"address":    "逊克县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	232626: {
		{
			"address":    "孙吴县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	232627: {
		{
			"address":    "通北县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	232700: {
		{
			"address":    "大兴安岭地区",
			"start_year": "",
			"end_year":   "",
		},
	},
	232701: {
		{
			"address":    "漠河市",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	232721: {
		{
			"address":    "呼玛县",
			"start_year": "",
			"end_year":   "",
		},
	},
	232722: {
		{
			"address":    "塔河县",
			"start_year": "1981",
			"end_year":   "",
		},
	},
	232723: {
		{
			"address":    "漠河县",
			"start_year": "1981",
			"end_year":   "2017",
		},
	},
	239001: {
		{
			"address":    "绥芬河市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	239002: {
		{
			"address":    "镜泊湖市",
			"start_year": "1986",
			"end_year":   "1986",
		},
		{
			"address":    "阿城市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	239003: {
		{
			"address":    "同江市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	239004: {
		{
			"address":    "富锦市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	239005: {
		{
			"address":    "铁力市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	239006: {
		{
			"address":    "密山市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	239007: {
		{
			"address":    "海林市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	239008: {
		{
			"address":    "讷河市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	239009: {
		{
			"address":    "北安市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	239010: {
		{
			"address":    "五大连池市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	239011: {
		{
			"address":    "宁安市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	310000: {
		{
			"address":    "上海市",
			"start_year": "",
			"end_year":   "",
		},
	},
	310101: {
		{
			"address":    "黄浦区",
			"start_year": "",
			"end_year":   "",
		},
	},
	310102: {
		{
			"address":    "南市区",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	310103: {
		{
			"address":    "卢湾区",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	310104: {
		{
			"address":    "徐汇区",
			"start_year": "",
			"end_year":   "",
		},
	},
	310105: {
		{
			"address":    "长宁区",
			"start_year": "",
			"end_year":   "",
		},
	},
	310106: {
		{
			"address":    "静安区",
			"start_year": "",
			"end_year":   "",
		},
	},
	310107: {
		{
			"address":    "普陀区",
			"start_year": "",
			"end_year":   "",
		},
	},
	310108: {
		{
			"address":    "闸北区",
			"start_year": "",
			"end_year":   "2014",
		},
	},
	310109: {
		{
			"address":    "虹口区",
			"start_year": "",
			"end_year":   "",
		},
	},
	310110: {
		{
			"address":    "杨浦区",
			"start_year": "",
			"end_year":   "",
		},
	},
	310111: {
		{
			"address":    "吴淞区",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	310112: {
		{
			"address":    "闵行区",
			"start_year": "1981",
			"end_year":   "",
		},
	},
	310113: {
		{
			"address":    "宝山区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	310114: {
		{
			"address":    "嘉定区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	310115: {
		{
			"address":    "浦东新区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	310116: {
		{
			"address":    "金山区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	310117: {
		{
			"address":    "松江区",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	310118: {
		{
			"address":    "青浦区",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	310119: {
		{
			"address":    "南汇区",
			"start_year": "2001",
			"end_year":   "2008",
		},
	},
	310120: {
		{
			"address":    "奉贤区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	310151: {
		{
			"address":    "崇明区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	310201: {
		{
			"address":    "上海县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	310202: {
		{
			"address":    "嘉定县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	310203: {
		{
			"address":    "宝山县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	310204: {
		{
			"address":    "川沙县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	310205: {
		{
			"address":    "南汇县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	310206: {
		{
			"address":    "奉贤县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	310207: {
		{
			"address":    "松江县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	310208: {
		{
			"address":    "金山县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	310209: {
		{
			"address":    "青浦县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	310210: {
		{
			"address":    "崇明县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	310221: {
		{
			"address":    "上海县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	310222: {
		{
			"address":    "嘉定县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	310223: {
		{
			"address":    "宝山县",
			"start_year": "1982",
			"end_year":   "1987",
		},
	},
	310224: {
		{
			"address":    "川沙县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	310225: {
		{
			"address":    "南汇县",
			"start_year": "1982",
			"end_year":   "2000",
		},
	},
	310226: {
		{
			"address":    "奉贤县",
			"start_year": "1982",
			"end_year":   "2000",
		},
	},
	310227: {
		{
			"address":    "松江县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	310228: {
		{
			"address":    "金山县",
			"start_year": "1982",
			"end_year":   "1996",
		},
	},
	310229: {
		{
			"address":    "青浦县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	310230: {
		{
			"address":    "崇明县",
			"start_year": "1982",
			"end_year":   "2015",
		},
	},
	320000: {
		{
			"address":    "江苏省",
			"start_year": "",
			"end_year":   "",
		},
	},
	320100: {
		{
			"address":    "南京市",
			"start_year": "",
			"end_year":   "",
		},
	},
	320102: {
		{
			"address":    "玄武区",
			"start_year": "",
			"end_year":   "",
		},
	},
	320103: {
		{
			"address":    "白下区",
			"start_year": "",
			"end_year":   "2012",
		},
	},
	320104: {
		{
			"address":    "秦淮区",
			"start_year": "",
			"end_year":   "",
		},
	},
	320105: {
		{
			"address":    "建邺区",
			"start_year": "",
			"end_year":   "",
		},
	},
	320106: {
		{
			"address":    "鼓楼区",
			"start_year": "",
			"end_year":   "",
		},
	},
	320107: {
		{
			"address":    "下关区",
			"start_year": "",
			"end_year":   "2012",
		},
	},
	320111: {
		{
			"address":    "浦口区",
			"start_year": "",
			"end_year":   "",
		},
	},
	320112: {
		{
			"address":    "大厂区",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	320113: {
		{
			"address":    "栖霞区",
			"start_year": "",
			"end_year":   "",
		},
	},
	320114: {
		{
			"address":    "雨花区",
			"start_year": "",
			"end_year":   "1983",
		},
		{
			"address":    "雨花台区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	320115: {
		{
			"address":    "江宁区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	320116: {
		{
			"address":    "六合区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	320117: {
		{
			"address":    "溧水区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	320118: {
		{
			"address":    "高淳区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	320121: {
		{
			"address":    "江宁县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	320122: {
		{
			"address":    "江浦县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	320123: {
		{
			"address":    "六合县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	320124: {
		{
			"address":    "溧水县",
			"start_year": "1983",
			"end_year":   "2012",
		},
	},
	320125: {
		{
			"address":    "高淳县",
			"start_year": "1983",
			"end_year":   "2012",
		},
	},
	320200: {
		{
			"address":    "无锡市",
			"start_year": "",
			"end_year":   "",
		},
	},
	320202: {
		{
			"address":    "崇安区",
			"start_year": "",
			"end_year":   "2014",
		},
	},
	320203: {
		{
			"address":    "南长区",
			"start_year": "",
			"end_year":   "2014",
		},
	},
	320204: {
		{
			"address":    "北塘区",
			"start_year": "",
			"end_year":   "2014",
		},
	},
	320205: {
		{
			"address":    "马山区",
			"start_year": "1987",
			"end_year":   "1999",
		},
		{
			"address":    "锡山区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	320206: {
		{
			"address":    "惠山区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	320211: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "1999",
		},
		{
			"address":    "滨湖区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	320213: {
		{
			"address":    "梁溪区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	320214: {
		{
			"address":    "新吴区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	320221: {
		{
			"address":    "江阴县",
			"start_year": "1983",
			"end_year":   "1986",
		},
	},
	320222: {
		{
			"address":    "无锡县",
			"start_year": "1983",
			"end_year":   "1994",
		},
	},
	320223: {
		{
			"address":    "宜兴县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	320281: {
		{
			"address":    "江阴市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	320282: {
		{
			"address":    "宜兴市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	320283: {
		{
			"address":    "锡山市",
			"start_year": "1995",
			"end_year":   "1999",
		},
	},
	320300: {
		{
			"address":    "徐州市",
			"start_year": "",
			"end_year":   "",
		},
	},
	320302: {
		{
			"address":    "鼓楼区",
			"start_year": "",
			"end_year":   "",
		},
	},
	320303: {
		{
			"address":    "云龙区",
			"start_year": "",
			"end_year":   "",
		},
	},
	320304: {
		{
			"address":    "矿区",
			"start_year": "",
			"end_year":   "1994",
		},
		{
			"address":    "九里区",
			"start_year": "1995",
			"end_year":   "2009",
		},
	},
	320305: {
		{
			"address":    "贾汪区",
			"start_year": "",
			"end_year":   "",
		},
	},
	320311: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1992",
		},
		{
			"address":    "泉山区",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	320312: {
		{
			"address":    "铜山区",
			"start_year": "2010",
			"end_year":   "",
		},
	},
	320321: {
		{
			"address":    "丰县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320322: {
		{
			"address":    "沛县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320323: {
		{
			"address":    "铜山县",
			"start_year": "1983",
			"end_year":   "2009",
		},
	},
	320324: {
		{
			"address":    "睢宁县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320325: {
		{
			"address":    "邳县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	320326: {
		{
			"address":    "新沂县",
			"start_year": "1983",
			"end_year":   "1989",
		},
	},
	320381: {
		{
			"address":    "新沂市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	320382: {
		{
			"address":    "邳州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	320400: {
		{
			"address":    "常州市",
			"start_year": "",
			"end_year":   "",
		},
	},
	320402: {
		{
			"address":    "天宁区",
			"start_year": "",
			"end_year":   "",
		},
	},
	320403: {
		{
			"address":    "广化区",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	320404: {
		{
			"address":    "钟楼区",
			"start_year": "",
			"end_year":   "",
		},
	},
	320405: {
		{
			"address":    "戚墅堰区",
			"start_year": "",
			"end_year":   "2014",
		},
	},
	320411: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "2001",
		},
		{
			"address":    "新北区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	320412: {
		{
			"address":    "武进区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	320413: {
		{
			"address":    "金坛区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	320421: {
		{
			"address":    "武进县",
			"start_year": "1983",
			"end_year":   "1994",
		},
	},
	320422: {
		{
			"address":    "金坛县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	320423: {
		{
			"address":    "溧阳县",
			"start_year": "1983",
			"end_year":   "1989",
		},
	},
	320481: {
		{
			"address":    "溧阳市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	320482: {
		{
			"address":    "金坛市",
			"start_year": "1995",
			"end_year":   "2014",
		},
	},
	320483: {
		{
			"address":    "武进市",
			"start_year": "1995",
			"end_year":   "2001",
		},
	},
	320500: {
		{
			"address":    "苏州市",
			"start_year": "",
			"end_year":   "",
		},
	},
	320502: {
		{
			"address":    "沧浪区",
			"start_year": "",
			"end_year":   "2011",
		},
	},
	320503: {
		{
			"address":    "平江区",
			"start_year": "",
			"end_year":   "2011",
		},
	},
	320504: {
		{
			"address":    "金阊区",
			"start_year": "",
			"end_year":   "2011",
		},
	},
	320505: {
		{
			"address":    "虎丘区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	320506: {
		{
			"address":    "吴中区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	320507: {
		{
			"address":    "相城区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	320508: {
		{
			"address":    "姑苏区",
			"start_year": "2012",
			"end_year":   "",
		},
	},
	320509: {
		{
			"address":    "吴江区",
			"start_year": "2012",
			"end_year":   "",
		},
	},
	320511: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "1999",
		},
	},
	320521: {
		{
			"address":    "沙洲县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	320522: {
		{
			"address":    "太仓县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	320523: {
		{
			"address":    "昆山县",
			"start_year": "1983",
			"end_year":   "1988",
		},
	},
	320524: {
		{
			"address":    "吴县",
			"start_year": "1983",
			"end_year":   "1994",
		},
	},
	320525: {
		{
			"address":    "吴江县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	320581: {
		{
			"address":    "常熟市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	320582: {
		{
			"address":    "张家港市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	320583: {
		{
			"address":    "昆山市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	320584: {
		{
			"address":    "吴江市",
			"start_year": "1995",
			"end_year":   "2011",
		},
	},
	320585: {
		{
			"address":    "太仓市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	320586: {
		{
			"address":    "吴县市",
			"start_year": "1995",
			"end_year":   "1999",
		},
	},
	320600: {
		{
			"address":    "南通市",
			"start_year": "",
			"end_year":   "",
		},
	},
	320602: {
		{
			"address":    "城中区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "城区",
			"start_year": "1983",
			"end_year":   "1990",
		},
		{
			"address":    "崇川区",
			"start_year": "1991",
			"end_year":   "",
		},
	},
	320603: {
		{
			"address":    "港闸区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	320611: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1990",
		},
		{
			"address":    "港闸区",
			"start_year": "1991",
			"end_year":   "",
		},
	},
	320612: {
		{
			"address":    "通州区",
			"start_year": "2009",
			"end_year":   "",
		},
	},
	320621: {
		{
			"address":    "海安县",
			"start_year": "1983",
			"end_year":   "2017",
		},
	},
	320622: {
		{
			"address":    "如皋县",
			"start_year": "1983",
			"end_year":   "1990",
		},
	},
	320623: {
		{
			"address":    "如东县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320624: {
		{
			"address":    "南通县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	320625: {
		{
			"address":    "海门县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	320626: {
		{
			"address":    "启东县",
			"start_year": "1983",
			"end_year":   "1988",
		},
	},
	320681: {
		{
			"address":    "启东市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	320682: {
		{
			"address":    "如皋市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	320683: {
		{
			"address":    "通州市",
			"start_year": "1995",
			"end_year":   "2008",
		},
	},
	320684: {
		{
			"address":    "海门市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	320685: {
		{
			"address":    "海安市",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	320700: {
		{
			"address":    "连云港市",
			"start_year": "",
			"end_year":   "",
		},
	},
	320702: {
		{
			"address":    "连云港区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "新海区",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	320703: {
		{
			"address":    "盐区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "连云区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320704: {
		{
			"address":    "新浦区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "云台区",
			"start_year": "1983",
			"end_year":   "2000",
		},
	},
	320705: {
		{
			"address":    "海州区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "新浦区",
			"start_year": "1986",
			"end_year":   "2013",
		},
	},
	320706: {
		{
			"address":    "海州区",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	320707: {
		{
			"address":    "赣榆区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	320711: {
		{
			"address":    "郊区",
			"start_year": "1981",
			"end_year":   "1982",
		},
	},
	320721: {
		{
			"address":    "赣榆县",
			"start_year": "1983",
			"end_year":   "2013",
		},
	},
	320722: {
		{
			"address":    "东海县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320723: {
		{
			"address":    "灌云县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320724: {
		{
			"address":    "灌南县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	320800: {
		{
			"address":    "淮阴市",
			"start_year": "1983",
			"end_year":   "1999",
		},
		{
			"address":    "淮安市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	320802: {
		{
			"address":    "清河区",
			"start_year": "1983",
			"end_year":   "2015",
		},
	},
	320803: {
		{
			"address":    "楚州区",
			"start_year": "2000",
			"end_year":   "2010",
		},
		{
			"address":    "淮安区",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	320804: {
		{
			"address":    "淮阴区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	320811: {
		{
			"address":    "清浦区",
			"start_year": "1983",
			"end_year":   "2015",
		},
	},
	320812: {
		{
			"address":    "清江浦区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	320813: {
		{
			"address":    "洪泽区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	320821: {
		{
			"address":    "淮阴县",
			"start_year": "1983",
			"end_year":   "1999",
		},
	},
	320822: {
		{
			"address":    "灌南县",
			"start_year": "1983",
			"end_year":   "1995",
		},
	},
	320823: {
		{
			"address":    "沭阳县",
			"start_year": "1983",
			"end_year":   "1995",
		},
	},
	320824: {
		{
			"address":    "宿迁县",
			"start_year": "1983",
			"end_year":   "1986",
		},
	},
	320825: {
		{
			"address":    "泗阳县",
			"start_year": "1983",
			"end_year":   "1995",
		},
	},
	320826: {
		{
			"address":    "涟水县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320827: {
		{
			"address":    "泗洪县",
			"start_year": "1983",
			"end_year":   "1995",
		},
	},
	320828: {
		{
			"address":    "淮安县",
			"start_year": "1983",
			"end_year":   "1986",
		},
	},
	320829: {
		{
			"address":    "洪泽县",
			"start_year": "1983",
			"end_year":   "2015",
		},
	},
	320830: {
		{
			"address":    "盱眙县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320831: {
		{
			"address":    "金湖县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320881: {
		{
			"address":    "宿迁市",
			"start_year": "1995",
			"end_year":   "1995",
		},
	},
	320882: {
		{
			"address":    "淮安市",
			"start_year": "1995",
			"end_year":   "1999",
		},
	},
	320900: {
		{
			"address":    "盐城市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320902: {
		{
			"address":    "城区",
			"start_year": "1983",
			"end_year":   "2002",
		},
		{
			"address":    "亭湖区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	320903: {
		{
			"address":    "盐都区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	320904: {
		{
			"address":    "大丰区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	320911: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "1995",
		},
	},
	320921: {
		{
			"address":    "响水县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320922: {
		{
			"address":    "滨海县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320923: {
		{
			"address":    "阜宁县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320924: {
		{
			"address":    "射阳县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320925: {
		{
			"address":    "建湖县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	320926: {
		{
			"address":    "大丰县",
			"start_year": "1983",
			"end_year":   "1995",
		},
	},
	320927: {
		{
			"address":    "东台县",
			"start_year": "1983",
			"end_year":   "1986",
		},
	},
	320928: {
		{
			"address":    "盐都县",
			"start_year": "1996",
			"end_year":   "2002",
		},
	},
	320981: {
		{
			"address":    "东台市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	320982: {
		{
			"address":    "大丰市",
			"start_year": "1996",
			"end_year":   "2014",
		},
	},
	321000: {
		{
			"address":    "扬州市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	321002: {
		{
			"address":    "广陵区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	321003: {
		{
			"address":    "邗江区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	321011: {
		{
			"address":    "郊区",
			"start_year": "1984",
			"end_year":   "2001",
		},
		{
			"address":    "维扬区",
			"start_year": "2002",
			"end_year":   "2010",
		},
	},
	321012: {
		{
			"address":    "江都区",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	321021: {
		{
			"address":    "兴化县",
			"start_year": "1983",
			"end_year":   "1986",
		},
	},
	321022: {
		{
			"address":    "高邮县",
			"start_year": "1983",
			"end_year":   "1990",
		},
	},
	321023: {
		{
			"address":    "宝应县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	321024: {
		{
			"address":    "靖江县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	321025: {
		{
			"address":    "泰兴县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	321026: {
		{
			"address":    "江都县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	321027: {
		{
			"address":    "邗江县",
			"start_year": "1983",
			"end_year":   "1999",
		},
	},
	321028: {
		{
			"address":    "泰县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	321029: {
		{
			"address":    "仪征县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	321081: {
		{
			"address":    "仪征市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	321082: {
		{
			"address":    "泰州市",
			"start_year": "1995",
			"end_year":   "1995",
		},
	},
	321083: {
		{
			"address":    "兴化市",
			"start_year": "1995",
			"end_year":   "1995",
		},
	},
	321084: {
		{
			"address":    "高邮市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	321085: {
		{
			"address":    "靖江市",
			"start_year": "1995",
			"end_year":   "1995",
		},
	},
	321086: {
		{
			"address":    "泰兴市",
			"start_year": "1995",
			"end_year":   "1995",
		},
	},
	321087: {
		{
			"address":    "姜堰市",
			"start_year": "1995",
			"end_year":   "1995",
		},
	},
	321088: {
		{
			"address":    "江都市",
			"start_year": "1995",
			"end_year":   "2010",
		},
	},
	321100: {
		{
			"address":    "镇江市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	321102: {
		{
			"address":    "城区",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "京口区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	321111: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "润州区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	321112: {
		{
			"address":    "丹徒区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	321121: {
		{
			"address":    "丹徒县",
			"start_year": "1983",
			"end_year":   "2001",
		},
	},
	321122: {
		{
			"address":    "丹阳县",
			"start_year": "1983",
			"end_year":   "1986",
		},
	},
	321123: {
		{
			"address":    "句容县",
			"start_year": "1983",
			"end_year":   "1994",
		},
	},
	321124: {
		{
			"address":    "扬中县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	321181: {
		{
			"address":    "丹阳市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	321182: {
		{
			"address":    "扬中市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	321183: {
		{
			"address":    "句容市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	321200: {
		{
			"address":    "泰州市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	321201: {
		{
			"address":    "泰州市",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	321202: {
		{
			"address":    "海陵区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	321203: {
		{
			"address":    "高港区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	321204: {
		{
			"address":    "姜堰区",
			"start_year": "2012",
			"end_year":   "",
		},
	},
	321281: {
		{
			"address":    "兴化市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	321282: {
		{
			"address":    "靖江市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	321283: {
		{
			"address":    "泰兴市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	321284: {
		{
			"address":    "姜堰市",
			"start_year": "1996",
			"end_year":   "2011",
		},
	},
	321300: {
		{
			"address":    "宿迁市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	321301: {
		{
			"address":    "常熟市",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	321302: {
		{
			"address":    "宿城区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	321311: {
		{
			"address":    "宿豫区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	321321: {
		{
			"address":    "宿豫县",
			"start_year": "1996",
			"end_year":   "2003",
		},
	},
	321322: {
		{
			"address":    "沭阳县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	321323: {
		{
			"address":    "泗阳县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	321324: {
		{
			"address":    "泗洪县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	322100: {
		{
			"address":    "徐州地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322121: {
		{
			"address":    "丰县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322122: {
		{
			"address":    "沛县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322123: {
		{
			"address":    "铜山县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322124: {
		{
			"address":    "睢宁县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322125: {
		{
			"address":    "邳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322126: {
		{
			"address":    "新沂县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322127: {
		{
			"address":    "东海县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322128: {
		{
			"address":    "赣榆县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322200: {
		{
			"address":    "淮阴地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322201: {
		{
			"address":    "清江市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322221: {
		{
			"address":    "淮阴县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322222: {
		{
			"address":    "灌云县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322223: {
		{
			"address":    "灌南县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322224: {
		{
			"address":    "沭阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322225: {
		{
			"address":    "宿迁县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322226: {
		{
			"address":    "泗阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322227: {
		{
			"address":    "涟水县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322228: {
		{
			"address":    "泗洪县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322229: {
		{
			"address":    "淮安县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322230: {
		{
			"address":    "洪泽县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322231: {
		{
			"address":    "盱眙县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322232: {
		{
			"address":    "金湖县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322300: {
		{
			"address":    "盐城地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322321: {
		{
			"address":    "响水县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322322: {
		{
			"address":    "滨海县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322323: {
		{
			"address":    "阜宁县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322324: {
		{
			"address":    "射阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322325: {
		{
			"address":    "建湖县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322326: {
		{
			"address":    "盐城县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322327: {
		{
			"address":    "大丰县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322328: {
		{
			"address":    "东台县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322400: {
		{
			"address":    "扬州地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322401: {
		{
			"address":    "扬州市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322402: {
		{
			"address":    "泰州市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322421: {
		{
			"address":    "兴化县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322422: {
		{
			"address":    "高邮县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322423: {
		{
			"address":    "宝应县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322424: {
		{
			"address":    "靖江县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322425: {
		{
			"address":    "泰兴县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322426: {
		{
			"address":    "江都县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322427: {
		{
			"address":    "邗江县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322428: {
		{
			"address":    "泰县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322429: {
		{
			"address":    "仪征县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322500: {
		{
			"address":    "南通地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322521: {
		{
			"address":    "海安县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322522: {
		{
			"address":    "如皋县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322523: {
		{
			"address":    "如东县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322524: {
		{
			"address":    "南通县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322525: {
		{
			"address":    "海门县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322526: {
		{
			"address":    "启东县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322600: {
		{
			"address":    "镇江地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322601: {
		{
			"address":    "镇江市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322621: {
		{
			"address":    "丹徒县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322622: {
		{
			"address":    "武进县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322623: {
		{
			"address":    "丹阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322624: {
		{
			"address":    "句容县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322625: {
		{
			"address":    "金坛县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322626: {
		{
			"address":    "溧水县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322627: {
		{
			"address":    "高淳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322628: {
		{
			"address":    "溧阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322629: {
		{
			"address":    "宜兴县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322630: {
		{
			"address":    "扬中县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322700: {
		{
			"address":    "苏州地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322721: {
		{
			"address":    "江阴县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322722: {
		{
			"address":    "无锡县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322723: {
		{
			"address":    "沙洲县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322724: {
		{
			"address":    "常熟县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322725: {
		{
			"address":    "太仓县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322726: {
		{
			"address":    "昆山县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322727: {
		{
			"address":    "吴县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	322728: {
		{
			"address":    "吴江县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	329001: {
		{
			"address":    "泰州市",
			"start_year": "1984",
			"end_year":   "1994",
		},
	},
	329002: {
		{
			"address":    "仪征市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	329003: {
		{
			"address":    "常熟市",
			"start_year": "1984",
			"end_year":   "1994",
		},
	},
	329004: {
		{
			"address":    "张家港市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	329005: {
		{
			"address":    "江阴市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	329006: {
		{
			"address":    "宿迁市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	329007: {
		{
			"address":    "丹阳市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	329008: {
		{
			"address":    "东台市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	329009: {
		{
			"address":    "兴化市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	329010: {
		{
			"address":    "淮安市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	329011: {
		{
			"address":    "宜兴市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	329012: {
		{
			"address":    "昆山市",
			"start_year": "1989",
			"end_year":   "1994",
		},
	},
	329013: {
		{
			"address":    "启东市",
			"start_year": "1989",
			"end_year":   "1994",
		},
	},
	329014: {
		{
			"address":    "新沂市",
			"start_year": "1990",
			"end_year":   "1994",
		},
	},
	329015: {
		{
			"address":    "溧阳市",
			"start_year": "1990",
			"end_year":   "1994",
		},
	},
	329016: {
		{
			"address":    "如皋市",
			"start_year": "1991",
			"end_year":   "1994",
		},
	},
	329017: {
		{
			"address":    "高邮市",
			"start_year": "1991",
			"end_year":   "1994",
		},
	},
	329018: {
		{
			"address":    "吴江市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	329019: {
		{
			"address":    "邳州市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	329020: {
		{
			"address":    "泰兴市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	329021: {
		{
			"address":    "通州市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	329022: {
		{
			"address":    "太仓市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	329023: {
		{
			"address":    "靖江市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	329024: {
		{
			"address":    "金坛市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	329025: {
		{
			"address":    "江都市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	329026: {
		{
			"address":    "海门市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	329027: {
		{
			"address":    "扬中市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	329028: {
		{
			"address":    "姜堰市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	330000: {
		{
			"address":    "浙江省",
			"start_year": "",
			"end_year":   "",
		},
	},
	330100: {
		{
			"address":    "杭州市",
			"start_year": "",
			"end_year":   "",
		},
	},
	330102: {
		{
			"address":    "上城区",
			"start_year": "",
			"end_year":   "",
		},
	},
	330103: {
		{
			"address":    "下城区",
			"start_year": "",
			"end_year":   "",
		},
	},
	330104: {
		{
			"address":    "江干区",
			"start_year": "",
			"end_year":   "",
		},
	},
	330105: {
		{
			"address":    "拱墅区",
			"start_year": "",
			"end_year":   "",
		},
	},
	330106: {
		{
			"address":    "西湖区",
			"start_year": "",
			"end_year":   "",
		},
	},
	330107: {
		{
			"address":    "半山区",
			"start_year": "",
			"end_year":   "1989",
		},
	},
	330108: {
		{
			"address":    "滨江区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	330109: {
		{
			"address":    "萧山区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	330110: {
		{
			"address":    "余杭区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	330111: {
		{
			"address":    "富阳区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	330112: {
		{
			"address":    "临安区",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	330121: {
		{
			"address":    "萧山县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	330122: {
		{
			"address":    "桐庐县",
			"start_year": "",
			"end_year":   "",
		},
	},
	330123: {
		{
			"address":    "富阳县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	330124: {
		{
			"address":    "临安县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	330125: {
		{
			"address":    "余杭县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	330126: {
		{
			"address":    "建德县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	330127: {
		{
			"address":    "淳安县",
			"start_year": "",
			"end_year":   "",
		},
	},
	330181: {
		{
			"address":    "萧山市",
			"start_year": "1995",
			"end_year":   "2000",
		},
	},
	330182: {
		{
			"address":    "建德市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	330183: {
		{
			"address":    "富阳市",
			"start_year": "1995",
			"end_year":   "2013",
		},
	},
	330184: {
		{
			"address":    "余杭市",
			"start_year": "1995",
			"end_year":   "2000",
		},
	},
	330185: {
		{
			"address":    "临安市",
			"start_year": "1996",
			"end_year":   "2016",
		},
	},
	330200: {
		{
			"address":    "宁波市",
			"start_year": "",
			"end_year":   "",
		},
	},
	330202: {
		{
			"address":    "镇明区",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	330203: {
		{
			"address":    "海曙区",
			"start_year": "",
			"end_year":   "",
		},
	},
	330204: {
		{
			"address":    "江东区",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	330205: {
		{
			"address":    "江北区",
			"start_year": "",
			"end_year":   "",
		},
	},
	330206: {
		{
			"address":    "滨海区",
			"start_year": "1985",
			"end_year":   "1986",
		},
		{
			"address":    "北仑区",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	330211: {
		{
			"address":    "镇海区",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	330212: {
		{
			"address":    "鄞州区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	330213: {
		{
			"address":    "奉化区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	330219: {
		{
			"address":    "余姚市",
			"start_year": "1985",
			"end_year":   "1985",
		},
	},
	330221: {
		{
			"address":    "镇海县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	330222: {
		{
			"address":    "慈溪县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	330223: {
		{
			"address":    "余姚县",
			"start_year": "1983",
			"end_year":   "1984",
		},
	},
	330224: {
		{
			"address":    "奉化县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	330225: {
		{
			"address":    "象山县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	330226: {
		{
			"address":    "宁海县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	330227: {
		{
			"address":    "鄞县",
			"start_year": "1983",
			"end_year":   "2001",
		},
	},
	330281: {
		{
			"address":    "余姚市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	330282: {
		{
			"address":    "慈溪市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	330283: {
		{
			"address":    "奉化市",
			"start_year": "1995",
			"end_year":   "2015",
		},
	},
	330300: {
		{
			"address":    "温州市",
			"start_year": "",
			"end_year":   "",
		},
	},
	330301: {
		{
			"address":    "东城区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	330302: {
		{
			"address":    "南城区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "城区",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "鹿城区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	330303: {
		{
			"address":    "西城区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "龙湾区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	330304: {
		{
			"address":    "瓯海区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	330305: {
		{
			"address":    "洞头区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	330321: {
		{
			"address":    "洞头县",
			"start_year": "1981",
			"end_year":   "1981",
		},
		{
			"address":    "瓯海县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	330322: {
		{
			"address":    "永嘉县",
			"start_year": "1981",
			"end_year":   "1981",
		},
		{
			"address":    "洞头县",
			"start_year": "1982",
			"end_year":   "2014",
		},
	},
	330323: {
		{
			"address":    "瑞安县",
			"start_year": "1981",
			"end_year":   "1981",
		},
		{
			"address":    "乐清县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	330324: {
		{
			"address":    "文成县",
			"start_year": "1981",
			"end_year":   "1981",
		},
		{
			"address":    "永嘉县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	330325: {
		{
			"address":    "平阳县",
			"start_year": "1981",
			"end_year":   "1981",
		},
		{
			"address":    "瑞安县",
			"start_year": "1982",
			"end_year":   "1986",
		},
	},
	330326: {
		{
			"address":    "乐清县",
			"start_year": "1981",
			"end_year":   "1981",
		},
		{
			"address":    "平阳县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	330327: {
		{
			"address":    "泰顺县",
			"start_year": "1981",
			"end_year":   "1981",
		},
		{
			"address":    "苍南县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	330328: {
		{
			"address":    "瓯海县",
			"start_year": "1981",
			"end_year":   "1981",
		},
		{
			"address":    "文成县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	330329: {
		{
			"address":    "苍南县",
			"start_year": "1981",
			"end_year":   "1981",
		},
		{
			"address":    "泰顺县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	330381: {
		{
			"address":    "瑞安市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	330382: {
		{
			"address":    "乐清市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	330383: {
		{
			"address":    "龙港市",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	330400: {
		{
			"address":    "嘉兴市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	330402: {
		{
			"address":    "城区",
			"start_year": "1983",
			"end_year":   "1992",
		},
		{
			"address":    "秀城区",
			"start_year": "1993",
			"end_year":   "2004",
		},
		{
			"address":    "南湖区",
			"start_year": "2005",
			"end_year":   "",
		},
	},
	330411: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "1998",
		},
		{
			"address":    "秀洲区",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	330421: {
		{
			"address":    "嘉善县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	330422: {
		{
			"address":    "平湖县",
			"start_year": "1983",
			"end_year":   "1990",
		},
	},
	330423: {
		{
			"address":    "海宁县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	330424: {
		{
			"address":    "海盐县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	330425: {
		{
			"address":    "桐乡县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	330481: {
		{
			"address":    "海宁市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	330482: {
		{
			"address":    "平湖市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	330483: {
		{
			"address":    "桐乡市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	330500: {
		{
			"address":    "湖州市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	330502: {
		{
			"address":    "城区",
			"start_year": "1983",
			"end_year":   "1987",
		},
		{
			"address":    "吴兴区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	330503: {
		{
			"address":    "南浔区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	330511: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	330521: {
		{
			"address":    "德清县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	330522: {
		{
			"address":    "长兴县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	330523: {
		{
			"address":    "安吉县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	330600: {
		{
			"address":    "绍兴市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	330602: {
		{
			"address":    "越城区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	330603: {
		{
			"address":    "柯桥区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	330604: {
		{
			"address":    "上虞区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	330621: {
		{
			"address":    "绍兴县",
			"start_year": "1983",
			"end_year":   "2012",
		},
	},
	330622: {
		{
			"address":    "上虞县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	330623: {
		{
			"address":    "嵊县",
			"start_year": "1983",
			"end_year":   "1994",
		},
	},
	330624: {
		{
			"address":    "新昌县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	330625: {
		{
			"address":    "诸暨县",
			"start_year": "1983",
			"end_year":   "1988",
		},
	},
	330681: {
		{
			"address":    "诸暨市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	330682: {
		{
			"address":    "上虞市",
			"start_year": "1995",
			"end_year":   "2012",
		},
	},
	330683: {
		{
			"address":    "嵊州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	330700: {
		{
			"address":    "金华市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	330701: {
		{
			"address":    "兰溪市",
			"start_year": "1985",
			"end_year":   "1985",
		},
	},
	330702: {
		{
			"address":    "婺城区",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	330703: {
		{
			"address":    "金东区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	330721: {
		{
			"address":    "金华县",
			"start_year": "1985",
			"end_year":   "1999",
		},
	},
	330722: {
		{
			"address":    "永康县",
			"start_year": "1985",
			"end_year":   "1991",
		},
	},
	330723: {
		{
			"address":    "武义县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	330724: {
		{
			"address":    "东阳县",
			"start_year": "1985",
			"end_year":   "1987",
		},
	},
	330725: {
		{
			"address":    "义乌县",
			"start_year": "1985",
			"end_year":   "1987",
		},
	},
	330726: {
		{
			"address":    "浦江县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	330727: {
		{
			"address":    "磐安县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	330781: {
		{
			"address":    "兰溪市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	330782: {
		{
			"address":    "义乌市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	330783: {
		{
			"address":    "东阳市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	330784: {
		{
			"address":    "永康市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	330800: {
		{
			"address":    "衢州市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	330802: {
		{
			"address":    "柯城区",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	330803: {
		{
			"address":    "衢江区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	330821: {
		{
			"address":    "衢县",
			"start_year": "1985",
			"end_year":   "2000",
		},
	},
	330822: {
		{
			"address":    "常山县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	330823: {
		{
			"address":    "江山县",
			"start_year": "1985",
			"end_year":   "1986",
		},
	},
	330824: {
		{
			"address":    "开化县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	330825: {
		{
			"address":    "龙游县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	330881: {
		{
			"address":    "江山市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	330900: {
		{
			"address":    "舟山市",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	330902: {
		{
			"address":    "定海区",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	330903: {
		{
			"address":    "普陀区",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	330921: {
		{
			"address":    "岱山县",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	330922: {
		{
			"address":    "嵊泗县",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	331000: {
		{
			"address":    "台州市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	331002: {
		{
			"address":    "椒江区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	331003: {
		{
			"address":    "黄岩区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	331004: {
		{
			"address":    "路桥区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	331021: {
		{
			"address":    "玉环县",
			"start_year": "1994",
			"end_year":   "2016",
		},
	},
	331022: {
		{
			"address":    "三门县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	331023: {
		{
			"address":    "天台县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	331024: {
		{
			"address":    "仙居县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	331081: {
		{
			"address":    "温岭市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	331082: {
		{
			"address":    "临海市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	331083: {
		{
			"address":    "玉环市",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	331100: {
		{
			"address":    "丽水市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	331102: {
		{
			"address":    "莲都区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	331121: {
		{
			"address":    "青田县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	331122: {
		{
			"address":    "缙云县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	331123: {
		{
			"address":    "遂昌县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	331124: {
		{
			"address":    "松阳县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	331125: {
		{
			"address":    "云和县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	331126: {
		{
			"address":    "庆元县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	331127: {
		{
			"address":    "景宁畲族自治县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	331181: {
		{
			"address":    "龙泉市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	332100: {
		{
			"address":    "嘉兴地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	332101: {
		{
			"address":    "湖州市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	332102: {
		{
			"address":    "嘉兴市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	332121: {
		{
			"address":    "嘉兴县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "嘉善县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	332122: {
		{
			"address":    "嘉善县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "平湖县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	332123: {
		{
			"address":    "平湖县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "海宁县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	332124: {
		{
			"address":    "海宁县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "海盐县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	332125: {
		{
			"address":    "海盐县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "桐乡县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	332126: {
		{
			"address":    "桐乡县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "德清县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	332127: {
		{
			"address":    "德清县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "长兴县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	332128: {
		{
			"address":    "吴兴县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "安吉县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	332129: {
		{
			"address":    "长兴县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	332130: {
		{
			"address":    "安吉县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	332200: {
		{
			"address":    "宁波地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	332221: {
		{
			"address":    "慈溪县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	332222: {
		{
			"address":    "余姚县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	332223: {
		{
			"address":    "奉化县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	332224: {
		{
			"address":    "象山县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	332225: {
		{
			"address":    "宁海县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	332226: {
		{
			"address":    "鄞县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	332227: {
		{
			"address":    "镇海县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	332300: {
		{
			"address":    "绍兴地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	332301: {
		{
			"address":    "绍兴市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	332321: {
		{
			"address":    "绍兴县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "上虞县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	332322: {
		{
			"address":    "上虞县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "嵊县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	332323: {
		{
			"address":    "嵊县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "新昌县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	332324: {
		{
			"address":    "新昌县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "诸暨县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	332325: {
		{
			"address":    "诸暨县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	332400: {
		{
			"address":    "温州地区",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "金华地区",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	332401: {
		{
			"address":    "金华市",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	332402: {
		{
			"address":    "衢州市",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	332421: {
		{
			"address":    "洞头县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "兰溪县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	332422: {
		{
			"address":    "永嘉县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "永康县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	332423: {
		{
			"address":    "瑞安县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "武义县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	332424: {
		{
			"address":    "文成县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "东阳县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	332425: {
		{
			"address":    "平阳县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "义乌县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	332426: {
		{
			"address":    "乐清县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "浦江县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	332427: {
		{
			"address":    "泰顺县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "常山县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	332428: {
		{
			"address":    "江山县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	332429: {
		{
			"address":    "开化县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	332430: {
		{
			"address":    "龙游县",
			"start_year": "1983",
			"end_year":   "1984",
		},
	},
	332431: {
		{
			"address":    "磐安县",
			"start_year": "1983",
			"end_year":   "1984",
		},
	},
	332500: {
		{
			"address":    "金华地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "丽水地区",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	332501: {
		{
			"address":    "金华市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "丽水市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	332502: {
		{
			"address":    "衢州市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "龙泉市",
			"start_year": "1990",
			"end_year":   "1994",
		},
	},
	332521: {
		{
			"address":    "金华县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "丽水县",
			"start_year": "1982",
			"end_year":   "1985",
		},
	},
	332522: {
		{
			"address":    "兰溪县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "青田县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	332523: {
		{
			"address":    "永康县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "云和县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	332524: {
		{
			"address":    "武义县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "龙泉县",
			"start_year": "1982",
			"end_year":   "1989",
		},
	},
	332525: {
		{
			"address":    "东阳县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "庆元县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	332526: {
		{
			"address":    "义乌县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "缙云县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	332527: {
		{
			"address":    "浦江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "遂昌县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	332528: {
		{
			"address":    "衢县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "松阳县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	332529: {
		{
			"address":    "常山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "景宁畲族自治县",
			"start_year": "1984",
			"end_year":   "1999",
		},
	},
	332530: {
		{
			"address":    "江山县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	332531: {
		{
			"address":    "开化县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	332581: {
		{
			"address":    "龙泉市",
			"start_year": "1995",
			"end_year":   "1999",
		},
	},
	332582: {
		{
			"address":    "丽水市",
			"start_year": "1995",
			"end_year":   "1999",
		},
	},
	332600: {
		{
			"address":    "丽水地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "台州地区",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	332601: {
		{
			"address":    "椒江市",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	332602: {
		{
			"address":    "临海市",
			"start_year": "1986",
			"end_year":   "1993",
		},
	},
	332603: {
		{
			"address":    "黄岩市",
			"start_year": "1989",
			"end_year":   "1993",
		},
	},
	332621: {
		{
			"address":    "丽水县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "临海县",
			"start_year": "1982",
			"end_year":   "1985",
		},
	},
	332622: {
		{
			"address":    "青田县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "黄岩县",
			"start_year": "1982",
			"end_year":   "1988",
		},
	},
	332623: {
		{
			"address":    "云和县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "温岭县",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	332624: {
		{
			"address":    "龙泉县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "仙居县",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	332625: {
		{
			"address":    "庆元县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "天台县",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	332626: {
		{
			"address":    "缙云县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "三门县",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	332627: {
		{
			"address":    "遂昌县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "玉环县",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	332700: {
		{
			"address":    "台州地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "舟山地区",
			"start_year": "1982",
			"end_year":   "1986",
		},
	},
	332701: {
		{
			"address":    "椒江市",
			"start_year": "1981",
			"end_year":   "1981",
		},
	},
	332721: {
		{
			"address":    "临海县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "定海县",
			"start_year": "1982",
			"end_year":   "1986",
		},
	},
	332722: {
		{
			"address":    "黄岩县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "普陀县",
			"start_year": "1982",
			"end_year":   "1986",
		},
	},
	332723: {
		{
			"address":    "温岭县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "岱山县",
			"start_year": "1982",
			"end_year":   "1986",
		},
	},
	332724: {
		{
			"address":    "仙居县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "嵊泗县",
			"start_year": "1982",
			"end_year":   "1986",
		},
	},
	332725: {
		{
			"address":    "天台县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	332726: {
		{
			"address":    "三门县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	332727: {
		{
			"address":    "玉环县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	332800: {
		{
			"address":    "舟山地区",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	332821: {
		{
			"address":    "定海县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	332822: {
		{
			"address":    "普陀县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	332823: {
		{
			"address":    "岱山县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	332824: {
		{
			"address":    "嵊泗县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	339001: {
		{
			"address":    "余姚市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	339002: {
		{
			"address":    "海宁市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	339003: {
		{
			"address":    "兰溪市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	339004: {
		{
			"address":    "瑞安市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	339005: {
		{
			"address":    "萧山市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	339006: {
		{
			"address":    "江山市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	339007: {
		{
			"address":    "安徽省",
			"start_year": "1987",
			"end_year":   "1987",
		},
		{
			"address":    "义乌市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	339008: {
		{
			"address":    "合肥市",
			"start_year": "1987",
			"end_year":   "1987",
		},
		{
			"address":    "东阳市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	339009: {
		{
			"address":    "慈溪市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	339010: {
		{
			"address":    "奉化市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	339011: {
		{
			"address":    "诸暨市",
			"start_year": "1989",
			"end_year":   "1994",
		},
	},
	339012: {
		{
			"address":    "平湖市",
			"start_year": "1991",
			"end_year":   "1994",
		},
	},
	339013: {
		{
			"address":    "建德市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	339014: {
		{
			"address":    "永康市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	339015: {
		{
			"address":    "上虞市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	339016: {
		{
			"address":    "桐乡市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	339017: {
		{
			"address":    "乐清市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	339018: {
		{
			"address":    "临海市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	339019: {
		{
			"address":    "富阳市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	339020: {
		{
			"address":    "温岭市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	339021: {
		{
			"address":    "余杭市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	340000: {
		{
			"address":    "安徽省",
			"start_year": "",
			"end_year":   "",
		},
	},
	340100: {
		{
			"address":    "合肥市",
			"start_year": "",
			"end_year":   "",
		},
	},
	340102: {
		{
			"address":    "东市区",
			"start_year": "",
			"end_year":   "2001",
		},
		{
			"address":    "瑶海区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	340103: {
		{
			"address":    "中市区",
			"start_year": "",
			"end_year":   "2001",
		},
		{
			"address":    "庐阳区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	340104: {
		{
			"address":    "西市区",
			"start_year": "",
			"end_year":   "2001",
		},
		{
			"address":    "蜀山区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	340111: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "2001",
		},
		{
			"address":    "包河区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	340121: {
		{
			"address":    "长丰县",
			"start_year": "",
			"end_year":   "",
		},
	},
	340122: {
		{
			"address":    "肥东县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	340123: {
		{
			"address":    "肥西县",
			"start_year": "",
			"end_year":   "",
		},
	},
	340124: {
		{
			"address":    "庐江县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	340181: {
		{
			"address":    "巢湖市",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	340200: {
		{
			"address":    "芜湖市",
			"start_year": "",
			"end_year":   "",
		},
	},
	340202: {
		{
			"address":    "镜湖区",
			"start_year": "",
			"end_year":   "",
		},
	},
	340203: {
		{
			"address":    "马塘区",
			"start_year": "",
			"end_year":   "2004",
		},
		{
			"address":    "弋江区",
			"start_year": "2005",
			"end_year":   "",
		},
	},
	340204: {
		{
			"address":    "新芜区",
			"start_year": "",
			"end_year":   "2004",
		},
	},
	340205: {
		{
			"address":    "裕溪口区",
			"start_year": "",
			"end_year":   "1989",
		},
		{
			"address":    "鸠江区",
			"start_year": "1990",
			"end_year":   "2004",
		},
	},
	340206: {
		{
			"address":    "四褐山区",
			"start_year": "",
			"end_year":   "1989",
		},
	},
	340207: {
		{
			"address":    "鸠江区",
			"start_year": "2005",
			"end_year":   "",
		},
	},
	340208: {
		{
			"address":    "三山区",
			"start_year": "2005",
			"end_year":   "",
		},
	},
	340211: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1989",
		},
	},
	340221: {
		{
			"address":    "芜湖县",
			"start_year": "",
			"end_year":   "",
		},
	},
	340222: {
		{
			"address":    "繁昌县",
			"start_year": "",
			"end_year":   "",
		},
	},
	340223: {
		{
			"address":    "南陵县",
			"start_year": "",
			"end_year":   "",
		},
	},
	340224: {
		{
			"address":    "青阳县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	340225: {
		{
			"address":    "无为县",
			"start_year": "2011",
			"end_year":   "2018",
		},
	},
	340281: {
		{
			"address":    "无为市",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	340300: {
		{
			"address":    "蚌埠市",
			"start_year": "",
			"end_year":   "",
		},
	},
	340302: {
		{
			"address":    "东市区",
			"start_year": "",
			"end_year":   "2003",
		},
		{
			"address":    "龙子湖区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	340303: {
		{
			"address":    "中市区",
			"start_year": "",
			"end_year":   "2003",
		},
		{
			"address":    "蚌山区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	340304: {
		{
			"address":    "西市区",
			"start_year": "",
			"end_year":   "2003",
		},
		{
			"address":    "禹会区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	340311: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "2003",
		},
		{
			"address":    "淮上区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	340321: {
		{
			"address":    "怀远县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	340322: {
		{
			"address":    "五河县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	340323: {
		{
			"address":    "固镇县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	340400: {
		{
			"address":    "淮南市",
			"start_year": "",
			"end_year":   "",
		},
	},
	340402: {
		{
			"address":    "大通区",
			"start_year": "",
			"end_year":   "",
		},
	},
	340403: {
		{
			"address":    "田家庵区",
			"start_year": "",
			"end_year":   "",
		},
	},
	340404: {
		{
			"address":    "谢家集区",
			"start_year": "",
			"end_year":   "",
		},
	},
	340405: {
		{
			"address":    "八公山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	340406: {
		{
			"address":    "潘集区",
			"start_year": "",
			"end_year":   "",
		},
	},
	340421: {
		{
			"address":    "凤台县",
			"start_year": "",
			"end_year":   "",
		},
	},
	340422: {
		{
			"address":    "寿县",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	340500: {
		{
			"address":    "马鞍山市",
			"start_year": "",
			"end_year":   "",
		},
	},
	340502: {
		{
			"address":    "金家庄区",
			"start_year": "",
			"end_year":   "2011",
		},
	},
	340503: {
		{
			"address":    "花山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	340504: {
		{
			"address":    "雨山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	340505: {
		{
			"address":    "向山区",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	340506: {
		{
			"address":    "博望区",
			"start_year": "2012",
			"end_year":   "",
		},
	},
	340511: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	340521: {
		{
			"address":    "当涂县",
			"start_year": "",
			"end_year":   "",
		},
	},
	340522: {
		{
			"address":    "含山县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	340523: {
		{
			"address":    "和县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	340600: {
		{
			"address":    "淮北市",
			"start_year": "",
			"end_year":   "",
		},
	},
	340602: {
		{
			"address":    "杜集区",
			"start_year": "",
			"end_year":   "",
		},
	},
	340603: {
		{
			"address":    "相山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	340604: {
		{
			"address":    "烈山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	340611: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	340621: {
		{
			"address":    "濉溪县",
			"start_year": "",
			"end_year":   "",
		},
	},
	340700: {
		{
			"address":    "铜陵市",
			"start_year": "",
			"end_year":   "",
		},
	},
	340702: {
		{
			"address":    "铜官山区",
			"start_year": "",
			"end_year":   "2014",
		},
	},
	340703: {
		{
			"address":    "狮子山区",
			"start_year": "",
			"end_year":   "2014",
		},
	},
	340704: {
		{
			"address":    "铜山区",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	340705: {
		{
			"address":    "铜官区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	340706: {
		{
			"address":    "义安区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	340711: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "",
		},
	},
	340721: {
		{
			"address":    "铜陵县",
			"start_year": "",
			"end_year":   "2014",
		},
	},
	340722: {
		{
			"address":    "枞阳县",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	340800: {
		{
			"address":    "安庆市",
			"start_year": "",
			"end_year":   "",
		},
	},
	340802: {
		{
			"address":    "迎江区",
			"start_year": "",
			"end_year":   "",
		},
	},
	340803: {
		{
			"address":    "大观区",
			"start_year": "",
			"end_year":   "",
		},
	},
	340811: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "2004",
		},
		{
			"address":    "宜秀区",
			"start_year": "2005",
			"end_year":   "",
		},
	},
	340821: {
		{
			"address":    "桐城县",
			"start_year": "1988",
			"end_year":   "1995",
		},
	},
	340822: {
		{
			"address":    "怀宁县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	340823: {
		{
			"address":    "枞阳县",
			"start_year": "1988",
			"end_year":   "2014",
		},
	},
	340824: {
		{
			"address":    "潜山县",
			"start_year": "1988",
			"end_year":   "2017",
		},
	},
	340825: {
		{
			"address":    "太湖县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	340826: {
		{
			"address":    "宿松县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	340827: {
		{
			"address":    "望江县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	340828: {
		{
			"address":    "岳西县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	340881: {
		{
			"address":    "桐城市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	340882: {
		{
			"address":    "潜山市",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	340901: {
		{
			"address":    "黄山市",
			"start_year": "1983",
			"end_year":   "1986",
		},
	},
	341000: {
		{
			"address":    "黄山市",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	341002: {
		{
			"address":    "屯溪区",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	341003: {
		{
			"address":    "黄山区",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	341004: {
		{
			"address":    "徽州区",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	341021: {
		{
			"address":    "歙县",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	341022: {
		{
			"address":    "休宁县",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	341023: {
		{
			"address":    "黟县",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	341024: {
		{
			"address":    "祁门县",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	341100: {
		{
			"address":    "滁州市",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	341102: {
		{
			"address":    "琅琊区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	341103: {
		{
			"address":    "南谯区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	341121: {
		{
			"address":    "天长县",
			"start_year": "1992",
			"end_year":   "1992",
		},
	},
	341122: {
		{
			"address":    "来安县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	341124: {
		{
			"address":    "全椒县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	341125: {
		{
			"address":    "定远县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	341126: {
		{
			"address":    "凤阳县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	341127: {
		{
			"address":    "嘉山县",
			"start_year": "1992",
			"end_year":   "1993",
		},
	},
	341181: {
		{
			"address":    "天长市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	341182: {
		{
			"address":    "明光市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	341200: {
		{
			"address":    "阜阳市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	341202: {
		{
			"address":    "颍州区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	341203: {
		{
			"address":    "颍东区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	341204: {
		{
			"address":    "颍泉区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	341221: {
		{
			"address":    "临泉县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	341222: {
		{
			"address":    "太和县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	341223: {
		{
			"address":    "涡阳县",
			"start_year": "1996",
			"end_year":   "1999",
		},
	},
	341224: {
		{
			"address":    "蒙城县",
			"start_year": "1996",
			"end_year":   "1999",
		},
	},
	341225: {
		{
			"address":    "阜南县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	341226: {
		{
			"address":    "颍上县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	341227: {
		{
			"address":    "利辛县",
			"start_year": "1996",
			"end_year":   "1999",
		},
	},
	341281: {
		{
			"address":    "亳州市",
			"start_year": "1996",
			"end_year":   "1999",
		},
	},
	341282: {
		{
			"address":    "界首市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	341300: {
		{
			"address":    "宿州市",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	341302: {
		{
			"address":    "埇桥区",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	341321: {
		{
			"address":    "砀山县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	341322: {
		{
			"address":    "萧县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	341323: {
		{
			"address":    "灵璧县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	341324: {
		{
			"address":    "泗县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	341400: {
		{
			"address":    "巢湖市",
			"start_year": "1999",
			"end_year":   "2010",
		},
	},
	341402: {
		{
			"address":    "居巢区",
			"start_year": "1999",
			"end_year":   "2010",
		},
	},
	341421: {
		{
			"address":    "庐江县",
			"start_year": "1999",
			"end_year":   "2010",
		},
	},
	341422: {
		{
			"address":    "无为县",
			"start_year": "1999",
			"end_year":   "2010",
		},
	},
	341423: {
		{
			"address":    "含山县",
			"start_year": "1999",
			"end_year":   "2010",
		},
	},
	341424: {
		{
			"address":    "和县",
			"start_year": "1999",
			"end_year":   "2010",
		},
	},
	341500: {
		{
			"address":    "六安市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	341502: {
		{
			"address":    "金安区",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	341503: {
		{
			"address":    "裕安区",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	341504: {
		{
			"address":    "叶集区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	341521: {
		{
			"address":    "寿县",
			"start_year": "1999",
			"end_year":   "2014",
		},
	},
	341522: {
		{
			"address":    "霍邱县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	341523: {
		{
			"address":    "舒城县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	341524: {
		{
			"address":    "金寨县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	341525: {
		{
			"address":    "霍山县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	341600: {
		{
			"address":    "亳州市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341602: {
		{
			"address":    "谯城区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341621: {
		{
			"address":    "涡阳县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341622: {
		{
			"address":    "蒙城县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341623: {
		{
			"address":    "利辛县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341700: {
		{
			"address":    "池州市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341702: {
		{
			"address":    "贵池区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341721: {
		{
			"address":    "东至县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341722: {
		{
			"address":    "石台县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341723: {
		{
			"address":    "青阳县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341800: {
		{
			"address":    "宣城市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341802: {
		{
			"address":    "宣州区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341821: {
		{
			"address":    "郎溪县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341822: {
		{
			"address":    "广德县",
			"start_year": "2000",
			"end_year":   "2018",
		},
	},
	341823: {
		{
			"address":    "泾县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341824: {
		{
			"address":    "绩溪县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341825: {
		{
			"address":    "旌德县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341881: {
		{
			"address":    "宁国市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	341882: {
		{
			"address":    "广德市",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	342100: {
		{
			"address":    "阜阳地区",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	342101: {
		{
			"address":    "阜阳市",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	342102: {
		{
			"address":    "亳州市",
			"start_year": "1986",
			"end_year":   "1995",
		},
	},
	342103: {
		{
			"address":    "界首市",
			"start_year": "1989",
			"end_year":   "1995",
		},
	},
	342121: {
		{
			"address":    "阜阳县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	342122: {
		{
			"address":    "临泉县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	342123: {
		{
			"address":    "太和县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	342124: {
		{
			"address":    "涡阳县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	342125: {
		{
			"address":    "蒙城县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	342126: {
		{
			"address":    "亳县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	342127: {
		{
			"address":    "阜南县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	342128: {
		{
			"address":    "颍上县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	342129: {
		{
			"address":    "界首县",
			"start_year": "",
			"end_year":   "1988",
		},
	},
	342130: {
		{
			"address":    "利辛县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	342200: {
		{
			"address":    "宿县地区",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	342201: {
		{
			"address":    "宿州市",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	342221: {
		{
			"address":    "砀山县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	342222: {
		{
			"address":    "萧县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	342223: {
		{
			"address":    "宿县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	342224: {
		{
			"address":    "灵璧县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	342225: {
		{
			"address":    "泗县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	342226: {
		{
			"address":    "怀远县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	342227: {
		{
			"address":    "五河县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	342228: {
		{
			"address":    "固镇县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	342300: {
		{
			"address":    "滁县地区",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	342301: {
		{
			"address":    "滁州市",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	342321: {
		{
			"address":    "天长县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	342322: {
		{
			"address":    "来安县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	342323: {
		{
			"address":    "滁县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	342324: {
		{
			"address":    "全椒县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	342325: {
		{
			"address":    "定远县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	342326: {
		{
			"address":    "凤阳县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	342327: {
		{
			"address":    "嘉山县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	342400: {
		{
			"address":    "六安地区",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	342401: {
		{
			"address":    "六安市",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	342421: {
		{
			"address":    "六安县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	342422: {
		{
			"address":    "寿县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	342423: {
		{
			"address":    "霍邱县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	342424: {
		{
			"address":    "肥西县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	342425: {
		{
			"address":    "舒城县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	342426: {
		{
			"address":    "金寨县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	342427: {
		{
			"address":    "霍山县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	342500: {
		{
			"address":    "宣城地区",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	342501: {
		{
			"address":    "宣城市",
			"start_year": "1987",
			"end_year":   "1999",
		},
	},
	342502: {
		{
			"address":    "宁国市",
			"start_year": "1997",
			"end_year":   "1999",
		},
	},
	342521: {
		{
			"address":    "宣城县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	342522: {
		{
			"address":    "郎溪县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	342523: {
		{
			"address":    "广德县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	342524: {
		{
			"address":    "宁国县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	342525: {
		{
			"address":    "当涂县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	342526: {
		{
			"address":    "繁昌县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	342527: {
		{
			"address":    "南陵县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	342528: {
		{
			"address":    "青阳县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	342529: {
		{
			"address":    "泾县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	342530: {
		{
			"address":    "旌德县",
			"start_year": "1987",
			"end_year":   "1999",
		},
	},
	342531: {
		{
			"address":    "绩溪县",
			"start_year": "1987",
			"end_year":   "1999",
		},
	},
	342600: {
		{
			"address":    "巢湖地区",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	342601: {
		{
			"address":    "巢湖市",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	342621: {
		{
			"address":    "肥东县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	342622: {
		{
			"address":    "庐江县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	342623: {
		{
			"address":    "无为县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	342624: {
		{
			"address":    "巢县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	342625: {
		{
			"address":    "含山县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	342626: {
		{
			"address":    "和县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	342700: {
		{
			"address":    "徽州地区",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	342701: {
		{
			"address":    "屯溪市",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	342721: {
		{
			"address":    "绩溪县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	342722: {
		{
			"address":    "旌德县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	342723: {
		{
			"address":    "歙县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	342724: {
		{
			"address":    "休宁县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	342725: {
		{
			"address":    "黟县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	342726: {
		{
			"address":    "祁门县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	342727: {
		{
			"address":    "太平县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	342728: {
		{
			"address":    "石台县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	342800: {
		{
			"address":    "安庆地区",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	342821: {
		{
			"address":    "怀宁县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	342822: {
		{
			"address":    "桐城县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	342823: {
		{
			"address":    "枞阳县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	342824: {
		{
			"address":    "潜山县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	342825: {
		{
			"address":    "太湖县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	342826: {
		{
			"address":    "宿松县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	342827: {
		{
			"address":    "望江县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	342828: {
		{
			"address":    "岳西县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	342829: {
		{
			"address":    "东至县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	342830: {
		{
			"address":    "贵池县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	342831: {
		{
			"address":    "石台县",
			"start_year": "1987",
			"end_year":   "1987",
		},
	},
	342900: {
		{
			"address":    "池州地区",
			"start_year": "1988",
			"end_year":   "1999",
		},
	},
	342901: {
		{
			"address":    "贵池市",
			"start_year": "1988",
			"end_year":   "1999",
		},
	},
	342921: {
		{
			"address":    "东至县",
			"start_year": "1988",
			"end_year":   "1999",
		},
	},
	342922: {
		{
			"address":    "石台县",
			"start_year": "1988",
			"end_year":   "1999",
		},
	},
	342923: {
		{
			"address":    "青阳县",
			"start_year": "1988",
			"end_year":   "1999",
		},
	},
	349001: {
		{
			"address":    "天长市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	349002: {
		{
			"address":    "明光市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	350000: {
		{
			"address":    "福建省",
			"start_year": "",
			"end_year":   "",
		},
	},
	350100: {
		{
			"address":    "福州市",
			"start_year": "",
			"end_year":   "",
		},
	},
	350102: {
		{
			"address":    "鼓楼区",
			"start_year": "",
			"end_year":   "",
		},
	},
	350103: {
		{
			"address":    "台江区",
			"start_year": "",
			"end_year":   "",
		},
	},
	350104: {
		{
			"address":    "仓山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	350105: {
		{
			"address":    "马尾区",
			"start_year": "",
			"end_year":   "",
		},
	},
	350111: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1994",
		},
		{
			"address":    "晋安区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	350112: {
		{
			"address":    "长乐区",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	350121: {
		{
			"address":    "闽侯县",
			"start_year": "",
			"end_year":   "",
		},
	},
	350122: {
		{
			"address":    "连江县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350123: {
		{
			"address":    "罗源县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350124: {
		{
			"address":    "闽清县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350125: {
		{
			"address":    "永泰县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350126: {
		{
			"address":    "长乐县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	350127: {
		{
			"address":    "福清县",
			"start_year": "1983",
			"end_year":   "1989",
		},
	},
	350128: {
		{
			"address":    "平潭县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350181: {
		{
			"address":    "福清市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	350182: {
		{
			"address":    "长乐市",
			"start_year": "1995",
			"end_year":   "2016",
		},
	},
	350200: {
		{
			"address":    "厦门市",
			"start_year": "",
			"end_year":   "",
		},
	},
	350202: {
		{
			"address":    "鼓浪屿区",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	350203: {
		{
			"address":    "思明区",
			"start_year": "",
			"end_year":   "",
		},
	},
	350204: {
		{
			"address":    "开元区",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	350205: {
		{
			"address":    "杏林区",
			"start_year": "",
			"end_year":   "2002",
		},
		{
			"address":    "海沧区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	350206: {
		{
			"address":    "湖里区",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	350211: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1986",
		},
		{
			"address":    "集美区",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	350212: {
		{
			"address":    "同安区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	350213: {
		{
			"address":    "翔安区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	350221: {
		{
			"address":    "同安县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	350300: {
		{
			"address":    "莆田市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350302: {
		{
			"address":    "城厢区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350303: {
		{
			"address":    "涵江区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350304: {
		{
			"address":    "荔城区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	350305: {
		{
			"address":    "秀屿区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	350321: {
		{
			"address":    "莆田县",
			"start_year": "1983",
			"end_year":   "2001",
		},
	},
	350322: {
		{
			"address":    "仙游县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350400: {
		{
			"address":    "三明市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350402: {
		{
			"address":    "梅列区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350403: {
		{
			"address":    "三元区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350420: {
		{
			"address":    "永安市",
			"start_year": "1985",
			"end_year":   "1985",
		},
	},
	350421: {
		{
			"address":    "明溪县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350422: {
		{
			"address":    "永安县",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	350423: {
		{
			"address":    "清流县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350424: {
		{
			"address":    "宁化县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350425: {
		{
			"address":    "大田县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350426: {
		{
			"address":    "尤溪县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350427: {
		{
			"address":    "沙县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350428: {
		{
			"address":    "将乐县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350429: {
		{
			"address":    "泰宁县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350430: {
		{
			"address":    "建宁县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	350481: {
		{
			"address":    "永安市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	350500: {
		{
			"address":    "泉州市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350501: {
		{
			"address":    "永安市",
			"start_year": "1984",
			"end_year":   "1984",
		},
	},
	350502: {
		{
			"address":    "鲤城区",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350503: {
		{
			"address":    "丰泽区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	350504: {
		{
			"address":    "洛江区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	350505: {
		{
			"address":    "泉港区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	350521: {
		{
			"address":    "惠安县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350522: {
		{
			"address":    "晋江县",
			"start_year": "1985",
			"end_year":   "1991",
		},
	},
	350523: {
		{
			"address":    "南安县",
			"start_year": "1985",
			"end_year":   "1992",
		},
	},
	350524: {
		{
			"address":    "安溪县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350525: {
		{
			"address":    "永春县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350526: {
		{
			"address":    "德化县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350527: {
		{
			"address":    "金门县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350581: {
		{
			"address":    "石狮市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	350582: {
		{
			"address":    "晋江市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	350583: {
		{
			"address":    "南安市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	350600: {
		{
			"address":    "漳州市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350602: {
		{
			"address":    "芗城区",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350603: {
		{
			"address":    "龙文区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	350621: {
		{
			"address":    "龙海县",
			"start_year": "1985",
			"end_year":   "1992",
		},
	},
	350622: {
		{
			"address":    "云霄县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350623: {
		{
			"address":    "漳浦县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350624: {
		{
			"address":    "诏安县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350625: {
		{
			"address":    "长泰县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350626: {
		{
			"address":    "东山县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350627: {
		{
			"address":    "南靖县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350628: {
		{
			"address":    "平和县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350629: {
		{
			"address":    "华安县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	350681: {
		{
			"address":    "龙海市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	350700: {
		{
			"address":    "南平市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	350702: {
		{
			"address":    "延平区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	350703: {
		{
			"address":    "建阳区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	350721: {
		{
			"address":    "顺昌县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	350722: {
		{
			"address":    "浦城县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	350723: {
		{
			"address":    "光泽县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	350724: {
		{
			"address":    "松溪县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	350725: {
		{
			"address":    "政和县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	350781: {
		{
			"address":    "邵武市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	350782: {
		{
			"address":    "武夷山市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	350783: {
		{
			"address":    "建瓯市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	350784: {
		{
			"address":    "建阳市",
			"start_year": "1995",
			"end_year":   "2013",
		},
	},
	350800: {
		{
			"address":    "龙岩市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	350802: {
		{
			"address":    "新罗区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	350803: {
		{
			"address":    "永定区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	350821: {
		{
			"address":    "长汀县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	350822: {
		{
			"address":    "永定县",
			"start_year": "1996",
			"end_year":   "2013",
		},
	},
	350823: {
		{
			"address":    "上杭县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	350824: {
		{
			"address":    "武平县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	350825: {
		{
			"address":    "连城县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	350881: {
		{
			"address":    "漳平市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	350900: {
		{
			"address":    "宁德市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	350902: {
		{
			"address":    "蕉城区",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	350921: {
		{
			"address":    "霞浦县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	350922: {
		{
			"address":    "古田县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	350923: {
		{
			"address":    "屏南县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	350924: {
		{
			"address":    "寿宁县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	350925: {
		{
			"address":    "周宁县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	350926: {
		{
			"address":    "柘荣县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	350981: {
		{
			"address":    "福安市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	350982: {
		{
			"address":    "福鼎市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	352100: {
		{
			"address":    "建阳地区",
			"start_year": "1982",
			"end_year":   "1987",
		},
		{
			"address":    "南平地区",
			"start_year": "1988",
			"end_year":   "1993",
		},
	},
	352101: {
		{
			"address":    "南平市",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	352102: {
		{
			"address":    "邵武市",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	352103: {
		{
			"address":    "武夷山市",
			"start_year": "1989",
			"end_year":   "1993",
		},
	},
	352104: {
		{
			"address":    "建瓯市",
			"start_year": "1992",
			"end_year":   "1993",
		},
	},
	352121: {
		{
			"address":    "顺昌县",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	352122: {
		{
			"address":    "建阳县",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	352123: {
		{
			"address":    "建瓯县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	352124: {
		{
			"address":    "浦城县",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	352125: {
		{
			"address":    "邵武县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352126: {
		{
			"address":    "崇安县",
			"start_year": "1982",
			"end_year":   "1988",
		},
	},
	352127: {
		{
			"address":    "光泽县",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	352128: {
		{
			"address":    "松溪县",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	352129: {
		{
			"address":    "政和县",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	352200: {
		{
			"address":    "建阳地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宁德地区",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	352201: {
		{
			"address":    "南平市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宁德市",
			"start_year": "1988",
			"end_year":   "1998",
		},
	},
	352202: {
		{
			"address":    "福安市",
			"start_year": "1989",
			"end_year":   "1998",
		},
	},
	352203: {
		{
			"address":    "福鼎市",
			"start_year": "1995",
			"end_year":   "1998",
		},
	},
	352221: {
		{
			"address":    "顺昌县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宁德县",
			"start_year": "1982",
			"end_year":   "1987",
		},
	},
	352222: {
		{
			"address":    "建阳县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "连江县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352223: {
		{
			"address":    "建瓯县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "罗源县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352224: {
		{
			"address":    "浦城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "福鼎县",
			"start_year": "1982",
			"end_year":   "1994",
		},
	},
	352225: {
		{
			"address":    "邵武县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "霞浦县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	352226: {
		{
			"address":    "崇安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "福安县",
			"start_year": "1982",
			"end_year":   "1988",
		},
	},
	352227: {
		{
			"address":    "光泽县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "古田县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	352228: {
		{
			"address":    "松溪县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "屏南县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	352229: {
		{
			"address":    "政和县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "寿宁县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	352230: {
		{
			"address":    "周宁县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	352231: {
		{
			"address":    "柘荣县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	352300: {
		{
			"address":    "宁德地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "莆田地区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352321: {
		{
			"address":    "宁德县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "闽清县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352322: {
		{
			"address":    "连江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "永泰县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352323: {
		{
			"address":    "罗源县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "长乐县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352324: {
		{
			"address":    "福鼎县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "福清县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352325: {
		{
			"address":    "霞浦县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "平潭县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352326: {
		{
			"address":    "福安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "莆田县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352327: {
		{
			"address":    "古田县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "仙游县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352328: {
		{
			"address":    "屏南县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352329: {
		{
			"address":    "寿宁县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352330: {
		{
			"address":    "周宁县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352331: {
		{
			"address":    "柘荣县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352400: {
		{
			"address":    "莆田地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "晋江地区",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352401: {
		{
			"address":    "泉州市",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352421: {
		{
			"address":    "闽清县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "惠安县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352422: {
		{
			"address":    "永泰县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "晋江县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352423: {
		{
			"address":    "长乐县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "南安县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352424: {
		{
			"address":    "福清县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "安溪县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352425: {
		{
			"address":    "平潭县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "永春县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352426: {
		{
			"address":    "莆田县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "德化县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352427: {
		{
			"address":    "仙游县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "金门县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352500: {
		{
			"address":    "晋江地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "龙溪地区",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352501: {
		{
			"address":    "泉州市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "漳州市",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352521: {
		{
			"address":    "惠安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "龙海县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352522: {
		{
			"address":    "晋江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "云霄县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352523: {
		{
			"address":    "南安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "漳浦县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352524: {
		{
			"address":    "安溪县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "诏安县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352525: {
		{
			"address":    "永春县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "长泰县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352526: {
		{
			"address":    "德化县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "东山县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352527: {
		{
			"address":    "金门县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "南靖县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352528: {
		{
			"address":    "平和县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352529: {
		{
			"address":    "华安县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	352600: {
		{
			"address":    "龙溪地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "龙岩地区",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	352601: {
		{
			"address":    "漳州市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "龙岩市",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	352602: {
		{
			"address":    "漳平市",
			"start_year": "1990",
			"end_year":   "1995",
		},
	},
	352621: {
		{
			"address":    "龙海县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352622: {
		{
			"address":    "云霄县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "长汀县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	352623: {
		{
			"address":    "漳浦县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "永定县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	352624: {
		{
			"address":    "诏安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "上杭县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	352625: {
		{
			"address":    "长泰县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "武平县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	352626: {
		{
			"address":    "东山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "漳平县",
			"start_year": "1982",
			"end_year":   "1989",
		},
	},
	352627: {
		{
			"address":    "南靖县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "连城县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	352628: {
		{
			"address":    "平和县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352629: {
		{
			"address":    "华安县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352700: {
		{
			"address":    "龙岩地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "三明地区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352701: {
		{
			"address":    "龙岩市",
			"start_year": "1981",
			"end_year":   "1981",
		},
		{
			"address":    "三明市",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352721: {
		{
			"address":    "龙岩县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "明溪县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352722: {
		{
			"address":    "长汀县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "永安县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352723: {
		{
			"address":    "永定县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "清流县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352724: {
		{
			"address":    "上杭县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宁化县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352725: {
		{
			"address":    "武平县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "大田县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352726: {
		{
			"address":    "漳平县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "尤溪县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352727: {
		{
			"address":    "连城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "沙县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352728: {
		{
			"address":    "将乐县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352729: {
		{
			"address":    "泰宁县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352730: {
		{
			"address":    "建宁县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	352800: {
		{
			"address":    "三明地区",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352801: {
		{
			"address":    "三明市",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352821: {
		{
			"address":    "明溪县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352822: {
		{
			"address":    "永安县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352823: {
		{
			"address":    "清流县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352824: {
		{
			"address":    "宁化县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352825: {
		{
			"address":    "大田县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352826: {
		{
			"address":    "尤溪县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352827: {
		{
			"address":    "沙县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352828: {
		{
			"address":    "将乐县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352829: {
		{
			"address":    "泰宁县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	352830: {
		{
			"address":    "建宁县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	359001: {
		{
			"address":    "永安市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	359002: {
		{
			"address":    "石狮市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	359003: {
		{
			"address":    "福清市",
			"start_year": "1990",
			"end_year":   "1994",
		},
	},
	359004: {
		{
			"address":    "晋江市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	359005: {
		{
			"address":    "南安市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	359006: {
		{
			"address":    "龙海市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	359007: {
		{
			"address":    "邵武市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	359008: {
		{
			"address":    "武夷山市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	359009: {
		{
			"address":    "建瓯市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	359010: {
		{
			"address":    "建阳市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	359011: {
		{
			"address":    "长乐市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	360000: {
		{
			"address":    "江西省",
			"start_year": "",
			"end_year":   "",
		},
	},
	360100: {
		{
			"address":    "南昌市",
			"start_year": "",
			"end_year":   "",
		},
	},
	360102: {
		{
			"address":    "东湖区",
			"start_year": "",
			"end_year":   "",
		},
	},
	360103: {
		{
			"address":    "西湖区",
			"start_year": "",
			"end_year":   "",
		},
	},
	360104: {
		{
			"address":    "青云谱区",
			"start_year": "",
			"end_year":   "",
		},
	},
	360105: {
		{
			"address":    "湾里区",
			"start_year": "",
			"end_year":   "2018",
		},
	},
	360111: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "2001",
		},
		{
			"address":    "青山湖区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	360112: {
		{
			"address":    "新建区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	360113: {
		{
			"address":    "红谷滩区",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	360121: {
		{
			"address":    "新建县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "南昌县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	360122: {
		{
			"address":    "南昌县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "新建县",
			"start_year": "1982",
			"end_year":   "2014",
		},
	},
	360123: {
		{
			"address":    "安义县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	360124: {
		{
			"address":    "进贤县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	360200: {
		{
			"address":    "景德镇市",
			"start_year": "",
			"end_year":   "",
		},
	},
	360202: {
		{
			"address":    "昌江区",
			"start_year": "",
			"end_year":   "",
		},
	},
	360203: {
		{
			"address":    "珠山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	360211: {
		{
			"address":    "鹅湖区",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	360212: {
		{
			"address":    "蛟潭区",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	360221: {
		{
			"address":    "乐平县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	360222: {
		{
			"address":    "浮梁县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	360281: {
		{
			"address":    "乐平市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	360300: {
		{
			"address":    "萍乡市",
			"start_year": "",
			"end_year":   "",
		},
	},
	360302: {
		{
			"address":    "城关区",
			"start_year": "",
			"end_year":   "1992",
		},
		{
			"address":    "安源区",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	360311: {
		{
			"address":    "上栗区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	360312: {
		{
			"address":    "芦溪区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	360313: {
		{
			"address":    "湘东区",
			"start_year": "",
			"end_year":   "",
		},
	},
	360321: {
		{
			"address":    "莲花县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	360322: {
		{
			"address":    "上栗县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	360323: {
		{
			"address":    "芦溪县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	360400: {
		{
			"address":    "九江市",
			"start_year": "",
			"end_year":   "",
		},
	},
	360402: {
		{
			"address":    "庐山区",
			"start_year": "",
			"end_year":   "2015",
		},
		{
			"address":    "濂溪区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	360403: {
		{
			"address":    "浔阳区",
			"start_year": "",
			"end_year":   "",
		},
	},
	360404: {
		{
			"address":    "柴桑区",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	360411: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	360421: {
		{
			"address":    "九江县",
			"start_year": "1983",
			"end_year":   "2016",
		},
	},
	360422: {
		{
			"address":    "瑞昌县",
			"start_year": "1983",
			"end_year":   "1988",
		},
	},
	360423: {
		{
			"address":    "武宁县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	360424: {
		{
			"address":    "修水县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	360425: {
		{
			"address":    "永修县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	360426: {
		{
			"address":    "德安县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	360427: {
		{
			"address":    "星子县",
			"start_year": "1983",
			"end_year":   "2015",
		},
	},
	360428: {
		{
			"address":    "都昌县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	360429: {
		{
			"address":    "湖口县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	360430: {
		{
			"address":    "彭泽县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	360481: {
		{
			"address":    "瑞昌市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	360482: {
		{
			"address":    "共青城市",
			"start_year": "2010",
			"end_year":   "",
		},
	},
	360483: {
		{
			"address":    "庐山市",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	360500: {
		{
			"address":    "新余市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	360502: {
		{
			"address":    "渝水区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	360521: {
		{
			"address":    "井冈山",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "分宜县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	360600: {
		{
			"address":    "鹰潭市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	360602: {
		{
			"address":    "月湖区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	360603: {
		{
			"address":    "余江区",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	360621: {
		{
			"address":    "贵溪县",
			"start_year": "1983",
			"end_year":   "1995",
		},
	},
	360622: {
		{
			"address":    "余江县",
			"start_year": "1983",
			"end_year":   "2017",
		},
	},
	360681: {
		{
			"address":    "贵溪市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	360700: {
		{
			"address":    "赣州市",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360702: {
		{
			"address":    "章贡区",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360703: {
		{
			"address":    "南康区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	360704: {
		{
			"address":    "赣县区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	360721: {
		{
			"address":    "赣县",
			"start_year": "1998",
			"end_year":   "2015",
		},
	},
	360722: {
		{
			"address":    "信丰县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360723: {
		{
			"address":    "大余县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360724: {
		{
			"address":    "上犹县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360725: {
		{
			"address":    "崇义县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360726: {
		{
			"address":    "安远县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360727: {
		{
			"address":    "龙南县",
			"start_year": "1998",
			"end_year":   "2019",
		},
	},
	360728: {
		{
			"address":    "定南县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360729: {
		{
			"address":    "全南县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360730: {
		{
			"address":    "宁都县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360731: {
		{
			"address":    "于都县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360732: {
		{
			"address":    "兴国县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360733: {
		{
			"address":    "会昌县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360734: {
		{
			"address":    "寻乌县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360735: {
		{
			"address":    "石城县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360781: {
		{
			"address":    "瑞金市",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	360782: {
		{
			"address":    "南康市",
			"start_year": "1998",
			"end_year":   "2012",
		},
	},
	360783: {
		{
			"address":    "龙南市",
			"start_year": "2020",
			"end_year":   "",
		},
	},
	360800: {
		{
			"address":    "吉安市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360802: {
		{
			"address":    "吉州区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360803: {
		{
			"address":    "青原区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360821: {
		{
			"address":    "吉安县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360822: {
		{
			"address":    "吉水县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360823: {
		{
			"address":    "峡江县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360824: {
		{
			"address":    "新干县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360825: {
		{
			"address":    "永丰县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360826: {
		{
			"address":    "泰和县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360827: {
		{
			"address":    "遂川县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360828: {
		{
			"address":    "万安县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360829: {
		{
			"address":    "安福县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360830: {
		{
			"address":    "永新县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360881: {
		{
			"address":    "井冈山市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360900: {
		{
			"address":    "宜春市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360902: {
		{
			"address":    "袁州区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360921: {
		{
			"address":    "奉新县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360922: {
		{
			"address":    "万载县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360923: {
		{
			"address":    "上高县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360924: {
		{
			"address":    "宜丰县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360925: {
		{
			"address":    "靖安县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360926: {
		{
			"address":    "铜鼓县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360981: {
		{
			"address":    "丰城市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360982: {
		{
			"address":    "樟树市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	360983: {
		{
			"address":    "高安市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361000: {
		{
			"address":    "抚州市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361002: {
		{
			"address":    "临川区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361003: {
		{
			"address":    "东乡区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	361021: {
		{
			"address":    "南城县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361022: {
		{
			"address":    "黎川县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361023: {
		{
			"address":    "南丰县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361024: {
		{
			"address":    "崇仁县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361025: {
		{
			"address":    "乐安县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361026: {
		{
			"address":    "宜黄县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361027: {
		{
			"address":    "金溪县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361028: {
		{
			"address":    "资溪县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361029: {
		{
			"address":    "东乡县",
			"start_year": "2000",
			"end_year":   "2015",
		},
	},
	361030: {
		{
			"address":    "广昌县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361100: {
		{
			"address":    "上饶市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361102: {
		{
			"address":    "信州区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361103: {
		{
			"address":    "广丰区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	361104: {
		{
			"address":    "广信区",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	361121: {
		{
			"address":    "上饶县",
			"start_year": "2000",
			"end_year":   "2018",
		},
	},
	361122: {
		{
			"address":    "广丰县",
			"start_year": "2000",
			"end_year":   "2014",
		},
	},
	361123: {
		{
			"address":    "玉山县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361124: {
		{
			"address":    "铅山县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361125: {
		{
			"address":    "横峰县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361126: {
		{
			"address":    "弋阳县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361127: {
		{
			"address":    "余干县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361128: {
		{
			"address":    "波阳县",
			"start_year": "2000",
			"end_year":   "2002",
		},
		{
			"address":    "鄱阳县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	361129: {
		{
			"address":    "万年县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361130: {
		{
			"address":    "婺源县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	361181: {
		{
			"address":    "德兴市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	362100: {
		{
			"address":    "九江地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "赣州地区",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362101: {
		{
			"address":    "赣州市",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362102: {
		{
			"address":    "瑞金市",
			"start_year": "1994",
			"end_year":   "1997",
		},
	},
	362103: {
		{
			"address":    "南康市",
			"start_year": "1995",
			"end_year":   "1997",
		},
	},
	362121: {
		{
			"address":    "九江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "赣县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362122: {
		{
			"address":    "永修县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "南康县",
			"start_year": "1982",
			"end_year":   "1994",
		},
	},
	362123: {
		{
			"address":    "彭泽县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "信丰县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362124: {
		{
			"address":    "德安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "大余县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362125: {
		{
			"address":    "湖口县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "上犹县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362126: {
		{
			"address":    "瑞昌县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "崇义县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362127: {
		{
			"address":    "都昌县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "安远县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362128: {
		{
			"address":    "武宁县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "龙南县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362129: {
		{
			"address":    "星子县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "定南县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362130: {
		{
			"address":    "修水县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "全南县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362131: {
		{
			"address":    "宁都县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362132: {
		{
			"address":    "于都县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362133: {
		{
			"address":    "兴国县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362134: {
		{
			"address":    "瑞金县",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	362135: {
		{
			"address":    "会昌县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362136: {
		{
			"address":    "寻乌县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362137: {
		{
			"address":    "石城县",
			"start_year": "1982",
			"end_year":   "1997",
		},
	},
	362138: {
		{
			"address":    "广昌县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362200: {
		{
			"address":    "上饶地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宜春地区",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362201: {
		{
			"address":    "上饶市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宜春市",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362202: {
		{
			"address":    "鹰潭市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "丰城市",
			"start_year": "1988",
			"end_year":   "1999",
		},
	},
	362203: {
		{
			"address":    "樟树市",
			"start_year": "1988",
			"end_year":   "1999",
		},
	},
	362204: {
		{
			"address":    "高安市",
			"start_year": "1993",
			"end_year":   "1999",
		},
	},
	362221: {
		{
			"address":    "上饶县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "丰城县",
			"start_year": "1982",
			"end_year":   "1987",
		},
	},
	362222: {
		{
			"address":    "贵溪县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "高安县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	362223: {
		{
			"address":    "婺源县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "清江县",
			"start_year": "1982",
			"end_year":   "1987",
		},
	},
	362224: {
		{
			"address":    "余江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "新余县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362225: {
		{
			"address":    "德兴县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宜春县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	362226: {
		{
			"address":    "万年县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "奉新县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362227: {
		{
			"address":    "玉山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "万载县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362228: {
		{
			"address":    "乐平县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "上高县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362229: {
		{
			"address":    "广丰县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宜丰县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362230: {
		{
			"address":    "波阳县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "分宜县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362231: {
		{
			"address":    "铅山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "安义县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362232: {
		{
			"address":    "余干县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "靖安县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362233: {
		{
			"address":    "横峰县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "铜鼓县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362234: {
		{
			"address":    "弋阳县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	362300: {
		{
			"address":    "宜春地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "上饶地区",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362301: {
		{
			"address":    "宜春市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "上饶市",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362302: {
		{
			"address":    "鹰潭市",
			"start_year": "1982",
			"end_year":   "1982",
		},
		{
			"address":    "德兴市",
			"start_year": "1990",
			"end_year":   "1999",
		},
	},
	362321: {
		{
			"address":    "宜春县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "上饶县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362322: {
		{
			"address":    "高安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "广丰县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362323: {
		{
			"address":    "万载县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "玉山县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362324: {
		{
			"address":    "丰城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "铅山县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362325: {
		{
			"address":    "铜鼓县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "横峰县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362326: {
		{
			"address":    "清江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "弋阳县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362327: {
		{
			"address":    "宜丰县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "贵溪县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362328: {
		{
			"address":    "新余县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "余江县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362329: {
		{
			"address":    "上高县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "余干县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362330: {
		{
			"address":    "分宜县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "波阳县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362331: {
		{
			"address":    "安义县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "万年县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362332: {
		{
			"address":    "靖安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "乐平县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362333: {
		{
			"address":    "奉新县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "德兴县",
			"start_year": "1982",
			"end_year":   "1989",
		},
	},
	362334: {
		{
			"address":    "婺源县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362400: {
		{
			"address":    "抚州地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "吉安地区",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362401: {
		{
			"address":    "抚州市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "吉安市",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362402: {
		{
			"address":    "井冈山市",
			"start_year": "1984",
			"end_year":   "1999",
		},
	},
	362421: {
		{
			"address":    "临川县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "吉安县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362422: {
		{
			"address":    "宜黄县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "吉水县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362423: {
		{
			"address":    "金溪县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "峡江县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362424: {
		{
			"address":    "崇仁县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "新干县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362425: {
		{
			"address":    "资溪县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "永丰县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362426: {
		{
			"address":    "乐安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "泰和县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362427: {
		{
			"address":    "黎川县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "遂川县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362428: {
		{
			"address":    "东乡县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "万安县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362429: {
		{
			"address":    "南丰县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "安福县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362430: {
		{
			"address":    "进贤县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "永新县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362431: {
		{
			"address":    "南城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "莲花县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	362432: {
		{
			"address":    "宁冈县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362433: {
		{
			"address":    "井冈山县",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	362500: {
		{
			"address":    "吉安地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "抚州地区",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362501: {
		{
			"address":    "吉安市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "抚州市",
			"start_year": "1982",
			"end_year":   "1986",
		},
		{
			"address":    "临川市",
			"start_year": "1987",
			"end_year":   "1999",
		},
	},
	362521: {
		{
			"address":    "吉安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "临川县",
			"start_year": "1982",
			"end_year":   "1986",
		},
	},
	362522: {
		{
			"address":    "万安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "南城县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362523: {
		{
			"address":    "新干县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "黎川县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362524: {
		{
			"address":    "遂川县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "南丰县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362525: {
		{
			"address":    "峡江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "崇仁县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362526: {
		{
			"address":    "宁冈县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "乐安县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362527: {
		{
			"address":    "吉水县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宜黄县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362528: {
		{
			"address":    "永新县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "金溪县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362529: {
		{
			"address":    "永丰县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "资溪县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362530: {
		{
			"address":    "莲花县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "进贤县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362531: {
		{
			"address":    "泰和县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "东乡县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	362532: {
		{
			"address":    "安福县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "广昌县",
			"start_year": "1983",
			"end_year":   "1999",
		},
	},
	362533: {
		{
			"address":    "井冈山县",
			"start_year": "1981",
			"end_year":   "1981",
		},
	},
	362600: {
		{
			"address":    "赣州地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "九江地区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362601: {
		{
			"address":    "赣州市",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	362621: {
		{
			"address":    "广昌县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "九江县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362622: {
		{
			"address":    "定南县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "瑞昌县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362623: {
		{
			"address":    "石城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "武宁县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362624: {
		{
			"address":    "龙南县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "修水县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362625: {
		{
			"address":    "宁都县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "永修县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362626: {
		{
			"address":    "全南县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "德安县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362627: {
		{
			"address":    "兴国县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "星子县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362628: {
		{
			"address":    "信丰县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "都昌县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362629: {
		{
			"address":    "于都县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "湖口县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362630: {
		{
			"address":    "赣县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "彭泽县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	362631: {
		{
			"address":    "瑞金县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	362632: {
		{
			"address":    "南康县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	362633: {
		{
			"address":    "会昌县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	362634: {
		{
			"address":    "上犹县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	362635: {
		{
			"address":    "安远县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	362636: {
		{
			"address":    "崇义县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	362637: {
		{
			"address":    "寻乌县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	362638: {
		{
			"address":    "大余县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	369001: {
		{
			"address":    "瑞昌市",
			"start_year": "1989",
			"end_year":   "1994",
		},
	},
	369002: {
		{
			"address":    "乐平市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	370000: {
		{
			"address":    "山东省",
			"start_year": "",
			"end_year":   "",
		},
	},
	370100: {
		{
			"address":    "济南市",
			"start_year": "",
			"end_year":   "",
		},
	},
	370102: {
		{
			"address":    "历下区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370103: {
		{
			"address":    "市中区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370104: {
		{
			"address":    "槐荫区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370105: {
		{
			"address":    "天桥区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370111: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	370112: {
		{
			"address":    "历城区",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	370113: {
		{
			"address":    "长清区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	370114: {
		{
			"address":    "章丘区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	370115: {
		{
			"address":    "济阳区",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	370116: {
		{
			"address":    "莱芜区",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	370117: {
		{
			"address":    "钢城区",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	370121: {
		{
			"address":    "历城县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	370122: {
		{
			"address":    "章丘县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	370123: {
		{
			"address":    "长清县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	370124: {
		{
			"address":    "平阴县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	370125: {
		{
			"address":    "济阳县",
			"start_year": "1989",
			"end_year":   "2017",
		},
	},
	370126: {
		{
			"address":    "商河县",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	370181: {
		{
			"address":    "章丘市",
			"start_year": "1995",
			"end_year":   "2015",
		},
	},
	370200: {
		{
			"address":    "青岛市",
			"start_year": "",
			"end_year":   "",
		},
	},
	370202: {
		{
			"address":    "市南区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370203: {
		{
			"address":    "市北区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370204: {
		{
			"address":    "台东区",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	370205: {
		{
			"address":    "四方区",
			"start_year": "",
			"end_year":   "2011",
		},
	},
	370206: {
		{
			"address":    "沧口区",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	370211: {
		{
			"address":    "黄岛区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370212: {
		{
			"address":    "崂山区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	370213: {
		{
			"address":    "李沧区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	370214: {
		{
			"address":    "城阳区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	370215: {
		{
			"address":    "即墨区",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	370221: {
		{
			"address":    "崂山县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	370222: {
		{
			"address":    "即墨县",
			"start_year": "",
			"end_year":   "1988",
		},
	},
	370223: {
		{
			"address":    "胶南县",
			"start_year": "",
			"end_year":   "1989",
		},
	},
	370224: {
		{
			"address":    "胶县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	370225: {
		{
			"address":    "莱西县",
			"start_year": "1983",
			"end_year":   "1989",
		},
	},
	370226: {
		{
			"address":    "平度县",
			"start_year": "1983",
			"end_year":   "1988",
		},
	},
	370281: {
		{
			"address":    "胶州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370282: {
		{
			"address":    "即墨市",
			"start_year": "1995",
			"end_year":   "2016",
		},
	},
	370283: {
		{
			"address":    "平度市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370284: {
		{
			"address":    "胶南市",
			"start_year": "1995",
			"end_year":   "2011",
		},
	},
	370285: {
		{
			"address":    "莱西市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370300: {
		{
			"address":    "淄博市",
			"start_year": "",
			"end_year":   "",
		},
	},
	370302: {
		{
			"address":    "淄川区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370303: {
		{
			"address":    "张店区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370304: {
		{
			"address":    "博山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370305: {
		{
			"address":    "临淄区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370306: {
		{
			"address":    "周村区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370321: {
		{
			"address":    "桓台县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370322: {
		{
			"address":    "高青县",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	370323: {
		{
			"address":    "沂源县",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	370400: {
		{
			"address":    "枣庄市",
			"start_year": "",
			"end_year":   "",
		},
	},
	370402: {
		{
			"address":    "市中区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370403: {
		{
			"address":    "薛城区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370404: {
		{
			"address":    "峄城区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370405: {
		{
			"address":    "台儿庄区",
			"start_year": "",
			"end_year":   "",
		},
	},
	370406: {
		{
			"address":    "齐村区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "山亭区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370421: {
		{
			"address":    "滕县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	370481: {
		{
			"address":    "滕州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370500: {
		{
			"address":    "东营市",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	370502: {
		{
			"address":    "东营区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	370503: {
		{
			"address":    "河口区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	370504: {
		{
			"address":    "牛庄区",
			"start_year": "1982",
			"end_year":   "1986",
		},
	},
	370505: {
		{
			"address":    "垦利区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	370521: {
		{
			"address":    "垦利县",
			"start_year": "1982",
			"end_year":   "2015",
		},
	},
	370522: {
		{
			"address":    "利津县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	370523: {
		{
			"address":    "广饶县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370600: {
		{
			"address":    "烟台市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370602: {
		{
			"address":    "芝罘区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370611: {
		{
			"address":    "福山区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370612: {
		{
			"address":    "牟平区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	370613: {
		{
			"address":    "莱山区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	370620: {
		{
			"address":    "威海市",
			"start_year": "1985",
			"end_year":   "1985",
		},
	},
	370622: {
		{
			"address":    "蓬莱县",
			"start_year": "1983",
			"end_year":   "1990",
		},
	},
	370623: {
		{
			"address":    "黄县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	370624: {
		{
			"address":    "招远县",
			"start_year": "1983",
			"end_year":   "1990",
		},
	},
	370625: {
		{
			"address":    "掖县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	370627: {
		{
			"address":    "莱阳县",
			"start_year": "1983",
			"end_year":   "1986",
		},
	},
	370628: {
		{
			"address":    "栖霞县",
			"start_year": "1983",
			"end_year":   "1994",
		},
	},
	370629: {
		{
			"address":    "海阳县",
			"start_year": "1983",
			"end_year":   "1995",
		},
	},
	370630: {
		{
			"address":    "乳山县",
			"start_year": "1983",
			"end_year":   "1986",
		},
	},
	370631: {
		{
			"address":    "牟平县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	370632: {
		{
			"address":    "文登县",
			"start_year": "1983",
			"end_year":   "1986",
		},
	},
	370633: {
		{
			"address":    "荣成县",
			"start_year": "1983",
			"end_year":   "1986",
		},
	},
	370634: {
		{
			"address":    "长岛县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370681: {
		{
			"address":    "龙口市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370682: {
		{
			"address":    "莱阳市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370683: {
		{
			"address":    "莱州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370684: {
		{
			"address":    "蓬莱市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370685: {
		{
			"address":    "招远市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370686: {
		{
			"address":    "栖霞市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370687: {
		{
			"address":    "海阳市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	370700: {
		{
			"address":    "潍坊市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370702: {
		{
			"address":    "潍城区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370703: {
		{
			"address":    "寒亭区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370704: {
		{
			"address":    "坊子区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370705: {
		{
			"address":    "奎文区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	370721: {
		{
			"address":    "益都县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	370722: {
		{
			"address":    "安丘县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	370723: {
		{
			"address":    "寿光县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	370724: {
		{
			"address":    "临朐县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370725: {
		{
			"address":    "昌乐县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370726: {
		{
			"address":    "昌邑县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	370727: {
		{
			"address":    "高密县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	370728: {
		{
			"address":    "诸城县",
			"start_year": "1983",
			"end_year":   "1986",
		},
	},
	370729: {
		{
			"address":    "五莲县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	370781: {
		{
			"address":    "青州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370782: {
		{
			"address":    "诸城市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370783: {
		{
			"address":    "寿光市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370784: {
		{
			"address":    "安丘市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370785: {
		{
			"address":    "高密市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370786: {
		{
			"address":    "昌邑市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370800: {
		{
			"address":    "济宁市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370802: {
		{
			"address":    "市中区",
			"start_year": "1983",
			"end_year":   "2012",
		},
	},
	370811: {
		{
			"address":    "市郊区",
			"start_year": "1983",
			"end_year":   "1992",
		},
		{
			"address":    "任城区",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	370812: {
		{
			"address":    "兖州区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	370822: {
		{
			"address":    "兖州县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	370823: {
		{
			"address":    "曲阜县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	370825: {
		{
			"address":    "邹县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	370826: {
		{
			"address":    "微山县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370827: {
		{
			"address":    "鱼台县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370828: {
		{
			"address":    "金乡县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370829: {
		{
			"address":    "嘉祥县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	370830: {
		{
			"address":    "汶上县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	370831: {
		{
			"address":    "泗水县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	370832: {
		{
			"address":    "梁山县",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	370881: {
		{
			"address":    "曲阜市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370882: {
		{
			"address":    "兖州市",
			"start_year": "1995",
			"end_year":   "2012",
		},
	},
	370883: {
		{
			"address":    "邹城市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	370900: {
		{
			"address":    "泰安市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	370901: {
		{
			"address":    "威海市",
			"start_year": "1983",
			"end_year":   "1984",
		},
	},
	370902: {
		{
			"address":    "泰山区",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	370911: {
		{
			"address":    "郊区",
			"start_year": "1985",
			"end_year":   "1999",
		},
		{
			"address":    "岱岳区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	370921: {
		{
			"address":    "宁阳县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	370922: {
		{
			"address":    "肥城县",
			"start_year": "1985",
			"end_year":   "1991",
		},
	},
	370923: {
		{
			"address":    "东平县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	370981: {
		{
			"address":    "莱芜市",
			"start_year": "1985",
			"end_year":   "1985",
		},
	},
	370982: {
		{
			"address":    "新泰市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	370983: {
		{
			"address":    "肥城市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	371000: {
		{
			"address":    "威海市",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	371002: {
		{
			"address":    "环翠区",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	371003: {
		{
			"address":    "文登区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	371021: {
		{
			"address":    "乳山县",
			"start_year": "1987",
			"end_year":   "1992",
		},
	},
	371022: {
		{
			"address":    "文登县",
			"start_year": "1987",
			"end_year":   "1987",
		},
	},
	371023: {
		{
			"address":    "荣成县",
			"start_year": "1987",
			"end_year":   "1987",
		},
	},
	371081: {
		{
			"address":    "文登市",
			"start_year": "1995",
			"end_year":   "2013",
		},
	},
	371082: {
		{
			"address":    "荣成市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	371083: {
		{
			"address":    "乳山市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	371100: {
		{
			"address":    "日照市",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	371102: {
		{
			"address":    "东港区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	371103: {
		{
			"address":    "岚山区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	371121: {
		{
			"address":    "五莲县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	371122: {
		{
			"address":    "莒县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	371200: {
		{
			"address":    "莱芜市",
			"start_year": "1992",
			"end_year":   "2017",
		},
	},
	371202: {
		{
			"address":    "莱城区",
			"start_year": "1992",
			"end_year":   "2017",
		},
	},
	371203: {
		{
			"address":    "钢城区",
			"start_year": "1992",
			"end_year":   "2017",
		},
	},
	371300: {
		{
			"address":    "临沂市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371302: {
		{
			"address":    "兰山区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371311: {
		{
			"address":    "罗庄区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371312: {
		{
			"address":    "河东区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371321: {
		{
			"address":    "沂南县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371322: {
		{
			"address":    "郯城县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371323: {
		{
			"address":    "沂水县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371324: {
		{
			"address":    "苍山县",
			"start_year": "1994",
			"end_year":   "2013",
		},
		{
			"address":    "兰陵县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	371325: {
		{
			"address":    "费县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371326: {
		{
			"address":    "平邑县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371327: {
		{
			"address":    "莒南县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371328: {
		{
			"address":    "蒙阴县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371329: {
		{
			"address":    "临沭县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371400: {
		{
			"address":    "德州市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371402: {
		{
			"address":    "德城区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371403: {
		{
			"address":    "陵城区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	371421: {
		{
			"address":    "陵县",
			"start_year": "1994",
			"end_year":   "2013",
		},
	},
	371422: {
		{
			"address":    "宁津县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371423: {
		{
			"address":    "庆云县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371424: {
		{
			"address":    "临邑县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371425: {
		{
			"address":    "齐河县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371426: {
		{
			"address":    "平原县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371427: {
		{
			"address":    "夏津县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371428: {
		{
			"address":    "武城县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	371481: {
		{
			"address":    "乐陵市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	371482: {
		{
			"address":    "禹城市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	371500: {
		{
			"address":    "聊城市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	371502: {
		{
			"address":    "东昌府区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	371503: {
		{
			"address":    "茌平区",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	371521: {
		{
			"address":    "阳谷县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	371522: {
		{
			"address":    "莘县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	371523: {
		{
			"address":    "茌平县",
			"start_year": "1997",
			"end_year":   "2018",
		},
	},
	371524: {
		{
			"address":    "东阿县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	371525: {
		{
			"address":    "冠县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	371526: {
		{
			"address":    "高唐县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	371581: {
		{
			"address":    "临清市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	371600: {
		{
			"address":    "滨州市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	371602: {
		{
			"address":    "滨城区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	371603: {
		{
			"address":    "沾化区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	371621: {
		{
			"address":    "惠民县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	371622: {
		{
			"address":    "阳信县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	371623: {
		{
			"address":    "无棣县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	371624: {
		{
			"address":    "沾化县",
			"start_year": "2000",
			"end_year":   "2013",
		},
	},
	371625: {
		{
			"address":    "博兴县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	371626: {
		{
			"address":    "邹平县",
			"start_year": "2000",
			"end_year":   "2017",
		},
	},
	371681: {
		{
			"address":    "邹平市",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	371700: {
		{
			"address":    "菏泽市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	371702: {
		{
			"address":    "牡丹区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	371703: {
		{
			"address":    "定陶区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	371721: {
		{
			"address":    "曹县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	371722: {
		{
			"address":    "单县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	371723: {
		{
			"address":    "成武县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	371724: {
		{
			"address":    "巨野县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	371725: {
		{
			"address":    "郓城县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	371726: {
		{
			"address":    "鄄城县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	371727: {
		{
			"address":    "定陶县",
			"start_year": "2000",
			"end_year":   "2015",
		},
	},
	371728: {
		{
			"address":    "东明县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	372100: {
		{
			"address":    "烟台地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372101: {
		{
			"address":    "烟台市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372102: {
		{
			"address":    "威海市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372121: {
		{
			"address":    "福山县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372122: {
		{
			"address":    "蓬莱县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372123: {
		{
			"address":    "黄县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372124: {
		{
			"address":    "招远县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372125: {
		{
			"address":    "掖县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372126: {
		{
			"address":    "莱西县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372127: {
		{
			"address":    "莱阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372128: {
		{
			"address":    "栖霞县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372129: {
		{
			"address":    "海阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372130: {
		{
			"address":    "乳山县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372131: {
		{
			"address":    "牟平县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372132: {
		{
			"address":    "文登县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372133: {
		{
			"address":    "荣成县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372134: {
		{
			"address":    "长岛县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372200: {
		{
			"address":    "昌潍地区",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "潍坊地区",
			"start_year": "1981",
			"end_year":   "1982",
		},
	},
	372201: {
		{
			"address":    "潍坊市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372221: {
		{
			"address":    "益都县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372222: {
		{
			"address":    "安丘县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372223: {
		{
			"address":    "寿光县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372224: {
		{
			"address":    "临朐县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372225: {
		{
			"address":    "昌乐县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372226: {
		{
			"address":    "昌邑县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372227: {
		{
			"address":    "高密县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372228: {
		{
			"address":    "诸城县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372229: {
		{
			"address":    "五莲县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372230: {
		{
			"address":    "平度县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372231: {
		{
			"address":    "潍县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372300: {
		{
			"address":    "惠民地区",
			"start_year": "",
			"end_year":   "1991",
		},
		{
			"address":    "滨州地区",
			"start_year": "1992",
			"end_year":   "1999",
		},
	},
	372301: {
		{
			"address":    "滨州市",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	372321: {
		{
			"address":    "惠民县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	372322: {
		{
			"address":    "滨县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	372323: {
		{
			"address":    "阳信县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	372324: {
		{
			"address":    "无棣县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	372325: {
		{
			"address":    "沾化县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	372326: {
		{
			"address":    "利津县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	372327: {
		{
			"address":    "广饶县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372328: {
		{
			"address":    "博兴县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	372329: {
		{
			"address":    "桓台县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372330: {
		{
			"address":    "邹平县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	372331: {
		{
			"address":    "高青县",
			"start_year": "",
			"end_year":   "1988",
		},
	},
	372332: {
		{
			"address":    "垦利县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	372400: {
		{
			"address":    "德州地区",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372401: {
		{
			"address":    "德州市",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372402: {
		{
			"address":    "乐陵市",
			"start_year": "1988",
			"end_year":   "1993",
		},
	},
	372403: {
		{
			"address":    "禹城市",
			"start_year": "1993",
			"end_year":   "1993",
		},
	},
	372421: {
		{
			"address":    "陵县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372422: {
		{
			"address":    "平原县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372423: {
		{
			"address":    "夏津县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372424: {
		{
			"address":    "武城县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372425: {
		{
			"address":    "齐河县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372426: {
		{
			"address":    "禹城县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	372427: {
		{
			"address":    "乐陵县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	372428: {
		{
			"address":    "临邑县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372429: {
		{
			"address":    "商河县",
			"start_year": "",
			"end_year":   "1988",
		},
	},
	372430: {
		{
			"address":    "济阳县",
			"start_year": "",
			"end_year":   "1988",
		},
	},
	372431: {
		{
			"address":    "宁津县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372432: {
		{
			"address":    "庆云县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372500: {
		{
			"address":    "聊城地区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	372501: {
		{
			"address":    "聊城市",
			"start_year": "1983",
			"end_year":   "1996",
		},
	},
	372502: {
		{
			"address":    "临清市",
			"start_year": "1983",
			"end_year":   "1996",
		},
	},
	372521: {
		{
			"address":    "阳谷县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "聊城县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	372522: {
		{
			"address":    "莘县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "阳谷县",
			"start_year": "1982",
			"end_year":   "1996",
		},
	},
	372523: {
		{
			"address":    "茌平县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "莘县",
			"start_year": "1982",
			"end_year":   "1996",
		},
	},
	372524: {
		{
			"address":    "东阿县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "茌平县",
			"start_year": "1982",
			"end_year":   "1996",
		},
	},
	372525: {
		{
			"address":    "冠县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "东阿县",
			"start_year": "1982",
			"end_year":   "1996",
		},
	},
	372526: {
		{
			"address":    "高唐县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "冠县",
			"start_year": "1982",
			"end_year":   "1996",
		},
	},
	372527: {
		{
			"address":    "临清县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "高唐县",
			"start_year": "1982",
			"end_year":   "1996",
		},
	},
	372528: {
		{
			"address":    "聊城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "临清县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	372600: {
		{
			"address":    "泰安地区",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	372601: {
		{
			"address":    "泰安市",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	372602: {
		{
			"address":    "新汶市",
			"start_year": "1982",
			"end_year":   "1982",
		},
		{
			"address":    "莱芜市",
			"start_year": "1983",
			"end_year":   "1984",
		},
	},
	372603: {
		{
			"address":    "新泰市",
			"start_year": "1983",
			"end_year":   "1984",
		},
	},
	372621: {
		{
			"address":    "泰安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "莱芜县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	372622: {
		{
			"address":    "莱芜县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "新泰县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	372623: {
		{
			"address":    "新泰县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	372624: {
		{
			"address":    "宁阳县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	372625: {
		{
			"address":    "肥城县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	372626: {
		{
			"address":    "东平县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	372627: {
		{
			"address":    "平阴县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	372628: {
		{
			"address":    "新汶县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	372629: {
		{
			"address":    "汶上县",
			"start_year": "1983",
			"end_year":   "1984",
		},
	},
	372630: {
		{
			"address":    "泗水县",
			"start_year": "1983",
			"end_year":   "1984",
		},
	},
	372700: {
		{
			"address":    "济宁地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372701: {
		{
			"address":    "济宁市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372721: {
		{
			"address":    "济宁县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372722: {
		{
			"address":    "兖州县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372723: {
		{
			"address":    "曲阜县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372724: {
		{
			"address":    "泗水县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372725: {
		{
			"address":    "邹县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372726: {
		{
			"address":    "微山县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372727: {
		{
			"address":    "鱼台县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372728: {
		{
			"address":    "金乡县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372729: {
		{
			"address":    "嘉祥县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372730: {
		{
			"address":    "汶上县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372800: {
		{
			"address":    "临沂地区",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372801: {
		{
			"address":    "临沂市",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	372821: {
		{
			"address":    "临沂县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372822: {
		{
			"address":    "郯城县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372823: {
		{
			"address":    "苍山县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372824: {
		{
			"address":    "莒南县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372825: {
		{
			"address":    "日照县",
			"start_year": "",
			"end_year":   "1988",
		},
	},
	372826: {
		{
			"address":    "莒县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	372827: {
		{
			"address":    "沂水县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372828: {
		{
			"address":    "沂源县",
			"start_year": "",
			"end_year":   "1988",
		},
	},
	372829: {
		{
			"address":    "蒙阴县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372830: {
		{
			"address":    "平邑县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372831: {
		{
			"address":    "费县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372832: {
		{
			"address":    "沂南县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372833: {
		{
			"address":    "临沭县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	372900: {
		{
			"address":    "菏泽地区",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	372901: {
		{
			"address":    "菏泽市",
			"start_year": "1983",
			"end_year":   "1999",
		},
	},
	372921: {
		{
			"address":    "菏泽县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	372922: {
		{
			"address":    "曹县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	372923: {
		{
			"address":    "定陶县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	372924: {
		{
			"address":    "成武县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	372925: {
		{
			"address":    "单县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	372926: {
		{
			"address":    "巨野县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	372927: {
		{
			"address":    "梁山县",
			"start_year": "",
			"end_year":   "1988",
		},
	},
	372928: {
		{
			"address":    "郓城县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	372929: {
		{
			"address":    "鄄城县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	372930: {
		{
			"address":    "东明县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	379001: {
		{
			"address":    "青州市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	379002: {
		{
			"address":    "龙口市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	379003: {
		{
			"address":    "曲阜市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	379004: {
		{
			"address":    "莱芜市",
			"start_year": "1986",
			"end_year":   "1991",
		},
	},
	379005: {
		{
			"address":    "新泰市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	379006: {
		{
			"address":    "威海市",
			"start_year": "1986",
			"end_year":   "1986",
		},
		{
			"address":    "胶州市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	379007: {
		{
			"address":    "诸城市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	379008: {
		{
			"address":    "莱阳市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	379009: {
		{
			"address":    "莱州市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	379010: {
		{
			"address":    "滕州市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	379011: {
		{
			"address":    "文登市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	379012: {
		{
			"address":    "荣成市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	379013: {
		{
			"address":    "即墨市",
			"start_year": "1989",
			"end_year":   "1994",
		},
	},
	379014: {
		{
			"address":    "平度市",
			"start_year": "1989",
			"end_year":   "1994",
		},
	},
	379015: {
		{
			"address":    "莱西市",
			"start_year": "1990",
			"end_year":   "1994",
		},
	},
	379016: {
		{
			"address":    "胶南市",
			"start_year": "1990",
			"end_year":   "1994",
		},
	},
	379017: {
		{
			"address":    "蓬莱市",
			"start_year": "1991",
			"end_year":   "1994",
		},
	},
	379018: {
		{
			"address":    "招远市",
			"start_year": "1991",
			"end_year":   "1994",
		},
	},
	379019: {
		{
			"address":    "肥城市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	379020: {
		{
			"address":    "章丘市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	379021: {
		{
			"address":    "兖州市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	379022: {
		{
			"address":    "邹城市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	379023: {
		{
			"address":    "寿光市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	379024: {
		{
			"address":    "乳山市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	379025: {
		{
			"address":    "乐陵市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	379026: {
		{
			"address":    "禹城市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	379027: {
		{
			"address":    "安丘市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	379028: {
		{
			"address":    "昌邑市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	379029: {
		{
			"address":    "高密市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	410000: {
		{
			"address":    "河南省",
			"start_year": "",
			"end_year":   "",
		},
	},
	410100: {
		{
			"address":    "郑州市",
			"start_year": "",
			"end_year":   "",
		},
	},
	410102: {
		{
			"address":    "中原区",
			"start_year": "",
			"end_year":   "",
		},
	},
	410103: {
		{
			"address":    "二七区",
			"start_year": "",
			"end_year":   "",
		},
	},
	410104: {
		{
			"address":    "向阳回族区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "管城回族区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410105: {
		{
			"address":    "金水区",
			"start_year": "",
			"end_year":   "",
		},
	},
	410106: {
		{
			"address":    "上街区",
			"start_year": "",
			"end_year":   "",
		},
	},
	410107: {
		{
			"address":    "新密区",
			"start_year": "1982",
			"end_year":   "1986",
		},
	},
	410108: {
		{
			"address":    "邙山区",
			"start_year": "1987",
			"end_year":   "2002",
		},
		{
			"address":    "惠济区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	410111: {
		{
			"address":    "金海区",
			"start_year": "1981",
			"end_year":   "1986",
		},
	},
	410112: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	410121: {
		{
			"address":    "荥阳县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	410122: {
		{
			"address":    "中牟县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410123: {
		{
			"address":    "新郑县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	410124: {
		{
			"address":    "巩县",
			"start_year": "1983",
			"end_year":   "1990",
		},
	},
	410125: {
		{
			"address":    "登封县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	410126: {
		{
			"address":    "密县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	410181: {
		{
			"address":    "巩义市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	410182: {
		{
			"address":    "荥阳市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	410183: {
		{
			"address":    "新密市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	410184: {
		{
			"address":    "新郑市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	410185: {
		{
			"address":    "登封市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	410200: {
		{
			"address":    "开封市",
			"start_year": "",
			"end_year":   "",
		},
	},
	410202: {
		{
			"address":    "龙亭区",
			"start_year": "",
			"end_year":   "",
		},
	},
	410203: {
		{
			"address":    "顺河回族区",
			"start_year": "",
			"end_year":   "",
		},
	},
	410204: {
		{
			"address":    "古楼区",
			"start_year": "",
			"end_year":   "1988",
		},
		{
			"address":    "鼓楼区",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	410205: {
		{
			"address":    "南关区",
			"start_year": "",
			"end_year":   "2004",
		},
		{
			"address":    "禹王台区",
			"start_year": "2005",
			"end_year":   "",
		},
	},
	410211: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "2004",
		},
		{
			"address":    "金明区",
			"start_year": "2005",
			"end_year":   "2013",
		},
	},
	410212: {
		{
			"address":    "祥符区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	410221: {
		{
			"address":    "杞县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410222: {
		{
			"address":    "通许县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410223: {
		{
			"address":    "尉氏县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410224: {
		{
			"address":    "开封县",
			"start_year": "1983",
			"end_year":   "2013",
		},
	},
	410225: {
		{
			"address":    "兰考县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410300: {
		{
			"address":    "洛阳市",
			"start_year": "",
			"end_year":   "",
		},
	},
	410302: {
		{
			"address":    "洛北区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "老城区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410303: {
		{
			"address":    "西工区",
			"start_year": "",
			"end_year":   "",
		},
	},
	410304: {
		{
			"address":    "瀍河回族区",
			"start_year": "",
			"end_year":   "",
		},
	},
	410305: {
		{
			"address":    "涧西区",
			"start_year": "",
			"end_year":   "",
		},
	},
	410306: {
		{
			"address":    "吉利区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	410311: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1999",
		},
		{
			"address":    "洛龙区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	410321: {
		{
			"address":    "偃师县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	410322: {
		{
			"address":    "孟津县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410323: {
		{
			"address":    "新安县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410324: {
		{
			"address":    "栾川县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410325: {
		{
			"address":    "嵩县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410326: {
		{
			"address":    "汝阳县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410327: {
		{
			"address":    "宜阳县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410328: {
		{
			"address":    "洛宁县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410329: {
		{
			"address":    "伊川县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410381: {
		{
			"address":    "偃师市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	410400: {
		{
			"address":    "平顶山市",
			"start_year": "",
			"end_year":   "",
		},
	},
	410402: {
		{
			"address":    "新华区",
			"start_year": "",
			"end_year":   "",
		},
	},
	410403: {
		{
			"address":    "卫东区",
			"start_year": "",
			"end_year":   "",
		},
	},
	410404: {
		{
			"address":    "石龙区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	410411: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1993",
		},
		{
			"address":    "湛河区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	410412: {
		{
			"address":    "舞钢区",
			"start_year": "1982",
			"end_year":   "1989",
		},
	},
	410421: {
		{
			"address":    "宝丰县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410422: {
		{
			"address":    "叶县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410423: {
		{
			"address":    "鲁山县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410424: {
		{
			"address":    "临汝县",
			"start_year": "1986",
			"end_year":   "1987",
		},
	},
	410425: {
		{
			"address":    "郏县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410426: {
		{
			"address":    "襄城县",
			"start_year": "1986",
			"end_year":   "1996",
		},
	},
	410481: {
		{
			"address":    "舞钢市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	410482: {
		{
			"address":    "汝州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	410500: {
		{
			"address":    "鹤壁市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "安阳市",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	410502: {
		{
			"address":    "鹤山区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "文峰区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	410503: {
		{
			"address":    "山城区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "北关区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	410504: {
		{
			"address":    "铁西区",
			"start_year": "1984",
			"end_year":   "2001",
		},
	},
	410505: {
		{
			"address":    "殷都区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	410506: {
		{
			"address":    "龙安区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	410511: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	410521: {
		{
			"address":    "林县",
			"start_year": "1984",
			"end_year":   "1993",
		},
	},
	410522: {
		{
			"address":    "安阳县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	410523: {
		{
			"address":    "汤阴县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	410524: {
		{
			"address":    "淇县",
			"start_year": "1984",
			"end_year":   "1985",
		},
	},
	410525: {
		{
			"address":    "浚县",
			"start_year": "1984",
			"end_year":   "1985",
		},
	},
	410526: {
		{
			"address":    "滑县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410527: {
		{
			"address":    "内黄县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410581: {
		{
			"address":    "林州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	410600: {
		{
			"address":    "焦作市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "鹤壁市",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	410602: {
		{
			"address":    "解放区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "鹤山区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	410603: {
		{
			"address":    "中站区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "山城区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	410604: {
		{
			"address":    "马村区",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	410611: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "2000",
		},
		{
			"address":    "淇滨区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	410621: {
		{
			"address":    "浚县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410622: {
		{
			"address":    "淇县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410700: {
		{
			"address":    "新乡市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410702: {
		{
			"address":    "红旗区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410703: {
		{
			"address":    "新华区",
			"start_year": "1983",
			"end_year":   "2002",
		},
		{
			"address":    "卫滨区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	410704: {
		{
			"address":    "北站区",
			"start_year": "1983",
			"end_year":   "2002",
		},
		{
			"address":    "凤泉区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	410711: {
		{
			"address":    "郊区",
			"start_year": "1983",
			"end_year":   "2002",
		},
		{
			"address":    "牧野区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	410721: {
		{
			"address":    "新乡县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410722: {
		{
			"address":    "汲县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	410723: {
		{
			"address":    "辉县",
			"start_year": "1986",
			"end_year":   "1987",
		},
	},
	410724: {
		{
			"address":    "获嘉县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410725: {
		{
			"address":    "原阳县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410726: {
		{
			"address":    "延津县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410727: {
		{
			"address":    "封丘县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410728: {
		{
			"address":    "长垣县",
			"start_year": "1986",
			"end_year":   "2018",
		},
	},
	410781: {
		{
			"address":    "卫辉市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	410782: {
		{
			"address":    "辉县市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	410783: {
		{
			"address":    "长垣市",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	410800: {
		{
			"address":    "焦作市",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	410802: {
		{
			"address":    "解放区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	410803: {
		{
			"address":    "中站区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	410804: {
		{
			"address":    "马村区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	410811: {
		{
			"address":    "郊区",
			"start_year": "1982",
			"end_year":   "1989",
		},
		{
			"address":    "山阳区",
			"start_year": "1990",
			"end_year":   "",
		},
	},
	410821: {
		{
			"address":    "修武县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410822: {
		{
			"address":    "博爱县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410823: {
		{
			"address":    "武陟县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410824: {
		{
			"address":    "沁阳县",
			"start_year": "1986",
			"end_year":   "1988",
		},
	},
	410825: {
		{
			"address":    "温县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	410826: {
		{
			"address":    "孟县",
			"start_year": "1986",
			"end_year":   "1995",
		},
	},
	410827: {
		{
			"address":    "济源县",
			"start_year": "1986",
			"end_year":   "1987",
		},
	},
	410881: {
		{
			"address":    "济源市",
			"start_year": "1995",
			"end_year":   "1996",
		},
	},
	410882: {
		{
			"address":    "沁阳市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	410883: {
		{
			"address":    "孟州市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	410900: {
		{
			"address":    "濮阳市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410902: {
		{
			"address":    "市区",
			"start_year": "1985",
			"end_year":   "2001",
		},
		{
			"address":    "华龙区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	410911: {
		{
			"address":    "郊区",
			"start_year": "1984",
			"end_year":   "1986",
		},
	},
	410921: {
		{
			"address":    "滑县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	410922: {
		{
			"address":    "清丰县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410923: {
		{
			"address":    "南乐县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410924: {
		{
			"address":    "内黄县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	410925: {
		{
			"address":    "长垣县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	410926: {
		{
			"address":    "范县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410927: {
		{
			"address":    "台前县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	410928: {
		{
			"address":    "濮阳县",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	411000: {
		{
			"address":    "许昌市",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	411002: {
		{
			"address":    "魏都区",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	411003: {
		{
			"address":    "建安区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	411021: {
		{
			"address":    "禹县",
			"start_year": "1986",
			"end_year":   "1987",
		},
	},
	411022: {
		{
			"address":    "长葛县",
			"start_year": "1986",
			"end_year":   "1992",
		},
	},
	411023: {
		{
			"address":    "许昌县",
			"start_year": "1986",
			"end_year":   "2015",
		},
	},
	411024: {
		{
			"address":    "鄢陵县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	411025: {
		{
			"address":    "襄城县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	411081: {
		{
			"address":    "禹州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	411082: {
		{
			"address":    "长葛市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	411100: {
		{
			"address":    "漯河市",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	411102: {
		{
			"address":    "源汇区",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	411103: {
		{
			"address":    "郾城区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	411104: {
		{
			"address":    "召陵区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	411121: {
		{
			"address":    "舞阳县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	411122: {
		{
			"address":    "临颍县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	411123: {
		{
			"address":    "郾城县",
			"start_year": "1986",
			"end_year":   "2003",
		},
	},
	411200: {
		{
			"address":    "三门峡市",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	411202: {
		{
			"address":    "湖滨区",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	411203: {
		{
			"address":    "陕州区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	411221: {
		{
			"address":    "渑池县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	411222: {
		{
			"address":    "陕县",
			"start_year": "1986",
			"end_year":   "2014",
		},
	},
	411223: {
		{
			"address":    "灵宝县",
			"start_year": "1986",
			"end_year":   "1992",
		},
	},
	411224: {
		{
			"address":    "卢氏县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	411281: {
		{
			"address":    "义马市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	411282: {
		{
			"address":    "灵宝市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	411300: {
		{
			"address":    "南阳市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	411302: {
		{
			"address":    "宛城区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	411303: {
		{
			"address":    "卧龙区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	411321: {
		{
			"address":    "南召县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	411322: {
		{
			"address":    "方城县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	411323: {
		{
			"address":    "西峡县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	411324: {
		{
			"address":    "镇平县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	411325: {
		{
			"address":    "内乡县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	411326: {
		{
			"address":    "淅川县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	411327: {
		{
			"address":    "社旗县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	411328: {
		{
			"address":    "唐河县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	411329: {
		{
			"address":    "新野县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	411330: {
		{
			"address":    "桐柏县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	411381: {
		{
			"address":    "邓州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	411400: {
		{
			"address":    "商丘市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	411402: {
		{
			"address":    "梁园区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	411403: {
		{
			"address":    "睢阳区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	411421: {
		{
			"address":    "民权县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	411422: {
		{
			"address":    "睢县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	411423: {
		{
			"address":    "宁陵县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	411424: {
		{
			"address":    "柘城县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	411425: {
		{
			"address":    "虞城县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	411426: {
		{
			"address":    "夏邑县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	411481: {
		{
			"address":    "永城市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	411500: {
		{
			"address":    "信阳市",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	411502: {
		{
			"address":    "浉河区",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	411503: {
		{
			"address":    "平桥区",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	411521: {
		{
			"address":    "罗山县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	411522: {
		{
			"address":    "光山县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	411523: {
		{
			"address":    "新县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	411524: {
		{
			"address":    "商城县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	411525: {
		{
			"address":    "固始县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	411526: {
		{
			"address":    "潢川县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	411527: {
		{
			"address":    "淮滨县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	411528: {
		{
			"address":    "息县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	411600: {
		{
			"address":    "周口市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411602: {
		{
			"address":    "川汇区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411603: {
		{
			"address":    "淮阳区",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	411621: {
		{
			"address":    "扶沟县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411622: {
		{
			"address":    "西华县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411623: {
		{
			"address":    "商水县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411624: {
		{
			"address":    "沈丘县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411625: {
		{
			"address":    "郸城县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411626: {
		{
			"address":    "淮阳县",
			"start_year": "2000",
			"end_year":   "2018",
		},
	},
	411627: {
		{
			"address":    "太康县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411628: {
		{
			"address":    "鹿邑县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411681: {
		{
			"address":    "项城市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411700: {
		{
			"address":    "驻马店市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411702: {
		{
			"address":    "驿城区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411721: {
		{
			"address":    "西平县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411722: {
		{
			"address":    "上蔡县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411723: {
		{
			"address":    "平舆县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411724: {
		{
			"address":    "正阳县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411725: {
		{
			"address":    "确山县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411726: {
		{
			"address":    "泌阳县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411727: {
		{
			"address":    "汝南县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411728: {
		{
			"address":    "遂平县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	411729: {
		{
			"address":    "新蔡县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	412100: {
		{
			"address":    "安阳地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412101: {
		{
			"address":    "安阳市",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	412102: {
		{
			"address":    "文峰区",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	412103: {
		{
			"address":    "北关区",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	412104: {
		{
			"address":    "铁西区",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	412111: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	412121: {
		{
			"address":    "林县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	412122: {
		{
			"address":    "安阳县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	412123: {
		{
			"address":    "汤阴县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	412124: {
		{
			"address":    "淇县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	412125: {
		{
			"address":    "浚县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	412126: {
		{
			"address":    "濮阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412127: {
		{
			"address":    "滑县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412128: {
		{
			"address":    "清丰县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412129: {
		{
			"address":    "南乐县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412130: {
		{
			"address":    "内黄县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412131: {
		{
			"address":    "长垣县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412132: {
		{
			"address":    "范县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412133: {
		{
			"address":    "台前县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412200: {
		{
			"address":    "新乡地区",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412201: {
		{
			"address":    "新乡市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412202: {
		{
			"address":    "红旗区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412203: {
		{
			"address":    "新华区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412204: {
		{
			"address":    "北站区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	412211: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412221: {
		{
			"address":    "沁阳县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412222: {
		{
			"address":    "博爱县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412223: {
		{
			"address":    "济源县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412224: {
		{
			"address":    "孟县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412225: {
		{
			"address":    "温县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412226: {
		{
			"address":    "武陟县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412227: {
		{
			"address":    "修武县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412228: {
		{
			"address":    "获嘉县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412229: {
		{
			"address":    "新乡县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412230: {
		{
			"address":    "辉县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412231: {
		{
			"address":    "汲县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412232: {
		{
			"address":    "原阳县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412233: {
		{
			"address":    "延津县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412234: {
		{
			"address":    "封丘县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412300: {
		{
			"address":    "商丘地区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	412301: {
		{
			"address":    "商丘市",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	412302: {
		{
			"address":    "永城市",
			"start_year": "1996",
			"end_year":   "1996",
		},
	},
	412321: {
		{
			"address":    "虞城县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	412322: {
		{
			"address":    "商丘县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	412323: {
		{
			"address":    "民权县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	412324: {
		{
			"address":    "宁陵县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	412325: {
		{
			"address":    "睢县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	412326: {
		{
			"address":    "夏邑县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	412327: {
		{
			"address":    "柘城县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	412328: {
		{
			"address":    "永城县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	412400: {
		{
			"address":    "开封地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412421: {
		{
			"address":    "杞县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412422: {
		{
			"address":    "通许县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412423: {
		{
			"address":    "尉氏县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412424: {
		{
			"address":    "开封县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412425: {
		{
			"address":    "中牟县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412426: {
		{
			"address":    "新郑县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "巩县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	412427: {
		{
			"address":    "巩县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "登封县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	412428: {
		{
			"address":    "登封县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "新郑县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	412429: {
		{
			"address":    "密县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412430: {
		{
			"address":    "兰考县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412500: {
		{
			"address":    "洛阳地区",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412501: {
		{
			"address":    "三门峡市",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412502: {
		{
			"address":    "义马市",
			"start_year": "1981",
			"end_year":   "1985",
		},
	},
	412521: {
		{
			"address":    "偃师县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412522: {
		{
			"address":    "孟津县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412523: {
		{
			"address":    "新安县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412524: {
		{
			"address":    "渑池县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412525: {
		{
			"address":    "陕县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412526: {
		{
			"address":    "灵宝县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412527: {
		{
			"address":    "伊川县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412528: {
		{
			"address":    "汝阳县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412529: {
		{
			"address":    "嵩县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412530: {
		{
			"address":    "洛宁县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412531: {
		{
			"address":    "卢氏县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412532: {
		{
			"address":    "栾川县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412533: {
		{
			"address":    "临汝县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412534: {
		{
			"address":    "宜阳县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412600: {
		{
			"address":    "许昌地区",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412601: {
		{
			"address":    "许昌市",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412602: {
		{
			"address":    "漯河市",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412621: {
		{
			"address":    "长葛县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412622: {
		{
			"address":    "禹县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412623: {
		{
			"address":    "鄢陵县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412624: {
		{
			"address":    "许昌县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412625: {
		{
			"address":    "郏县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412626: {
		{
			"address":    "临颍县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412627: {
		{
			"address":    "襄城县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412628: {
		{
			"address":    "宝丰县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412629: {
		{
			"address":    "郾城县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412630: {
		{
			"address":    "叶县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412631: {
		{
			"address":    "鲁山县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	412632: {
		{
			"address":    "舞阳县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	412700: {
		{
			"address":    "周口地区",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412701: {
		{
			"address":    "周口市",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412702: {
		{
			"address":    "项城市",
			"start_year": "1993",
			"end_year":   "1999",
		},
	},
	412721: {
		{
			"address":    "扶沟县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412722: {
		{
			"address":    "西华县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412723: {
		{
			"address":    "商水县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412724: {
		{
			"address":    "太康县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412725: {
		{
			"address":    "鹿邑县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412726: {
		{
			"address":    "郸城县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412727: {
		{
			"address":    "淮阳县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412728: {
		{
			"address":    "沈丘县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412729: {
		{
			"address":    "项城县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	412800: {
		{
			"address":    "驻马店地区",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412801: {
		{
			"address":    "驻马店市",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412821: {
		{
			"address":    "确山县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412822: {
		{
			"address":    "泌阳县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412823: {
		{
			"address":    "遂平县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412824: {
		{
			"address":    "西平县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412825: {
		{
			"address":    "上蔡县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412826: {
		{
			"address":    "汝南县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412827: {
		{
			"address":    "平舆县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412828: {
		{
			"address":    "新蔡县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412829: {
		{
			"address":    "正阳县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	412900: {
		{
			"address":    "南阳地区",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	412901: {
		{
			"address":    "南阳市",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	412902: {
		{
			"address":    "邓州市",
			"start_year": "1988",
			"end_year":   "1993",
		},
	},
	412921: {
		{
			"address":    "南召县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	412922: {
		{
			"address":    "方城县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	412923: {
		{
			"address":    "西峡县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	412924: {
		{
			"address":    "南阳县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	412925: {
		{
			"address":    "镇平县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	412926: {
		{
			"address":    "内乡县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	412927: {
		{
			"address":    "淅川县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	412928: {
		{
			"address":    "社旗县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	412929: {
		{
			"address":    "唐河县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	412930: {
		{
			"address":    "邓县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	412931: {
		{
			"address":    "新野县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	412932: {
		{
			"address":    "桐柏县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	413000: {
		{
			"address":    "信阳地区",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	413001: {
		{
			"address":    "信阳市",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	413021: {
		{
			"address":    "息县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	413022: {
		{
			"address":    "淮滨县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	413023: {
		{
			"address":    "信阳县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	413024: {
		{
			"address":    "潢川县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	413025: {
		{
			"address":    "光山县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	413026: {
		{
			"address":    "固始县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	413027: {
		{
			"address":    "商城县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	413028: {
		{
			"address":    "罗山县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	413029: {
		{
			"address":    "新县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	419001: {
		{
			"address":    "义马市",
			"start_year": "1986",
			"end_year":   "1994",
		},
		{
			"address":    "济源市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	419002: {
		{
			"address":    "汝州市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	419003: {
		{
			"address":    "济源市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	419004: {
		{
			"address":    "禹州市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	419005: {
		{
			"address":    "卫辉市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	419006: {
		{
			"address":    "辉县市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	419007: {
		{
			"address":    "沁阳市",
			"start_year": "1989",
			"end_year":   "1994",
		},
	},
	419008: {
		{
			"address":    "舞钢市",
			"start_year": "1990",
			"end_year":   "1994",
		},
	},
	419009: {
		{
			"address":    "巩义市",
			"start_year": "1991",
			"end_year":   "1994",
		},
	},
	419010: {
		{
			"address":    "灵宝市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	419011: {
		{
			"address":    "长葛市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	419012: {
		{
			"address":    "偃师市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	419013: {
		{
			"address":    "邓州市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	419014: {
		{
			"address":    "林州市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	419015: {
		{
			"address":    "新密市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	419016: {
		{
			"address":    "荥阳市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	419017: {
		{
			"address":    "新郑市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	419018: {
		{
			"address":    "登封市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	420000: {
		{
			"address":    "湖北省",
			"start_year": "",
			"end_year":   "",
		},
	},
	420100: {
		{
			"address":    "武汉市",
			"start_year": "",
			"end_year":   "",
		},
	},
	420102: {
		{
			"address":    "江岸区",
			"start_year": "",
			"end_year":   "",
		},
	},
	420103: {
		{
			"address":    "江汉区",
			"start_year": "",
			"end_year":   "",
		},
	},
	420104: {
		{
			"address":    "硚口区",
			"start_year": "",
			"end_year":   "",
		},
	},
	420105: {
		{
			"address":    "汉阳区",
			"start_year": "",
			"end_year":   "",
		},
	},
	420106: {
		{
			"address":    "武昌区",
			"start_year": "",
			"end_year":   "",
		},
	},
	420107: {
		{
			"address":    "青山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	420111: {
		{
			"address":    "洪山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	420112: {
		{
			"address":    "东西湖区",
			"start_year": "",
			"end_year":   "",
		},
	},
	420113: {
		{
			"address":    "汉南区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	420114: {
		{
			"address":    "蔡甸区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	420115: {
		{
			"address":    "江夏区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	420116: {
		{
			"address":    "黄陂区",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	420117: {
		{
			"address":    "新洲区",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	420121: {
		{
			"address":    "汉阳县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	420122: {
		{
			"address":    "武昌县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	420123: {
		{
			"address":    "黄陂县",
			"start_year": "1983",
			"end_year":   "1997",
		},
	},
	420124: {
		{
			"address":    "新洲县",
			"start_year": "1983",
			"end_year":   "1997",
		},
	},
	420200: {
		{
			"address":    "黄石市",
			"start_year": "",
			"end_year":   "",
		},
	},
	420202: {
		{
			"address":    "黄石港区",
			"start_year": "",
			"end_year":   "",
		},
	},
	420203: {
		{
			"address":    "石灰窑区",
			"start_year": "",
			"end_year":   "2000",
		},
		{
			"address":    "西塞山区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	420204: {
		{
			"address":    "下陆区",
			"start_year": "",
			"end_year":   "",
		},
	},
	420205: {
		{
			"address":    "铁山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	420211: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	420221: {
		{
			"address":    "大冶县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	420222: {
		{
			"address":    "阳新县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	420281: {
		{
			"address":    "大冶市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	420300: {
		{
			"address":    "十堰市",
			"start_year": "",
			"end_year":   "",
		},
	},
	420302: {
		{
			"address":    "茅箭区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	420303: {
		{
			"address":    "张湾区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	420304: {
		{
			"address":    "郧阳区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	420321: {
		{
			"address":    "郧县",
			"start_year": "1994",
			"end_year":   "2013",
		},
	},
	420322: {
		{
			"address":    "郧西县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	420323: {
		{
			"address":    "竹山县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	420324: {
		{
			"address":    "竹溪县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	420325: {
		{
			"address":    "房县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	420381: {
		{
			"address":    "丹江口市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	420400: {
		{
			"address":    "沙市市",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	420500: {
		{
			"address":    "宜昌市",
			"start_year": "",
			"end_year":   "",
		},
	},
	420502: {
		{
			"address":    "西陵区",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	420503: {
		{
			"address":    "伍家岗区",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	420504: {
		{
			"address":    "点军区",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	420505: {
		{
			"address":    "猇亭区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	420506: {
		{
			"address":    "夷陵区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	420521: {
		{
			"address":    "宜昌县",
			"start_year": "1992",
			"end_year":   "2000",
		},
	},
	420523: {
		{
			"address":    "枝江县",
			"start_year": "1992",
			"end_year":   "1995",
		},
	},
	420525: {
		{
			"address":    "远安县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	420526: {
		{
			"address":    "兴山县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	420527: {
		{
			"address":    "秭归县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	420528: {
		{
			"address":    "长阳土家族自治县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	420529: {
		{
			"address":    "五峰土家族自治县",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	420581: {
		{
			"address":    "枝城市",
			"start_year": "1996",
			"end_year":   "1997",
		},
		{
			"address":    "宜都市",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	420582: {
		{
			"address":    "当阳市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	420583: {
		{
			"address":    "枝城市",
			"start_year": "1995",
			"end_year":   "1995",
		},
		{
			"address":    "枝江市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	420600: {
		{
			"address":    "襄樊市",
			"start_year": "",
			"end_year":   "2009",
		},
		{
			"address":    "襄阳市",
			"start_year": "2010",
			"end_year":   "",
		},
	},
	420602: {
		{
			"address":    "襄城区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	420603: {
		{
			"address":    "樊东区",
			"start_year": "1984",
			"end_year":   "1994",
		},
	},
	420604: {
		{
			"address":    "樊西区",
			"start_year": "1984",
			"end_year":   "1994",
		},
	},
	420606: {
		{
			"address":    "樊城区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	420607: {
		{
			"address":    "襄阳区",
			"start_year": "2001",
			"end_year":   "2009",
		},
		{
			"address":    "襄州区",
			"start_year": "2010",
			"end_year":   "",
		},
	},
	420611: {
		{
			"address":    "郊区",
			"start_year": "1984",
			"end_year":   "1994",
		},
	},
	420621: {
		{
			"address":    "襄阳县",
			"start_year": "1983",
			"end_year":   "2000",
		},
	},
	420622: {
		{
			"address":    "枣阳县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	420623: {
		{
			"address":    "宜城县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	420624: {
		{
			"address":    "南漳县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	420625: {
		{
			"address":    "谷城县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	420626: {
		{
			"address":    "保康县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	420682: {
		{
			"address":    "老河口市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	420683: {
		{
			"address":    "枣阳市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	420684: {
		{
			"address":    "宜城市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	420700: {
		{
			"address":    "鄂州市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	420702: {
		{
			"address":    "梁子湖区",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	420703: {
		{
			"address":    "黄州区",
			"start_year": "1984",
			"end_year":   "1986",
		},
		{
			"address":    "华容区",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	420704: {
		{
			"address":    "鄂城区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	420800: {
		{
			"address":    "荆门市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	420802: {
		{
			"address":    "东宝区",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	420803: {
		{
			"address":    "沙洋区",
			"start_year": "1985",
			"end_year":   "1997",
		},
	},
	420804: {
		{
			"address":    "掇刀区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	420821: {
		{
			"address":    "京山县",
			"start_year": "1996",
			"end_year":   "2017",
		},
	},
	420822: {
		{
			"address":    "沙洋县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	420881: {
		{
			"address":    "钟祥市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	420882: {
		{
			"address":    "京山市",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	420900: {
		{
			"address":    "孝感市",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	420901: {
		{
			"address":    "随州市",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	420902: {
		{
			"address":    "孝南区",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	420921: {
		{
			"address":    "孝昌县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	420922: {
		{
			"address":    "大悟县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	420923: {
		{
			"address":    "云梦县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	420924: {
		{
			"address":    "汉川县",
			"start_year": "1993",
			"end_year":   "1996",
		},
	},
	420981: {
		{
			"address":    "应城市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	420982: {
		{
			"address":    "安陆市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	420983: {
		{
			"address":    "广水市",
			"start_year": "1995",
			"end_year":   "1999",
		},
	},
	420984: {
		{
			"address":    "汉川市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	421000: {
		{
			"address":    "荆沙市",
			"start_year": "1994",
			"end_year":   "1995",
		},
		{
			"address":    "荆州市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	421001: {
		{
			"address":    "老河口市",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	421002: {
		{
			"address":    "沙市区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	421003: {
		{
			"address":    "荆州区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	421004: {
		{
			"address":    "江陵区",
			"start_year": "1994",
			"end_year":   "1997",
		},
	},
	421022: {
		{
			"address":    "公安县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	421023: {
		{
			"address":    "监利县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	421024: {
		{
			"address":    "松滋县",
			"start_year": "1994",
			"end_year":   "1994",
		},
		{
			"address":    "江陵县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	421025: {
		{
			"address":    "京山县",
			"start_year": "1994",
			"end_year":   "1995",
		},
	},
	421081: {
		{
			"address":    "石首市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	421082: {
		{
			"address":    "钟祥市",
			"start_year": "1995",
			"end_year":   "1995",
		},
	},
	421083: {
		{
			"address":    "洪湖市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	421087: {
		{
			"address":    "松滋市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	421100: {
		{
			"address":    "黄冈市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	421102: {
		{
			"address":    "黄州区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	421121: {
		{
			"address":    "团风县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	421122: {
		{
			"address":    "红安县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	421123: {
		{
			"address":    "罗田县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	421124: {
		{
			"address":    "英山县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	421125: {
		{
			"address":    "浠水县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	421126: {
		{
			"address":    "蕲春县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	421127: {
		{
			"address":    "黄梅县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	421181: {
		{
			"address":    "麻城市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	421182: {
		{
			"address":    "武穴市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	421200: {
		{
			"address":    "咸宁市",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	421202: {
		{
			"address":    "咸安区",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	421221: {
		{
			"address":    "嘉鱼县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	421222: {
		{
			"address":    "通城县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	421223: {
		{
			"address":    "崇阳县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	421224: {
		{
			"address":    "通山县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	421281: {
		{
			"address":    "赤壁市",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	421300: {
		{
			"address":    "随州市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	421303: {
		{
			"address":    "曾都区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	421321: {
		{
			"address":    "随县",
			"start_year": "2009",
			"end_year":   "",
		},
	},
	421381: {
		{
			"address":    "广水市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	422100: {
		{
			"address":    "黄冈地区",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	422101: {
		{
			"address":    "鄂城市",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "麻城市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	422102: {
		{
			"address":    "武穴市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	422103: {
		{
			"address":    "黄州市",
			"start_year": "1990",
			"end_year":   "1994",
		},
	},
	422121: {
		{
			"address":    "黄冈县",
			"start_year": "",
			"end_year":   "1989",
		},
	},
	422122: {
		{
			"address":    "新洲县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422123: {
		{
			"address":    "红安县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	422124: {
		{
			"address":    "麻城县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	422125: {
		{
			"address":    "罗田县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	422126: {
		{
			"address":    "英山县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	422127: {
		{
			"address":    "浠水县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	422128: {
		{
			"address":    "蕲春县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	422129: {
		{
			"address":    "广济县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	422130: {
		{
			"address":    "黄梅县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	422131: {
		{
			"address":    "鄂城县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422200: {
		{
			"address":    "孝感地区",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	422201: {
		{
			"address":    "孝感市",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	422202: {
		{
			"address":    "应城市",
			"start_year": "1986",
			"end_year":   "1992",
		},
	},
	422203: {
		{
			"address":    "安陆市",
			"start_year": "1987",
			"end_year":   "1992",
		},
	},
	422204: {
		{
			"address":    "广水市",
			"start_year": "1988",
			"end_year":   "1992",
		},
	},
	422221: {
		{
			"address":    "孝感县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422222: {
		{
			"address":    "黄陂县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422223: {
		{
			"address":    "大悟县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	422224: {
		{
			"address":    "应山县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	422225: {
		{
			"address":    "安陆县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	422226: {
		{
			"address":    "云梦县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	422227: {
		{
			"address":    "应城县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	422228: {
		{
			"address":    "汉川县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	422300: {
		{
			"address":    "咸宁地区",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	422301: {
		{
			"address":    "咸宁市",
			"start_year": "1983",
			"end_year":   "1997",
		},
	},
	422302: {
		{
			"address":    "蒲圻市",
			"start_year": "1986",
			"end_year":   "1997",
		},
	},
	422321: {
		{
			"address":    "咸宁县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422322: {
		{
			"address":    "嘉鱼县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	422323: {
		{
			"address":    "蒲圻县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	422324: {
		{
			"address":    "通城县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	422325: {
		{
			"address":    "崇阳县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	422326: {
		{
			"address":    "通山县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	422327: {
		{
			"address":    "阳新县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	422400: {
		{
			"address":    "荆州地区",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	422401: {
		{
			"address":    "荆门市",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "仙桃市",
			"start_year": "1986",
			"end_year":   "1993",
		},
	},
	422402: {
		{
			"address":    "石首市",
			"start_year": "1986",
			"end_year":   "1993",
		},
	},
	422403: {
		{
			"address":    "洪湖市",
			"start_year": "1987",
			"end_year":   "1993",
		},
	},
	422404: {
		{
			"address":    "天门市",
			"start_year": "1987",
			"end_year":   "1993",
		},
	},
	422405: {
		{
			"address":    "潜江市",
			"start_year": "1988",
			"end_year":   "1993",
		},
	},
	422406: {
		{
			"address":    "钟祥市",
			"start_year": "1992",
			"end_year":   "1993",
		},
	},
	422421: {
		{
			"address":    "江陵县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	422422: {
		{
			"address":    "松滋县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	422423: {
		{
			"address":    "公安县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	422424: {
		{
			"address":    "石首县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	422425: {
		{
			"address":    "监利县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	422426: {
		{
			"address":    "洪湖县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	422427: {
		{
			"address":    "沔阳县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	422428: {
		{
			"address":    "天门县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	422429: {
		{
			"address":    "潜江县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	422430: {
		{
			"address":    "荆门县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422431: {
		{
			"address":    "钟祥县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	422432: {
		{
			"address":    "京山县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	422500: {
		{
			"address":    "襄阳地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422501: {
		{
			"address":    "随州市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422502: {
		{
			"address":    "老河口市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422521: {
		{
			"address":    "樊阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422522: {
		{
			"address":    "枣阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422523: {
		{
			"address":    "随县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422524: {
		{
			"address":    "宜城县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422525: {
		{
			"address":    "南漳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422526: {
		{
			"address":    "光化县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422527: {
		{
			"address":    "谷城县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422528: {
		{
			"address":    "保康县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422600: {
		{
			"address":    "郧阳地区",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	422601: {
		{
			"address":    "丹江口市",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	422621: {
		{
			"address":    "均县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422622: {
		{
			"address":    "郧县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	422623: {
		{
			"address":    "郧西县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	422624: {
		{
			"address":    "竹山县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	422625: {
		{
			"address":    "竹溪县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	422626: {
		{
			"address":    "房县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	422627: {
		{
			"address":    "神农架林区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422700: {
		{
			"address":    "宜昌地区",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	422701: {
		{
			"address":    "枝城市",
			"start_year": "1987",
			"end_year":   "1991",
		},
	},
	422702: {
		{
			"address":    "当阳市",
			"start_year": "1988",
			"end_year":   "1991",
		},
	},
	422721: {
		{
			"address":    "宜昌县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	422722: {
		{
			"address":    "宜都县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	422723: {
		{
			"address":    "枝江县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	422724: {
		{
			"address":    "当阳县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	422725: {
		{
			"address":    "远安县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	422726: {
		{
			"address":    "兴山县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	422727: {
		{
			"address":    "秭归县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	422728: {
		{
			"address":    "长阳县",
			"start_year": "",
			"end_year":   "1983",
		},
		{
			"address":    "长阳土家族自治县",
			"start_year": "1984",
			"end_year":   "1991",
		},
	},
	422729: {
		{
			"address":    "五峰县",
			"start_year": "",
			"end_year":   "1983",
		},
		{
			"address":    "五峰土家族自治县",
			"start_year": "1984",
			"end_year":   "1991",
		},
	},
	422800: {
		{
			"address":    "恩施地区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "鄂西土家族苗族自治州",
			"start_year": "1983",
			"end_year":   "1992",
		},
		{
			"address":    "恩施土家族苗族自治州",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	422801: {
		{
			"address":    "恩施市",
			"start_year": "1981",
			"end_year":   "",
		},
	},
	422802: {
		{
			"address":    "利川市",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	422821: {
		{
			"address":    "恩施县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	422822: {
		{
			"address":    "建始县",
			"start_year": "",
			"end_year":   "",
		},
	},
	422823: {
		{
			"address":    "巴东县",
			"start_year": "",
			"end_year":   "",
		},
	},
	422824: {
		{
			"address":    "利川县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	422825: {
		{
			"address":    "宣恩县",
			"start_year": "",
			"end_year":   "",
		},
	},
	422826: {
		{
			"address":    "咸丰县",
			"start_year": "",
			"end_year":   "",
		},
	},
	422827: {
		{
			"address":    "来凤土家族自治县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "来凤县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	422828: {
		{
			"address":    "鹤峰土家族自治县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "鹤峰县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	422921: {
		{
			"address":    "神农架林区",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	429001: {
		{
			"address":    "随州市",
			"start_year": "1984",
			"end_year":   "1999",
		},
	},
	429002: {
		{
			"address":    "老河口市",
			"start_year": "1984",
			"end_year":   "1994",
		},
	},
	429003: {
		{
			"address":    "枣阳市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	429004: {
		{
			"address":    "仙桃市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	429005: {
		{
			"address":    "潜江市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	429006: {
		{
			"address":    "天门市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	429007: {
		{
			"address":    "枝城市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	429008: {
		{
			"address":    "当阳市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	429009: {
		{
			"address":    "应城市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	429010: {
		{
			"address":    "安陆市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	429011: {
		{
			"address":    "广水市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	429012: {
		{
			"address":    "石首市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	429013: {
		{
			"address":    "洪湖市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	429014: {
		{
			"address":    "钟祥市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	429015: {
		{
			"address":    "丹江口市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	429016: {
		{
			"address":    "大冶市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	429017: {
		{
			"address":    "宜城市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	429021: {
		{
			"address":    "神农架林区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	430000: {
		{
			"address":    "湖南省",
			"start_year": "",
			"end_year":   "",
		},
	},
	430100: {
		{
			"address":    "长沙市",
			"start_year": "",
			"end_year":   "",
		},
	},
	430102: {
		{
			"address":    "城东区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "东区",
			"start_year": "1983",
			"end_year":   "1995",
		},
		{
			"address":    "芙蓉区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	430103: {
		{
			"address":    "城南区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "南区",
			"start_year": "1983",
			"end_year":   "1995",
		},
		{
			"address":    "天心区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	430104: {
		{
			"address":    "城西区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "西区",
			"start_year": "1983",
			"end_year":   "1995",
		},
		{
			"address":    "岳麓区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	430105: {
		{
			"address":    "城北区",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "北区",
			"start_year": "1983",
			"end_year":   "1995",
		},
		{
			"address":    "开福区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	430111: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1995",
		},
		{
			"address":    "雨花区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	430112: {
		{
			"address":    "望城区",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	430121: {
		{
			"address":    "长沙县",
			"start_year": "",
			"end_year":   "",
		},
	},
	430122: {
		{
			"address":    "望城县",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	430123: {
		{
			"address":    "浏阳县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	430124: {
		{
			"address":    "宁乡县",
			"start_year": "1983",
			"end_year":   "2016",
		},
	},
	430181: {
		{
			"address":    "浏阳市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	430182: {
		{
			"address":    "宁乡市",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	430200: {
		{
			"address":    "株洲市",
			"start_year": "",
			"end_year":   "",
		},
	},
	430202: {
		{
			"address":    "东区",
			"start_year": "",
			"end_year":   "1996",
		},
		{
			"address":    "荷塘区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	430203: {
		{
			"address":    "北区",
			"start_year": "",
			"end_year":   "1996",
		},
		{
			"address":    "芦淞区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	430204: {
		{
			"address":    "南区",
			"start_year": "",
			"end_year":   "1996",
		},
		{
			"address":    "石峰区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	430211: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1996",
		},
		{
			"address":    "天元区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	430212: {
		{
			"address":    "渌口区",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	430219: {
		{
			"address":    "醴陵市",
			"start_year": "1985",
			"end_year":   "1985",
		},
	},
	430221: {
		{
			"address":    "株洲县",
			"start_year": "",
			"end_year":   "2017",
		},
	},
	430222: {
		{
			"address":    "醴陵县",
			"start_year": "1983",
			"end_year":   "1984",
		},
	},
	430223: {
		{
			"address":    "攸县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	430224: {
		{
			"address":    "茶陵县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	430225: {
		{
			"address":    "酃县",
			"start_year": "1983",
			"end_year":   "1993",
		},
		{
			"address":    "炎陵县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	430281: {
		{
			"address":    "醴陵市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	430300: {
		{
			"address":    "湘潭市",
			"start_year": "",
			"end_year":   "",
		},
		{
			"address":    "邵阳市",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	430302: {
		{
			"address":    "雨湖区",
			"start_year": "",
			"end_year":   "",
		},
		{
			"address":    "东区",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	430303: {
		{
			"address":    "湘江区",
			"start_year": "",
			"end_year":   "1991",
		},
		{
			"address":    "西区",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	430304: {
		{
			"address":    "岳塘区",
			"start_year": "",
			"end_year":   "",
		},
		{
			"address":    "桥头区",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	430305: {
		{
			"address":    "板塘区",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	430306: {
		{
			"address":    "韶山区",
			"start_year": "1988",
			"end_year":   "1989",
		},
	},
	430311: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	430321: {
		{
			"address":    "邵东县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "湘潭县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	430322: {
		{
			"address":    "新邵县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "湘乡县",
			"start_year": "1984",
			"end_year":   "1985",
		},
	},
	430381: {
		{
			"address":    "湘乡市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	430382: {
		{
			"address":    "韶山市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	430400: {
		{
			"address":    "衡阳市",
			"start_year": "",
			"end_year":   "",
		},
		{
			"address":    "湘潭市",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	430402: {
		{
			"address":    "江东区",
			"start_year": "",
			"end_year":   "2000",
		},
		{
			"address":    "雨湖区",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	430403: {
		{
			"address":    "城南区",
			"start_year": "",
			"end_year":   "2000",
		},
		{
			"address":    "湘江区",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	430404: {
		{
			"address":    "城北区",
			"start_year": "",
			"end_year":   "2000",
		},
		{
			"address":    "岳塘区",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	430405: {
		{
			"address":    "板塘区",
			"start_year": "1982",
			"end_year":   "1983",
		},
		{
			"address":    "珠晖区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	430406: {
		{
			"address":    "雁峰区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	430407: {
		{
			"address":    "石鼓区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	430408: {
		{
			"address":    "蒸湘区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	430411: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	430412: {
		{
			"address":    "南岳区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	430421: {
		{
			"address":    "湘潭县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "衡阳县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	430422: {
		{
			"address":    "湘乡县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "衡南县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	430423: {
		{
			"address":    "衡山县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	430424: {
		{
			"address":    "衡东县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	430425: {
		{
			"address":    "常宁县",
			"start_year": "1984",
			"end_year":   "1995",
		},
	},
	430426: {
		{
			"address":    "祁东县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	430427: {
		{
			"address":    "耒阳县",
			"start_year": "1984",
			"end_year":   "1985",
		},
	},
	430481: {
		{
			"address":    "耒阳市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	430482: {
		{
			"address":    "常宁市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	430500: {
		{
			"address":    "邵阳市",
			"start_year": "",
			"end_year":   "",
		},
		{
			"address":    "衡阳市",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	430502: {
		{
			"address":    "东区",
			"start_year": "",
			"end_year":   "1996",
		},
		{
			"address":    "江东区",
			"start_year": "1982",
			"end_year":   "1983",
		},
		{
			"address":    "双清区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	430503: {
		{
			"address":    "西区",
			"start_year": "",
			"end_year":   "1996",
		},
		{
			"address":    "城南区",
			"start_year": "1982",
			"end_year":   "1983",
		},
		{
			"address":    "大祥区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	430504: {
		{
			"address":    "桥头区",
			"start_year": "",
			"end_year":   "1986",
		},
		{
			"address":    "城北区",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	430511: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1996",
		},
		{
			"address":    "北塔区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	430521: {
		{
			"address":    "衡阳县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "邵东县",
			"start_year": "1984",
			"end_year":   "2018",
		},
	},
	430522: {
		{
			"address":    "衡南县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "新邵县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	430523: {
		{
			"address":    "衡山县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "邵阳县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	430524: {
		{
			"address":    "衡东县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "隆回县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	430525: {
		{
			"address":    "常宁县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "洞口县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	430526: {
		{
			"address":    "祁东县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "武冈县",
			"start_year": "1986",
			"end_year":   "1993",
		},
	},
	430527: {
		{
			"address":    "耒阳县",
			"start_year": "1983",
			"end_year":   "1983",
		},
		{
			"address":    "绥宁县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	430528: {
		{
			"address":    "新宁县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	430529: {
		{
			"address":    "城步苗族自治县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	430581: {
		{
			"address":    "武冈市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	430582: {
		{
			"address":    "邵东市",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	430600: {
		{
			"address":    "岳阳市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	430602: {
		{
			"address":    "南区",
			"start_year": "1984",
			"end_year":   "1995",
		},
		{
			"address":    "岳阳楼区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	430603: {
		{
			"address":    "北区",
			"start_year": "1984",
			"end_year":   "1995",
		},
		{
			"address":    "云溪区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	430611: {
		{
			"address":    "郊区",
			"start_year": "1984",
			"end_year":   "1995",
		},
		{
			"address":    "君山区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	430621: {
		{
			"address":    "岳阳县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	430622: {
		{
			"address":    "临湘县",
			"start_year": "1986",
			"end_year":   "1991",
		},
	},
	430623: {
		{
			"address":    "华容县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	430624: {
		{
			"address":    "湘阴县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	430626: {
		{
			"address":    "平江县",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	430627: {
		{
			"address":    "汨罗县",
			"start_year": "1986",
			"end_year":   "1986",
		},
	},
	430681: {
		{
			"address":    "汨罗市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	430682: {
		{
			"address":    "临湘市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	430700: {
		{
			"address":    "常德市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	430702: {
		{
			"address":    "武陵区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	430703: {
		{
			"address":    "鼎城区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	430721: {
		{
			"address":    "安乡县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	430722: {
		{
			"address":    "汉寿县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	430723: {
		{
			"address":    "澧县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	430724: {
		{
			"address":    "临澧县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	430725: {
		{
			"address":    "桃源县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	430726: {
		{
			"address":    "石门县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	430781: {
		{
			"address":    "津市市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	430800: {
		{
			"address":    "大庸市",
			"start_year": "1988",
			"end_year":   "1993",
		},
		{
			"address":    "张家界市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	430802: {
		{
			"address":    "永定区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	430811: {
		{
			"address":    "武陵源区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	430821: {
		{
			"address":    "慈利县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	430822: {
		{
			"address":    "桑植县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	430900: {
		{
			"address":    "益阳市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	430902: {
		{
			"address":    "资阳区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	430903: {
		{
			"address":    "赫山区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	430921: {
		{
			"address":    "南县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	430922: {
		{
			"address":    "桃江县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	430923: {
		{
			"address":    "安化县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	430981: {
		{
			"address":    "沅江市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	431000: {
		{
			"address":    "郴州市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	431002: {
		{
			"address":    "北湖区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	431003: {
		{
			"address":    "苏仙区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	431021: {
		{
			"address":    "桂阳县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	431022: {
		{
			"address":    "宜章县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	431023: {
		{
			"address":    "永兴县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	431024: {
		{
			"address":    "嘉禾县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	431025: {
		{
			"address":    "临武县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	431026: {
		{
			"address":    "汝城县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	431027: {
		{
			"address":    "桂东县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	431028: {
		{
			"address":    "安仁县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	431081: {
		{
			"address":    "资兴市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	431100: {
		{
			"address":    "永州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	431102: {
		{
			"address":    "芝山区",
			"start_year": "1995",
			"end_year":   "2004",
		},
		{
			"address":    "零陵区",
			"start_year": "2005",
			"end_year":   "",
		},
	},
	431103: {
		{
			"address":    "冷水滩区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	431121: {
		{
			"address":    "祁阳县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	431122: {
		{
			"address":    "东安县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	431123: {
		{
			"address":    "双牌县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	431124: {
		{
			"address":    "道县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	431125: {
		{
			"address":    "江永县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	431126: {
		{
			"address":    "宁远县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	431127: {
		{
			"address":    "蓝山县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	431128: {
		{
			"address":    "新田县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	431129: {
		{
			"address":    "江华瑶族自治县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	431200: {
		{
			"address":    "怀化市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	431202: {
		{
			"address":    "鹤城区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	431221: {
		{
			"address":    "中方县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	431222: {
		{
			"address":    "沅陵县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	431223: {
		{
			"address":    "辰溪县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	431224: {
		{
			"address":    "溆浦县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	431225: {
		{
			"address":    "会同县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	431226: {
		{
			"address":    "麻阳苗族自治县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	431227: {
		{
			"address":    "新晃侗族自治县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	431228: {
		{
			"address":    "芷江侗族自治县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	431229: {
		{
			"address":    "靖州苗族侗族自治县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	431230: {
		{
			"address":    "通道侗族自治县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	431281: {
		{
			"address":    "洪江市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	431300: {
		{
			"address":    "娄底市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	431302: {
		{
			"address":    "娄星区",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	431321: {
		{
			"address":    "双峰县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	431322: {
		{
			"address":    "新化县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	431381: {
		{
			"address":    "冷水江市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	431382: {
		{
			"address":    "涟源市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	432100: {
		{
			"address":    "湘潭地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432121: {
		{
			"address":    "湘潭县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432122: {
		{
			"address":    "湘乡县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432123: {
		{
			"address":    "醴陵县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432124: {
		{
			"address":    "浏阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432125: {
		{
			"address":    "攸县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432126: {
		{
			"address":    "茶陵县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432127: {
		{
			"address":    "酃县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432200: {
		{
			"address":    "岳阳地区",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	432201: {
		{
			"address":    "岳阳市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432221: {
		{
			"address":    "岳阳县",
			"start_year": "",
			"end_year":   "1980",
		},
	},
	432222: {
		{
			"address":    "平江县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	432223: {
		{
			"address":    "湘阴县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	432224: {
		{
			"address":    "汨罗县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	432225: {
		{
			"address":    "临湘县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	432226: {
		{
			"address":    "华容县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	432300: {
		{
			"address":    "益阳地区",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432301: {
		{
			"address":    "益阳市",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432302: {
		{
			"address":    "沅江市",
			"start_year": "1988",
			"end_year":   "1993",
		},
	},
	432321: {
		{
			"address":    "益阳县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432322: {
		{
			"address":    "南县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432323: {
		{
			"address":    "沅江县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	432324: {
		{
			"address":    "宁乡县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432325: {
		{
			"address":    "桃江县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432326: {
		{
			"address":    "安化县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432400: {
		{
			"address":    "常德地区",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	432401: {
		{
			"address":    "常德市",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	432402: {
		{
			"address":    "津市市",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	432421: {
		{
			"address":    "常德县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	432422: {
		{
			"address":    "安乡县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	432423: {
		{
			"address":    "汉寿县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	432424: {
		{
			"address":    "澧县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	432425: {
		{
			"address":    "临澧县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	432426: {
		{
			"address":    "桃源县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	432427: {
		{
			"address":    "石门县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	432428: {
		{
			"address":    "慈利县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	432500: {
		{
			"address":    "涟源地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "娄底地区",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	432501: {
		{
			"address":    "娄底市",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	432502: {
		{
			"address":    "冷水江市",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	432503: {
		{
			"address":    "涟源市",
			"start_year": "1987",
			"end_year":   "1998",
		},
	},
	432521: {
		{
			"address":    "涟源县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	432522: {
		{
			"address":    "双峰县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	432523: {
		{
			"address":    "邵东县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432524: {
		{
			"address":    "新化县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	432525: {
		{
			"address":    "新邵县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432600: {
		{
			"address":    "邵阳地区",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	432621: {
		{
			"address":    "邵阳县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	432622: {
		{
			"address":    "隆回县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	432623: {
		{
			"address":    "武冈县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	432624: {
		{
			"address":    "洞口县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	432625: {
		{
			"address":    "新宁县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	432626: {
		{
			"address":    "绥宁县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	432627: {
		{
			"address":    "城步苗族自治县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	432700: {
		{
			"address":    "衡阳地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432721: {
		{
			"address":    "衡阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432722: {
		{
			"address":    "衡南县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432723: {
		{
			"address":    "衡山县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432724: {
		{
			"address":    "衡东县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432725: {
		{
			"address":    "常宁县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432726: {
		{
			"address":    "祁东县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432727: {
		{
			"address":    "祁阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432800: {
		{
			"address":    "郴州地区",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432801: {
		{
			"address":    "郴州市",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432802: {
		{
			"address":    "资兴市",
			"start_year": "1984",
			"end_year":   "1993",
		},
	},
	432821: {
		{
			"address":    "郴县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432822: {
		{
			"address":    "桂阳县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432823: {
		{
			"address":    "永兴县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432824: {
		{
			"address":    "宜章县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432825: {
		{
			"address":    "资兴县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	432826: {
		{
			"address":    "嘉禾县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432827: {
		{
			"address":    "临武县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432828: {
		{
			"address":    "汝城县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432829: {
		{
			"address":    "桂东县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432830: {
		{
			"address":    "耒阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	432831: {
		{
			"address":    "安仁县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	432900: {
		{
			"address":    "零陵地区",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	432901: {
		{
			"address":    "永州市",
			"start_year": "1982",
			"end_year":   "1994",
		},
	},
	432902: {
		{
			"address":    "冷水滩市",
			"start_year": "1984",
			"end_year":   "1994",
		},
	},
	432921: {
		{
			"address":    "零陵县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	432922: {
		{
			"address":    "东安县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	432923: {
		{
			"address":    "道县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	432924: {
		{
			"address":    "宁远县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	432925: {
		{
			"address":    "江永县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	432926: {
		{
			"address":    "江华瑶族自治县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	432927: {
		{
			"address":    "蓝山县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	432928: {
		{
			"address":    "新田县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	432929: {
		{
			"address":    "双牌县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	432930: {
		{
			"address":    "祁阳县",
			"start_year": "1983",
			"end_year":   "1994",
		},
	},
	433000: {
		{
			"address":    "黔阳地区",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "怀化地区",
			"start_year": "1981",
			"end_year":   "1996",
		},
	},
	433001: {
		{
			"address":    "洪江市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "怀化市",
			"start_year": "1982",
			"end_year":   "1996",
		},
	},
	433002: {
		{
			"address":    "怀化市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "洪江市",
			"start_year": "1982",
			"end_year":   "1996",
		},
	},
	433021: {
		{
			"address":    "黔阳县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	433022: {
		{
			"address":    "沅陵县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	433023: {
		{
			"address":    "辰溪县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	433024: {
		{
			"address":    "溆浦县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	433025: {
		{
			"address":    "麻阳县",
			"start_year": "",
			"end_year":   "1987",
		},
		{
			"address":    "麻阳苗族自治县",
			"start_year": "1988",
			"end_year":   "1996",
		},
	},
	433026: {
		{
			"address":    "新晃侗族自治县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	433027: {
		{
			"address":    "芷江县",
			"start_year": "",
			"end_year":   "1985",
		},
		{
			"address":    "芷江侗族自治县",
			"start_year": "1986",
			"end_year":   "1996",
		},
	},
	433028: {
		{
			"address":    "怀化县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	433029: {
		{
			"address":    "会同县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	433030: {
		{
			"address":    "靖县",
			"start_year": "",
			"end_year":   "1986",
		},
		{
			"address":    "靖州苗族侗族自治县",
			"start_year": "1987",
			"end_year":   "1996",
		},
	},
	433031: {
		{
			"address":    "通道侗族自治县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	433100: {
		{
			"address":    "湘西土家族苗族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	433101: {
		{
			"address":    "吉首市",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	433102: {
		{
			"address":    "大庸市",
			"start_year": "1985",
			"end_year":   "1987",
		},
	},
	433121: {
		{
			"address":    "吉首县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	433122: {
		{
			"address":    "泸溪县",
			"start_year": "",
			"end_year":   "",
		},
	},
	433123: {
		{
			"address":    "凤凰县",
			"start_year": "",
			"end_year":   "",
		},
	},
	433124: {
		{
			"address":    "花垣县",
			"start_year": "",
			"end_year":   "",
		},
	},
	433125: {
		{
			"address":    "保靖县",
			"start_year": "",
			"end_year":   "",
		},
	},
	433126: {
		{
			"address":    "古丈县",
			"start_year": "",
			"end_year":   "",
		},
	},
	433127: {
		{
			"address":    "永顺县",
			"start_year": "",
			"end_year":   "",
		},
	},
	433128: {
		{
			"address":    "大庸县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	433129: {
		{
			"address":    "桑植县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	433130: {
		{
			"address":    "龙山县",
			"start_year": "",
			"end_year":   "",
		},
	},
	439001: {
		{
			"address":    "醴陵市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	439002: {
		{
			"address":    "湘乡市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	439003: {
		{
			"address":    "耒阳市",
			"start_year": "1986",
			"end_year":   "1994",
		},
	},
	439004: {
		{
			"address":    "汨罗市",
			"start_year": "1987",
			"end_year":   "1994",
		},
	},
	439005: {
		{
			"address":    "津市市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	439006: {
		{
			"address":    "韶山市",
			"start_year": "1990",
			"end_year":   "1994",
		},
	},
	439007: {
		{
			"address":    "临湘市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	439008: {
		{
			"address":    "浏阳市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	439009: {
		{
			"address":    "资兴市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	439010: {
		{
			"address":    "沅江市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	439011: {
		{
			"address":    "武冈市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	440000: {
		{
			"address":    "广东省",
			"start_year": "",
			"end_year":   "",
		},
	},
	440100: {
		{
			"address":    "广州市",
			"start_year": "",
			"end_year":   "",
		},
	},
	440102: {
		{
			"address":    "东山区",
			"start_year": "",
			"end_year":   "2004",
		},
	},
	440103: {
		{
			"address":    "荔湾区",
			"start_year": "",
			"end_year":   "",
		},
	},
	440104: {
		{
			"address":    "越秀区",
			"start_year": "",
			"end_year":   "",
		},
	},
	440105: {
		{
			"address":    "海珠区",
			"start_year": "",
			"end_year":   "",
		},
	},
	440106: {
		{
			"address":    "芳村区",
			"start_year": "1985",
			"end_year":   "1987",
		},
		{
			"address":    "天河区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	440107: {
		{
			"address":    "天河区",
			"start_year": "1985",
			"end_year":   "1987",
		},
		{
			"address":    "芳村区",
			"start_year": "1988",
			"end_year":   "2004",
		},
	},
	440111: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1986",
		},
		{
			"address":    "白云区",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	440112: {
		{
			"address":    "黄埔区",
			"start_year": "",
			"end_year":   "",
		},
	},
	440113: {
		{
			"address":    "番禺区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	440114: {
		{
			"address":    "花都区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	440115: {
		{
			"address":    "南沙区",
			"start_year": "2005",
			"end_year":   "",
		},
	},
	440116: {
		{
			"address":    "萝岗区",
			"start_year": "2005",
			"end_year":   "2013",
		},
	},
	440117: {
		{
			"address":    "从化区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	440118: {
		{
			"address":    "增城区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	440121: {
		{
			"address":    "花县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	440122: {
		{
			"address":    "从化县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	440123: {
		{
			"address":    "新丰县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	440124: {
		{
			"address":    "龙门县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	440125: {
		{
			"address":    "增城县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	440126: {
		{
			"address":    "番禺县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	440127: {
		{
			"address":    "清远县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	440128: {
		{
			"address":    "佛冈县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	440181: {
		{
			"address":    "番禺市",
			"start_year": "1995",
			"end_year":   "1999",
		},
	},
	440182: {
		{
			"address":    "花都市",
			"start_year": "1995",
			"end_year":   "1999",
		},
	},
	440183: {
		{
			"address":    "增城市",
			"start_year": "1995",
			"end_year":   "2013",
		},
	},
	440184: {
		{
			"address":    "从化市",
			"start_year": "1995",
			"end_year":   "2013",
		},
	},
	440200: {
		{
			"address":    "韶关市",
			"start_year": "",
			"end_year":   "",
		},
	},
	440202: {
		{
			"address":    "北江区",
			"start_year": "1984",
			"end_year":   "2003",
		},
	},
	440203: {
		{
			"address":    "浈江区",
			"start_year": "1984",
			"end_year":   "2002",
		},
		{
			"address":    "武江区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	440204: {
		{
			"address":    "武江区",
			"start_year": "1984",
			"end_year":   "2002",
		},
		{
			"address":    "浈江区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	440205: {
		{
			"address":    "曲江区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	440221: {
		{
			"address":    "曲江县",
			"start_year": "",
			"end_year":   "2003",
		},
	},
	440222: {
		{
			"address":    "始兴县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	440223: {
		{
			"address":    "南雄县",
			"start_year": "1983",
			"end_year":   "1995",
		},
	},
	440224: {
		{
			"address":    "仁化县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	440225: {
		{
			"address":    "乐昌县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	440226: {
		{
			"address":    "连县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	440227: {
		{
			"address":    "阳山县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	440228: {
		{
			"address":    "英德县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	440229: {
		{
			"address":    "翁源县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	440230: {
		{
			"address":    "连山壮族瑶族自治县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	440231: {
		{
			"address":    "连南瑶族自治县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	440232: {
		{
			"address":    "乳源瑶族自治县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	440233: {
		{
			"address":    "新丰县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	440281: {
		{
			"address":    "乐昌市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	440282: {
		{
			"address":    "南雄市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	440300: {
		{
			"address":    "深圳市",
			"start_year": "",
			"end_year":   "",
		},
	},
	440303: {
		{
			"address":    "罗湖区",
			"start_year": "1990",
			"end_year":   "",
		},
	},
	440304: {
		{
			"address":    "福田区",
			"start_year": "1990",
			"end_year":   "",
		},
	},
	440305: {
		{
			"address":    "南山区",
			"start_year": "1990",
			"end_year":   "",
		},
	},
	440306: {
		{
			"address":    "宝安区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	440307: {
		{
			"address":    "龙岗区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	440308: {
		{
			"address":    "盐田区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	440309: {
		{
			"address":    "龙华区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	440310: {
		{
			"address":    "坪山区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	440311: {
		{
			"address":    "光明区",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	440321: {
		{
			"address":    "宝安县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	440400: {
		{
			"address":    "珠海市",
			"start_year": "",
			"end_year":   "",
		},
	},
	440402: {
		{
			"address":    "香洲区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	440403: {
		{
			"address":    "斗门区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	440404: {
		{
			"address":    "金湾区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	440421: {
		{
			"address":    "斗门县",
			"start_year": "1983",
			"end_year":   "2000",
		},
	},
	440500: {
		{
			"address":    "汕头市",
			"start_year": "",
			"end_year":   "",
		},
	},
	440502: {
		{
			"address":    "同平区",
			"start_year": "",
			"end_year":   "1990",
		},
		{
			"address":    "龙湖区",
			"start_year": "1991",
			"end_year":   "1993",
		},
	},
	440503: {
		{
			"address":    "安平区",
			"start_year": "",
			"end_year":   "1990",
		},
		{
			"address":    "金园区",
			"start_year": "1991",
			"end_year":   "1993",
		},
	},
	440504: {
		{
			"address":    "公园区",
			"start_year": "",
			"end_year":   "1990",
		},
		{
			"address":    "升平区",
			"start_year": "1991",
			"end_year":   "1993",
		},
	},
	440505: {
		{
			"address":    "金沙区",
			"start_year": "",
			"end_year":   "1990",
		},
	},
	440506: {
		{
			"address":    "达豪区",
			"start_year": "1984",
			"end_year":   "2002",
		},
	},
	440507: {
		{
			"address":    "龙湖区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	440508: {
		{
			"address":    "金园区",
			"start_year": "1994",
			"end_year":   "2002",
		},
	},
	440509: {
		{
			"address":    "升平区",
			"start_year": "1994",
			"end_year":   "2002",
		},
	},
	440510: {
		{
			"address":    "河浦区",
			"start_year": "1994",
			"end_year":   "2002",
		},
	},
	440511: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1990",
		},
		{
			"address":    "金平区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	440512: {
		{
			"address":    "濠江区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	440513: {
		{
			"address":    "潮阳区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	440514: {
		{
			"address":    "潮南区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	440515: {
		{
			"address":    "澄海区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	440521: {
		{
			"address":    "澄海县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	440522: {
		{
			"address":    "饶平县",
			"start_year": "1983",
			"end_year":   "1990",
		},
	},
	440523: {
		{
			"address":    "南澳县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	440524: {
		{
			"address":    "潮阳县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	440525: {
		{
			"address":    "揭阳县",
			"start_year": "1983",
			"end_year":   "1990",
		},
	},
	440526: {
		{
			"address":    "揭西县",
			"start_year": "1983",
			"end_year":   "1990",
		},
	},
	440527: {
		{
			"address":    "普宁县",
			"start_year": "1983",
			"end_year":   "1990",
		},
	},
	440528: {
		{
			"address":    "惠来县",
			"start_year": "1983",
			"end_year":   "1990",
		},
	},
	440582: {
		{
			"address":    "潮阳市",
			"start_year": "1995",
			"end_year":   "2002",
		},
	},
	440583: {
		{
			"address":    "澄海市",
			"start_year": "1995",
			"end_year":   "2002",
		},
	},
	440600: {
		{
			"address":    "佛山市",
			"start_year": "",
			"end_year":   "",
		},
	},
	440602: {
		{
			"address":    "汾江区",
			"start_year": "1984",
			"end_year":   "1986",
		},
		{
			"address":    "城区",
			"start_year": "1987",
			"end_year":   "2001",
		},
	},
	440603: {
		{
			"address":    "石湾区",
			"start_year": "1984",
			"end_year":   "2001",
		},
	},
	440604: {
		{
			"address":    "禅城区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	440605: {
		{
			"address":    "南海区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	440606: {
		{
			"address":    "顺德区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	440607: {
		{
			"address":    "三水区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	440608: {
		{
			"address":    "高明区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	440620: {
		{
			"address":    "中山市",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	440621: {
		{
			"address":    "三水县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	440622: {
		{
			"address":    "南海县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	440623: {
		{
			"address":    "顺德县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	440624: {
		{
			"address":    "高明县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	440681: {
		{
			"address":    "顺德市",
			"start_year": "1995",
			"end_year":   "2001",
		},
	},
	440682: {
		{
			"address":    "南海市",
			"start_year": "1995",
			"end_year":   "2001",
		},
	},
	440683: {
		{
			"address":    "三水市",
			"start_year": "1995",
			"end_year":   "2001",
		},
	},
	440684: {
		{
			"address":    "高明市",
			"start_year": "1995",
			"end_year":   "2001",
		},
	},
	440700: {
		{
			"address":    "江门市",
			"start_year": "",
			"end_year":   "",
		},
	},
	440702: {
		{
			"address":    "城区",
			"start_year": "1984",
			"end_year":   "1993",
		},
	},
	440703: {
		{
			"address":    "蓬江区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	440704: {
		{
			"address":    "江海区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	440705: {
		{
			"address":    "新会区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	440711: {
		{
			"address":    "郊区",
			"start_year": "1984",
			"end_year":   "1993",
		},
	},
	440721: {
		{
			"address":    "新会县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	440722: {
		{
			"address":    "台山县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	440723: {
		{
			"address":    "恩平县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	440724: {
		{
			"address":    "开平县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	440725: {
		{
			"address":    "鹤山县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	440726: {
		{
			"address":    "阳江县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	440727: {
		{
			"address":    "阳春县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	440781: {
		{
			"address":    "台山市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	440782: {
		{
			"address":    "新会市",
			"start_year": "1995",
			"end_year":   "2001",
		},
	},
	440783: {
		{
			"address":    "开平市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	440784: {
		{
			"address":    "鹤山市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	440785: {
		{
			"address":    "恩平市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	440800: {
		{
			"address":    "湛江市",
			"start_year": "",
			"end_year":   "",
		},
	},
	440802: {
		{
			"address":    "赤坎区",
			"start_year": "",
			"end_year":   "",
		},
	},
	440803: {
		{
			"address":    "霞山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	440804: {
		{
			"address":    "坡头区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	440811: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1993",
		},
		{
			"address":    "麻章区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	440821: {
		{
			"address":    "吴川县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	440822: {
		{
			"address":    "廉江县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	440823: {
		{
			"address":    "遂溪县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	440824: {
		{
			"address":    "海康县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	440825: {
		{
			"address":    "徐闻县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	440881: {
		{
			"address":    "廉江市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	440882: {
		{
			"address":    "雷州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	440883: {
		{
			"address":    "吴川市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	440900: {
		{
			"address":    "茂名市",
			"start_year": "",
			"end_year":   "",
		},
	},
	440902: {
		{
			"address":    "茂南区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	440903: {
		{
			"address":    "茂港区",
			"start_year": "2001",
			"end_year":   "2013",
		},
	},
	440904: {
		{
			"address":    "电白区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	440921: {
		{
			"address":    "信宜县",
			"start_year": "1983",
			"end_year":   "1994",
		},
	},
	440922: {
		{
			"address":    "高州县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	440923: {
		{
			"address":    "电白县",
			"start_year": "1983",
			"end_year":   "2013",
		},
	},
	440924: {
		{
			"address":    "化州县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	440981: {
		{
			"address":    "高州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	440982: {
		{
			"address":    "化州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	440983: {
		{
			"address":    "信宜市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	441000: {
		{
			"address":    "海口市",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	441001: {
		{
			"address":    "潮州市",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	441002: {
		{
			"address":    "新华区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	441003: {
		{
			"address":    "立新区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	441004: {
		{
			"address":    "东方红区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	441005: {
		{
			"address":    "秀英区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	441100: {
		{
			"address":    "三亚市",
			"start_year": "1987",
			"end_year":   "1987",
		},
	},
	441200: {
		{
			"address":    "肇庆市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441202: {
		{
			"address":    "端州区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441203: {
		{
			"address":    "鼎湖区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441204: {
		{
			"address":    "高要区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	441221: {
		{
			"address":    "高要县",
			"start_year": "1988",
			"end_year":   "1992",
		},
	},
	441222: {
		{
			"address":    "四会县",
			"start_year": "1988",
			"end_year":   "1992",
		},
	},
	441223: {
		{
			"address":    "广宁县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441224: {
		{
			"address":    "怀集县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441225: {
		{
			"address":    "封开县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441226: {
		{
			"address":    "德庆县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441227: {
		{
			"address":    "云浮县",
			"start_year": "1988",
			"end_year":   "1991",
		},
	},
	441228: {
		{
			"address":    "新兴县",
			"start_year": "1988",
			"end_year":   "1993",
		},
	},
	441229: {
		{
			"address":    "郁南县",
			"start_year": "1988",
			"end_year":   "1993",
		},
	},
	441230: {
		{
			"address":    "罗定县",
			"start_year": "1988",
			"end_year":   "1992",
		},
	},
	441283: {
		{
			"address":    "高要市",
			"start_year": "1995",
			"end_year":   "2014",
		},
	},
	441284: {
		{
			"address":    "四会市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	441300: {
		{
			"address":    "惠州市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441302: {
		{
			"address":    "惠城区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441303: {
		{
			"address":    "惠阳区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	441321: {
		{
			"address":    "惠阳县",
			"start_year": "1988",
			"end_year":   "1993",
		},
	},
	441322: {
		{
			"address":    "博罗县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441323: {
		{
			"address":    "惠东县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441324: {
		{
			"address":    "龙门县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441381: {
		{
			"address":    "惠阳市",
			"start_year": "1995",
			"end_year":   "2002",
		},
	},
	441400: {
		{
			"address":    "梅州市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441402: {
		{
			"address":    "梅江区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441403: {
		{
			"address":    "梅县区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	441421: {
		{
			"address":    "梅县",
			"start_year": "1988",
			"end_year":   "2012",
		},
	},
	441422: {
		{
			"address":    "大埔县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441423: {
		{
			"address":    "丰顺县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441424: {
		{
			"address":    "五华县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441425: {
		{
			"address":    "兴宁县",
			"start_year": "1988",
			"end_year":   "1993",
		},
	},
	441426: {
		{
			"address":    "平远县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441427: {
		{
			"address":    "蕉岭县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441481: {
		{
			"address":    "兴宁市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	441500: {
		{
			"address":    "汕尾市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441502: {
		{
			"address":    "城区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441521: {
		{
			"address":    "海丰县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441522: {
		{
			"address":    "陆丰县",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	441523: {
		{
			"address":    "陆河县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441581: {
		{
			"address":    "陆丰市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	441600: {
		{
			"address":    "河源市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441602: {
		{
			"address":    "源城区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441611: {
		{
			"address":    "郊区",
			"start_year": "1988",
			"end_year":   "1992",
		},
	},
	441621: {
		{
			"address":    "紫金县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441622: {
		{
			"address":    "龙川县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441623: {
		{
			"address":    "连平县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441624: {
		{
			"address":    "和平县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441625: {
		{
			"address":    "东源县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	441700: {
		{
			"address":    "阳江市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441702: {
		{
			"address":    "江城区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441703: {
		{
			"address":    "阳东区",
			"start_year": "1988",
			"end_year":   "1990",
		},
	},
	441704: {
		{
			"address":    "阳东区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	441721: {
		{
			"address":    "阳西县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441722: {
		{
			"address":    "阳春县",
			"start_year": "1988",
			"end_year":   "1993",
		},
	},
	441723: {
		{
			"address":    "阳东县",
			"start_year": "1991",
			"end_year":   "2013",
		},
	},
	441781: {
		{
			"address":    "阳春市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	441800: {
		{
			"address":    "清远市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441802: {
		{
			"address":    "清城区",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441803: {
		{
			"address":    "清新区",
			"start_year": "2012",
			"end_year":   "",
		},
	},
	441811: {
		{
			"address":    "清郊区",
			"start_year": "1988",
			"end_year":   "1991",
		},
	},
	441821: {
		{
			"address":    "佛冈县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441822: {
		{
			"address":    "英德县",
			"start_year": "1988",
			"end_year":   "1993",
		},
	},
	441823: {
		{
			"address":    "阳山县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441824: {
		{
			"address":    "连县",
			"start_year": "1988",
			"end_year":   "1993",
		},
	},
	441825: {
		{
			"address":    "连山壮族瑶族自治县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441826: {
		{
			"address":    "连南瑶族自治县",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	441827: {
		{
			"address":    "清新县",
			"start_year": "1992",
			"end_year":   "2011",
		},
	},
	441881: {
		{
			"address":    "英德市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	441882: {
		{
			"address":    "连州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	441900: {
		{
			"address":    "东莞市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	442000: {
		{
			"address":    "中山市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	442100: {
		{
			"address":    "海南行政区",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442101: {
		{
			"address":    "海口市",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	442102: {
		{
			"address":    "通什市",
			"start_year": "1987",
			"end_year":   "1987",
		},
	},
	442121: {
		{
			"address":    "琼山县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442122: {
		{
			"address":    "文昌县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442123: {
		{
			"address":    "琼海县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442124: {
		{
			"address":    "万宁县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442125: {
		{
			"address":    "定安县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442126: {
		{
			"address":    "屯昌县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442127: {
		{
			"address":    "澄迈县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442128: {
		{
			"address":    "临高县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442129: {
		{
			"address":    "儋县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442130: {
		{
			"address":    "东方黎族自治县",
			"start_year": "1987",
			"end_year":   "1987",
		},
	},
	442131: {
		{
			"address":    "乐东黎族自治县",
			"start_year": "1987",
			"end_year":   "1987",
		},
	},
	442132: {
		{
			"address":    "琼中黎族苗族自治县",
			"start_year": "1987",
			"end_year":   "1987",
		},
	},
	442133: {
		{
			"address":    "保亭黎族苗族自治县",
			"start_year": "1987",
			"end_year":   "1987",
		},
	},
	442134: {
		{
			"address":    "陵水黎族自治县",
			"start_year": "1987",
			"end_year":   "1987",
		},
	},
	442135: {
		{
			"address":    "白沙黎族自治县",
			"start_year": "1987",
			"end_year":   "1987",
		},
	},
	442136: {
		{
			"address":    "昌江黎族自治县",
			"start_year": "1987",
			"end_year":   "1987",
		},
	},
	442200: {
		{
			"address":    "海南黎族苗族自治州",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	442201: {
		{
			"address":    "三亚市",
			"start_year": "1984",
			"end_year":   "1986",
		},
	},
	442202: {
		{
			"address":    "通什市",
			"start_year": "1986",
			"end_year":   "1986",
		},
	},
	442221: {
		{
			"address":    "崖县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	442222: {
		{
			"address":    "东方县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	442223: {
		{
			"address":    "乐东县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	442224: {
		{
			"address":    "琼中县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	442225: {
		{
			"address":    "保亭县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	442226: {
		{
			"address":    "陵水县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	442227: {
		{
			"address":    "白沙县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	442228: {
		{
			"address":    "昌江县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	442300: {
		{
			"address":    "汕头地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442301: {
		{
			"address":    "潮州市",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	442302: {
		{
			"address":    "潮州市",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	442321: {
		{
			"address":    "潮安县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442322: {
		{
			"address":    "澄海县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442323: {
		{
			"address":    "饶平县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442324: {
		{
			"address":    "南澳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442325: {
		{
			"address":    "潮阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442326: {
		{
			"address":    "揭阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442327: {
		{
			"address":    "揭西县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442328: {
		{
			"address":    "普宁县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442329: {
		{
			"address":    "惠来县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442330: {
		{
			"address":    "陆丰县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442331: {
		{
			"address":    "海丰县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442400: {
		{
			"address":    "梅县地区",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442401: {
		{
			"address":    "梅州市",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "梅县市",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	442421: {
		{
			"address":    "梅县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442422: {
		{
			"address":    "大埔县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442423: {
		{
			"address":    "丰顺县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442424: {
		{
			"address":    "五华县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442425: {
		{
			"address":    "兴宁县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442426: {
		{
			"address":    "平远县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442427: {
		{
			"address":    "蕉岭县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442500: {
		{
			"address":    "惠阳地区",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442501: {
		{
			"address":    "惠州市",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442502: {
		{
			"address":    "东莞市",
			"start_year": "1985",
			"end_year":   "1987",
		},
	},
	442521: {
		{
			"address":    "惠阳县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442522: {
		{
			"address":    "紫金县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442523: {
		{
			"address":    "和平县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442524: {
		{
			"address":    "连平县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442525: {
		{
			"address":    "河源县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442526: {
		{
			"address":    "博罗县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442527: {
		{
			"address":    "东莞县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	442528: {
		{
			"address":    "惠东县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442529: {
		{
			"address":    "龙川县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442530: {
		{
			"address":    "陆丰县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	442531: {
		{
			"address":    "海丰县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	442600: {
		{
			"address":    "韶关地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442621: {
		{
			"address":    "始兴县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "三水县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	442622: {
		{
			"address":    "南雄县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "南海县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	442623: {
		{
			"address":    "仁化县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "顺德县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	442624: {
		{
			"address":    "乐昌县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "中山县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	442625: {
		{
			"address":    "连县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "斗门县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	442626: {
		{
			"address":    "阳山县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "新会县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	442627: {
		{
			"address":    "英德县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "台山县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	442628: {
		{
			"address":    "清远县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "恩平县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	442629: {
		{
			"address":    "佛冈县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "开平县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	442630: {
		{
			"address":    "翁源县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442631: {
		{
			"address":    "连山壮族瑶族自治县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "鹤山县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	442632: {
		{
			"address":    "连南瑶族自治县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "高明县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	442633: {
		{
			"address":    "乳源瑶族自治县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442700: {
		{
			"address":    "佛山地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442721: {
		{
			"address":    "三水县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	442722: {
		{
			"address":    "南海县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	442723: {
		{
			"address":    "顺德县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	442724: {
		{
			"address":    "中山县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	442725: {
		{
			"address":    "斗门县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	442726: {
		{
			"address":    "新会县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	442727: {
		{
			"address":    "台山县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	442728: {
		{
			"address":    "恩平县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	442729: {
		{
			"address":    "开平县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	442730: {
		{
			"address":    "高鹤县",
			"start_year": "",
			"end_year":   "1980",
		},
	},
	442731: {
		{
			"address":    "鹤山县",
			"start_year": "1981",
			"end_year":   "1981",
		},
	},
	442732: {
		{
			"address":    "高明县",
			"start_year": "1981",
			"end_year":   "1981",
		},
	},
	442800: {
		{
			"address":    "肇庆地区",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442801: {
		{
			"address":    "肇庆市",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442821: {
		{
			"address":    "高要县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442822: {
		{
			"address":    "四会县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442823: {
		{
			"address":    "广宁县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442824: {
		{
			"address":    "怀集县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442825: {
		{
			"address":    "封开县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442826: {
		{
			"address":    "德庆县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442827: {
		{
			"address":    "云浮县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442828: {
		{
			"address":    "新兴县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442829: {
		{
			"address":    "郁南县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442830: {
		{
			"address":    "罗定县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	442900: {
		{
			"address":    "湛江地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442921: {
		{
			"address":    "阳江县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442922: {
		{
			"address":    "阳春县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442923: {
		{
			"address":    "信宜县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442924: {
		{
			"address":    "高州县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442925: {
		{
			"address":    "电白县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442926: {
		{
			"address":    "吴川县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442927: {
		{
			"address":    "化州县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442928: {
		{
			"address":    "廉江县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442929: {
		{
			"address":    "遂溪县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442930: {
		{
			"address":    "海康县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	442931: {
		{
			"address":    "徐闻县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	445100: {
		{
			"address":    "潮州市",
			"start_year": "1991",
			"end_year":   "",
		},
	},
	445102: {
		{
			"address":    "湘桥区",
			"start_year": "1991",
			"end_year":   "",
		},
	},
	445103: {
		{
			"address":    "潮安区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	445121: {
		{
			"address":    "潮安县",
			"start_year": "1991",
			"end_year":   "2012",
		},
	},
	445122: {
		{
			"address":    "饶平县",
			"start_year": "1991",
			"end_year":   "",
		},
	},
	445200: {
		{
			"address":    "揭阳市",
			"start_year": "1991",
			"end_year":   "",
		},
	},
	445202: {
		{
			"address":    "榕城区",
			"start_year": "1991",
			"end_year":   "",
		},
	},
	445203: {
		{
			"address":    "揭东区",
			"start_year": "2012",
			"end_year":   "",
		},
	},
	445221: {
		{
			"address":    "揭东县",
			"start_year": "1991",
			"end_year":   "2011",
		},
	},
	445222: {
		{
			"address":    "揭西县",
			"start_year": "1991",
			"end_year":   "",
		},
	},
	445223: {
		{
			"address":    "普宁县",
			"start_year": "1991",
			"end_year":   "1992",
		},
	},
	445224: {
		{
			"address":    "惠来县",
			"start_year": "1991",
			"end_year":   "",
		},
	},
	445281: {
		{
			"address":    "普宁市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	445300: {
		{
			"address":    "云浮市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	445302: {
		{
			"address":    "云城区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	445303: {
		{
			"address":    "云安区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	445321: {
		{
			"address":    "新兴县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	445322: {
		{
			"address":    "郁南县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	445323: {
		{
			"address":    "云安县",
			"start_year": "1996",
			"end_year":   "2013",
		},
	},
	445381: {
		{
			"address":    "罗定市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	449001: {
		{
			"address":    "潮州市",
			"start_year": "1984",
			"end_year":   "1990",
		},
		{
			"address":    "顺德市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	449002: {
		{
			"address":    "中山市",
			"start_year": "1984",
			"end_year":   "1987",
		},
		{
			"address":    "台山市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	449003: {
		{
			"address":    "海口市",
			"start_year": "1986",
			"end_year":   "1986",
		},
		{
			"address":    "番禺市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	449004: {
		{
			"address":    "南海市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	449005: {
		{
			"address":    "云浮市",
			"start_year": "1992",
			"end_year":   "1993",
		},
	},
	449006: {
		{
			"address":    "新会市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	449007: {
		{
			"address":    "开平市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	449008: {
		{
			"address":    "三水市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	449009: {
		{
			"address":    "普宁市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	449010: {
		{
			"address":    "罗定市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	449011: {
		{
			"address":    "潮阳市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	449012: {
		{
			"address":    "高州市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	449013: {
		{
			"address":    "花都市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	449014: {
		{
			"address":    "高要市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	449015: {
		{
			"address":    "鹤山市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	449016: {
		{
			"address":    "四会市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	449017: {
		{
			"address":    "增城市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	449018: {
		{
			"address":    "廉江市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	449019: {
		{
			"address":    "英德市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	449020: {
		{
			"address":    "恩平市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	449021: {
		{
			"address":    "从化市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	449022: {
		{
			"address":    "澄海市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	449023: {
		{
			"address":    "高明市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	449024: {
		{
			"address":    "连州市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	449025: {
		{
			"address":    "雷州市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	449026: {
		{
			"address":    "乐昌市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	449027: {
		{
			"address":    "阳春市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	449028: {
		{
			"address":    "惠阳市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	449029: {
		{
			"address":    "吴川市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	449030: {
		{
			"address":    "兴宁市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	449031: {
		{
			"address":    "化州市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	450000: {
		{
			"address":    "广西壮族自治区",
			"start_year": "",
			"end_year":   "",
		},
	},
	450100: {
		{
			"address":    "南宁市",
			"start_year": "",
			"end_year":   "",
		},
	},
	450102: {
		{
			"address":    "兴宁区",
			"start_year": "",
			"end_year":   "",
		},
	},
	450103: {
		{
			"address":    "新城区",
			"start_year": "",
			"end_year":   "2003",
		},
		{
			"address":    "青秀区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	450104: {
		{
			"address":    "城北区",
			"start_year": "",
			"end_year":   "2003",
		},
	},
	450105: {
		{
			"address":    "江南区",
			"start_year": "",
			"end_year":   "",
		},
	},
	450106: {
		{
			"address":    "永新区",
			"start_year": "",
			"end_year":   "2003",
		},
	},
	450107: {
		{
			"address":    "市郊区",
			"start_year": "1984",
			"end_year":   "2000",
		},
		{
			"address":    "西乡塘区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	450108: {
		{
			"address":    "良庆区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	450109: {
		{
			"address":    "邕宁区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	450110: {
		{
			"address":    "武鸣区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	450121: {
		{
			"address":    "邕宁县",
			"start_year": "1983",
			"end_year":   "2003",
		},
	},
	450122: {
		{
			"address":    "武鸣县",
			"start_year": "1983",
			"end_year":   "2014",
		},
	},
	450123: {
		{
			"address":    "隆安县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	450124: {
		{
			"address":    "马山县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	450125: {
		{
			"address":    "上林县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	450126: {
		{
			"address":    "宾阳县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	450127: {
		{
			"address":    "横县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	450200: {
		{
			"address":    "柳州市",
			"start_year": "",
			"end_year":   "",
		},
	},
	450202: {
		{
			"address":    "城中区",
			"start_year": "",
			"end_year":   "",
		},
	},
	450203: {
		{
			"address":    "鱼峰区",
			"start_year": "",
			"end_year":   "",
		},
	},
	450204: {
		{
			"address":    "柳南区",
			"start_year": "",
			"end_year":   "",
		},
	},
	450205: {
		{
			"address":    "柳北区",
			"start_year": "",
			"end_year":   "",
		},
	},
	450206: {
		{
			"address":    "市郊区",
			"start_year": "1984",
			"end_year":   "2001",
		},
		{
			"address":    "柳江区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	450221: {
		{
			"address":    "柳江县",
			"start_year": "1983",
			"end_year":   "2015",
		},
	},
	450222: {
		{
			"address":    "柳城县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	450223: {
		{
			"address":    "鹿寨县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	450224: {
		{
			"address":    "融安县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	450225: {
		{
			"address":    "融水苗族自治县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	450226: {
		{
			"address":    "三江侗族自治县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	450300: {
		{
			"address":    "桂林市",
			"start_year": "",
			"end_year":   "",
		},
	},
	450302: {
		{
			"address":    "秀峰区",
			"start_year": "",
			"end_year":   "",
		},
	},
	450303: {
		{
			"address":    "叠彩区",
			"start_year": "",
			"end_year":   "",
		},
	},
	450304: {
		{
			"address":    "象山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	450305: {
		{
			"address":    "七星区",
			"start_year": "",
			"end_year":   "",
		},
	},
	450306: {
		{
			"address":    "市郊区",
			"start_year": "1984",
			"end_year":   "1995",
		},
	},
	450311: {
		{
			"address":    "雁山区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	450312: {
		{
			"address":    "临桂区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	450321: {
		{
			"address":    "阳朔县",
			"start_year": "1981",
			"end_year":   "",
		},
	},
	450322: {
		{
			"address":    "临桂县",
			"start_year": "1983",
			"end_year":   "2012",
		},
	},
	450323: {
		{
			"address":    "灵川县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	450324: {
		{
			"address":    "全州县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	450325: {
		{
			"address":    "兴安县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	450326: {
		{
			"address":    "永福县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	450327: {
		{
			"address":    "灌阳县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	450328: {
		{
			"address":    "龙胜各族自治县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	450329: {
		{
			"address":    "资源县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	450330: {
		{
			"address":    "平乐县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	450331: {
		{
			"address":    "荔浦县",
			"start_year": "1998",
			"end_year":   "2017",
		},
	},
	450332: {
		{
			"address":    "恭城瑶族自治县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	450381: {
		{
			"address":    "荔浦市",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	450400: {
		{
			"address":    "梧州市",
			"start_year": "",
			"end_year":   "",
		},
	},
	450402: {
		{
			"address":    "白云区",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	450403: {
		{
			"address":    "万秀区",
			"start_year": "",
			"end_year":   "",
		},
	},
	450404: {
		{
			"address":    "蝶山区",
			"start_year": "",
			"end_year":   "2012",
		},
	},
	450405: {
		{
			"address":    "鸳江区",
			"start_year": "",
			"end_year":   "1983",
		},
		{
			"address":    "市郊区",
			"start_year": "1984",
			"end_year":   "2002",
		},
		{
			"address":    "长洲区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	450406: {
		{
			"address":    "龙圩区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	450421: {
		{
			"address":    "苍梧县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	450422: {
		{
			"address":    "藤县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	450423: {
		{
			"address":    "蒙山县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	450481: {
		{
			"address":    "岑溪市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	450500: {
		{
			"address":    "北海市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	450502: {
		{
			"address":    "海城区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	450503: {
		{
			"address":    "市郊区",
			"start_year": "1984",
			"end_year":   "1993",
		},
		{
			"address":    "银海区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	450512: {
		{
			"address":    "铁山港区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	450521: {
		{
			"address":    "合浦县",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	450600: {
		{
			"address":    "防城港市",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	450602: {
		{
			"address":    "港口区",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	450603: {
		{
			"address":    "防城区",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	450621: {
		{
			"address":    "上思县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	450681: {
		{
			"address":    "东兴市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	450700: {
		{
			"address":    "钦州市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	450702: {
		{
			"address":    "钦南区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	450703: {
		{
			"address":    "钦北区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	450721: {
		{
			"address":    "灵山县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	450722: {
		{
			"address":    "浦北县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	450800: {
		{
			"address":    "贵港市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	450802: {
		{
			"address":    "港北区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	450803: {
		{
			"address":    "港南区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	450804: {
		{
			"address":    "覃塘区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	450821: {
		{
			"address":    "平南县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	450881: {
		{
			"address":    "桂平市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	450900: {
		{
			"address":    "玉林市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	450902: {
		{
			"address":    "玉州区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	450903: {
		{
			"address":    "福绵区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	450921: {
		{
			"address":    "容县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	450922: {
		{
			"address":    "陆川县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	450923: {
		{
			"address":    "博白县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	450924: {
		{
			"address":    "兴业县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	450981: {
		{
			"address":    "北流市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	451000: {
		{
			"address":    "百色市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451002: {
		{
			"address":    "右江区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451003: {
		{
			"address":    "田阳区",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	451021: {
		{
			"address":    "田阳县",
			"start_year": "2002",
			"end_year":   "2018",
		},
	},
	451022: {
		{
			"address":    "田东县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451023: {
		{
			"address":    "平果县",
			"start_year": "2002",
			"end_year":   "2018",
		},
	},
	451024: {
		{
			"address":    "德保县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451025: {
		{
			"address":    "靖西县",
			"start_year": "2002",
			"end_year":   "2014",
		},
	},
	451026: {
		{
			"address":    "那坡县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451027: {
		{
			"address":    "凌云县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451028: {
		{
			"address":    "乐业县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451029: {
		{
			"address":    "田林县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451030: {
		{
			"address":    "西林县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451031: {
		{
			"address":    "隆林各族自治县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451081: {
		{
			"address":    "靖西市",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	451082: {
		{
			"address":    "平果市",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	451100: {
		{
			"address":    "贺州市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451102: {
		{
			"address":    "八步区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451103: {
		{
			"address":    "平桂区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	451121: {
		{
			"address":    "昭平县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451122: {
		{
			"address":    "钟山县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451123: {
		{
			"address":    "富川瑶族自治县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451200: {
		{
			"address":    "河池市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451202: {
		{
			"address":    "金城江区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451203: {
		{
			"address":    "宜州区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	451221: {
		{
			"address":    "南丹县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451222: {
		{
			"address":    "天峨县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451223: {
		{
			"address":    "凤山县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451224: {
		{
			"address":    "东兰县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451225: {
		{
			"address":    "罗城仫佬族自治县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451226: {
		{
			"address":    "环江毛南族自治县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451227: {
		{
			"address":    "巴马瑶族自治县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451228: {
		{
			"address":    "都安瑶族自治县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451229: {
		{
			"address":    "大化瑶族自治县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451281: {
		{
			"address":    "宜州市",
			"start_year": "2002",
			"end_year":   "2015",
		},
	},
	451300: {
		{
			"address":    "来宾市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451302: {
		{
			"address":    "兴宾区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451321: {
		{
			"address":    "忻城县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451322: {
		{
			"address":    "象州县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451323: {
		{
			"address":    "武宣县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451324: {
		{
			"address":    "金秀瑶族自治县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451381: {
		{
			"address":    "合山市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451400: {
		{
			"address":    "崇左市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451402: {
		{
			"address":    "江州区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451421: {
		{
			"address":    "扶绥县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451422: {
		{
			"address":    "宁明县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451423: {
		{
			"address":    "龙州县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451424: {
		{
			"address":    "大新县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451425: {
		{
			"address":    "天等县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	451481: {
		{
			"address":    "凭祥市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	452100: {
		{
			"address":    "南宁地区",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452101: {
		{
			"address":    "凭祥市",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452121: {
		{
			"address":    "邕宁县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	452122: {
		{
			"address":    "横县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452123: {
		{
			"address":    "宾阳县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452124: {
		{
			"address":    "上林县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452125: {
		{
			"address":    "武鸣县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	452126: {
		{
			"address":    "隆安县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452127: {
		{
			"address":    "马山县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452128: {
		{
			"address":    "扶绥县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452129: {
		{
			"address":    "崇左县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452130: {
		{
			"address":    "大新县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452131: {
		{
			"address":    "天等县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452132: {
		{
			"address":    "宁明县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452133: {
		{
			"address":    "龙州县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452200: {
		{
			"address":    "柳州地区",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452201: {
		{
			"address":    "合山市",
			"start_year": "1981",
			"end_year":   "2001",
		},
	},
	452221: {
		{
			"address":    "柳江县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	452222: {
		{
			"address":    "柳城县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	452223: {
		{
			"address":    "鹿寨县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452224: {
		{
			"address":    "象州县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452225: {
		{
			"address":    "武宣县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452226: {
		{
			"address":    "来宾县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452227: {
		{
			"address":    "融安县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452228: {
		{
			"address":    "三江侗族自治县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452229: {
		{
			"address":    "融水苗族自治县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452230: {
		{
			"address":    "金秀瑶族自治县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452231: {
		{
			"address":    "忻城县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452300: {
		{
			"address":    "桂林地区",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	452321: {
		{
			"address":    "临桂县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	452322: {
		{
			"address":    "灵川县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	452323: {
		{
			"address":    "全州县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	452324: {
		{
			"address":    "兴安县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	452325: {
		{
			"address":    "永福县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	452326: {
		{
			"address":    "阳朔县",
			"start_year": "",
			"end_year":   "1980",
		},
	},
	452327: {
		{
			"address":    "灌阳县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	452328: {
		{
			"address":    "龙胜各族自治县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	452329: {
		{
			"address":    "资源县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	452330: {
		{
			"address":    "平乐县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	452331: {
		{
			"address":    "荔浦县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	452332: {
		{
			"address":    "恭城县",
			"start_year": "",
			"end_year":   "1989",
		},
		{
			"address":    "恭城瑶族自治县",
			"start_year": "1990",
			"end_year":   "1997",
		},
	},
	452400: {
		{
			"address":    "梧州地区",
			"start_year": "",
			"end_year":   "1996",
		},
		{
			"address":    "贺州地区",
			"start_year": "1997",
			"end_year":   "2001",
		},
	},
	452401: {
		{
			"address":    "岑溪市",
			"start_year": "1995",
			"end_year":   "1996",
		},
	},
	452402: {
		{
			"address":    "贺州市",
			"start_year": "1997",
			"end_year":   "2001",
		},
	},
	452421: {
		{
			"address":    "岑溪县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	452422: {
		{
			"address":    "苍梧县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	452423: {
		{
			"address":    "藤县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	452424: {
		{
			"address":    "昭平县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452425: {
		{
			"address":    "蒙山县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	452426: {
		{
			"address":    "贺县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	452427: {
		{
			"address":    "钟山县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452428: {
		{
			"address":    "富川县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "富川瑶族自治县",
			"start_year": "1983",
			"end_year":   "2001",
		},
	},
	452500: {
		{
			"address":    "玉林地区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	452501: {
		{
			"address":    "玉林市",
			"start_year": "1983",
			"end_year":   "1996",
		},
	},
	452502: {
		{
			"address":    "贵港市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	452503: {
		{
			"address":    "桂平市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	452504: {
		{
			"address":    "北流市",
			"start_year": "1994",
			"end_year":   "1996",
		},
	},
	452521: {
		{
			"address":    "玉林县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	452522: {
		{
			"address":    "贵县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	452523: {
		{
			"address":    "桂平县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	452524: {
		{
			"address":    "平南县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	452525: {
		{
			"address":    "容县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	452526: {
		{
			"address":    "北流县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	452527: {
		{
			"address":    "陆川县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	452528: {
		{
			"address":    "博白县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	452600: {
		{
			"address":    "百色地区",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452601: {
		{
			"address":    "百色市",
			"start_year": "1983",
			"end_year":   "2001",
		},
	},
	452621: {
		{
			"address":    "百色县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	452622: {
		{
			"address":    "田阳县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452623: {
		{
			"address":    "田东县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452624: {
		{
			"address":    "平果县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452625: {
		{
			"address":    "德保县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452626: {
		{
			"address":    "靖西县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452627: {
		{
			"address":    "那坡县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452628: {
		{
			"address":    "凌云县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452629: {
		{
			"address":    "乐业县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452630: {
		{
			"address":    "田林县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452631: {
		{
			"address":    "隆林各族自治县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452632: {
		{
			"address":    "西林县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452700: {
		{
			"address":    "河池地区",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452701: {
		{
			"address":    "河池市",
			"start_year": "1983",
			"end_year":   "2001",
		},
	},
	452702: {
		{
			"address":    "宜州市",
			"start_year": "1993",
			"end_year":   "2001",
		},
	},
	452721: {
		{
			"address":    "河池县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	452722: {
		{
			"address":    "宜山县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	452723: {
		{
			"address":    "罗城县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "罗城仫佬族自治县",
			"start_year": "1983",
			"end_year":   "2001",
		},
	},
	452724: {
		{
			"address":    "环江县",
			"start_year": "",
			"end_year":   "1985",
		},
		{
			"address":    "环江毛南族自治县",
			"start_year": "1986",
			"end_year":   "2001",
		},
	},
	452725: {
		{
			"address":    "南丹县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452726: {
		{
			"address":    "天峨县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452727: {
		{
			"address":    "凤山县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452728: {
		{
			"address":    "东兰县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452729: {
		{
			"address":    "巴马瑶族自治县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452730: {
		{
			"address":    "都安瑶族自治县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	452731: {
		{
			"address":    "大化瑶族自治县",
			"start_year": "1987",
			"end_year":   "2001",
		},
	},
	452800: {
		{
			"address":    "钦州地区",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	452801: {
		{
			"address":    "北海市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	452802: {
		{
			"address":    "钦州市",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	452821: {
		{
			"address":    "上思县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	452822: {
		{
			"address":    "防城各族自治县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	452823: {
		{
			"address":    "钦州县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	452824: {
		{
			"address":    "灵山县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	452825: {
		{
			"address":    "合浦县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	452826: {
		{
			"address":    "浦北县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	460000: {
		{
			"address":    "海南省",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	460001: {
		{
			"address":    "通什市",
			"start_year": "1988",
			"end_year":   "2000",
		},
		{
			"address":    "五指山市",
			"start_year": "2001",
			"end_year":   "2002",
		},
	},
	460002: {
		{
			"address":    "琼海市",
			"start_year": "1992",
			"end_year":   "2002",
		},
	},
	460003: {
		{
			"address":    "儋州市",
			"start_year": "1993",
			"end_year":   "2002",
		},
	},
	460004: {
		{
			"address":    "琼山市",
			"start_year": "1994",
			"end_year":   "2001",
		},
	},
	460005: {
		{
			"address":    "文昌市",
			"start_year": "1995",
			"end_year":   "2002",
		},
	},
	460006: {
		{
			"address":    "万宁市",
			"start_year": "1996",
			"end_year":   "2002",
		},
	},
	460007: {
		{
			"address":    "东方市",
			"start_year": "1997",
			"end_year":   "2002",
		},
	},
	460021: {
		{
			"address":    "琼山县",
			"start_year": "1988",
			"end_year":   "1993",
		},
	},
	460022: {
		{
			"address":    "文昌县",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	460023: {
		{
			"address":    "琼海县",
			"start_year": "1988",
			"end_year":   "1991",
		},
	},
	460024: {
		{
			"address":    "万宁县",
			"start_year": "1988",
			"end_year":   "1995",
		},
	},
	460025: {
		{
			"address":    "定安县",
			"start_year": "1988",
			"end_year":   "2002",
		},
	},
	460026: {
		{
			"address":    "屯昌县",
			"start_year": "1988",
			"end_year":   "2002",
		},
	},
	460027: {
		{
			"address":    "澄迈县",
			"start_year": "1988",
			"end_year":   "2002",
		},
	},
	460028: {
		{
			"address":    "临高县",
			"start_year": "1988",
			"end_year":   "2002",
		},
	},
	460029: {
		{
			"address":    "儋县",
			"start_year": "1988",
			"end_year":   "1992",
		},
	},
	460030: {
		{
			"address":    "白沙黎族自治县",
			"start_year": "1988",
			"end_year":   "2002",
		},
	},
	460031: {
		{
			"address":    "昌江黎族自治县",
			"start_year": "1988",
			"end_year":   "2002",
		},
	},
	460032: {
		{
			"address":    "东方黎族自治县",
			"start_year": "1988",
			"end_year":   "1996",
		},
	},
	460033: {
		{
			"address":    "乐东黎族自治县",
			"start_year": "1988",
			"end_year":   "2002",
		},
	},
	460034: {
		{
			"address":    "陵水黎族自治县",
			"start_year": "1988",
			"end_year":   "2002",
		},
	},
	460035: {
		{
			"address":    "保亭黎族苗族自治县",
			"start_year": "1988",
			"end_year":   "2002",
		},
	},
	460036: {
		{
			"address":    "琼中黎族苗族自治县",
			"start_year": "1988",
			"end_year":   "2002",
		},
	},
	460100: {
		{
			"address":    "海口市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	460102: {
		{
			"address":    "振东区",
			"start_year": "1990",
			"end_year":   "2001",
		},
	},
	460103: {
		{
			"address":    "新华区",
			"start_year": "1990",
			"end_year":   "2001",
		},
	},
	460104: {
		{
			"address":    "秀英区",
			"start_year": "1990",
			"end_year":   "2001",
		},
	},
	460105: {
		{
			"address":    "秀英区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	460106: {
		{
			"address":    "龙华区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	460107: {
		{
			"address":    "琼山区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	460108: {
		{
			"address":    "美兰区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	460200: {
		{
			"address":    "三亚市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	460202: {
		{
			"address":    "海棠区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	460203: {
		{
			"address":    "吉阳区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	460204: {
		{
			"address":    "天涯区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	460205: {
		{
			"address":    "崖州区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	460300: {
		{
			"address":    "三沙市",
			"start_year": "2012",
			"end_year":   "",
		},
	},
	460400: {
		{
			"address":    "儋州市",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	469001: {
		{
			"address":    "五指山市",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	469002: {
		{
			"address":    "琼海市",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	469003: {
		{
			"address":    "儋州市",
			"start_year": "2003",
			"end_year":   "2014",
		},
	},
	469005: {
		{
			"address":    "文昌市",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	469006: {
		{
			"address":    "万宁市",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	469007: {
		{
			"address":    "东方市",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	469021: {
		{
			"address":    "定安县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	469022: {
		{
			"address":    "屯昌县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	469023: {
		{
			"address":    "澄迈县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	469024: {
		{
			"address":    "临高县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	469025: {
		{
			"address":    "白沙黎族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	469026: {
		{
			"address":    "昌江黎族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	469027: {
		{
			"address":    "乐东黎族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	469028: {
		{
			"address":    "陵水黎族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	469029: {
		{
			"address":    "保亭黎族苗族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	469030: {
		{
			"address":    "琼中黎族苗族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	500000: {
		{
			"address":    "重庆市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500101: {
		{
			"address":    "万县区",
			"start_year": "1997",
			"end_year":   "1997",
		},
		{
			"address":    "万州区",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	500102: {
		{
			"address":    "涪陵区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500103: {
		{
			"address":    "渝中区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500104: {
		{
			"address":    "大渡口区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500105: {
		{
			"address":    "江北区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500106: {
		{
			"address":    "沙坪坝区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500107: {
		{
			"address":    "九龙坡区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500108: {
		{
			"address":    "南岸区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500109: {
		{
			"address":    "北碚区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500110: {
		{
			"address":    "万盛区",
			"start_year": "1997",
			"end_year":   "2010",
		},
		{
			"address":    "綦江区",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	500111: {
		{
			"address":    "双桥区",
			"start_year": "1997",
			"end_year":   "2010",
		},
		{
			"address":    "大足区",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	500112: {
		{
			"address":    "渝北区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500113: {
		{
			"address":    "巴南区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500114: {
		{
			"address":    "黔江区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	500115: {
		{
			"address":    "长寿区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	500116: {
		{
			"address":    "江津区",
			"start_year": "2006",
			"end_year":   "",
		},
	},
	500117: {
		{
			"address":    "合川区",
			"start_year": "2006",
			"end_year":   "",
		},
	},
	500118: {
		{
			"address":    "永川区",
			"start_year": "2006",
			"end_year":   "",
		},
	},
	500119: {
		{
			"address":    "南川区",
			"start_year": "2006",
			"end_year":   "",
		},
	},
	500120: {
		{
			"address":    "璧山区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	500151: {
		{
			"address":    "铜梁区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	500152: {
		{
			"address":    "潼南区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	500153: {
		{
			"address":    "荣昌区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	500154: {
		{
			"address":    "开州区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	500155: {
		{
			"address":    "梁平区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	500156: {
		{
			"address":    "武隆区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	500221: {
		{
			"address":    "长寿县",
			"start_year": "1997",
			"end_year":   "2000",
		},
	},
	500222: {
		{
			"address":    "綦江县",
			"start_year": "1997",
			"end_year":   "2010",
		},
	},
	500223: {
		{
			"address":    "潼南县",
			"start_year": "1997",
			"end_year":   "2014",
		},
	},
	500224: {
		{
			"address":    "铜梁县",
			"start_year": "1997",
			"end_year":   "2013",
		},
	},
	500225: {
		{
			"address":    "大足县",
			"start_year": "1997",
			"end_year":   "2010",
		},
	},
	500226: {
		{
			"address":    "荣昌县",
			"start_year": "1997",
			"end_year":   "2014",
		},
	},
	500227: {
		{
			"address":    "璧山县",
			"start_year": "1997",
			"end_year":   "2013",
		},
	},
	500228: {
		{
			"address":    "梁平县",
			"start_year": "1997",
			"end_year":   "2015",
		},
	},
	500229: {
		{
			"address":    "城口县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500230: {
		{
			"address":    "丰都县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500231: {
		{
			"address":    "垫江县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500232: {
		{
			"address":    "武隆县",
			"start_year": "1997",
			"end_year":   "2015",
		},
	},
	500233: {
		{
			"address":    "忠县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500234: {
		{
			"address":    "开县",
			"start_year": "1997",
			"end_year":   "2015",
		},
	},
	500235: {
		{
			"address":    "云阳县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500236: {
		{
			"address":    "奉节县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500237: {
		{
			"address":    "巫山县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500238: {
		{
			"address":    "巫溪县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500239: {
		{
			"address":    "黔江土家族苗族自治县",
			"start_year": "1997",
			"end_year":   "1999",
		},
	},
	500240: {
		{
			"address":    "石柱土家族自治县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500241: {
		{
			"address":    "秀山土家族苗族自治县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500242: {
		{
			"address":    "酉阳土家族苗族自治县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500243: {
		{
			"address":    "彭水苗族土家族自治县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	500381: {
		{
			"address":    "江津市",
			"start_year": "1997",
			"end_year":   "2005",
		},
	},
	500382: {
		{
			"address":    "合川市",
			"start_year": "1997",
			"end_year":   "2005",
		},
	},
	500383: {
		{
			"address":    "永川市",
			"start_year": "1997",
			"end_year":   "2005",
		},
	},
	500384: {
		{
			"address":    "南川市",
			"start_year": "1997",
			"end_year":   "2005",
		},
	},
	510000: {
		{
			"address":    "四川省",
			"start_year": "",
			"end_year":   "",
		},
	},
	510100: {
		{
			"address":    "成都市",
			"start_year": "",
			"end_year":   "",
		},
	},
	510102: {
		{
			"address":    "东城区",
			"start_year": "",
			"end_year":   "1989",
		},
	},
	510103: {
		{
			"address":    "西城区",
			"start_year": "",
			"end_year":   "1989",
		},
	},
	510104: {
		{
			"address":    "锦江区",
			"start_year": "1990",
			"end_year":   "",
		},
	},
	510105: {
		{
			"address":    "青羊区",
			"start_year": "1990",
			"end_year":   "",
		},
	},
	510106: {
		{
			"address":    "金牛区",
			"start_year": "1990",
			"end_year":   "",
		},
	},
	510107: {
		{
			"address":    "武侯区",
			"start_year": "1990",
			"end_year":   "",
		},
	},
	510108: {
		{
			"address":    "成华区",
			"start_year": "1990",
			"end_year":   "",
		},
	},
	510111: {
		{
			"address":    "金牛区",
			"start_year": "",
			"end_year":   "1989",
		},
	},
	510112: {
		{
			"address":    "龙泉驿区",
			"start_year": "",
			"end_year":   "",
		},
	},
	510113: {
		{
			"address":    "青白江区",
			"start_year": "",
			"end_year":   "",
		},
	},
	510114: {
		{
			"address":    "新都区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	510115: {
		{
			"address":    "温江区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	510116: {
		{
			"address":    "双流区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	510117: {
		{
			"address":    "郫都区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	510118: {
		{
			"address":    "新津区",
			"start_year": "2020",
			"end_year":   "",
		},
	},
	510121: {
		{
			"address":    "金堂县",
			"start_year": "",
			"end_year":   "",
		},
	},
	510122: {
		{
			"address":    "双流县",
			"start_year": "",
			"end_year":   "2014",
		},
	},
	510123: {
		{
			"address":    "温江县",
			"start_year": "1983",
			"end_year":   "2001",
		},
	},
	510124: {
		{
			"address":    "郫县",
			"start_year": "1983",
			"end_year":   "2015",
		},
	},
	510125: {
		{
			"address":    "新都县",
			"start_year": "1983",
			"end_year":   "2000",
		},
	},
	510126: {
		{
			"address":    "彭县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	510127: {
		{
			"address":    "灌县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	510128: {
		{
			"address":    "崇庆县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	510129: {
		{
			"address":    "大邑县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	510130: {
		{
			"address":    "邛崃县",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	510131: {
		{
			"address":    "蒲江县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	510132: {
		{
			"address":    "新津县",
			"start_year": "1983",
			"end_year":   "2019",
		},
	},
	510181: {
		{
			"address":    "都江堰市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	510182: {
		{
			"address":    "彭州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	510183: {
		{
			"address":    "邛崃市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	510184: {
		{
			"address":    "崇州市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	510185: {
		{
			"address":    "简阳市",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	510200: {
		{
			"address":    "重庆市",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	510202: {
		{
			"address":    "市中区",
			"start_year": "",
			"end_year":   "1993",
		},
		{
			"address":    "渝中区",
			"start_year": "1994",
			"end_year":   "1996",
		},
	},
	510203: {
		{
			"address":    "大渡口区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	510211: {
		{
			"address":    "江北区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	510212: {
		{
			"address":    "沙坪坝区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	510213: {
		{
			"address":    "九龙坡区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	510214: {
		{
			"address":    "南岸区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	510215: {
		{
			"address":    "北碚区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	510216: {
		{
			"address":    "南桐矿区",
			"start_year": "",
			"end_year":   "1992",
		},
		{
			"address":    "万盛区",
			"start_year": "1993",
			"end_year":   "1996",
		},
	},
	510217: {
		{
			"address":    "双桥区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	510218: {
		{
			"address":    "渝北区",
			"start_year": "1994",
			"end_year":   "1996",
		},
	},
	510219: {
		{
			"address":    "巴南区",
			"start_year": "1994",
			"end_year":   "1996",
		},
	},
	510221: {
		{
			"address":    "长寿县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	510222: {
		{
			"address":    "巴县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	510223: {
		{
			"address":    "綦江县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	510224: {
		{
			"address":    "江北县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	510225: {
		{
			"address":    "江津县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	510226: {
		{
			"address":    "合川县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	510227: {
		{
			"address":    "潼南县",
			"start_year": "1983",
			"end_year":   "1996",
		},
	},
	510228: {
		{
			"address":    "铜梁县",
			"start_year": "1983",
			"end_year":   "1996",
		},
	},
	510229: {
		{
			"address":    "永川县",
			"start_year": "1983",
			"end_year":   "1991",
		},
	},
	510230: {
		{
			"address":    "大足县",
			"start_year": "1983",
			"end_year":   "1996",
		},
	},
	510231: {
		{
			"address":    "荣昌县",
			"start_year": "1983",
			"end_year":   "1996",
		},
	},
	510232: {
		{
			"address":    "璧山县",
			"start_year": "1983",
			"end_year":   "1996",
		},
	},
	510281: {
		{
			"address":    "永川市",
			"start_year": "1995",
			"end_year":   "1996",
		},
	},
	510282: {
		{
			"address":    "合川市",
			"start_year": "1995",
			"end_year":   "1996",
		},
	},
	510283: {
		{
			"address":    "江津市",
			"start_year": "1995",
			"end_year":   "1996",
		},
	},
	510300: {
		{
			"address":    "自贡市",
			"start_year": "",
			"end_year":   "",
		},
	},
	510302: {
		{
			"address":    "自流井区",
			"start_year": "",
			"end_year":   "",
		},
	},
	510303: {
		{
			"address":    "贡井区",
			"start_year": "",
			"end_year":   "",
		},
	},
	510304: {
		{
			"address":    "大安区",
			"start_year": "",
			"end_year":   "",
		},
	},
	510311: {
		{
			"address":    "沿滩区",
			"start_year": "",
			"end_year":   "",
		},
	},
	510321: {
		{
			"address":    "荣县",
			"start_year": "",
			"end_year":   "",
		},
	},
	510322: {
		{
			"address":    "富顺县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	510400: {
		{
			"address":    "渡口市",
			"start_year": "",
			"end_year":   "1986",
		},
		{
			"address":    "攀枝花市",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	510402: {
		{
			"address":    "东区",
			"start_year": "",
			"end_year":   "",
		},
	},
	510403: {
		{
			"address":    "西区",
			"start_year": "",
			"end_year":   "",
		},
	},
	510411: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "仁和区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	510421: {
		{
			"address":    "米易县",
			"start_year": "",
			"end_year":   "",
		},
	},
	510422: {
		{
			"address":    "盐边县",
			"start_year": "",
			"end_year":   "",
		},
	},
	510500: {
		{
			"address":    "泸州市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	510502: {
		{
			"address":    "市中区",
			"start_year": "1984",
			"end_year":   "1994",
		},
		{
			"address":    "江阳区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	510503: {
		{
			"address":    "纳溪区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	510504: {
		{
			"address":    "龙马潭区",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	510521: {
		{
			"address":    "泸县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	510522: {
		{
			"address":    "合江县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	510523: {
		{
			"address":    "纳溪县",
			"start_year": "1983",
			"end_year":   "1994",
		},
	},
	510524: {
		{
			"address":    "叙永县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	510525: {
		{
			"address":    "古蔺县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	510600: {
		{
			"address":    "德阳市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	510602: {
		{
			"address":    "市中区",
			"start_year": "1984",
			"end_year":   "1995",
		},
	},
	510603: {
		{
			"address":    "旌阳区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	510604: {
		{
			"address":    "罗江区",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	510621: {
		{
			"address":    "德阳县",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	510622: {
		{
			"address":    "绵竹县",
			"start_year": "1983",
			"end_year":   "1995",
		},
	},
	510623: {
		{
			"address":    "中江县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	510624: {
		{
			"address":    "广汉县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	510625: {
		{
			"address":    "什邡县",
			"start_year": "1983",
			"end_year":   "1994",
		},
	},
	510626: {
		{
			"address":    "罗江县",
			"start_year": "1996",
			"end_year":   "2016",
		},
	},
	510681: {
		{
			"address":    "广汉市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	510682: {
		{
			"address":    "什邡市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	510683: {
		{
			"address":    "绵竹市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	510700: {
		{
			"address":    "绵阳市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	510702: {
		{
			"address":    "市中区",
			"start_year": "1985",
			"end_year":   "1991",
		},
	},
	510703: {
		{
			"address":    "涪城区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	510704: {
		{
			"address":    "游仙区",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	510705: {
		{
			"address":    "安州区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	510721: {
		{
			"address":    "江油县",
			"start_year": "1985",
			"end_year":   "1987",
		},
	},
	510722: {
		{
			"address":    "三台县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	510723: {
		{
			"address":    "盐亭县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	510724: {
		{
			"address":    "安县",
			"start_year": "1985",
			"end_year":   "2015",
		},
	},
	510725: {
		{
			"address":    "梓潼县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	510726: {
		{
			"address":    "北川县",
			"start_year": "1985",
			"end_year":   "2002",
		},
		{
			"address":    "北川羌族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	510727: {
		{
			"address":    "平武县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	510781: {
		{
			"address":    "江油市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	510800: {
		{
			"address":    "广元市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	510802: {
		{
			"address":    "市中区",
			"start_year": "1985",
			"end_year":   "2006",
		},
		{
			"address":    "利州区",
			"start_year": "2007",
			"end_year":   "",
		},
	},
	510811: {
		{
			"address":    "元坝区",
			"start_year": "1989",
			"end_year":   "2012",
		},
		{
			"address":    "昭化区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	510812: {
		{
			"address":    "朝天区",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	510821: {
		{
			"address":    "旺苍县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	510822: {
		{
			"address":    "青川县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	510823: {
		{
			"address":    "剑阁县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	510824: {
		{
			"address":    "苍溪县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	510900: {
		{
			"address":    "遂宁市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	510902: {
		{
			"address":    "市中区",
			"start_year": "1985",
			"end_year":   "2002",
		},
	},
	510903: {
		{
			"address":    "船山区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	510904: {
		{
			"address":    "安居区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	510921: {
		{
			"address":    "蓬溪县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	510922: {
		{
			"address":    "射洪县",
			"start_year": "1985",
			"end_year":   "2018",
		},
	},
	510923: {
		{
			"address":    "大英县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	510981: {
		{
			"address":    "射洪市",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	511000: {
		{
			"address":    "内江市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	511002: {
		{
			"address":    "市中区",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	511011: {
		{
			"address":    "东兴区",
			"start_year": "1989",
			"end_year":   "",
		},
	},
	511021: {
		{
			"address":    "内江县",
			"start_year": "1985",
			"end_year":   "1988",
		},
	},
	511022: {
		{
			"address":    "乐至县",
			"start_year": "1985",
			"end_year":   "1997",
		},
	},
	511023: {
		{
			"address":    "安岳县",
			"start_year": "1985",
			"end_year":   "1997",
		},
	},
	511024: {
		{
			"address":    "威远县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	511025: {
		{
			"address":    "资中县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	511026: {
		{
			"address":    "资阳县",
			"start_year": "1985",
			"end_year":   "1992",
		},
	},
	511027: {
		{
			"address":    "简阳县",
			"start_year": "1985",
			"end_year":   "1993",
		},
	},
	511028: {
		{
			"address":    "隆昌县",
			"start_year": "1985",
			"end_year":   "2016",
		},
	},
	511081: {
		{
			"address":    "资阳市",
			"start_year": "1995",
			"end_year":   "1997",
		},
	},
	511082: {
		{
			"address":    "简阳市",
			"start_year": "1995",
			"end_year":   "1997",
		},
	},
	511083: {
		{
			"address":    "隆昌市",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	511100: {
		{
			"address":    "乐山市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	511102: {
		{
			"address":    "市中区",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	511111: {
		{
			"address":    "沙湾区",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	511112: {
		{
			"address":    "五通桥区",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	511113: {
		{
			"address":    "金口河区",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	511121: {
		{
			"address":    "仁寿县",
			"start_year": "1985",
			"end_year":   "1996",
		},
	},
	511122: {
		{
			"address":    "眉山县",
			"start_year": "1985",
			"end_year":   "1996",
		},
	},
	511123: {
		{
			"address":    "犍为县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	511124: {
		{
			"address":    "井研县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	511125: {
		{
			"address":    "峨眉县",
			"start_year": "1985",
			"end_year":   "1987",
		},
	},
	511126: {
		{
			"address":    "夹江县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	511127: {
		{
			"address":    "洪雅县",
			"start_year": "1985",
			"end_year":   "1996",
		},
	},
	511128: {
		{
			"address":    "彭山县",
			"start_year": "1985",
			"end_year":   "1996",
		},
	},
	511129: {
		{
			"address":    "沐川县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	511130: {
		{
			"address":    "青神县",
			"start_year": "1985",
			"end_year":   "1996",
		},
	},
	511131: {
		{
			"address":    "丹棱县",
			"start_year": "1985",
			"end_year":   "1996",
		},
	},
	511132: {
		{
			"address":    "峨边彝族自治县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	511133: {
		{
			"address":    "马边彝族自治县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	511181: {
		{
			"address":    "峨眉山市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	511200: {
		{
			"address":    "万县市",
			"start_year": "1992",
			"end_year":   "1996",
		},
	},
	511202: {
		{
			"address":    "龙宝区",
			"start_year": "1992",
			"end_year":   "1996",
		},
	},
	511203: {
		{
			"address":    "天城区",
			"start_year": "1992",
			"end_year":   "1996",
		},
	},
	511204: {
		{
			"address":    "五桥区",
			"start_year": "1992",
			"end_year":   "1996",
		},
	},
	511221: {
		{
			"address":    "开县",
			"start_year": "1992",
			"end_year":   "1996",
		},
	},
	511222: {
		{
			"address":    "忠县",
			"start_year": "1992",
			"end_year":   "1996",
		},
	},
	511223: {
		{
			"address":    "梁平县",
			"start_year": "1992",
			"end_year":   "1996",
		},
	},
	511224: {
		{
			"address":    "云阳县",
			"start_year": "1992",
			"end_year":   "1996",
		},
	},
	511225: {
		{
			"address":    "奉节县",
			"start_year": "1992",
			"end_year":   "1996",
		},
	},
	511226: {
		{
			"address":    "巫山县",
			"start_year": "1992",
			"end_year":   "1996",
		},
	},
	511227: {
		{
			"address":    "巫溪县",
			"start_year": "1992",
			"end_year":   "1996",
		},
	},
	511228: {
		{
			"address":    "城口县",
			"start_year": "1992",
			"end_year":   "1996",
		},
	},
	511300: {
		{
			"address":    "南充市",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	511302: {
		{
			"address":    "顺庆区",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	511303: {
		{
			"address":    "高坪区",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	511304: {
		{
			"address":    "嘉陵区",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	511321: {
		{
			"address":    "南部县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	511322: {
		{
			"address":    "营山县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	511323: {
		{
			"address":    "蓬安县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	511324: {
		{
			"address":    "仪陇县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	511325: {
		{
			"address":    "西充县",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	511381: {
		{
			"address":    "阆中市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	511400: {
		{
			"address":    "眉山市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511402: {
		{
			"address":    "东坡区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511403: {
		{
			"address":    "彭山区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	511421: {
		{
			"address":    "仁寿县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511422: {
		{
			"address":    "彭山县",
			"start_year": "2000",
			"end_year":   "2013",
		},
	},
	511423: {
		{
			"address":    "洪雅县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511424: {
		{
			"address":    "丹棱县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511425: {
		{
			"address":    "青神县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511500: {
		{
			"address":    "宜宾市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	511502: {
		{
			"address":    "翠屏区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	511503: {
		{
			"address":    "南溪区",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	511504: {
		{
			"address":    "叙州区",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	511521: {
		{
			"address":    "宜宾县",
			"start_year": "1996",
			"end_year":   "2017",
		},
	},
	511522: {
		{
			"address":    "南溪县",
			"start_year": "1996",
			"end_year":   "2010",
		},
	},
	511523: {
		{
			"address":    "江安县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	511524: {
		{
			"address":    "长宁县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	511525: {
		{
			"address":    "高县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	511526: {
		{
			"address":    "珙县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	511527: {
		{
			"address":    "筠连县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	511528: {
		{
			"address":    "兴文县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	511529: {
		{
			"address":    "屏山县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	511600: {
		{
			"address":    "广安市",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	511602: {
		{
			"address":    "广安区",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	511603: {
		{
			"address":    "前锋区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	511621: {
		{
			"address":    "岳池县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	511622: {
		{
			"address":    "武胜县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	511623: {
		{
			"address":    "邻水县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	511681: {
		{
			"address":    "华蓥市",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	511700: {
		{
			"address":    "达州市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	511702: {
		{
			"address":    "通川区",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	511703: {
		{
			"address":    "达川区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	511721: {
		{
			"address":    "达县",
			"start_year": "1999",
			"end_year":   "2012",
		},
	},
	511722: {
		{
			"address":    "宣汉县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	511723: {
		{
			"address":    "开江县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	511724: {
		{
			"address":    "大竹县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	511725: {
		{
			"address":    "渠县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	511781: {
		{
			"address":    "万源市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	511800: {
		{
			"address":    "雅安市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511802: {
		{
			"address":    "雨城区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511803: {
		{
			"address":    "名山区",
			"start_year": "2012",
			"end_year":   "",
		},
	},
	511821: {
		{
			"address":    "名山县",
			"start_year": "2000",
			"end_year":   "2011",
		},
	},
	511822: {
		{
			"address":    "荥经县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511823: {
		{
			"address":    "汉源县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511824: {
		{
			"address":    "石棉县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511825: {
		{
			"address":    "天全县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511826: {
		{
			"address":    "芦山县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511827: {
		{
			"address":    "宝兴县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511900: {
		{
			"address":    "巴中市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511902: {
		{
			"address":    "巴州区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511903: {
		{
			"address":    "恩阳区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	511921: {
		{
			"address":    "通江县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511922: {
		{
			"address":    "南江县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	511923: {
		{
			"address":    "平昌县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	512000: {
		{
			"address":    "资阳市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	512002: {
		{
			"address":    "雁江区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	512021: {
		{
			"address":    "安岳县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	512022: {
		{
			"address":    "乐至县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	512081: {
		{
			"address":    "简阳市",
			"start_year": "2000",
			"end_year":   "2015",
		},
	},
	512100: {
		{
			"address":    "温江地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "永川地区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512121: {
		{
			"address":    "温江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "江津县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512122: {
		{
			"address":    "郫县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "合川县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512123: {
		{
			"address":    "灌县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "潼南县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512124: {
		{
			"address":    "彭县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "铜梁县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512125: {
		{
			"address":    "什邡县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "永川县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512126: {
		{
			"address":    "广汉县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "大足县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512127: {
		{
			"address":    "新都县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "荣昌县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512128: {
		{
			"address":    "新津县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "璧山县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512129: {
		{
			"address":    "蒲江县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512130: {
		{
			"address":    "邛崃县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512131: {
		{
			"address":    "大邑县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512132: {
		{
			"address":    "崇庆县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512200: {
		{
			"address":    "绵阳地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "万县地区",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	512201: {
		{
			"address":    "绵阳市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "万县市",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	512202: {
		{
			"address":    "市中区",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512221: {
		{
			"address":    "江油县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "万县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	512222: {
		{
			"address":    "青川县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "开县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	512223: {
		{
			"address":    "平武县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "忠县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	512224: {
		{
			"address":    "广元县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "梁平县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	512225: {
		{
			"address":    "旺苍县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "云阳县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	512226: {
		{
			"address":    "剑阁县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "奉节县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	512227: {
		{
			"address":    "梓潼县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "巫山县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	512228: {
		{
			"address":    "三台县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "巫溪县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	512229: {
		{
			"address":    "盐亭县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "城口县",
			"start_year": "1982",
			"end_year":   "1991",
		},
	},
	512230: {
		{
			"address":    "射洪县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512231: {
		{
			"address":    "遂宁县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512232: {
		{
			"address":    "蓬溪县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512233: {
		{
			"address":    "中江县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512234: {
		{
			"address":    "德阳县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512235: {
		{
			"address":    "绵竹县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512236: {
		{
			"address":    "安县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512237: {
		{
			"address":    "北川县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512300: {
		{
			"address":    "内江地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "涪陵地区",
			"start_year": "1982",
			"end_year":   "1994",
		},
	},
	512301: {
		{
			"address":    "内江市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "涪陵市",
			"start_year": "1983",
			"end_year":   "1994",
		},
	},
	512302: {
		{
			"address":    "南川市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	512321: {
		{
			"address":    "内江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "涪陵县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512322: {
		{
			"address":    "乐至县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "垫江县",
			"start_year": "1982",
			"end_year":   "1994",
		},
	},
	512323: {
		{
			"address":    "安岳县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "南川县",
			"start_year": "1982",
			"end_year":   "1993",
		},
	},
	512324: {
		{
			"address":    "威远县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "丰都县",
			"start_year": "1982",
			"end_year":   "1994",
		},
	},
	512325: {
		{
			"address":    "资中县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "石柱县",
			"start_year": "1982",
			"end_year":   "1982",
		},
		{
			"address":    "石柱土家族自治县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	512326: {
		{
			"address":    "资阳县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "武隆县",
			"start_year": "1982",
			"end_year":   "1994",
		},
	},
	512327: {
		{
			"address":    "简阳县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "彭水县",
			"start_year": "1982",
			"end_year":   "1982",
		},
		{
			"address":    "彭水苗族土家族自治县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	512328: {
		{
			"address":    "黔江县",
			"start_year": "1982",
			"end_year":   "1982",
		},
		{
			"address":    "黔江土家族苗族自治县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	512329: {
		{
			"address":    "酉阳县",
			"start_year": "1982",
			"end_year":   "1982",
		},
		{
			"address":    "酉阳土家族苗族自治县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	512330: {
		{
			"address":    "秀山县",
			"start_year": "1982",
			"end_year":   "1982",
		},
		{
			"address":    "秀山土家族苗族自治县",
			"start_year": "1983",
			"end_year":   "1987",
		},
	},
	512400: {
		{
			"address":    "宜宾地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "内江地区",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512401: {
		{
			"address":    "宜宾市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "内江市",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512402: {
		{
			"address":    "泸州市",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512421: {
		{
			"address":    "宜宾县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "内江县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512422: {
		{
			"address":    "富顺县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "资中县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512423: {
		{
			"address":    "隆昌县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "资阳县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512424: {
		{
			"address":    "南溪县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "简阳县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512425: {
		{
			"address":    "江安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "威远县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512426: {
		{
			"address":    "纳溪县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "隆昌县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512427: {
		{
			"address":    "泸县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "安岳县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512428: {
		{
			"address":    "合江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "乐至县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512429: {
		{
			"address":    "古蔺县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512430: {
		{
			"address":    "叙永县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512431: {
		{
			"address":    "长宁县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512432: {
		{
			"address":    "兴文县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512433: {
		{
			"address":    "珙县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512434: {
		{
			"address":    "高县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512435: {
		{
			"address":    "筠连县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512436: {
		{
			"address":    "屏山县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	512500: {
		{
			"address":    "乐山地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宜宾地区",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	512501: {
		{
			"address":    "乐山市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宜宾市",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	512502: {
		{
			"address":    "泸州市",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512521: {
		{
			"address":    "夹江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "泸县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512522: {
		{
			"address":    "洪雅县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "富顺县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512523: {
		{
			"address":    "丹棱县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "合江县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512524: {
		{
			"address":    "青神县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "纳溪县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512525: {
		{
			"address":    "眉山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "叙永县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512526: {
		{
			"address":    "彭山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "古蔺县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512527: {
		{
			"address":    "井研县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宜宾县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	512528: {
		{
			"address":    "仁寿县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "南溪县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	512529: {
		{
			"address":    "犍为县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "江安县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	512530: {
		{
			"address":    "沐川县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "长宁县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	512531: {
		{
			"address":    "峨眉县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "高县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	512532: {
		{
			"address":    "金口河工农区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "筠连县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	512533: {
		{
			"address":    "峨边县",
			"start_year": "1981",
			"end_year":   "1981",
		},
		{
			"address":    "珙县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	512534: {
		{
			"address":    "马边县",
			"start_year": "1981",
			"end_year":   "1981",
		},
		{
			"address":    "兴文县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	512535: {
		{
			"address":    "屏山县",
			"start_year": "1982",
			"end_year":   "1995",
		},
	},
	512600: {
		{
			"address":    "江津地区",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "永川地区",
			"start_year": "1981",
			"end_year":   "1981",
		},
		{
			"address":    "乐山地区",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512601: {
		{
			"address":    "乐山市",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512621: {
		{
			"address":    "永川县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "仁寿县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512622: {
		{
			"address":    "大足县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "眉山县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512623: {
		{
			"address":    "铜梁县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "犍为县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512624: {
		{
			"address":    "合川县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "井研县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512625: {
		{
			"address":    "潼南县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "峨眉县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512626: {
		{
			"address":    "璧山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "夹江县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512627: {
		{
			"address":    "江津县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "洪雅县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512628: {
		{
			"address":    "荣昌县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "彭山县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512629: {
		{
			"address":    "沐川县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512630: {
		{
			"address":    "青神县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512631: {
		{
			"address":    "丹棱县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512632: {
		{
			"address":    "峨边县",
			"start_year": "1982",
			"end_year":   "1982",
		},
		{
			"address":    "峨边彝族自治县",
			"start_year": "1983",
			"end_year":   "1984",
		},
	},
	512633: {
		{
			"address":    "马边县",
			"start_year": "1982",
			"end_year":   "1982",
		},
		{
			"address":    "马边彝族自治县",
			"start_year": "1983",
			"end_year":   "1984",
		},
	},
	512634: {
		{
			"address":    "金口河工农区",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512700: {
		{
			"address":    "涪陵地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "温江地区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512721: {
		{
			"address":    "涪陵县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "温江县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512722: {
		{
			"address":    "垫江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "郫县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512723: {
		{
			"address":    "丰都县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "新都县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512724: {
		{
			"address":    "石柱县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "广汉县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512725: {
		{
			"address":    "秀山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "什邡县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512726: {
		{
			"address":    "酉阳县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "彭县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512727: {
		{
			"address":    "黔江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "灌县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512728: {
		{
			"address":    "彭水县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "崇庆县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512729: {
		{
			"address":    "武隆县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "大邑县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512730: {
		{
			"address":    "南川县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "邛崃县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512731: {
		{
			"address":    "蒲江县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512732: {
		{
			"address":    "新津县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512800: {
		{
			"address":    "万县地区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "绵阳地区",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512801: {
		{
			"address":    "万县市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "绵阳市",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512802: {
		{
			"address":    "市中区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512821: {
		{
			"address":    "万县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "德阳县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512822: {
		{
			"address":    "开县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "绵竹县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512823: {
		{
			"address":    "城口县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "安县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512824: {
		{
			"address":    "巫溪县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "江油县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512825: {
		{
			"address":    "巫山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "梓潼县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512826: {
		{
			"address":    "奉节县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "剑阁县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512827: {
		{
			"address":    "云阳县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "广元县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512828: {
		{
			"address":    "忠县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "旺苍县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512829: {
		{
			"address":    "梁平县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "青川县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512830: {
		{
			"address":    "平武县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512831: {
		{
			"address":    "北川县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512832: {
		{
			"address":    "遂宁县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512833: {
		{
			"address":    "三台县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512834: {
		{
			"address":    "中江县",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	512835: {
		{
			"address":    "蓬溪县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512836: {
		{
			"address":    "射洪县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512837: {
		{
			"address":    "盐亭县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512900: {
		{
			"address":    "南充地区",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	512901: {
		{
			"address":    "南充市",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	512902: {
		{
			"address":    "华蓥市",
			"start_year": "1985",
			"end_year":   "1992",
		},
	},
	512903: {
		{
			"address":    "阆中市",
			"start_year": "1991",
			"end_year":   "1992",
		},
	},
	512921: {
		{
			"address":    "南充县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	512922: {
		{
			"address":    "苍溪县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "南部县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	512923: {
		{
			"address":    "阆中县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "岳池县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	512924: {
		{
			"address":    "仪陇县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "营山县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	512925: {
		{
			"address":    "南部县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "广安县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	512926: {
		{
			"address":    "西充县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "蓬安县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	512927: {
		{
			"address":    "营山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "仪陇县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	512928: {
		{
			"address":    "蓬安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "武胜县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	512929: {
		{
			"address":    "广安县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "西充县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	512930: {
		{
			"address":    "岳池县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "阆中县",
			"start_year": "1982",
			"end_year":   "1990",
		},
	},
	512931: {
		{
			"address":    "武胜县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "苍溪县",
			"start_year": "1982",
			"end_year":   "1984",
		},
	},
	512932: {
		{
			"address":    "华云工农区",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	513000: {
		{
			"address":    "达县地区",
			"start_year": "",
			"end_year":   "1992",
		},
		{
			"address":    "达川地区",
			"start_year": "1993",
			"end_year":   "1998",
		},
	},
	513001: {
		{
			"address":    "达县市",
			"start_year": "",
			"end_year":   "1992",
		},
		{
			"address":    "达川市",
			"start_year": "1993",
			"end_year":   "1998",
		},
	},
	513002: {
		{
			"address":    "万源市",
			"start_year": "1993",
			"end_year":   "1998",
		},
	},
	513021: {
		{
			"address":    "达县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	513022: {
		{
			"address":    "万源县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "宣汉县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	513023: {
		{
			"address":    "宣汉县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "开江县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	513024: {
		{
			"address":    "开江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "万源县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	513025: {
		{
			"address":    "邻水县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "通江县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	513026: {
		{
			"address":    "大竹县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "南江县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	513027: {
		{
			"address":    "渠县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "巴中县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	513028: {
		{
			"address":    "南江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "平昌县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	513029: {
		{
			"address":    "巴中县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "大竹县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	513030: {
		{
			"address":    "平昌县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "渠县",
			"start_year": "1982",
			"end_year":   "1998",
		},
	},
	513031: {
		{
			"address":    "通江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "邻水县",
			"start_year": "1982",
			"end_year":   "1992",
		},
	},
	513032: {
		{
			"address":    "白沙工农区",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	513100: {
		{
			"address":    "雅安地区",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	513101: {
		{
			"address":    "雅安市",
			"start_year": "1983",
			"end_year":   "1999",
		},
	},
	513121: {
		{
			"address":    "雅安县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	513122: {
		{
			"address":    "芦山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "名山县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	513123: {
		{
			"address":    "名山县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "荥经县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	513124: {
		{
			"address":    "荥经县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "汉源县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	513125: {
		{
			"address":    "汉源县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "石棉县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	513126: {
		{
			"address":    "石棉县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "天全县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	513127: {
		{
			"address":    "天全县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "芦山县",
			"start_year": "1982",
			"end_year":   "1999",
		},
	},
	513128: {
		{
			"address":    "宝兴县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	513200: {
		{
			"address":    "阿坝藏族自治州",
			"start_year": "",
			"end_year":   "1986",
		},
		{
			"address":    "阿坝藏族羌族自治州",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	513201: {
		{
			"address":    "马尔康区",
			"start_year": "2015",
			"end_year":   "2015",
		},
		{
			"address":    "马尔康市",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	513221: {
		{
			"address":    "马尔康县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "汶川县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513222: {
		{
			"address":    "红原县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "理县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513223: {
		{
			"address":    "阿坝县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "茂汶羌族自治县",
			"start_year": "1982",
			"end_year":   "1986",
		},
		{
			"address":    "茂县",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	513224: {
		{
			"address":    "若尔盖县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "松潘县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513225: {
		{
			"address":    "黑水县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "南坪县",
			"start_year": "1982",
			"end_year":   "1996",
		},
		{
			"address":    "九寨沟县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	513226: {
		{
			"address":    "松潘县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "金川县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513227: {
		{
			"address":    "南坪县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "小金县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513228: {
		{
			"address":    "茂汶羌族自治县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "黑水县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513229: {
		{
			"address":    "汶川县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "马尔康县",
			"start_year": "1982",
			"end_year":   "2014",
		},
	},
	513230: {
		{
			"address":    "理县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "壤塘县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513231: {
		{
			"address":    "小金县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "阿坝县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513232: {
		{
			"address":    "金川县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "若尔盖县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513233: {
		{
			"address":    "壤塘县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "红原县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513300: {
		{
			"address":    "甘孜藏族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	513301: {
		{
			"address":    "康定市",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	513321: {
		{
			"address":    "康定县",
			"start_year": "",
			"end_year":   "2014",
		},
	},
	513322: {
		{
			"address":    "炉霍县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "泸定县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513323: {
		{
			"address":    "甘孜县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "丹巴县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513324: {
		{
			"address":    "新龙县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "九龙县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513325: {
		{
			"address":    "白玉县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "雅江县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513326: {
		{
			"address":    "德格县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "道孚县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513327: {
		{
			"address":    "石渠县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "炉霍县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513328: {
		{
			"address":    "色达县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "甘孜县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513329: {
		{
			"address":    "泸定县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "新龙县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513330: {
		{
			"address":    "丹巴县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "德格县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513331: {
		{
			"address":    "九龙县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "白玉县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513332: {
		{
			"address":    "雅江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "石渠县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513333: {
		{
			"address":    "道孚县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "色达县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513334: {
		{
			"address":    "理塘县",
			"start_year": "",
			"end_year":   "",
		},
	},
	513335: {
		{
			"address":    "乡城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "巴塘县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513336: {
		{
			"address":    "稻城县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "乡城县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513337: {
		{
			"address":    "巴塘县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "稻城县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513338: {
		{
			"address":    "得荣县",
			"start_year": "",
			"end_year":   "",
		},
	},
	513400: {
		{
			"address":    "凉山彝族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	513401: {
		{
			"address":    "西昌市",
			"start_year": "",
			"end_year":   "",
		},
	},
	513421: {
		{
			"address":    "西昌县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	513422: {
		{
			"address":    "昭觉县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "木里藏族自治县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513423: {
		{
			"address":    "甘洛县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "盐源县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513424: {
		{
			"address":    "峨边县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "德昌县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513425: {
		{
			"address":    "马边县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "会理县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513426: {
		{
			"address":    "雷波县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "会东县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513427: {
		{
			"address":    "宁南县",
			"start_year": "",
			"end_year":   "",
		},
	},
	513428: {
		{
			"address":    "会东县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "普格县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513429: {
		{
			"address":    "会理县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "布拖县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513430: {
		{
			"address":    "德昌县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "金阳县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513431: {
		{
			"address":    "美姑县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "昭觉县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513432: {
		{
			"address":    "金阳县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "喜德县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513433: {
		{
			"address":    "布拖县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "冕宁县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513434: {
		{
			"address":    "普格县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "越西县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513435: {
		{
			"address":    "喜德县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "甘洛县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513436: {
		{
			"address":    "越西县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "美姑县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513437: {
		{
			"address":    "盐源县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "雷波县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	513438: {
		{
			"address":    "木里藏族自治县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	513439: {
		{
			"address":    "冕宁县",
			"start_year": "",
			"end_year":   "1981",
		},
	},
	513500: {
		{
			"address":    "黔江地区",
			"start_year": "1988",
			"end_year":   "1996",
		},
	},
	513521: {
		{
			"address":    "石柱土家族自治县",
			"start_year": "1988",
			"end_year":   "1996",
		},
	},
	513522: {
		{
			"address":    "秀山土家族苗族自治县",
			"start_year": "1988",
			"end_year":   "1996",
		},
	},
	513523: {
		{
			"address":    "黔江土家族苗族自治县",
			"start_year": "1988",
			"end_year":   "1996",
		},
	},
	513524: {
		{
			"address":    "酉阳土家族苗族自治县",
			"start_year": "1988",
			"end_year":   "1996",
		},
	},
	513525: {
		{
			"address":    "彭水苗族土家族自治县",
			"start_year": "1988",
			"end_year":   "1996",
		},
	},
	513600: {
		{
			"address":    "广安地区",
			"start_year": "1993",
			"end_year":   "1997",
		},
	},
	513601: {
		{
			"address":    "华蓥市",
			"start_year": "1993",
			"end_year":   "1997",
		},
	},
	513621: {
		{
			"address":    "岳池县",
			"start_year": "1993",
			"end_year":   "1997",
		},
	},
	513622: {
		{
			"address":    "广安县",
			"start_year": "1993",
			"end_year":   "1997",
		},
	},
	513623: {
		{
			"address":    "武胜县",
			"start_year": "1993",
			"end_year":   "1997",
		},
	},
	513624: {
		{
			"address":    "邻水县",
			"start_year": "1993",
			"end_year":   "1997",
		},
	},
	513700: {
		{
			"address":    "巴中地区",
			"start_year": "1993",
			"end_year":   "1999",
		},
	},
	513701: {
		{
			"address":    "巴中市",
			"start_year": "1993",
			"end_year":   "1999",
		},
	},
	513721: {
		{
			"address":    "通江县",
			"start_year": "1993",
			"end_year":   "1999",
		},
	},
	513722: {
		{
			"address":    "南江县",
			"start_year": "1993",
			"end_year":   "1999",
		},
	},
	513723: {
		{
			"address":    "平昌县",
			"start_year": "1993",
			"end_year":   "1999",
		},
	},
	513800: {
		{
			"address":    "眉山地区",
			"start_year": "1997",
			"end_year":   "1999",
		},
	},
	513821: {
		{
			"address":    "眉山县",
			"start_year": "1997",
			"end_year":   "1999",
		},
	},
	513822: {
		{
			"address":    "仁寿县",
			"start_year": "1997",
			"end_year":   "1999",
		},
	},
	513823: {
		{
			"address":    "彭山县",
			"start_year": "1997",
			"end_year":   "1999",
		},
	},
	513824: {
		{
			"address":    "洪雅县",
			"start_year": "1997",
			"end_year":   "1999",
		},
	},
	513825: {
		{
			"address":    "丹棱县",
			"start_year": "1997",
			"end_year":   "1999",
		},
	},
	513826: {
		{
			"address":    "青神县",
			"start_year": "1997",
			"end_year":   "1999",
		},
	},
	513900: {
		{
			"address":    "资阳地区",
			"start_year": "1998",
			"end_year":   "1999",
		},
	},
	513901: {
		{
			"address":    "资阳市",
			"start_year": "1998",
			"end_year":   "1999",
		},
	},
	513902: {
		{
			"address":    "简阳市",
			"start_year": "1998",
			"end_year":   "1999",
		},
	},
	513921: {
		{
			"address":    "安岳县",
			"start_year": "1998",
			"end_year":   "1999",
		},
	},
	513922: {
		{
			"address":    "乐至县",
			"start_year": "1998",
			"end_year":   "1999",
		},
	},
	517000: {
		{
			"address":    "涪陵市",
			"start_year": "1995",
			"end_year":   "1996",
		},
	},
	517002: {
		{
			"address":    "枳城区",
			"start_year": "1995",
			"end_year":   "1996",
		},
	},
	517003: {
		{
			"address":    "李渡区",
			"start_year": "1995",
			"end_year":   "1996",
		},
	},
	517021: {
		{
			"address":    "垫江县",
			"start_year": "1995",
			"end_year":   "1996",
		},
	},
	517022: {
		{
			"address":    "丰都县",
			"start_year": "1995",
			"end_year":   "1996",
		},
	},
	517023: {
		{
			"address":    "武隆县",
			"start_year": "1995",
			"end_year":   "1996",
		},
	},
	517081: {
		{
			"address":    "南川市",
			"start_year": "1995",
			"end_year":   "1996",
		},
	},
	519001: {
		{
			"address":    "广汉市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	519002: {
		{
			"address":    "江油市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	519003: {
		{
			"address":    "都江堰市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	519004: {
		{
			"address":    "峨眉山市",
			"start_year": "1988",
			"end_year":   "1994",
		},
	},
	519005: {
		{
			"address":    "永川市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	519006: {
		{
			"address":    "合川市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	519007: {
		{
			"address":    "江津市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	519008: {
		{
			"address":    "阆中市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	519009: {
		{
			"address":    "资阳市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	519010: {
		{
			"address":    "彭州市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	519011: {
		{
			"address":    "简阳市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	519012: {
		{
			"address":    "邛崃市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	519013: {
		{
			"address":    "崇州市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	520000: {
		{
			"address":    "贵州省",
			"start_year": "",
			"end_year":   "",
		},
	},
	520100: {
		{
			"address":    "贵阳市",
			"start_year": "",
			"end_year":   "",
		},
	},
	520102: {
		{
			"address":    "南明区",
			"start_year": "",
			"end_year":   "",
		},
	},
	520103: {
		{
			"address":    "云岩区",
			"start_year": "",
			"end_year":   "",
		},
	},
	520111: {
		{
			"address":    "花溪区",
			"start_year": "",
			"end_year":   "",
		},
	},
	520112: {
		{
			"address":    "乌当区",
			"start_year": "",
			"end_year":   "",
		},
	},
	520113: {
		{
			"address":    "白云区",
			"start_year": "",
			"end_year":   "",
		},
	},
	520114: {
		{
			"address":    "小河区",
			"start_year": "2000",
			"end_year":   "2011",
		},
	},
	520115: {
		{
			"address":    "观山湖区",
			"start_year": "2012",
			"end_year":   "",
		},
	},
	520121: {
		{
			"address":    "开阳县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	520122: {
		{
			"address":    "息烽县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	520123: {
		{
			"address":    "修文县",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	520181: {
		{
			"address":    "清镇市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	520200: {
		{
			"address":    "六盘水市",
			"start_year": "",
			"end_year":   "",
		},
	},
	520201: {
		{
			"address":    "水城特区",
			"start_year": "",
			"end_year":   "1986",
		},
		{
			"address":    "钟山区",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	520202: {
		{
			"address":    "盘县特区",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	520203: {
		{
			"address":    "六枝特区",
			"start_year": "",
			"end_year":   "",
		},
	},
	520221: {
		{
			"address":    "水城县",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	520222: {
		{
			"address":    "盘县",
			"start_year": "1999",
			"end_year":   "2016",
		},
	},
	520281: {
		{
			"address":    "盘州市",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	520300: {
		{
			"address":    "遵义市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	520302: {
		{
			"address":    "红花岗区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	520303: {
		{
			"address":    "汇川区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	520304: {
		{
			"address":    "播州区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	520321: {
		{
			"address":    "遵义县",
			"start_year": "1997",
			"end_year":   "2015",
		},
	},
	520322: {
		{
			"address":    "桐梓县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	520323: {
		{
			"address":    "绥阳县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	520324: {
		{
			"address":    "正安县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	520325: {
		{
			"address":    "道真仡佬族苗族自治县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	520326: {
		{
			"address":    "务川仡佬族苗族自治县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	520327: {
		{
			"address":    "凤冈县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	520328: {
		{
			"address":    "湄潭县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	520329: {
		{
			"address":    "余庆县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	520330: {
		{
			"address":    "习水县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	520381: {
		{
			"address":    "赤水市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	520382: {
		{
			"address":    "仁怀市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	520400: {
		{
			"address":    "安顺市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	520402: {
		{
			"address":    "西秀区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	520403: {
		{
			"address":    "平坝区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	520421: {
		{
			"address":    "平坝县",
			"start_year": "2000",
			"end_year":   "2013",
		},
	},
	520422: {
		{
			"address":    "普定县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	520423: {
		{
			"address":    "镇宁布依族苗族自治县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	520424: {
		{
			"address":    "关岭布依族苗族自治县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	520425: {
		{
			"address":    "紫云苗族布依族自治县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	520500: {
		{
			"address":    "毕节市",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520502: {
		{
			"address":    "七星关区",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520521: {
		{
			"address":    "大方县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520522: {
		{
			"address":    "黔西县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520523: {
		{
			"address":    "金沙县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520524: {
		{
			"address":    "织金县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520525: {
		{
			"address":    "纳雍县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520526: {
		{
			"address":    "威宁彝族回族苗族自治县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520527: {
		{
			"address":    "赫章县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520600: {
		{
			"address":    "铜仁市",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520602: {
		{
			"address":    "碧江区",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520603: {
		{
			"address":    "万山区",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520621: {
		{
			"address":    "江口县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520622: {
		{
			"address":    "玉屏侗族自治县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520623: {
		{
			"address":    "石阡县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520624: {
		{
			"address":    "思南县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520625: {
		{
			"address":    "印江土家族苗族自治县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520626: {
		{
			"address":    "德江县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520627: {
		{
			"address":    "沿河土家族自治县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	520628: {
		{
			"address":    "松桃苗族自治县",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	522100: {
		{
			"address":    "遵义地区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	522101: {
		{
			"address":    "遵义市",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	522102: {
		{
			"address":    "赤水市",
			"start_year": "1990",
			"end_year":   "1996",
		},
	},
	522103: {
		{
			"address":    "仁怀市",
			"start_year": "1995",
			"end_year":   "1996",
		},
	},
	522121: {
		{
			"address":    "遵义县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	522122: {
		{
			"address":    "桐梓县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	522123: {
		{
			"address":    "绥阳县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	522124: {
		{
			"address":    "正安县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	522125: {
		{
			"address":    "道真县",
			"start_year": "",
			"end_year":   "1985",
		},
		{
			"address":    "道真仡佬族苗族自治县",
			"start_year": "1986",
			"end_year":   "1996",
		},
	},
	522126: {
		{
			"address":    "务川县",
			"start_year": "",
			"end_year":   "1985",
		},
		{
			"address":    "务川仡佬族苗族自治县",
			"start_year": "1986",
			"end_year":   "1996",
		},
	},
	522127: {
		{
			"address":    "凤冈县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	522128: {
		{
			"address":    "湄潭县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	522129: {
		{
			"address":    "余庆县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	522130: {
		{
			"address":    "仁怀县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	522131: {
		{
			"address":    "赤水县",
			"start_year": "",
			"end_year":   "1989",
		},
	},
	522132: {
		{
			"address":    "习水县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	522200: {
		{
			"address":    "铜仁地区",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	522201: {
		{
			"address":    "铜仁市",
			"start_year": "1987",
			"end_year":   "2010",
		},
	},
	522221: {
		{
			"address":    "铜仁县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	522222: {
		{
			"address":    "江口县",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	522223: {
		{
			"address":    "玉屏县",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "玉屏侗族自治县",
			"start_year": "1983",
			"end_year":   "2010",
		},
	},
	522224: {
		{
			"address":    "石阡县",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	522225: {
		{
			"address":    "思南县",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	522226: {
		{
			"address":    "印江县",
			"start_year": "",
			"end_year":   "1985",
		},
		{
			"address":    "印江土家族苗族自治县",
			"start_year": "1986",
			"end_year":   "2010",
		},
	},
	522227: {
		{
			"address":    "德江县",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	522228: {
		{
			"address":    "沿河县",
			"start_year": "",
			"end_year":   "1985",
		},
		{
			"address":    "沿河土家族自治县",
			"start_year": "1986",
			"end_year":   "2010",
		},
	},
	522229: {
		{
			"address":    "松桃苗族自治县",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	522230: {
		{
			"address":    "万山特区",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	522300: {
		{
			"address":    "兴义地区",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "黔西南布依族苗族自治州",
			"start_year": "1981",
			"end_year":   "",
		},
	},
	522301: {
		{
			"address":    "兴义市",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	522302: {
		{
			"address":    "兴仁市",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	522321: {
		{
			"address":    "兴义县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	522322: {
		{
			"address":    "兴仁县",
			"start_year": "",
			"end_year":   "2017",
		},
	},
	522323: {
		{
			"address":    "普安县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522324: {
		{
			"address":    "晴隆县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522325: {
		{
			"address":    "贞丰布依族苗族自治县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "贞丰县",
			"start_year": "1981",
			"end_year":   "",
		},
	},
	522326: {
		{
			"address":    "望谟布依族苗族自治县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "望谟县",
			"start_year": "1981",
			"end_year":   "",
		},
	},
	522327: {
		{
			"address":    "册亨布依族自治县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "册亨县",
			"start_year": "1981",
			"end_year":   "",
		},
	},
	522328: {
		{
			"address":    "安龙布依族苗族自治县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "安龙县",
			"start_year": "1981",
			"end_year":   "",
		},
	},
	522400: {
		{
			"address":    "毕节地区",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	522401: {
		{
			"address":    "毕节市",
			"start_year": "1993",
			"end_year":   "2010",
		},
	},
	522421: {
		{
			"address":    "毕节县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	522422: {
		{
			"address":    "大方县",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	522423: {
		{
			"address":    "黔西县",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	522424: {
		{
			"address":    "金沙县",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	522425: {
		{
			"address":    "织金县",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	522426: {
		{
			"address":    "纳雍县",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	522427: {
		{
			"address":    "威宁彝族回族苗族自治县",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	522428: {
		{
			"address":    "赫章县",
			"start_year": "",
			"end_year":   "2010",
		},
	},
	522500: {
		{
			"address":    "安顺地区",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	522501: {
		{
			"address":    "安顺市",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	522502: {
		{
			"address":    "清镇市",
			"start_year": "1992",
			"end_year":   "1994",
		},
	},
	522521: {
		{
			"address":    "安顺县",
			"start_year": "",
			"end_year":   "1989",
		},
	},
	522522: {
		{
			"address":    "开阳县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	522523: {
		{
			"address":    "息烽县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	522524: {
		{
			"address":    "修文县",
			"start_year": "",
			"end_year":   "1994",
		},
	},
	522525: {
		{
			"address":    "清镇县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	522526: {
		{
			"address":    "平坝县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	522527: {
		{
			"address":    "普定县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	522528: {
		{
			"address":    "关岭县",
			"start_year": "",
			"end_year":   "1980",
		},
		{
			"address":    "关岭布依族苗族自治县",
			"start_year": "1981",
			"end_year":   "1999",
		},
	},
	522529: {
		{
			"address":    "镇宁布依族苗族自治县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	522530: {
		{
			"address":    "紫云苗族布依族自治县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	522600: {
		{
			"address":    "黔东南苗族侗族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	522601: {
		{
			"address":    "凯里市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	522621: {
		{
			"address":    "凯里县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	522622: {
		{
			"address":    "黄平县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522623: {
		{
			"address":    "施秉县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522624: {
		{
			"address":    "三穗县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522625: {
		{
			"address":    "镇远县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522626: {
		{
			"address":    "岑巩县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522627: {
		{
			"address":    "天柱县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522628: {
		{
			"address":    "锦屏县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522629: {
		{
			"address":    "剑河县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522630: {
		{
			"address":    "台江县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522631: {
		{
			"address":    "黎平县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522632: {
		{
			"address":    "榕江县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522633: {
		{
			"address":    "从江县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522634: {
		{
			"address":    "雷山县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522635: {
		{
			"address":    "麻江县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522636: {
		{
			"address":    "丹寨县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522700: {
		{
			"address":    "黔南布依族苗族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	522701: {
		{
			"address":    "都匀市",
			"start_year": "",
			"end_year":   "",
		},
	},
	522702: {
		{
			"address":    "福泉市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	522721: {
		{
			"address":    "都匀县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	522722: {
		{
			"address":    "荔波县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522723: {
		{
			"address":    "贵定县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522724: {
		{
			"address":    "福泉县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	522725: {
		{
			"address":    "瓮安县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522726: {
		{
			"address":    "独山县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522727: {
		{
			"address":    "平塘县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522728: {
		{
			"address":    "罗甸县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522729: {
		{
			"address":    "长顺县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522730: {
		{
			"address":    "龙里县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522731: {
		{
			"address":    "惠水县",
			"start_year": "",
			"end_year":   "",
		},
	},
	522732: {
		{
			"address":    "三都水族自治县",
			"start_year": "",
			"end_year":   "",
		},
	},
	530000: {
		{
			"address":    "云南省",
			"start_year": "",
			"end_year":   "",
		},
	},
	530100: {
		{
			"address":    "昆明市",
			"start_year": "",
			"end_year":   "",
		},
	},
	530102: {
		{
			"address":    "五华区",
			"start_year": "",
			"end_year":   "",
		},
	},
	530103: {
		{
			"address":    "盘龙区",
			"start_year": "",
			"end_year":   "",
		},
	},
	530111: {
		{
			"address":    "官渡区",
			"start_year": "",
			"end_year":   "",
		},
	},
	530112: {
		{
			"address":    "西山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	530113: {
		{
			"address":    "东川区",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	530114: {
		{
			"address":    "呈贡区",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	530115: {
		{
			"address":    "晋宁区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	530121: {
		{
			"address":    "安宁县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "呈贡县",
			"start_year": "1982",
			"end_year":   "2010",
		},
	},
	530122: {
		{
			"address":    "呈贡县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "晋宁县",
			"start_year": "1982",
			"end_year":   "2015",
		},
	},
	530123: {
		{
			"address":    "富民县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "安宁县",
			"start_year": "1982",
			"end_year":   "1994",
		},
	},
	530124: {
		{
			"address":    "晋宁县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "富民县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	530125: {
		{
			"address":    "宜良县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	530126: {
		{
			"address":    "路南彝族自治县",
			"start_year": "1983",
			"end_year":   "1997",
		},
		{
			"address":    "石林彝族自治县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	530127: {
		{
			"address":    "嵩明县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	530128: {
		{
			"address":    "禄劝县",
			"start_year": "1983",
			"end_year":   "1984",
		},
		{
			"address":    "禄劝彝族苗族自治县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	530129: {
		{
			"address":    "寻甸回族彝族自治县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	530181: {
		{
			"address":    "安宁市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	530200: {
		{
			"address":    "东川市",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	530300: {
		{
			"address":    "曲靖市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530302: {
		{
			"address":    "麒麟区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530303: {
		{
			"address":    "沾益区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	530304: {
		{
			"address":    "马龙区",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	530321: {
		{
			"address":    "马龙县",
			"start_year": "1997",
			"end_year":   "2017",
		},
	},
	530322: {
		{
			"address":    "陆良县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530323: {
		{
			"address":    "师宗县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530324: {
		{
			"address":    "罗平县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530325: {
		{
			"address":    "富源县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530326: {
		{
			"address":    "会泽县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530327: {
		{
			"address":    "寻甸回族彝族自治县",
			"start_year": "1997",
			"end_year":   "1997",
		},
	},
	530328: {
		{
			"address":    "沾益县",
			"start_year": "1997",
			"end_year":   "2015",
		},
	},
	530381: {
		{
			"address":    "宣威市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530400: {
		{
			"address":    "玉溪市",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530402: {
		{
			"address":    "红塔区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530403: {
		{
			"address":    "江川区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	530421: {
		{
			"address":    "江川县",
			"start_year": "1997",
			"end_year":   "2014",
		},
	},
	530422: {
		{
			"address":    "澄江县",
			"start_year": "1997",
			"end_year":   "2018",
		},
	},
	530423: {
		{
			"address":    "通海县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530424: {
		{
			"address":    "华宁县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530425: {
		{
			"address":    "易门县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530426: {
		{
			"address":    "峨山彝族自治县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530427: {
		{
			"address":    "新平彝族傣族自治县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530428: {
		{
			"address":    "元江哈尼族彝族傣族自治县",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	530481: {
		{
			"address":    "澄江市",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	530500: {
		{
			"address":    "保山市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	530502: {
		{
			"address":    "隆阳区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	530521: {
		{
			"address":    "施甸县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	530522: {
		{
			"address":    "腾冲县",
			"start_year": "2000",
			"end_year":   "2014",
		},
	},
	530523: {
		{
			"address":    "龙陵县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	530524: {
		{
			"address":    "昌宁县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	530581: {
		{
			"address":    "腾冲市",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	530600: {
		{
			"address":    "昭通市",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	530602: {
		{
			"address":    "昭阳区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	530621: {
		{
			"address":    "鲁甸县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	530622: {
		{
			"address":    "巧家县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	530623: {
		{
			"address":    "盐津县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	530624: {
		{
			"address":    "大关县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	530625: {
		{
			"address":    "永善县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	530626: {
		{
			"address":    "绥江县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	530627: {
		{
			"address":    "镇雄县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	530628: {
		{
			"address":    "彝良县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	530629: {
		{
			"address":    "威信县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	530630: {
		{
			"address":    "水富县",
			"start_year": "2001",
			"end_year":   "2017",
		},
	},
	530681: {
		{
			"address":    "水富市",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	530700: {
		{
			"address":    "丽江市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	530702: {
		{
			"address":    "古城区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	530721: {
		{
			"address":    "玉龙纳西族自治县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	530722: {
		{
			"address":    "永胜县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	530723: {
		{
			"address":    "华坪县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	530724: {
		{
			"address":    "宁蒗彝族自治县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	530800: {
		{
			"address":    "思茅市",
			"start_year": "2003",
			"end_year":   "2006",
		},
		{
			"address":    "普洱市",
			"start_year": "2007",
			"end_year":   "",
		},
	},
	530802: {
		{
			"address":    "翠云区",
			"start_year": "2003",
			"end_year":   "2006",
		},
		{
			"address":    "思茅区",
			"start_year": "2007",
			"end_year":   "",
		},
	},
	530821: {
		{
			"address":    "普洱哈尼族彝族自治县",
			"start_year": "2003",
			"end_year":   "2006",
		},
		{
			"address":    "宁洱哈尼族彝族自治县",
			"start_year": "2007",
			"end_year":   "",
		},
	},
	530822: {
		{
			"address":    "墨江哈尼族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530823: {
		{
			"address":    "景东彝族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530824: {
		{
			"address":    "景谷傣族彝族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530825: {
		{
			"address":    "镇沅彝族哈尼族拉祜族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530826: {
		{
			"address":    "江城哈尼族彝族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530827: {
		{
			"address":    "孟连傣族拉祜族佤族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530828: {
		{
			"address":    "澜沧拉祜族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530829: {
		{
			"address":    "西盟佤族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530900: {
		{
			"address":    "临沧市",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530902: {
		{
			"address":    "临翔区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530921: {
		{
			"address":    "凤庆县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530922: {
		{
			"address":    "云县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530923: {
		{
			"address":    "永德县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530924: {
		{
			"address":    "镇康县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530925: {
		{
			"address":    "双江拉祜族佤族布朗族傣族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530926: {
		{
			"address":    "耿马傣族佤族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	530927: {
		{
			"address":    "沧源佤族自治县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	532100: {
		{
			"address":    "昭通地区",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	532101: {
		{
			"address":    "昭通市",
			"start_year": "1981",
			"end_year":   "2000",
		},
	},
	532121: {
		{
			"address":    "昭通县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	532122: {
		{
			"address":    "鲁甸县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	532123: {
		{
			"address":    "巧家县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	532124: {
		{
			"address":    "盐津县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	532125: {
		{
			"address":    "大关县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	532126: {
		{
			"address":    "永善县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	532127: {
		{
			"address":    "绥江县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	532128: {
		{
			"address":    "镇雄县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	532129: {
		{
			"address":    "彝良县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	532130: {
		{
			"address":    "威信县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	532131: {
		{
			"address":    "水富县",
			"start_year": "1981",
			"end_year":   "2000",
		},
	},
	532200: {
		{
			"address":    "曲靖地区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532201: {
		{
			"address":    "曲靖市",
			"start_year": "1983",
			"end_year":   "1996",
		},
	},
	532202: {
		{
			"address":    "宣威市",
			"start_year": "1994",
			"end_year":   "1996",
		},
	},
	532221: {
		{
			"address":    "曲靖县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	532222: {
		{
			"address":    "沾益县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	532223: {
		{
			"address":    "马龙县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532224: {
		{
			"address":    "宣威县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	532225: {
		{
			"address":    "富源县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532226: {
		{
			"address":    "罗平县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532227: {
		{
			"address":    "师宗县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532228: {
		{
			"address":    "陆良县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532229: {
		{
			"address":    "宜良县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	532230: {
		{
			"address":    "路南彝族自治县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	532231: {
		{
			"address":    "寻甸回族彝族自治县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532232: {
		{
			"address":    "嵩明县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	532233: {
		{
			"address":    "会泽县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532300: {
		{
			"address":    "楚雄彝族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	532301: {
		{
			"address":    "楚雄市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	532321: {
		{
			"address":    "楚雄县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	532322: {
		{
			"address":    "双柏县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532323: {
		{
			"address":    "牟定县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532324: {
		{
			"address":    "南华县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532325: {
		{
			"address":    "姚安县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532326: {
		{
			"address":    "大姚县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532327: {
		{
			"address":    "永仁县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532328: {
		{
			"address":    "元谋县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532329: {
		{
			"address":    "武定县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532330: {
		{
			"address":    "禄劝县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	532331: {
		{
			"address":    "禄丰县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532400: {
		{
			"address":    "玉溪地区",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532401: {
		{
			"address":    "玉溪市",
			"start_year": "1983",
			"end_year":   "1996",
		},
	},
	532421: {
		{
			"address":    "玉溪县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	532422: {
		{
			"address":    "江川县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532423: {
		{
			"address":    "澄江县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532424: {
		{
			"address":    "通海县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532425: {
		{
			"address":    "华宁县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532426: {
		{
			"address":    "易门县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532427: {
		{
			"address":    "峨山彝族自治县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532428: {
		{
			"address":    "新平彝族傣族自治县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532429: {
		{
			"address":    "元江哈尼族彝族傣族自治县",
			"start_year": "",
			"end_year":   "1996",
		},
	},
	532500: {
		{
			"address":    "红河哈尼族彝族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	532501: {
		{
			"address":    "个旧市",
			"start_year": "",
			"end_year":   "",
		},
	},
	532502: {
		{
			"address":    "开远市",
			"start_year": "1981",
			"end_year":   "",
		},
	},
	532503: {
		{
			"address":    "蒙自市",
			"start_year": "2010",
			"end_year":   "",
		},
	},
	532504: {
		{
			"address":    "弥勒市",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	532521: {
		{
			"address":    "开远县",
			"start_year": "",
			"end_year":   "1980",
		},
	},
	532522: {
		{
			"address":    "蒙自县",
			"start_year": "",
			"end_year":   "2009",
		},
	},
	532523: {
		{
			"address":    "屏边苗族自治县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532524: {
		{
			"address":    "建水县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532525: {
		{
			"address":    "石屏县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532526: {
		{
			"address":    "弥勒县",
			"start_year": "",
			"end_year":   "2012",
		},
	},
	532527: {
		{
			"address":    "泸西县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532528: {
		{
			"address":    "元阳县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532529: {
		{
			"address":    "红河县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532530: {
		{
			"address":    "金平县",
			"start_year": "",
			"end_year":   "1984",
		},
		{
			"address":    "金平苗族瑶族傣族自治县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	532531: {
		{
			"address":    "绿春县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532532: {
		{
			"address":    "河口瑶族自治县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532600: {
		{
			"address":    "文山壮族苗族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	532601: {
		{
			"address":    "文山市",
			"start_year": "2010",
			"end_year":   "",
		},
	},
	532621: {
		{
			"address":    "文山县",
			"start_year": "",
			"end_year":   "2009",
		},
	},
	532622: {
		{
			"address":    "砚山县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532623: {
		{
			"address":    "西畴县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532624: {
		{
			"address":    "麻栗坡县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532625: {
		{
			"address":    "马关县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532626: {
		{
			"address":    "丘北县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532627: {
		{
			"address":    "广南县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532628: {
		{
			"address":    "富宁县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532700: {
		{
			"address":    "思茅地区",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	532701: {
		{
			"address":    "思茅市",
			"start_year": "1993",
			"end_year":   "2002",
		},
	},
	532721: {
		{
			"address":    "思茅县",
			"start_year": "1981",
			"end_year":   "1992",
		},
	},
	532722: {
		{
			"address":    "普洱县",
			"start_year": "",
			"end_year":   "1984",
		},
		{
			"address":    "普洱哈尼族彝族自治县",
			"start_year": "1985",
			"end_year":   "2002",
		},
	},
	532723: {
		{
			"address":    "墨江哈尼族自治县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	532724: {
		{
			"address":    "景东县",
			"start_year": "",
			"end_year":   "1984",
		},
		{
			"address":    "景东彝族自治县",
			"start_year": "1985",
			"end_year":   "2002",
		},
	},
	532725: {
		{
			"address":    "景谷县",
			"start_year": "",
			"end_year":   "1984",
		},
		{
			"address":    "景谷傣族彝族自治县",
			"start_year": "1985",
			"end_year":   "2002",
		},
	},
	532726: {
		{
			"address":    "镇沅县",
			"start_year": "",
			"end_year":   "1989",
		},
		{
			"address":    "镇沅彝族哈尼族拉祜族自治县",
			"start_year": "1990",
			"end_year":   "2002",
		},
	},
	532727: {
		{
			"address":    "江城哈尼族彝族自治县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	532728: {
		{
			"address":    "孟连傣族拉祜族佤族自治县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	532729: {
		{
			"address":    "澜沧拉祜族自治县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	532730: {
		{
			"address":    "西盟佤族自治县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	532800: {
		{
			"address":    "西双版纳傣族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	532801: {
		{
			"address":    "景洪市",
			"start_year": "1993",
			"end_year":   "",
		},
	},
	532821: {
		{
			"address":    "景洪县",
			"start_year": "",
			"end_year":   "1992",
		},
	},
	532822: {
		{
			"address":    "勐海县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532823: {
		{
			"address":    "勐腊县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532900: {
		{
			"address":    "大理白族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	532901: {
		{
			"address":    "下关市",
			"start_year": "",
			"end_year":   "1982",
		},
		{
			"address":    "大理市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	532921: {
		{
			"address":    "大理县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	532922: {
		{
			"address":    "漾濞县",
			"start_year": "",
			"end_year":   "1984",
		},
		{
			"address":    "漾濞彝族自治县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	532923: {
		{
			"address":    "祥云县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532924: {
		{
			"address":    "宾川县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532925: {
		{
			"address":    "弥渡县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532926: {
		{
			"address":    "南涧彝族自治县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532927: {
		{
			"address":    "巍山彝族回族自治县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532928: {
		{
			"address":    "永平县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532929: {
		{
			"address":    "云龙县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532930: {
		{
			"address":    "洱源县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532931: {
		{
			"address":    "剑川县",
			"start_year": "",
			"end_year":   "",
		},
	},
	532932: {
		{
			"address":    "鹤庆县",
			"start_year": "",
			"end_year":   "",
		},
	},
	533000: {
		{
			"address":    "保山地区",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	533001: {
		{
			"address":    "保山市",
			"start_year": "1983",
			"end_year":   "1999",
		},
	},
	533021: {
		{
			"address":    "保山县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	533022: {
		{
			"address":    "施甸县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	533023: {
		{
			"address":    "腾冲县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	533024: {
		{
			"address":    "龙陵县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	533025: {
		{
			"address":    "昌宁县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	533100: {
		{
			"address":    "德宏傣族景颇族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	533101: {
		{
			"address":    "畹町市",
			"start_year": "1985",
			"end_year":   "1998",
		},
	},
	533102: {
		{
			"address":    "瑞丽市",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	533103: {
		{
			"address":    "潞西市",
			"start_year": "1996",
			"end_year":   "2009",
		},
		{
			"address":    "芒市",
			"start_year": "2010",
			"end_year":   "",
		},
	},
	533121: {
		{
			"address":    "潞西县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	533122: {
		{
			"address":    "梁河县",
			"start_year": "",
			"end_year":   "",
		},
	},
	533123: {
		{
			"address":    "盈江县",
			"start_year": "",
			"end_year":   "",
		},
	},
	533124: {
		{
			"address":    "陇川县",
			"start_year": "",
			"end_year":   "",
		},
	},
	533125: {
		{
			"address":    "瑞丽县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	533126: {
		{
			"address":    "畹町镇",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	533200: {
		{
			"address":    "丽江地区",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	533221: {
		{
			"address":    "丽江纳西族自治县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	533222: {
		{
			"address":    "永胜县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	533223: {
		{
			"address":    "华坪县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	533224: {
		{
			"address":    "宁蒗彝族自治县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	533300: {
		{
			"address":    "怒江傈僳族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	533301: {
		{
			"address":    "泸水市",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	533321: {
		{
			"address":    "碧江县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "泸水县",
			"start_year": "1982",
			"end_year":   "2015",
		},
	},
	533322: {
		{
			"address":    "福贡县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "碧江县",
			"start_year": "1982",
			"end_year":   "1985",
		},
	},
	533323: {
		{
			"address":    "贡山独龙族怒族自治县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "福贡县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	533324: {
		{
			"address":    "泸水县",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "贡山独龙族怒族自治县",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	533325: {
		{
			"address":    "兰坪县",
			"start_year": "",
			"end_year":   "1986",
		},
		{
			"address":    "兰坪白族普米族自治县",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	533400: {
		{
			"address":    "迪庆藏族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	533401: {
		{
			"address":    "香格里拉市",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	533421: {
		{
			"address":    "中甸县",
			"start_year": "",
			"end_year":   "2000",
		},
		{
			"address":    "香格里拉县",
			"start_year": "2001",
			"end_year":   "2013",
		},
	},
	533422: {
		{
			"address":    "德钦县",
			"start_year": "",
			"end_year":   "",
		},
	},
	533423: {
		{
			"address":    "维西县",
			"start_year": "",
			"end_year":   "1984",
		},
		{
			"address":    "维西傈僳族自治县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	533500: {
		{
			"address":    "临沧地区",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	533521: {
		{
			"address":    "临沧县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	533522: {
		{
			"address":    "凤庆县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	533523: {
		{
			"address":    "云县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	533524: {
		{
			"address":    "永德县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	533525: {
		{
			"address":    "镇康县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	533526: {
		{
			"address":    "双江县",
			"start_year": "",
			"end_year":   "1984",
		},
		{
			"address":    "双江拉祜族佤族布朗族傣族自治县",
			"start_year": "1985",
			"end_year":   "2002",
		},
	},
	533527: {
		{
			"address":    "耿马傣族佤族自治县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	533528: {
		{
			"address":    "沧源佤族自治县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	540000: {
		{
			"address":    "西藏自治区",
			"start_year": "",
			"end_year":   "",
		},
	},
	540100: {
		{
			"address":    "拉萨市",
			"start_year": "",
			"end_year":   "",
		},
	},
	540102: {
		{
			"address":    "城关区",
			"start_year": "",
			"end_year":   "",
		},
	},
	540103: {
		{
			"address":    "堆龙德庆区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	540104: {
		{
			"address":    "达孜区",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	540121: {
		{
			"address":    "林周县",
			"start_year": "",
			"end_year":   "",
		},
	},
	540122: {
		{
			"address":    "当雄县",
			"start_year": "",
			"end_year":   "",
		},
	},
	540123: {
		{
			"address":    "尼木县",
			"start_year": "",
			"end_year":   "",
		},
	},
	540124: {
		{
			"address":    "曲水县",
			"start_year": "",
			"end_year":   "",
		},
	},
	540125: {
		{
			"address":    "堆龙德庆县",
			"start_year": "",
			"end_year":   "2014",
		},
	},
	540126: {
		{
			"address":    "达孜县",
			"start_year": "",
			"end_year":   "2016",
		},
	},
	540127: {
		{
			"address":    "墨竹工卡县",
			"start_year": "",
			"end_year":   "",
		},
	},
	540128: {
		{
			"address":    "工布江达县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	540129: {
		{
			"address":    "林芝县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	540130: {
		{
			"address":    "米林县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	540131: {
		{
			"address":    "墨脱县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	540200: {
		{
			"address":    "日喀则市",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540202: {
		{
			"address":    "桑珠孜区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540221: {
		{
			"address":    "南木林县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540222: {
		{
			"address":    "江孜县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540223: {
		{
			"address":    "定日县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540224: {
		{
			"address":    "萨迦县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540225: {
		{
			"address":    "拉孜县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540226: {
		{
			"address":    "昂仁县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540227: {
		{
			"address":    "谢通门县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540228: {
		{
			"address":    "白朗县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540229: {
		{
			"address":    "仁布县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540230: {
		{
			"address":    "康马县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540231: {
		{
			"address":    "定结县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540232: {
		{
			"address":    "仲巴县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540233: {
		{
			"address":    "亚东县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540234: {
		{
			"address":    "吉隆县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540235: {
		{
			"address":    "聂拉木县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540236: {
		{
			"address":    "萨嘎县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540237: {
		{
			"address":    "岗巴县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540300: {
		{
			"address":    "昌都市",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540302: {
		{
			"address":    "卡若区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540321: {
		{
			"address":    "江达县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540322: {
		{
			"address":    "贡觉县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540323: {
		{
			"address":    "类乌齐县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540324: {
		{
			"address":    "丁青县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540325: {
		{
			"address":    "察雅县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540326: {
		{
			"address":    "八宿县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540327: {
		{
			"address":    "左贡县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540328: {
		{
			"address":    "芒康县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540329: {
		{
			"address":    "洛隆县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540330: {
		{
			"address":    "边坝县",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	540400: {
		{
			"address":    "林芝市",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	540402: {
		{
			"address":    "巴宜区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	540421: {
		{
			"address":    "工布江达县",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	540422: {
		{
			"address":    "米林县",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	540423: {
		{
			"address":    "墨脱县",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	540424: {
		{
			"address":    "波密县",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	540425: {
		{
			"address":    "察隅县",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	540426: {
		{
			"address":    "朗县",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	540500: {
		{
			"address":    "山南市",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	540502: {
		{
			"address":    "乃东区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	540521: {
		{
			"address":    "扎囊县",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	540522: {
		{
			"address":    "贡嘎县",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	540523: {
		{
			"address":    "桑日县",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	540524: {
		{
			"address":    "琼结县",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	540525: {
		{
			"address":    "曲松县",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	540526: {
		{
			"address":    "措美县",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	540527: {
		{
			"address":    "洛扎县",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	540528: {
		{
			"address":    "加查县",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	540529: {
		{
			"address":    "隆子县",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	540530: {
		{
			"address":    "错那县",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	540531: {
		{
			"address":    "浪卡子县",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	540600: {
		{
			"address":    "那曲市",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	540602: {
		{
			"address":    "色尼区",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	540621: {
		{
			"address":    "嘉黎县",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	540622: {
		{
			"address":    "比如县",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	540623: {
		{
			"address":    "聂荣县",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	540624: {
		{
			"address":    "安多县",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	540625: {
		{
			"address":    "申扎县",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	540626: {
		{
			"address":    "索县",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	540627: {
		{
			"address":    "班戈县",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	540628: {
		{
			"address":    "巴青县",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	540629: {
		{
			"address":    "尼玛县",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	540630: {
		{
			"address":    "双湖县",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	542100: {
		{
			"address":    "昌都地区",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542121: {
		{
			"address":    "昌都县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542122: {
		{
			"address":    "江达县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542123: {
		{
			"address":    "贡觉县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542124: {
		{
			"address":    "类乌齐县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542125: {
		{
			"address":    "丁青县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542126: {
		{
			"address":    "察雅县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542127: {
		{
			"address":    "八宿县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542128: {
		{
			"address":    "左贡县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542129: {
		{
			"address":    "芒康县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542130: {
		{
			"address":    "波密县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	542131: {
		{
			"address":    "察隅县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	542132: {
		{
			"address":    "洛隆县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542133: {
		{
			"address":    "边坝县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542134: {
		{
			"address":    "盐井县",
			"start_year": "1983",
			"end_year":   "1998",
		},
	},
	542135: {
		{
			"address":    "碧土县",
			"start_year": "1983",
			"end_year":   "1998",
		},
	},
	542136: {
		{
			"address":    "妥坝县",
			"start_year": "1983",
			"end_year":   "1998",
		},
	},
	542137: {
		{
			"address":    "生达县",
			"start_year": "1983",
			"end_year":   "1998",
		},
	},
	542200: {
		{
			"address":    "山南地区",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	542221: {
		{
			"address":    "乃东县",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	542222: {
		{
			"address":    "扎囊县",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	542223: {
		{
			"address":    "贡嘎县",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	542224: {
		{
			"address":    "桑日县",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	542225: {
		{
			"address":    "穷结县",
			"start_year": "",
			"end_year":   "1985",
		},
		{
			"address":    "琼结县",
			"start_year": "1986",
			"end_year":   "2015",
		},
	},
	542226: {
		{
			"address":    "曲松县",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	542227: {
		{
			"address":    "措美县",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	542228: {
		{
			"address":    "洛扎县",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	542229: {
		{
			"address":    "加查县",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	542230: {
		{
			"address":    "朗县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	542231: {
		{
			"address":    "隆子县",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	542232: {
		{
			"address":    "错那县",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	542233: {
		{
			"address":    "浪卡子县",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	542300: {
		{
			"address":    "日喀则地区",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542301: {
		{
			"address":    "日喀则市",
			"start_year": "1986",
			"end_year":   "2013",
		},
	},
	542321: {
		{
			"address":    "日喀则县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	542322: {
		{
			"address":    "南木林县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542323: {
		{
			"address":    "江孜县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542324: {
		{
			"address":    "定日县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542325: {
		{
			"address":    "萨迦县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542326: {
		{
			"address":    "拉孜县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542327: {
		{
			"address":    "昂仁县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542328: {
		{
			"address":    "谢通门县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542329: {
		{
			"address":    "白朗县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542330: {
		{
			"address":    "仁布县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542331: {
		{
			"address":    "康马县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542332: {
		{
			"address":    "定结县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542333: {
		{
			"address":    "仲巴县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542334: {
		{
			"address":    "亚东县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542335: {
		{
			"address":    "吉隆县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542336: {
		{
			"address":    "聂拉木县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542337: {
		{
			"address":    "萨嘎县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542338: {
		{
			"address":    "岗巴县",
			"start_year": "",
			"end_year":   "2013",
		},
	},
	542400: {
		{
			"address":    "那曲地区",
			"start_year": "",
			"end_year":   "2016",
		},
	},
	542421: {
		{
			"address":    "那曲县",
			"start_year": "",
			"end_year":   "2016",
		},
	},
	542422: {
		{
			"address":    "嘉黎县",
			"start_year": "",
			"end_year":   "2016",
		},
	},
	542423: {
		{
			"address":    "比如县",
			"start_year": "",
			"end_year":   "2016",
		},
	},
	542424: {
		{
			"address":    "聂荣县",
			"start_year": "",
			"end_year":   "2016",
		},
	},
	542425: {
		{
			"address":    "安多县",
			"start_year": "",
			"end_year":   "2016",
		},
	},
	542426: {
		{
			"address":    "申扎县",
			"start_year": "",
			"end_year":   "2016",
		},
	},
	542427: {
		{
			"address":    "索县",
			"start_year": "",
			"end_year":   "2016",
		},
	},
	542428: {
		{
			"address":    "班戈县",
			"start_year": "",
			"end_year":   "2016",
		},
	},
	542429: {
		{
			"address":    "巴青县",
			"start_year": "",
			"end_year":   "2016",
		},
	},
	542430: {
		{
			"address":    "尼玛县",
			"start_year": "1983",
			"end_year":   "2016",
		},
	},
	542431: {
		{
			"address":    "双湖县",
			"start_year": "2012",
			"end_year":   "2016",
		},
	},
	542500: {
		{
			"address":    "阿里地区",
			"start_year": "",
			"end_year":   "",
		},
	},
	542521: {
		{
			"address":    "普兰县",
			"start_year": "",
			"end_year":   "",
		},
	},
	542522: {
		{
			"address":    "札达县",
			"start_year": "",
			"end_year":   "",
		},
	},
	542523: {
		{
			"address":    "噶尔县",
			"start_year": "",
			"end_year":   "",
		},
	},
	542524: {
		{
			"address":    "日土县",
			"start_year": "",
			"end_year":   "",
		},
	},
	542525: {
		{
			"address":    "革吉县",
			"start_year": "",
			"end_year":   "",
		},
	},
	542526: {
		{
			"address":    "改则县",
			"start_year": "",
			"end_year":   "",
		},
	},
	542527: {
		{
			"address":    "措勤县",
			"start_year": "",
			"end_year":   "",
		},
	},
	542528: {
		{
			"address":    "隆格尔县",
			"start_year": "1983",
			"end_year":   "1998",
		},
	},
	542600: {
		{
			"address":    "林芝地区",
			"start_year": "1983",
			"end_year":   "2014",
		},
	},
	542621: {
		{
			"address":    "林芝县",
			"start_year": "1983",
			"end_year":   "2014",
		},
	},
	542622: {
		{
			"address":    "工布江达县",
			"start_year": "1983",
			"end_year":   "2014",
		},
	},
	542623: {
		{
			"address":    "米林县",
			"start_year": "1983",
			"end_year":   "2014",
		},
	},
	542624: {
		{
			"address":    "墨脱县",
			"start_year": "1983",
			"end_year":   "2014",
		},
	},
	542625: {
		{
			"address":    "波密县",
			"start_year": "1983",
			"end_year":   "2014",
		},
	},
	542626: {
		{
			"address":    "察隅县",
			"start_year": "1983",
			"end_year":   "2014",
		},
	},
	542627: {
		{
			"address":    "朗县",
			"start_year": "1983",
			"end_year":   "2014",
		},
	},
	542700: {
		{
			"address":    "江孜地区",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	542721: {
		{
			"address":    "江孜县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	542722: {
		{
			"address":    "浪卡子县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	542723: {
		{
			"address":    "白朗县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	542724: {
		{
			"address":    "仁布县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	542725: {
		{
			"address":    "康马县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	542726: {
		{
			"address":    "亚东县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	542727: {
		{
			"address":    "岗巴县",
			"start_year": "1983",
			"end_year":   "1985",
		},
	},
	610000: {
		{
			"address":    "陕西省",
			"start_year": "",
			"end_year":   "",
		},
	},
	610100: {
		{
			"address":    "西安市",
			"start_year": "",
			"end_year":   "",
		},
	},
	610102: {
		{
			"address":    "新城区",
			"start_year": "",
			"end_year":   "",
		},
	},
	610103: {
		{
			"address":    "碑林区",
			"start_year": "",
			"end_year":   "",
		},
	},
	610104: {
		{
			"address":    "莲湖区",
			"start_year": "",
			"end_year":   "",
		},
	},
	610111: {
		{
			"address":    "灞桥区",
			"start_year": "",
			"end_year":   "",
		},
	},
	610112: {
		{
			"address":    "未央区",
			"start_year": "",
			"end_year":   "",
		},
	},
	610113: {
		{
			"address":    "雁塔区",
			"start_year": "",
			"end_year":   "",
		},
	},
	610114: {
		{
			"address":    "阎良区",
			"start_year": "",
			"end_year":   "",
		},
	},
	610115: {
		{
			"address":    "临潼区",
			"start_year": "1997",
			"end_year":   "",
		},
	},
	610116: {
		{
			"address":    "长安区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	610117: {
		{
			"address":    "高陵区",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	610118: {
		{
			"address":    "鄠邑区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	610121: {
		{
			"address":    "长安县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	610122: {
		{
			"address":    "蓝田县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	610123: {
		{
			"address":    "临潼县",
			"start_year": "1983",
			"end_year":   "1996",
		},
	},
	610124: {
		{
			"address":    "周至县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	610125: {
		{
			"address":    "户县",
			"start_year": "1983",
			"end_year":   "2015",
		},
	},
	610126: {
		{
			"address":    "高陵县",
			"start_year": "1983",
			"end_year":   "2013",
		},
	},
	610200: {
		{
			"address":    "铜川市",
			"start_year": "",
			"end_year":   "",
		},
	},
	610202: {
		{
			"address":    "城区",
			"start_year": "",
			"end_year":   "1999",
		},
		{
			"address":    "王益区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	610203: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1999",
		},
		{
			"address":    "印台区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	610204: {
		{
			"address":    "耀州区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	610221: {
		{
			"address":    "耀县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	610222: {
		{
			"address":    "宜君县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	610300: {
		{
			"address":    "宝鸡市",
			"start_year": "",
			"end_year":   "",
		},
	},
	610302: {
		{
			"address":    "渭滨区",
			"start_year": "",
			"end_year":   "",
		},
	},
	610303: {
		{
			"address":    "金台区",
			"start_year": "",
			"end_year":   "",
		},
	},
	610304: {
		{
			"address":    "杨陵区",
			"start_year": "1982",
			"end_year":   "1982",
		},
		{
			"address":    "陈仓区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	610321: {
		{
			"address":    "宝鸡县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	610322: {
		{
			"address":    "凤翔县",
			"start_year": "",
			"end_year":   "",
		},
	},
	610323: {
		{
			"address":    "岐山县",
			"start_year": "",
			"end_year":   "",
		},
	},
	610324: {
		{
			"address":    "扶风县",
			"start_year": "",
			"end_year":   "",
		},
	},
	610325: {
		{
			"address":    "武功县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	610326: {
		{
			"address":    "眉县",
			"start_year": "",
			"end_year":   "",
		},
	},
	610327: {
		{
			"address":    "陇县",
			"start_year": "",
			"end_year":   "",
		},
	},
	610328: {
		{
			"address":    "千阳县",
			"start_year": "",
			"end_year":   "",
		},
	},
	610329: {
		{
			"address":    "麟游县",
			"start_year": "",
			"end_year":   "",
		},
	},
	610330: {
		{
			"address":    "凤县",
			"start_year": "",
			"end_year":   "",
		},
	},
	610331: {
		{
			"address":    "太白县",
			"start_year": "",
			"end_year":   "",
		},
	},
	610400: {
		{
			"address":    "咸阳市",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	610401: {
		{
			"address":    "咸阳市",
			"start_year": "1983",
			"end_year":   "1983",
		},
	},
	610402: {
		{
			"address":    "秦都区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	610403: {
		{
			"address":    "杨陵区",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	610404: {
		{
			"address":    "渭城区",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	610421: {
		{
			"address":    "兴平县",
			"start_year": "1983",
			"end_year":   "1992",
		},
	},
	610422: {
		{
			"address":    "三原县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	610423: {
		{
			"address":    "泾阳县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	610424: {
		{
			"address":    "乾县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	610425: {
		{
			"address":    "礼泉县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	610426: {
		{
			"address":    "永寿县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	610427: {
		{
			"address":    "彬县",
			"start_year": "1983",
			"end_year":   "2017",
		},
	},
	610428: {
		{
			"address":    "长武县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	610429: {
		{
			"address":    "旬邑县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	610430: {
		{
			"address":    "淳化县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	610431: {
		{
			"address":    "武功县",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	610481: {
		{
			"address":    "兴平市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	610482: {
		{
			"address":    "彬州市",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	610500: {
		{
			"address":    "渭南市",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	610502: {
		{
			"address":    "临渭区",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	610503: {
		{
			"address":    "华州区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	610521: {
		{
			"address":    "华县",
			"start_year": "1994",
			"end_year":   "2014",
		},
	},
	610522: {
		{
			"address":    "潼关县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	610523: {
		{
			"address":    "大荔县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	610524: {
		{
			"address":    "合阳县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	610525: {
		{
			"address":    "澄城县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	610526: {
		{
			"address":    "蒲城县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	610527: {
		{
			"address":    "白水县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	610528: {
		{
			"address":    "富平县",
			"start_year": "1994",
			"end_year":   "",
		},
	},
	610581: {
		{
			"address":    "韩城市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	610582: {
		{
			"address":    "华阴市",
			"start_year": "1995",
			"end_year":   "",
		},
	},
	610600: {
		{
			"address":    "延安市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610602: {
		{
			"address":    "宝塔区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610603: {
		{
			"address":    "安塞区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	610621: {
		{
			"address":    "延长县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610622: {
		{
			"address":    "延川县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610623: {
		{
			"address":    "子长县",
			"start_year": "1996",
			"end_year":   "2018",
		},
	},
	610624: {
		{
			"address":    "安塞县",
			"start_year": "1996",
			"end_year":   "2015",
		},
	},
	610625: {
		{
			"address":    "志丹县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610626: {
		{
			"address":    "吴旗县",
			"start_year": "1996",
			"end_year":   "2004",
		},
		{
			"address":    "吴起县",
			"start_year": "2005",
			"end_year":   "",
		},
	},
	610627: {
		{
			"address":    "甘泉县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610628: {
		{
			"address":    "富县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610629: {
		{
			"address":    "洛川县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610630: {
		{
			"address":    "宜川县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610631: {
		{
			"address":    "黄龙县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610632: {
		{
			"address":    "黄陵县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610681: {
		{
			"address":    "子长市",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	610700: {
		{
			"address":    "汉中市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610702: {
		{
			"address":    "汉台区",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610703: {
		{
			"address":    "南郑区",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	610721: {
		{
			"address":    "南郑县",
			"start_year": "1996",
			"end_year":   "2016",
		},
	},
	610722: {
		{
			"address":    "城固县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610723: {
		{
			"address":    "洋县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610724: {
		{
			"address":    "西乡县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610725: {
		{
			"address":    "勉县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610726: {
		{
			"address":    "宁强县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610727: {
		{
			"address":    "略阳县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610728: {
		{
			"address":    "镇巴县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610729: {
		{
			"address":    "留坝县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610730: {
		{
			"address":    "佛坪县",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	610800: {
		{
			"address":    "榆林市",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	610802: {
		{
			"address":    "榆阳区",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	610803: {
		{
			"address":    "横山区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	610821: {
		{
			"address":    "神木县",
			"start_year": "1999",
			"end_year":   "2016",
		},
	},
	610822: {
		{
			"address":    "府谷县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	610823: {
		{
			"address":    "横山县",
			"start_year": "1999",
			"end_year":   "2014",
		},
	},
	610824: {
		{
			"address":    "靖边县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	610825: {
		{
			"address":    "定边县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	610826: {
		{
			"address":    "绥德县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	610827: {
		{
			"address":    "米脂县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	610828: {
		{
			"address":    "佳县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	610829: {
		{
			"address":    "吴堡县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	610830: {
		{
			"address":    "清涧县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	610831: {
		{
			"address":    "子洲县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	610881: {
		{
			"address":    "神木市",
			"start_year": "2017",
			"end_year":   "",
		},
	},
	610900: {
		{
			"address":    "安康市",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	610902: {
		{
			"address":    "汉滨区",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	610921: {
		{
			"address":    "汉阴县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	610922: {
		{
			"address":    "石泉县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	610923: {
		{
			"address":    "宁陕县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	610924: {
		{
			"address":    "紫阳县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	610925: {
		{
			"address":    "岚皋县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	610926: {
		{
			"address":    "平利县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	610927: {
		{
			"address":    "镇坪县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	610928: {
		{
			"address":    "旬阳县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	610929: {
		{
			"address":    "白河县",
			"start_year": "2000",
			"end_year":   "",
		},
	},
	611000: {
		{
			"address":    "商洛市",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	611002: {
		{
			"address":    "商州区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	611021: {
		{
			"address":    "洛南县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	611022: {
		{
			"address":    "丹凤县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	611023: {
		{
			"address":    "商南县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	611024: {
		{
			"address":    "山阳县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	611025: {
		{
			"address":    "镇安县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	611026: {
		{
			"address":    "柞水县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	612100: {
		{
			"address":    "渭南地区",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	612101: {
		{
			"address":    "渭南市",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	612102: {
		{
			"address":    "韩城市",
			"start_year": "1983",
			"end_year":   "1993",
		},
	},
	612103: {
		{
			"address":    "华阴市",
			"start_year": "1990",
			"end_year":   "1993",
		},
	},
	612121: {
		{
			"address":    "蓝田县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612122: {
		{
			"address":    "临潼县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612123: {
		{
			"address":    "渭南县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612124: {
		{
			"address":    "华县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	612125: {
		{
			"address":    "华阴县",
			"start_year": "",
			"end_year":   "1989",
		},
	},
	612126: {
		{
			"address":    "潼关县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	612127: {
		{
			"address":    "大荔县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	612128: {
		{
			"address":    "蒲城县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	612129: {
		{
			"address":    "澄城县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	612130: {
		{
			"address":    "白水县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	612131: {
		{
			"address":    "韩城县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612132: {
		{
			"address":    "合阳县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	612133: {
		{
			"address":    "富平县",
			"start_year": "",
			"end_year":   "1993",
		},
	},
	612200: {
		{
			"address":    "咸阳地区",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612201: {
		{
			"address":    "咸阳市",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612221: {
		{
			"address":    "兴平县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612222: {
		{
			"address":    "周至县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612223: {
		{
			"address":    "户县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612224: {
		{
			"address":    "三原县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612225: {
		{
			"address":    "泾阳县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612226: {
		{
			"address":    "高陵县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612227: {
		{
			"address":    "乾县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612228: {
		{
			"address":    "礼泉县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612229: {
		{
			"address":    "永寿县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612230: {
		{
			"address":    "彬县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612231: {
		{
			"address":    "长武县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612232: {
		{
			"address":    "旬邑县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612233: {
		{
			"address":    "淳化县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612300: {
		{
			"address":    "汉中地区",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612301: {
		{
			"address":    "汉中市",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612321: {
		{
			"address":    "南郑县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612322: {
		{
			"address":    "城固县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612323: {
		{
			"address":    "洋县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612324: {
		{
			"address":    "西乡县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612325: {
		{
			"address":    "勉县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612326: {
		{
			"address":    "宁强县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612327: {
		{
			"address":    "略阳县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612328: {
		{
			"address":    "镇巴县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612329: {
		{
			"address":    "留坝县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612330: {
		{
			"address":    "佛坪县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612400: {
		{
			"address":    "安康地区",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	612401: {
		{
			"address":    "安康市",
			"start_year": "1988",
			"end_year":   "1999",
		},
	},
	612421: {
		{
			"address":    "安康县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	612422: {
		{
			"address":    "汉阴县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	612423: {
		{
			"address":    "石泉县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	612424: {
		{
			"address":    "宁陕县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	612425: {
		{
			"address":    "紫阳县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	612426: {
		{
			"address":    "岚皋县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	612427: {
		{
			"address":    "平利县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	612428: {
		{
			"address":    "镇坪县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	612429: {
		{
			"address":    "旬阳县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	612430: {
		{
			"address":    "白河县",
			"start_year": "",
			"end_year":   "1999",
		},
	},
	612500: {
		{
			"address":    "商洛地区",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	612501: {
		{
			"address":    "商州市",
			"start_year": "1988",
			"end_year":   "2000",
		},
	},
	612521: {
		{
			"address":    "商县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	612522: {
		{
			"address":    "洛南县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	612523: {
		{
			"address":    "丹凤县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	612524: {
		{
			"address":    "商南县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	612525: {
		{
			"address":    "山阳县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	612526: {
		{
			"address":    "镇安县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	612527: {
		{
			"address":    "柞水县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	612600: {
		{
			"address":    "延安地区",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612601: {
		{
			"address":    "延安市",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612621: {
		{
			"address":    "延长县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612622: {
		{
			"address":    "延川县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612623: {
		{
			"address":    "子长县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612624: {
		{
			"address":    "安塞县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612625: {
		{
			"address":    "志丹县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612626: {
		{
			"address":    "吴旗县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612627: {
		{
			"address":    "甘泉县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612628: {
		{
			"address":    "富县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612629: {
		{
			"address":    "洛川县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612630: {
		{
			"address":    "宜川县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612631: {
		{
			"address":    "黄龙县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612632: {
		{
			"address":    "黄陵县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	612633: {
		{
			"address":    "宜君县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	612700: {
		{
			"address":    "榆林地区",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	612701: {
		{
			"address":    "榆林市",
			"start_year": "1988",
			"end_year":   "1998",
		},
	},
	612721: {
		{
			"address":    "榆林县",
			"start_year": "",
			"end_year":   "1987",
		},
	},
	612722: {
		{
			"address":    "神木县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	612723: {
		{
			"address":    "府谷县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	612724: {
		{
			"address":    "横山县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	612725: {
		{
			"address":    "靖边县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	612726: {
		{
			"address":    "定边县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	612727: {
		{
			"address":    "绥德县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	612728: {
		{
			"address":    "米脂县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	612729: {
		{
			"address":    "佳县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	612730: {
		{
			"address":    "吴堡县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	612731: {
		{
			"address":    "清涧县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	612732: {
		{
			"address":    "子洲县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	619001: {
		{
			"address":    "兴平市",
			"start_year": "1993",
			"end_year":   "1994",
		},
	},
	619002: {
		{
			"address":    "韩城市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	619003: {
		{
			"address":    "华阴市",
			"start_year": "1994",
			"end_year":   "1994",
		},
	},
	620000: {
		{
			"address":    "甘肃省",
			"start_year": "",
			"end_year":   "",
		},
	},
	620100: {
		{
			"address":    "兰州市",
			"start_year": "",
			"end_year":   "",
		},
	},
	620102: {
		{
			"address":    "城关区",
			"start_year": "",
			"end_year":   "",
		},
	},
	620103: {
		{
			"address":    "七里河区",
			"start_year": "",
			"end_year":   "",
		},
	},
	620104: {
		{
			"address":    "西固区",
			"start_year": "",
			"end_year":   "",
		},
	},
	620105: {
		{
			"address":    "安宁区",
			"start_year": "",
			"end_year":   "",
		},
	},
	620111: {
		{
			"address":    "红古区",
			"start_year": "",
			"end_year":   "",
		},
	},
	620112: {
		{
			"address":    "白银区",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	620121: {
		{
			"address":    "永登县",
			"start_year": "",
			"end_year":   "",
		},
	},
	620122: {
		{
			"address":    "皋兰县",
			"start_year": "",
			"end_year":   "",
		},
	},
	620123: {
		{
			"address":    "榆中县",
			"start_year": "",
			"end_year":   "",
		},
	},
	620200: {
		{
			"address":    "嘉峪关市",
			"start_year": "",
			"end_year":   "",
		},
	},
	620300: {
		{
			"address":    "金昌市",
			"start_year": "1981",
			"end_year":   "",
		},
	},
	620302: {
		{
			"address":    "金川区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	620321: {
		{
			"address":    "永昌县",
			"start_year": "1981",
			"end_year":   "",
		},
	},
	620400: {
		{
			"address":    "白银市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	620402: {
		{
			"address":    "白银区",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	620403: {
		{
			"address":    "平川区",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	620421: {
		{
			"address":    "靖远县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	620422: {
		{
			"address":    "会宁县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	620423: {
		{
			"address":    "景泰县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	620500: {
		{
			"address":    "天水市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	620502: {
		{
			"address":    "秦城区",
			"start_year": "1985",
			"end_year":   "2003",
		},
		{
			"address":    "秦州区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	620503: {
		{
			"address":    "北道区",
			"start_year": "1985",
			"end_year":   "2003",
		},
		{
			"address":    "麦积区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	620521: {
		{
			"address":    "清水县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	620522: {
		{
			"address":    "秦安县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	620523: {
		{
			"address":    "甘谷县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	620524: {
		{
			"address":    "武山县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	620525: {
		{
			"address":    "张家川回族自治县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	620600: {
		{
			"address":    "武威市",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	620602: {
		{
			"address":    "凉州区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	620621: {
		{
			"address":    "民勤县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	620622: {
		{
			"address":    "古浪县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	620623: {
		{
			"address":    "天祝藏族自治县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	620700: {
		{
			"address":    "张掖市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620702: {
		{
			"address":    "甘州区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620721: {
		{
			"address":    "肃南裕固族自治县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620722: {
		{
			"address":    "民乐县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620723: {
		{
			"address":    "临泽县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620724: {
		{
			"address":    "高台县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620725: {
		{
			"address":    "山丹县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620800: {
		{
			"address":    "平凉市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620802: {
		{
			"address":    "崆峒区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620821: {
		{
			"address":    "泾川县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620822: {
		{
			"address":    "灵台县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620823: {
		{
			"address":    "崇信县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620824: {
		{
			"address":    "华亭县",
			"start_year": "2002",
			"end_year":   "2017",
		},
	},
	620825: {
		{
			"address":    "庄浪县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620826: {
		{
			"address":    "静宁县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620881: {
		{
			"address":    "华亭市",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	620900: {
		{
			"address":    "酒泉市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620902: {
		{
			"address":    "肃州区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620921: {
		{
			"address":    "金塔县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620922: {
		{
			"address":    "瓜州县",
			"start_year": "2006",
			"end_year":   "",
		},
	},
	620923: {
		{
			"address":    "肃北蒙古族自治县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620924: {
		{
			"address":    "阿克塞哈萨克族自治县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620925: {
		{
			"address":    "安西县",
			"start_year": "2002",
			"end_year":   "2005",
		},
	},
	620981: {
		{
			"address":    "玉门市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	620982: {
		{
			"address":    "敦煌市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	621000: {
		{
			"address":    "庆阳市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	621002: {
		{
			"address":    "西峰区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	621021: {
		{
			"address":    "庆城县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	621022: {
		{
			"address":    "环县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	621023: {
		{
			"address":    "华池县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	621024: {
		{
			"address":    "合水县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	621025: {
		{
			"address":    "正宁县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	621026: {
		{
			"address":    "宁县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	621027: {
		{
			"address":    "镇原县",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	621100: {
		{
			"address":    "定西市",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	621102: {
		{
			"address":    "安定区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	621121: {
		{
			"address":    "通渭县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	621122: {
		{
			"address":    "陇西县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	621123: {
		{
			"address":    "渭源县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	621124: {
		{
			"address":    "临洮县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	621125: {
		{
			"address":    "漳县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	621126: {
		{
			"address":    "岷县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	621200: {
		{
			"address":    "陇南市",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	621202: {
		{
			"address":    "武都区",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	621221: {
		{
			"address":    "成县",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	621222: {
		{
			"address":    "文县",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	621223: {
		{
			"address":    "宕昌县",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	621224: {
		{
			"address":    "康县",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	621225: {
		{
			"address":    "西和县",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	621226: {
		{
			"address":    "礼县",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	621227: {
		{
			"address":    "徽县",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	621228: {
		{
			"address":    "两当县",
			"start_year": "2004",
			"end_year":   "",
		},
	},
	622100: {
		{
			"address":    "酒泉地区",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622101: {
		{
			"address":    "玉门市",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622102: {
		{
			"address":    "酒泉市",
			"start_year": "1985",
			"end_year":   "2001",
		},
	},
	622103: {
		{
			"address":    "敦煌市",
			"start_year": "1987",
			"end_year":   "2001",
		},
	},
	622121: {
		{
			"address":    "酒泉县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622122: {
		{
			"address":    "敦煌县",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	622123: {
		{
			"address":    "金塔县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622124: {
		{
			"address":    "肃北蒙古族自治县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622125: {
		{
			"address":    "阿克塞哈萨克族自治县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622126: {
		{
			"address":    "安西县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622200: {
		{
			"address":    "张掖地区",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622201: {
		{
			"address":    "张掖市",
			"start_year": "1985",
			"end_year":   "2001",
		},
	},
	622221: {
		{
			"address":    "张掖县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622222: {
		{
			"address":    "肃南裕固族自治县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622223: {
		{
			"address":    "民乐县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622224: {
		{
			"address":    "临泽县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622225: {
		{
			"address":    "高台县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622226: {
		{
			"address":    "山丹县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622300: {
		{
			"address":    "武威地区",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	622301: {
		{
			"address":    "武威市",
			"start_year": "1985",
			"end_year":   "2000",
		},
	},
	622321: {
		{
			"address":    "武威县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622322: {
		{
			"address":    "民勤县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	622323: {
		{
			"address":    "古浪县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	622324: {
		{
			"address":    "景泰县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622325: {
		{
			"address":    "永昌县",
			"start_year": "",
			"end_year":   "1980",
		},
	},
	622326: {
		{
			"address":    "天祝藏族自治县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	622400: {
		{
			"address":    "定西地区",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	622421: {
		{
			"address":    "定西县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	622422: {
		{
			"address":    "靖远县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622423: {
		{
			"address":    "会宁县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622424: {
		{
			"address":    "通渭县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	622425: {
		{
			"address":    "陇西县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	622426: {
		{
			"address":    "渭源县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	622427: {
		{
			"address":    "临洮县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	622428: {
		{
			"address":    "漳县",
			"start_year": "1985",
			"end_year":   "2002",
		},
	},
	622429: {
		{
			"address":    "岷县",
			"start_year": "1985",
			"end_year":   "2002",
		},
	},
	622500: {
		{
			"address":    "天水地区",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622501: {
		{
			"address":    "天水市",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622521: {
		{
			"address":    "张家川回族自治县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622522: {
		{
			"address":    "天水县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622523: {
		{
			"address":    "清水县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622524: {
		{
			"address":    "徽县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622525: {
		{
			"address":    "两当县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622526: {
		{
			"address":    "礼县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622527: {
		{
			"address":    "西和县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622528: {
		{
			"address":    "武山县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622529: {
		{
			"address":    "甘谷县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622530: {
		{
			"address":    "秦安县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622531: {
		{
			"address":    "漳县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622600: {
		{
			"address":    "武都地区",
			"start_year": "",
			"end_year":   "1984",
		},
		{
			"address":    "陇南地区",
			"start_year": "1985",
			"end_year":   "2003",
		},
	},
	622621: {
		{
			"address":    "武都县",
			"start_year": "",
			"end_year":   "2003",
		},
	},
	622622: {
		{
			"address":    "岷县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	622623: {
		{
			"address":    "宕昌县",
			"start_year": "",
			"end_year":   "2003",
		},
	},
	622624: {
		{
			"address":    "成县",
			"start_year": "",
			"end_year":   "2003",
		},
	},
	622625: {
		{
			"address":    "康县",
			"start_year": "",
			"end_year":   "2003",
		},
	},
	622626: {
		{
			"address":    "文县",
			"start_year": "",
			"end_year":   "2003",
		},
	},
	622627: {
		{
			"address":    "西和县",
			"start_year": "1985",
			"end_year":   "2003",
		},
	},
	622628: {
		{
			"address":    "礼县",
			"start_year": "1985",
			"end_year":   "2003",
		},
	},
	622629: {
		{
			"address":    "两当县",
			"start_year": "1985",
			"end_year":   "2003",
		},
	},
	622630: {
		{
			"address":    "徽县",
			"start_year": "1985",
			"end_year":   "2003",
		},
	},
	622700: {
		{
			"address":    "平凉地区",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622701: {
		{
			"address":    "平凉市",
			"start_year": "1983",
			"end_year":   "2001",
		},
	},
	622721: {
		{
			"address":    "平凉县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	622722: {
		{
			"address":    "泾川县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622723: {
		{
			"address":    "灵台县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622724: {
		{
			"address":    "崇信县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622725: {
		{
			"address":    "华亭县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622726: {
		{
			"address":    "庄浪县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622727: {
		{
			"address":    "静宁县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622800: {
		{
			"address":    "庆阳地区",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622801: {
		{
			"address":    "西峰市",
			"start_year": "1985",
			"end_year":   "2001",
		},
	},
	622821: {
		{
			"address":    "庆阳县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622822: {
		{
			"address":    "环县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622823: {
		{
			"address":    "华池县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622824: {
		{
			"address":    "合水县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622825: {
		{
			"address":    "正宁县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622826: {
		{
			"address":    "宁县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622827: {
		{
			"address":    "镇原县",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	622900: {
		{
			"address":    "临夏回族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	622901: {
		{
			"address":    "临夏市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	622921: {
		{
			"address":    "临夏县",
			"start_year": "",
			"end_year":   "",
		},
	},
	622922: {
		{
			"address":    "康乐县",
			"start_year": "",
			"end_year":   "",
		},
	},
	622923: {
		{
			"address":    "永靖县",
			"start_year": "",
			"end_year":   "",
		},
	},
	622924: {
		{
			"address":    "广河县",
			"start_year": "",
			"end_year":   "",
		},
	},
	622925: {
		{
			"address":    "和政县",
			"start_year": "",
			"end_year":   "",
		},
	},
	622926: {
		{
			"address":    "东乡族自治县",
			"start_year": "",
			"end_year":   "",
		},
	},
	622927: {
		{
			"address":    "积石山保安族东乡族撒拉族自治县",
			"start_year": "",
			"end_year":   "",
		},
	},
	623000: {
		{
			"address":    "甘南藏族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	623001: {
		{
			"address":    "合作市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	623021: {
		{
			"address":    "临潭县",
			"start_year": "",
			"end_year":   "",
		},
	},
	623022: {
		{
			"address":    "卓尼县",
			"start_year": "",
			"end_year":   "",
		},
	},
	623023: {
		{
			"address":    "舟曲县",
			"start_year": "",
			"end_year":   "",
		},
	},
	623024: {
		{
			"address":    "迭部县",
			"start_year": "",
			"end_year":   "",
		},
	},
	623025: {
		{
			"address":    "玛曲县",
			"start_year": "",
			"end_year":   "",
		},
	},
	623026: {
		{
			"address":    "碌曲县",
			"start_year": "",
			"end_year":   "",
		},
	},
	623027: {
		{
			"address":    "夏河县",
			"start_year": "",
			"end_year":   "",
		},
	},
	630000: {
		{
			"address":    "青海省",
			"start_year": "",
			"end_year":   "",
		},
	},
	630100: {
		{
			"address":    "西宁市",
			"start_year": "",
			"end_year":   "",
		},
	},
	630102: {
		{
			"address":    "城东区",
			"start_year": "",
			"end_year":   "",
		},
	},
	630103: {
		{
			"address":    "城中区",
			"start_year": "",
			"end_year":   "",
		},
	},
	630104: {
		{
			"address":    "城西区",
			"start_year": "",
			"end_year":   "",
		},
	},
	630105: {
		{
			"address":    "城北区",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	630106: {
		{
			"address":    "湟中区",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	630111: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	630121: {
		{
			"address":    "大通县",
			"start_year": "",
			"end_year":   "1984",
		},
		{
			"address":    "大通回族土族自治县",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	630122: {
		{
			"address":    "湟中县",
			"start_year": "1999",
			"end_year":   "2018",
		},
	},
	630123: {
		{
			"address":    "湟源县",
			"start_year": "1999",
			"end_year":   "",
		},
	},
	630200: {
		{
			"address":    "海东市",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	630202: {
		{
			"address":    "乐都区",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	630203: {
		{
			"address":    "平安区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	630221: {
		{
			"address":    "平安县",
			"start_year": "2013",
			"end_year":   "2014",
		},
	},
	630222: {
		{
			"address":    "民和回族土族自治县",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	630223: {
		{
			"address":    "互助土族自治县",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	630224: {
		{
			"address":    "化隆回族自治县",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	630225: {
		{
			"address":    "循化撒拉族自治县",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	632100: {
		{
			"address":    "海东地区",
			"start_year": "",
			"end_year":   "2012",
		},
	},
	632121: {
		{
			"address":    "平安县",
			"start_year": "",
			"end_year":   "2012",
		},
	},
	632122: {
		{
			"address":    "民和县",
			"start_year": "",
			"end_year":   "1984",
		},
		{
			"address":    "民和回族土族自治县",
			"start_year": "1985",
			"end_year":   "2012",
		},
	},
	632123: {
		{
			"address":    "乐都县",
			"start_year": "",
			"end_year":   "2012",
		},
	},
	632124: {
		{
			"address":    "湟中县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	632125: {
		{
			"address":    "湟源县",
			"start_year": "",
			"end_year":   "1998",
		},
	},
	632126: {
		{
			"address":    "互助土族自治县",
			"start_year": "",
			"end_year":   "2012",
		},
	},
	632127: {
		{
			"address":    "化隆回族自治县",
			"start_year": "",
			"end_year":   "2012",
		},
	},
	632128: {
		{
			"address":    "循化撒拉族自治县",
			"start_year": "",
			"end_year":   "2012",
		},
	},
	632200: {
		{
			"address":    "海北藏族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	632221: {
		{
			"address":    "门源回族自治县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632222: {
		{
			"address":    "祁连县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632223: {
		{
			"address":    "海晏县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632224: {
		{
			"address":    "刚察县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632300: {
		{
			"address":    "黄南藏族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	632301: {
		{
			"address":    "同仁市",
			"start_year": "2020",
			"end_year":   "",
		},
	},
	632321: {
		{
			"address":    "同仁县",
			"start_year": "",
			"end_year":   "2019",
		},
	},
	632322: {
		{
			"address":    "尖扎县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632323: {
		{
			"address":    "泽库县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632324: {
		{
			"address":    "河南蒙古族自治县",
			"start_year": "1981",
			"end_year":   "",
		},
	},
	632421: {
		{
			"address":    "河南蒙古族自治县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	632500: {
		{
			"address":    "海南藏族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	632521: {
		{
			"address":    "共和县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632522: {
		{
			"address":    "同德县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632523: {
		{
			"address":    "贵德县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632524: {
		{
			"address":    "兴海县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632525: {
		{
			"address":    "贵南县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632600: {
		{
			"address":    "果洛藏族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	632621: {
		{
			"address":    "玛沁县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632622: {
		{
			"address":    "班玛县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632623: {
		{
			"address":    "甘德县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632624: {
		{
			"address":    "达日县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632625: {
		{
			"address":    "久治县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632626: {
		{
			"address":    "玛多县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632700: {
		{
			"address":    "玉树藏族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	632701: {
		{
			"address":    "玉树市",
			"start_year": "2013",
			"end_year":   "",
		},
	},
	632721: {
		{
			"address":    "玉树县",
			"start_year": "",
			"end_year":   "2012",
		},
	},
	632722: {
		{
			"address":    "杂多县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632723: {
		{
			"address":    "称多县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632724: {
		{
			"address":    "治多县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632725: {
		{
			"address":    "囊谦县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632726: {
		{
			"address":    "曲麻莱县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632800: {
		{
			"address":    "海西蒙古族藏族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	632801: {
		{
			"address":    "格尔木市",
			"start_year": "",
			"end_year":   "",
		},
	},
	632802: {
		{
			"address":    "德令哈市",
			"start_year": "1988",
			"end_year":   "",
		},
	},
	632803: {
		{
			"address":    "茫崖市",
			"start_year": "2018",
			"end_year":   "",
		},
	},
	632821: {
		{
			"address":    "乌兰县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632822: {
		{
			"address":    "都兰县",
			"start_year": "",
			"end_year":   "",
		},
	},
	632823: {
		{
			"address":    "天峻县",
			"start_year": "",
			"end_year":   "",
		},
	},
	640000: {
		{
			"address":    "宁夏回族自治区",
			"start_year": "",
			"end_year":   "",
		},
	},
	640100: {
		{
			"address":    "银川市",
			"start_year": "",
			"end_year":   "",
		},
	},
	640102: {
		{
			"address":    "城区",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	640103: {
		{
			"address":    "新城区",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	640104: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "2001",
		},
		{
			"address":    "兴庆区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	640105: {
		{
			"address":    "西夏区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	640106: {
		{
			"address":    "金凤区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	640121: {
		{
			"address":    "永宁县",
			"start_year": "",
			"end_year":   "",
		},
	},
	640122: {
		{
			"address":    "贺兰县",
			"start_year": "",
			"end_year":   "",
		},
	},
	640181: {
		{
			"address":    "灵武市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	640200: {
		{
			"address":    "石咀山市",
			"start_year": "",
			"end_year":   "1986",
		},
		{
			"address":    "石嘴山市",
			"start_year": "1987",
			"end_year":   "",
		},
	},
	640202: {
		{
			"address":    "大武口区",
			"start_year": "",
			"end_year":   "",
		},
	},
	640204: {
		{
			"address":    "石炭井区",
			"start_year": "",
			"end_year":   "2001",
		},
	},
	640205: {
		{
			"address":    "石咀山区",
			"start_year": "",
			"end_year":   "1986",
		},
		{
			"address":    "石嘴山区",
			"start_year": "1987",
			"end_year":   "2002",
		},
		{
			"address":    "惠农区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	640211: {
		{
			"address":    "郊区",
			"start_year": "",
			"end_year":   "1986",
		},
	},
	640221: {
		{
			"address":    "平罗县",
			"start_year": "",
			"end_year":   "",
		},
	},
	640222: {
		{
			"address":    "陶乐县",
			"start_year": "",
			"end_year":   "2002",
		},
	},
	640223: {
		{
			"address":    "惠农县",
			"start_year": "1987",
			"end_year":   "2002",
		},
	},
	640300: {
		{
			"address":    "吴忠市",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	640302: {
		{
			"address":    "利通区",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	640303: {
		{
			"address":    "红寺堡区",
			"start_year": "2009",
			"end_year":   "",
		},
	},
	640321: {
		{
			"address":    "中卫县",
			"start_year": "1998",
			"end_year":   "2002",
		},
	},
	640322: {
		{
			"address":    "中宁县",
			"start_year": "1998",
			"end_year":   "2002",
		},
	},
	640323: {
		{
			"address":    "盐池县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	640324: {
		{
			"address":    "同心县",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	640381: {
		{
			"address":    "青铜峡市",
			"start_year": "1998",
			"end_year":   "",
		},
	},
	640382: {
		{
			"address":    "灵武市",
			"start_year": "1998",
			"end_year":   "2001",
		},
	},
	640400: {
		{
			"address":    "固原市",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	640402: {
		{
			"address":    "原州区",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	640421: {
		{
			"address":    "海原县",
			"start_year": "2001",
			"end_year":   "2002",
		},
	},
	640422: {
		{
			"address":    "西吉县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	640423: {
		{
			"address":    "隆德县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	640424: {
		{
			"address":    "泾源县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	640425: {
		{
			"address":    "彭阳县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	640500: {
		{
			"address":    "中卫市",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	640502: {
		{
			"address":    "沙坡头区",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	640521: {
		{
			"address":    "中宁县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	640522: {
		{
			"address":    "海原县",
			"start_year": "2003",
			"end_year":   "",
		},
	},
	642100: {
		{
			"address":    "银南地区",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	642101: {
		{
			"address":    "吴忠市",
			"start_year": "1983",
			"end_year":   "1997",
		},
	},
	642102: {
		{
			"address":    "青铜峡市",
			"start_year": "1984",
			"end_year":   "1997",
		},
	},
	642103: {
		{
			"address":    "灵武市",
			"start_year": "1996",
			"end_year":   "1997",
		},
	},
	642121: {
		{
			"address":    "吴忠县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	642122: {
		{
			"address":    "青铜峡县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	642123: {
		{
			"address":    "中卫县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	642124: {
		{
			"address":    "中宁县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	642125: {
		{
			"address":    "灵武县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	642126: {
		{
			"address":    "盐池县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	642127: {
		{
			"address":    "同心县",
			"start_year": "",
			"end_year":   "1997",
		},
	},
	642200: {
		{
			"address":    "固原地区",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	642221: {
		{
			"address":    "固原县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	642222: {
		{
			"address":    "海原县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	642223: {
		{
			"address":    "西吉县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	642224: {
		{
			"address":    "隆德县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	642225: {
		{
			"address":    "泾源县",
			"start_year": "",
			"end_year":   "2000",
		},
	},
	642226: {
		{
			"address":    "彭阳县",
			"start_year": "1983",
			"end_year":   "2000",
		},
	},
	650000: {
		{
			"address":    "新疆维吾尔自治区",
			"start_year": "",
			"end_year":   "",
		},
	},
	650100: {
		{
			"address":    "乌鲁木齐市",
			"start_year": "",
			"end_year":   "",
		},
	},
	650102: {
		{
			"address":    "天山区",
			"start_year": "",
			"end_year":   "",
		},
	},
	650103: {
		{
			"address":    "沙依巴克区",
			"start_year": "",
			"end_year":   "",
		},
	},
	650104: {
		{
			"address":    "新市区",
			"start_year": "",
			"end_year":   "",
		},
	},
	650105: {
		{
			"address":    "水磨沟区",
			"start_year": "",
			"end_year":   "",
		},
	},
	650106: {
		{
			"address":    "头屯河区",
			"start_year": "",
			"end_year":   "",
		},
	},
	650107: {
		{
			"address":    "南山区",
			"start_year": "",
			"end_year":   "1988",
		},
		{
			"address":    "南山矿区",
			"start_year": "1989",
			"end_year":   "1998",
		},
		{
			"address":    "南泉区",
			"start_year": "1999",
			"end_year":   "2001",
		},
		{
			"address":    "达坂城区",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	650108: {
		{
			"address":    "东山区",
			"start_year": "1987",
			"end_year":   "2006",
		},
	},
	650109: {
		{
			"address":    "米东区",
			"start_year": "2007",
			"end_year":   "",
		},
	},
	650121: {
		{
			"address":    "乌鲁木齐县",
			"start_year": "",
			"end_year":   "",
		},
	},
	650200: {
		{
			"address":    "克拉玛依市",
			"start_year": "",
			"end_year":   "",
		},
	},
	650202: {
		{
			"address":    "独山子区",
			"start_year": "",
			"end_year":   "",
		},
	},
	650203: {
		{
			"address":    "克拉玛依区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	650204: {
		{
			"address":    "白碱滩区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	650205: {
		{
			"address":    "乌尔禾区",
			"start_year": "1982",
			"end_year":   "",
		},
	},
	650300: {
		{
			"address":    "石河子市",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	650400: {
		{
			"address":    "吐鲁番市",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	650402: {
		{
			"address":    "高昌区",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	650421: {
		{
			"address":    "鄯善县",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	650422: {
		{
			"address":    "托克逊县",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	650500: {
		{
			"address":    "哈密市",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	650502: {
		{
			"address":    "伊州区",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	650521: {
		{
			"address":    "巴里坤哈萨克自治县",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	650522: {
		{
			"address":    "伊吾县",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	652100: {
		{
			"address":    "吐鲁番地区",
			"start_year": "",
			"end_year":   "2014",
		},
	},
	652101: {
		{
			"address":    "吐鲁番市",
			"start_year": "1984",
			"end_year":   "2014",
		},
	},
	652121: {
		{
			"address":    "吐鲁番县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652122: {
		{
			"address":    "鄯善县",
			"start_year": "",
			"end_year":   "2014",
		},
	},
	652123: {
		{
			"address":    "托克逊县",
			"start_year": "",
			"end_year":   "2014",
		},
	},
	652200: {
		{
			"address":    "哈密地区",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	652201: {
		{
			"address":    "哈密市",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	652221: {
		{
			"address":    "哈密县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	652222: {
		{
			"address":    "巴里坤哈萨克自治县",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	652223: {
		{
			"address":    "伊吾县",
			"start_year": "",
			"end_year":   "2015",
		},
	},
	652300: {
		{
			"address":    "昌吉回族自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	652301: {
		{
			"address":    "昌吉市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	652302: {
		{
			"address":    "阜康市",
			"start_year": "1992",
			"end_year":   "",
		},
	},
	652303: {
		{
			"address":    "米泉市",
			"start_year": "1996",
			"end_year":   "2006",
		},
	},
	652321: {
		{
			"address":    "昌吉县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	652322: {
		{
			"address":    "米泉县",
			"start_year": "",
			"end_year":   "1995",
		},
	},
	652323: {
		{
			"address":    "呼图壁县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652324: {
		{
			"address":    "玛纳斯县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652325: {
		{
			"address":    "奇台县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652326: {
		{
			"address":    "阜康县",
			"start_year": "",
			"end_year":   "1991",
		},
	},
	652327: {
		{
			"address":    "吉木萨尔县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652328: {
		{
			"address":    "木垒哈萨克自治县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652400: {
		{
			"address":    "伊犁哈萨克自治州",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652401: {
		{
			"address":    "奎屯市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "伊宁市",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	652402: {
		{
			"address":    "伊宁市",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "一区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	652403: {
		{
			"address":    "一区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "二区",
			"start_year": "1982",
			"end_year":   "1982",
		},
	},
	652404: {
		{
			"address":    "二区",
			"start_year": "",
			"end_year":   "1981",
		},
		{
			"address":    "奎屯市",
			"start_year": "1982",
			"end_year":   "1983",
		},
	},
	652421: {
		{
			"address":    "伊宁县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652422: {
		{
			"address":    "察布查尔锡伯自治县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652423: {
		{
			"address":    "霍城县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652424: {
		{
			"address":    "巩留县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652425: {
		{
			"address":    "新源县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652426: {
		{
			"address":    "昭苏县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652427: {
		{
			"address":    "特克斯县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652428: {
		{
			"address":    "尼勒克县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652500: {
		{
			"address":    "塔城地区",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652521: {
		{
			"address":    "塔城县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652522: {
		{
			"address":    "额敏县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652523: {
		{
			"address":    "乌苏县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652524: {
		{
			"address":    "沙湾县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652525: {
		{
			"address":    "托里县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652526: {
		{
			"address":    "裕民县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652527: {
		{
			"address":    "和布克赛尔蒙古自治县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652600: {
		{
			"address":    "阿勒泰地区",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652621: {
		{
			"address":    "阿勒泰县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652622: {
		{
			"address":    "布尔津县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652623: {
		{
			"address":    "富蕴县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652624: {
		{
			"address":    "福海县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652625: {
		{
			"address":    "哈巴河县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652626: {
		{
			"address":    "青河县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652627: {
		{
			"address":    "吉木乃县",
			"start_year": "",
			"end_year":   "1983",
		},
	},
	652700: {
		{
			"address":    "博尔塔拉蒙古自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	652701: {
		{
			"address":    "博乐市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	652702: {
		{
			"address":    "阿拉山口市",
			"start_year": "2012",
			"end_year":   "",
		},
	},
	652721: {
		{
			"address":    "博乐县",
			"start_year": "",
			"end_year":   "1984",
		},
	},
	652722: {
		{
			"address":    "精河县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652723: {
		{
			"address":    "温泉县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652800: {
		{
			"address":    "巴音郭楞蒙古自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	652801: {
		{
			"address":    "库尔勒市",
			"start_year": "",
			"end_year":   "",
		},
	},
	652821: {
		{
			"address":    "库尔勒县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	652822: {
		{
			"address":    "轮台县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652823: {
		{
			"address":    "尉犁县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652824: {
		{
			"address":    "若羌县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652825: {
		{
			"address":    "且末县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652826: {
		{
			"address":    "焉耆回族自治县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652827: {
		{
			"address":    "和静县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652828: {
		{
			"address":    "和硕县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652829: {
		{
			"address":    "博湖县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652900: {
		{
			"address":    "阿克苏地区",
			"start_year": "",
			"end_year":   "",
		},
	},
	652901: {
		{
			"address":    "阿克苏市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	652902: {
		{
			"address":    "库车市",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	652921: {
		{
			"address":    "阿克苏县",
			"start_year": "",
			"end_year":   "1982",
		},
	},
	652922: {
		{
			"address":    "温宿县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652923: {
		{
			"address":    "库车县",
			"start_year": "",
			"end_year":   "2018",
		},
	},
	652924: {
		{
			"address":    "沙雅县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652925: {
		{
			"address":    "新和县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652926: {
		{
			"address":    "拜城县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652927: {
		{
			"address":    "乌什县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652928: {
		{
			"address":    "阿瓦提县",
			"start_year": "",
			"end_year":   "",
		},
	},
	652929: {
		{
			"address":    "柯坪县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653000: {
		{
			"address":    "克孜勒苏柯尔克孜自治州",
			"start_year": "",
			"end_year":   "",
		},
	},
	653001: {
		{
			"address":    "阿图什市",
			"start_year": "1986",
			"end_year":   "",
		},
	},
	653021: {
		{
			"address":    "阿图什县",
			"start_year": "",
			"end_year":   "1985",
		},
	},
	653022: {
		{
			"address":    "阿克陶县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653023: {
		{
			"address":    "阿合奇县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653024: {
		{
			"address":    "乌恰县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653100: {
		{
			"address":    "喀什地区",
			"start_year": "",
			"end_year":   "",
		},
	},
	653101: {
		{
			"address":    "喀什市",
			"start_year": "",
			"end_year":   "",
		},
	},
	653121: {
		{
			"address":    "疏附县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653122: {
		{
			"address":    "疏勒县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653123: {
		{
			"address":    "英吉沙县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653124: {
		{
			"address":    "泽普县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653125: {
		{
			"address":    "莎车县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653126: {
		{
			"address":    "叶城县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653127: {
		{
			"address":    "麦盖提县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653128: {
		{
			"address":    "岳普湖县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653129: {
		{
			"address":    "伽师县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653130: {
		{
			"address":    "巴楚县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653131: {
		{
			"address":    "塔什库尔干塔吉克自治县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653200: {
		{
			"address":    "和田地区",
			"start_year": "",
			"end_year":   "",
		},
	},
	653201: {
		{
			"address":    "和田市",
			"start_year": "1983",
			"end_year":   "",
		},
	},
	653221: {
		{
			"address":    "和田县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653222: {
		{
			"address":    "墨玉县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653223: {
		{
			"address":    "皮山县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653224: {
		{
			"address":    "洛浦县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653225: {
		{
			"address":    "策勒县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653226: {
		{
			"address":    "于田县",
			"start_year": "",
			"end_year":   "",
		},
	},
	653227: {
		{
			"address":    "民丰县",
			"start_year": "",
			"end_year":   "",
		},
	},
	654000: {
		{
			"address":    "伊犁哈萨克自治州",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	654001: {
		{
			"address":    "奎屯市",
			"start_year": "1984",
			"end_year":   "2000",
		},
	},
	654002: {
		{
			"address":    "伊宁市",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	654003: {
		{
			"address":    "奎屯市",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	654004: {
		{
			"address":    "霍尔果斯市",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	654021: {
		{
			"address":    "伊宁县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	654022: {
		{
			"address":    "察布查尔锡伯自治县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	654023: {
		{
			"address":    "霍城县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	654024: {
		{
			"address":    "巩留县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	654025: {
		{
			"address":    "新源县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	654026: {
		{
			"address":    "昭苏县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	654027: {
		{
			"address":    "特克斯县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	654028: {
		{
			"address":    "尼勒克县",
			"start_year": "2001",
			"end_year":   "",
		},
	},
	654100: {
		{
			"address":    "伊犁地区",
			"start_year": "1984",
			"end_year":   "2000",
		},
	},
	654101: {
		{
			"address":    "伊宁市",
			"start_year": "1984",
			"end_year":   "2000",
		},
	},
	654121: {
		{
			"address":    "伊宁县",
			"start_year": "1984",
			"end_year":   "2000",
		},
	},
	654122: {
		{
			"address":    "察布查尔锡伯自治县",
			"start_year": "1984",
			"end_year":   "2000",
		},
	},
	654123: {
		{
			"address":    "霍城县",
			"start_year": "1984",
			"end_year":   "2000",
		},
	},
	654124: {
		{
			"address":    "巩留县",
			"start_year": "1984",
			"end_year":   "2000",
		},
	},
	654125: {
		{
			"address":    "新源县",
			"start_year": "1984",
			"end_year":   "2000",
		},
	},
	654126: {
		{
			"address":    "昭苏县",
			"start_year": "1984",
			"end_year":   "2000",
		},
	},
	654127: {
		{
			"address":    "特克斯县",
			"start_year": "1984",
			"end_year":   "2000",
		},
	},
	654128: {
		{
			"address":    "尼勒克县",
			"start_year": "1984",
			"end_year":   "2000",
		},
	},
	654200: {
		{
			"address":    "塔城地区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	654201: {
		{
			"address":    "塔城市",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	654202: {
		{
			"address":    "乌苏市",
			"start_year": "1996",
			"end_year":   "",
		},
	},
	654221: {
		{
			"address":    "额敏县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	654222: {
		{
			"address":    "乌苏县",
			"start_year": "1984",
			"end_year":   "1995",
		},
	},
	654223: {
		{
			"address":    "沙湾县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	654224: {
		{
			"address":    "托里县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	654225: {
		{
			"address":    "裕民县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	654226: {
		{
			"address":    "和布克赛尔蒙古自治县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	654300: {
		{
			"address":    "阿勒泰地区",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	654301: {
		{
			"address":    "阿勒泰市",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	654321: {
		{
			"address":    "布尔津县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	654322: {
		{
			"address":    "富蕴县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	654323: {
		{
			"address":    "福海县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	654324: {
		{
			"address":    "哈巴河县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	654325: {
		{
			"address":    "青河县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	654326: {
		{
			"address":    "吉木乃县",
			"start_year": "1984",
			"end_year":   "",
		},
	},
	659001: {
		{
			"address":    "石河子市",
			"start_year": "1985",
			"end_year":   "",
		},
	},
	659002: {
		{
			"address":    "阿拉尔市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	659003: {
		{
			"address":    "图木舒克市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	659004: {
		{
			"address":    "五家渠市",
			"start_year": "2002",
			"end_year":   "",
		},
	},
	659005: {
		{
			"address":    "北屯市",
			"start_year": "2011",
			"end_year":   "",
		},
	},
	659006: {
		{
			"address":    "铁门关市",
			"start_year": "2012",
			"end_year":   "",
		},
	},
	659007: {
		{
			"address":    "双河市",
			"start_year": "2014",
			"end_year":   "",
		},
	},
	659008: {
		{
			"address":    "可克达拉市",
			"start_year": "2015",
			"end_year":   "",
		},
	},
	659009: {
		{
			"address":    "昆玉市",
			"start_year": "2016",
			"end_year":   "",
		},
	},
	659010: {
		{
			"address":    "胡杨河市",
			"start_year": "2019",
			"end_year":   "",
		},
	},
	810000: {
		{
			"address":    "香港特别行政区",
			"start_year": "",
			"end_year":   "",
		},
	},
	820000: {
		{
			"address":    "澳门特别行政区",
			"start_year": "",
			"end_year":   "",
		},
	},
	830000: {
		{
			"address":    "台湾省",
			"start_year": "",
			"end_year":   "",
		},
	},
}
