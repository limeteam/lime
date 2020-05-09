<template>
  <div>{{ data }}</div>
</template>
<script>
import { getConfigList } from '@/api/lime-admin/config'
export default {
  name: 'Configlist',
  data() {
    return {
      data: ''
    }
  },
  created() {
    this.getConfigList()
  },
  methods: {
    async getConfigList() {
      getConfigList().then(response => {
        const data = response.data
        this.data = data
      })
    }
  }
}
</script>
