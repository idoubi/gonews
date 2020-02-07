<template>
  <div id="app">
    <header class="header">
      <nav class="inner">
        <router-link to="/" exact id="logo">
          <img class="logo" src="./assets/logo.png" alt="logo" />
        </router-link>
        <a href="https://idoubi.cc" target="_blank" class>作者博客</a>
        <a href="https://github.com/idoubi/gonews" target="_blank" class>项目地址</a>
        <a href="https://gocn.vip" target="_blank" class>GoCN社区</a>

        <div class="github">
          <input type="text" v-model="keyword" id="search" placeholder="keyword..." />
        </div>
      </nav>
    </header>
    <transition name="fade" mode="out-in">
      <router-view class="view"></router-view>
    </transition>
  </div>
</template>

<script>
export default {
  name: "app",
  data() {
    return {
      keyword: ""
    };
  },
  created() {
    this.$store.commit("SET_PAGE", Number(this.$route.params.page || 1));
    this.$store.commit("SET_NEWS");
  },
  watch: {
    $route() {
      // 分页查询
      let page = Number(this.$route.params.page) || 1;
      this.$store.commit("SET_PAGE", page);
      this.$store.commit("SET_NEWS");
    },
    keyword(value) {
      // 关键词搜索
      this.$store.commit("SET_PAGE", 1);
      this.$store.commit("SET_KEYWORD", value);
      this.$store.commit("SET_NEWS");
    }
  }
};
</script>

<style lang="stylus">
body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif;
  font-size: 15px;
  background-color: lighten(#eceef1, 30%);
  margin: 0;
  padding-top: 55px;
  color: #34495e;
  overflow-y: scroll;
}

a {
  color: #34495e;
  text-decoration: none;
}

.header {
  background-color: #499ef3;
  position: fixed;
  z-index: 999;
  height: 55px;
  top: 0;
  left: 0;
  right: 0;

  .inner {
    max-width: 800px;
    box-sizing: border-box;
    margin: 0px auto;
    padding: 15px 5px;
  }

  a {
    color: rgba(255, 255, 255, 0.8);
    line-height: 24px;
    transition: color 0.15s ease;
    display: inline-block;
    vertical-align: middle;
    font-weight: 300;
    letter-spacing: 0.075em;
    margin-right: 1.8em;

    &:hover {
      color: #fff;
    }

    &.router-link-active {
      color: #fff;
      font-weight: 400;
    }

    &:nth-child(6) {
      margin-right: 0;
    }
  }

  .github {
    color: #fff;
    font-size: 0.9em;
    margin: 0;
    float: right;
  }

  #search {
    border: 1px solid #ccc;
    padding: 8px 10px;
    color: #333;
    margin-top: -5px;
  }
}

.logo {
  position: absolute;
  top: 5px;
  width: 44px;
  height: 44px;
  margin-right: 10px;
  display: inline-block;
  vertical-align: middle;
}

#logo {
  margin-right: 4em;
}

.view {
  max-width: 800px;
  margin: 0 auto;
  position: relative;
}

.fade-enter-active, .fade-leave-active {
  transition: all 0.2s ease;
}

.fade-enter, .fade-leave-active {
  opacity: 0;
}

@media (max-width: 860px) {
  .header .inner {
    padding: 15px 30px;
  }
}

@media (max-width: 600px) {
  .header {
    .inner {
      padding: 15px;
    }

    a {
      margin-right: 1em;
    }

    .github {
      display: none;
    }
  }
}
</style>