import ICMSLayout from "../components/ILayout/ICMSLayout"
import Carousel from "../components/CMS/Carousel"
import Catalog from "../components/CMS/Catalog"
import Placement from "../components/CMS/Placement"

export const ICMSReouter = [{
  path: '/cms',
  component: ICMSLayout,
  children: [
    {path: 'carousel_list',component: Carousel},
    {path: 'catalog_list',component: Catalog},
    {path: 'placement_list',component: Placement},
  ]
}];
