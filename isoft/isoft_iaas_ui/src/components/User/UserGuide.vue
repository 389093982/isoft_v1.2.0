<template>
  <div>
    <ElementsLoader :placement_name="GLOBAL.placement_user_guide" @onLoadElement="onLoadElement">
      <IBeautifulCard :title="placement_label" v-if="elements.length > 0">
        <div slot="content">
          <Row style="min-height: 500px;padding: 20px;">
            <Col span="6">
              <div style="min-height: 500px;border: 2px solid rgba(223,223,223,0.5);padding: 20px;margin: 2px;">
                <p v-for="(element, index) in elements" style="font-size: 15px;">
                  <IBeautifulLink @onclick="showElement(index)">{{element.element_label}}</IBeautifulLink>
                </p>
              </div>
            </Col>
            <Col span="18">
              <div style="min-height: 500px;border:2px solid rgba(223,223,223,0.5);padding: 20px;margin: 2px;">
                <IShowMarkdown v-if="elements[showElementIndex] && elements[showElementIndex].md_content" :content="elements[showElementIndex].md_content"/>
              </div>
            </Col>
          </Row>
        </div>
      </IBeautifulCard>
    </ElementsLoader>
  </div>
</template>

<script>
  import IShowMarkdown from "../Common/markdown/IShowMarkdown"
  import ElementsLoader from "../Background/CMS/ElementsLoader";
  import IBeautifulCard from "../../components/Common/card/IBeautifulCard";
  export default {
    name: "UserGuide",
    components: {IShowMarkdown, IBeautifulCard, ElementsLoader},
    props:{
      placement_name:{
        type:String,
        default: '',
      }
    },
    data(){
      return {
        elements:[],
        placement_label:'',
        showElementIndex: 0,
      }
    },
    methods:{
      showElement:function(index){
        this.showElementIndex = index;
      },
      onLoadElement:function (placement_label, elements) {
        this.placement_label = placement_label;
        this.elements = elements;
      }
    }
  }
</script>

<style scoped>

</style>
