
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="充值者ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否有效:" prop="isEffective">
          <el-switch v-model="formData.isEffective" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="交易hash:" prop="blockHash">
          <el-input v-model="formData.blockHash" :clearable="true"  placeholder="请输入交易hash" />
       </el-form-item>
        <el-form-item label="交易网络数:" prop="blockNumber">
          <el-input v-model="formData.blockNumber" :clearable="true"  placeholder="请输入交易网络数" />
       </el-form-item>
       <el-form-item label="BAB令牌:" prop="bab">
          <el-input v-model="formData.bab" :clearable="true"  placeholder="BAB令牌" />
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
  createVipRecord,
  updateVipRecord,
  findVipRecord
} from '@/api/Vip/vip_record'

defineOptions({
    name: 'VipRecordForm'
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
            userId: undefined,
            isEffective: false,
            blockHash: '',
            blockNumber: '',
            bab: '',
        })
// 验证规则
const rule = reactive({
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               isEffective : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               blockHash : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               blockNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               bab : [{
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
      const res = await findVipRecord({ ID: route.query.id })
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
               res = await createVipRecord(formData.value)
               break
             case 'update':
               res = await updateVipRecord(formData.value)
               break
             default:
               res = await createVipRecord(formData.value)
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
