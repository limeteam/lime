# 图片上传组件（拖动排序、删除图片、预览图片）

> 功能拖动排序、删除图片、预览图片

## 使用方法
最简单方法使用v-model双向绑定你要绑定的值
```html
<template>
    <image-uploader v-model="yourData"/>
</template   
```
或者使用value值输入，事件绑定修改值
```html
<template>
    <image-uploader :value="yourData" @change="bindYourDataFunction"/>
</template   
```
```es6
<script>
    import ImageUploader from '@/components/ImageUploader'
    export default {
        // ...
        components: {
            ImageUploader
        },
        // ...
    }
</script> 
```

## 属性
*   value - 绑定值 类型:Array **必填** 
*   disabled - 是否禁用 类型:Boolean|String 默认:false
*   sortable - 是否可以拖拽 类型:Boolean|String 默认:true
*   sizeLimit - 文件大小限制 类型:Number 默认:5 备注：单位是mb
*   extLimit - 后缀名 类型:Array 默认:['pdf', 'jpg', 'jpeg', 'png']
*   isPrivate - 是否私有库 类型:Boolean 默认:false

## 组件依赖
element-ui
@/utils/upLoadFile(oss上传文件公共方法)
vue-js-grid(拖拽排序组件)