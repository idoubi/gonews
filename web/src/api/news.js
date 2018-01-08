import Mock from 'mockjs'

let Random = Mock.Random
let per = 20  // 每页显示数据数量 
let total = 101  // 数据总量
function getNews() {
	return Mock.mock({
		'total': total,
		'per': per,
		'items|20': [
		    {
		      	'id': function() {
		        	return Random.int(1, 5)
		      	},
		      	url: function() {
		        	return Random.url('http')
		      	},
		      	title: function() {
		        	return Random.cword(20, 40)
		      	},
		      	ctime: function() {
		        	return Random.date('yyyyMMdd')
		      	}
		    }
		  ]
	})
}
Mock.mock(/api\/news\?page=\d+/, function(options) {
	return getNews()
})
