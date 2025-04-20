
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户id:" prop="user_id">
          <el-input v-model.number="formData.user_id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邀请者用户名:" prop="username">
          <el-input v-model="formData.username" :clearable="true"  placeholder="请输入邀请者用户名" />
       </el-form-item>
        <el-form-item label="被邀请者用户id:" prop="newUserId">
          <el-input v-model.number="formData.newUserId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="被邀请者用户名:" prop="newUsername">
          <el-input v-model="formData.newUsername" :clearable="true"  placeholder="请输入被邀请者用户名" />
       </el-form-item>
        <el-form-item label="邀请码:" prop="inviteCode">
          <el-input v-model="formData.inviteCode" :clearable="true"  placeholder="请输入邀请码" />
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createInvitationRecord,
  updateInvitationRecord,
  findInvitationRecord
} from '@/api/InvitationRecord/invitationRecord'

defineOptions({
    name: 'InvitationRecordForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            user_id: undefined,
            username: '',
            newUserId: undefined,
            newUsername: '',
            inviteCode: '',
        })
// 验证规则
const rule = reactive({
               user_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               username : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               newUserId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               newUsername : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findInvitationRecord({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createInvitationRecord(formData.value)
               break
             case 'update':
               res = await updateInvitationRecord(formData.value)
               break
             default:
               res = await createInvitationRecord(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
