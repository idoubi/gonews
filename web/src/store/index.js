import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import moment from 'moment'
// import '../api/news'  

Vue.use(Vuex)

export function createStore() {
	return new Vuex.Store({
		state: {
			page: 1,
			keyword: '',
			newsTotal: 0,
			newsPer: 10,
			newsItems: []
		},
		mutations: {
			SET_PAGE(state, page) {  // 设置页面
				state.page = page
			},
			SET_KEYWORD(state, keyword) {  // 设置搜索关键词
				state.keyword = keyword
			},
			SET_NEWS(state) {   // 设置新闻数据
				let page = state.page || 1
				let keyword = state.keyword
				var url = '/api/news?page=' + page + '&keyword=' + keyword
				axios.get(url).then(function(resp) {
				  	var items = resp.data.items
				  	var news = []
				  	for (let i = 0, len = items.length; i < len; i++) {
				    	var sNews = {
					      	id: items[i].id,
					      	title: items[i].title,
					      	url: items[i].link,
					      	ctime: moment(items[i].ctime, 'YYYYMMDD').fromNow(),
					      	weekday: moment(items[i].ctime).format('dddd').substr(0, 3)
					    }
				    	news.push(sNews)
				  	}
				  	state.newsTotal = resp.data.total
				  	state.newsPer = resp.data.per
				  	state.newsItems = news
				}).catch(function() {
				  	console.log('error')
				})
			}
		}
	})
}
