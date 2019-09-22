<template>
  <div style="margin: 0 15px;background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;min-height: 500px;">

    <Row style="padding: 50px;">
      <Col span="16">
        <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
          <FormItem label="商品名称" prop="good_name">
            <Input v-model="formValidate.good_name" placeholder="Enter blog good_name..."/>
          </FormItem>
          <FormItem label="商品描述" prop="good_desc">
            <Input v-model="formValidate.good_desc" type="textarea" :rows="5" placeholder="Enter blog good_desc..."/>
          </FormItem>
          <FormItem label="商品金额" prop="good_price">
            <Input v-model="formValidate.good_price" placeholder="Enter blog good_price..."/>
          </FormItem>

          <FormItem label="商品图片" prop="good_images">
            <Scroll height="160">
              <span style="height: 140px;">
                <img v-for="good_image in formValidate.good_images" :src="good_image"
                     width="120px" height="90px" style="margin: 5px;"/>
              </span>
            </Scroll>
            <IFileUpload ref="fileUpload" :auto-hide-modal="true" btn-size="small"
                         @uploadComplete="uploadComplete" action="/api/iwork/fileUpload/default" uploadLabel="上传图片"/>
          </FormItem>

          <FormItem>
            <Button type="success" @click="handleSubmit('formValidate')">提交</Button>
            <Button type="error" v-if="formValidate.article_id > 0"
                    style="margin-left: 8px" @click="handleDelete('formValidate')">删除该条目</Button>
          </FormItem>
        </Form>
      </Col>
      <Col span="8">
        AAAAAAAAAAAAAA
      </Col>
    </Row>
  </div>
</template>

<script>
  import IFileUpload from "../Common/file/IFileUpload";
  import {GoodEdit} from "../../api"

  export default {
    name: "GoodEdit",
    components:{IFileUpload},
    data(){
      return {
        formValidate: {
          good_id:-1,
          good_name: '',
          good_desc: '',
          good_price: 0,     // 负数表示暂无报价
          good_images:[],
        },
        ruleValidate: {
          good_name: [
            { required: true, message: '商品名称不能为空', trigger: 'blur' }
          ],
          good_desc: [
            { required: true, message: '商品描述不能为空', trigger: 'blur' }
          ],
          good_price: [
            { required: true, message: '商品价格不能为空', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      convertObjToArr(obj){
        if(typeof obj == typeof []){
          return obj;
        }
        let arr = [];
        for(key in obj){
          arr.push(obj[key]);
        }
        return arr;
      },
      uploadComplete(result, file) {
        if(result.status=="SUCCESS"){
          let _good_images = this.convertObjToArr(this.formValidate.good_images);
          _good_images.push(result.fileServerPath);
          this.$set(this.formValidate, "good_images", _good_images);
        }
      },
      handleDelete: async function(name){
        // if(this.formValidate.article_id > 0){
        //   const result = await ArticleDelete(this.formValidate.article_id);
        //   if(result.status == "SUCCESS"){
        //     this.$refs[name].resetFields();
        //     if(this.successEmit){
        //       this.$emit("successEmitFunc");
        //     }
        //   }
        // }
      },
      handleSubmit: function(name) {
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await GoodEdit(_this.formValidate.good_id,
              _this.formValidate.good_name,_this.formValidate.good_desc,
              _this.formValidate.good_price, JSON.stringify(_this.formValidate.good_images));
            if(result.status == "SUCCESS"){
                this.$router.push({ path: '/igood/good_list'});
            }else{
              _this.$Message.error('提交失败!');
            }
          } else {
            _this.$Message.error('验证失败!');
          }
        })
      },
    }
  }
</script>

<style scoped>

</style>
