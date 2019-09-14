<template>
  <Row>
    <Col span="6">
      图书名称：<Input v-model="book_name"/>
      <Button @click="addChapter">新增目录</Button>
      <Input v-for="(chapter, index) in book_chapters" type="text" style="margin-bottom: 5px;" :value="chapter.chapter_name"/>
      <Button size="small" type="success" @click="editBook">提交</Button>
    </Col>
    <Col span="18">
    </Col>
  </Row>
</template>

<script>
  import {EditBook} from "../../../api"
  import BlogEdit from "../BlogEdit";

  export default {
    name: "BookEdit",
    components: {BlogEdit},
    data(){
      return {
        book_name:'',
        book_chapters:[],
      }
    },
    methods:{
      addChapter:function () {
        this.book_chapters.push({'chapter_name':'测试目录'});
      },
      editBook:async function () {
        const result = await EditBook(-1, this.book_name, JSON.stringify(this.book_chapters));
        if(result.status == "SUCCESS"){
          alert(11111);
        }
      }
    }
  }
</script>

<style scoped>

</style>
