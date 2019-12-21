<template>
  <!-- 图文链接类型占位符 -->
  <div>
    <ElementsLoader :placement_name="placement_name" @onLoadElement="onLoadElement">
      <ul v-if="elements.length > 0">
        <li v-for="element in elements" class="li">
          <a :href="element.linked_refer" target="_blank">
            <img :src="element.img_path" width="280px" height="150px"/>
            <p style="font-size: 14px;">{{element.element_label}}</p>
          </a>
        </li>
      </ul>
    </ElementsLoader>
  </div>
</template>

<script>
  import ElementsLoader from "../../Background/CMS/ElementsLoader"

  export default {
    name: "GraphicLink",
    components:{ElementsLoader},
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
      }
    },
    methods:{
      onLoadElement:function (placement_label, elements) {
        this.placement_label = placement_label;
        this.elements = elements;
      }
    },
  }
</script>

<style scoped>
  a{
    color: black;
  }
  .li {
    list-style: none;
    float: left;
    padding: 5px 5px;
    text-align: center;
    border: 1px solid transparent;
    cursor: pointer;
  }
  li:hover{
    background-color: #f4f4f4;
    border: 1px solid #b2de8c;
  }
  li:hover a{
    color:red;
  }
</style>
