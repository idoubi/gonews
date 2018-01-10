<template>
  <li class="news-item">
    <span class="score" v-if="item.newDay">{{ item.weekday }}</span>
    <span class="title">
      <template v-if="item.url">
        <a :href="item.url" target="_blank" rel="noopener">{{ item.title ? item.title : item.url }}</a>
      </template>
    </span>
    <br>
    <span class="meta">
      <span class="time">
        {{ item.ctime }}
      </span>
    </span>
  </li>
</template>

<script>
export default {
  name: 'news-item',
  props: ['item']
}
</script>

<style lang="stylus">
.news-item
  background-color #fff
  padding 20px 30px 20px 80px
  border-bottom 1px solid #eee
  position relative
  line-height 20px
  .score
    color #ff6600
    font-size 1.1em
    font-weight 700
    position absolute
    top 50%
    left 0
    width 80px
    text-align center
    margin-top -10px
  .meta, .host
    font-size .85em
    color #828282
    a
      color #828282
      text-decoration underline
      &:hover
        color #ff6600
</style>