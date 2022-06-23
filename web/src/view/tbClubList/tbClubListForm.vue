<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="俱乐部名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创始人id:">
          <el-input v-model.number="formData.ownerId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="俱乐部公告:">
          <el-input v-model="formData.clubMsg" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="累计服务费:">
          <el-input-number v-model="formData.feeSum" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="累计分成:">
          <el-input-number v-model="formData.cutSum" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'TbClubList'
}
</script>

<script setup>
import {
  createTbClubList,
  updateTbClubList,
  findTbClubList
} from '@/api/tbClubList'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        name: '',
        ownerId: 0,
        clubMsg: '',
        feeSum: 0,
        cutSum: 0,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTbClubList({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retbClubList
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createTbClubList(formData.value)
          break
        case 'update':
          res = await updateTbClubList(formData.value)
          break
        default:
          res = await createTbClubList(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
