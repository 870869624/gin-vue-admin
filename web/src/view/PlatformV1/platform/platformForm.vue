
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="平台名称:" prop="platformName">
          <el-input v-model="formData.platformName" :clearable="true"  placeholder="请输入平台名称" />
       </el-form-item>
        <el-form-item label="平台图片:" prop="platformPicture">
          <SelectImage v-model="formData.platformPicture" file-type="image"/>
       </el-form-item>
        <el-form-item label="平台是否展示:" prop="platformIsShow">
          <el-switch v-model="formData.platformIsShow" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="公链ID:" prop="publicChainId">
           <el-select v-model="formData.publicChainId" placeholder="请选择公链ID" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in PublicChainOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="平台链接:" prop="platformUrl">
          <el-input v-model="formData.platformUrl" :clearable="true"  placeholder="请输入平台链接" />
       </el-form-item>
        <el-form-item label="平台描述:" prop="brief">
          <el-input v-model="formData.brief" :clearable="true"  placeholder="请输入平台描述" />
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
  createPlatform,
  updatePlatform,
  findPlatform
} from '@/api/PlatformV1/platform'

defineOptions({
    name: 'PlatformForm'
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
const PublicChainOptions = ref([])
const formData = ref({
            platformName: '',
            platformPicture: "",
            platformIsShow: false,
            publicChainId: '',
            platformUrl: '',
            brief: '',
        })
// 验证规则
const rule = reactive({
               platformName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               platformPicture : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               publicChainId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               platformUrl : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               brief : [{
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
      const res = await findPlatform({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    PublicChainOptions.value = await getDictFunc('PublicChain')
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
               res = await createPlatform(formData.value)
               break
             case 'update':
               res = await updatePlatform(formData.value)
               break
             default:
               res = await createPlatform(formData.value)
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
