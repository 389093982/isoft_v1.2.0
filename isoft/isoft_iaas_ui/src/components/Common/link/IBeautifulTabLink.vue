<template>
  <IBeautifulCard :title="pData.title">
    <div slot="content" style="padding: 20px;">
      <Tabs :animated="false">
        <TabPane :label="itemData.itemTitle" v-for="(itemData, index) in pData.pItemData">
          <Row :gutter="10">
            <Col span="6" v-for="item in itemData.items">
              <IBeautifulHoverBox :item="item"/>
            </Col>
          </Row>
        </TabPane>
      </Tabs>
    </div>
  </IBeautifulCard>
</template>

<script>
  import IBeautifulHoverBox from "./IBeautifulHoverBox"
  import IBeautifulCard from "../../Common/card/IBeautifulCard"

  export default {
    name: "IBeautifulTabLink",
    components:{IBeautifulHoverBox,IBeautifulCard},
    data(){
      return {
        current_index : 0,
      }
    },
    props: {
      // 带有默认值的对象
      pData: {
        type: Object,
        // 对象或数组默认值必须从一个工厂函数获取
        default: function () {
          return {
            title: '热门产品',
            pItemData:[
              {
                itemTitle:'热门',
                items:[
                  {item_label:'云服务器1',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器1',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器1',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器1',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器1',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器1',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器1',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                ],
              },
              {
                itemTitle:'计算',
                items:[
                  {item_label:'云服务器2',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器2',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器2',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器2',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器2',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器2',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器2',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                ],
              },
              {
                itemTitle:'存储',
                items:[
                  {item_label:'云服务器3',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器3',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器3',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器3',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器3',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器3',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器3',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                ],
              },
              {
                itemTitle:'视频服务',
                items:[
                  {item_label:'云服务器4',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器4',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器4',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器4',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器4',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器4',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器4',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                ],
              },
              {
                itemTitle:'数据库',
                items:[
                  {item_label:'云服务器5',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器5',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器5',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器5',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器5',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器5',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                  {item_label:'云服务器5',item_desc:'稳定安全，高易用可弹性伸缩的计算服务',item_link:'https://www.baidu.com',},
                ],
              }
            ]
          }
        }
      },
    },
    methods:{

    },
  }
</script>

<style scoped>


</style>
