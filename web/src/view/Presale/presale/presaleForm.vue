
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="预售项目名称:" prop="presaleName">
          <el-input v-model="formData.presaleName" :clearable="true"  placeholder="请输入预售项目名称" />
       </el-form-item>
        <el-form-item label="预售项目图片:" prop="presalePicture">
          <SelectImage v-model="formData.presalePicture" file-type="image"/>
       </el-form-item>
        <el-form-item label="预售开始时间:" prop="presaleStartTime">
          <el-date-picker v-model="formData.presaleStartTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="上传者用户ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="预售项目是否审核通过:" prop="presaleIsPass">
          <el-switch v-model="formData.presaleIsPass" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="预售项目是否展示:" prop="presaleIsShow">
          <el-switch v-model="formData.presaleIsShow" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="公链ID:" prop="publicChainId">
           <el-select v-model="formData.publicChainId" placeholder="请选择公链ID" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in PublicChainOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="预售项目链接:" prop="presaleurl">
          <el-input v-model="formData.presaleurl" :clearable="true"  placeholder="请输入预售项目链接" />
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
  createPresale,
  updatePresale,
  findPresale
} from '@/api/Presale/presale'

defineOptions({
    name: 'PresaleForm'
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
            presaleName: '',
            presalePicture: "",
            presaleStartTime: new Date(),
            userId: undefined,
            presaleIsPass: false,
            presaleIsShow: false,
            publicChainId: '',
            presaleurl: '',
        })
// 验证规则
const rule = reactive({
               presaleName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               presalePicture : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               presaleStartTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               presaleIsPass : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               presaleIsShow : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               publicChainId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               presaleurl : [{
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
      const res = await findPresale({ ID: route.query.id })
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
               res = await createPresale(formData.value)
               break
             case 'update':
               res = await updatePresale(formData.value)
               break
             default:
               res = await createPresale(formData.value)
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
