<template>
  <div class="news-view">
    <div class="news-list-nav">
      <a v-if="page > 1" @click="changePage(page - 1)">&lt; prev</a>
      <a v-else class="disabled">&lt; prev</a>
      <span>{{ page }}/{{ maxPage }}</span>
      <a v-if="hasMore" @click="changePage(page + 1)">more &gt;</a>
      <a v-else class="disabled">more &gt;</a>
    </div>
    <transition :name="transition">
      <div class="news-list" :key="page" v-if="page > 0">
        <transition-group tag="ul" name="item">
          <Item v-for="item in newsItems" :key="item.title+item.id" :item="item"></Item>
        </transition-group>
        <div id="copyright">
          Powered by ©
          <a href="https://idoubi.cc" target="_blank">艾逗笔</a>
        </div>
      </div>
    </transition>
  </div>
</template>

<script>
import Item from "../components/Item.vue";
export default {
  name: "item-list",
  components: {
    Item
  },
  props: {},
  data() {
    return {
      transition: "slide-right",
      type: "news"
    };
  },
  computed: {
    maxPage() {
      return Math.ceil(this.newsTotal / this.newsPer);
    },
    hasMore() {
      return this.page < this.maxPage;
    },
    page() {
      return this.$store.state.page;
    },
    newsTotal() {
      return this.$store.state.newsTotal;
    },
    newsPer() {
      return this.$store.state.newsPer;
    },
    newsItems() {
      return this.$store.state.newsItems;
    }
  },
  methods: {
    changePage(page) {
      this.$store.commit("SET_PAGE", page);
      this.$store.commit("SET_NEWS");
    }
  }
};
</script>

<style lang="stylus">
#copyright {
  position: absolute;
  width: 100%;
  text-align: center;
  height: 60px;
  line-height: 60px;

  a {
    color: #499ef3;
  }
}

.news-view {
  padding-top: 45px;
}

.news-list-nav, .news-list {
  background-color: #fff;
  border-radius: 2px;
}

.news-list-nav {
  padding: 15px 30px;
  position: fixed;
  text-align: center;
  top: 55px;
  left: 0;
  right: 0;
  z-index: 998;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);

  a {
    margin: 0 1em;
    cursor: pointer;
  }

  .disabled {
    color: #ccc;
  }
}

.news-list {
  position: absolute;
  margin: 30px 0;
  width: 100%;
  transition: all 0.5s cubic-bezier(0.55, 0, 0.1, 1);

  ul {
    list-style-type: none;
    padding: 0;
    margin: 0;
  }
}

.slide-left-enter, .slide-right-leave-to {
  opacity: 0;
  transform: translate(30px, 0);
}

.slide-left-leave-to, .slide-right-enter {
  opacity: 0;
  transform: translate(-30px, 0);
}

.item-move, .item-enter-active, .item-leave-active {
  transition: all 0.5s cubic-bezier(0.55, 0, 0.1, 1);
}

.item-enter {
  opacity: 0;
  transform: translate(30px, 0);
}

.item-leave-active {
  position: absolute;
  opacity: 0;
  transform: translate(30px, 0);
}

@media (max-width: 600px) {
  .news-list {
    margin: 10px 0;
  }
}
</style>
