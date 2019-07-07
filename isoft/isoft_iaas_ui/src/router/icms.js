import ICMSLayout from "../components/ILayout/ICMSLayout"
import Element from "../components/CMS/Element"
import Catalog from "../components/CMS/Catalog"
import Placement from "../components/CMS/Placement"

export const ICMSReouter = [{
  path: '/cms',
  component: ICMSLayout,
  children: [
    {path: 'element_list',component: Element},
    {path: 'catalog_list',component: Catalog},
    {path: 'placement_list',component: Placement},
  ]
}];
