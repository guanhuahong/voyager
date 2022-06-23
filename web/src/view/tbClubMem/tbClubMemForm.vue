<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="俱乐部id:">
          <el-input v-model.number="formData.clubId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="成员id:">
          <el-input v-model.number="formData.playerId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="成员类型:">
          <el-select v-model="formData.memType" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in club_member_typeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="上级合伙人:">
          <el-input v-model.number="formData.parentId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分成比例:">
          <el-input-number v-model="formData.cutRate" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="邀请码:">
          <el-input v-model="formData.inviteCode" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合伙人公告:">
          <el-input v-model="formData.partnerMsg" clearable placeholder="请输入" />
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
  name: 'TbClubMem'
}
</script>

<script setup>
import {
  createTbClubMem,
  updateTbClubMem,
  findTbClubMem
} from '@/api/tbClubMem'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const club_member_typeOptions = ref([])
const formData = ref({
        clubId: 0,
        playerId: 0,
        memType: undefined,
        parentId: 0,
        cutRate: 0,
        inviteCode: '',
        partnerMsg: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTbClubMem({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retbClubMem
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    club_member_typeOptions.value = await getDictFunc('club_member_type')
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createTbClubMem(formData.value)
          break
        case 'update':
          res = await updateTbClubMem(formData.value)
          break
        default:
          res = await createTbClubMem(formData.value)
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
