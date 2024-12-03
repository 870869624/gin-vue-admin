
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="公链名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入公链名称" />
       </el-form-item>
        <el-form-item label="公链图片:" prop="picture">
          <SelectImage v-model="formData.picture" file-type="image"/>
       </el-form-item>
        <el-form-item label="公链logo:" prop="logo">
          <SelectImage v-model="formData.logo" file-type="image"/>
       </el-form-item>
        <el-form-item label="公链简介:" prop="introduction">
          <el-input v-model="formData.introduction" :clearable="true"  placeholder="请输入公链简介" />
       </el-form-item>
        <el-form-item label="公链网址:" prop="url">
          <el-input v-model="formData.url" :clearable="true"  placeholder="请输入公链网址" />
       </el-form-item>
        <el-form-item label="是否展示:" prop="is_Show">
          <el-switch v-model="formData.is_Show" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createPublicChain,
  updatePublicChain,
  findPublicChain
} from '@/api/PublicChain/publicChain'

defineOptions({
    name: 'PublicChainForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            name: '',
            picture: "",
            logo: "",
            introduction: '',
            url: '',
            is_Show: false,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               picture : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               introduction : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               url : [{
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
      const res = await findPublicChain({ ID: route.query.id })
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
               res = await createPublicChain(formData.value)
               break
             case 'update':
               res = await updatePublicChain(formData.value)
               break
             default:
               res = await createPublicChain(formData.value)
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
