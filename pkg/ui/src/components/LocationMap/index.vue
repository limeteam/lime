<template>
  <baidu-map
    v-bind="$attrs"
    :scroll-wheel-zoom="!disabled"
    :dragging="!disabled"
    :zoom="15"
    :center="position"
    ak="EEbiVGRFiZnERXVsgzRdQa2rVVXLRNOq"
    class="baidu-map"
    @ready="onMapReady"
  >
    <bm-navigation anchor="BMAP_ANCHOR_TOP_RIGHT" />
    <bm-marker
      v-if="pointText"
      :position="position"
      :dragging="false"
      animation="BMAP_ANIMATION_BOUNCE"
    >
      <bm-label
        :label-style="{color: 'red', fontSize : '12px'}"
        :offset="{ height: 30}"
        :content="pointText"
      />
    </bm-marker>
  </baidu-map>
</template>

<script>
import Vue from 'vue'
import BaiduMap from 'vue-baidu-map'

Vue.use(BaiduMap, {
  ak: 'EEbiVGRFiZnERXVsgzRdQa2rVVXLRNOq'
})
export default {
  name: 'LocationMap',
  inheritAttrs: false,
  props: {
    disabled: {
      type: Boolean,
      default: false
    },
    lng: {
      type: [Number, String],
      required: true,
      default: 0
    },
    lat: {
      type: [Number, String],
      required: true,
      default: 0
    },
    pointText: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      status: {
        mapReady: false
      }
    }
  },
  computed: {
    position() {
      if (this.status.mapReady) {
        return { lng: parseFloat(this.lng) || 0, lat: parseFloat(this.lat) || 0 }
      } else {
        return { lng: 0, lat: 0 }
      }
    }
  },
  methods: {
    onMapReady() {
      this.status.mapReady = true
    }
  }
}
</script>

<style lang="scss">
.baidu-map {
    width: 100%;
    height: 400px;
    box-shadow: 0px 0px 2px 1px rgba(0,0,0,.175);
}
</style>
