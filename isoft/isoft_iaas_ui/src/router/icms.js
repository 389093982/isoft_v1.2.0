import IEmptyLayout from "../components/ILayout/IEmptyLayout"
import Index from "../components/CMS/Index"

export const ICMSReouter = [{
  path: '/cms',
  component: IEmptyLayout,
  children: [
    {path: 'index',component: Index},
  ]
}];
