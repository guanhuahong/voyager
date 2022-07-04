<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="玩家id:">
          <el-input v-model.number="formData.playerId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="玩家昵称:">
          <el-input v-model="formData.nick" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头像:">
          <el-input v-model="formData.avatar" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="玩家账户金额:">
          <el-input-number v-model="formData.gold" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="玩家状态:">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="累计输赢:">
          <el-input-number v-model="formData.totalWinlost" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="累计抽水:">
          <el-input-number v-model="formData.totalSysfee" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="累计投注:">
          <el-input-number v-model="formData.totalBet" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="累计有效投注:">
          <el-input-number v-model="formData.totalEffbet" :precision="2" clearable></el-input-number>
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
  name: 'PlayerProfile'
}
</script>

<script setup>
import {
  createPlayerProfile,
  updatePlayerProfile,
  findPlayerProfile
} from '@/api/playerProfile'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        playerId: 0,
        nick: '',
        avatar: '',
        gold: 0,
        status: false,
        totalWinlost: 0,
        totalSysfee: 0,
        totalBet: 0,
        totalEffbet: 0,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPlayerProfile({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.replayerProfile
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
          res = await createPlayerProfile(formData.value)
          break
        case 'update':
          res = await updatePlayerProfile(formData.value)
          break
        default:
          res = await createPlayerProfile(formData.value)
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
