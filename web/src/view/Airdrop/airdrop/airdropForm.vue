
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="空投名称:" prop="airdropName">
          <el-input v-model="formData.airdropName" :clearable="true"  placeholder="请输入空投名称" />
       </el-form-item>
        <el-form-item label="空投项目图片:" prop="airdropPicture">
          <SelectImage v-model="formData.airdropPicture" file-type="image"/>
       </el-form-item>
        <el-form-item label="空投项目价值:" prop="airdropValue">
          <el-input-number v-model="formData.airdropValue" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="空投项目结束日期:" prop="airdropEndtime">
          <el-date-picker v-model="formData.airdropEndtime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="提交者用户id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="公链id:" prop="publicChainId">
           <el-select v-model="formData.publicChainId" placeholder="请选择公链id" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in PublicChainOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="是否通过审核:" prop="airdropIsPass">
          <el-switch v-model="formData.airdropIsPass" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="是否展示:" prop="airdropIsShow">
          <el-switch v-model="formData.airdropIsShow" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="空投项目网址:" prop="airdropUrl">
          <el-input v-model="formData.airdropUrl" :clearable="true"  placeholder="请输入空投项目网址" />
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
  createAirdrop,
  updateAirdrop,
  findAirdrop
} from '@/api/Airdrop/airdrop'

defineOptions({
    name: 'AirdropForm'
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
            airdropName: '',
            airdropPicture: "",
            airdropValue: 0,
            airdropEndtime: new Date(),
            userId: undefined,
            publicChainId: '',
            airdropIsPass: false,
            airdropIsShow: false,
            airdropUrl: '',
        })
// 验证规则
const rule = reactive({
               airdropName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               airdropPicture : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               airdropValue : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               airdropEndtime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               publicChainId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               airdropUrl : [{
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
      const res = await findAirdrop({ ID: route.query.id })
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
               res = await createAirdrop(formData.value)
               break
             case 'update':
               res = await updateAirdrop(formData.value)
               break
             default:
               res = await createAirdrop(formData.value)
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
