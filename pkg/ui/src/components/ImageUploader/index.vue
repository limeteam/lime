<template>
  <div class="image-uploader">
    <div class="image-uploader__scrollbar">
      <el-upload
        v-if="!_disabled"
        :show-file-list="false"
        :http-request="uploadImage"
        action
        class="image-uploader__el-upload"
      >
        <i class="el-icon-plus avatar-uploader-icon" />
      </el-upload>
      <div class="image-uploader__image-container">
        <div v-if="value.length === 0" class="image-uploader__contain--no-data">
          <p>暂无数据</p>
        </div>
        <grid
          ref="grid"
          :style="_styleGrid"
          :draggable="!_disabled"
          :sortable="_sortable"
          :items="imagePreviewList"
          :cell-width="imgSize"
          :cell-height="imgSize"
          @dragend="onDragend"
        >
          <template slot="cell" slot-scope="props">
            <div class="image-uploader__image-wrapper">
              <el-badge v-if="props.sort === 0" value="封面" class="image-uploader__mark" />
              <el-button
                class="image-uploader__preview-button"
                size="mini"
                @click="toPreview($event,getGridItemBySlot(props.sort))"
              >
                预览
              </el-button>
              <i
                v-if="!_disabled"
                class="el-icon-circle-close image-uploader__close-icon"
                @click="onDeleteImage(props)"
                @mousedown="onMouseDownDelete($event)"
              />
              <img
                :title="props.item.url"
                :alt="props.item.url"
                :src="getImageSrc(getGridItemBySlot(props.sort))"
                class="image-uploader__image"
                @mousedown="onMouseDownImg($event)"
              >
            </div>
          </template>
        </grid>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
import Grid from 'vue-js-grid'
import { uploadFile } from '@/utils/upLoadFile'
import { Loading } from 'element-ui'

Vue.use(Grid)

export default {
  name: 'ImageUploader',
  model: {
    prop: 'value',
    event: 'change'
  },
  props: {
    value: {
      type: Array,
      required: true
    },
    disabled: {
      type: [Boolean, String],
      default: false
    },
    sortable: {
      type: [Boolean, String],
      default: true
    },
    // 上传大小限制单位mb
    sizeLimit: {
      type: Number,
      default: 5
    },
    isPrivate: {
      type: Boolean,
      default: false
    },
    extLimit: {
      type: Array,
      default: () => ['pdf', 'jpg', 'jpeg', 'png']
    }
  },
  data() {
    return {
      // 是否来自父组件改变
      isPropChange: true,
      hasDeletedItem: false,
      imagePreviewList: [],
      styleGrid: {
        height: '0'
      }
    }
  },
  computed: {
    imgSize: () => 188,
    _disabled() {
      return this.disabled || this.disabled === ''
    },
    _sortable() {
      if (this._disabled) {
        return false
      } else {
        return this.sortable || this.sortable === ''
      }
    },
    _styleGrid() {
      const isShow = this.value.length !== 0
      if (isShow) {
        return { visibility: 'visible', height: this.styleGrid.height }
      } else {
        return { visibility: 'hidden', height: '0' }
      }
    }
  },
  watch: {
    value: 'watcherValue'
  },
  mounted() {
    this.init()
    this.getWindowSize()
    window.addEventListener('resize', this.getWindowSize)
  },
  updated() {
    this.getWindowSize()
  },
  methods: {
    init() {
      this.watcherValue(this.value)
    },
    watcherValue(val, oldVal) {
      if (this.isPropChange && val.length > 0) {
        // 根据value插入显示数据
        this.imagePreviewList.push(...val)
      }
    },
    /**
     * updateData 根据Grid组件的Data，更新参数value
     */
    updateData(gridArr) {
      this.isPropChange = false
      if (gridArr) {
        const deletedData = []

        if (this.hasDeletedItem) {
          const delIndex = this.value.findIndex(item => item.hasOwnProperty('id') && item.del)
          // 当找到有删除值时候分离，正常值和删除值
          if (delIndex > -1) {
            deletedData.push(...this.value.splice(delIndex, this.value.length - delIndex))
          }
        }
        // 把gird的data.list复制一份
        gridArr = gridArr.slice().sort((a, b) => a.sort - b.sort)
        gridArr.forEach((gridArrItem, index) => {
          if (this.value[index]) {
            this.value[index] = gridArrItem.item
            this.value[index].sort = gridArrItem.sort
          } else {
            this.value.push(gridArrItem.item)
          }
        })

        if (this.hasDeletedItem) {
          this.value.push(...deletedData)
        }

        this.$emit('change', this.value)
      }
    },
    deleteData({ item }) {
      const valueItem = this.value.find(_item => _item === item)
      if (valueItem != null) {
        this.isPropChange = false
        this.value.splice(this.value.indexOf(valueItem), 1)
        if (valueItem.hasOwnProperty('id')) {
          valueItem.del = true
          delete valueItem.sort
          this.value.push(valueItem)
        }
        if (this.$refs.grid.list.length) {
          // 整理顺序
          this.$refs.grid.list.forEach((gridArrItem, index) => {
            // 输出数据顺序
            this.value[index].sort = gridArrItem.sort
          })
        }

        this.$emit('change', this.value)
      }
    },
    getWindowSize() {
      // 等待vue-dom ready
      this.$nextTick(() => {
        if (this.$refs.grid && this.$refs.grid.windowWidth != null) {
          // 当前第三方组件的宽度
          const containerWidth = this.$refs.grid.$el.offsetWidth
          // 计算当前组件每一行最多可以放多少图片
          const rowPictureCount = Math.floor(containerWidth / this.imgSize)
          // 第三方组件grid的图片list
          const list = this.$refs.grid.list

          // 动态计算拖到grid的宽度（int-组件规定数据类型）（每一行最多可放图片数 * 每张图片宽度）
          this.$refs.grid.windowWidth = rowPictureCount * this.imgSize
          // 动态计算拖到grid的高度（string-需要添加px才生效）（ceil(当前图片数量 / 每一行最多可放图片数)  * 每张图片宽度）
          this.styleGrid.height = Math.ceil(list.length / rowPictureCount) * this.imgSize + 'px'
        }
      })
    },
    async uploadImage(context) {
      if (!this.checkFile({ file: context.file, handler: this.fileLimitHandler })) {
        return
      }
      if (!this.checkFile({ file: context.file, handler: this.fileExtLimitHandler })) {
        return
      }
      const loadingInstance = Loading.service({
        lock: true,
        background: 'transparent'
      })
      try {
        const upLoadPreviewUrl = await this.getPreviewImgSrc(context.file)
        const uploadRsp = await uploadFile({
          file: context.file,
          name: context.file.fileName,
          // 1 公有库 2 私有库
          type: this.isPrivate ? 2 : 1
        }, {
          beforeGetAuth: loadingInstance.setText.bind(loadingInstance, '正在获取上传权限'),
          beforeUpload: loadingInstance.setText.bind(loadingInstance, '正在上传文件...')
        })

        if (uploadRsp) {
          const grid = this.$refs.grid
          const length = grid.list.length
          // 这里需要通过访问子组件插入数据，不能通过imagePreviewList插入或删除数据，因为会导致子组件的位置复位bug
          let uploadUrl = uploadRsp.res.requestUrls[0]
          const index = uploadUrl.indexOf('?')
          if (index !== -1) {
            uploadUrl = uploadUrl.slice(0, index)
          }

          grid.list.push({
            // 保证index唯一性，防止出现index相同情况导致拖到触发多个子组件bug
            index: Math.floor(Math.random() * 1000000),
            sort: length,
            item: {
              url: uploadUrl, sort: length
            },
            // 上传预览图片地址
            upLoadPreviewUrl
          })
          if (this.$refs.grid) {
            // 插入新图片
            this.updateData(this.$refs.grid.list)
          }
        }
      } finally {
        this.$nextTick(() => {
          // 任何情况下，隐藏loading
          loadingInstance.close()
        })
      }
    },
    getPreviewImgSrc(file) {
      const reader = new FileReader()
      const requestImgSrc = new Promise((resolve, reject) => {
        reader.addEventListener('load', () => resolve(reader.result), false)
        reader.addEventListener('error', error => reject(error), false)
      })
      const timeout = new Promise((resolve, reject) => setTimeout(reject.bind(null, new Error('读取本地图片失败')), 5000))
      reader.readAsDataURL(file)
      return Promise.race([requestImgSrc, timeout])
    },
    getGridItemBySlot(sort) {
      return this.$refs.grid.list.find(listItem => listItem.item.sort === sort)
    },
    getImageSrc(gridItem) {
      if (gridItem.upLoadPreviewUrl) {
        return gridItem.upLoadPreviewUrl
        // 私有域图片预览地址 previewUrl
      } else if (gridItem.item.previewUrl) {
        return gridItem.item.previewUrl
      } else {
        return gridItem.item.url
      }
    },
    checkFile({ file, handler }) {
      const { valid, message } = handler(file)
      if (!valid) {
        this.$message({ message, type: 'error' })
      }
      return valid
    },
    fileLimitHandler(file) {
      // 换算为字节数
      const fileSizeLimit = this.sizeLimit * 1000000
      const valid = file.size <= fileSizeLimit
      return { valid, message: `上传数据不能超过${this.sizeLimit.toFixed(2)}M` }
    },
    fileExtLimitHandler(file) {
      const point = file.name.lastIndexOf('.')
      const ext = file.name.substr(point + 1)
      const valid = this.extLimit.indexOf(ext) > -1
      return { valid, message: `上传的文件仅支持${this.extLimit.join(',')}等${this.extLimit.length}种格式` }
    },
    onDragend(event) {
      const item = event.items
      this.updateData(item)
    },
    onDeleteImage(slotProp) {
      this.hasDeletedItem = true
      // 删除第三方组件的数据
      slotProp.remove()
      this.deleteData(slotProp)
    },
    onMouseDownDelete(event) {
      event.stopPropagation()
    },
    onMouseDownImg(event) {
      event.preventDefault()
    },
    toPreview(event, gridItem) {
      event.stopPropagation()
      const previewUrl = gridItem.upLoadPreviewUrl || gridItem.item.previewUrl

      if (previewUrl) {
        // 在新的标签页预览图片
        const image = new Image()
        const newWindow = window.open('', Date.now())
        image.src = previewUrl
        newWindow.document.write('')
        newWindow.document.write(image.outerHTML)
      } else {
        window.open(gridItem.item.url)
      }
    }
  }
}
</script>

<style lang="scss">
$imgSize: 178px;

.image-uploader {
  min-height: $imgSize;

  &__scrollbar {
    display: flex;
  }

  &__image-container {
    flex: 1 0 auto;
  }
  &__el-upload {
    width: $imgSize;
    height: $imgSize;
    flex: 0 0 $imgSize;
    margin: 10px;
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  &__el-upload:hover {
    border-color: #409eff;
  }
  &__close-icon {
    position: absolute;
    right: 0;
    top: 0;
    width: 20px;
    height: 20px;
    cursor: pointer;

    &::before {
      font-size: 20px;
      color: red;
    }
  }
  &__image {
    width: 100%;
    height: 100%;
    overflow: hidden;
    display: block;
  }
  &__image-wrapper {
    margin: 10px;
    width: $imgSize;
    height: $imgSize;
    box-sizing: border-box;
    padding: 8px;
    background-color: #fff;
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    overflow: hidden;
    position: relative;
  }
  &__contain--no-data {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;

    p {
      color: #ccc;
      font-size: 16px;
      text-align: center;
    }
  }
  &__mark {
    position: absolute;
    left: 0;
    top: 5px;
    .el-badge__content {
      background-color: blue;
    }
  }
  &__preview-button {
    position: absolute;
    right: 0;
    bottom: 0;
  }
  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: $imgSize;
    height: $imgSize;
    line-height: $imgSize;
    text-align: center;
  }
  .avatar {
    width: $imgSize;
    height: $imgSize;
    display: block;
  }
}
</style>
