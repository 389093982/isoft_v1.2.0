<template>
  <div style="margin: 0 15px;background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;min-height: 500px;">

    <Row style="padding: 50px;">
      <Col span="16">
        <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="100">
          <FormItem label="商品名称" prop="good_name">
            <Input v-model.trim="formValidate.good_name" placeholder="Enter good_name..."/>
          </FormItem>
          <FormItem label="商品描述" prop="good_desc">
            <Input v-model.trim="formValidate.good_desc" type="textarea" :rows="5" placeholder="Enter good_desc..."/>
          </FormItem>
          <FormItem label="商品金额" prop="good_price">
            <Input v-model.trim="formValidate.good_price" placeholder="Enter good_price..."/>
          </FormItem>
          <FormItem label="卖家姓名" prop="good_seller">
            <Input v-model.trim="formValidate.good_seller" :readonly="true"/>
          </FormItem>
          <FormItem label="卖家联系方式" prop="seller_contact">
            <Input v-model.trim="formValidate.seller_contact" placeholder="Enter seller_contact..."/>
          </FormItem>

          <FormItem label="商品图片" prop="good_images">
            <Scroll height="160">
              <span style="height: 140px;">
                <img v-for="good_image in formValidate.good_images" :src="good_image"
                     width="120px" height="90px" style="margin: 5px;"/>
              </span>
            </Scroll>
            <IFileUpload ref="fileUpload" :auto-hide-modal="true" btn-size="small"
                         @uploadComplete="uploadComplete" action="/api/iwork/httpservice/fileUpload" uploadLabel="上传图片"/>
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
  import {GoodEdit,GetGoodDetail} from "../../api"
  import {GetLoginUserName} from "../../tools"

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
          good_seller:'',
          seller_contact:'',
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
          seller_contact: [
            { required: true, message: '卖家联系方式不能为空', trigger: 'blur' }
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
        if(this.formValidate.good_desc.length < 50){
          this.$Message.error('商品描述太短，不能少于 50 个字符!');
          return;
        }
        if(this.formValidate.good_images.length == 0){
          this.$Message.error('必须上传一张图片!');
          return;
        }
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await GoodEdit(_this.formValidate.good_id,
              _this.formValidate.good_name,_this.formValidate.good_desc,
              _this.formValidate.good_price, _this.formValidate.good_seller,
              _this.formValidate.seller_contact, JSON.stringify(_this.formValidate.good_images));
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
      refreshGoodDetail:async function (good_id) {
        const result = await GetGoodDetail(good_id);
        if(result.status == "SUCCESS"){
          this.formValidate.good_id = result.good.id;
          this.formValidate.good_name = result.good.good_name;
          this.formValidate.good_desc = result.good.good_desc;
          this.formValidate.good_desc = result.good.good_desc;
          this.formValidate.good_price = result.good.good_price;
          this.formValidate.good_seller = result.good.good_seller;
          this.formValidate.seller_contact = result.good.seller_contact;
          this.formValidate.good_images = JSON.parse(result.good.good_images);
        }
      }
    },
    mounted(){
      this.formValidate.good_seller = GetLoginUserName();
      if(this.$route.query.id != undefined && this.$route.query.id > 0){
        this.refreshGoodDetail(this.$route.query.id);
      }
    }
  }
</script>

<style scoped>

</style>
